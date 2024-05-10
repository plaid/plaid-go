package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v25/plaid"
	"github.com/stretchr/testify/assert"
)

func TestAuthGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		FIRST_PLATYPUS_BANK,
		[]plaid.Products{plaid.PRODUCTS_AUTH},
	)

	// Get auth for all accounts
	authGetResp, _, err := testClient.PlaidApi.AuthGet(ctx).AuthGetRequest(
		*plaid.NewAuthGetRequest(accessToken),
	).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, authGetResp.GetAccounts())
	assert.NotEqual(t, authGetResp.GetNumbers(), plaid.AuthGetNumbers{})
	assert.NotEqual(t, authGetResp.GetItem(), plaid.Item{})

	// Get auth for single account
	authGetRequest := plaid.NewAuthGetRequest(accessToken)
	authGetRequestOptions := plaid.NewAuthGetRequestOptions()
	authGetRequestOptions.SetAccountIds([]string{
		authGetResp.GetAccounts()[1].GetAccountId(),
	})
	authGetRequest.SetOptions(*authGetRequestOptions)
	authGetResp, _, err = testClient.PlaidApi.AuthGet(ctx).AuthGetRequest(
		*authGetRequest,
	).Execute()

	assert.NoError(t, err)
	assert.Equal(t, len(authGetResp.GetAccounts()), 1)
	assert.NotEqual(t, authGetResp.GetItem(), plaid.Item{})
}
