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

// Connection It's information about a connected train at a particular stop.
//
// swagger:model connection
type Connection struct {

	// Connection status.
	// Required: true
	Cs ConnectionStatus `json:"cs" xml:"cs"`

	// EVA station number.
	Eva int64 `json:"eva,omitempty" xml:"eva"`

	// Id.
	// Required: true
	ID *string `json:"id" xml:"id"`

	// Timetable stop of missed trip.
	Ref *TimetableStop `json:"ref,omitempty"`

	// Timetable stop.
	// Required: true
	S *TimetableStop `json:"s"`

	// Time stamp. The time, in ten digit 'YYMMddHHmm' format, e.g. '1404011437' for 14:37 on April the 1st of 2014.
	// Required: true
	Ts *string `json:"ts" xml:"ts"`
}

// Validate validates this connection
func (m *Connection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Connection) validateCs(formats strfmt.Registry) error {

	if err := m.Cs.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cs")
		}
		return err
	}

	return nil
}

func (m *Connection) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Connection) validateRef(formats strfmt.Registry) error {

	if swag.IsZero(m.Ref) { // not required
		return nil
	}

	if m.Ref != nil {
		if err := m.Ref.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ref")
			}
			return err
		}
	}

	return nil
}

func (m *Connection) validateS(formats strfmt.Registry) error {

	if err := validate.Required("s", "body", m.S); err != nil {
		return err
	}

	if m.S != nil {
		if err := m.S.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s")
			}
			return err
		}
	}

	return nil
}

func (m *Connection) validateTs(formats strfmt.Registry) error {

	if err := validate.Required("ts", "body", m.Ts); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Connection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Connection) UnmarshalBinary(b []byte) error {
	var res Connection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
