package internal

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/plaid/plaid-go/v49/core"
)

// Every Plaid API request carries the client ID and secret, so the SDK only
// sends requests to Plaid's public API hosts unless the caller names another
// endpoint with option.WithCustomEndpoint.
var plaidAPIHosts = map[string]bool{
	"production.plaid.com": true,
	"sandbox.plaid.com":    true,
}

// errRedirect is returned instead of following a redirect, which would resend
// the credential headers, possibly to another host. Plaid's API doesn't redirect.
var errRedirect = errors.New("plaid: refusing to follow a redirect")

// checkDestination allows a Plaid API host over HTTPS on the default port or
// the endpoint a caller explicitly allowed with option.WithCustomEndpoint.
func checkDestination(u *url.URL, client core.HTTPClient) error {
	host := strings.TrimSuffix(strings.ToLower(u.Hostname()), ".")
	if u.Scheme == "https" && plaidAPIHosts[host] && (u.Port() == "" || u.Port() == "443") {
		return nil
	}
	if customClient, ok := client.(*core.CustomEndpointClient); ok && sameOrigin(u, customClient.BaseURL) {
		return nil
	}
	return fmt.Errorf("plaid: refusing to send a request to %q: the base URL must be https://production.plaid.com or https://sandbox.plaid.com, or an endpoint named with option.WithCustomEndpoint", u.Redacted())
}

func sameOrigin(u *url.URL, baseURL string) bool {
	allowed, err := url.Parse(baseURL)
	if err != nil || allowed.Host == "" {
		return false
	}
	return strings.EqualFold(u.Scheme, allowed.Scheme) && strings.EqualFold(u.Host, allowed.Host)
}

// withoutRedirects returns a copy of an *http.Client that refuses redirects,
// leaving the caller's client unchanged, including inside the client installed
// by option.WithCustomEndpoint. Other HTTPClient implementations are returned
// as is and are responsible for their own redirect handling.
func withoutRedirects(client core.HTTPClient) core.HTTPClient {
	if customClient, ok := client.(*core.CustomEndpointClient); ok {
		inner := customClient.HTTPClient
		if inner == nil {
			inner = http.DefaultClient
		}
		return &core.CustomEndpointClient{BaseURL: customClient.BaseURL, HTTPClient: withoutRedirects(inner)}
	}
	httpClient, ok := client.(*http.Client)
	if !ok {
		return client
	}
	clone := *httpClient
	clone.CheckRedirect = func(*http.Request, []*http.Request) error {
		return errRedirect
	}
	return &clone
}
