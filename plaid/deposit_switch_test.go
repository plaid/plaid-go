package plaid

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCreateDepositSwitch(t *testing.T) {
	accessTokenResp, err := testClient.ImportItem([]string{"identity", "auth"}, map[string]interface{}{
		"user_id":    "user_good",
		"auth_token": "pass_good",
	}, importItemRequestOptions{
		Webhook: "https://plaid.com/webhook-test",
	})
	assert.Nil(t, err)
	assert.NotNil(t, accessTokenResp)
	assert.True(t, strings.HasPrefix(accessTokenResp.AccessToken, "access-sandbox"))

	getAccountsResp, err := testClient.GetAccounts(accessTokenResp.AccessToken)
	assert.Nil(t, err)
	targetAccountID := getAccountsResp.Accounts[0].AccountID
	depositSwitchResp, err := testClient.CreateDepositSwitch(targetAccountID, accessTokenResp.AccessToken)
	assert.Nil(t, err)
	assert.True(t, depositSwitchResp.DepositSwitchID != "")
}
