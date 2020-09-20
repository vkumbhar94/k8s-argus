// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// NewAddDeviceGroupPropertyParams creates a new AddDeviceGroupPropertyParams object
// with the default values initialized.
func NewAddDeviceGroupPropertyParams() *AddDeviceGroupPropertyParams {
	var ()
	return &AddDeviceGroupPropertyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddDeviceGroupPropertyParamsWithTimeout creates a new AddDeviceGroupPropertyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddDeviceGroupPropertyParamsWithTimeout(timeout time.Duration) *AddDeviceGroupPropertyParams {
	var ()
	return &AddDeviceGroupPropertyParams{

		timeout: timeout,
	}
}

// NewAddDeviceGroupPropertyParamsWithContext creates a new AddDeviceGroupPropertyParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddDeviceGroupPropertyParamsWithContext(ctx context.Context) *AddDeviceGroupPropertyParams {
	var ()
	return &AddDeviceGroupPropertyParams{

		Context: ctx,
	}
}

// NewAddDeviceGroupPropertyParamsWithHTTPClient creates a new AddDeviceGroupPropertyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddDeviceGroupPropertyParamsWithHTTPClient(client *http.Client) *AddDeviceGroupPropertyParams {
	var ()
	return &AddDeviceGroupPropertyParams{
		HTTPClient: client,
	}
}

/*AddDeviceGroupPropertyParams contains all the parameters to send to the API endpoint
for the add device group property operation typically these are written to a http.Request
*/
type AddDeviceGroupPropertyParams struct {

	/*Body*/
	Body *models.EntityProperty
	/*Gid
	  group ID

	*/
	Gid int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add device group property params
func (o *AddDeviceGroupPropertyParams) WithTimeout(timeout time.Duration) *AddDeviceGroupPropertyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add device group property params
func (o *AddDeviceGroupPropertyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add device group property params
func (o *AddDeviceGroupPropertyParams) WithContext(ctx context.Context) *AddDeviceGroupPropertyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add device group property params
func (o *AddDeviceGroupPropertyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add device group property params
func (o *AddDeviceGroupPropertyParams) WithHTTPClient(client *http.Client) *AddDeviceGroupPropertyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add device group property params
func (o *AddDeviceGroupPropertyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add device group property params
func (o *AddDeviceGroupPropertyParams) WithBody(body *models.EntityProperty) *AddDeviceGroupPropertyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add device group property params
func (o *AddDeviceGroupPropertyParams) SetBody(body *models.EntityProperty) {
	o.Body = body
}

// WithGid adds the gid to the add device group property params
func (o *AddDeviceGroupPropertyParams) WithGid(gid int32) *AddDeviceGroupPropertyParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the add device group property params
func (o *AddDeviceGroupPropertyParams) SetGid(gid int32) {
	o.Gid = gid
}

// WriteToRequest writes these params to a swagger request
func (o *AddDeviceGroupPropertyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param gid
	if err := r.SetPathParam("gid", swag.FormatInt32(o.Gid)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
