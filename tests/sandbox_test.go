package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v11/plaid"
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

func TestSandboxIncomeVerificationItem(t *testing.T) {
	// TODO (czhou): Unskip when test is fixed.
	t.Skip()
	testClient := NewTestClient()
	ctx := context.Background()

	daysRequested := int32(180)
	sandboxPublicTokenResp, _, err := testClient.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(
		plaid.SandboxPublicTokenCreateRequest{
			InstitutionId:   FIRST_PLATYPUS_BANK,
			InitialProducts: []plaid.Products{plaid.PRODUCTS_INCOME_VERIFICATION},
			Options: &plaid.SandboxPublicTokenCreateRequestOptions{
				IncomeVerification: &plaid.SandboxPublicTokenCreateRequestOptionsIncomeVerification{
					BankIncome: &plaid.SandboxPublicTokenCreateRequestIncomeVerificationBankIncome{
						DaysRequested: &daysRequested,
					},
					IncomeSourceTypes: &[]plaid.IncomeVerificationSourceType{
						"bank",
					},
				},
			},
		},
	).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, sandboxPublicTokenResp)
	assert.NotEmpty(t, sandboxPublicTokenResp.GetPublicToken())
}
