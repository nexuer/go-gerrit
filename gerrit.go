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
	Accounts *AccountsService
	Changes  *ChangesService
	Config   *ConfigService
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
	c.Accounts = (*AccountsService)(&c.common)
	c.Changes = (*ChangesService)(&c.common)
	c.Config = (*ConfigService)(&c.common)

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

func (c *Client) InvokeByCredential(ctx context.Context, method, path string, args any, reply any,
	delContentType ...bool) (*http.Response, error) {
	if c.credential != nil {
		path = c.credential.AuthURL(path)
	}
	callOpts := &ghttp.CallOptions{
		BeforeHook: func(request *http.Request) error {
			// errors api:
			// https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#set-active
			// https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#delete-active
			if len(delContentType) > 0 && delContentType[0] {
				request.Header.Del("Content-Type")
			}

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
	msg string
}

func (e *Error) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	e.msg = string(data)
	return nil
}

func (e *Error) Error() string {
	return e.msg
}

type ListOptions struct {
	// Limit the number of projects to be included in the results.
	Limit int `query:"n,omitempty"`

	// Skip the given number of branches from the beginning of the list.
	Skip int `query:"S,omitempty"`
}

const defaultLimit = 25

func NewListOptions(skip int, limits ...int) ListOptions {
	l := defaultLimit
	if len(limits) > 0 && limits[0] > 0 {
		l = limits[0]
	}
	return ListOptions{
		Skip:  skip,
		Limit: l,
	}
}

func IsNotFound(err error) bool {
	code, ok := StatusForErr(err)
	if ok && code == http.StatusNotFound {
		return true
	}
	return false
}

func IsForbidden(err error) bool {
	code, ok := StatusForErr(err)
	if ok && code == http.StatusForbidden {
		return true
	}
	return false
}

func IsUnauthorized(err error) bool {
	code, ok := StatusForErr(err)
	if ok && code == http.StatusUnauthorized {
		return true
	}
	return false
}

func IsTimeout(err error) bool {
	return ghttp.IsTimeout(err)
}

func StatusForErr(err error) (int, bool) {
	return ghttp.StatusForErr(err)
}
