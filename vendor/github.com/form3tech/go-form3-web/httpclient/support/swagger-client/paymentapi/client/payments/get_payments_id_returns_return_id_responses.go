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

// GetPaymentsIDReturnsReturnIDReader is a Reader for the GetPaymentsIDReturnsReturnID structure.
type GetPaymentsIDReturnsReturnIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsIDReturnsReturnIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsIDReturnsReturnIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsIDReturnsReturnIDOK creates a GetPaymentsIDReturnsReturnIDOK with default headers values
func NewGetPaymentsIDReturnsReturnIDOK() *GetPaymentsIDReturnsReturnIDOK {
	return &GetPaymentsIDReturnsReturnIDOK{}
}

/*GetPaymentsIDReturnsReturnIDOK handles this case with default header values.

Return details
*/
type GetPaymentsIDReturnsReturnIDOK struct {
	Payload *models.ReturnDetailsResponse
}

func (o *GetPaymentsIDReturnsReturnIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/{id}/returns/{returnId}][%d] getPaymentsIdReturnsReturnIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsIDReturnsReturnIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
