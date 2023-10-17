package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v17/plaid"
	"github.com/stretchr/testify/assert"
)

func TestHoldingsGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		FIRST_PLATYPUS_BANK,
		[]plaid.Products{plaid.PRODUCTS_INVESTMENTS},
	)

	// Get holdings for all accounts
	holdingsGetResp, _, err := testClient.PlaidApi.InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(
		*plaid.NewInvestmentsHoldingsGetRequest(accessToken),
	).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, holdingsGetResp.GetAccounts())
	assert.NotEmpty(t, holdingsGetResp.GetHoldings())
	assert.NotEmpty(t, holdingsGetResp.GetSecurities())

	// Get holdings for single account
	holdingsGetReq := plaid.NewInvestmentsHoldingsGetRequest(accessToken)
	holdingsGetReqOptions := plaid.NewInvestmentHoldingsGetRequestOptions()
	holdingsGetReqOptions.SetAccountIds([]string{holdingsGetResp.GetAccounts()[0].GetAccountId()})
	holdingsGetReq.SetOptions(*holdingsGetReqOptions)

	holdingsGetResp, _, err = testClient.PlaidApi.InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(
		*holdingsGetReq,
	).Execute()

	assert.NoError(t, err)
	assert.Len(t, holdingsGetResp.GetAccounts(), 1)
}
