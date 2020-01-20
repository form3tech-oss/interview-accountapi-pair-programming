// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PaymentRelationshipsPaymentSubmission payment relationships payment submission
// swagger:model paymentRelationshipsPaymentSubmission
type PaymentRelationshipsPaymentSubmission struct {

	// data
	Data PaymentRelationshipsPaymentSubmissionData `json:"data"`
}

// Validate validates this payment relationships payment submission
func (m *PaymentRelationshipsPaymentSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentRelationshipsPaymentSubmission) UnmarshalBinary(b []byte) error {
	var res PaymentRelationshipsPaymentSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
