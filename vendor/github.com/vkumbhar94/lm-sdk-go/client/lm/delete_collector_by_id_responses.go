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

// DeleteCollectorByIDReader is a Reader for the DeleteCollectorByID structure.
type DeleteCollectorByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCollectorByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCollectorByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteCollectorByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCollectorByIDOK creates a DeleteCollectorByIDOK with default headers values
func NewDeleteCollectorByIDOK() *DeleteCollectorByIDOK {
	return &DeleteCollectorByIDOK{}
}

/*DeleteCollectorByIDOK handles this case with default header values.

successful operation
*/
type DeleteCollectorByIDOK struct {
	Payload interface{}
}

func (o *DeleteCollectorByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /setting/collector/collectors/{id}][%d] deleteCollectorByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteCollectorByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCollectorByIDDefault creates a DeleteCollectorByIDDefault with default headers values
func NewDeleteCollectorByIDDefault(code int) *DeleteCollectorByIDDefault {
	return &DeleteCollectorByIDDefault{
		_statusCode: code,
	}
}

/*DeleteCollectorByIDDefault handles this case with default header values.

Error
*/
type DeleteCollectorByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete collector by Id default response
func (o *DeleteCollectorByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCollectorByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /setting/collector/collectors/{id}][%d] deleteCollectorById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCollectorByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
