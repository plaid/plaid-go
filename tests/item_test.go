package tests

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/plaid/plaid-go/v39/plaid"
	"github.com/stretchr/testify/assert"
)

func TestItemGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)
	request := plaid.NewItemGetRequest(accessToken)
	resp, _, err := testClient.PlaidApi.ItemGet(ctx).ItemGetRequest(*request).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.GetItem().ItemId)
}

func TestItemRemove(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)
	request := plaid.NewItemRemoveRequest(accessToken)
	_, _, err := testClient.PlaidApi.ItemRemove(ctx).ItemRemoveRequest(*request).Execute()
	assert.NoError(t, err)
}

func TestItemWebhookUpdate(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)
	request := plaid.NewItemWebhookUpdateRequest(accessToken)
	request.SetWebhook("https://plaid.com/webhook-test")
	resp, _, err := testClient.PlaidApi.ItemWebhookUpdate(ctx).ItemWebhookUpdateRequest(*request).Execute()

	assert.NoError(t, err)
	assert.NotEqual(t, plaid.Item{}, resp.GetItem())
	assert.Equal(t, resp.Item.GetWebhook(), "https://plaid.com/webhook-test")
}

func TestItemAccessTokenInvalidate(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)
	request := plaid.NewItemAccessTokenInvalidateRequest(accessToken)
	resp, _, err := testClient.PlaidApi.ItemAccessTokenInvalidate(ctx).ItemAccessTokenInvalidateRequest(*request).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, resp.GetNewAccessToken())
}

func TestItemPublicTokenCreate(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, testProducts)
	request := plaid.NewItemPublicTokenCreateRequest(accessToken)
	resp, _, err := testClient.PlaidApi.ItemCreatePublicToken(ctx).ItemPublicTokenCreateRequest(*request).Execute()

	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(resp.GetPublicToken(), fmt.Sprintf("public-%s", PlaidEnv())))
}

func TestItemPublicTokenExchange(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	sandboxReq := plaid.NewSandboxPublicTokenCreateRequest(FIRST_PLATYPUS_BANK, testProducts)
	sandboxResp, _, err := testClient.PlaidApi.SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(*sandboxReq).Execute()
	assert.NoError(t, err)

	tokenReq := plaid.NewItemPublicTokenExchangeRequest(sandboxResp.PublicToken)
	tokenResp, _, err := testClient.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(*tokenReq).Execute()
	assert.NoError(t, err)

	assert.True(t, strings.HasPrefix(tokenResp.GetAccessToken(), fmt.Sprintf("access-%s", PlaidEnv())))
	assert.NotEmpty(t, tokenResp.GetItemId())
}

func TestItemImportWithoutOptions(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	request := plaid.NewItemImportRequest(
		[]plaid.Products{plaid.PRODUCTS_IDENTITY, plaid.PRODUCTS_AUTH},
		*plaid.NewItemImportRequestUserAuth("user_good", "pass_good"))
	resp, _, err := testClient.PlaidApi.ItemImport(ctx).ItemImportRequest(*request).Execute()
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(resp.GetAccessToken(), fmt.Sprintf("access-%s", PlaidEnv())))
}

func TestItemImportWithOptions(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	request := plaid.NewItemImportRequest(
		[]plaid.Products{plaid.PRODUCTS_IDENTITY, plaid.PRODUCTS_AUTH},
		*plaid.NewItemImportRequestUserAuth("user_good", "pass_good"))
	webhook := "https://plaid.com/webhook-test"
	request.SetOptions(plaid.ItemImportRequestOptions{
		Webhook: &webhook,
	})
	resp, _, err := testClient.PlaidApi.ItemImport(ctx).ItemImportRequest(*request).Execute()
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(resp.GetAccessToken(), fmt.Sprintf("access-%s", PlaidEnv())))
}
