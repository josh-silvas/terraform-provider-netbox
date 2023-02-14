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

// NewSecretsSecretsListParams creates a new SecretsSecretsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretsSecretsListParams() *SecretsSecretsListParams {
	return &SecretsSecretsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretsListParamsWithTimeout creates a new SecretsSecretsListParams object
// with the ability to set a timeout on a request.
func NewSecretsSecretsListParamsWithTimeout(timeout time.Duration) *SecretsSecretsListParams {
	return &SecretsSecretsListParams{
		timeout: timeout,
	}
}

// NewSecretsSecretsListParamsWithContext creates a new SecretsSecretsListParams object
// with the ability to set a context for a request.
func NewSecretsSecretsListParamsWithContext(ctx context.Context) *SecretsSecretsListParams {
	return &SecretsSecretsListParams{
		Context: ctx,
	}
}

// NewSecretsSecretsListParamsWithHTTPClient creates a new SecretsSecretsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretsSecretsListParamsWithHTTPClient(client *http.Client) *SecretsSecretsListParams {
	return &SecretsSecretsListParams{
		HTTPClient: client,
	}
}

/*
SecretsSecretsListParams contains all the parameters to send to the API endpoint

	for the secrets secrets list operation.

	Typically these are written to a http.Request.
*/
type SecretsSecretsListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Device.
	Device *string

	// Devicen.
	Devicen *string

	// DeviceID.
	DeviceID *string

	// DeviceIDn.
	DeviceIDn *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameEmpty.
	NameEmpty *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// Role.
	Role *string

	// Rolen.
	Rolen *string

	// RoleID.
	RoleID *string

	// RoleIDn.
	RoleIDn *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// VirtualMachine.
	VirtualMachine *string

	// VirtualMachinen.
	VirtualMachinen *string

	// VirtualMachineID.
	VirtualMachineID *string

	// VirtualMachineIDn.
	VirtualMachineIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secrets secrets list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretsSecretsListParams) WithDefaults() *SecretsSecretsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secrets secrets list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretsSecretsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secrets secrets list params
