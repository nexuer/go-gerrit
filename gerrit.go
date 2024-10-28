package gerrit

import (
	"bytes"
	"context"
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

func (c *Client) InvokeByCredential(ctx context.Context, method, path string, args any, reply any) error {
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

func (c *Client) Invoke(ctx context.Context, method, path string, args any, reply any, callOpts ...*ghttp.CallOptions) error {
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

	_, err := c.cc.Invoke(ctx, method, path, args, reply, callOpt)
	return err
}
