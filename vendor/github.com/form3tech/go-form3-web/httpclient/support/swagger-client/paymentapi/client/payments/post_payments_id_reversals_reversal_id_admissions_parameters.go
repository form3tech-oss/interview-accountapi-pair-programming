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

	"github.com/form3tech/go-form3-web/httpclient/support/swagger-client/paymentapi/models"
)

// NewPostPaymentsIDReversalsReversalIDAdmissionsParams creates a new PostPaymentsIDReversalsReversalIDAdmissionsParams object
// with the default values initialized.
func NewPostPaymentsIDReversalsReversalIDAdmissionsParams() *PostPaymentsIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDReversalsReversalIDAdmissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsIDReversalsReversalIDAdmissionsParamsWithTimeout creates a new PostPaymentsIDReversalsReversalIDAdmissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsIDReversalsReversalIDAdmissionsParamsWithTimeout(timeout time.Duration) *PostPaymentsIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDReversalsReversalIDAdmissionsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsIDReversalsReversalIDAdmissionsParamsWithContext creates a new PostPaymentsIDReversalsReversalIDAdmissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsIDReversalsReversalIDAdmissionsParamsWithContext(ctx context.Context) *PostPaymentsIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDReversalsReversalIDAdmissionsParams{

		Context: ctx,
	}
}

// NewPostPaymentsIDReversalsReversalIDAdmissionsParamsWithHTTPClient creates a new PostPaymentsIDReversalsReversalIDAdmissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsIDReversalsReversalIDAdmissionsParamsWithHTTPClient(client *http.Client) *PostPaymentsIDReversalsReversalIDAdmissionsParams {
	var ()
	return &PostPaymentsIDReversalsReversalIDAdmissionsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsIDReversalsReversalIDAdmissionsParams contains all the parameters to send to the API endpoint
for the post payments ID reversals reversal ID admissions operation typically these are written to a http.Request
*/
type PostPaymentsIDReversalsReversalIDAdmissionsParams struct {

	/*ReversalAdmissionCreationRequest*/
	ReversalAdmissionCreationRequest *models.ReversalAdmissionCreation
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*ReversalID
	  Reversal Id

	*/
	ReversalID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) WithTimeout(timeout time.Duration) *PostPaymentsIDReversalsReversalIDAdmissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) WithContext(ctx context.Context) *PostPaymentsIDReversalsReversalIDAdmissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) WithHTTPClient(client *http.Client) *PostPaymentsIDReversalsReversalIDAdmissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReversalAdmissionCreationRequest adds the reversalAdmissionCreationRequest to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) WithReversalAdmissionCreationRequest(reversalAdmissionCreationRequest *models.ReversalAdmissionCreation) *PostPaymentsIDReversalsReversalIDAdmissionsParams {
	o.SetReversalAdmissionCreationRequest(reversalAdmissionCreationRequest)
	return o
}

// SetReversalAdmissionCreationRequest adds the reversalAdmissionCreationRequest to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) SetReversalAdmissionCreationRequest(reversalAdmissionCreationRequest *models.ReversalAdmissionCreation) {
	o.ReversalAdmissionCreationRequest = reversalAdmissionCreationRequest
}

// WithID adds the id to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) WithID(id strfmt.UUID) *PostPaymentsIDReversalsReversalIDAdmissionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithReversalID adds the reversalID to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) WithReversalID(reversalID strfmt.UUID) *PostPaymentsIDReversalsReversalIDAdmissionsParams {
	o.SetReversalID(reversalID)
	return o
}

// SetReversalID adds the reversalId to the post payments ID reversals reversal ID admissions params
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) SetReversalID(reversalID strfmt.UUID) {
	o.ReversalID = reversalID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsIDReversalsReversalIDAdmissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReversalAdmissionCreationRequest != nil {
		if err := r.SetBodyParam(o.ReversalAdmissionCreationRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
