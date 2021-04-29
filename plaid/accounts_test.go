package plaid

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func TestGetAccounts(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)

	// get all accounts
	options := GetAccountsOptions{
		AccountIDs: []string{},
	}
	accountsResp, err := testClient.GetAccountsWithOptions(tokenResp.AccessToken, options)
	assert.Nil(t, err)
	assert.NotNil(t, accountsResp.Accounts)
	assert.True(t, len(accountsResp.Accounts) > 1)
	assert.NotNil(t, accountsResp.Item)

	// get selected accounts
	options = GetAccountsOptions{
		AccountIDs: []string{accountsResp.Accounts[0].AccountID},
	}
	accountsResp, err = testClient.GetAccountsWithOptions(tokenResp.AccessToken, options)
	assert.Nil(t, err)
	assert.Equal(t, len(accountsResp.Accounts), 1)
	assert.NotNil(t, accountsResp.Item)
}

func TestGetBalances(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)

	// get all balances
	balanceResp, err := testClient.GetBalances(tokenResp.AccessToken)
	assert.Nil(t, err)
	assert.NotNil(t, balanceResp.Accounts)

	// get selected accounts
	options := GetBalancesOptions{
		AccountIDs: []string{balanceResp.Accounts[0].AccountID},
	}
	balanceResp, err = testClient.GetBalancesWithOptions(tokenResp.AccessToken, options)
	assert.Nil(t, err)
	assert.Equal(t, len(balanceResp.Accounts), 1)
}

func TestGetBalancesWithMinDatetime(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(balanceDatetimeInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)

	format := "2006-01-02T15:04:05-07:00"

	// get all balances and check for datetime
	options := GetBalancesOptions{
		MinLastUpdatedDatetime: time.Now().AddDate(0, 0, -2).UTC().Format(format),
	}
	balanceResp, err := testClient.GetBalancesWithOptions(tokenResp.AccessToken, options)
	assert.Nil(t, err)
	assert.NotNil(t, balanceResp.Accounts)
	assert.NotEqual(t, balanceResp.Accounts[0].Balances.LastUpdatedDatetime, "")

	// assert a min last updated datetime that will fail
	options.MinLastUpdatedDatetime = time.Now().UTC().Format(format)
	_, err = testClient.GetBalancesWithOptions(tokenResp.AccessToken, options)
	assert.NotNil(t, err)
	assert.Equal(t, err.(Error).ErrorCode, "LAST_UPDATED_DATETIME_OUT_OF_RANGE")
}
