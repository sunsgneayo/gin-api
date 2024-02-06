package client

import (
	"bytes"
	"io"
	"net/http"
)

// HTTPClient is a simple HTTP client wrapper.
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new instance of HTTPClient.
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}

// Get sends an HTTP GET request and returns the response body and status code.
func (c *HTTPClient) Get(url string) ([]byte, int, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

// Post sends an HTTP POST request with the given payload and returns the response body and status code.
func (c *HTTPClient) Post(url string, payload []byte) ([]byte, int, error) {
	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
