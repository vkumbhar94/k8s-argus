// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/vkumbhar94/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Netscan netscan
// swagger:discriminator Netscan method
type Netscan interface {
	runtime.Validatable

	// The ID of the Collector associated with this Netscan
	// Required: true
	Collector() *int32
	SetCollector(*int32)

	// The description of the Collector associated with this Netscan
	CollectorDescription() string
	SetCollectorDescription(string)

	// The ID of the group of the Collector associated with this Netscan
	CollectorGroup() int32
	SetCollectorGroup(int32)

	// The name of the group of the Collector associated with this Netscan
	CollectorGroupName() string
	SetCollectorGroupName(string)

	// The user that created the policy
	Creator() string
	SetCreator(string)

	// The description of the Netscan Policy
	Description() string
	SetDescription(string)

	// Information that determines how duplicate discovered devices should be handled
	// Required: true
	Duplicate() *ExcludeDuplicateIps
	SetDuplicate(*ExcludeDuplicateIps)

	// The group the Netscan policy should belong to
	Group() string
	SetGroup(string)

	// The ID of the Netscan Policy
	ID() int32
	SetID(int32)

	// The method that should be used to discover devices. Options are nmap (ICMP Ping), nec2 (EC2), and script
	// Required: true
	Method() string
	SetMethod(string)

	// The name of the Netscan Policy
	// Required: true
	Name() *string
	SetName(*string)

	// The date and time of the next start time of the scan - displayed as manual if the scan does not run on a schedule
	NextStart() string
	SetNextStart(string)

	// The epoch of the next start time of the scan - displayed as 0 if the scan does not run on a schedule
	NextStartEpoch() int64
	SetNextStartEpoch(int64)

	// The ID of the group the policy belongs to
	NsgID() int32
	SetNsgID(int32)

	// Information related to the recurring execution schedule for the Netscan Policy
	Schedule() *NetScanSchedule
	SetSchedule(*NetScanSchedule)

	// The Id of the device
	Version() int32
	SetVersion(int32)
}

type netscan struct {
	collectorField *int32

	collectorDescriptionField string

	collectorGroupField int32

	collectorGroupNameField string

	creatorField string

	descriptionField string

	duplicateField *ExcludeDuplicateIps

	groupField string

	idField int32

	methodField string

	nameField *string

	nextStartField string

	nextStartEpochField int64

	nsgIdField int32

	scheduleField *NetScanSchedule

	versionField int32
}

// Collector gets the collector of this polymorphic type
func (m *netscan) Collector() *int32 {
	return m.collectorField
}

// SetCollector sets the collector of this polymorphic type
func (m *netscan) SetCollector(val *int32) {
	m.collectorField = val
}

// CollectorDescription gets the collector description of this polymorphic type
func (m *netscan) CollectorDescription() string {
	return m.collectorDescriptionField
}

// SetCollectorDescription sets the collector description of this polymorphic type
func (m *netscan) SetCollectorDescription(val string) {
	m.collectorDescriptionField = val
}

// CollectorGroup gets the collector group of this polymorphic type
func (m *netscan) CollectorGroup() int32 {
	return m.collectorGroupField
}

// SetCollectorGroup sets the collector group of this polymorphic type
func (m *netscan) SetCollectorGroup(val int32) {
	m.collectorGroupField = val
}

// CollectorGroupName gets the collector group name of this polymorphic type
func (m *netscan) CollectorGroupName() string {
	return m.collectorGroupNameField
}

// SetCollectorGroupName sets the collector group name of this polymorphic type
func (m *netscan) SetCollectorGroupName(val string) {
	m.collectorGroupNameField = val
}

// Creator gets the creator of this polymorphic type
func (m *netscan) Creator() string {
	return m.creatorField
}

// SetCreator sets the creator of this polymorphic type
func (m *netscan) SetCreator(val string) {
	m.creatorField = val
}

// Description gets the description of this polymorphic type
func (m *netscan) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *netscan) SetDescription(val string) {
	m.descriptionField = val
}

// Duplicate gets the duplicate of this polymorphic type
func (m *netscan) Duplicate() *ExcludeDuplicateIps {
	return m.duplicateField
}

