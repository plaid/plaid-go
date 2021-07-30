package tests

import (
	"context"
	"testing"

	plaid "github.com/plaid/plaid-go"
	"github.com/stretchr/testify/require"
)

func TestGetAccounts(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		FIRST_PLATYPUS_BANK,
		[]plaid.Products{plaid.PRODUCTS_TRANSACTIONS},
	)

	// Get all accounts
	accountsGetResp, _, err := testClient.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*plaid.NewAccountsGetRequest(accessToken),
	).Execute()

	item, _ := accountsGetResp.GetItemOk()
	require.Nil(t, err)
	require.NotNil(t, item)
	require.True(t, len(accountsGetResp.GetAccounts()) > 1)

	// Get single account
	accountsGetRequest := plaid.NewAccountsGetRequest(accessToken)
	accountsGetRequest.SetOptions(plaid.AccountsGetRequestOptions{
		AccountIds: &[]string{accountsGetResp.GetAccounts()[0].GetAccountId()},
	})

	accountsGetResp, _, err = testClient.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*accountsGetRequest,
	).Execute()

	item, _ = accountsGetResp.GetItemOk()
	require.Nil(t, err)
	require.NotNil(t, item)
	require.Equal(t, len(accountsGetResp.GetAccounts()), 1)
}
