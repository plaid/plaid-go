package option

import (
	core "github.com/plaid/plaid-go/v49/core"
)

// WithCustomEndpoint sends requests to baseURL instead of a public Plaid
// environment, using httpClient (nil for http.DefaultClient). Use it for a
// dedicated endpoint Plaid has set you up on, such as an mTLS endpoint, or for
// an endpoint you control. Requests carry your client ID and secret, and only
// that endpoint is allowed. This opt-in is also required for local mock servers.
// Redirects are refused when httpClient is nil or an *http.Client. Other
// core.HTTPClient implementations are trusted to enforce their own redirect policy.
//
// Pass your HTTP client here rather than with WithHTTPClient, which would
// replace it and remove the permission.
func WithCustomEndpoint(baseURL string, httpClient core.HTTPClient) *core.CustomEndpointOption {
	return &core.CustomEndpointOption{
		BaseURL:    baseURL,
		HTTPClient: httpClient,
	}
}
