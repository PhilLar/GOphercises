package teamuser

import (
	"github.com/lokalise/go-lokalise-api/Env"
	"github.com/lokalise/go-lokalise-api/handlers"
	"io"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	timeout    time.Duration
	baseURL    string
	apiToken   string
	retryCount int
	httpClient *resty.Client
	logger     io.Writer

	TeamUsers Service
}

type ClientOption func(*Client) error

//noinspection GoUnusedExportedFunction
func NewClient(apiToken string, options ...ClientOption) (*Client, error) {
	c := Client{
		apiToken:   apiToken,
		retryCount: 3,
		baseURL:    Env.DefaultBaseURL,
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
		SetHeader(Env.ApiTokenHeader, c.apiToken).
		// SetLogger(c.logger).
		SetError(handlers.ErrorResponse{})

	c.TeamUsers = Service{Client: &c}

	// Add other services here

	return &c, nil
}
