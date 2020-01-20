/* #nosec */// Code generated by go-swagger; DO NOT EDIT.

package account_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/interview-accountapi-pair-programming/internal/swagger-client/interview-accountapi-pair-programming/models"
)

// GetOrganisationAccountsIDReader is a Reader for the GetOrganisationAccountsID structure.
type GetOrganisationAccountsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganisationAccountsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganisationAccountsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrganisationAccountsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrganisationAccountsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganisationAccountsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetOrganisationAccountsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganisationAccountsIDOK creates a GetOrganisationAccountsIDOK with default headers values
func NewGetOrganisationAccountsIDOK() *GetOrganisationAccountsIDOK {
	return &GetOrganisationAccountsIDOK{}
}

/*GetOrganisationAccountsIDOK handles this case with default header values.

Accounts details
*/
type GetOrganisationAccountsIDOK struct {
	Payload *models.AccountDetailsResponse
}

func (o *GetOrganisationAccountsIDOK) Error() string {
	return fmt.Sprintf("[GET /organisation/accounts/{id}][%d] getOrganisationAccountsIdOK  %+v", 200, o.Payload)
}

func (o *GetOrganisationAccountsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationAccountsIDBadRequest creates a GetOrganisationAccountsIDBadRequest with default headers values
func NewGetOrganisationAccountsIDBadRequest() *GetOrganisationAccountsIDBadRequest {
	return &GetOrganisationAccountsIDBadRequest{}
}

/*GetOrganisationAccountsIDBadRequest handles this case with default header values.

Bad Request
*/
type GetOrganisationAccountsIDBadRequest struct {
	Payload *models.APIError
}

func (o *GetOrganisationAccountsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /organisation/accounts/{id}][%d] getOrganisationAccountsIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganisationAccountsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationAccountsIDForbidden creates a GetOrganisationAccountsIDForbidden with default headers values
func NewGetOrganisationAccountsIDForbidden() *GetOrganisationAccountsIDForbidden {
	return &GetOrganisationAccountsIDForbidden{}
}

/*GetOrganisationAccountsIDForbidden handles this case with default header values.

Forbidden
*/
type GetOrganisationAccountsIDForbidden struct {
	Payload *models.APIError
}

func (o *GetOrganisationAccountsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /organisation/accounts/{id}][%d] getOrganisationAccountsIdForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganisationAccountsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationAccountsIDNotFound creates a GetOrganisationAccountsIDNotFound with default headers values
func NewGetOrganisationAccountsIDNotFound() *GetOrganisationAccountsIDNotFound {
	return &GetOrganisationAccountsIDNotFound{}
}

/*GetOrganisationAccountsIDNotFound handles this case with default header values.

Not Found
*/
type GetOrganisationAccountsIDNotFound struct {
	Payload *models.APIError
}

func (o *GetOrganisationAccountsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /organisation/accounts/{id}][%d] getOrganisationAccountsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganisationAccountsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganisationAccountsIDInternalServerError creates a GetOrganisationAccountsIDInternalServerError with default headers values
func NewGetOrganisationAccountsIDInternalServerError() *GetOrganisationAccountsIDInternalServerError {
	return &GetOrganisationAccountsIDInternalServerError{}
}

/*GetOrganisationAccountsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetOrganisationAccountsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *GetOrganisationAccountsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organisation/accounts/{id}][%d] getOrganisationAccountsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganisationAccountsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
