/* #nosec */// Code generated by go-swagger; DO NOT EDIT.

package account_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/interview-accountapi-pair-programming/internal/swagger-client/interview-accountapi/models"
)

// DeleteOrganisationAccountsIDReader is a Reader for the DeleteOrganisationAccountsID structure.
type DeleteOrganisationAccountsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganisationAccountsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteOrganisationAccountsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteOrganisationAccountsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteOrganisationAccountsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteOrganisationAccountsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteOrganisationAccountsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOrganisationAccountsIDNoContent creates a DeleteOrganisationAccountsIDNoContent with default headers values
func NewDeleteOrganisationAccountsIDNoContent() *DeleteOrganisationAccountsIDNoContent {
	return &DeleteOrganisationAccountsIDNoContent{}
}

/*DeleteOrganisationAccountsIDNoContent handles this case with default header values.

Account deleted
*/
type DeleteOrganisationAccountsIDNoContent struct {
}

func (o *DeleteOrganisationAccountsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organisation/accounts/{id}][%d] deleteOrganisationAccountsIdNoContent ", 204)
}

func (o *DeleteOrganisationAccountsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganisationAccountsIDBadRequest creates a DeleteOrganisationAccountsIDBadRequest with default headers values
func NewDeleteOrganisationAccountsIDBadRequest() *DeleteOrganisationAccountsIDBadRequest {
	return &DeleteOrganisationAccountsIDBadRequest{}
}

/*DeleteOrganisationAccountsIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteOrganisationAccountsIDBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteOrganisationAccountsIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /organisation/accounts/{id}][%d] deleteOrganisationAccountsIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrganisationAccountsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganisationAccountsIDForbidden creates a DeleteOrganisationAccountsIDForbidden with default headers values
func NewDeleteOrganisationAccountsIDForbidden() *DeleteOrganisationAccountsIDForbidden {
	return &DeleteOrganisationAccountsIDForbidden{}
}

/*DeleteOrganisationAccountsIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteOrganisationAccountsIDForbidden struct {
	Payload *models.APIError
}

func (o *DeleteOrganisationAccountsIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organisation/accounts/{id}][%d] deleteOrganisationAccountsIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrganisationAccountsIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganisationAccountsIDNotFound creates a DeleteOrganisationAccountsIDNotFound with default headers values
func NewDeleteOrganisationAccountsIDNotFound() *DeleteOrganisationAccountsIDNotFound {
	return &DeleteOrganisationAccountsIDNotFound{}
}

/*DeleteOrganisationAccountsIDNotFound handles this case with default header values.

Not Found
*/
type DeleteOrganisationAccountsIDNotFound struct {
	Payload *models.APIError
}

func (o *DeleteOrganisationAccountsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organisation/accounts/{id}][%d] deleteOrganisationAccountsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrganisationAccountsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganisationAccountsIDInternalServerError creates a DeleteOrganisationAccountsIDInternalServerError with default headers values
func NewDeleteOrganisationAccountsIDInternalServerError() *DeleteOrganisationAccountsIDInternalServerError {
	return &DeleteOrganisationAccountsIDInternalServerError{}
}

/*DeleteOrganisationAccountsIDInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteOrganisationAccountsIDInternalServerError struct {
	Payload *models.APIError
}

func (o *DeleteOrganisationAccountsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /organisation/accounts/{id}][%d] deleteOrganisationAccountsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOrganisationAccountsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
