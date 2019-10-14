package plaid

import (
	"context"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetIdentity(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(context.Background(), sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(context.Background(), sandboxResp.PublicToken)
	identityResp, err := testClient.GetIdentity(context.Background(), tokenResp.AccessToken)

	assert.Nil(t, err)
	assert.NotNil(t, identityResp.Accounts)
	for _, acc := range identityResp.Accounts {
		assert.NotNil(t, acc.Owners)
		assert.True(t, len(acc.Owners) > 0)
	}
}
