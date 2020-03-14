// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Timetable A timetable is made of a set of TimetableStops and a potential Disruption.
//
// swagger:model timetable
type Timetable struct {

	// EVA station number.
	Eva int64 `json:"eva,omitempty" xml:"eva"`

	// List of Message.
	M []*Message `json:"m"`

	// List of TimetableStop.
	S []*TimetableStop `json:"s"`

	// Station name.
	Station string `json:"station,omitempty" xml:"station"`
}

// Validate validates this timetable
func (m *Timetable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateM(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Timetable) validateM(formats strfmt.Registry) error {

	if swag.IsZero(m.M) { // not required
		return nil
	}

	for i := 0; i < len(m.M); i++ {
		if swag.IsZero(m.M[i]) { // not required
			continue
		}

		if m.M[i] != nil {
			if err := m.M[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("m" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Timetable) validateS(formats strfmt.Registry) error {

	if swag.IsZero(m.S) { // not required
		return nil
	}

	for i := 0; i < len(m.S); i++ {
		if swag.IsZero(m.S[i]) { // not required
			continue
		}

		if m.S[i] != nil {
			if err := m.S[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("s" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Timetable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Timetable) UnmarshalBinary(b []byte) error {
	var res Timetable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
