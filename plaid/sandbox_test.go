package plaid

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCreateSandboxPublicToken(t *testing.T) {
	sandboxResp, err := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(sandboxResp.PublicToken, "public-sandbox"))
}

func TestResetSandboxItem(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	resetResp, err := testClient.ResetSandboxItem(tokenResp.AccessToken)

	assert.Nil(t, err)
	assert.True(t, resetResp.ResetLogin)
}
