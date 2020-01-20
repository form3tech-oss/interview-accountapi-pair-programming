// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReturnSubmissionValidationAttributes return submission validation attributes
// swagger:model returnSubmissionValidationAttributes
type ReturnSubmissionValidationAttributes struct {

	// response
	Response string `json:"response,omitempty"`

	// skip limit check
	SkipLimitCheck *bool `json:"skip_limit_check,omitempty"`

	// source
	Source ValidationSource `json:"source,omitempty"`

	// status
	Status ValidationStatus `json:"status,omitempty"`
}

// Validate validates this return submission validation attributes
func (m *ReturnSubmissionValidationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnSubmissionValidationAttributes) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if err := m.Source.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source")
		}
		return err
	}

	return nil
}

func (m *ReturnSubmissionValidationAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnSubmissionValidationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionValidationAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionValidationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
