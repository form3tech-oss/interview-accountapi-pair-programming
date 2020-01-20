// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReversalSubmissionRelationships reversal submission relationships
// swagger:model ReversalSubmissionRelationships
type ReversalSubmissionRelationships struct {

	// reversal
	Reversal *RelationshipLinks `json:"reversal,omitempty"`

	// validations
	Validations *RelationshipLinks `json:"validations,omitempty"`
}

// Validate validates this reversal submission relationships
func (m *ReversalSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReversal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValidations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalSubmissionRelationships) validateReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.Reversal) { // not required
		return nil
	}

	if m.Reversal != nil {

		if err := m.Reversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reversal")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalSubmissionRelationships) validateValidations(formats strfmt.Registry) error {

	if swag.IsZero(m.Validations) { // not required
		return nil
	}

	if m.Validations != nil {

		if err := m.Validations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res ReversalSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
