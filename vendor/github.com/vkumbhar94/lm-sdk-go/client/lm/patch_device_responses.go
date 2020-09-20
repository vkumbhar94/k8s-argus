// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// PatchDeviceReader is a Reader for the PatchDevice structure.
type PatchDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchDeviceDefault(response.Code())
		if result.Code() == 429 {
                        errResp :=  &models.ErrorResponse{
                                                ErrorCode: 429,
                                                ErrorDetail: map[string]interface{}{
                                                        "x-rate-limit-limit": response.GetHeader("x-rate-limit-limit"),
                                                        "x-rate-limit-remaining": response.GetHeader("x-rate-limit-remaining"),
                                                        "x-rate-limit-window": response.GetHeader("x-rate-limit-window"),
                                                },
                                                ErrorMessage: "Customized response from argus sdk",
                                        }
			result.Payload = errResp
                        return nil, result
                }
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDeviceOK creates a PatchDeviceOK with default headers values
func NewPatchDeviceOK() *PatchDeviceOK {
	return &PatchDeviceOK{}
}

/*PatchDeviceOK handles this case with default header values.

successful operation
*/
type PatchDeviceOK struct {
	Payload *models.Device
}

func (o *PatchDeviceOK) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{id}][%d] patchDeviceOK  %+v", 200, o.Payload)
}

func (o *PatchDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Device)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDeviceDefault creates a PatchDeviceDefault with default headers values
func NewPatchDeviceDefault(code int) *PatchDeviceDefault {
	return &PatchDeviceDefault{
		_statusCode: code,
	}
}

/*PatchDeviceDefault handles this case with default header values.

Error
*/
type PatchDeviceDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch device default response
func (o *PatchDeviceDefault) Code() int {
	return o._statusCode
}

func (o *PatchDeviceDefault) Error() string {
	return fmt.Sprintf("[PATCH /device/devices/{id}][%d] patchDevice default  %+v", o._statusCode, o.Payload)
}

func (o *PatchDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