func (o *SecretsSecretsListParams) WithTimeout(timeout time.Duration) *SecretsSecretsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secrets list params
func (o *SecretsSecretsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secrets list params
func (o *SecretsSecretsListParams) WithContext(ctx context.Context) *SecretsSecretsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secrets list params
func (o *SecretsSecretsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secrets list params
func (o *SecretsSecretsListParams) WithHTTPClient(client *http.Client) *SecretsSecretsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secrets list params
func (o *SecretsSecretsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the secrets secrets list params
func (o *SecretsSecretsListParams) WithCreated(created *string) *SecretsSecretsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the secrets secrets list params
func (o *SecretsSecretsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the secrets secrets list params
func (o *SecretsSecretsListParams) WithCreatedGte(createdGte *string) *SecretsSecretsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the secrets secrets list params
func (o *SecretsSecretsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the secrets secrets list params
func (o *SecretsSecretsListParams) WithCreatedLte(createdLte *string) *SecretsSecretsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the secrets secrets list params
func (o *SecretsSecretsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDevice adds the device to the secrets secrets list params
func (o *SecretsSecretsListParams) WithDevice(device *string) *SecretsSecretsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the secrets secrets list params
func (o *SecretsSecretsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDevicen adds the devicen to the secrets secrets list params
func (o *SecretsSecretsListParams) WithDevicen(devicen *string) *SecretsSecretsListParams {
	o.SetDevicen(devicen)
	return o
}

// SetDevicen adds the deviceN to the secrets secrets list params
func (o *SecretsSecretsListParams) SetDevicen(devicen *string) {
	o.Devicen = devicen
}

// WithDeviceID adds the deviceID to the secrets secrets list params
func (o *SecretsSecretsListParams) WithDeviceID(deviceID *string) *SecretsSecretsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the secrets secrets list params
func (o *SecretsSecretsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceIDn adds the deviceIDn to the secrets secrets list params
func (o *SecretsSecretsListParams) WithDeviceIDn(deviceIDn *string) *SecretsSecretsListParams {
	o.SetDeviceIDn(deviceIDn)
	return o
}

// SetDeviceIDn adds the deviceIdN to the secrets secrets list params
func (o *SecretsSecretsListParams) SetDeviceIDn(deviceIDn *string) {
	o.DeviceIDn = deviceIDn
}

// WithID adds the id to the secrets secrets list params
func (o *SecretsSecretsListParams) WithID(id *string) *SecretsSecretsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets secrets list params
func (o *SecretsSecretsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the secrets secrets list params
func (o *SecretsSecretsListParams) WithIDGt(iDGt *string) *SecretsSecretsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the secrets secrets list params
func (o *SecretsSecretsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the secrets secrets list params
func (o *SecretsSecretsListParams) WithIDGte(iDGte *string) *SecretsSecretsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the secrets secrets list params
func (o *SecretsSecretsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the secrets secrets list params
func (o *SecretsSecretsListParams) WithIDLt(iDLt *string) *SecretsSecretsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the secrets secrets list params
func (o *SecretsSecretsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the secrets secrets list params
func (o *SecretsSecretsListParams) WithIDLte(iDLte *string) *SecretsSecretsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the secrets secrets list params
func (o *SecretsSecretsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the secrets secrets list params
func (o *SecretsSecretsListParams) WithIDn(iDn *string) *SecretsSecretsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the secrets secrets list params
func (o *SecretsSecretsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the secrets secrets list params
func (o *SecretsSecretsListParams) WithLastUpdated(lastUpdated *string) *SecretsSecretsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the secrets secrets list params
func (o *SecretsSecretsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the secrets secrets list params
func (o *SecretsSecretsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *SecretsSecretsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the secrets secrets list params
func (o *SecretsSecretsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the secrets secrets list params
func (o *SecretsSecretsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *SecretsSecretsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the secrets secrets list params
func (o *SecretsSecretsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the secrets secrets list params
func (o *SecretsSecretsListParams) WithLimit(limit *int64) *SecretsSecretsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the secrets secrets list params
func (o *SecretsSecretsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the secrets secrets list params
func (o *SecretsSecretsListParams) WithName(name *string) *SecretsSecretsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the secrets secrets list params
func (o *SecretsSecretsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the secrets secrets list params
func (o *SecretsSecretsListParams) WithNameEmpty(nameEmpty *string) *SecretsSecretsListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the secrets secrets list params
func (o *SecretsSecretsListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the secrets secrets list params
func (o *SecretsSecretsListParams) WithNameIc(nameIc *string) *SecretsSecretsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the secrets secrets list params
func (o *SecretsSecretsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the secrets secrets list params
func (o *SecretsSecretsListParams) WithNameIe(nameIe *string) *SecretsSecretsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the secrets secrets list params
func (o *SecretsSecretsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the secrets secrets list params
func (o *SecretsSecretsListParams) WithNameIew(nameIew *string) *SecretsSecretsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the secrets secrets list params
func (o *SecretsSecretsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the secrets secrets list params
func (o *SecretsSecretsListParams) WithNameIsw(nameIsw *string) *SecretsSecretsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the secrets secrets list params
func (o *SecretsSecretsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the secrets secrets list params
func (o *SecretsSecretsListParams) WithNamen(namen *string) *SecretsSecretsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the secrets secrets list params
func (o *SecretsSecretsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the secrets secrets list params
func (o *SecretsSecretsListParams) WithNameNic(nameNic *string) *SecretsSecretsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the secrets secrets list params
func (o *SecretsSecretsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the secrets secrets list params
func (o *SecretsSecretsListParams) WithNameNie(nameNie *string) *SecretsSecretsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the secrets secrets list params
func (o *SecretsSecretsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the secrets secrets list params
func (o *SecretsSecretsListParams) WithNameNiew(nameNiew *string) *SecretsSecretsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the secrets secrets list params
func (o *SecretsSecretsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the secrets secrets list params
func (o *SecretsSecretsListParams) WithNameNisw(nameNisw *string) *SecretsSecretsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the secrets secrets list params
func (o *SecretsSecretsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the secrets secrets list params
func (o *SecretsSecretsListParams) WithOffset(offset *int64) *SecretsSecretsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the secrets secrets list params
func (o *SecretsSecretsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the secrets secrets list params
func (o *SecretsSecretsListParams) WithQ(q *string) *SecretsSecretsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the secrets secrets list params
func (o *SecretsSecretsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the secrets secrets list params
func (o *SecretsSecretsListParams) WithRole(role *string) *SecretsSecretsListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the secrets secrets list params
func (o *SecretsSecretsListParams) SetRole(role *string) {
	o.Role = role
}

// WithRolen adds the rolen to the secrets secrets list params
func (o *SecretsSecretsListParams) WithRolen(rolen *string) *SecretsSecretsListParams {
	o.SetRolen(rolen)
	return o
}

// SetRolen adds the roleN to the secrets secrets list params
func (o *SecretsSecretsListParams) SetRolen(rolen *string) {
	o.Rolen = rolen
}

// WithRoleID adds the roleID to the secrets secrets list params
func (o *SecretsSecretsListParams) WithRoleID(roleID *string) *SecretsSecretsListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the secrets secrets list params
func (o *SecretsSecretsListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithRoleIDn adds the roleIDn to the secrets secrets list params
func (o *SecretsSecretsListParams) WithRoleIDn(roleIDn *string) *SecretsSecretsListParams {
	o.SetRoleIDn(roleIDn)
	return o
}

// SetRoleIDn adds the roleIdN to the secrets secrets list params
func (o *SecretsSecretsListParams) SetRoleIDn(roleIDn *string) {
	o.RoleIDn = roleIDn
}

// WithTag adds the tag to the secrets secrets list params
func (o *SecretsSecretsListParams) WithTag(tag *string) *SecretsSecretsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the secrets secrets list params
func (o *SecretsSecretsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the secrets secrets list params
func (o *SecretsSecretsListParams) WithTagn(tagn *string) *SecretsSecretsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the secrets secrets list params
func (o *SecretsSecretsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithVirtualMachine adds the virtualMachine to the secrets secrets list params
func (o *SecretsSecretsListParams) WithVirtualMachine(virtualMachine *string) *SecretsSecretsListParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the secrets secrets list params
func (o *SecretsSecretsListParams) SetVirtualMachine(virtualMachine *string) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachinen adds the virtualMachinen to the secrets secrets list params
func (o *SecretsSecretsListParams) WithVirtualMachinen(virtualMachinen *string) *SecretsSecretsListParams {
	o.SetVirtualMachinen(virtualMachinen)
	return o
}

// SetVirtualMachinen adds the virtualMachineN to the secrets secrets list params
func (o *SecretsSecretsListParams) SetVirtualMachinen(virtualMachinen *string) {
	o.VirtualMachinen = virtualMachinen
}

// WithVirtualMachineID adds the virtualMachineID to the secrets secrets list params
func (o *SecretsSecretsListParams) WithVirtualMachineID(virtualMachineID *string) *SecretsSecretsListParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the secrets secrets list params
func (o *SecretsSecretsListParams) SetVirtualMachineID(virtualMachineID *string) {
	o.VirtualMachineID = virtualMachineID
}

// WithVirtualMachineIDn adds the virtualMachineIDn to the secrets secrets list params
func (o *SecretsSecretsListParams) WithVirtualMachineIDn(virtualMachineIDn *string) *SecretsSecretsListParams {
	o.SetVirtualMachineIDn(virtualMachineIDn)
	return o
}

// SetVirtualMachineIDn adds the virtualMachineIdN to the secrets secrets list params
func (o *SecretsSecretsListParams) SetVirtualMachineIDn(virtualMachineIDn *string) {
	o.VirtualMachineIDn = virtualMachineIDn
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.Device != nil {

		// query param device
		var qrDevice string

		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {

			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}
	}

	if o.Devicen != nil {

		// query param device__n
		var qrDevicen string

		if o.Devicen != nil {
			qrDevicen = *o.Devicen
		}
		qDevicen := qrDevicen
		if qDevicen != "" {

			if err := r.SetQueryParam("device__n", qDevicen); err != nil {
				return err
			}
		}
	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string

		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {

			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}
	}

	if o.DeviceIDn != nil {

		// query param device_id__n
		var qrDeviceIDn string

		if o.DeviceIDn != nil {
			qrDeviceIDn = *o.DeviceIDn
		}
		qDeviceIDn := qrDeviceIDn
		if qDeviceIDn != "" {

			if err := r.SetQueryParam("device_id__n", qDeviceIDn); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
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

	if o.NameEmpty != nil {

		// query param name__empty
		var qrNameEmpty string

		if o.NameEmpty != nil {
			qrNameEmpty = *o.NameEmpty
		}
		qNameEmpty := qrNameEmpty
		if qNameEmpty != "" {

			if err := r.SetQueryParam("name__empty", qNameEmpty); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
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

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Role != nil {

		// query param role
		var qrRole string

		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {

			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}
	}

	if o.Rolen != nil {

		// query param role__n
		var qrRolen string

		if o.Rolen != nil {
			qrRolen = *o.Rolen
		}
		qRolen := qrRolen
		if qRolen != "" {

			if err := r.SetQueryParam("role__n", qRolen); err != nil {
				return err
			}
		}
	}

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID string

		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {

			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
				return err
			}
		}
	}

	if o.RoleIDn != nil {

		// query param role_id__n
		var qrRoleIDn string

		if o.RoleIDn != nil {
			qrRoleIDn = *o.RoleIDn
		}
		qRoleIDn := qrRoleIDn
		if qRoleIDn != "" {

			if err := r.SetQueryParam("role_id__n", qRoleIDn); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachine != nil {

		// query param virtual_machine
		var qrVirtualMachine string

		if o.VirtualMachine != nil {
			qrVirtualMachine = *o.VirtualMachine
		}
		qVirtualMachine := qrVirtualMachine
		if qVirtualMachine != "" {

			if err := r.SetQueryParam("virtual_machine", qVirtualMachine); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachinen != nil {

		// query param virtual_machine__n
		var qrVirtualMachinen string

		if o.VirtualMachinen != nil {
			qrVirtualMachinen = *o.VirtualMachinen
		}
		qVirtualMachinen := qrVirtualMachinen
		if qVirtualMachinen != "" {

			if err := r.SetQueryParam("virtual_machine__n", qVirtualMachinen); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachineID != nil {

		// query param virtual_machine_id
		var qrVirtualMachineID string

		if o.VirtualMachineID != nil {
			qrVirtualMachineID = *o.VirtualMachineID
		}
		qVirtualMachineID := qrVirtualMachineID
		if qVirtualMachineID != "" {

			if err := r.SetQueryParam("virtual_machine_id", qVirtualMachineID); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachineIDn != nil {

		// query param virtual_machine_id__n
		var qrVirtualMachineIDn string

		if o.VirtualMachineIDn != nil {
			qrVirtualMachineIDn = *o.VirtualMachineIDn
		}
		qVirtualMachineIDn := qrVirtualMachineIDn
		if qVirtualMachineIDn != "" {

			if err := r.SetQueryParam("virtual_machine_id__n", qVirtualMachineIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
