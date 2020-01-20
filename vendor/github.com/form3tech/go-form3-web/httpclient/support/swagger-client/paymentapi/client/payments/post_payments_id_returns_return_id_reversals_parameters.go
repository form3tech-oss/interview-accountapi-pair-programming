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

// NewPostPaymentsIDReturnsReturnIDReversalsParams creates a new PostPaymentsIDReturnsReturnIDReversalsParams object
// with the default values initialized.
func NewPostPaymentsIDReturnsReturnIDReversalsParams() *PostPaymentsIDReturnsReturnIDReversalsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDReversalsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsIDReturnsReturnIDReversalsParamsWithTimeout creates a new PostPaymentsIDReturnsReturnIDReversalsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsIDReturnsReturnIDReversalsParamsWithTimeout(timeout time.Duration) *PostPaymentsIDReturnsReturnIDReversalsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDReversalsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsIDReturnsReturnIDReversalsParamsWithContext creates a new PostPaymentsIDReturnsReturnIDReversalsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsIDReturnsReturnIDReversalsParamsWithContext(ctx context.Context) *PostPaymentsIDReturnsReturnIDReversalsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDReversalsParams{

		Context: ctx,
	}
}

// NewPostPaymentsIDReturnsReturnIDReversalsParamsWithHTTPClient creates a new PostPaymentsIDReturnsReturnIDReversalsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsIDReturnsReturnIDReversalsParamsWithHTTPClient(client *http.Client) *PostPaymentsIDReturnsReturnIDReversalsParams {
	var ()
	return &PostPaymentsIDReturnsReturnIDReversalsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsIDReturnsReturnIDReversalsParams contains all the parameters to send to the API endpoint
for the post payments ID returns return ID reversals operation typically these are written to a http.Request
*/
type PostPaymentsIDReturnsReturnIDReversalsParams struct {

	/*ReturnReversalCreationRequest*/
	ReturnReversalCreationRequest *models.ReturnReversalCreation
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

// WithTimeout adds the timeout to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) WithTimeout(timeout time.Duration) *PostPaymentsIDReturnsReturnIDReversalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) WithContext(ctx context.Context) *PostPaymentsIDReturnsReturnIDReversalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) WithHTTPClient(client *http.Client) *PostPaymentsIDReturnsReturnIDReversalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnReversalCreationRequest adds the returnReversalCreationRequest to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) WithReturnReversalCreationRequest(returnReversalCreationRequest *models.ReturnReversalCreation) *PostPaymentsIDReturnsReturnIDReversalsParams {
	o.SetReturnReversalCreationRequest(returnReversalCreationRequest)
	return o
}

// SetReturnReversalCreationRequest adds the returnReversalCreationRequest to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) SetReturnReversalCreationRequest(returnReversalCreationRequest *models.ReturnReversalCreation) {
	o.ReturnReversalCreationRequest = returnReversalCreationRequest
}

// WithID adds the id to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) WithID(id strfmt.UUID) *PostPaymentsIDReturnsReturnIDReversalsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithReturnID adds the returnID to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) WithReturnID(returnID strfmt.UUID) *PostPaymentsIDReturnsReturnIDReversalsParams {
	o.SetReturnID(returnID)
	return o
}

// SetReturnID adds the returnId to the post payments ID returns return ID reversals params
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) SetReturnID(returnID strfmt.UUID) {
	o.ReturnID = returnID
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsIDReturnsReturnIDReversalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReturnReversalCreationRequest != nil {
		if err := r.SetBodyParam(o.ReturnReversalCreationRequest); err != nil {
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
