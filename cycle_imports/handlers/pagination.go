package handlers

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type OptionsApplier interface {
	Apply(req *resty.Request)
}

type PageOptions struct {
	Limit int64
	Page  int64
}

func (options *PageOptions) Apply(req *resty.Request) {
	if options.Limit != 0 {
		req.SetQueryParam("limit", fmt.Sprintf("%d", options.Limit))
	}
	if options.Page != 0 {
		req.SetQueryParam("page", fmt.Sprintf("%d", options.Page))
	}
}

type Paged struct {
	TotalCount int64
	PageCount  int64
	Limit      int64
	Page       int64
}
