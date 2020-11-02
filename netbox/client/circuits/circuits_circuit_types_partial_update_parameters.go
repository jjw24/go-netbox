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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/jjw24/go-netbox/netbox/models"
)

// NewCircuitsCircuitTypesPartialUpdateParams creates a new CircuitsCircuitTypesPartialUpdateParams object
// with the default values initialized.
func NewCircuitsCircuitTypesPartialUpdateParams() *CircuitsCircuitTypesPartialUpdateParams {
	var ()
	return &CircuitsCircuitTypesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesPartialUpdateParamsWithTimeout creates a new CircuitsCircuitTypesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitTypesPartialUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesPartialUpdateParams {
	var ()
	return &CircuitsCircuitTypesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesPartialUpdateParamsWithContext creates a new CircuitsCircuitTypesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitTypesPartialUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitTypesPartialUpdateParams {
	var ()
	return &CircuitsCircuitTypesPartialUpdateParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitTypesPartialUpdateParamsWithHTTPClient creates a new CircuitsCircuitTypesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitTypesPartialUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesPartialUpdateParams {
	var ()
	return &CircuitsCircuitTypesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitTypesPartialUpdateParams contains all the parameters to send to the API endpoint
for the circuits circuit types partial update operation typically these are written to a http.Request
*/
type CircuitsCircuitTypesPartialUpdateParams struct {

	/*Data*/
	Data *models.CircuitType
	/*ID
	  A unique integer value identifying this circuit type.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitTypesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) WithData(data *models.CircuitType) *CircuitsCircuitTypesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) SetData(data *models.CircuitType) {
	o.Data = data
}

// WithID adds the id to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) WithID(id int64) *CircuitsCircuitTypesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
