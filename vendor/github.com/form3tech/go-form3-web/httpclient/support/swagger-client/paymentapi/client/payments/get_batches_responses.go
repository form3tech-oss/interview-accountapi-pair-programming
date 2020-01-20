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

// GetBatchesReader is a Reader for the GetBatches structure.
type GetBatchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBatchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBatchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBatchesOK creates a GetBatchesOK with default headers values
func NewGetBatchesOK() *GetBatchesOK {
	return &GetBatchesOK{}
}

/*GetBatchesOK handles this case with default header values.

List of batch details
*/
type GetBatchesOK struct {
	Payload *models.PaymentBatchDetailsListResponse
}

func (o *GetBatchesOK) Error() string {
	return fmt.Sprintf("[GET /batches][%d] getBatchesOK  %+v", 200, o.Payload)
}

func (o *GetBatchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentBatchDetailsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
