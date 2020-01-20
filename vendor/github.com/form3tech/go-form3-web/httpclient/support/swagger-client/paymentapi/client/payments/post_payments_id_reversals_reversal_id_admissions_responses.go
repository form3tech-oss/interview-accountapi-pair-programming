// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech/go-form3-web/httpclient/support/swagger-client/paymentapi/models"
)

// PostPaymentsIDReversalsReversalIDAdmissionsReader is a Reader for the PostPaymentsIDReversalsReversalIDAdmissions structure.
type PostPaymentsIDReversalsReversalIDAdmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDReversalsReversalIDAdmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDReversalsReversalIDAdmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDReversalsReversalIDAdmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDReversalsReversalIDAdmissionsCreated creates a PostPaymentsIDReversalsReversalIDAdmissionsCreated with default headers values
func NewPostPaymentsIDReversalsReversalIDAdmissionsCreated() *PostPaymentsIDReversalsReversalIDAdmissionsCreated {
	return &PostPaymentsIDReversalsReversalIDAdmissionsCreated{}
}

/*PostPaymentsIDReversalsReversalIDAdmissionsCreated handles this case with default header values.

Reversal admission creation response
*/
type PostPaymentsIDReversalsReversalIDAdmissionsCreated struct {
	Payload *models.ReversalAdmissionCreationResponse
}

func (o *PostPaymentsIDReversalsReversalIDAdmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals/{reversalId}/admissions][%d] postPaymentsIdReversalsReversalIdAdmissionsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDReversalsReversalIDAdmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReversalAdmissionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReversalsReversalIDAdmissionsBadRequest creates a PostPaymentsIDReversalsReversalIDAdmissionsBadRequest with default headers values
func NewPostPaymentsIDReversalsReversalIDAdmissionsBadRequest() *PostPaymentsIDReversalsReversalIDAdmissionsBadRequest {
	return &PostPaymentsIDReversalsReversalIDAdmissionsBadRequest{}
}

/*PostPaymentsIDReversalsReversalIDAdmissionsBadRequest handles this case with default header values.

Reversal admission creation error
*/
type PostPaymentsIDReversalsReversalIDAdmissionsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReversalsReversalIDAdmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals/{reversalId}/admissions][%d] postPaymentsIdReversalsReversalIdAdmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDReversalsReversalIDAdmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
