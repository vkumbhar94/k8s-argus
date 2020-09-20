// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GaugeDataPoint gauge data point
// swagger:model GaugeDataPoint
type GaugeDataPoint struct {

	// aggregate function
	AggregateFunction string `json:"aggregateFunction,omitempty"`

	// data point Id
	DataPointID int32 `json:"dataPointId,omitempty"`

	// data point name
	DataPointName string `json:"dataPointName,omitempty"`

	// data series
	DataSeries string `json:"dataSeries,omitempty"`

	// data source full name
	DataSourceFullName string `json:"dataSourceFullName,omitempty"`

	// data source Id
	DataSourceID int32 `json:"dataSourceId,omitempty"`

	// device display name
	// Required: true
	DeviceDisplayName *string `json:"deviceDisplayName"`

	// device group full path
	// Required: true
	DeviceGroupFullPath *string `json:"deviceGroupFullPath"`

	// instance name
	// Required: true
	InstanceName *string `json:"instanceName"`

	// rpn
	Rpn string `json:"rpn,omitempty"`
}

// Validate validates this gauge data point
func (m *GaugeDataPoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceGroupFullPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GaugeDataPoint) validateDeviceDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("deviceDisplayName", "body", m.DeviceDisplayName); err != nil {
		return err
	}

	return nil
}

func (m *GaugeDataPoint) validateDeviceGroupFullPath(formats strfmt.Registry) error {

	if err := validate.Required("deviceGroupFullPath", "body", m.DeviceGroupFullPath); err != nil {
		return err
	}

	return nil
}

func (m *GaugeDataPoint) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instanceName", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GaugeDataPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GaugeDataPoint) UnmarshalBinary(b []byte) error {
	var res GaugeDataPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
