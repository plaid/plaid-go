package tests

import (
	"context"
	"os"
	"testing"

	"github.com/plaid/plaid-go/v16/plaid"
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
