package social_channel

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

// NewSocialChannelFetchPopularTopicsParams creates a new SocialChannelFetchPopularTopicsParams object
// with the default values initialized.
func NewSocialChannelFetchPopularTopicsParams() *SocialChannelFetchPopularTopicsParams {
	var ()
	return &SocialChannelFetchPopularTopicsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSocialChannelFetchPopularTopicsParamsWithTimeout creates a new SocialChannelFetchPopularTopicsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSocialChannelFetchPopularTopicsParamsWithTimeout(timeout time.Duration) *SocialChannelFetchPopularTopicsParams {
	var ()
	return &SocialChannelFetchPopularTopicsParams{

		timeout: timeout,
	}
}

// NewSocialChannelFetchPopularTopicsParamsWithContext creates a new SocialChannelFetchPopularTopicsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSocialChannelFetchPopularTopicsParamsWithContext(ctx context.Context) *SocialChannelFetchPopularTopicsParams {
	var ()
	return &SocialChannelFetchPopularTopicsParams{

		Context: ctx,
	}
}

/*SocialChannelFetchPopularTopicsParams contains all the parameters to send to the API endpoint
for the social channel fetch popular topics operation typically these are written to a http.Request
*/
type SocialChannelFetchPopularTopicsParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the social channel fetch popular topics params
func (o *SocialChannelFetchPopularTopicsParams) WithTimeout(timeout time.Duration) *SocialChannelFetchPopularTopicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the social channel fetch popular topics params
func (o *SocialChannelFetchPopularTopicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the social channel fetch popular topics params
func (o *SocialChannelFetchPopularTopicsParams) WithContext(ctx context.Context) *SocialChannelFetchPopularTopicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the social channel fetch popular topics params
func (o *SocialChannelFetchPopularTopicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the social channel fetch popular topics params
func (o *SocialChannelFetchPopularTopicsParams) WithBody(body models.DefaultSelector) *SocialChannelFetchPopularTopicsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the social channel fetch popular topics params
func (o *SocialChannelFetchPopularTopicsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SocialChannelFetchPopularTopicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
