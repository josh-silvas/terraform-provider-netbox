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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InterfaceConnection interface connection
//
// swagger:model InterfaceConnection
type InterfaceConnection struct {

	// Connected endpoint reachable
	// Read Only: true
	ConnectedEndpointReachable *bool `json:"connected_endpoint_reachable,omitempty"`

	// interface a
	Interfacea *NestedInterface `json:"interface_a,omitempty"`

	// interface b
	// Required: true
	Interfaceb *NestedInterface `json:"interface_b"`
}

// Validate validates this interface connection
func (m *InterfaceConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterfacea(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaceb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceConnection) validateInterfacea(formats strfmt.Registry) error {
	if swag.IsZero(m.Interfacea) { // not required
		return nil
	}

	if m.Interfacea != nil {
		if err := m.Interfacea.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_a")
			}
			return err
		}
	}

	return nil
}

func (m *InterfaceConnection) validateInterfaceb(formats strfmt.Registry) error {

	if err := validate.Required("interface_b", "body", m.Interfaceb); err != nil {
		return err
	}

	if m.Interfaceb != nil {
		if err := m.Interfaceb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_b")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this interface connection based on the context it is used
func (m *InterfaceConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectedEndpointReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfacea(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaceb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InterfaceConnection) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *InterfaceConnection) contextValidateInterfacea(ctx context.Context, formats strfmt.Registry) error {

	if m.Interfacea != nil {
		if err := m.Interfacea.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_a")
			}
			return err
		}
	}

	return nil
}

func (m *InterfaceConnection) contextValidateInterfaceb(ctx context.Context, formats strfmt.Registry) error {

	if m.Interfaceb != nil {
		if err := m.Interfaceb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_b")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceConnection) UnmarshalBinary(b []byte) error {
	var res InterfaceConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
