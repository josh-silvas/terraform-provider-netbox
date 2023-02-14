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

// NewDcimPowerOutletTemplatesBulkUpdateParams creates a new DcimPowerOutletTemplatesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletTemplatesBulkUpdateParams() *DcimPowerOutletTemplatesBulkUpdateParams {
	return &DcimPowerOutletTemplatesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletTemplatesBulkUpdateParamsWithTimeout creates a new DcimPowerOutletTemplatesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletTemplatesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesBulkUpdateParams {
	return &DcimPowerOutletTemplatesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletTemplatesBulkUpdateParamsWithContext creates a new DcimPowerOutletTemplatesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletTemplatesBulkUpdateParamsWithContext(ctx context.Context) *DcimPowerOutletTemplatesBulkUpdateParams {
	return &DcimPowerOutletTemplatesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletTemplatesBulkUpdateParamsWithHTTPClient creates a new DcimPowerOutletTemplatesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletTemplatesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesBulkUpdateParams {
	return &DcimPowerOutletTemplatesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimPowerOutletTemplatesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim power outlet templates bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimPowerOutletTemplatesBulkUpdateParams struct {

	// Data.
	Data *models.WritablePowerOutletTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlet templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletTemplatesBulkUpdateParams) WithDefaults() *DcimPowerOutletTemplatesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlet templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletTemplatesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlet templates bulk update params
func (o *DcimPowerOutletTemplatesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlet templates bulk update params
func (o *DcimPowerOutletTemplatesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlet templates bulk update params
func (o *DcimPowerOutletTemplatesBulkUpdateParams) WithContext(ctx context.Context) *DcimPowerOutletTemplatesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlet templates bulk update params
func (o *DcimPowerOutletTemplatesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlet templates bulk update params
func (o *DcimPowerOutletTemplatesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlet templates bulk update params
func (o *DcimPowerOutletTemplatesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power outlet templates bulk update params
func (o *DcimPowerOutletTemplatesBulkUpdateParams) WithData(data *models.WritablePowerOutletTemplate) *DcimPowerOutletTemplatesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power outlet templates bulk update params
func (o *DcimPowerOutletTemplatesBulkUpdateParams) SetData(data *models.WritablePowerOutletTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletTemplatesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
