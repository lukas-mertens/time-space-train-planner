// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TripLabel It's a compound data type that contains common data items that characterize a Trip. The contents is represented as a compact 6-tuple in XML.
//
// swagger:model tripLabel
type TripLabel struct {

	// Category. Trip category, e.g. "ICE" or "RE".
	// Required: true
	C *string `json:"c" xml:"c"`

	// Filter flags.
	F string `json:"f,omitempty" xml:"f"`

	// Trip/train number, e.g. "4523".
	// Required: true
	N *string `json:"n" xml:"n"`

	// Owner. A unique short-form and only intended to map a trip to specific evu.
	// Required: true
	O *string `json:"o" xml:"o"`

	// Trip type.
	T TripType `json:"t,omitempty" xml:"t"`
}

// Validate validates this trip label
func (m *TripLabel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateC(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateO(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateT(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TripLabel) validateC(formats strfmt.Registry) error {

	if err := validate.Required("c", "body", m.C); err != nil {
		return err
	}

	return nil
}

func (m *TripLabel) validateN(formats strfmt.Registry) error {

	if err := validate.Required("n", "body", m.N); err != nil {
		return err
	}

	return nil
}

func (m *TripLabel) validateO(formats strfmt.Registry) error {

	if err := validate.Required("o", "body", m.O); err != nil {
		return err
	}

	return nil
}

func (m *TripLabel) validateT(formats strfmt.Registry) error {

	if swag.IsZero(m.T) { // not required
		return nil
	}

	if err := m.T.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("t")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TripLabel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TripLabel) UnmarshalBinary(b []byte) error {
	var res TripLabel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
