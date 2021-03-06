package j_user

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

// NewPostRemoteAPIJUserAuthenticateWithOauthParams creates a new PostRemoteAPIJUserAuthenticateWithOauthParams object
// with the default values initialized.
func NewPostRemoteAPIJUserAuthenticateWithOauthParams() *PostRemoteAPIJUserAuthenticateWithOauthParams {
	var ()
	return &PostRemoteAPIJUserAuthenticateWithOauthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJUserAuthenticateWithOauthParamsWithTimeout creates a new PostRemoteAPIJUserAuthenticateWithOauthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJUserAuthenticateWithOauthParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJUserAuthenticateWithOauthParams {
	var ()
	return &PostRemoteAPIJUserAuthenticateWithOauthParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJUserAuthenticateWithOauthParamsWithContext creates a new PostRemoteAPIJUserAuthenticateWithOauthParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJUserAuthenticateWithOauthParamsWithContext(ctx context.Context) *PostRemoteAPIJUserAuthenticateWithOauthParams {
	var ()
	return &PostRemoteAPIJUserAuthenticateWithOauthParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJUserAuthenticateWithOauthParams contains all the parameters to send to the API endpoint
for the post remote API j user authenticate with oauth operation typically these are written to a http.Request
*/
type PostRemoteAPIJUserAuthenticateWithOauthParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j user authenticate with oauth params
func (o *PostRemoteAPIJUserAuthenticateWithOauthParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJUserAuthenticateWithOauthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j user authenticate with oauth params
func (o *PostRemoteAPIJUserAuthenticateWithOauthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j user authenticate with oauth params
func (o *PostRemoteAPIJUserAuthenticateWithOauthParams) WithContext(ctx context.Context) *PostRemoteAPIJUserAuthenticateWithOauthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j user authenticate with oauth params
func (o *PostRemoteAPIJUserAuthenticateWithOauthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j user authenticate with oauth params
func (o *PostRemoteAPIJUserAuthenticateWithOauthParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJUserAuthenticateWithOauthParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j user authenticate with oauth params
func (o *PostRemoteAPIJUserAuthenticateWithOauthParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJUserAuthenticateWithOauthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
