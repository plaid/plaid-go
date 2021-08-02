package tests

import (
	"context"
	"testing"

	plaid "github.com/plaid/plaid-go"
	assert "github.com/stretchr/testify/require"
)

func TestSandboxItemResetLogin(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)

	resetResp, _, err := testClient.PlaidApi.SandboxItemResetLogin(ctx).SandboxItemResetLoginRequest(*plaid.NewSandboxItemResetLoginRequest(accessToken)).Execute()
	assert.NoError(t, err)
	assert.True(t, resetResp.ResetLogin)
}
