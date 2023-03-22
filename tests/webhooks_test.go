package tests

import (
	"context"
	"os"
	"testing"

	"github.com/plaid/plaid-go/v11/plaid"
	"github.com/stretchr/testify/assert"
)

// WebhookVerificationKeyId provides a valid key ID for testing webhook verification. It defaults
// to the key corresponding to the sandbox environment.
func WebhookVerificationKeyId() string {
	if keyId, ok := os.LookupEnv("WEBHOOK_VERIFICATION_KEY_ID"); ok {
		return keyId
	} else {
		return "6c5516e1-92dc-479e-a8ff-5a51992e0001"
	}
}

func TestWebhookVerificationKeyGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	webhookResp, _, err := testClient.PlaidApi.WebhookVerificationKeyGet(ctx).WebhookVerificationKeyGetRequest(*plaid.NewWebhookVerificationKeyGetRequest(WebhookVerificationKeyId())).Execute()

	assert.Nil(t, err)
	assert.NotEmpty(t, webhookResp.Key.Alg)
	assert.NotZero(t, webhookResp.Key.CreatedAt)
	assert.NotEmpty(t, webhookResp.Key.Crv)
	assert.NotEmpty(t, webhookResp.Key.Kid)
	assert.NotEmpty(t, webhookResp.Key.Kty)
	assert.NotEmpty(t, webhookResp.Key.Use)
	assert.NotEmpty(t, webhookResp.Key.X)
	assert.NotEmpty(t, webhookResp.Key.Y)
}
