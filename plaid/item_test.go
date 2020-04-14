package plaid

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"testing"
	"time"

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
	assert.NotEmpty(t, itemResp.Status)
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

func TestCreateItemAddToken(t *testing.T) {
	fakeClientUserID, _ := randomHex(12)
	itemAddTokenResp, err := testClient.CreateItemAddToken(ItemAddTokenUserFields{
		ClientUserID: fakeClientUserID,
	})

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(itemAddTokenResp.AddToken, "item-add-sandbox"))
	assert.NotZero(t, itemAddTokenResp.Expiration)
}

func TestCreateItemAddTokenWithUserFields(t *testing.T) {
	fakeClientUserID, _ := randomHex(12)
	timeA := time.Date(2020, 5, 4, 12, 4, 2, 0, time.UTC)

	itemAddTokenResp, err := testClient.CreateItemAddToken(ItemAddTokenUserFields{
		ClientUserID:          fakeClientUserID,
		EmailAddress:          "hdcase@example.com",
		EmailAddressVerified:  true,
		PhoneNumber:           "+1 (415) 555-0333",
		PhoneNumberVerifiedAt: &timeA,
	})

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(itemAddTokenResp.AddToken, "item-add-sandbox"))
	assert.NotZero(t, itemAddTokenResp.Expiration)
}

func TestExchangePublicToken(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, err := testClient.ExchangePublicToken(sandboxResp.PublicToken)

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(tokenResp.AccessToken, "access-sandbox"))
	assert.NotNil(t, tokenResp.ItemID)
}

func TestImportItemWithoutOptions(t *testing.T) {
	accessTokenResp, err := testClient.ImportItem([]string{"identity", "auth"}, map[string]interface{}{
		"user_id":    "user_good",
		"auth_token": "pass_good",
	}, importItemRequestOptions{})
	assert.Nil(t, err)
	assert.NotNil(t, accessTokenResp)
	assert.True(t, strings.HasPrefix(accessTokenResp.AccessToken, "access-sandbox"))
}

func TestImportItemWithOptions(t *testing.T) {
	accessTokenResp, err := testClient.ImportItem([]string{"identity", "auth"}, map[string]interface{}{
		"user_id":    "user_good",
		"auth_token": "pass_good",
	}, importItemRequestOptions{
		Webhook: "https://plaid.com/webhook-test",
	})
	assert.Nil(t, err)
	assert.NotNil(t, accessTokenResp)
	assert.True(t, strings.HasPrefix(accessTokenResp.AccessToken, "access-sandbox"))
}

func Test_prepareUserFieldsForSend(t *testing.T) {
	// no panic
	prepareUserFieldsForSend(nil)

	timeA := time.Unix(50, 0)
	uf := ItemAddTokenUserFields{
		EmailAddressVerified:  true,
		PhoneNumberVerifiedAt: &timeA,
		PhoneNumberVerified:   true,
	}
	prepareUserFieldsForSend(&uf)
	assert.Equal(t, &timeA, uf.PhoneNumberVerifiedAt)
	assert.Equal(t, &verificationDateUnknown, uf.EmailAddressVerifiedAt)
}
