// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/vkumbhar94/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// AddDeviceGroupPropertyReader is a Reader for the AddDeviceGroupProperty structure.
type AddDeviceGroupPropertyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDeviceGroupPropertyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddDeviceGroupPropertyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddDeviceGroupPropertyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddDeviceGroupPropertyOK creates a AddDeviceGroupPropertyOK with default headers values
func NewAddDeviceGroupPropertyOK() *AddDeviceGroupPropertyOK {
	return &AddDeviceGroupPropertyOK{}
}

/*AddDeviceGroupPropertyOK handles this case with default header values.

successful operation
*/
type AddDeviceGroupPropertyOK struct {
	Payload *models.EntityProperty
}

func (o *AddDeviceGroupPropertyOK) Error() string {
	return fmt.Sprintf("[POST /device/groups/{gid}/properties][%d] addDeviceGroupPropertyOK  %+v", 200, o.Payload)
}

func (o *AddDeviceGroupPropertyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntityProperty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDeviceGroupPropertyDefault creates a AddDeviceGroupPropertyDefault with default headers values
func NewAddDeviceGroupPropertyDefault(code int) *AddDeviceGroupPropertyDefault {
	return &AddDeviceGroupPropertyDefault{
		_statusCode: code,
	}
}

/*AddDeviceGroupPropertyDefault handles this case with default header values.

Error
*/
type AddDeviceGroupPropertyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add device group property default response
func (o *AddDeviceGroupPropertyDefault) Code() int {
	return o._statusCode
}

func (o *AddDeviceGroupPropertyDefault) Error() string {
	return fmt.Sprintf("[POST /device/groups/{gid}/properties][%d] addDeviceGroupProperty default  %+v", o._statusCode, o.Payload)
}

func (o *AddDeviceGroupPropertyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
