// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBatchesIDParams creates a new GetBatchesIDParams object
// with the default values initialized.
func NewGetBatchesIDParams() *GetBatchesIDParams {
	var ()
	return &GetBatchesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBatchesIDParamsWithTimeout creates a new GetBatchesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBatchesIDParamsWithTimeout(timeout time.Duration) *GetBatchesIDParams {
	var ()
	return &GetBatchesIDParams{

		timeout: timeout,
	}
}

// NewGetBatchesIDParamsWithContext creates a new GetBatchesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBatchesIDParamsWithContext(ctx context.Context) *GetBatchesIDParams {
	var ()
	return &GetBatchesIDParams{

		Context: ctx,
	}
}

// NewGetBatchesIDParamsWithHTTPClient creates a new GetBatchesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBatchesIDParamsWithHTTPClient(client *http.Client) *GetBatchesIDParams {
	var ()
	return &GetBatchesIDParams{
		HTTPClient: client,
	}
}

/*GetBatchesIDParams contains all the parameters to send to the API endpoint
for the get batches ID operation typically these are written to a http.Request
*/
type GetBatchesIDParams struct {

	/*ID
	  Batch Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get batches ID params
func (o *GetBatchesIDParams) WithTimeout(timeout time.Duration) *GetBatchesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get batches ID params
func (o *GetBatchesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get batches ID params
func (o *GetBatchesIDParams) WithContext(ctx context.Context) *GetBatchesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get batches ID params
func (o *GetBatchesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get batches ID params
func (o *GetBatchesIDParams) WithHTTPClient(client *http.Client) *GetBatchesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get batches ID params
func (o *GetBatchesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get batches ID params
func (o *GetBatchesIDParams) WithID(id strfmt.UUID) *GetBatchesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get batches ID params
func (o *GetBatchesIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBatchesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
