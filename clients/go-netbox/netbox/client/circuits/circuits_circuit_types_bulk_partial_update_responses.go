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

// CircuitsCircuitTypesBulkPartialUpdateReader is a Reader for the CircuitsCircuitTypesBulkPartialUpdate structure.
type CircuitsCircuitTypesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTypesBulkPartialUpdateOK creates a CircuitsCircuitTypesBulkPartialUpdateOK with default headers values
func NewCircuitsCircuitTypesBulkPartialUpdateOK() *CircuitsCircuitTypesBulkPartialUpdateOK {
	return &CircuitsCircuitTypesBulkPartialUpdateOK{}
}

/*
	CircuitsCircuitTypesBulkPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTypesBulkPartialUpdateOK circuits circuit types bulk partial update o k
*/
type CircuitsCircuitTypesBulkPartialUpdateOK struct {
	Payload *models.CircuitType
}

func (o *CircuitsCircuitTypesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/circuit-types/][%d] circuitsCircuitTypesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitTypesBulkPartialUpdateOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
