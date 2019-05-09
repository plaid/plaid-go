package plaid

import (
	"testing"

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
	assert.Equal(t, len(accountsResp.Accounts), 7)

	// get selected accounts
	options = GetAccountsOptions{
		AccountIDs: []string{accountsResp.Accounts[0].AccountID},
	}
	accountsResp, err = testClient.GetAccountsWithOptions(tokenResp.AccessToken, options)
	assert.Nil(t, err)
	assert.Equal(t, len(accountsResp.Accounts), 1)
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
