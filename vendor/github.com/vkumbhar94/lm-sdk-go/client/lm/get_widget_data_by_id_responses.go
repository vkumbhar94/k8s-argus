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

// GetWidgetDataByIDReader is a Reader for the GetWidgetDataByID structure.
type GetWidgetDataByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWidgetDataByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWidgetDataByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetWidgetDataByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWidgetDataByIDOK creates a GetWidgetDataByIDOK with default headers values
func NewGetWidgetDataByIDOK() *GetWidgetDataByIDOK {
	return &GetWidgetDataByIDOK{}
}

/*GetWidgetDataByIDOK handles this case with default header values.

successful operation
*/
type GetWidgetDataByIDOK struct {
	Payload models.WidgetData
}

func (o *GetWidgetDataByIDOK) Error() string {
	return fmt.Sprintf("[GET /dashboard/widgets/{id}/data][%d] getWidgetDataByIdOK  %+v", 200, o.Payload)
}

func (o *GetWidgetDataByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalWidgetData(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetWidgetDataByIDDefault creates a GetWidgetDataByIDDefault with default headers values
func NewGetWidgetDataByIDDefault(code int) *GetWidgetDataByIDDefault {
	return &GetWidgetDataByIDDefault{
		_statusCode: code,
	}
}

/*GetWidgetDataByIDDefault handles this case with default header values.

Error
*/
type GetWidgetDataByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get widget data by Id default response
func (o *GetWidgetDataByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWidgetDataByIDDefault) Error() string {
	return fmt.Sprintf("[GET /dashboard/widgets/{id}/data][%d] getWidgetDataById default  %+v", o._statusCode, o.Payload)
}

func (o *GetWidgetDataByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
