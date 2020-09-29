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

// GetNetscanListReader is a Reader for the GetNetscanList structure.
type GetNetscanListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetscanListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetscanListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNetscanListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetscanListOK creates a GetNetscanListOK with default headers values
func NewGetNetscanListOK() *GetNetscanListOK {
	return &GetNetscanListOK{}
}

/*GetNetscanListOK handles this case with default header values.

successful operation
*/
type GetNetscanListOK struct {
	Payload *models.NetscanPaginationResponse
}

func (o *GetNetscanListOK) Error() string {
	return fmt.Sprintf("[GET /setting/netscans][%d] getNetscanListOK  %+v", 200, o.Payload)
}

func (o *GetNetscanListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetscanPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetscanListDefault creates a GetNetscanListDefault with default headers values
func NewGetNetscanListDefault(code int) *GetNetscanListDefault {
	return &GetNetscanListDefault{
		_statusCode: code,
	}
}

/*GetNetscanListDefault handles this case with default header values.

Error
*/
type GetNetscanListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get netscan list default response
func (o *GetNetscanListDefault) Code() int {
	return o._statusCode
}

func (o *GetNetscanListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/netscans][%d] getNetscanList default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetscanListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
