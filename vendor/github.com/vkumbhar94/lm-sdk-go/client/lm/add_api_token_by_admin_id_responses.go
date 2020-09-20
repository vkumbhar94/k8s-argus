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

// AddAPITokenByAdminIDReader is a Reader for the AddAPITokenByAdminID structure.
type AddAPITokenByAdminIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAPITokenByAdminIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddAPITokenByAdminIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddAPITokenByAdminIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddAPITokenByAdminIDOK creates a AddAPITokenByAdminIDOK with default headers values
func NewAddAPITokenByAdminIDOK() *AddAPITokenByAdminIDOK {
	return &AddAPITokenByAdminIDOK{}
}

/*AddAPITokenByAdminIDOK handles this case with default header values.

successful operation
*/
type AddAPITokenByAdminIDOK struct {
	Payload *models.APIToken
}

func (o *AddAPITokenByAdminIDOK) Error() string {
	return fmt.Sprintf("[POST /setting/admins/{adminId}/apitokens][%d] addApiTokenByAdminIdOK  %+v", 200, o.Payload)
}

func (o *AddAPITokenByAdminIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAPITokenByAdminIDDefault creates a AddAPITokenByAdminIDDefault with default headers values
func NewAddAPITokenByAdminIDDefault(code int) *AddAPITokenByAdminIDDefault {
	return &AddAPITokenByAdminIDDefault{
		_statusCode: code,
	}
}

/*AddAPITokenByAdminIDDefault handles this case with default header values.

Error
*/
type AddAPITokenByAdminIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add Api token by admin Id default response
func (o *AddAPITokenByAdminIDDefault) Code() int {
	return o._statusCode
}

func (o *AddAPITokenByAdminIDDefault) Error() string {
	return fmt.Sprintf("[POST /setting/admins/{adminId}/apitokens][%d] addApiTokenByAdminId default  %+v", o._statusCode, o.Payload)
}

func (o *AddAPITokenByAdminIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
