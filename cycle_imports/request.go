package lokalise

import (
	"context"

	"github.com/go-resty/resty/v2"
)

const (
	apiTokenHeader = "X-Api-Token"
	defaultBaseURL = "https://api.lokalise.co/api2"
)

func (c *Client) Get(ctx context.Context, path string, res interface{}) (*resty.Response, error) {
	return c.req(ctx, path, res).Get(path)
}

func (c *Client) GetList(ctx context.Context, path string, res interface{}, options OptionsApplier) (*resty.Response, error) {
	req := c.req(ctx, path, res)
	options.Apply(req)
	return req.Get(path)
}

func (c *Client) Post(ctx context.Context, path string, res, body interface{}) (*resty.Response, error) {
	return c.reqWithBody(ctx, path, res, body).Post(path)
}

func (c *Client) Put(ctx context.Context, path string, res, body interface{}) (*resty.Response, error) {
	return c.reqWithBody(ctx, path, res, body).Put(path)
}

func (c *Client) Delete(ctx context.Context, path string, res interface{}) (*resty.Response, error) {
	return c.req(ctx, path, res).Delete(path)
}

func (c *Client) req(ctx context.Context, path string, res interface{}) *resty.Request {
	return c.httpClient.R().
		SetResult(&res).
		SetContext(ctx)
}

func (c *Client) reqWithBody(ctx context.Context, path string, res, body interface{}) *resty.Request {
	return c.req(ctx, path, res).SetBody(body)
}
