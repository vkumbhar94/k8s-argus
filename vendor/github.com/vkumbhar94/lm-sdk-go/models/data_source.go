// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/vkumbhar94/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataSource data source
// swagger:model DataSource
type DataSource struct {

	// applies to
	AppliesTo string `json:"appliesTo,omitempty"`

	// audit version
	// Read Only: true
	AuditVersion int64 `json:"auditVersion,omitempty"`

	// auto discovery config
	AutoDiscoveryConfig *AutoDiscoveryConfiguration `json:"autoDiscoveryConfig,omitempty"`

	// collect interval
	// Required: true
	CollectInterval *int32 `json:"collectInterval"`

	// collect method
	// Required: true
	CollectMethod *string `json:"collectMethod"`

	collectorAttributeField CollectorAttribute

	// data points
	DataPoints []*DataPoint `json:"dataPoints,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// enable auto discovery
	EnableAutoDiscovery bool `json:"enableAutoDiscovery,omitempty"`

	// enable eri discovery
	EnableEriDiscovery bool `json:"enableEriDiscovery,omitempty"`

	// eri discovery config
	EriDiscoveryConfig *ScriptERIdiscoveryAttributeV2 `json:"eriDiscoveryConfig,omitempty"`

	// eri discovery interval
	EriDiscoveryInterval int32 `json:"eriDiscoveryInterval,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// has multi instances
	// Read Only: true
	HasMultiInstances *bool `json:"hasMultiInstances,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// tags
	Tags string `json:"tags,omitempty"`

	// technology
	Technology string `json:"technology,omitempty"`

	// version
	// Read Only: true
	Version int64 `json:"version,omitempty"`
}

// CollectorAttribute gets the collector attribute of this base type
func (m *DataSource) CollectorAttribute() CollectorAttribute {
	return m.collectorAttributeField
}

// SetCollectorAttribute sets the collector attribute of this base type
func (m *DataSource) SetCollectorAttribute(val CollectorAttribute) {
	m.collectorAttributeField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DataSource) UnmarshalJSON(raw []byte) error {
	var data struct {
		AppliesTo string `json:"appliesTo,omitempty"`

		AuditVersion int64 `json:"auditVersion,omitempty"`

		AutoDiscoveryConfig *AutoDiscoveryConfiguration `json:"autoDiscoveryConfig,omitempty"`

		CollectInterval *int32 `json:"collectInterval"`

		CollectMethod *string `json:"collectMethod"`

		CollectorAttribute json.RawMessage `json:"collectorAttribute"`

		DataPoints []*DataPoint `json:"dataPoints,omitempty"`

		Description string `json:"description,omitempty"`

		DisplayName string `json:"displayName,omitempty"`

		EnableAutoDiscovery bool `json:"enableAutoDiscovery,omitempty"`

		EnableEriDiscovery bool `json:"enableEriDiscovery,omitempty"`

		EriDiscoveryConfig *ScriptERIdiscoveryAttributeV2 `json:"eriDiscoveryConfig,omitempty"`

		EriDiscoveryInterval int32 `json:"eriDiscoveryInterval,omitempty"`

		Group string `json:"group,omitempty"`

		HasMultiInstances *bool `json:"hasMultiInstances,omitempty"`

		ID int32 `json:"id,omitempty"`

		Name *string `json:"name"`

		Tags string `json:"tags,omitempty"`

		Technology string `json:"technology,omitempty"`

		Version int64 `json:"version,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	propCollectorAttribute, err := UnmarshalCollectorAttribute(bytes.NewBuffer(data.CollectorAttribute), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result DataSource

	// appliesTo
	result.AppliesTo = data.AppliesTo

	// auditVersion
	result.AuditVersion = data.AuditVersion

	// autoDiscoveryConfig
	result.AutoDiscoveryConfig = data.AutoDiscoveryConfig

	// collectInterval
	result.CollectInterval = data.CollectInterval

	// collectMethod
	result.CollectMethod = data.CollectMethod

	// collectorAttribute
	result.collectorAttributeField = propCollectorAttribute

	// dataPoints
	result.DataPoints = data.DataPoints

	// description
	result.Description = data.Description

	// displayName
	result.DisplayName = data.DisplayName

	// enableAutoDiscovery
	result.EnableAutoDiscovery = data.EnableAutoDiscovery

	// enableEriDiscovery
	result.EnableEriDiscovery = data.EnableEriDiscovery

	// eriDiscoveryConfig
	result.EriDiscoveryConfig = data.EriDiscoveryConfig

	// eriDiscoveryInterval
	result.EriDiscoveryInterval = data.EriDiscoveryInterval

	// group
	result.Group = data.Group

	// hasMultiInstances
	result.HasMultiInstances = data.HasMultiInstances

	// id
	result.ID = data.ID

	// name
	result.Name = data.Name

	// tags
	result.Tags = data.Tags

	// technology
	result.Technology = data.Technology

	// version
	result.Version = data.Version

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DataSource) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AppliesTo string `json:"appliesTo,omitempty"`

		AuditVersion int64 `json:"auditVersion,omitempty"`

		AutoDiscoveryConfig *AutoDiscoveryConfiguration `json:"autoDiscoveryConfig,omitempty"`

		CollectInterval *int32 `json:"collectInterval"`

		CollectMethod *string `json:"collectMethod"`

		DataPoints []*DataPoint `json:"dataPoints,omitempty"`

		Description string `json:"description,omitempty"`

		DisplayName string `json:"displayName,omitempty"`

		EnableAutoDiscovery bool `json:"enableAutoDiscovery,omitempty"`

		EnableEriDiscovery bool `json:"enableEriDiscovery,omitempty"`

		EriDiscoveryConfig *ScriptERIdiscoveryAttributeV2 `json:"eriDiscoveryConfig,omitempty"`

		EriDiscoveryInterval int32 `json:"eriDiscoveryInterval,omitempty"`

		Group string `json:"group,omitempty"`

		HasMultiInstances *bool `json:"hasMultiInstances,omitempty"`

		ID int32 `json:"id,omitempty"`

		Name *string `json:"name"`

		Tags string `json:"tags,omitempty"`

		Technology string `json:"technology,omitempty"`

		Version int64 `json:"version,omitempty"`
	}{

		AppliesTo: m.AppliesTo,

		AuditVersion: m.AuditVersion,

		AutoDiscoveryConfig: m.AutoDiscoveryConfig,

		CollectInterval: m.CollectInterval,

		CollectMethod: m.CollectMethod,

		DataPoints: m.DataPoints,

		Description: m.Description,

		DisplayName: m.DisplayName,

		EnableAutoDiscovery: m.EnableAutoDiscovery,

		EnableEriDiscovery: m.EnableEriDiscovery,

		EriDiscoveryConfig: m.EriDiscoveryConfig,

		EriDiscoveryInterval: m.EriDiscoveryInterval,

		Group: m.Group,

		HasMultiInstances: m.HasMultiInstances,

		ID: m.ID,

		Name: m.Name,

		Tags: m.Tags,

		Technology: m.Technology,

		Version: m.Version,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		CollectorAttribute CollectorAttribute `json:"collectorAttribute"`
	}{

		CollectorAttribute: m.collectorAttributeField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this data source
func (m *DataSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoDiscoveryConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectorAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEriDiscoveryConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSource) validateAutoDiscoveryConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoDiscoveryConfig) { // not required
		return nil
	}

	if m.AutoDiscoveryConfig != nil {
		if err := m.AutoDiscoveryConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoDiscoveryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *DataSource) validateCollectInterval(formats strfmt.Registry) error {

	if err := validate.Required("collectInterval", "body", m.CollectInterval); err != nil {
		return err
	}

	return nil
}

func (m *DataSource) validateCollectMethod(formats strfmt.Registry) error {

	if err := validate.Required("collectMethod", "body", m.CollectMethod); err != nil {
		return err
	}

	return nil
}

func (m *DataSource) validateCollectorAttribute(formats strfmt.Registry) error {

	if err := m.CollectorAttribute().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("collectorAttribute")
		}
		return err
	}

	return nil
}

func (m *DataSource) validateDataPoints(formats strfmt.Registry) error {

	if swag.IsZero(m.DataPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.DataPoints); i++ {
		if swag.IsZero(m.DataPoints[i]) { // not required
			continue
		}

		if m.DataPoints[i] != nil {
			if err := m.DataPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DataSource) validateEriDiscoveryConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.EriDiscoveryConfig) { // not required
		return nil
	}

	if m.EriDiscoveryConfig != nil {
		if err := m.EriDiscoveryConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eriDiscoveryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *DataSource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSource) UnmarshalBinary(b []byte) error {
	var res DataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
