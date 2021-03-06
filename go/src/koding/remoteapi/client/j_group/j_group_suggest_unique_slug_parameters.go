package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewJGroupSuggestUniqueSlugParams creates a new JGroupSuggestUniqueSlugParams object
// with the default values initialized.
func NewJGroupSuggestUniqueSlugParams() *JGroupSuggestUniqueSlugParams {
	var ()
	return &JGroupSuggestUniqueSlugParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupSuggestUniqueSlugParamsWithTimeout creates a new JGroupSuggestUniqueSlugParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupSuggestUniqueSlugParamsWithTimeout(timeout time.Duration) *JGroupSuggestUniqueSlugParams {
	var ()
	return &JGroupSuggestUniqueSlugParams{

		timeout: timeout,
	}
}

// NewJGroupSuggestUniqueSlugParamsWithContext creates a new JGroupSuggestUniqueSlugParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupSuggestUniqueSlugParamsWithContext(ctx context.Context) *JGroupSuggestUniqueSlugParams {
	var ()
	return &JGroupSuggestUniqueSlugParams{

		Context: ctx,
	}
}

/*JGroupSuggestUniqueSlugParams contains all the parameters to send to the API endpoint
for the j group suggest unique slug operation typically these are written to a http.Request
*/
type JGroupSuggestUniqueSlugParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j group suggest unique slug params
func (o *JGroupSuggestUniqueSlugParams) WithTimeout(timeout time.Duration) *JGroupSuggestUniqueSlugParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group suggest unique slug params
func (o *JGroupSuggestUniqueSlugParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group suggest unique slug params
func (o *JGroupSuggestUniqueSlugParams) WithContext(ctx context.Context) *JGroupSuggestUniqueSlugParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group suggest unique slug params
func (o *JGroupSuggestUniqueSlugParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group suggest unique slug params
func (o *JGroupSuggestUniqueSlugParams) WithBody(body models.DefaultSelector) *JGroupSuggestUniqueSlugParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group suggest unique slug params
func (o *JGroupSuggestUniqueSlugParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupSuggestUniqueSlugParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
