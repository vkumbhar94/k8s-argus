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

// NewAddAlertRuleParams creates a new AddAlertRuleParams object
// with the default values initialized.
func NewAddAlertRuleParams() *AddAlertRuleParams {
	var ()
	return &AddAlertRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAlertRuleParamsWithTimeout creates a new AddAlertRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAlertRuleParamsWithTimeout(timeout time.Duration) *AddAlertRuleParams {
	var ()
	return &AddAlertRuleParams{

		timeout: timeout,
	}
}

// NewAddAlertRuleParamsWithContext creates a new AddAlertRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAlertRuleParamsWithContext(ctx context.Context) *AddAlertRuleParams {
	var ()
	return &AddAlertRuleParams{

		Context: ctx,
	}
}

// NewAddAlertRuleParamsWithHTTPClient creates a new AddAlertRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAlertRuleParamsWithHTTPClient(client *http.Client) *AddAlertRuleParams {
	var ()
	return &AddAlertRuleParams{
		HTTPClient: client,
	}
}

/*AddAlertRuleParams contains all the parameters to send to the API endpoint
for the add alert rule operation typically these are written to a http.Request
*/
type AddAlertRuleParams struct {

	/*Body*/
	Body *models.AlertRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add alert rule params
func (o *AddAlertRuleParams) WithTimeout(timeout time.Duration) *AddAlertRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add alert rule params
func (o *AddAlertRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add alert rule params
func (o *AddAlertRuleParams) WithContext(ctx context.Context) *AddAlertRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add alert rule params
func (o *AddAlertRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add alert rule params
func (o *AddAlertRuleParams) WithHTTPClient(client *http.Client) *AddAlertRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add alert rule params
func (o *AddAlertRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add alert rule params
func (o *AddAlertRuleParams) WithBody(body *models.AlertRule) *AddAlertRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add alert rule params
func (o *AddAlertRuleParams) SetBody(body *models.AlertRule) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAlertRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
