// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PaymentDefaultsAttributes payment defaults attributes
// swagger:model paymentDefaultsAttributes
type PaymentDefaultsAttributes struct {

	// default payment scheme
	DefaultPaymentScheme string `json:"default_payment_scheme,omitempty"`
}

// Validate validates this payment defaults attributes
func (m *PaymentDefaultsAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PaymentDefaultsAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentDefaultsAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentDefaultsAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
