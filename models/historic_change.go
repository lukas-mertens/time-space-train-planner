// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HistoricChange Historic change.
//
// swagger:model historicChange
type HistoricChange struct {

	// Timestamp. The time, in ten digit 'YYMMddHHmm' format, e.g. '1404011437' for 14:37 on April the 1st of 2014.
	Ts string `json:"ts,omitempty" xml:"ts"`
}

// Validate validates this historic change
func (m *HistoricChange) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoricChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoricChange) UnmarshalBinary(b []byte) error {
	var res HistoricChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
