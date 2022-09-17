package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/xruins/NatureRemoWebUI/nature-remo-api-server/pkg/logger"
	remoclient "github.com/xruins/NatureRemoWebUI/nature-remo-api-server/pkg/remo/client"
	"github.com/xruins/NatureRemoWebUI/nature-remo-api-server/pkg/remo/client/operations"
	"github.com/xruins/NatureRemoWebUI/nature-remo-api-server/pkg/remo/models"
)

type Server struct {
	decoder                  *schema.Decoder
	logger                   logger.Logger
	Token                    string
	RemoClient               *remoclient.NatureAPI
	appliancesCache          []*models.Appliance
	appliancesCacheExpiredAt time.Time
	mu                       sync.RWMutex
	prefix                   string
}

var (
	remoAPITimeout              = time.Duration(30 * time.Second)
	remoAPIApplianceCacheExpire = time.Duration(30 * time.Minute)
)

type RemoAPIStatus struct {
	OK          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	ErrorDetail string `json:"error_detail"`
}

func returnError(w http.ResponseWriter, err error) {
	response := &RemoAPIStatus{}

	if err != nil {
		response.OK = false
		response.ErrorDetail = err.Error()
	}

	if apiErr, ok := err.(*runtime.APIError); ok {
		response.ErrorCode = apiErr.Code
	}

	b, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if !response.OK {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(b)
}

func NewServer(logger logger.Logger, token, prefix string) *Server {
	client := remoclient.NewHTTPClient(strfmt.Default)
	return &Server{
		decoder:    schema.NewDecoder(),
		logger:     logger,
		Token:      token,
		RemoClient: client,
		prefix:     prefix,
		mu:         sync.RWMutex{},
	}
}

func (s *Server) setAppliancesCache(ctx context.Context, now time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	p := &operations.Get1AppliancesParams{}
	p = p.WithTimeout(remoAPITimeout)
	p = p.WithContext(ctx)
	response, err := s.RemoClient.Operations.Get1Appliances(p, s.bearerTokenAuthInfo())
	if err != nil {
		s.logger.Errorf("failed to get Appliances. err: %s", err)
		return fmt.Errorf("failed to update appliances cache: %w", err)
	}

	s.appliancesCache = response.Payload
	s.appliancesCacheExpiredAt = now.Add(remoAPIApplianceCacheExpire)
	return nil
}

func (s *Server) validAppliancesCache(now time.Time) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.appliancesCache != nil && now.Before(s.appliancesCacheExpiredAt)
}

type getAppliancesHandlerParam struct {
	Force string `schema:"force"`
}

func (s *Server) getAppliancesHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	param := &getAppliancesHandlerParam{}
	err := s.decoder.Decode(param, r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var force bool
	if param.Force == "true" || param.Force == "1" {
		force = true
	}

	appendCORSHeaders(w)

	now := time.Now()
	if force || !s.validAppliancesCache(now) {
		err = s.setAppliancesCache(r.Context(), now)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = json.NewEncoder(w).Encode(s.appliancesCache)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func appendCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Max-Age", "86400")
}

func readJSONRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("failed to decode request as JSON: %w", err)
	}
	return nil
}

func (s *Server) bearerTokenAuthInfo() runtime.ClientAuthInfoWriter {
	return client.BearerToken(s.Token)
}

func (s *Server) optionsPreflightHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	appendCORSHeaders(w)
}

func (s *Server) postSendSignalHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	appendCORSHeaders(w)

	vars := mux.Vars(r)
	signalID := vars["id"]
	if signalID == "" {
		http.Error(w, "missing signalID", http.StatusBadRequest)
		return
	}

	p := &operations.Post1SignalsSignalSendParams{
		Signal: signalID,
	}
	p = p.WithTimeout(remoAPITimeout)
	p = p.WithContext(r.Context())
	_, err := s.RemoClient.Operations.Post1SignalsSignalSend(p, s.bearerTokenAuthInfo())
	if err != nil {
		s.logger.Errorf("failed to send Signal. err: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) postSendLightButtonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	appendCORSHeaders(w)

	vars := mux.Vars(r)
	applianceID := vars["id"]
	buttonID := vars["button_id"]

	if applianceID == "" {
		http.Error(w, "missing applianceID", http.StatusBadRequest)
		return
	}
	if buttonID == "" {
		http.Error(w, "missing buttonID", http.StatusBadRequest)
		return
	}

	p := &operations.Post1AppliancesApplianceLightParams{
		Appliance: applianceID,
		Button:    buttonID,
	}
	p = p.WithTimeout(remoAPITimeout)
	p = p.WithContext(r.Context())
	_, err := s.RemoClient.Operations.Post1AppliancesApplianceLight(p, s.bearerTokenAuthInfo())
	if err != nil {
		s.logger.Errorf("failed to send light Signal. err: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) getHealthHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
}

func (s *Server) preflightHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	appendCORSHeaders(w)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.WriteHeader(http.StatusNotFound)
}

func (s *Server) CreateRouter() http.Handler {
	r := mux.NewRouter()
	if s.prefix != "" {
		r = r.PathPrefix("/" + s.prefix).Subrouter()
	}
	r.PathPrefix("/").Methods("OPTIONS").HandlerFunc(s.preflightHandler)
	r.HandleFunc("/api/v1/health", s.getHealthHandler).Methods("GET")
	r.HandleFunc("/api/v1/appliances", s.getAppliancesHandler).Methods("GET").Queries("force", "{force}")
	r.HandleFunc("/api/v1/appliances/{id:[0-9a-z\\-]+}/buttons/{button_id:[0-9a-z\\-]+}/send", s.postSendLightButtonHandler).Methods("POST")
	r.HandleFunc("/api/v1/signals/{id:[0-9a-z\\-]+}/send", s.postSendSignalHandler).Methods("POST")
	r.PathPrefix("/api/").HandlerFunc(s.notFoundHandler)

	if _, err := staticFiles.ReadFile("index.html"); err != nil {
		log.Printf("started to serve WebUI")

		spa := spaHandler{staticFS: staticFiles, staticPath: "build", indexPath: "index.html"}
		r.PathPrefix("/").Handler(spa)
	} else {
		s.logger.Warnf("skipped to serve WebUI, because it has not been built")
	}

	return handlers.LoggingHandler(os.Stdout, r)
}
