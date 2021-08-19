package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/plaid"
	"github.com/stretchr/testify/assert"
)

func TestWebhookVerificationKeyGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	webhookResp, _, err := testClient.PlaidApi.WebhookVerificationKeyGet(ctx).WebhookVerificationKeyGetRequest(*plaid.NewWebhookVerificationKeyGetRequest("6c5516e1-92dc-479e-a8ff-5a51992e0001")).Execute()

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
