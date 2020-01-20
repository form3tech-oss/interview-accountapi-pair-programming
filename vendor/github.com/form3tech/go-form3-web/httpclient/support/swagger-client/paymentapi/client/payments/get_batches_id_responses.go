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

// GetBatchesIDReader is a Reader for the GetBatchesID structure.
type GetBatchesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBatchesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBatchesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBatchesIDOK creates a GetBatchesIDOK with default headers values
func NewGetBatchesIDOK() *GetBatchesIDOK {
	return &GetBatchesIDOK{}
}

/*GetBatchesIDOK handles this case with default header values.

Batch details
*/
type GetBatchesIDOK struct {
	Payload *models.PaymentBatchDetailsResponse
}

func (o *GetBatchesIDOK) Error() string {
	return fmt.Sprintf("[GET /batches/{id}][%d] getBatchesIdOK  %+v", 200, o.Payload)
}

func (o *GetBatchesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentBatchDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
