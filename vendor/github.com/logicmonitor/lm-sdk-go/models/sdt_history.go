// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SDTHistory SDT history
// swagger:model SDTHistory
type SDTHistory struct {

	// The user that added the SDT
	// Read Only: true
	Admin string `json:"admin,omitempty"`

	// The comment associated with the SDT
	// Read Only: true
	Comment string `json:"comment,omitempty"`

	// The duration of the SDT, in minutes
	// Read Only: true
	Duration int64 `json:"duration,omitempty"`

	// The end epoch for the SDT
	// Read Only: true
	EndEpoch int64 `json:"endEpoch,omitempty"`

	// The ID of the SDT
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The ID of the resource in SDT, e.g. the group or device in SDT
	// Read Only: true
	ItemID int32 `json:"itemId,omitempty"`

	// The start epoch for the SDT
	// Read Only: true
	StartEpoch int64 `json:"startEpoch,omitempty"`

	// The SDT type
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this SDT history
func (m *SDTHistory) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SDTHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SDTHistory) UnmarshalBinary(b []byte) error {
	var res SDTHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
