package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetWebhookVerificationKey(t *testing.T) {
	webhookResp, err := testClient.GetWebhookVerificationKey("6c5516e1-92dc-479e-a8ff-5a51992e0001")

	assert.Nil(t, err)
	assert.NotNil(t, webhookResp.Key)
	assert.NotNil(t, webhookResp.Key.Kid)
}
