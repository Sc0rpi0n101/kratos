// Code generated by go-swagger; DO NOT EDIT.

package public

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

// NewInitializeSelfServiceVerificationViaBrowserFlowParams creates a new InitializeSelfServiceVerificationViaBrowserFlowParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitializeSelfServiceVerificationViaBrowserFlowParams() *InitializeSelfServiceVerificationViaBrowserFlowParams {
	return &InitializeSelfServiceVerificationViaBrowserFlowParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitializeSelfServiceVerificationViaBrowserFlowParamsWithTimeout creates a new InitializeSelfServiceVerificationViaBrowserFlowParams object
// with the ability to set a timeout on a request.
func NewInitializeSelfServiceVerificationViaBrowserFlowParamsWithTimeout(timeout time.Duration) *InitializeSelfServiceVerificationViaBrowserFlowParams {
	return &InitializeSelfServiceVerificationViaBrowserFlowParams{
		timeout: timeout,
	}
}

// NewInitializeSelfServiceVerificationViaBrowserFlowParamsWithContext creates a new InitializeSelfServiceVerificationViaBrowserFlowParams object
// with the ability to set a context for a request.
func NewInitializeSelfServiceVerificationViaBrowserFlowParamsWithContext(ctx context.Context) *InitializeSelfServiceVerificationViaBrowserFlowParams {
	return &InitializeSelfServiceVerificationViaBrowserFlowParams{
		Context: ctx,
	}
}

// NewInitializeSelfServiceVerificationViaBrowserFlowParamsWithHTTPClient creates a new InitializeSelfServiceVerificationViaBrowserFlowParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitializeSelfServiceVerificationViaBrowserFlowParamsWithHTTPClient(client *http.Client) *InitializeSelfServiceVerificationViaBrowserFlowParams {
	return &InitializeSelfServiceVerificationViaBrowserFlowParams{
		HTTPClient: client,
	}
}

/* InitializeSelfServiceVerificationViaBrowserFlowParams contains all the parameters to send to the API endpoint
   for the initialize self service verification via browser flow operation.

   Typically these are written to a http.Request.
*/
type InitializeSelfServiceVerificationViaBrowserFlowParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initialize self service verification via browser flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeSelfServiceVerificationViaBrowserFlowParams) WithDefaults() *InitializeSelfServiceVerificationViaBrowserFlowParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initialize self service verification via browser flow params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeSelfServiceVerificationViaBrowserFlowParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initialize self service verification via browser flow params
func (o *InitializeSelfServiceVerificationViaBrowserFlowParams) WithTimeout(timeout time.Duration) *InitializeSelfServiceVerificationViaBrowserFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initialize self service verification via browser flow params
func (o *InitializeSelfServiceVerificationViaBrowserFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initialize self service verification via browser flow params
func (o *InitializeSelfServiceVerificationViaBrowserFlowParams) WithContext(ctx context.Context) *InitializeSelfServiceVerificationViaBrowserFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initialize self service verification via browser flow params
func (o *InitializeSelfServiceVerificationViaBrowserFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initialize self service verification via browser flow params
func (o *InitializeSelfServiceVerificationViaBrowserFlowParams) WithHTTPClient(client *http.Client) *InitializeSelfServiceVerificationViaBrowserFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initialize self service verification via browser flow params
func (o *InitializeSelfServiceVerificationViaBrowserFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InitializeSelfServiceVerificationViaBrowserFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}