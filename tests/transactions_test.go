package tests

import (
	"context"
	"testing"
	"time"

	plaid "github.com/plaid/plaid-go"
	assert "github.com/stretchr/testify/require"
)

func TestTransactionsGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)

	startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
	endDate := time.Now().Format(iso8601TimeFormat)
	transactionsResp, _, err := testClient.PlaidApi.TransactionsGet(ctx).TransactionsGetRequest(*plaid.NewTransactionsGetRequest(
		accessToken,
		startDate,
		endDate,
	)).Execute()

	assert.NoError(t, err)
	assert.NotNil(t, transactionsResp.Accounts)
	assert.NotNil(t, transactionsResp.Transactions)
}

func TestTransactionsGetWithOptions(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)

	startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
	endDate := time.Now().Format(iso8601TimeFormat)
	options := plaid.TransactionsGetRequestOptions{
		Count:  plaid.PtrInt32(2),
		Offset: plaid.PtrInt32(1),
	}
	request := plaid.NewTransactionsGetRequest(
		accessToken,
		startDate,
		endDate,
	)
	request.SetOptions(options)
	transactionsResp, _, err := testClient.PlaidApi.TransactionsGet(ctx).TransactionsGetRequest(*request).Execute()

	assert.NoError(t, err)
	assert.NotNil(t, transactionsResp.Accounts)
	assert.NotNil(t, transactionsResp.Transactions)

}

func TestTransactionsRefresh(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)
	_, _, err := testClient.PlaidApi.TransactionsRefresh(ctx).TransactionsRefreshRequest(*plaid.NewTransactionsRefreshRequest(accessToken)).Execute()

	assert.NoError(t, err)
}
