package plaid

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func TestGetTransactions(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	transactionsResp, err := testClient.GetTransactions(tokenResp.AccessToken, "2016-01-01", "2017-01-01")

	if plaidErr, ok := err.(Error); ok {
		for ok && plaidErr.ErrorCode == "PRODUCT_NOT_READY" {
			time.Sleep(5 * time.Second)
			transactionsResp, err = testClient.GetTransactions(tokenResp.AccessToken, "2016-01-01", "2017-01-01")
			plaidErr, ok = err.(Error)
		}
	}

	assert.Nil(t, err)
	assert.NotNil(t, transactionsResp.Accounts)
	assert.NotNil(t, transactionsResp.Transactions)
}

func TestGetTransactionsWithOptions(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)

	options := GetTransactionsOptions{
		StartDate:  "2016-01-01",
		EndDate:    "2017-01-01",
		AccountIDs: []string{},
		Count:      2,
		Offset:     1,
	}
	transactionsResp, err := testClient.GetTransactionsWithOptions(tokenResp.AccessToken, options)

	if plaidErr, ok := err.(Error); ok {
		for ok && plaidErr.ErrorCode == "PRODUCT_NOT_READY" {
			time.Sleep(5 * time.Second)
			transactionsResp, err = testClient.GetTransactionsWithOptions(tokenResp.AccessToken, options)
			plaidErr, ok = err.(Error)
		}
	}

	assert.Nil(t, err)
	assert.NotNil(t, transactionsResp.Transactions)
}
