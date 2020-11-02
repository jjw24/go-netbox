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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jjw24/go-netbox/netbox/models"
)

// DcimConsoleServerPortTemplatesPartialUpdateReader is a Reader for the DcimConsoleServerPortTemplatesPartialUpdate structure.
type DcimConsoleServerPortTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortTemplatesPartialUpdateOK creates a DcimConsoleServerPortTemplatesPartialUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesPartialUpdateOK() *DcimConsoleServerPortTemplatesPartialUpdateOK {
	return &DcimConsoleServerPortTemplatesPartialUpdateOK{}
}

/*DcimConsoleServerPortTemplatesPartialUpdateOK handles this case with default header values.

DcimConsoleServerPortTemplatesPartialUpdateOK dcim console server port templates partial update o k
*/
type DcimConsoleServerPortTemplatesPartialUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortTemplatesPartialUpdateDefault creates a DcimConsoleServerPortTemplatesPartialUpdateDefault with default headers values
func NewDcimConsoleServerPortTemplatesPartialUpdateDefault(code int) *DcimConsoleServerPortTemplatesPartialUpdateDefault {
	return &DcimConsoleServerPortTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimConsoleServerPortTemplatesPartialUpdateDefault handles this case with default header values.

DcimConsoleServerPortTemplatesPartialUpdateDefault dcim console server port templates partial update default
*/
type DcimConsoleServerPortTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server port templates partial update default response
func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
