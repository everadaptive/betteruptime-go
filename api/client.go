package api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context/ctxhttp"
)

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
	userAgent  string
}

type option func(c *Client)

func WithHTTPClient(httpClient *http.Client) option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithUserAgent(userAgent string) option {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

func NewClient(token string, opts ...option) (*Client, error) {
	baseURL := "https://betteruptime.com"

	c := Client{
		baseURL:    baseURL,
		token:      token,
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return &c, nil
}

func (c *Client) Get(ctx context.Context, path string) (*http.Response, error) {
	return c.do(ctx, http.MethodGet, path, nil)
}

func (c *Client) Post(ctx context.Context, path string, body io.Reader) (*http.Response, error) {
	return c.do(ctx, http.MethodPost, path, body)
}

func (c *Client) Patch(ctx context.Context, path string, body io.Reader) (*http.Response, error) {
	return c.do(ctx, http.MethodPatch, path, body)
}

func (c *Client) Delete(ctx context.Context, path string) (*http.Response, error) {
	return c.do(ctx, http.MethodDelete, path, nil)
}

func (c *Client) do(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.baseURL, path), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	if method == http.MethodPost || method == http.MethodPatch {
		req.Header.Set("Content-Type", "application/json")
	}
	return ctxhttp.Do(ctx, c.httpClient, req)
}
