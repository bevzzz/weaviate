//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// SchemaActionsPropertiesAddReader is a Reader for the SchemaActionsPropertiesAdd structure.
type SchemaActionsPropertiesAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaActionsPropertiesAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSchemaActionsPropertiesAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSchemaActionsPropertiesAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSchemaActionsPropertiesAddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewSchemaActionsPropertiesAddUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSchemaActionsPropertiesAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaActionsPropertiesAddOK creates a SchemaActionsPropertiesAddOK with default headers values
func NewSchemaActionsPropertiesAddOK() *SchemaActionsPropertiesAddOK {
	return &SchemaActionsPropertiesAddOK{}
}

/*SchemaActionsPropertiesAddOK handles this case with default header values.

Added the property.
*/
type SchemaActionsPropertiesAddOK struct {
	Payload *models.Property
}

func (o *SchemaActionsPropertiesAddOK) Error() string {
	return fmt.Sprintf("[POST /schema/actions/{className}/properties][%d] schemaActionsPropertiesAddOK  %+v", 200, o.Payload)
}

func (o *SchemaActionsPropertiesAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Property)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaActionsPropertiesAddUnauthorized creates a SchemaActionsPropertiesAddUnauthorized with default headers values
func NewSchemaActionsPropertiesAddUnauthorized() *SchemaActionsPropertiesAddUnauthorized {
	return &SchemaActionsPropertiesAddUnauthorized{}
}

/*SchemaActionsPropertiesAddUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type SchemaActionsPropertiesAddUnauthorized struct {
}

func (o *SchemaActionsPropertiesAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schema/actions/{className}/properties][%d] schemaActionsPropertiesAddUnauthorized ", 401)
}

func (o *SchemaActionsPropertiesAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaActionsPropertiesAddForbidden creates a SchemaActionsPropertiesAddForbidden with default headers values
func NewSchemaActionsPropertiesAddForbidden() *SchemaActionsPropertiesAddForbidden {
	return &SchemaActionsPropertiesAddForbidden{}
}

/*SchemaActionsPropertiesAddForbidden handles this case with default header values.

Forbidden
*/
type SchemaActionsPropertiesAddForbidden struct {
	Payload *models.ErrorResponse
}

func (o *SchemaActionsPropertiesAddForbidden) Error() string {
	return fmt.Sprintf("[POST /schema/actions/{className}/properties][%d] schemaActionsPropertiesAddForbidden  %+v", 403, o.Payload)
}

func (o *SchemaActionsPropertiesAddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaActionsPropertiesAddUnprocessableEntity creates a SchemaActionsPropertiesAddUnprocessableEntity with default headers values
func NewSchemaActionsPropertiesAddUnprocessableEntity() *SchemaActionsPropertiesAddUnprocessableEntity {
	return &SchemaActionsPropertiesAddUnprocessableEntity{}
}

/*SchemaActionsPropertiesAddUnprocessableEntity handles this case with default header values.

Invalid property.
*/
type SchemaActionsPropertiesAddUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *SchemaActionsPropertiesAddUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /schema/actions/{className}/properties][%d] schemaActionsPropertiesAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SchemaActionsPropertiesAddUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaActionsPropertiesAddInternalServerError creates a SchemaActionsPropertiesAddInternalServerError with default headers values
func NewSchemaActionsPropertiesAddInternalServerError() *SchemaActionsPropertiesAddInternalServerError {
	return &SchemaActionsPropertiesAddInternalServerError{}
}

/*SchemaActionsPropertiesAddInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaActionsPropertiesAddInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SchemaActionsPropertiesAddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /schema/actions/{className}/properties][%d] schemaActionsPropertiesAddInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaActionsPropertiesAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}