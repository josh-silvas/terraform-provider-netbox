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

package users

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

// NewUsersPermissionsUpdateParams creates a new UsersPermissionsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersPermissionsUpdateParams() *UsersPermissionsUpdateParams {
	return &UsersPermissionsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersPermissionsUpdateParamsWithTimeout creates a new UsersPermissionsUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersPermissionsUpdateParamsWithTimeout(timeout time.Duration) *UsersPermissionsUpdateParams {
	return &UsersPermissionsUpdateParams{
		timeout: timeout,
	}
}

// NewUsersPermissionsUpdateParamsWithContext creates a new UsersPermissionsUpdateParams object
// with the ability to set a context for a request.
func NewUsersPermissionsUpdateParamsWithContext(ctx context.Context) *UsersPermissionsUpdateParams {
	return &UsersPermissionsUpdateParams{
		Context: ctx,
	}
}

// NewUsersPermissionsUpdateParamsWithHTTPClient creates a new UsersPermissionsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersPermissionsUpdateParamsWithHTTPClient(client *http.Client) *UsersPermissionsUpdateParams {
	return &UsersPermissionsUpdateParams{
		HTTPClient: client,
	}
}

/*
UsersPermissionsUpdateParams contains all the parameters to send to the API endpoint

	for the users permissions update operation.

	Typically these are written to a http.Request.
*/
type UsersPermissionsUpdateParams struct {

	// Data.
	Data *models.WritableObjectPermission

	/* ID.

	   A unique integer value identifying this permission.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users permissions update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsUpdateParams) WithDefaults() *UsersPermissionsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users permissions update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users permissions update params
func (o *UsersPermissionsUpdateParams) WithTimeout(timeout time.Duration) *UsersPermissionsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users permissions update params
func (o *UsersPermissionsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users permissions update params
func (o *UsersPermissionsUpdateParams) WithContext(ctx context.Context) *UsersPermissionsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users permissions update params
func (o *UsersPermissionsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users permissions update params
func (o *UsersPermissionsUpdateParams) WithHTTPClient(client *http.Client) *UsersPermissionsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users permissions update params
func (o *UsersPermissionsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users permissions update params
func (o *UsersPermissionsUpdateParams) WithData(data *models.WritableObjectPermission) *UsersPermissionsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users permissions update params
func (o *UsersPermissionsUpdateParams) SetData(data *models.WritableObjectPermission) {
	o.Data = data
}

// WithID adds the id to the users permissions update params
func (o *UsersPermissionsUpdateParams) WithID(id int64) *UsersPermissionsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users permissions update params
func (o *UsersPermissionsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersPermissionsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
