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

// ReferenceTrip A reference trip is another real trip, but it doesn't have its own stops and events. It refers only to its ref-erenced regular trip. The reference trip collects mainly all different attributes of the referenced regular trip.
//
// swagger:model referenceTrip
type ReferenceTrip struct {

	// The cancellation flag. True means, the reference trip is cancelled.
	// Required: true
	C *bool `json:"c" xml:"c"`

	// Reference trip stop label of the end arrival event.
	// Required: true
	Ea *ReferenceTripStopLabel `json:"ea"`

	// An id that uniquely identifies the reference trip. It consists of the following two elements separated by dashes:
	//
	// * A 'daily trip id' that uniquely identifies a reference trip within one day. This id is typically reused on subsequent days. This could be negative.
	// * A 10-digit date specifier (YYMMddHHmm) that indicates the planned departure date of the referenced regular trip from its start station.
	//
	// Example:
	//
	// '-7874571842864554321-1403311221' would be used for a trip with daily trip id '-7874571842864554321' that starts on march the 31th 2014.
	//
	// Required: true
	ID *string `json:"id" xml:"id"`

	// Reference trip label.
	// Required: true
	Rtl *ReferenceTripLabel `json:"rtl"`

	// Reference trip stop label of the start departure event.
	// Required: true
	Sd *ReferenceTripStopLabel `json:"sd"`
}

// Validate validates this reference trip
func (m *ReferenceTrip) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateC(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRtl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReferenceTrip) validateC(formats strfmt.Registry) error {

	if err := validate.Required("c", "body", m.C); err != nil {
		return err
	}

	return nil
}

func (m *ReferenceTrip) validateEa(formats strfmt.Registry) error {

	if err := validate.Required("ea", "body", m.Ea); err != nil {
		return err
	}

	if m.Ea != nil {
		if err := m.Ea.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ea")
			}
			return err
		}
	}

	return nil
}

func (m *ReferenceTrip) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ReferenceTrip) validateRtl(formats strfmt.Registry) error {

	if err := validate.Required("rtl", "body", m.Rtl); err != nil {
		return err
	}

	if m.Rtl != nil {
		if err := m.Rtl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rtl")
			}
			return err
		}
	}

	return nil
}

func (m *ReferenceTrip) validateSd(formats strfmt.Registry) error {

	if err := validate.Required("sd", "body", m.Sd); err != nil {
		return err
	}

	if m.Sd != nil {
		if err := m.Sd.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sd")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReferenceTrip) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReferenceTrip) UnmarshalBinary(b []byte) error {
	var res ReferenceTrip
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
