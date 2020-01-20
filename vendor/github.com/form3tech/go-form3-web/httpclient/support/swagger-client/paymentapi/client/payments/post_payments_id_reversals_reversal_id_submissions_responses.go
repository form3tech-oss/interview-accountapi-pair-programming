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

// PostPaymentsIDReversalsReversalIDSubmissionsReader is a Reader for the PostPaymentsIDReversalsReversalIDSubmissions structure.
type PostPaymentsIDReversalsReversalIDSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDReversalsReversalIDSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDReversalsReversalIDSubmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDReversalsReversalIDSubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDReversalsReversalIDSubmissionsCreated creates a PostPaymentsIDReversalsReversalIDSubmissionsCreated with default headers values
func NewPostPaymentsIDReversalsReversalIDSubmissionsCreated() *PostPaymentsIDReversalsReversalIDSubmissionsCreated {
	return &PostPaymentsIDReversalsReversalIDSubmissionsCreated{}
}

/*PostPaymentsIDReversalsReversalIDSubmissionsCreated handles this case with default header values.

Reversal submission creation response
*/
type PostPaymentsIDReversalsReversalIDSubmissionsCreated struct {
	Payload *models.ReversalSubmissionCreationResponse
}

func (o *PostPaymentsIDReversalsReversalIDSubmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals/{reversalId}/submissions][%d] postPaymentsIdReversalsReversalIdSubmissionsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDReversalsReversalIDSubmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReversalSubmissionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDReversalsReversalIDSubmissionsBadRequest creates a PostPaymentsIDReversalsReversalIDSubmissionsBadRequest with default headers values
func NewPostPaymentsIDReversalsReversalIDSubmissionsBadRequest() *PostPaymentsIDReversalsReversalIDSubmissionsBadRequest {
	return &PostPaymentsIDReversalsReversalIDSubmissionsBadRequest{}
}

/*PostPaymentsIDReversalsReversalIDSubmissionsBadRequest handles this case with default header values.

Reversal submission creation error
*/
type PostPaymentsIDReversalsReversalIDSubmissionsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDReversalsReversalIDSubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/reversals/{reversalId}/submissions][%d] postPaymentsIdReversalsReversalIdSubmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDReversalsReversalIDSubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
