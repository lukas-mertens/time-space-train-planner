// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TripReference It's a reference to another trip, which holds its label and reference trips, if available.
//
// swagger:model tripReference
type TripReference struct {

	// The referred trips reference trip elements.
	Rt []*TripLabel `json:"rt"`

	// The referred trips label.
	// Required: true
	Tl *TripLabel `json:"tl"`
}

// Validate validates this trip reference
func (m *TripReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TripReference) validateRt(formats strfmt.Registry) error {

	if swag.IsZero(m.Rt) { // not required
		return nil
	}

	for i := 0; i < len(m.Rt); i++ {
		if swag.IsZero(m.Rt[i]) { // not required
			continue
		}

		if m.Rt[i] != nil {
			if err := m.Rt[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rt" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TripReference) validateTl(formats strfmt.Registry) error {

	if err := validate.Required("tl", "body", m.Tl); err != nil {
		return err
	}

	if m.Tl != nil {
		if err := m.Tl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tl")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TripReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TripReference) UnmarshalBinary(b []byte) error {
	var res TripReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
