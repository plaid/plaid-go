package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetLiabilities(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, []string{"liabilities"})
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	liabilitiesResp, err := testClient.GetLiabilities(tokenResp.AccessToken)

	assert.Nil(t, err)
	assert.NotNil(t, liabilitiesResp.Accounts)
	assert.NotNil(t, liabilitiesResp.Item)
	assert.NotNil(t, liabilitiesResp.Liabilities)
	assert.Len(t, liabilitiesResp.Liabilities.Student, 1)

	accountID := liabilitiesResp.Accounts[7].AccountID
	liabilitiesResp, err = testClient.GetLiabilitiesWithOptions(tokenResp.AccessToken, GetLiabilitiesOptions{
		AccountIDs: []string{accountID},
	})
	assert.Nil(t, err)
	assert.Len(t, liabilitiesResp.Accounts, 1)
	assert.Len(t, liabilitiesResp.Liabilities.Student, 1)
}
