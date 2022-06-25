package stormglass

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// BaseURLV2 https://docs.stormglass.io/#/?id=api-endpoint.
	BaseURLV2 = "https://api.stormglass.io/v2"
)

// Client for accessing StormGlass API.
type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

// NewClient returns a new Client with default config.
func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURLV2,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", c.apiKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return NewError(res)
	}

	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return fmt.Errorf("decode error %w", err)
	}

	return nil
}
