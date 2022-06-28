package tests

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/plaid/plaid-go/v8/plaid"
	"github.com/stretchr/testify/assert"
)

func createBankTransfer(
	t *testing.T,
	ctx context.Context,
	testClient *plaid.APIClient,
) plaid.BankTransferCreateResponse {
	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		FIRST_PLATYPUS_BANK,
		[]plaid.Products{plaid.PRODUCTS_AUTH},
	)

	accountsRequest := plaid.NewAccountsGetRequest(accessToken)
	accountsResponse, _, err := testClient.PlaidApi.AccountsGet(ctx).AccountsGetRequest(*accountsRequest).Execute()
	assert.NoError(t, err)
	var accountID string
	for _, account := range accountsResponse.Accounts {
		if account.Type == plaid.ACCOUNTTYPE_DEPOSITORY {
			accountID = account.AccountId
			break
		}
	}
	assert.NotEmpty(t, accountID)

	btCreateReq := plaid.NewBankTransferCreateRequest(
		fmt.Sprintf("%f", rand.Float32()),
		accessToken,
		accountID,
		plaid.BANKTRANSFERTYPE_CREDIT,
		plaid.BANKTRANSFERNETWORK_ACH,
		"1.00",
		"USD",
		"test",
		*plaid.NewBankTransferUser("Firstname Lastname"),
	)
	btCreateReq.SetAchClass(plaid.ACHCLASS_PPD)
	btCreateResp, _, err := testClient.PlaidApi.BankTransferCreate(ctx).BankTransferCreateRequest(*btCreateReq).Execute()
	if err != nil {
		plaidErr, _ := plaid.ToPlaidError(err)
		t.Log(plaidErr)
	}
	assert.NoError(t, err)
	assert.NotEmpty(t, btCreateResp.BankTransfer.Id)
	return btCreateResp
}

func TestBankTransferCreate(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	createBankTransfer(t, ctx, testClient)
}

func TestBankTransferCancel(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	btResp := createBankTransfer(t, ctx, testClient)

	cancelReq := plaid.NewBankTransferCancelRequest(btResp.BankTransfer.Id)
	_, _, err := testClient.PlaidApi.BankTransferCancel(ctx).BankTransferCancelRequest(*cancelReq).Execute()
	assert.NoError(t, err)
}

func TestBankTransferEventList(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	createBankTransfer(t, ctx, testClient)

	listReq := plaid.NewBankTransferEventListRequest()
	listReq.SetCount(1)
	listResp, _, err := testClient.PlaidApi.BankTransferEventList(ctx).BankTransferEventListRequest(*listReq).Execute()
	assert.NoError(t, err)
	assert.Len(t, listResp.BankTransferEvents, 1)
}

func TestBankTransferEventSync(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	createBankTransfer(t, ctx, testClient)

	syncReq := plaid.NewBankTransferEventSyncRequest(0)
	syncResp, _, err := testClient.PlaidApi.BankTransferEventSync(ctx).BankTransferEventSyncRequest(*syncReq).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, syncResp.BankTransferEvents)
}

func TestBankTransferGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	btResp := createBankTransfer(t, ctx, testClient)

	getReq := plaid.NewBankTransferGetRequest(btResp.BankTransfer.Id)
	getResp, _, err := testClient.PlaidApi.BankTransferGet(ctx).BankTransferGetRequest(*getReq).Execute()
	assert.NoError(t, err)
	assert.Equal(t, btResp.BankTransfer.Id, getResp.BankTransfer.Id)
}

func TestBankTransferList(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	createBankTransfer(t, ctx, testClient)

	listReq := plaid.NewBankTransferListRequest()
	listReq.SetCount(1)
	listResp, _, err := testClient.PlaidApi.BankTransferList(ctx).BankTransferListRequest(*listReq).Execute()
	assert.NoError(t, err)
	assert.Len(t, listResp.BankTransfers, 1)
}

func TestBankTransferBalanceGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	balReq := plaid.NewBankTransferBalanceGetRequest()
	balResp, _, err := testClient.PlaidApi.BankTransferBalanceGet(ctx).BankTransferBalanceGetRequest(*balReq).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, balResp.Balance.Available)
}

func TestBankTransferMigrateAccount(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	migrateReq := plaid.NewBankTransferMigrateAccountRequest(
		"100000000",
		"121122676",
		"checking",
	)
	migrateResp, _, err := testClient.PlaidApi.BankTransferMigrateAccount(ctx).BankTransferMigrateAccountRequest(*migrateReq).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, migrateResp.AccessToken)
}
