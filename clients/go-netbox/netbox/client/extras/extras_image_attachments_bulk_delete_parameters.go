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
)

// NewExtrasImageAttachmentsBulkDeleteParams creates a new ExtrasImageAttachmentsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasImageAttachmentsBulkDeleteParams() *ExtrasImageAttachmentsBulkDeleteParams {
	return &ExtrasImageAttachmentsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsBulkDeleteParamsWithTimeout creates a new ExtrasImageAttachmentsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasImageAttachmentsBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsBulkDeleteParams {
	return &ExtrasImageAttachmentsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsBulkDeleteParamsWithContext creates a new ExtrasImageAttachmentsBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasImageAttachmentsBulkDeleteParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsBulkDeleteParams {
	return &ExtrasImageAttachmentsBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasImageAttachmentsBulkDeleteParamsWithHTTPClient creates a new ExtrasImageAttachmentsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasImageAttachmentsBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsBulkDeleteParams {
	return &ExtrasImageAttachmentsBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
ExtrasImageAttachmentsBulkDeleteParams contains all the parameters to send to the API endpoint

	for the extras image attachments bulk delete operation.

	Typically these are written to a http.Request.
*/
type ExtrasImageAttachmentsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras image attachments bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsBulkDeleteParams) WithDefaults() *ExtrasImageAttachmentsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras image attachments bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras image attachments bulk delete params
func (o *ExtrasImageAttachmentsBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments bulk delete params
func (o *ExtrasImageAttachmentsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments bulk delete params
func (o *ExtrasImageAttachmentsBulkDeleteParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments bulk delete params
func (o *ExtrasImageAttachmentsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments bulk delete params
func (o *ExtrasImageAttachmentsBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments bulk delete params
func (o *ExtrasImageAttachmentsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}