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

package ipam

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

	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

// NewIpamPrefixesPartialUpdateParams creates a new IpamPrefixesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesPartialUpdateParams() *IpamPrefixesPartialUpdateParams {
	return &IpamPrefixesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesPartialUpdateParamsWithTimeout creates a new IpamPrefixesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamPrefixesPartialUpdateParams {
	return &IpamPrefixesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesPartialUpdateParamsWithContext creates a new IpamPrefixesPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamPrefixesPartialUpdateParamsWithContext(ctx context.Context) *IpamPrefixesPartialUpdateParams {
	return &IpamPrefixesPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamPrefixesPartialUpdateParamsWithHTTPClient creates a new IpamPrefixesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamPrefixesPartialUpdateParams {
	return &IpamPrefixesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamPrefixesPartialUpdateParams contains all the parameters to send to the API endpoint

	for the ipam prefixes partial update operation.

	Typically these are written to a http.Request.
*/
type IpamPrefixesPartialUpdateParams struct {

	// Data.
	Data *models.WritablePrefix

	/* ID.

	   A unique integer value identifying this prefix.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesPartialUpdateParams) WithDefaults() *IpamPrefixesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes partial update params
func (o *IpamPrefixesPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamPrefixesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes partial update params
func (o *IpamPrefixesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes partial update params
func (o *IpamPrefixesPartialUpdateParams) WithContext(ctx context.Context) *IpamPrefixesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes partial update params
func (o *IpamPrefixesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes partial update params
func (o *IpamPrefixesPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamPrefixesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes partial update params
func (o *IpamPrefixesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes partial update params
func (o *IpamPrefixesPartialUpdateParams) WithData(data *models.WritablePrefix) *IpamPrefixesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes partial update params
func (o *IpamPrefixesPartialUpdateParams) SetData(data *models.WritablePrefix) {
	o.Data = data
}

// WithID adds the id to the ipam prefixes partial update params
func (o *IpamPrefixesPartialUpdateParams) WithID(id int64) *IpamPrefixesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes partial update params
func (o *IpamPrefixesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
