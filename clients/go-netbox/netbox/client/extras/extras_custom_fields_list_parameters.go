// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewExtrasCustomFieldsListParams creates a new ExtrasCustomFieldsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldsListParams() *ExtrasCustomFieldsListParams {
	return &ExtrasCustomFieldsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldsListParamsWithTimeout creates a new ExtrasCustomFieldsListParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldsListParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldsListParams {
	return &ExtrasCustomFieldsListParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldsListParamsWithContext creates a new ExtrasCustomFieldsListParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldsListParamsWithContext(ctx context.Context) *ExtrasCustomFieldsListParams {
	return &ExtrasCustomFieldsListParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldsListParamsWithHTTPClient creates a new ExtrasCustomFieldsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldsListParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldsListParams {
	return &ExtrasCustomFieldsListParams{
		HTTPClient: client,
	}
}

/*
ExtrasCustomFieldsListParams contains all the parameters to send to the API endpoint

	for the extras custom fields list operation.

	Typically these are written to a http.Request.
*/
type ExtrasCustomFieldsListParams struct {

	// ContentTypes.
	ContentTypes *string

	// FilterLogic.
	FilterLogic *string

	// ID.
	ID *float64

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Required.
	Required *string

	// Weight.
	Weight *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom fields list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsListParams) WithDefaults() *ExtrasCustomFieldsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom fields list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithContext(ctx context.Context) *ExtrasCustomFieldsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentTypes adds the contentTypes to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithContentTypes(contentTypes *string) *ExtrasCustomFieldsListParams {
	o.SetContentTypes(contentTypes)
	return o
}

// SetContentTypes adds the contentTypes to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetContentTypes(contentTypes *string) {
	o.ContentTypes = contentTypes
}

// WithFilterLogic adds the filterLogic to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithFilterLogic(filterLogic *string) *ExtrasCustomFieldsListParams {
	o.SetFilterLogic(filterLogic)
	return o
}

// SetFilterLogic adds the filterLogic to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetFilterLogic(filterLogic *string) {
	o.FilterLogic = filterLogic
}

// WithID adds the id to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithID(id *float64) *ExtrasCustomFieldsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetID(id *float64) {
	o.ID = id
}

// WithLimit adds the limit to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithLimit(limit *int64) *ExtrasCustomFieldsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithName(name *string) *ExtrasCustomFieldsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithOffset(offset *int64) *ExtrasCustomFieldsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithRequired adds the required to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithRequired(required *string) *ExtrasCustomFieldsListParams {
	o.SetRequired(required)
	return o
}

// SetRequired adds the required to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetRequired(required *string) {
	o.Required = required
}

// WithWeight adds the weight to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) WithWeight(weight *float64) *ExtrasCustomFieldsListParams {
	o.SetWeight(weight)
	return o
}

// SetWeight adds the weight to the extras custom fields list params
func (o *ExtrasCustomFieldsListParams) SetWeight(weight *float64) {
	o.Weight = weight
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentTypes != nil {

		// query param content_types
		var qrContentTypes string

		if o.ContentTypes != nil {
			qrContentTypes = *o.ContentTypes
		}
		qContentTypes := qrContentTypes
		if qContentTypes != "" {

			if err := r.SetQueryParam("content_types", qContentTypes); err != nil {
				return err
			}
		}
	}

	if o.FilterLogic != nil {

		// query param filter_logic
		var qrFilterLogic string

		if o.FilterLogic != nil {
			qrFilterLogic = *o.FilterLogic
		}
		qFilterLogic := qrFilterLogic
		if qFilterLogic != "" {

			if err := r.SetQueryParam("filter_logic", qFilterLogic); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID float64

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatFloat64(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Required != nil {

		// query param required
		var qrRequired string

		if o.Required != nil {
			qrRequired = *o.Required
		}
		qRequired := qrRequired
		if qRequired != "" {

			if err := r.SetQueryParam("required", qRequired); err != nil {
				return err
			}
		}
	}

	if o.Weight != nil {

		// query param weight
		var qrWeight float64

		if o.Weight != nil {
			qrWeight = *o.Weight
		}
		qWeight := swag.FormatFloat64(qrWeight)
		if qWeight != "" {

			if err := r.SetQueryParam("weight", qWeight); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
