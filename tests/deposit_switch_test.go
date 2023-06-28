package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v14/plaid"
	"github.com/stretchr/testify/assert"
)

func createDepositSwitchID(t *testing.T, ctx context.Context, testClient *plaid.APIClient) (depositSwitchID string) {
	accessToken := importSandboxItem(
		t,
		ctx,
		testClient,
		"user_good",
		"pass_good",
		[]plaid.Products{plaid.PRODUCTS_IDENTITY, plaid.PRODUCTS_AUTH},
	)

	// Get all accounts
	accountsGetResp, _, err := testClient.PlaidApi.AccountsGet(ctx).AccountsGetRequest(
		*plaid.NewAccountsGetRequest(accessToken),
	).Execute()

	item, _ := accountsGetResp.GetItemOk()
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Greater(t, len(accountsGetResp.GetAccounts()), 1)

	// Create deposit switch
	depositSwitchResp, _, err := testClient.PlaidApi.DepositSwitchCreate(ctx).DepositSwitchCreateRequest(
		*plaid.NewDepositSwitchCreateRequest(
			accessToken,
			accountsGetResp.GetAccounts()[0].GetAccountId(),
		),
	).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, depositSwitchResp.GetDepositSwitchId())

	return depositSwitchResp.GetDepositSwitchId()
}

func TestDepositSwitchCreate(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	_ = createDepositSwitchID(t, ctx, testClient)
}

func TestDepositSwitchGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	depositSwitchID := createDepositSwitchID(t, ctx, testClient)

	// Get deposit switch
	depositSwitchGetResp, _, err := testClient.PlaidApi.DepositSwitchGet(ctx).DepositSwitchGetRequest(
		*plaid.NewDepositSwitchGetRequest(depositSwitchID),
	).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, depositSwitchGetResp.GetDepositSwitchId())
	assert.NotEmpty(t, depositSwitchGetResp.GetTargetItemId())
	assert.NotEmpty(t, depositSwitchGetResp.GetTargetAccountId())
	assert.NotEmpty(t, depositSwitchGetResp.GetDateCreated())
	assert.NotEmpty(t, depositSwitchGetResp.GetState())
}

func TestDepositSwitchTokenCreate(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	depositSwitchID := createDepositSwitchID(t, ctx, testClient)

	// Create deposit switch token
	depositSwitchTokenResp, _, err := testClient.PlaidApi.DepositSwitchTokenCreate(ctx).DepositSwitchTokenCreateRequest(
		*plaid.NewDepositSwitchTokenCreateRequest(depositSwitchID),
	).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, depositSwitchTokenResp.GetDepositSwitchToken())
	assert.NotEmpty(t, depositSwitchTokenResp.GetDepositSwitchTokenExpirationTime())
}
