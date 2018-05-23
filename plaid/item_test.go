package plaid

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func TestGetItem(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	itemResp, err := testClient.GetItem(tokenResp.AccessToken)

	assert.Nil(t, err)
	assert.NotNil(t, itemResp.Item)
}

func TestRemoveItem(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	itemResp, err := testClient.RemoveItem(tokenResp.AccessToken)

	assert.Nil(t, err)
	assert.True(t, itemResp.Removed)
}

func TestUpdateItemWebhook(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	itemResp, err := testClient.UpdateItemWebhook(tokenResp.AccessToken, "https://plaid.com/webhook-test")

	assert.Nil(t, err)
	assert.NotNil(t, itemResp.Item)
	assert.Equal(t, itemResp.Item.Webhook, "https://plaid.com/webhook-test")
}

func TestInvalidateAccessToken(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	newTokenResp, err := testClient.InvalidateAccessToken(tokenResp.AccessToken)

	assert.Nil(t, err)
	assert.NotNil(t, newTokenResp.NewAccessToken)
}

func TestUpdateAccessTokenVersion(t *testing.T) {
	invalidToken, _ := randomHex(80)
	newTokenResp, err := testClient.InvalidateAccessToken(invalidToken)
	assert.NotNil(t, err)
	assert.True(t, newTokenResp.NewAccessToken == "")
}

func TestCreatePublicToken(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	publicTokenResp, err := testClient.CreatePublicToken(tokenResp.AccessToken)

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(publicTokenResp.PublicToken, "public-sandbox"))
}

func TestExchangePublicToken(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, err := testClient.ExchangePublicToken(sandboxResp.PublicToken)

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(tokenResp.AccessToken, "access-sandbox"))
	assert.NotNil(t, tokenResp.ItemID)
}
