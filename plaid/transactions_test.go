package plaid

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

const iso8601TimeFormat = "2006-01-02"

func TestGetTransactions(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
	endDate := time.Now().Format(iso8601TimeFormat)
	transactionsResp, err := testClient.GetTransactions(tokenResp.AccessToken, startDate, endDate)

	if plaidErr, ok := err.(Error); ok {
		for ok && plaidErr.ErrorCode == "PRODUCT_NOT_READY" {
			time.Sleep(5 * time.Second)
			transactionsResp, err = testClient.GetTransactions(tokenResp.AccessToken, startDate, endDate)
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

	startDate := time.Now().Add(-365 * 24 * time.Hour).Format(iso8601TimeFormat)
	endDate := time.Now().Format(iso8601TimeFormat)
	options := GetTransactionsOptions{
		StartDate:  startDate,
		EndDate:    endDate,
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
