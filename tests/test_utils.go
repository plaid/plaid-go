package tests

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/plaid/plaid-go/v42/plaid"
	"github.com/stretchr/testify/assert"
)

const (
	FIRST_PLATYPUS_BANK = "ins_109508"
	PLATYPUS_OAUTH_BANK = "ins_127287"
	iso8601TimeFormat   = "2006-01-02"
	RFC3339Nano         = "2006-01-02T15:04:05.999999999Z07:00"
)

var testProducts = []plaid.Products{
	plaid.PRODUCTS_AUTH,
	plaid.PRODUCTS_IDENTITY,
	plaid.PRODUCTS_TRANSACTIONS,
}

func NewTestClient() *plaid.APIClient {
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", os.Getenv("CLIENT_ID"))
	configuration.AddDefaultHeader("PLAID-SECRET", os.Getenv("SECRET"))
	if plaidApiUri, ok := os.LookupEnv("PLAID_API_URI"); ok {
		configuration.UseEnvironment(plaid.Environment(plaidApiUri))
	} else {
		configuration.UseEnvironment(plaid.Sandbox)
	}

	return plaid.NewAPIClient(configuration)
}

func PlaidEnv() string {
	if plaidEnv, ok := os.LookupEnv("PLAID_ENV"); ok {
		return plaidEnv
	} else {
		return "sandbox"
	}
}

// retryOnRateLimit retries an API call once after a 10s wait if it receives a
// 429 response. The caller passes a closure that executes the API call and
// returns the *http.Response; the closure should assign typed results to
// variables in the enclosing scope.
func retryOnRateLimit(t *testing.T, fn func() *http.Response) {
	httpResp := fn()
	if httpResp != nil && httpResp.StatusCode == 429 {
		t.Log("Rate limited (429), retrying after 10s...")
		time.Sleep(10 * time.Second)
		fn()
	}
}

func createSandboxItem(t *testing.T, ctx context.Context, client *plaid.APIClient, institutionID string, products []plaid.Products) string {
	// generate a sandbox public_token
	sandboxPublicTokenResp, _, err := client.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
		*plaid.NewSandboxPublicTokenCreateRequest(
			institutionID,
			products,
		),
	).Execute()

	assert.NotNil(t, sandboxPublicTokenResp)
	assert.NoError(t, err)

	// exchange the public_token for an access_token
	exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(sandboxPublicTokenResp.GetPublicToken()),
	).Execute()

	assert.NotNil(t, exchangePublicTokenResp)
	assert.NoError(t, err)

	return exchangePublicTokenResp.GetAccessToken()
}

func createSandboxItemWithOptions(t *testing.T, ctx context.Context, client *plaid.APIClient, institutionID string, products []plaid.Products, options *plaid.SandboxPublicTokenCreateRequestOptions) string {

	req := *plaid.NewSandboxPublicTokenCreateRequest(
		institutionID,
		products,
	)
	req.Options = options

	// generate a sandbox public_token
	sandboxPublicTokenResp, _, err := client.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(req).Execute()

	assert.NotNil(t, sandboxPublicTokenResp)
	assert.NoError(t, err)

	// exchange the public_token for an access_token
	exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(sandboxPublicTokenResp.GetPublicToken()),
	).Execute()

	assert.NotNil(t, exchangePublicTokenResp)
	assert.NoError(t, err)

	return exchangePublicTokenResp.GetAccessToken()
}

func importSandboxItem(t *testing.T, ctx context.Context, client *plaid.APIClient, userID string, authToken string, products []plaid.Products) string {
	accessTokenResp, _, err := client.PlaidApi.ItemImport(ctx).ItemImportRequest(
		*plaid.NewItemImportRequest(
			products,
			*plaid.NewItemImportRequestUserAuth(userID, authToken),
		),
	).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, accessTokenResp.GetAccessToken())
	return accessTokenResp.GetAccessToken()
}
