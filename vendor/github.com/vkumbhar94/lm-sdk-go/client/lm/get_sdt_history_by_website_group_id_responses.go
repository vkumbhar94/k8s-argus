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

// GetSDTHistoryByWebsiteGroupIDReader is a Reader for the GetSDTHistoryByWebsiteGroupID structure.
type GetSDTHistoryByWebsiteGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSDTHistoryByWebsiteGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSDTHistoryByWebsiteGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSDTHistoryByWebsiteGroupIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSDTHistoryByWebsiteGroupIDOK creates a GetSDTHistoryByWebsiteGroupIDOK with default headers values
func NewGetSDTHistoryByWebsiteGroupIDOK() *GetSDTHistoryByWebsiteGroupIDOK {
	return &GetSDTHistoryByWebsiteGroupIDOK{}
}

/*GetSDTHistoryByWebsiteGroupIDOK handles this case with default header values.

successful operation
*/
type GetSDTHistoryByWebsiteGroupIDOK struct {
	Payload *models.WebsiteGroupSDTHistoryPaginationResponse
}

func (o *GetSDTHistoryByWebsiteGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /website/groups/{id}/historysdts][%d] getSdtHistoryByWebsiteGroupIdOK  %+v", 200, o.Payload)
}

func (o *GetSDTHistoryByWebsiteGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebsiteGroupSDTHistoryPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSDTHistoryByWebsiteGroupIDDefault creates a GetSDTHistoryByWebsiteGroupIDDefault with default headers values
func NewGetSDTHistoryByWebsiteGroupIDDefault(code int) *GetSDTHistoryByWebsiteGroupIDDefault {
	return &GetSDTHistoryByWebsiteGroupIDDefault{
		_statusCode: code,
	}
}

/*GetSDTHistoryByWebsiteGroupIDDefault handles this case with default header values.

Error
*/
type GetSDTHistoryByWebsiteGroupIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get SDT history by website group Id default response
func (o *GetSDTHistoryByWebsiteGroupIDDefault) Code() int {
	return o._statusCode
}

func (o *GetSDTHistoryByWebsiteGroupIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/groups/{id}/historysdts][%d] getSDTHistoryByWebsiteGroupId default  %+v", o._statusCode, o.Payload)
}

func (o *GetSDTHistoryByWebsiteGroupIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
