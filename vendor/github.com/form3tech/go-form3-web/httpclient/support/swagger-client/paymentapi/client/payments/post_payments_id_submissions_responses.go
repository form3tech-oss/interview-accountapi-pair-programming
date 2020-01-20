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

// PostPaymentsIDSubmissionsReader is a Reader for the PostPaymentsIDSubmissions structure.
type PostPaymentsIDSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentsIDSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsIDSubmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPaymentsIDSubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsIDSubmissionsCreated creates a PostPaymentsIDSubmissionsCreated with default headers values
func NewPostPaymentsIDSubmissionsCreated() *PostPaymentsIDSubmissionsCreated {
	return &PostPaymentsIDSubmissionsCreated{}
}

/*PostPaymentsIDSubmissionsCreated handles this case with default header values.

Submission creation response
*/
type PostPaymentsIDSubmissionsCreated struct {
	Payload *models.PaymentSubmissionCreationResponse
}

func (o *PostPaymentsIDSubmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsIDSubmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSubmissionCreationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsIDSubmissionsBadRequest creates a PostPaymentsIDSubmissionsBadRequest with default headers values
func NewPostPaymentsIDSubmissionsBadRequest() *PostPaymentsIDSubmissionsBadRequest {
	return &PostPaymentsIDSubmissionsBadRequest{}
}

/*PostPaymentsIDSubmissionsBadRequest handles this case with default header values.

Submission creation error
*/
type PostPaymentsIDSubmissionsBadRequest struct {
	Payload *models.APIError
}

func (o *PostPaymentsIDSubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /payments/{id}/submissions][%d] postPaymentsIdSubmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPaymentsIDSubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
