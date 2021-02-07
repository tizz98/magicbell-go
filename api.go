package magicbell

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultAPIURL  = "https://api.magicbell.io"
	defaultTimeout = 5 * time.Second

	apiKeyHeader    = "X-MAGICBELL-API-KEY"
	apiSecretHeader = "X-MAGICBELL-API-SECRET"
)

// New instantiates a new API which implements the IAPI interface.
// Config.APIKey and Config.APISecret must be set in order for API requests to succeed.
func New(config Config) IAPI {
	if config.BaseURL == "" {
		config.BaseURL = defaultAPIURL
	}
	if config.Timeout == nil {
		config.Timeout = newDuration(defaultTimeout)
	}

	api := &API{config: config}
	api.client = &http.Client{
		Transport: api,
		Timeout:   *config.Timeout,
	}
	return api
}

// API implements the IAPI interface for making HTTP requests
// to the MagicBell API. Use New to instantiate this struct.
type API struct {
	config Config
	client *http.Client
}

// Config represents the required values to make HTTP requests to
// the MagicBell API.
type Config struct {
	// APIKey is the api key for your MagicBell account
	APIKey string
	// APISecret is the api secret for your MagicBell account
	APISecret string
	// BaseURL is the MagicBell API url, this is optional
	// and will default to https://api.magicbell.io
	BaseURL string
	// Timeout is an optional time.Duration to wait for HTTP requests to timeout.
	// If not provided, it will default to 5 seconds.
	Timeout *time.Duration // optional
}

func (c *Config) withBaseURL(url string) Config {
	c2 := *c
	c2.BaseURL = url
	return c2
}

// RoundTrip implements the http.RoundTripper interface and is used for the API's http.Client http.Transport.
// This method will automatically add the Config.APIKey and Config.APISecret as headers for all HTTP requests.
// After doing so, the http.DefaultTransport will be used to finish making the request.
func (a *API) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set(apiKeyHeader, a.config.APIKey)
	r.Header.Set(apiSecretHeader, a.config.APISecret)
	return http.DefaultTransport.RoundTrip(r)
}

func (a *API) makeRequest(ctx context.Context, method string, endpoint string, requestBody interface{}, out interface{}) error {
	var bodyReader io.Reader

	if requestBody != nil {
		serializedBody, err := json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("magicbell-go/api: error serializing request body: %w", err)
		}

		bodyReader = bytes.NewReader(serializedBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", a.config.BaseURL, endpoint), bodyReader)
	if err != nil {
		return fmt.Errorf("magicbell-go/api: error creating http request: %w", err)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("magicbell-go/api: error making http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		var bodyStr string
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			bodyStr = string(body)
		}

		return InternalServerError{
			StatusCode: resp.StatusCode,
			Body:       bodyStr,
		}
	} else if resp.StatusCode >= 400 {
		// todo: does this need to be handled explicitly?
	} else if resp.StatusCode == http.StatusNoContent {
		// special case with no response body
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("magicbell-go/api: error decoding response json: %w", err)
	}
	return nil
}

func newDuration(d time.Duration) *time.Duration { return &d }
