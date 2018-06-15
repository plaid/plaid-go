package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetAuth(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	authResp, err := testClient.GetAuth(tokenResp.AccessToken)

	// get auth for all accounts
	assert.Nil(t, err)
	assert.NotNil(t, authResp.Accounts)
	assert.NotNil(t, authResp.Numbers)

	// get auth for selected accounts
	options := GetAccountsOptions{
		AccountIDs: []string{authResp.Accounts[0].AccountID},
	}
	accountsResp, _ := testClient.GetAccountsWithOptions(tokenResp.AccessToken, options)
	assert.Equal(t, len(accountsResp.Accounts), 1)
}
