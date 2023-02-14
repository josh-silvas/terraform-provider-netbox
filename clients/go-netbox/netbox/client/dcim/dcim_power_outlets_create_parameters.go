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

// NewDcimPowerOutletsCreateParams creates a new DcimPowerOutletsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletsCreateParams() *DcimPowerOutletsCreateParams {
	return &DcimPowerOutletsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsCreateParamsWithTimeout creates a new DcimPowerOutletsCreateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletsCreateParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsCreateParams {
	return &DcimPowerOutletsCreateParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletsCreateParamsWithContext creates a new DcimPowerOutletsCreateParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletsCreateParamsWithContext(ctx context.Context) *DcimPowerOutletsCreateParams {
	return &DcimPowerOutletsCreateParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletsCreateParamsWithHTTPClient creates a new DcimPowerOutletsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletsCreateParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsCreateParams {
	return &DcimPowerOutletsCreateParams{
		HTTPClient: client,
	}
}

/*
DcimPowerOutletsCreateParams contains all the parameters to send to the API endpoint

	for the dcim power outlets create operation.

	Typically these are written to a http.Request.
*/
type DcimPowerOutletsCreateParams struct {

	// Data.
	Data *models.WritablePowerOutlet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlets create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsCreateParams) WithDefaults() *DcimPowerOutletsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlets create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlets create params
func (o *DcimPowerOutletsCreateParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets create params
func (o *DcimPowerOutletsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets create params
func (o *DcimPowerOutletsCreateParams) WithContext(ctx context.Context) *DcimPowerOutletsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets create params
func (o *DcimPowerOutletsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets create params
func (o *DcimPowerOutletsCreateParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets create params
func (o *DcimPowerOutletsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power outlets create params
func (o *DcimPowerOutletsCreateParams) WithData(data *models.WritablePowerOutlet) *DcimPowerOutletsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power outlets create params
func (o *DcimPowerOutletsCreateParams) SetData(data *models.WritablePowerOutlet) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
