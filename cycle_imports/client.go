package lokalise

import (
	"io"
	"os"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/lokalise/go-lokalise-api/teamuser"
)

type Client struct {
	timeout    time.Duration
	baseURL    string
	apiToken   string
	retryCount int
	httpClient *resty.Client
	logger     io.Writer

	TeamUsers teamuser.Service
}

type ClientOption func(*Client) error

//noinspection GoUnusedExportedFunction
func NewClient(apiToken string, options ...ClientOption) (*Client, error) {
	c := Client{
		apiToken:   apiToken,
		retryCount: 3,
		baseURL:    defaultBaseURL,
	}

	for _, o := range options {
		err := o(&c)
		if err != nil {
			return nil, err
		}
	}

	if c.logger == nil {
		c.logger = os.Stderr
	}

	c.httpClient = resty.New().
		SetHostURL(c.baseURL).
		SetRetryCount(c.retryCount).
		SetHeader(apiTokenHeader, c.apiToken).
		// SetLogger(c.logger).
		SetError(errorResponse{})

	c.TeamUsers = teamuser.Service{Client: &c}

	// Add other services here

	return &c, nil
}
