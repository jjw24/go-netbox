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

	"github.com/jjw24/go-netbox/netbox/models"
)

// CircuitsProvidersUpdateReader is a Reader for the CircuitsProvidersUpdate structure.
type CircuitsProvidersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsProvidersUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProvidersUpdateOK creates a CircuitsProvidersUpdateOK with default headers values
func NewCircuitsProvidersUpdateOK() *CircuitsProvidersUpdateOK {
	return &CircuitsProvidersUpdateOK{}
}

/*CircuitsProvidersUpdateOK handles this case with default header values.

CircuitsProvidersUpdateOK circuits providers update o k
*/
type CircuitsProvidersUpdateOK struct {
	Payload *models.Provider
}

func (o *CircuitsProvidersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/providers/{id}/][%d] circuitsProvidersUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersUpdateOK) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProvidersUpdateDefault creates a CircuitsProvidersUpdateDefault with default headers values
func NewCircuitsProvidersUpdateDefault(code int) *CircuitsProvidersUpdateDefault {
	return &CircuitsProvidersUpdateDefault{
		_statusCode: code,
	}
}

/*CircuitsProvidersUpdateDefault handles this case with default header values.

CircuitsProvidersUpdateDefault circuits providers update default
*/
type CircuitsProvidersUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits providers update default response
func (o *CircuitsProvidersUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsProvidersUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /circuits/providers/{id}/][%d] circuits_providers_update default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
