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

package secrets

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

// NewSecretsSecretRolesDeleteParams creates a new SecretsSecretRolesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretsSecretRolesDeleteParams() *SecretsSecretRolesDeleteParams {
	return &SecretsSecretRolesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretRolesDeleteParamsWithTimeout creates a new SecretsSecretRolesDeleteParams object
// with the ability to set a timeout on a request.
func NewSecretsSecretRolesDeleteParamsWithTimeout(timeout time.Duration) *SecretsSecretRolesDeleteParams {
	return &SecretsSecretRolesDeleteParams{
		timeout: timeout,
	}
}

// NewSecretsSecretRolesDeleteParamsWithContext creates a new SecretsSecretRolesDeleteParams object
// with the ability to set a context for a request.
func NewSecretsSecretRolesDeleteParamsWithContext(ctx context.Context) *SecretsSecretRolesDeleteParams {
	return &SecretsSecretRolesDeleteParams{
		Context: ctx,
	}
}

// NewSecretsSecretRolesDeleteParamsWithHTTPClient creates a new SecretsSecretRolesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretsSecretRolesDeleteParamsWithHTTPClient(client *http.Client) *SecretsSecretRolesDeleteParams {
	return &SecretsSecretRolesDeleteParams{
		HTTPClient: client,
	}
}

/*
SecretsSecretRolesDeleteParams contains all the parameters to send to the API endpoint

	for the secrets secret roles delete operation.

	Typically these are written to a http.Request.
*/
type SecretsSecretRolesDeleteParams struct {

	/* ID.

	   A unique integer value identifying this secret role.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secrets secret roles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretsSecretRolesDeleteParams) WithDefaults() *SecretsSecretRolesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secrets secret roles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretsSecretRolesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secrets secret roles delete params
func (o *SecretsSecretRolesDeleteParams) WithTimeout(timeout time.Duration) *SecretsSecretRolesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secret roles delete params
func (o *SecretsSecretRolesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secret roles delete params
func (o *SecretsSecretRolesDeleteParams) WithContext(ctx context.Context) *SecretsSecretRolesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secret roles delete params
func (o *SecretsSecretRolesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secret roles delete params
func (o *SecretsSecretRolesDeleteParams) WithHTTPClient(client *http.Client) *SecretsSecretRolesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secret roles delete params
func (o *SecretsSecretRolesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the secrets secret roles delete params
func (o *SecretsSecretRolesDeleteParams) WithID(id int64) *SecretsSecretRolesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets secret roles delete params
func (o *SecretsSecretRolesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretRolesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
