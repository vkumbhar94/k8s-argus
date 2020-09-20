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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// NewAddOpsNoteParams creates a new AddOpsNoteParams object
// with the default values initialized.
func NewAddOpsNoteParams() *AddOpsNoteParams {
	var ()
	return &AddOpsNoteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddOpsNoteParamsWithTimeout creates a new AddOpsNoteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddOpsNoteParamsWithTimeout(timeout time.Duration) *AddOpsNoteParams {
	var ()
	return &AddOpsNoteParams{

		timeout: timeout,
	}
}

// NewAddOpsNoteParamsWithContext creates a new AddOpsNoteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddOpsNoteParamsWithContext(ctx context.Context) *AddOpsNoteParams {
	var ()
	return &AddOpsNoteParams{

		Context: ctx,
	}
}

// NewAddOpsNoteParamsWithHTTPClient creates a new AddOpsNoteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddOpsNoteParamsWithHTTPClient(client *http.Client) *AddOpsNoteParams {
	var ()
	return &AddOpsNoteParams{
		HTTPClient: client,
	}
}

/*AddOpsNoteParams contains all the parameters to send to the API endpoint
for the add ops note operation typically these are written to a http.Request
*/
type AddOpsNoteParams struct {

	/*Body*/
	Body *models.OpsNote

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add ops note params
func (o *AddOpsNoteParams) WithTimeout(timeout time.Duration) *AddOpsNoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add ops note params
func (o *AddOpsNoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add ops note params
func (o *AddOpsNoteParams) WithContext(ctx context.Context) *AddOpsNoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add ops note params
func (o *AddOpsNoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add ops note params
func (o *AddOpsNoteParams) WithHTTPClient(client *http.Client) *AddOpsNoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add ops note params
func (o *AddOpsNoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add ops note params
func (o *AddOpsNoteParams) WithBody(body *models.OpsNote) *AddOpsNoteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add ops note params
func (o *AddOpsNoteParams) SetBody(body *models.OpsNote) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddOpsNoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
