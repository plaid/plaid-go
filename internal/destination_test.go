package internal

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"

	"github.com/plaid/plaid-go/v49/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckDestination(t *testing.T) {
	allowed := []string{
		"https://production.plaid.com",
		"https://sandbox.plaid.com/accounts/get",
		"https://production.plaid.com:443",
		"https://PRODUCTION.plaid.com",
		"https://production.plaid.com.",
	}
	for _, rawURL := range allowed {
		assert.NoError(t, checkDestination(mustParseURL(t, rawURL), nil), rawURL)
	}

	rejected := []string{
		"http://localhost:8080",
		"https://localhost:8443",
		"http://127.0.0.1:55000",
		"http://127.2.3.4:8080",
		"http://[::1]:9000",
		"http://[::ffff:127.0.0.1]:9000",
		"https://production2.plaid.com",
		"https://api-mtls-private.plaid.com",
		"http://production.plaid.com",
		"https://production.plaid.com:8443",
		"https://plaid.com",
		"https://dashboard.plaid.com",
		"https://production.plaid.com.example.com",
		"https://example.com",
		"https://localhost.example.com",
		"http://10.0.0.1",
		"production.plaid.com",
		"",
	}
	for _, rawURL := range rejected {
		assert.Error(t, checkDestination(mustParseURL(t, rawURL), nil), rawURL)
	}
}

func TestCallRejectsDestinationBeforeSending(t *testing.T) {
	var requests atomic.Int32
	client := doerFunc(func(*http.Request) (*http.Response, error) {
		requests.Add(1)
		return nil, errors.New("unexpected request")
	})
	caller := NewCaller(&CallerParams{Client: client})

	_, err := caller.Call(context.Background(), &CallParams{
		URL:     "https://example.com/accounts/get",
		Method:  http.MethodPost,
		Headers: http.Header{"Plaid-Secret": []string{"secret"}},
	})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "refusing to send a request")
	assert.Zero(t, requests.Load())
}

func TestCallRefusesRedirects(t *testing.T) {
	var targetRequests atomic.Int32
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		targetRequests.Add(1)
	}))
	defer target.Close()
	redirector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL+r.URL.Path, http.StatusTemporaryRedirect)
	}))
	defer redirector.Close()

	httpClient := &http.Client{}
	caller := NewCaller(&CallerParams{Client: &core.CustomEndpointClient{BaseURL: redirector.URL, HTTPClient: httpClient}})
	_, err := caller.Call(context.Background(), &CallParams{
		URL:     redirector.URL + "/accounts/get",
		Method:  http.MethodPost,
		Headers: http.Header{"Plaid-Secret": []string{"secret"}},
	})

	require.ErrorIs(t, err, errRedirect)
	assert.Zero(t, targetRequests.Load())
	assert.Nil(t, httpClient.CheckRedirect, "the caller's client must not be modified")
}

func TestCheckDestinationAllowsOnlyTheCustomEndpoint(t *testing.T) {
	internalAPI := &core.CustomEndpointClient{BaseURL: "http://api_v2"}
	assert.NoError(t, checkDestination(mustParseURL(t, "http://api_v2/accounts/get"), internalAPI))
	for _, rawURL := range []string{
		"https://api_v2/accounts/get",
		"http://api_v2:8080/accounts/get",
		"http://api_v3/accounts/get",
		"https://example.com/accounts/get",
	} {
		assert.Error(t, checkDestination(mustParseURL(t, rawURL), internalAPI), rawURL)
	}

	mtls := &core.CustomEndpointClient{BaseURL: "https://api-mtls-private.plaid.com"}
	assert.NoError(t, checkDestination(mustParseURL(t, "https://api-mtls-private.plaid.com/accounts/get"), mtls))
	assert.Error(t, checkDestination(mustParseURL(t, "https://production2.plaid.com/accounts/get"), mtls))

	testEnv := &core.CustomEndpointClient{BaseURL: "https://api-v2.example.test"}
	assert.NoError(t, checkDestination(mustParseURL(t, "https://api-v2.example.test/accounts/get"), testEnv))
	assert.NoError(t, checkDestination(mustParseURL(t, "https://sandbox.plaid.com/accounts/get"), testEnv))
	assert.Error(t, checkDestination(mustParseURL(t, "https://api-v2.example.test/accounts/get"), &core.CustomEndpointClient{BaseURL: "not a url"}))
}

func TestCustomEndpointStillRefusesRedirects(t *testing.T) {
	var targetRequests atomic.Int32
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		targetRequests.Add(1)
	}))
	defer target.Close()
	redirector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL, http.StatusFound)
	}))
	defer redirector.Close()

	for _, inner := range []core.HTTPClient{nil, &http.Client{}} {
		caller := NewCaller(&CallerParams{Client: &core.CustomEndpointClient{BaseURL: redirector.URL, HTTPClient: inner}})
		_, err := caller.Call(context.Background(), &CallParams{URL: redirector.URL + "/accounts/get", Method: http.MethodPost})
		require.ErrorIs(t, err, errRedirect)
	}
	assert.Zero(t, targetRequests.Load())
}

type doerFunc func(*http.Request) (*http.Response, error)

func (f doerFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func mustParseURL(t *testing.T, rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	require.NoError(t, err)
	return u
}
