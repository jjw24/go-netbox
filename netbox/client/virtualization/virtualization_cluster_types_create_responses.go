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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jjw24/go-netbox/netbox/models"
)

// VirtualizationClusterTypesCreateReader is a Reader for the VirtualizationClusterTypesCreate structure.
type VirtualizationClusterTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVirtualizationClusterTypesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterTypesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterTypesCreateCreated creates a VirtualizationClusterTypesCreateCreated with default headers values
func NewVirtualizationClusterTypesCreateCreated() *VirtualizationClusterTypesCreateCreated {
	return &VirtualizationClusterTypesCreateCreated{}
}

/*VirtualizationClusterTypesCreateCreated handles this case with default header values.

VirtualizationClusterTypesCreateCreated virtualization cluster types create created
*/
type VirtualizationClusterTypesCreateCreated struct {
	Payload *models.ClusterType
}

func (o *VirtualizationClusterTypesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /virtualization/cluster-types/][%d] virtualizationClusterTypesCreateCreated  %+v", 201, o.Payload)
}

func (o *VirtualizationClusterTypesCreateCreated) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterTypesCreateDefault creates a VirtualizationClusterTypesCreateDefault with default headers values
func NewVirtualizationClusterTypesCreateDefault(code int) *VirtualizationClusterTypesCreateDefault {
	return &VirtualizationClusterTypesCreateDefault{
		_statusCode: code,
	}
}

/*VirtualizationClusterTypesCreateDefault handles this case with default header values.

VirtualizationClusterTypesCreateDefault virtualization cluster types create default
*/
type VirtualizationClusterTypesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster types create default response
func (o *VirtualizationClusterTypesCreateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterTypesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /virtualization/cluster-types/][%d] virtualization_cluster-types_create default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterTypesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterTypesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
