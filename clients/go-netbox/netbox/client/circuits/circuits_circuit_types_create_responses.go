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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

// CircuitsCircuitTypesCreateReader is a Reader for the CircuitsCircuitTypesCreate structure.
type CircuitsCircuitTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCircuitsCircuitTypesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTypesCreateCreated creates a CircuitsCircuitTypesCreateCreated with default headers values
func NewCircuitsCircuitTypesCreateCreated() *CircuitsCircuitTypesCreateCreated {
	return &CircuitsCircuitTypesCreateCreated{}
}

/*
	CircuitsCircuitTypesCreateCreated describes a response with status code 201, with default header values.

CircuitsCircuitTypesCreateCreated circuits circuit types create created
*/
type CircuitsCircuitTypesCreateCreated struct {
	Payload *models.CircuitType
}

func (o *CircuitsCircuitTypesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/circuit-types/][%d] circuitsCircuitTypesCreateCreated  %+v", 201, o.Payload)
}
func (o *CircuitsCircuitTypesCreateCreated) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
