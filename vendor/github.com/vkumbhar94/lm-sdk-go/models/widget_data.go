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
	"github.com/go-openapi/validate"
)

// WidgetData widget data
// swagger:discriminator WidgetData type
type WidgetData interface {
	runtime.Validatable

	// title
	Title() string
	SetTitle(string)

	// type
	// Read Only: true
	Type() string
	SetType(string)
}

type widgetData struct {
	titleField string

	typeField string
}

// Title gets the title of this polymorphic type
func (m *widgetData) Title() string {
	return m.titleField
}

// SetTitle sets the title of this polymorphic type
func (m *widgetData) SetTitle(val string) {
	m.titleField = val
}

// Type gets the type of this polymorphic type
func (m *widgetData) Type() string {
	return "WidgetData"
}

// SetType sets the type of this polymorphic type
func (m *widgetData) SetType(val string) {

}

// UnmarshalWidgetDataSlice unmarshals polymorphic slices of WidgetData
func UnmarshalWidgetDataSlice(reader io.Reader, consumer runtime.Consumer) ([]WidgetData, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []WidgetData
	for _, element := range elements {
		obj, err := unmarshalWidgetData(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalWidgetData unmarshals polymorphic WidgetData
func UnmarshalWidgetData(reader io.Reader, consumer runtime.Consumer) (WidgetData, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalWidgetData(data, consumer)
}

func unmarshalWidgetData(data []byte, consumer runtime.Consumer) (WidgetData, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "WidgetData":
		var result widgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "alert":
		var result AlertWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "batchjob":
		var result BatchJobWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "bigNumber":
		var result BigNumberWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "deviceSLA":
		var result DeviceSLAWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "dynamicTable":
		var result DynamicTableWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "gauge":
		var result GaugeWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "gmap":
		var result GoogleMapWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "graph":
		var result GraphPlot
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "groupNetflow":
		var result NetflowGroupWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "netflow":
		var result NetflowWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "noc":
		var result NOCWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "pieChart":
		var result PieChartWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "table":
		var result TableWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "websiteSLA":
		var result WebsiteSLAWidgetData
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)

}

// Validate validates this widget data
func (m *widgetData) Validate(formats strfmt.Registry) error {
	return nil
}
