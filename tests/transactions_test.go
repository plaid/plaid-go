package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/plaid/plaid-go/v32/plaid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionsGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)

	startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
	endDate := time.Now().Format(iso8601TimeFormat)
	request := plaid.NewTransactionsGetRequest(
		accessToken,
		startDate,
		endDate,
	)

	transactionsResp, err := pollForTransactionsGet(t, ctx, testClient, request)

	require.NoError(t, err)
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

	transactionsResp, err := pollForTransactionsGet(t, ctx, testClient, request)

	require.NoError(t, err)
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

func pollForTransactionsGet(t *testing.T, ctx context.Context, testClient *plaid.APIClient, request *plaid.TransactionsGetRequest) (*plaid.TransactionsGetResponse, error) {

	for i := 0; i < 10; i++ {
		response, _, err := testClient.PlaidApi.TransactionsGet(ctx).TransactionsGetRequest(*request).Execute()

		if err == nil {
			return &response, nil
		}

		plaidErr, conversionErr := plaid.ToPlaidError(err)
		assert.NoError(t, conversionErr)
		if plaidErr.ErrorCode == "PRODUCT_NOT_READY" {
			time.Sleep(2 * time.Second)
			continue
		}

		return &response, err
	}

	return nil, errors.New("failed to get transactions")
}
