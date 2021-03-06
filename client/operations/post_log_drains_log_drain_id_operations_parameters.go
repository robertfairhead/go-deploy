// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aptible/go-deploy/models"
)

// NewPostLogDrainsLogDrainIDOperationsParams creates a new PostLogDrainsLogDrainIDOperationsParams object
// with the default values initialized.
func NewPostLogDrainsLogDrainIDOperationsParams() *PostLogDrainsLogDrainIDOperationsParams {
	var ()
	return &PostLogDrainsLogDrainIDOperationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLogDrainsLogDrainIDOperationsParamsWithTimeout creates a new PostLogDrainsLogDrainIDOperationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLogDrainsLogDrainIDOperationsParamsWithTimeout(timeout time.Duration) *PostLogDrainsLogDrainIDOperationsParams {
	var ()
	return &PostLogDrainsLogDrainIDOperationsParams{

		timeout: timeout,
	}
}

// NewPostLogDrainsLogDrainIDOperationsParamsWithContext creates a new PostLogDrainsLogDrainIDOperationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLogDrainsLogDrainIDOperationsParamsWithContext(ctx context.Context) *PostLogDrainsLogDrainIDOperationsParams {
	var ()
	return &PostLogDrainsLogDrainIDOperationsParams{

		Context: ctx,
	}
}

// NewPostLogDrainsLogDrainIDOperationsParamsWithHTTPClient creates a new PostLogDrainsLogDrainIDOperationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLogDrainsLogDrainIDOperationsParamsWithHTTPClient(client *http.Client) *PostLogDrainsLogDrainIDOperationsParams {
	var ()
	return &PostLogDrainsLogDrainIDOperationsParams{
		HTTPClient: client,
	}
}

/*PostLogDrainsLogDrainIDOperationsParams contains all the parameters to send to the API endpoint
for the post log drains log drain ID operations operation typically these are written to a http.Request
*/
type PostLogDrainsLogDrainIDOperationsParams struct {

	/*AppRequest*/
	AppRequest *models.AppRequest29
	/*LogDrainID
	  log_drain_id

	*/
	LogDrainID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post log drains log drain ID operations params
func (o *PostLogDrainsLogDrainIDOperationsParams) WithTimeout(timeout time.Duration) *PostLogDrainsLogDrainIDOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post log drains log drain ID operations params
func (o *PostLogDrainsLogDrainIDOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post log drains log drain ID operations params
func (o *PostLogDrainsLogDrainIDOperationsParams) WithContext(ctx context.Context) *PostLogDrainsLogDrainIDOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post log drains log drain ID operations params
func (o *PostLogDrainsLogDrainIDOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post log drains log drain ID operations params
func (o *PostLogDrainsLogDrainIDOperationsParams) WithHTTPClient(client *http.Client) *PostLogDrainsLogDrainIDOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post log drains log drain ID operations params
func (o *PostLogDrainsLogDrainIDOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppRequest adds the appRequest to the post log drains log drain ID operations params
func (o *PostLogDrainsLogDrainIDOperationsParams) WithAppRequest(appRequest *models.AppRequest29) *PostLogDrainsLogDrainIDOperationsParams {
	o.SetAppRequest(appRequest)
	return o
}

// SetAppRequest adds the appRequest to the post log drains log drain ID operations params
func (o *PostLogDrainsLogDrainIDOperationsParams) SetAppRequest(appRequest *models.AppRequest29) {
	o.AppRequest = appRequest
}

// WithLogDrainID adds the logDrainID to the post log drains log drain ID operations params
func (o *PostLogDrainsLogDrainIDOperationsParams) WithLogDrainID(logDrainID int64) *PostLogDrainsLogDrainIDOperationsParams {
	o.SetLogDrainID(logDrainID)
	return o
}

// SetLogDrainID adds the logDrainId to the post log drains log drain ID operations params
func (o *PostLogDrainsLogDrainIDOperationsParams) SetLogDrainID(logDrainID int64) {
	o.LogDrainID = logDrainID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLogDrainsLogDrainIDOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppRequest != nil {
		if err := r.SetBodyParam(o.AppRequest); err != nil {
			return err
		}
	}

	// path param log_drain_id
	if err := r.SetPathParam("log_drain_id", swag.FormatInt64(o.LogDrainID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
