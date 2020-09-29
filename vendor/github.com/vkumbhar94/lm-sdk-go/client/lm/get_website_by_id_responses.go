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

// GetWebsiteByIDReader is a Reader for the GetWebsiteByID structure.
type GetWebsiteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWebsiteByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetWebsiteByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteByIDOK creates a GetWebsiteByIDOK with default headers values
func NewGetWebsiteByIDOK() *GetWebsiteByIDOK {
	return &GetWebsiteByIDOK{}
}

/*GetWebsiteByIDOK handles this case with default header values.

successful operation
*/
type GetWebsiteByIDOK struct {
	Payload models.Website
}

func (o *GetWebsiteByIDOK) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}][%d] getWebsiteByIdOK  %+v", 200, o.Payload)
}

func (o *GetWebsiteByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalWebsite(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetWebsiteByIDDefault creates a GetWebsiteByIDDefault with default headers values
func NewGetWebsiteByIDDefault(code int) *GetWebsiteByIDDefault {
	return &GetWebsiteByIDDefault{
		_statusCode: code,
	}
}

/*GetWebsiteByIDDefault handles this case with default header values.

Error
*/
type GetWebsiteByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website by Id default response
func (o *GetWebsiteByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteByIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}][%d] getWebsiteById default  %+v", o._statusCode, o.Payload)
}

func (o *GetWebsiteByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
