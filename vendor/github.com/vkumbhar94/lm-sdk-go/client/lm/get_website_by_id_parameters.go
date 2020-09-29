// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/vkumbhar94/runtime"
	cr "github.com/vkumbhar94/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetWebsiteByIDParams creates a new GetWebsiteByIDParams object
// with the default values initialized.
func NewGetWebsiteByIDParams() *GetWebsiteByIDParams {
	var (
		formatDefault = string("json")
	)
	return &GetWebsiteByIDParams{
		Format: &formatDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebsiteByIDParamsWithTimeout creates a new GetWebsiteByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWebsiteByIDParamsWithTimeout(timeout time.Duration) *GetWebsiteByIDParams {
	var (
		formatDefault = string("json")
	)
	return &GetWebsiteByIDParams{
		Format: &formatDefault,

		timeout: timeout,
	}
}

// NewGetWebsiteByIDParamsWithContext creates a new GetWebsiteByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWebsiteByIDParamsWithContext(ctx context.Context) *GetWebsiteByIDParams {
	var (
		formatDefault = string("json")
	)
	return &GetWebsiteByIDParams{
		Format: &formatDefault,

		Context: ctx,
	}
}

// NewGetWebsiteByIDParamsWithHTTPClient creates a new GetWebsiteByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWebsiteByIDParamsWithHTTPClient(client *http.Client) *GetWebsiteByIDParams {
	var (
		formatDefault = string("json")
	)
	return &GetWebsiteByIDParams{
		Format:     &formatDefault,
		HTTPClient: client,
	}
}

/*GetWebsiteByIDParams contains all the parameters to send to the API endpoint
for the get website by Id operation typically these are written to a http.Request
*/
type GetWebsiteByIDParams struct {

	/*Format*/
	Format *string
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get website by Id params
func (o *GetWebsiteByIDParams) WithTimeout(timeout time.Duration) *GetWebsiteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get website by Id params
func (o *GetWebsiteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get website by Id params
func (o *GetWebsiteByIDParams) WithContext(ctx context.Context) *GetWebsiteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get website by Id params
func (o *GetWebsiteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get website by Id params
func (o *GetWebsiteByIDParams) WithHTTPClient(client *http.Client) *GetWebsiteByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get website by Id params
func (o *GetWebsiteByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the get website by Id params
func (o *GetWebsiteByIDParams) WithFormat(format *string) *GetWebsiteByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get website by Id params
func (o *GetWebsiteByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get website by Id params
func (o *GetWebsiteByIDParams) WithID(id int32) *GetWebsiteByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get website by Id params
func (o *GetWebsiteByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsiteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
