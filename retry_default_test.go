package plaid_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync/atomic"
	"testing"

	plaid "github.com/plaid/plaid-go/v49"
	"github.com/plaid/plaid-go/v49/client"
	"github.com/plaid/plaid-go/v49/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Retrying a non-idempotent call such as /transfer/create can repeat it, so
// the client makes one attempt unless the caller opts in with WithMaxAttempts.
func TestNonIdempotentCallIsAttemptedOnceByDefault(t *testing.T) {
	for _, status := range []int{
		http.StatusRequestTimeout,
		http.StatusTooManyRequests,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			var requests atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests.Add(1)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(status)
				_, _ = w.Write([]byte(`{"error_type":"API_ERROR","error_code":"INTERNAL_SERVER_ERROR"}`))
			}))
			defer server.Close()

			c := client.NewClient(
				option.WithCustomEndpoint(server.URL, server.Client()),
				option.WithClientID(plaid.String("client-id")),
				option.WithSecret("secret"),
			)
			_, err := c.Transfer.Create(context.Background(), &plaid.TransferCreateRequest{
				AccessToken:     "access-sandbox",
				AccountID:       "account-id",
				AuthorizationID: "authorization-id",
			})

			require.Error(t, err)
			assert.Equal(t, int32(1), requests.Load())
		})
	}
}
