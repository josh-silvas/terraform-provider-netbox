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
	"github.com/go-openapi/swag"
)

// NewDcimFrontPortTemplatesListParams creates a new DcimFrontPortTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortTemplatesListParams() *DcimFrontPortTemplatesListParams {
	return &DcimFrontPortTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesListParamsWithTimeout creates a new DcimFrontPortTemplatesListParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesListParams {
	return &DcimFrontPortTemplatesListParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesListParamsWithContext creates a new DcimFrontPortTemplatesListParams object
// with the ability to set a context for a request.
func NewDcimFrontPortTemplatesListParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesListParams {
	return &DcimFrontPortTemplatesListParams{
		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesListParamsWithHTTPClient creates a new DcimFrontPortTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesListParams {
	return &DcimFrontPortTemplatesListParams{
		HTTPClient: client,
	}
}

/*
DcimFrontPortTemplatesListParams contains all the parameters to send to the API endpoint

	for the dcim front port templates list operation.

	Typically these are written to a http.Request.
*/
type DcimFrontPortTemplatesListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// DevicetypeID.
	DevicetypeID *string

	// DevicetypeIDn.
	DevicetypeIDn *string

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

	// Type.
	Type *string

	// Typen.
	Typen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesListParams) WithDefaults() *DcimFrontPortTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithCreated(created *string) *DcimFrontPortTemplatesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithCreatedGte(createdGte *string) *DcimFrontPortTemplatesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithCreatedLte(createdLte *string) *DcimFrontPortTemplatesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDevicetypeID adds the devicetypeID to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimFrontPortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithDevicetypeIDn adds the devicetypeIDn to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithDevicetypeIDn(devicetypeIDn *string) *DcimFrontPortTemplatesListParams {
	o.SetDevicetypeIDn(devicetypeIDn)
	return o
}

// SetDevicetypeIDn adds the devicetypeIdN to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetDevicetypeIDn(devicetypeIDn *string) {
	o.DevicetypeIDn = devicetypeIDn
}

// WithID adds the id to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithID(id *string) *DcimFrontPortTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithIDGt(iDGt *string) *DcimFrontPortTemplatesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithIDGte(iDGte *string) *DcimFrontPortTemplatesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithIDLt(iDLt *string) *DcimFrontPortTemplatesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithIDLte(iDLte *string) *DcimFrontPortTemplatesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithIDn(iDn *string) *DcimFrontPortTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithLastUpdated(lastUpdated *string) *DcimFrontPortTemplatesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimFrontPortTemplatesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimFrontPortTemplatesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithLimit(limit *int64) *DcimFrontPortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithName(name *string) *DcimFrontPortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameEmpty(nameEmpty *string) *DcimFrontPortTemplatesListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameIc(nameIc *string) *DcimFrontPortTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameIe(nameIe *string) *DcimFrontPortTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameIew(nameIew *string) *DcimFrontPortTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameIsw(nameIsw *string) *DcimFrontPortTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNamen(namen *string) *DcimFrontPortTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameNic(nameNic *string) *DcimFrontPortTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameNie(nameNie *string) *DcimFrontPortTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameNiew(nameNiew *string) *DcimFrontPortTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameNisw(nameNisw *string) *DcimFrontPortTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithOffset(offset *int64) *DcimFrontPortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithType adds the typeVar to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithType(typeVar *string) *DcimFrontPortTemplatesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithTypen(typen *string) *DcimFrontPortTemplatesListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID string

		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := qrDevicetypeID
		if qDevicetypeID != "" {

			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
				return err
			}
		}
	}

	if o.DevicetypeIDn != nil {

		// query param devicetype_id__n
		var qrDevicetypeIDn string

		if o.DevicetypeIDn != nil {
			qrDevicetypeIDn = *o.DevicetypeIDn
		}
		qDevicetypeIDn := qrDevicetypeIDn
		if qDevicetypeIDn != "" {

			if err := r.SetQueryParam("devicetype_id__n", qDevicetypeIDn); err != nil {
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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Typen != nil {

		// query param type__n
		var qrTypen string

		if o.Typen != nil {
			qrTypen = *o.Typen
		}
		qTypen := qrTypen
		if qTypen != "" {

			if err := r.SetQueryParam("type__n", qTypen); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
