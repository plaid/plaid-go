package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v17/plaid"
	"github.com/stretchr/testify/assert"
)

func TestAccountsGet(t *testing.T) {
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
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Greater(t, len(accountsGetResp.GetAccounts()), 1)

	// Get single account
	accountsGetRequest := plaid.NewAccountsGetRequest(accessToken)
	accountsGetRequest.SetOptions(plaid.AccountsGetRequestOptions{
		AccountIds: &[]string{accountsGetResp.GetAccounts()[0].GetAccountId()},
	})

	accountsGetResp, _, err = testClient.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*accountsGetRequest,
	).Execute()

	item, _ = accountsGetResp.GetItemOk()
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, len(accountsGetResp.GetAccounts()), 1)
}

func TestAccountsGetWithPersistentAccountId(t *testing.T) {
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

	persistentAccountId := accountsGetResp.GetAccounts()[0].PersistentAccountId

	item, _ := accountsGetResp.GetItemOk()
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Greater(t, len(accountsGetResp.GetAccounts()), 1)

	// Get single account
	accountsGetRequest := plaid.NewAccountsGetRequest(accessToken)
	accountsGetRequest.SetOptions(plaid.AccountsGetRequestOptions{
		AccountIds: &[]string{accountsGetResp.GetAccounts()[0].GetAccountId()},
	})

	accountsGetResp, _, err = testClient.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*accountsGetRequest,
	).Execute()

	item, _ = accountsGetResp.GetItemOk()
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, len(accountsGetResp.GetAccounts()), 1)
	assert.Equal(
		t,
		accountsGetResp.GetAccounts()[0].PersistentAccountId,
		persistentAccountId,
	)
}
