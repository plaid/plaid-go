package plaid

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func generateDepositSwitchID(t *testing.T) (depositSwitchID string, err error) {
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
	return depositSwitchResp.DepositSwitchID, err
}

func TestCreateDepositSwitch(t *testing.T) {
	_, err := generateDepositSwitchID(t)
	assert.Nil(t, err)
}

func TestGetDepositSwitch(t *testing.T) {
	id, err := generateDepositSwitchID(t)
	assert.Nil(t, err)
	depositSwitch, err := testClient.GetDepositSwitch(id)
	assert.Nil(t, err)
	// Check that all of the fields we expect to be non-empty are non-empty.
	assert.True(t, depositSwitch.DepositSwitchID != "")
	assert.True(t, depositSwitch.TargetItemID != "")
	assert.True(t, depositSwitch.TargetAccountID != "")
	assert.True(t, depositSwitch.CreatedDate != "")
	assert.True(t, depositSwitch.State != "")
}
