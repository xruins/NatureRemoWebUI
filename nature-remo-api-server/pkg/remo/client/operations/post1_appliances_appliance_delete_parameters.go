// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPost1AppliancesApplianceDeleteParams creates a new Post1AppliancesApplianceDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPost1AppliancesApplianceDeleteParams() *Post1AppliancesApplianceDeleteParams {
	return &Post1AppliancesApplianceDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPost1AppliancesApplianceDeleteParamsWithTimeout creates a new Post1AppliancesApplianceDeleteParams object
// with the ability to set a timeout on a request.
func NewPost1AppliancesApplianceDeleteParamsWithTimeout(timeout time.Duration) *Post1AppliancesApplianceDeleteParams {
	return &Post1AppliancesApplianceDeleteParams{
		timeout: timeout,
	}
}

// NewPost1AppliancesApplianceDeleteParamsWithContext creates a new Post1AppliancesApplianceDeleteParams object
// with the ability to set a context for a request.
func NewPost1AppliancesApplianceDeleteParamsWithContext(ctx context.Context) *Post1AppliancesApplianceDeleteParams {
	return &Post1AppliancesApplianceDeleteParams{
		Context: ctx,
	}
}

// NewPost1AppliancesApplianceDeleteParamsWithHTTPClient creates a new Post1AppliancesApplianceDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPost1AppliancesApplianceDeleteParamsWithHTTPClient(client *http.Client) *Post1AppliancesApplianceDeleteParams {
	return &Post1AppliancesApplianceDeleteParams{
		HTTPClient: client,
	}
}

/* Post1AppliancesApplianceDeleteParams contains all the parameters to send to the API endpoint
   for the post1 appliances appliance delete operation.

   Typically these are written to a http.Request.
*/
type Post1AppliancesApplianceDeleteParams struct {

	/* Appliance.

	   Appliance ID.
	*/
	Appliance string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post1 appliances appliance delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Post1AppliancesApplianceDeleteParams) WithDefaults() *Post1AppliancesApplianceDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post1 appliances appliance delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Post1AppliancesApplianceDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post1 appliances appliance delete params
func (o *Post1AppliancesApplianceDeleteParams) WithTimeout(timeout time.Duration) *Post1AppliancesApplianceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post1 appliances appliance delete params
func (o *Post1AppliancesApplianceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post1 appliances appliance delete params
func (o *Post1AppliancesApplianceDeleteParams) WithContext(ctx context.Context) *Post1AppliancesApplianceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post1 appliances appliance delete params
func (o *Post1AppliancesApplianceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post1 appliances appliance delete params
func (o *Post1AppliancesApplianceDeleteParams) WithHTTPClient(client *http.Client) *Post1AppliancesApplianceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post1 appliances appliance delete params
func (o *Post1AppliancesApplianceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppliance adds the appliance to the post1 appliances appliance delete params
func (o *Post1AppliancesApplianceDeleteParams) WithAppliance(appliance string) *Post1AppliancesApplianceDeleteParams {
	o.SetAppliance(appliance)
	return o
}

// SetAppliance adds the appliance to the post1 appliances appliance delete params
func (o *Post1AppliancesApplianceDeleteParams) SetAppliance(appliance string) {
	o.Appliance = appliance
}

// WriteToRequest writes these params to a swagger request
func (o *Post1AppliancesApplianceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appliance
	if err := r.SetPathParam("appliance", o.Appliance); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
