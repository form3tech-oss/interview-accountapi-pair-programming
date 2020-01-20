// Code generated by go-swagger; DO NOT EDIT.

package payments

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
)

// NewGetBatchesParams creates a new GetBatchesParams object
// with the default values initialized.
func NewGetBatchesParams() *GetBatchesParams {
	var ()
	return &GetBatchesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBatchesParamsWithTimeout creates a new GetBatchesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBatchesParamsWithTimeout(timeout time.Duration) *GetBatchesParams {
	var ()
	return &GetBatchesParams{

		timeout: timeout,
	}
}

// NewGetBatchesParamsWithContext creates a new GetBatchesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBatchesParamsWithContext(ctx context.Context) *GetBatchesParams {
	var ()
	return &GetBatchesParams{

		Context: ctx,
	}
}

// NewGetBatchesParamsWithHTTPClient creates a new GetBatchesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBatchesParamsWithHTTPClient(client *http.Client) *GetBatchesParams {
	var ()
	return &GetBatchesParams{
		HTTPClient: client,
	}
}

/*GetBatchesParams contains all the parameters to send to the API endpoint
for the get batches operation typically these are written to a http.Request
*/
type GetBatchesParams struct {

	/*FilterCreateDatetimeFrom*/
	FilterCreateDatetimeFrom *strfmt.DateTime
	/*FilterCreateDatetimeTo*/
	FilterCreateDatetimeTo *strfmt.DateTime
	/*FilterFileFormat*/
	FilterFileFormat *string
	/*FilterOrganisationID
	  Filter by organisation id

	*/
	FilterOrganisationID []strfmt.UUID
	/*PageNumber
	  Which page to select

	*/
	PageNumber *string
	/*PageSize
	  Number of items to select

	*/
	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get batches params
func (o *GetBatchesParams) WithTimeout(timeout time.Duration) *GetBatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get batches params
func (o *GetBatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get batches params
func (o *GetBatchesParams) WithContext(ctx context.Context) *GetBatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get batches params
func (o *GetBatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get batches params
func (o *GetBatchesParams) WithHTTPClient(client *http.Client) *GetBatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get batches params
func (o *GetBatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterCreateDatetimeFrom adds the filterCreateDatetimeFrom to the get batches params
func (o *GetBatchesParams) WithFilterCreateDatetimeFrom(filterCreateDatetimeFrom *strfmt.DateTime) *GetBatchesParams {
	o.SetFilterCreateDatetimeFrom(filterCreateDatetimeFrom)
	return o
}

// SetFilterCreateDatetimeFrom adds the filterCreateDatetimeFrom to the get batches params
func (o *GetBatchesParams) SetFilterCreateDatetimeFrom(filterCreateDatetimeFrom *strfmt.DateTime) {
	o.FilterCreateDatetimeFrom = filterCreateDatetimeFrom
}

// WithFilterCreateDatetimeTo adds the filterCreateDatetimeTo to the get batches params
func (o *GetBatchesParams) WithFilterCreateDatetimeTo(filterCreateDatetimeTo *strfmt.DateTime) *GetBatchesParams {
	o.SetFilterCreateDatetimeTo(filterCreateDatetimeTo)
	return o
}

// SetFilterCreateDatetimeTo adds the filterCreateDatetimeTo to the get batches params
func (o *GetBatchesParams) SetFilterCreateDatetimeTo(filterCreateDatetimeTo *strfmt.DateTime) {
	o.FilterCreateDatetimeTo = filterCreateDatetimeTo
}

// WithFilterFileFormat adds the filterFileFormat to the get batches params
func (o *GetBatchesParams) WithFilterFileFormat(filterFileFormat *string) *GetBatchesParams {
	o.SetFilterFileFormat(filterFileFormat)
	return o
}

// SetFilterFileFormat adds the filterFileFormat to the get batches params
func (o *GetBatchesParams) SetFilterFileFormat(filterFileFormat *string) {
	o.FilterFileFormat = filterFileFormat
}

// WithFilterOrganisationID adds the filterOrganisationID to the get batches params
func (o *GetBatchesParams) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *GetBatchesParams {
	o.SetFilterOrganisationID(filterOrganisationID)
	return o
}

// SetFilterOrganisationID adds the filterOrganisationId to the get batches params
func (o *GetBatchesParams) SetFilterOrganisationID(filterOrganisationID []strfmt.UUID) {
	o.FilterOrganisationID = filterOrganisationID
}

// WithPageNumber adds the pageNumber to the get batches params
func (o *GetBatchesParams) WithPageNumber(pageNumber *string) *GetBatchesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get batches params
func (o *GetBatchesParams) SetPageNumber(pageNumber *string) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get batches params
func (o *GetBatchesParams) WithPageSize(pageSize *int64) *GetBatchesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get batches params
func (o *GetBatchesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetBatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterCreateDatetimeFrom != nil {

		// query param filter[create_datetime_from]
		var qrFilterCreateDatetimeFrom strfmt.DateTime
		if o.FilterCreateDatetimeFrom != nil {
			qrFilterCreateDatetimeFrom = *o.FilterCreateDatetimeFrom
		}
		qFilterCreateDatetimeFrom := qrFilterCreateDatetimeFrom.String()
		if qFilterCreateDatetimeFrom != "" {
			if err := r.SetQueryParam("filter[create_datetime_from]", qFilterCreateDatetimeFrom); err != nil {
				return err
			}
		}

	}

	if o.FilterCreateDatetimeTo != nil {

		// query param filter[create_datetime_to]
		var qrFilterCreateDatetimeTo strfmt.DateTime
		if o.FilterCreateDatetimeTo != nil {
			qrFilterCreateDatetimeTo = *o.FilterCreateDatetimeTo
		}
		qFilterCreateDatetimeTo := qrFilterCreateDatetimeTo.String()
		if qFilterCreateDatetimeTo != "" {
			if err := r.SetQueryParam("filter[create_datetime_to]", qFilterCreateDatetimeTo); err != nil {
				return err
			}
		}

	}

	if o.FilterFileFormat != nil {

		// query param filter[file_format]
		var qrFilterFileFormat string
		if o.FilterFileFormat != nil {
			qrFilterFileFormat = *o.FilterFileFormat
		}
		qFilterFileFormat := qrFilterFileFormat
		if qFilterFileFormat != "" {
			if err := r.SetQueryParam("filter[file_format]", qFilterFileFormat); err != nil {
				return err
			}
		}

	}

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber string
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := qrPageNumber
		if qPageNumber != "" {
			if err := r.SetQueryParam("page[number]", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page[size]
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page[size]", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
