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

// GetPaymentdefaultsIDReader is a Reader for the GetPaymentdefaultsID structure.
type GetPaymentdefaultsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentdefaultsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentdefaultsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentdefaultsIDOK creates a GetPaymentdefaultsIDOK with default headers values
func NewGetPaymentdefaultsIDOK() *GetPaymentdefaultsIDOK {
	return &GetPaymentdefaultsIDOK{}
}

/*GetPaymentdefaultsIDOK handles this case with default header values.

Payment default details
*/
type GetPaymentdefaultsIDOK struct {
	Payload *models.PaymentDefaultsResponse
}

func (o *GetPaymentdefaultsIDOK) Error() string {
	return fmt.Sprintf("[GET /paymentdefaults/{id}][%d] getPaymentdefaultsIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentdefaultsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentDefaultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
