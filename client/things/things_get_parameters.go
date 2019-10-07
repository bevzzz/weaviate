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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewThingsGetParams creates a new ThingsGetParams object
// with the default values initialized.
func NewThingsGetParams() *ThingsGetParams {
	var ()
	return &ThingsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThingsGetParamsWithTimeout creates a new ThingsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThingsGetParamsWithTimeout(timeout time.Duration) *ThingsGetParams {
	var ()
	return &ThingsGetParams{

		timeout: timeout,
	}
}

// NewThingsGetParamsWithContext creates a new ThingsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewThingsGetParamsWithContext(ctx context.Context) *ThingsGetParams {
	var ()
	return &ThingsGetParams{

		Context: ctx,
	}
}

// NewThingsGetParamsWithHTTPClient creates a new ThingsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThingsGetParamsWithHTTPClient(client *http.Client) *ThingsGetParams {
	var ()
	return &ThingsGetParams{
		HTTPClient: client,
	}
}

/*ThingsGetParams contains all the parameters to send to the API endpoint
for the things get operation typically these are written to a http.Request
*/
type ThingsGetParams struct {

	/*ID
	  Unique ID of the Thing.

	*/
	ID strfmt.UUID
	/*Meta
	  Should additional meta information (e.g. about classified properties) be included? Defaults to false.

	*/
	Meta *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the things get params
func (o *ThingsGetParams) WithTimeout(timeout time.Duration) *ThingsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the things get params
func (o *ThingsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the things get params
func (o *ThingsGetParams) WithContext(ctx context.Context) *ThingsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the things get params
func (o *ThingsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the things get params
func (o *ThingsGetParams) WithHTTPClient(client *http.Client) *ThingsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the things get params
func (o *ThingsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the things get params
func (o *ThingsGetParams) WithID(id strfmt.UUID) *ThingsGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the things get params
func (o *ThingsGetParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithMeta adds the meta to the things get params
func (o *ThingsGetParams) WithMeta(meta *bool) *ThingsGetParams {
	o.SetMeta(meta)
	return o
}

// SetMeta adds the meta to the things get params
func (o *ThingsGetParams) SetMeta(meta *bool) {
	o.Meta = meta
}

// WriteToRequest writes these params to a swagger request
func (o *ThingsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Meta != nil {

		// query param meta
		var qrMeta bool
		if o.Meta != nil {
			qrMeta = *o.Meta
		}
		qMeta := swag.FormatBool(qrMeta)
		if qMeta != "" {
			if err := r.SetQueryParam("meta", qMeta); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
