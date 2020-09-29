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

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// NewUpdateReportGroupByIDParams creates a new UpdateReportGroupByIDParams object
// with the default values initialized.
func NewUpdateReportGroupByIDParams() *UpdateReportGroupByIDParams {
	var ()
	return &UpdateReportGroupByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReportGroupByIDParamsWithTimeout creates a new UpdateReportGroupByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateReportGroupByIDParamsWithTimeout(timeout time.Duration) *UpdateReportGroupByIDParams {
	var ()
	return &UpdateReportGroupByIDParams{

		timeout: timeout,
	}
}

// NewUpdateReportGroupByIDParamsWithContext creates a new UpdateReportGroupByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateReportGroupByIDParamsWithContext(ctx context.Context) *UpdateReportGroupByIDParams {
	var ()
	return &UpdateReportGroupByIDParams{

		Context: ctx,
	}
}

// NewUpdateReportGroupByIDParamsWithHTTPClient creates a new UpdateReportGroupByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateReportGroupByIDParamsWithHTTPClient(client *http.Client) *UpdateReportGroupByIDParams {
	var ()
	return &UpdateReportGroupByIDParams{
		HTTPClient: client,
	}
}

/*UpdateReportGroupByIDParams contains all the parameters to send to the API endpoint
for the update report group by Id operation typically these are written to a http.Request
*/
type UpdateReportGroupByIDParams struct {

	/*Body*/
	Body *models.ReportGroup
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update report group by Id params
func (o *UpdateReportGroupByIDParams) WithTimeout(timeout time.Duration) *UpdateReportGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update report group by Id params
func (o *UpdateReportGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update report group by Id params
func (o *UpdateReportGroupByIDParams) WithContext(ctx context.Context) *UpdateReportGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update report group by Id params
func (o *UpdateReportGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update report group by Id params
func (o *UpdateReportGroupByIDParams) WithHTTPClient(client *http.Client) *UpdateReportGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update report group by Id params
func (o *UpdateReportGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update report group by Id params
func (o *UpdateReportGroupByIDParams) WithBody(body *models.ReportGroup) *UpdateReportGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update report group by Id params
func (o *UpdateReportGroupByIDParams) SetBody(body *models.ReportGroup) {
	o.Body = body
}

// WithID adds the id to the update report group by Id params
func (o *UpdateReportGroupByIDParams) WithID(id int32) *UpdateReportGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update report group by Id params
func (o *UpdateReportGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReportGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
