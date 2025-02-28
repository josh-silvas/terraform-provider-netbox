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

package dcim

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

	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

// NewDcimRackRolesBulkPartialUpdateParams creates a new DcimRackRolesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackRolesBulkPartialUpdateParams() *DcimRackRolesBulkPartialUpdateParams {
	return &DcimRackRolesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackRolesBulkPartialUpdateParamsWithTimeout creates a new DcimRackRolesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRackRolesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRackRolesBulkPartialUpdateParams {
	return &DcimRackRolesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRackRolesBulkPartialUpdateParamsWithContext creates a new DcimRackRolesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRackRolesBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimRackRolesBulkPartialUpdateParams {
	return &DcimRackRolesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRackRolesBulkPartialUpdateParamsWithHTTPClient creates a new DcimRackRolesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackRolesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRackRolesBulkPartialUpdateParams {
	return &DcimRackRolesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimRackRolesBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim rack roles bulk partial update operation.

	Typically these are written to a http.Request.
*/
type DcimRackRolesBulkPartialUpdateParams struct {

	// Data.
	Data *models.RackRole

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack roles bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesBulkPartialUpdateParams) WithDefaults() *DcimRackRolesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack roles bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack roles bulk partial update params
func (o *DcimRackRolesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRackRolesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack roles bulk partial update params
func (o *DcimRackRolesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack roles bulk partial update params
func (o *DcimRackRolesBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimRackRolesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack roles bulk partial update params
func (o *DcimRackRolesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack roles bulk partial update params
func (o *DcimRackRolesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRackRolesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack roles bulk partial update params
func (o *DcimRackRolesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack roles bulk partial update params
func (o *DcimRackRolesBulkPartialUpdateParams) WithData(data *models.RackRole) *DcimRackRolesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack roles bulk partial update params
func (o *DcimRackRolesBulkPartialUpdateParams) SetData(data *models.RackRole) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackRolesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
