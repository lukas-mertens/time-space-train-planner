// Code generated by go-swagger; DO NOT EDIT.

package operations

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
)

// NewGetRchgEvaNoParams creates a new GetRchgEvaNoParams object
// with the default values initialized.
func NewGetRchgEvaNoParams() *GetRchgEvaNoParams {
	var ()
	return &GetRchgEvaNoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRchgEvaNoParamsWithTimeout creates a new GetRchgEvaNoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRchgEvaNoParamsWithTimeout(timeout time.Duration) *GetRchgEvaNoParams {
	var ()
	return &GetRchgEvaNoParams{

		timeout: timeout,
	}
}

// NewGetRchgEvaNoParamsWithContext creates a new GetRchgEvaNoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRchgEvaNoParamsWithContext(ctx context.Context) *GetRchgEvaNoParams {
	var ()
	return &GetRchgEvaNoParams{

		Context: ctx,
	}
}

// NewGetRchgEvaNoParamsWithHTTPClient creates a new GetRchgEvaNoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRchgEvaNoParamsWithHTTPClient(client *http.Client) *GetRchgEvaNoParams {
	var ()
	return &GetRchgEvaNoParams{
		HTTPClient: client,
	}
}

/*GetRchgEvaNoParams contains all the parameters to send to the API endpoint
for the get rchg eva no operation typically these are written to a http.Request
*/
type GetRchgEvaNoParams struct {

	/*EvaNo
	  Station EVA-number.

	*/
	EvaNo string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rchg eva no params
func (o *GetRchgEvaNoParams) WithTimeout(timeout time.Duration) *GetRchgEvaNoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rchg eva no params
func (o *GetRchgEvaNoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rchg eva no params
func (o *GetRchgEvaNoParams) WithContext(ctx context.Context) *GetRchgEvaNoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rchg eva no params
func (o *GetRchgEvaNoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rchg eva no params
func (o *GetRchgEvaNoParams) WithHTTPClient(client *http.Client) *GetRchgEvaNoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rchg eva no params
func (o *GetRchgEvaNoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEvaNo adds the evaNo to the get rchg eva no params
func (o *GetRchgEvaNoParams) WithEvaNo(evaNo string) *GetRchgEvaNoParams {
	o.SetEvaNo(evaNo)
	return o
}

// SetEvaNo adds the evaNo to the get rchg eva no params
func (o *GetRchgEvaNoParams) SetEvaNo(evaNo string) {
	o.EvaNo = evaNo
}

// WriteToRequest writes these params to a swagger request
func (o *GetRchgEvaNoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param evaNo
	if err := r.SetPathParam("evaNo", o.EvaNo); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}