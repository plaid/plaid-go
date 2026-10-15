package plaid_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	plaid "github.com/plaid/plaid-go/v49"
	"github.com/plaid/plaid-go/v49/client"
	"github.com/plaid/plaid-go/v49/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Guards the destination check in internal/caller.go, a generated file kept
// through regeneration only by .fernignore, against being lost.
func TestClientRefusesNonPlaidBaseURL(t *testing.T) {
	var requests atomic.Int32
	httpClient := doerFunc(func(*http.Request) (*http.Response, error) {
		requests.Add(1)
		return nil, errors.New("unexpected request")
	})

	for _, baseURL := range []string{
		"https://example.com", "http://production.plaid.com",
		"http://localhost:8080", "https://localhost:8443",
		"http://127.0.0.1:8080", "http://127.2.3.4:8080",
		"http://[::1]:8080", "http://[::ffff:127.0.0.1]:8080",
	} {
		c := client.NewClient(
			option.WithBaseURL(baseURL),
			option.WithClientID(plaid.String("client-id")),
			option.WithSecret("secret"),
			option.WithHTTPClient(httpClient),
		)
		_, err := c.Accounts.Get(context.Background(), &plaid.AccountsGetRequest{AccessToken: "access-sandbox"})
		require.Error(t, err, baseURL)
		assert.Contains(t, err.Error(), "refusing to send a request", baseURL)
	}
	assert.Zero(t, requests.Load())
}

func TestLoopbackRequiresExplicitOptIn(t *testing.T) {
	t.Setenv("PLAID_CLIENT_ID", "canary-env-client-id")
	t.Setenv("PLAID_SECRET", "canary-env-secret")
	var requests atomic.Int32
	var clientID, secret string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID, secret = r.Header.Get("Plaid-Client-Id"), r.Header.Get("Plaid-Secret")
		requests.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"accounts":[],"item":{},"request_id":"fixture"}`))
	}))
	defer server.Close()
	request := &plaid.AccountsGetRequest{AccessToken: "access-sandbox"}

	defaultClient := client.NewClient(option.WithBaseURL(server.URL))
	_, err := defaultClient.Accounts.Get(context.Background(), request)
	require.Error(t, err, "a local server must not receive environment credentials without explicit opt-in")
	assert.Contains(t, err.Error(), "WithCustomEndpoint")
	assert.Zero(t, requests.Load())

	dummyID := "wiremock-client-id"
	explicitClient := client.NewClient(
		option.WithCustomEndpoint(server.URL, server.Client()),
		option.WithClientID(&dummyID),
		option.WithSecret("wiremock-secret"),
	)
	_, err = explicitClient.Accounts.Get(context.Background(), request)
	require.NoError(t, err)
	assert.Equal(t, int32(1), requests.Load())
	assert.Equal(t, dummyID, clientID)
	assert.Equal(t, "wiremock-secret", secret)
}

func TestCustomEndpointOption(t *testing.T) {
	var hosts []string
	httpClient := doerFunc(func(req *http.Request) (*http.Response, error) {
		hosts = append(hosts, req.URL.Host)
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		}, nil
	})
	const internalURL = "https://api-v2.example.test"
	request := &plaid.AccountsGetRequest{AccessToken: "access-sandbox"}

	c := client.NewClient(option.WithCustomEndpoint(internalURL, httpClient), option.WithSecret("secret"))
	_, err := c.Accounts.Get(context.Background(), request)
	require.NoError(t, err)

	_, err = c.Accounts.Get(context.Background(), request, option.WithBaseURL("https://example.com"))
	require.Error(t, err, "the permission covers only the endpoint it names")

	c = client.NewClient(option.WithBaseURL(internalURL), option.WithHTTPClient(httpClient), option.WithSecret("secret"))
	_, err = c.Accounts.Get(context.Background(), request)
	require.Error(t, err, "WithBaseURL alone must not allow a non-Plaid host")
	assert.Contains(t, err.Error(), "option.WithCustomEndpoint")

	assert.Equal(t, []string{"api-v2.example.test"}, hosts)
}

type doerFunc func(*http.Request) (*http.Response, error)

func (f doerFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}
