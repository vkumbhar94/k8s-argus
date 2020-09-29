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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// NewAddAdminParams creates a new AddAdminParams object
// with the default values initialized.
func NewAddAdminParams() *AddAdminParams {
	var ()
	return &AddAdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAdminParamsWithTimeout creates a new AddAdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAdminParamsWithTimeout(timeout time.Duration) *AddAdminParams {
	var ()
	return &AddAdminParams{

		timeout: timeout,
	}
}

// NewAddAdminParamsWithContext creates a new AddAdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAdminParamsWithContext(ctx context.Context) *AddAdminParams {
	var ()
	return &AddAdminParams{

		Context: ctx,
	}
}

// NewAddAdminParamsWithHTTPClient creates a new AddAdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAdminParamsWithHTTPClient(client *http.Client) *AddAdminParams {
	var ()
	return &AddAdminParams{
		HTTPClient: client,
	}
}

/*AddAdminParams contains all the parameters to send to the API endpoint
for the add admin operation typically these are written to a http.Request
*/
type AddAdminParams struct {

	/*Body*/
	Body *models.Admin

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add admin params
func (o *AddAdminParams) WithTimeout(timeout time.Duration) *AddAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add admin params
func (o *AddAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add admin params
func (o *AddAdminParams) WithContext(ctx context.Context) *AddAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add admin params
func (o *AddAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add admin params
func (o *AddAdminParams) WithHTTPClient(client *http.Client) *AddAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add admin params
func (o *AddAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add admin params
func (o *AddAdminParams) WithBody(body *models.Admin) *AddAdminParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add admin params
func (o *AddAdminParams) SetBody(body *models.Admin) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
