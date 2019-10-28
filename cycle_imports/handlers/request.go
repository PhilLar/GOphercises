package handlers

import (
	"context"
	"github.com/go-resty/resty/v2"
)

//type MockRestyResponse struct {
//	Request     *resty.Request
//	RawResponse *http.Response
//
//	body       []byte
//	size       int64
//	receivedAt time.Time
//}
//
//type MockReq struct {
//	Request	string
//	Path	string
//}
//
//func (req *MockReq) Get() *resty.Response {
//	resp := &MockRestyResponse{}
//	return resp
//}

type Client interface {
	req(c context.Context, path string, res interface{}) MockReq
}



func Get(c Client, ctx context.Context, path string, res interface{}) (*resty.Response, error) {
	return c.req(ctx, path, res).Get(path), nil
}

func GetList(c Client, ctx context.Context, path string, res interface{}, options OptionsApplier) (*resty.Response, error) {
	req := c.req(ctx, path, res)
	lokalise.Apply(req)
	return req.Get(path), nil
}

func Post(c Client, ctx context.Context, path string, res, body interface{}) (*resty.Response, error) {
	return c.reqWithBody(ctx, path, res, body).Post(path), nil
}

func Put(c Client, ctx context.Context, path string, res, body interface{}) (*resty.Response, error) {
	return c.reqWithBody(ctx, path, res, body).Put(path), nil
}

func Delete(c Client, ctx context.Context, path string, res interface{}) (*resty.Response, error) {
	return c.req(ctx, path, res).Delete(path), nil
}

func req(c Client, ctx context.Context, path string, res interface{}) *resty.Request {
	return c.httpClient.R().
		SetResult(&res).
		SetContext(ctx)
}

func reqWithBody(c Client, ctx context.Context, path string, res, body interface{}) *resty.Request {
	return c.req(ctx, path, res).SetBody(body)
}
