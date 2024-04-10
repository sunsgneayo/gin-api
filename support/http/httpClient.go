package http

import (
	"github.com/go-resty/resty/v2"
)

// HttpClient  is a simple HTTP wrapper.
type HttpClient struct {
	client *resty.Client
}

// NewHTTPClient  creates a new instance of HTTPClient.
func NewHTTPClient() *HttpClient {
	return &HttpClient{
		client: resty.New(),
	}
}

// Get sends an HTTP GET request and returns the response body and status code.
func (c *HttpClient) Get(url string) ([]byte, int, error) {
	resp, err := c.client.R().SetHeader("cookie", "6526522").Get(url)
	if err != nil {
		return nil, 0, err
	}
	return resp.Body(), resp.StatusCode(), nil
}
