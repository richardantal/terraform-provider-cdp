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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// NewInitializeAWSComputeClusterParams creates a new InitializeAWSComputeClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitializeAWSComputeClusterParams() *InitializeAWSComputeClusterParams {
	return &InitializeAWSComputeClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitializeAWSComputeClusterParamsWithTimeout creates a new InitializeAWSComputeClusterParams object
// with the ability to set a timeout on a request.
func NewInitializeAWSComputeClusterParamsWithTimeout(timeout time.Duration) *InitializeAWSComputeClusterParams {
	return &InitializeAWSComputeClusterParams{
		timeout: timeout,
	}
}

// NewInitializeAWSComputeClusterParamsWithContext creates a new InitializeAWSComputeClusterParams object
// with the ability to set a context for a request.
func NewInitializeAWSComputeClusterParamsWithContext(ctx context.Context) *InitializeAWSComputeClusterParams {
	return &InitializeAWSComputeClusterParams{
		Context: ctx,
	}
}

// NewInitializeAWSComputeClusterParamsWithHTTPClient creates a new InitializeAWSComputeClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitializeAWSComputeClusterParamsWithHTTPClient(client *http.Client) *InitializeAWSComputeClusterParams {
	return &InitializeAWSComputeClusterParams{
		HTTPClient: client,
	}
}

/*
InitializeAWSComputeClusterParams contains all the parameters to send to the API endpoint

	for the initialize a w s compute cluster operation.

	Typically these are written to a http.Request.
*/
type InitializeAWSComputeClusterParams struct {

	// Input.
	Input *models.InitializeAWSComputeClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initialize a w s compute cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeAWSComputeClusterParams) WithDefaults() *InitializeAWSComputeClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initialize a w s compute cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeAWSComputeClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initialize a w s compute cluster params
func (o *InitializeAWSComputeClusterParams) WithTimeout(timeout time.Duration) *InitializeAWSComputeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initialize a w s compute cluster params
func (o *InitializeAWSComputeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initialize a w s compute cluster params
func (o *InitializeAWSComputeClusterParams) WithContext(ctx context.Context) *InitializeAWSComputeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initialize a w s compute cluster params
func (o *InitializeAWSComputeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initialize a w s compute cluster params
func (o *InitializeAWSComputeClusterParams) WithHTTPClient(client *http.Client) *InitializeAWSComputeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initialize a w s compute cluster params
func (o *InitializeAWSComputeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the initialize a w s compute cluster params
func (o *InitializeAWSComputeClusterParams) WithInput(input *models.InitializeAWSComputeClusterRequest) *InitializeAWSComputeClusterParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the initialize a w s compute cluster params
func (o *InitializeAWSComputeClusterParams) SetInput(input *models.InitializeAWSComputeClusterRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *InitializeAWSComputeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
