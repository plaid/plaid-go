package plaid

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func TestGetTransactions(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	transactionsResp, err := testClient.GetTransactions(tokenResp.AccessToken, Date(2016, time.January, 1), Date(2017, time.January, 1))

	if plaidErr, ok := err.(Error); ok {
		for ok && plaidErr.ErrorCode == "PRODUCT_NOT_READY" {
			time.Sleep(5 * time.Second)
			transactionsResp, err = testClient.GetTransactions(tokenResp.AccessToken, Date(2016, time.January, 01), Date(2017, time.January, 1))
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
		StartDate:  Date(2016, time.January, 1),
		EndDate:    Date(2017, time.January, 1),
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
