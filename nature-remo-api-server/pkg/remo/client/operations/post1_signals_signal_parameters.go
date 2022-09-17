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

// NewPost1SignalsSignalParams creates a new Post1SignalsSignalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPost1SignalsSignalParams() *Post1SignalsSignalParams {
	return &Post1SignalsSignalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPost1SignalsSignalParamsWithTimeout creates a new Post1SignalsSignalParams object
// with the ability to set a timeout on a request.
func NewPost1SignalsSignalParamsWithTimeout(timeout time.Duration) *Post1SignalsSignalParams {
	return &Post1SignalsSignalParams{
		timeout: timeout,
	}
}

// NewPost1SignalsSignalParamsWithContext creates a new Post1SignalsSignalParams object
// with the ability to set a context for a request.
func NewPost1SignalsSignalParamsWithContext(ctx context.Context) *Post1SignalsSignalParams {
	return &Post1SignalsSignalParams{
		Context: ctx,
	}
}

// NewPost1SignalsSignalParamsWithHTTPClient creates a new Post1SignalsSignalParams object
// with the ability to set a custom HTTPClient for a request.
func NewPost1SignalsSignalParamsWithHTTPClient(client *http.Client) *Post1SignalsSignalParams {
	return &Post1SignalsSignalParams{
		HTTPClient: client,
	}
}

/* Post1SignalsSignalParams contains all the parameters to send to the API endpoint
   for the post1 signals signal operation.

   Typically these are written to a http.Request.
*/
type Post1SignalsSignalParams struct {

	/* Image.

	   Basename of the image file included in the app. Ex: "ico_io"

	*/
	Image string

	/* Name.

	   Signal name
	*/
	Name string

	/* Signal.

	   Signal ID.
	*/
	Signal string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post1 signals signal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Post1SignalsSignalParams) WithDefaults() *Post1SignalsSignalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post1 signals signal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Post1SignalsSignalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post1 signals signal params
func (o *Post1SignalsSignalParams) WithTimeout(timeout time.Duration) *Post1SignalsSignalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post1 signals signal params
func (o *Post1SignalsSignalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post1 signals signal params
func (o *Post1SignalsSignalParams) WithContext(ctx context.Context) *Post1SignalsSignalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post1 signals signal params
func (o *Post1SignalsSignalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post1 signals signal params
func (o *Post1SignalsSignalParams) WithHTTPClient(client *http.Client) *Post1SignalsSignalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post1 signals signal params
func (o *Post1SignalsSignalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImage adds the image to the post1 signals signal params
func (o *Post1SignalsSignalParams) WithImage(image string) *Post1SignalsSignalParams {
	o.SetImage(image)
	return o
}

// SetImage adds the image to the post1 signals signal params
func (o *Post1SignalsSignalParams) SetImage(image string) {
	o.Image = image
}

// WithName adds the name to the post1 signals signal params
func (o *Post1SignalsSignalParams) WithName(name string) *Post1SignalsSignalParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post1 signals signal params
func (o *Post1SignalsSignalParams) SetName(name string) {
	o.Name = name
}

// WithSignal adds the signal to the post1 signals signal params
func (o *Post1SignalsSignalParams) WithSignal(signal string) *Post1SignalsSignalParams {
	o.SetSignal(signal)
	return o
}

// SetSignal adds the signal to the post1 signals signal params
func (o *Post1SignalsSignalParams) SetSignal(signal string) {
	o.Signal = signal
}

// WriteToRequest writes these params to a swagger request
func (o *Post1SignalsSignalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param image
	frImage := o.Image
	fImage := frImage
	if fImage != "" {
		if err := r.SetFormParam("image", fImage); err != nil {
			return err
		}
	}

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
		}
	}

	// path param signal
	if err := r.SetPathParam("signal", o.Signal); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}