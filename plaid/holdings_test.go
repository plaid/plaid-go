package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetHoldings(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, []string{"investments"})
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	holdingsResp, err := testClient.GetHoldings(tokenResp.AccessToken)

	assert.Nil(t, err)
	assert.NotNil(t, holdingsResp.Accounts)
	assert.NotNil(t, holdingsResp.Securities)
	assert.NotNil(t, holdingsResp.Holdings)
	assert.NotNil(t, holdingsResp.Item)
}
