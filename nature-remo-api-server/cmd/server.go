package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xruins/NatureRemoWebUI/nature-remo-api-server/pkg/logger"
	"github.com/xruins/NatureRemoWebUI/nature-remo-api-server/pkg/server"
)

// serverCmd runs the server to host Nature Remo Proxy API and WebUI.
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Nature Remo Proxy & WebUI server",
	Run: func(cmd *cobra.Command, args []string) {
		l, err := logger.NewZapLogger(logLevel)
		if err != nil {
			log.Fatalf("failed to initialize logger. err: %s", err)
		}

		if token == "" {
			l.Fatalf("required to set token. set with `-t` option or `%s`", paramToken)
		}

		s := server.NewServer(l, token, pathPrefix)
		l.Infof("API server is listening on %s:%s", bindAddress, port)
		srv := &http.Server{
			Addr:         bindAddress + ":" + port,
			Handler:      s.CreateRouter(),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		l.Info(srv.ListenAndServe())
	},
}

const (
	paramBindAddress = "API_SERVER_ADDR"
	paramPort        = "API_SERVER_PORT"
	paramPathPrefix  = "API_SERVER_PATH_PREFIX"
	paramLogLevel    = "API_SERVER_LOG_LEVEL"
	paramToken       = "REMO_API_TOKEN"
)

var bindAddress, port, token, pathPrefix, logLevel string

func init() {
	rootCmd.AddCommand(serverCmd)

	viper.SetDefault(paramBindAddress, "0.0.0.0")
	viper.SetDefault(paramPort, "3000")
	viper.SetDefault(paramLogLevel, "info")
	viper.SetDefault(paramPathPrefix, "")

	serverCmd.Flags().StringVarP(&bindAddress, "bind-address", "b", viper.GetString(paramBindAddress), "Host the server listen on")
	serverCmd.Flags().StringVarP(&port, "port", "p", viper.GetString(paramPort), "Port the server listen on")
	serverCmd.Flags().StringVarP(&token, "token", "t", viper.GetString(paramToken), "API Token of Nature Remo Cloud API")
	serverCmd.Flags().StringVar(&pathPrefix, "prefix", viper.GetString(paramPathPrefix), "Prefix for server endpoints")
	serverCmd.Flags().StringVarP(&logLevel, "loglevel", "l", viper.GetString(paramLogLevel), "Logging level of API server")
}