// SetDuplicate sets the duplicate of this polymorphic type
func (m *netscan) SetDuplicate(val *ExcludeDuplicateIps) {
	m.duplicateField = val
}

// Group gets the group of this polymorphic type
func (m *netscan) Group() string {
	return m.groupField
}

// SetGroup sets the group of this polymorphic type
func (m *netscan) SetGroup(val string) {
	m.groupField = val
}

// ID gets the id of this polymorphic type
func (m *netscan) ID() int32 {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *netscan) SetID(val int32) {
	m.idField = val
}

// Method gets the method of this polymorphic type
func (m *netscan) Method() string {
	return "Netscan"
}

// SetMethod sets the method of this polymorphic type
func (m *netscan) SetMethod(val string) {

}

// Name gets the name of this polymorphic type
func (m *netscan) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *netscan) SetName(val *string) {
	m.nameField = val
}

// NextStart gets the next start of this polymorphic type
func (m *netscan) NextStart() string {
	return m.nextStartField
}

// SetNextStart sets the next start of this polymorphic type
func (m *netscan) SetNextStart(val string) {
	m.nextStartField = val
}

// NextStartEpoch gets the next start epoch of this polymorphic type
func (m *netscan) NextStartEpoch() int64 {
	return m.nextStartEpochField
}

// SetNextStartEpoch sets the next start epoch of this polymorphic type
func (m *netscan) SetNextStartEpoch(val int64) {
	m.nextStartEpochField = val
}

// NsgID gets the nsg Id of this polymorphic type
func (m *netscan) NsgID() int32 {
	return m.nsgIdField
}

// SetNsgID sets the nsg Id of this polymorphic type
func (m *netscan) SetNsgID(val int32) {
	m.nsgIdField = val
}

// Schedule gets the schedule of this polymorphic type
func (m *netscan) Schedule() *NetScanSchedule {
	return m.scheduleField
}

// SetSchedule sets the schedule of this polymorphic type
func (m *netscan) SetSchedule(val *NetScanSchedule) {
	m.scheduleField = val
}

// Version gets the version of this polymorphic type
func (m *netscan) Version() int32 {
	return m.versionField
}

// SetVersion sets the version of this polymorphic type
func (m *netscan) SetVersion(val int32) {
	m.versionField = val
}

// UnmarshalNetscanSlice unmarshals polymorphic slices of Netscan
func UnmarshalNetscanSlice(reader io.Reader, consumer runtime.Consumer) ([]Netscan, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Netscan
	for _, element := range elements {
		obj, err := unmarshalNetscan(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalNetscan unmarshals polymorphic Netscan
func UnmarshalNetscan(reader io.Reader, consumer runtime.Consumer) (Netscan, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalNetscan(data, consumer)
}

func unmarshalNetscan(data []byte, consumer runtime.Consumer) (Netscan, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the method property.
	var getType struct {
		Method string `json:"method"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("method", "body", getType.Method); err != nil {
		return nil, err
	}

	// The value of method is used to determine which type to create and unmarshal the data into
	switch getType.Method {
	case "Netscan":
		var result netscan
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "aws":
		var result AWSNetscan
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "azure":
		var result AzureNetscan
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "gcp":
		var result GCPNetscan
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "nec2":
		var result Ec2Netscan
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "nmap":
		var result NMapNetscan
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "script":
		var result ScriptNetscan
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid method value: %q", getType.Method)

}

// Validate validates this netscan
func (m *netscan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuplicate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *netscan) validateCollector(formats strfmt.Registry) error {

	if err := validate.Required("collector", "body", m.Collector()); err != nil {
		return err
	}

	return nil
}

func (m *netscan) validateDuplicate(formats strfmt.Registry) error {

	if err := validate.Required("duplicate", "body", m.Duplicate()); err != nil {
		return err
	}

	if m.Duplicate() != nil {
		if err := m.Duplicate().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duplicate")
			}
			return err
		}
	}

	return nil
}

func (m *netscan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *netscan) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule()) { // not required
		return nil
	}

	if m.Schedule() != nil {
		if err := m.Schedule().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}
