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

	"github.com/jjw24/go-netbox/netbox/models"
)

// NewDcimFrontPortTemplatesPartialUpdateParams creates a new DcimFrontPortTemplatesPartialUpdateParams object
// with the default values initialized.
func NewDcimFrontPortTemplatesPartialUpdateParams() *DcimFrontPortTemplatesPartialUpdateParams {
	var ()
	return &DcimFrontPortTemplatesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesPartialUpdateParamsWithTimeout creates a new DcimFrontPortTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimFrontPortTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesPartialUpdateParams {
	var ()
	return &DcimFrontPortTemplatesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesPartialUpdateParamsWithContext creates a new DcimFrontPortTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimFrontPortTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesPartialUpdateParams {
	var ()
	return &DcimFrontPortTemplatesPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimFrontPortTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimFrontPortTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesPartialUpdateParams {
	var ()
	return &DcimFrontPortTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimFrontPortTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim front port templates partial update operation typically these are written to a http.Request
*/
type DcimFrontPortTemplatesPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableFrontPortTemplate
	/*ID
	  A unique integer value identifying this front port template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithData(data *models.WritableFrontPortTemplate) *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetData(data *models.WritableFrontPortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithID(id int64) *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
