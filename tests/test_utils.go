package tests

import (
	"context"
	"os"
	"testing"

	plaid "github.com/plaid/plaid-go"
	"github.com/stretchr/testify/require"
)

const (
	FIRST_PLATYPUS_BANK = "ins_109508"
	PLATYPUS_OAUTH_BANK = "ins_127287"
)

func NewTestClient() *plaid.APIClient {
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", os.Getenv("CLIENT_ID"))
	configuration.AddDefaultHeader("PLAID-SECRET", os.Getenv("SECRET"))
	configuration.UseEnvironment(plaid.Sandbox)

	return plaid.NewAPIClient(configuration)
}

func createSandboxItem(t *testing.T, ctx context.Context, client *plaid.APIClient, institutionID string, products []plaid.Products) string {
	// generate a sandbox public_token
	sandboxPublicTokenResp, _, err := client.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
		*plaid.NewSandboxPublicTokenCreateRequest(
			institutionID,
			products,
		),
	).Execute()

	require.NotNil(t, sandboxPublicTokenResp)
	require.NoError(t, err)

	// exchange the public_token for an access_token
	exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(sandboxPublicTokenResp.GetPublicToken()),
	).Execute()

	require.NotNil(t, exchangePublicTokenResp)
	require.NoError(t, err)

	return exchangePublicTokenResp.GetAccessToken()
}
