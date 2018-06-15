package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetIdentity(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	identityResp, err := testClient.GetIdentity(tokenResp.AccessToken)

	assert.Nil(t, err)
	assert.NotNil(t, identityResp.Identity)
}
