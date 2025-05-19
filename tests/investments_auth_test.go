package tests

import (
	"context"
	"github.com/plaid/plaid-go/v35/plaid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInvestmentsAuthGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()
	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, []plaid.Products{plaid.PRODUCTS_INVESTMENTS_AUTH})
	request := plaid.NewInvestmentsAuthGetRequest(accessToken)
	response, _, err := testClient.PlaidApi.InvestmentsAuthGet(ctx).InvestmentsAuthGetRequest(*request).Execute()
	require.NoError(t, err)
	assert.NotEmpty(t, response.GetAccounts())
	assert.NotEmpty(t, response.GetSecurities())
	assert.NotEmpty(t, response.GetHoldings())
	assert.NotEmpty(t, response.GetOwners())
	assert.NotEmpty(t, response.GetNumbers())
}
