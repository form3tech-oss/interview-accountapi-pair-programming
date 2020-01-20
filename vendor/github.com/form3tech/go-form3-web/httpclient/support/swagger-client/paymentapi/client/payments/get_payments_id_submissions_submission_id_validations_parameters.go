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

// NewGetPaymentsIDSubmissionsSubmissionIDValidationsParams creates a new GetPaymentsIDSubmissionsSubmissionIDValidationsParams object
// with the default values initialized.
func NewGetPaymentsIDSubmissionsSubmissionIDValidationsParams() *GetPaymentsIDSubmissionsSubmissionIDValidationsParams {
	var ()
	return &GetPaymentsIDSubmissionsSubmissionIDValidationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsIDSubmissionsSubmissionIDValidationsParamsWithTimeout creates a new GetPaymentsIDSubmissionsSubmissionIDValidationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsIDSubmissionsSubmissionIDValidationsParamsWithTimeout(timeout time.Duration) *GetPaymentsIDSubmissionsSubmissionIDValidationsParams {
	var ()
	return &GetPaymentsIDSubmissionsSubmissionIDValidationsParams{

		timeout: timeout,
	}
}

// NewGetPaymentsIDSubmissionsSubmissionIDValidationsParamsWithContext creates a new GetPaymentsIDSubmissionsSubmissionIDValidationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentsIDSubmissionsSubmissionIDValidationsParamsWithContext(ctx context.Context) *GetPaymentsIDSubmissionsSubmissionIDValidationsParams {
	var ()
	return &GetPaymentsIDSubmissionsSubmissionIDValidationsParams{

		Context: ctx,
	}
}

// NewGetPaymentsIDSubmissionsSubmissionIDValidationsParamsWithHTTPClient creates a new GetPaymentsIDSubmissionsSubmissionIDValidationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentsIDSubmissionsSubmissionIDValidationsParamsWithHTTPClient(client *http.Client) *GetPaymentsIDSubmissionsSubmissionIDValidationsParams {
	var ()
	return &GetPaymentsIDSubmissionsSubmissionIDValidationsParams{
		HTTPClient: client,
	}
}

/*GetPaymentsIDSubmissionsSubmissionIDValidationsParams contains all the parameters to send to the API endpoint
for the get payments ID submissions submission ID validations operation typically these are written to a http.Request
*/
type GetPaymentsIDSubmissionsSubmissionIDValidationsParams struct {

	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*SubmissionID
	  Submission Id

	*/
	SubmissionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payments ID submissions submission ID validations params
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) WithTimeout(timeout time.Duration) *GetPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payments ID submissions submission ID validations params
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payments ID submissions submission ID validations params
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) WithContext(ctx context.Context) *GetPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payments ID submissions submission ID validations params
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payments ID submissions submission ID validations params
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) WithHTTPClient(client *http.Client) *GetPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payments ID submissions submission ID validations params
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get payments ID submissions submission ID validations params
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) WithID(id strfmt.UUID) *GetPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get payments ID submissions submission ID validations params
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithSubmissionID adds the submissionID to the get payments ID submissions submission ID validations params
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) WithSubmissionID(submissionID strfmt.UUID) *GetPaymentsIDSubmissionsSubmissionIDValidationsParams {
	o.SetSubmissionID(submissionID)
	return o
}

// SetSubmissionID adds the submissionId to the get payments ID submissions submission ID validations params
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) SetSubmissionID(submissionID strfmt.UUID) {
	o.SubmissionID = submissionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param submissionId
	if err := r.SetPathParam("submissionId", o.SubmissionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
