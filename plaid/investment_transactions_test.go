package plaid

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func TestGetInvestmentTransactions(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, []string{"investments"})
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	startDateString := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
	endDateString := time.Now().Format(iso8601TimeFormat)
	investmentTransactionsResp, err := testClient.GetInvestmentTransactions(tokenResp.AccessToken, startDateString, endDateString)

	if plaidErr, ok := err.(Error); ok {
		for ok && plaidErr.ErrorCode == "PRODUCT_NOT_READY" {
			time.Sleep(5 * time.Second)
			investmentTransactionsResp, err = testClient.GetInvestmentTransactions(tokenResp.AccessToken, startDateString, endDateString)
			plaidErr, ok = err.(Error)
		}
	}

	assert.Nil(t, err)
	assert.NotNil(t, investmentTransactionsResp.Accounts)
	assert.NotNil(t, investmentTransactionsResp.InvestmentTransactions)
	assert.NotNil(t, investmentTransactionsResp.Securities)
	assert.NotNil(t, investmentTransactionsResp.Item)
}
