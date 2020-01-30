package plaid

import (
	"context"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCreateSandboxPublicToken(t *testing.T) {
	sandboxResp, err := testClient.CreateSandboxPublicToken(context.Background(), sandboxInstitution, testProducts)

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(sandboxResp.PublicToken, "public-sandbox"))
}

func TestResetSandboxItem(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(context.Background(), sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(context.Background(), sandboxResp.PublicToken)
	resetResp, err := testClient.ResetSandboxItem(context.Background(), tokenResp.AccessToken)

	assert.Nil(t, err)
	assert.True(t, resetResp.ResetLogin)
}
