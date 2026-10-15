package core

import "net/http"

// CustomEndpointOption is returned by option.WithCustomEndpoint.
type CustomEndpointOption struct {
	BaseURL    string
	HTTPClient HTTPClient
}

func (c *CustomEndpointOption) applyRequestOptions(opts *RequestOptions) {
	opts.BaseURL = c.BaseURL
	opts.HTTPClient = &CustomEndpointClient{BaseURL: c.BaseURL, HTTPClient: c.HTTPClient}
}

// CustomEndpointClient is the HTTP client installed by option.WithCustomEndpoint.
// It permits requests to the origin of BaseURL, which needn't be a public Plaid host.
type CustomEndpointClient struct {
	BaseURL    string
	HTTPClient HTTPClient
}

func (c *CustomEndpointClient) Do(req *http.Request) (*http.Response, error) {
	if c.HTTPClient == nil {
		return http.DefaultClient.Do(req)
	}
	return c.HTTPClient.Do(req)
}
