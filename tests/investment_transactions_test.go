package tests

import (
	"context"
	"testing"
	"time"

	"github.com/plaid/plaid-go/v24/plaid"
	"github.com/stretchr/testify/assert"
)

func TestInvestmentsTransactionsGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, []plaid.Products{plaid.PRODUCTS_INVESTMENTS})
	startDateString := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
	endDateString := time.Now().Format(iso8601TimeFormat)
	request := plaid.NewInvestmentsTransactionsGetRequest(accessToken, startDateString, endDateString)
	response, _, err := testClient.PlaidApi.InvestmentsTransactionsGet(ctx).InvestmentsTransactionsGetRequest(*request).Execute()
	assert.NoError(t, err)

	assert.NotEmpty(t, response.GetAccounts())
	assert.NotEmpty(t, response.GetInvestmentTransactions())
	for _, transaction := range response.InvestmentTransactions {
		assert.NotEmpty(t, transaction.Subtype)
	}
	assert.NotEmpty(t, response.GetSecurities())
}
