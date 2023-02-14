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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/josh-silvas/terraform-provider-netbox/clients/go-netbox/netbox/models"
)

// SecretsSecretRolesCreateReader is a Reader for the SecretsSecretRolesCreate structure.
type SecretsSecretRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecretsSecretRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSecretsSecretRolesCreateCreated creates a SecretsSecretRolesCreateCreated with default headers values
func NewSecretsSecretRolesCreateCreated() *SecretsSecretRolesCreateCreated {
	return &SecretsSecretRolesCreateCreated{}
}

/*
	SecretsSecretRolesCreateCreated describes a response with status code 201, with default header values.

SecretsSecretRolesCreateCreated secrets secret roles create created
*/
type SecretsSecretRolesCreateCreated struct {
	Payload *models.SecretRole
}

func (o *SecretsSecretRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /secrets/secret-roles/][%d] secretsSecretRolesCreateCreated  %+v", 201, o.Payload)
}
func (o *SecretsSecretRolesCreateCreated) GetPayload() *models.SecretRole {
	return o.Payload
}

func (o *SecretsSecretRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
