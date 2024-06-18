package tests

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/plaid/plaid-go/v26/plaid"
	assert "github.com/stretchr/testify/require"
)

type processorTokenPrepValues struct {
	accessToken string
	accountID   string
}

func processorTokenTestPrep(t *testing.T, ctx context.Context, testClient *plaid.APIClient) processorTokenPrepValues {
	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)

	accountsResp, _, err := testClient.PlaidApi.AccountsGet(ctx).AccountsGetRequest(plaid.AccountsGetRequest{
		AccessToken: accessToken,
	}).Execute()
	assert.NoError(t, err)
	accountID := accountsResp.Accounts[0].AccountId
	return processorTokenPrepValues{
		accessToken: accessToken,
		accountID:   accountID,
	}
}

func TestProcessorCreateApexToken(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	prepValues := processorTokenTestPrep(t, ctx, testClient)
	apexTokenResp, _, err := testClient.PlaidApi.ProcessorApexProcessorTokenCreate(ctx).ProcessorApexProcessorTokenCreateRequest(*plaid.NewProcessorApexProcessorTokenCreateRequest(prepValues.accessToken, prepValues.accountID)).Execute()
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(apexTokenResp.ProcessorToken, fmt.Sprintf("processor-%s", PlaidEnv())))
}

func TestProcessorCreateDwollaToken(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	prepValues := processorTokenTestPrep(t, ctx, testClient)
	dwollaTokenResp, _, err := testClient.PlaidApi.ProcessorTokenCreate(ctx).ProcessorTokenCreateRequest(*plaid.NewProcessorTokenCreateRequest(prepValues.accessToken, prepValues.accountID, "dwolla")).Execute()
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(dwollaTokenResp.ProcessorToken, fmt.Sprintf("processor-%s", PlaidEnv())))
}
