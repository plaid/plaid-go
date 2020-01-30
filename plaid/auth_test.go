package plaid

import (
	"context"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetAuth(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(context.Background(), sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(context.Background(), sandboxResp.PublicToken)
	authResp, err := testClient.GetAuth(context.Background(), tokenResp.AccessToken)

	// get auth for all accounts
	assert.Nil(t, err)
	assert.NotNil(t, authResp.Accounts)
	assert.NotNil(t, authResp.Numbers)

	// get auth for selected accounts
	options := GetAccountsOptions{
		AccountIDs: []string{authResp.Accounts[0].AccountID},
	}
	accountsResp, _ := testClient.GetAccountsWithOptions(context.Background(), tokenResp.AccessToken, options)
	assert.Equal(t, len(accountsResp.Accounts), 1)
}
