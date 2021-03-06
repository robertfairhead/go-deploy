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
)

// NewGetDatabasesDatabaseIDDatabaseCredentialsParams creates a new GetDatabasesDatabaseIDDatabaseCredentialsParams object
// with the default values initialized.
func NewGetDatabasesDatabaseIDDatabaseCredentialsParams() *GetDatabasesDatabaseIDDatabaseCredentialsParams {
	var ()
	return &GetDatabasesDatabaseIDDatabaseCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDatabasesDatabaseIDDatabaseCredentialsParamsWithTimeout creates a new GetDatabasesDatabaseIDDatabaseCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDatabasesDatabaseIDDatabaseCredentialsParamsWithTimeout(timeout time.Duration) *GetDatabasesDatabaseIDDatabaseCredentialsParams {
	var ()
	return &GetDatabasesDatabaseIDDatabaseCredentialsParams{

		timeout: timeout,
	}
}

// NewGetDatabasesDatabaseIDDatabaseCredentialsParamsWithContext creates a new GetDatabasesDatabaseIDDatabaseCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDatabasesDatabaseIDDatabaseCredentialsParamsWithContext(ctx context.Context) *GetDatabasesDatabaseIDDatabaseCredentialsParams {
	var ()
	return &GetDatabasesDatabaseIDDatabaseCredentialsParams{

		Context: ctx,
	}
}

// NewGetDatabasesDatabaseIDDatabaseCredentialsParamsWithHTTPClient creates a new GetDatabasesDatabaseIDDatabaseCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDatabasesDatabaseIDDatabaseCredentialsParamsWithHTTPClient(client *http.Client) *GetDatabasesDatabaseIDDatabaseCredentialsParams {
	var ()
	return &GetDatabasesDatabaseIDDatabaseCredentialsParams{
		HTTPClient: client,
	}
}

/*GetDatabasesDatabaseIDDatabaseCredentialsParams contains all the parameters to send to the API endpoint
for the get databases database ID database credentials operation typically these are written to a http.Request
*/
type GetDatabasesDatabaseIDDatabaseCredentialsParams struct {

	/*DatabaseID
	  database_id

	*/
	DatabaseID int64
	/*Page
	  current page of results for pagination

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get databases database ID database credentials params
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) WithTimeout(timeout time.Duration) *GetDatabasesDatabaseIDDatabaseCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get databases database ID database credentials params
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get databases database ID database credentials params
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) WithContext(ctx context.Context) *GetDatabasesDatabaseIDDatabaseCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get databases database ID database credentials params
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get databases database ID database credentials params
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) WithHTTPClient(client *http.Client) *GetDatabasesDatabaseIDDatabaseCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get databases database ID database credentials params
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatabaseID adds the databaseID to the get databases database ID database credentials params
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) WithDatabaseID(databaseID int64) *GetDatabasesDatabaseIDDatabaseCredentialsParams {
	o.SetDatabaseID(databaseID)
	return o
}

// SetDatabaseID adds the databaseId to the get databases database ID database credentials params
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) SetDatabaseID(databaseID int64) {
	o.DatabaseID = databaseID
}

// WithPage adds the page to the get databases database ID database credentials params
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) WithPage(page *int64) *GetDatabasesDatabaseIDDatabaseCredentialsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get databases database ID database credentials params
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetDatabasesDatabaseIDDatabaseCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param database_id
	if err := r.SetPathParam("database_id", swag.FormatInt64(o.DatabaseID)); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
