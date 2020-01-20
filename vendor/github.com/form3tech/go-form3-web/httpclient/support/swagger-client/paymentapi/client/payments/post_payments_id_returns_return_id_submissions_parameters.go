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

// NewPostPaymentsIDReturnsReturnIDSubmissionsParams creates a new PostPaymentsIDReturnsReturnIDSubmissionsParams object
// with the default values initialized.
func NewPostPaymentsIDReturnsReturnIDSubmissionsParams() *PostPaymentsIDReturnsReturnIDSubmissionsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDSubmissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsParamsWithTimeout creates a new PostPaymentsIDReturnsReturnIDSubmissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsIDReturnsReturnIDSubmissionsParamsWithTimeout(timeout time.Duration) *PostPaymentsIDReturnsReturnIDSubmissionsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDSubmissionsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsParamsWithContext creates a new PostPaymentsIDReturnsReturnIDSubmissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsIDReturnsReturnIDSubmissionsParamsWithContext(ctx context.Context) *PostPaymentsIDReturnsReturnIDSubmissionsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDSubmissionsParams{

		Context: ctx,
	}
}

// NewPostPaymentsIDReturnsReturnIDSubmissionsParamsWithHTTPClient creates a new PostPaymentsIDReturnsReturnIDSubmissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsIDReturnsReturnIDSubmissionsParamsWithHTTPClient(client *http.Client) *PostPaymentsIDReturnsReturnIDSubmissionsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDSubmissionsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsIDReturnsReturnIDSubmissionsParams contains all the parameters to send to the API endpoint
for the post payments ID returns return ID submissions operation typically these are written to a http.Request
*/
type PostPaymentsIDReturnsReturnIDSubmissionsParams struct {

	/*ReturnSubmissionCreationRequest*/
	ReturnSubmissionCreationRequest *models.ReturnSubmissionCreation
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID
	/*ReturnID
	  Return Id

	*/
	ReturnID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) WithTimeout(timeout time.Duration) *PostPaymentsIDReturnsReturnIDSubmissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) WithContext(ctx context.Context) *PostPaymentsIDReturnsReturnIDSubmissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) WithHTTPClient(client *http.Client) *PostPaymentsIDReturnsReturnIDSubmissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnSubmissionCreationRequest adds the returnSubmissionCreationRequest to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) WithReturnSubmissionCreationRequest(returnSubmissionCreationRequest *models.ReturnSubmissionCreation) *PostPaymentsIDReturnsReturnIDSubmissionsParams {
	o.SetReturnSubmissionCreationRequest(returnSubmissionCreationRequest)
	return o
}

// SetReturnSubmissionCreationRequest adds the returnSubmissionCreationRequest to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) SetReturnSubmissionCreationRequest(returnSubmissionCreationRequest *models.ReturnSubmissionCreation) {
	o.ReturnSubmissionCreationRequest = returnSubmissionCreationRequest
}

// WithID adds the id to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) WithID(id strfmt.UUID) *PostPaymentsIDReturnsReturnIDSubmissionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithReturnID adds the returnID to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) WithReturnID(returnID strfmt.UUID) *PostPaymentsIDReturnsReturnIDSubmissionsParams {
	o.SetReturnID(returnID)
	return o
}

// SetReturnID adds the returnId to the post payments ID returns return ID submissions params
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) SetReturnID(returnID strfmt.UUID) {
	o.ReturnID = returnID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsIDReturnsReturnIDSubmissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReturnSubmissionCreationRequest != nil {
		if err := r.SetBodyParam(o.ReturnSubmissionCreationRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param returnId
	if err := r.SetPathParam("returnId", o.ReturnID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
