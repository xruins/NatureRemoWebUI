// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Get1Appliances(params *Get1AppliancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Get1AppliancesOK, error)

	Get1AppliancesApplianceSignals(params *Get1AppliancesApplianceSignalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Get1AppliancesApplianceSignalsOK, error)

	Get1Devices(params *Get1DevicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Get1DevicesOK, error)

	Get1UsersMe(params *Get1UsersMeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Get1UsersMeOK, error)

	Post1ApplianceOrders(params *Post1ApplianceOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1ApplianceOrdersOK, error)

	Post1Appliances(params *Post1AppliancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesCreated, error)

	Post1AppliancesAppliance(params *Post1AppliancesApplianceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceOK, error)

	Post1AppliancesApplianceAirconSettings(params *Post1AppliancesApplianceAirconSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceAirconSettingsOK, error)

	Post1AppliancesApplianceDelete(params *Post1AppliancesApplianceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceDeleteOK, error)

	Post1AppliancesApplianceLight(params *Post1AppliancesApplianceLightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceLightOK, error)

	Post1AppliancesApplianceSignalOrders(params *Post1AppliancesApplianceSignalOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceSignalOrdersOK, error)

	Post1AppliancesApplianceSignals(params *Post1AppliancesApplianceSignalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceSignalsCreated, error)

	Post1AppliancesApplianceTv(params *Post1AppliancesApplianceTvParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceTvOK, error)

	Post1Detectappliance(params *Post1DetectapplianceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1DetectapplianceOK, error)

	Post1DevicesDevice(params *Post1DevicesDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1DevicesDeviceOK, error)

	Post1DevicesDeviceDelete(params *Post1DevicesDeviceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1DevicesDeviceDeleteOK, error)

	Post1DevicesDeviceHumidityOffset(params *Post1DevicesDeviceHumidityOffsetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1DevicesDeviceHumidityOffsetOK, error)

	Post1DevicesDeviceTemperatureOffset(params *Post1DevicesDeviceTemperatureOffsetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1DevicesDeviceTemperatureOffsetOK, error)

	Post1SignalsSignal(params *Post1SignalsSignalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1SignalsSignalOK, error)

	Post1SignalsSignalDelete(params *Post1SignalsSignalDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1SignalsSignalDeleteOK, error)

	Post1SignalsSignalSend(params *Post1SignalsSignalSendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1SignalsSignalSendOK, error)

	Post1UsersMe(params *Post1UsersMeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1UsersMeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Get1Appliances Fetch the list of appliances.
*/
func (a *Client) Get1Appliances(params *Get1AppliancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Get1AppliancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet1AppliancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get1Appliances",
		Method:             "GET",
		PathPattern:        "/1/appliances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Get1AppliancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get1AppliancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get1Appliances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Get1AppliancesApplianceSignals Fetch signals registered under this appliance.
*/
func (a *Client) Get1AppliancesApplianceSignals(params *Get1AppliancesApplianceSignalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Get1AppliancesApplianceSignalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet1AppliancesApplianceSignalsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get1AppliancesApplianceSignals",
		Method:             "GET",
		PathPattern:        "/1/appliances/{appliance}/signals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Get1AppliancesApplianceSignalsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get1AppliancesApplianceSignalsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get1AppliancesApplianceSignals: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Get1Devices Fetch the list of Remo devices the user has access to.
*/
func (a *Client) Get1Devices(params *Get1DevicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Get1DevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet1DevicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get1Devices",
		Method:             "GET",
		PathPattern:        "/1/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Get1DevicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get1DevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get1Devices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Get1UsersMe Fetch the authenticated user's information.
*/
func (a *Client) Get1UsersMe(params *Get1UsersMeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Get1UsersMeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet1UsersMeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get1UsersMe",
		Method:             "GET",
		PathPattern:        "/1/users/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Get1UsersMeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get1UsersMeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get1UsersMe: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1ApplianceOrders Reorder appliances.
*/
func (a *Client) Post1ApplianceOrders(params *Post1ApplianceOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1ApplianceOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1ApplianceOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1ApplianceOrders",
		Method:             "POST",
		PathPattern:        "/1/appliance_orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1ApplianceOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1ApplianceOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1ApplianceOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1Appliances Create a new appliance.
*/
func (a *Client) Post1Appliances(params *Post1AppliancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1AppliancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1Appliances",
		Method:             "POST",
		PathPattern:        "/1/appliances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1AppliancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1AppliancesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1Appliances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1AppliancesAppliance Update appliance.
*/
func (a *Client) Post1AppliancesAppliance(params *Post1AppliancesApplianceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1AppliancesApplianceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1AppliancesAppliance",
		Method:             "POST",
		PathPattern:        "/1/appliances/{appliance}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1AppliancesApplianceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1AppliancesApplianceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1AppliancesAppliance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1AppliancesApplianceAirconSettings Update air conditioner settings.
*/
func (a *Client) Post1AppliancesApplianceAirconSettings(params *Post1AppliancesApplianceAirconSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceAirconSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1AppliancesApplianceAirconSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1AppliancesApplianceAirconSettings",
		Method:             "POST",
		PathPattern:        "/1/appliances/{appliance}/aircon_settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1AppliancesApplianceAirconSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1AppliancesApplianceAirconSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1AppliancesApplianceAirconSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1AppliancesApplianceDelete Delete appliance.
*/
func (a *Client) Post1AppliancesApplianceDelete(params *Post1AppliancesApplianceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1AppliancesApplianceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1AppliancesApplianceDelete",
		Method:             "POST",
		PathPattern:        "/1/appliances/{appliance}/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1AppliancesApplianceDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1AppliancesApplianceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1AppliancesApplianceDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1AppliancesApplianceLight Send light infrared signal.
*/
func (a *Client) Post1AppliancesApplianceLight(params *Post1AppliancesApplianceLightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceLightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1AppliancesApplianceLightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1AppliancesApplianceLight",
		Method:             "POST",
		PathPattern:        "/1/appliances/{appliance}/light",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1AppliancesApplianceLightReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1AppliancesApplianceLightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1AppliancesApplianceLight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1AppliancesApplianceSignalOrders Reorder signals under this appliance.
*/
func (a *Client) Post1AppliancesApplianceSignalOrders(params *Post1AppliancesApplianceSignalOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceSignalOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1AppliancesApplianceSignalOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1AppliancesApplianceSignalOrders",
		Method:             "POST",
		PathPattern:        "/1/appliances/{appliance}/signal_orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1AppliancesApplianceSignalOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1AppliancesApplianceSignalOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1AppliancesApplianceSignalOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1AppliancesApplianceSignals Create a signal under this appliance.
*/
func (a *Client) Post1AppliancesApplianceSignals(params *Post1AppliancesApplianceSignalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceSignalsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1AppliancesApplianceSignalsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1AppliancesApplianceSignals",
		Method:             "POST",
		PathPattern:        "/1/appliances/{appliance}/signals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1AppliancesApplianceSignalsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1AppliancesApplianceSignalsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1AppliancesApplianceSignals: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1AppliancesApplianceTv Send tv infrared signal.
*/
func (a *Client) Post1AppliancesApplianceTv(params *Post1AppliancesApplianceTvParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1AppliancesApplianceTvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1AppliancesApplianceTvParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1AppliancesApplianceTv",
		Method:             "POST",
		PathPattern:        "/1/appliances/{appliance}/tv",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1AppliancesApplianceTvReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1AppliancesApplianceTvOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1AppliancesApplianceTv: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1Detectappliance Find the air conditioner best matching the provided infrared signal.
*/
func (a *Client) Post1Detectappliance(params *Post1DetectapplianceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1DetectapplianceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1DetectapplianceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1Detectappliance",
		Method:             "POST",
		PathPattern:        "/1/detectappliance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1DetectapplianceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1DetectapplianceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1Detectappliance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1DevicesDevice Update Remo
*/
func (a *Client) Post1DevicesDevice(params *Post1DevicesDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1DevicesDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1DevicesDeviceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1DevicesDevice",
		Method:             "POST",
		PathPattern:        "/1/devices/{device}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1DevicesDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1DevicesDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1DevicesDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1DevicesDeviceDelete Delete Remo.
*/
func (a *Client) Post1DevicesDeviceDelete(params *Post1DevicesDeviceDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1DevicesDeviceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1DevicesDeviceDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1DevicesDeviceDelete",
		Method:             "POST",
		PathPattern:        "/1/devices/{device}/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1DevicesDeviceDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1DevicesDeviceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1DevicesDeviceDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1DevicesDeviceHumidityOffset Update humidity offset.
*/
func (a *Client) Post1DevicesDeviceHumidityOffset(params *Post1DevicesDeviceHumidityOffsetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1DevicesDeviceHumidityOffsetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1DevicesDeviceHumidityOffsetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1DevicesDeviceHumidityOffset",
		Method:             "POST",
		PathPattern:        "/1/devices/{device}/humidity_offset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1DevicesDeviceHumidityOffsetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1DevicesDeviceHumidityOffsetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1DevicesDeviceHumidityOffset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1DevicesDeviceTemperatureOffset Update temperature offset.
*/
func (a *Client) Post1DevicesDeviceTemperatureOffset(params *Post1DevicesDeviceTemperatureOffsetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1DevicesDeviceTemperatureOffsetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1DevicesDeviceTemperatureOffsetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1DevicesDeviceTemperatureOffset",
		Method:             "POST",
		PathPattern:        "/1/devices/{device}/temperature_offset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1DevicesDeviceTemperatureOffsetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1DevicesDeviceTemperatureOffsetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1DevicesDeviceTemperatureOffset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1SignalsSignal Update infrared signal.
*/
func (a *Client) Post1SignalsSignal(params *Post1SignalsSignalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1SignalsSignalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1SignalsSignalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1SignalsSignal",
		Method:             "POST",
		PathPattern:        "/1/signals/{signal}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1SignalsSignalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1SignalsSignalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1SignalsSignal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1SignalsSignalDelete Delete an infrared signal.
*/
func (a *Client) Post1SignalsSignalDelete(params *Post1SignalsSignalDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1SignalsSignalDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1SignalsSignalDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1SignalsSignalDelete",
		Method:             "POST",
		PathPattern:        "/1/signals/{signal}/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1SignalsSignalDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1SignalsSignalDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1SignalsSignalDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1SignalsSignalSend Send infrared signal.
*/
func (a *Client) Post1SignalsSignalSend(params *Post1SignalsSignalSendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1SignalsSignalSendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1SignalsSignalSendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1SignalsSignalSend",
		Method:             "POST",
		PathPattern:        "/1/signals/{signal}/send",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1SignalsSignalSendReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1SignalsSignalSendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1SignalsSignalSend: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Post1UsersMe Update authenticated user's information.
*/
func (a *Client) Post1UsersMe(params *Post1UsersMeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Post1UsersMeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPost1UsersMeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Post1UsersMe",
		Method:             "POST",
		PathPattern:        "/1/users/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Post1UsersMeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Post1UsersMeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Post1UsersMe: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
