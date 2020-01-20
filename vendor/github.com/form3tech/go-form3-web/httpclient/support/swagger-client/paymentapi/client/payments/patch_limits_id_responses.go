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

// PatchLimitsIDReader is a Reader for the PatchLimitsID structure.
type PatchLimitsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLimitsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchLimitsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLimitsIDOK creates a PatchLimitsIDOK with default headers values
func NewPatchLimitsIDOK() *PatchLimitsIDOK {
	return &PatchLimitsIDOK{}
}

/*PatchLimitsIDOK handles this case with default header values.

Limit updated
*/
type PatchLimitsIDOK struct {
	Payload *models.LimitDetailsResponse
}

func (o *PatchLimitsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /limits/{id}][%d] patchLimitsIdOK  %+v", 200, o.Payload)
}

func (o *PatchLimitsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LimitDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
