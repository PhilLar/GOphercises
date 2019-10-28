package handlers

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
)

const (
	headerTotalCount = "X-Pagination-Total-Count"
	headerPageCount  = "X-Pagination-Page-Count"
	headerLimit      = "X-Pagination-Limit"
	headerPage       = "X-Pagination-Page"
)

type ErrorResponse struct {
	Error Error `json:"error"`
}

// Error is an API error.
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r Error) Error() string {
	return fmt.Sprintf("API request error %d %s", r.Code, r.Message)
}

func ApiError(res *resty.Response) error {
	if !res.IsError() {
		return nil
	}
	responseError := res.Error()
	if responseError == nil {
		return errors.New("lokalise: response marked as error but no data returned")
	}
	responseErrorModel, ok := responseError.(*ErrorResponse)
	if !ok {
		return errors.New("lokalise: response error model unknown")
	}
	return responseErrorModel.Error
}

func ApplyPaged(res *resty.Response, paged *Paged) {
	headers := res.Header()
	paged.TotalCount = headerInt64(headers, headerTotalCount)
	paged.PageCount = headerInt64(headers, headerPageCount)
	paged.Limit = headerInt64(headers, headerLimit)
	paged.Page = headerInt64(headers, headerPage)
}

func headerInt64(headers http.Header, headerKey string) int64 {
	headerValue := headers.Get(headerKey)
	if headerValue == "" {
		return -1
	}
	value, err := strconv.ParseInt(headers.Get(headerKey), 10, 64)
	if err != nil {
		return -1
	}
	return value
}
