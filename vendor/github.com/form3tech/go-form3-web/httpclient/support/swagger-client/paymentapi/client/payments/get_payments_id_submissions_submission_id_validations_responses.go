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

// GetPaymentsIDSubmissionsSubmissionIDValidationsReader is a Reader for the GetPaymentsIDSubmissionsSubmissionIDValidations structure.
type GetPaymentsIDSubmissionsSubmissionIDValidationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsIDSubmissionsSubmissionIDValidationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsIDSubmissionsSubmissionIDValidationsOK creates a GetPaymentsIDSubmissionsSubmissionIDValidationsOK with default headers values
func NewGetPaymentsIDSubmissionsSubmissionIDValidationsOK() *GetPaymentsIDSubmissionsSubmissionIDValidationsOK {
	return &GetPaymentsIDSubmissionsSubmissionIDValidationsOK{}
}

/*GetPaymentsIDSubmissionsSubmissionIDValidationsOK handles this case with default header values.

Payment submission validations
*/
type GetPaymentsIDSubmissionsSubmissionIDValidationsOK struct {
	Payload *models.PaymentSubmissionValidationListResponse
}

func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/submissions/{submissionId}/validations][%d] getPaymentsIdSubmissionsSubmissionIdValidationsOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsIDSubmissionsSubmissionIDValidationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentSubmissionValidationListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
