package gerrit

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/nexuer/ghttp"
)

type service struct {
	client *Client
}

type Options struct {
	UserAgent string
	Timeout   time.Duration
	Proxy     func(*http.Request) (*url.URL, error)
	Debug     bool
	TLS       *tls.Config
	Limiter   ghttp.Limiter
}

type Client struct {
	cc *ghttp.Client

	credential Credential

	common   service
	Projects *ProjectsService
}

func NewClient(credential Credential, opts ...*Options) *Client {
	c := &Client{}

	clientOpts := c.parseOptions(opts...)
	clientOpts = append(clientOpts,
		ghttp.WithNot2xxError(func() error {
			return new(Error)
		}),
	)
	c.cc = ghttp.NewClient(clientOpts...)

	c.common.client = c
	c.Projects = (*ProjectsService)(&c.common)

	c.SetCredential(credential)
	return c
}

func (c *Client) parseOptions(opts ...*Options) []ghttp.ClientOption {
	var opt *Options
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	} else {
		opt = new(Options)
	}

	clientOpts := make([]ghttp.ClientOption, 0)

	if opt.UserAgent != "" {
		clientOpts = append(clientOpts, ghttp.WithUserAgent(opt.UserAgent))
	}

	if opt.Debug {
		clientOpts = append(clientOpts, ghttp.WithDebug(true))
	}

	if opt.Timeout > 0 {
		clientOpts = append(clientOpts, ghttp.WithTimeout(opt.Timeout))
	}

	if opt.Proxy != nil {
		clientOpts = append(clientOpts, ghttp.WithProxy(opt.Proxy))
	}

	if opt.TLS != nil {
		clientOpts = append(clientOpts, ghttp.WithTLSConfig(opt.TLS))
	}

	if opt.Limiter != nil {
		clientOpts = append(clientOpts, ghttp.WithLimiter(opt.Limiter))
	}

	return clientOpts
}

func (c *Client) SetCredential(credential Credential) {
	var endpoint string
	if credential != nil {
		endpoint = credential.GetEndpoint()
		c.cc.SetEndpoint(endpoint)
	}
	c.credential = credential
}

var magicPrefix = []byte(")]}'\n")

func (c *Client) InvokeByCredential(ctx context.Context, method, path string, args any, reply any) (*http.Response, error) {
	if c.credential != nil {
		path = c.credential.AuthURL(path)
	}
	callOpts := &ghttp.CallOptions{
		BeforeHook: func(request *http.Request) error {
			if c.credential != nil {
				return c.credential.Auth(request)
			}
			return nil
		},
	}
	return c.Invoke(ctx, method, path, args, reply, callOpts)
}

func (c *Client) Invoke(ctx context.Context, method, path string, args any, reply any, callOpts ...*ghttp.CallOptions) (*http.Response, error) {
	callOpt := &ghttp.CallOptions{}
	if len(callOpts) > 0 && callOpts[0] != nil {
		callOpt = callOpts[0]
	}

	if method == http.MethodGet && args != nil {
		callOpt.Query = args
		args = nil
	}
	// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api.html#output
	callOpt.AfterHook = func(response *http.Response) error {
		all, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		_ = response.Body.Close()
		if bytes.HasPrefix(all, magicPrefix) {
			all = all[len(magicPrefix):]
		}
		response.Body = io.NopCloser(bytes.NewReader(all))
		return nil
	}

	return c.cc.Invoke(ctx, method, path, args, reply, callOpt)
}

type Error struct {
	Message string
}

func (e *Error) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	e.Message = string(data)
	return nil
}

func (e *Error) Error() string {
	return e.Message
}

type ListOptions struct {
	// Limit the number of projects to be included in the results.
	Limit int `query:"n,omitempty"`

	// Skip the given number of branches from the beginning of the list.
	Skip int `query:"S,omitempty"`
}

const defaultLimit = 25

func NewListOptions(s int, l ...int) ListOptions {
	if s <= 0 {
		s = 1
	}
	ll := defaultLimit
	if len(l) > 0 && l[0] > 0 {
		ll = l[0]
	}
	return ListOptions{
		Skip:  s,
		Limit: ll,
	}
}
