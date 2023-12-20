package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v20/plaid"
	"github.com/stretchr/testify/assert"
)

func TestBalancesGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		FIRST_PLATYPUS_BANK,
		[]plaid.Products{plaid.PRODUCTS_TRANSACTIONS},
	)

	// Get all balances
	balancesGetResp, _, err := testClient.PlaidApi.AccountsBalanceGet(ctx).AccountsBalanceGetRequest(
		*plaid.NewAccountsBalanceGetRequest(accessToken),
	).Execute()

	item, _ := balancesGetResp.GetItemOk()
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Greater(t, len(balancesGetResp.GetAccounts()), 1)

	// Get single account
	balancesGetRequest := plaid.NewAccountsBalanceGetRequest(accessToken)
	balancesGetRequest.SetOptions(plaid.AccountsBalanceGetRequestOptions{
		AccountIds: &[]string{balancesGetResp.GetAccounts()[0].GetAccountId()},
	})

	balancesGetResp, _, err = testClient.PlaidApi.AccountsBalanceGet(ctx).AccountsBalanceGetRequest(
		*balancesGetRequest,
	).Execute()

	item, _ = balancesGetResp.GetItemOk()
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, len(balancesGetResp.GetAccounts()), 1)
}
