package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetHoldings(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, []string{"investments"})
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	options := GetHoldingsOptions{
		AccountIDs: []string{},
	}
	holdingsResp, err := testClient.GetHoldingsWithOptions(tokenResp.AccessToken, options)

	assert.Nil(t, err)
	assert.NotNil(t, holdingsResp.Accounts)
	assert.NotNil(t, holdingsResp.Securities)
	assert.NotNil(t, holdingsResp.Holdings)
	assert.NotNil(t, holdingsResp.Item)

	// Get only selected accounts.
	options = GetHoldingsOptions{
		AccountIDs: []string{holdingsResp.Accounts[0].AccountID},
	}
	holdingsResp, err = testClient.GetHoldingsWithOptions(tokenResp.AccessToken, options)
	assert.Nil(t, err)
	assert.Equal(t, len(holdingsResp.Accounts), 1)
	assert.NotNil(t, holdingsResp.Securities)
	assert.NotNil(t, holdingsResp.Holdings)
	assert.NotNil(t, holdingsResp.Item)
}
