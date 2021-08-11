package tests

import (
	"context"
	"testing"

	plaid "github.com/plaid/plaid-go"
	"github.com/stretchr/testify/assert"
)

func TestAssetReportCreate(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		FIRST_PLATYPUS_BANK,
		[]plaid.Products{plaid.PRODUCTS_ASSETS},
	)

	assetReportCreateResp, _, err := testClient.PlaidApi.AssetReportCreate(ctx).AssetReportCreateRequest(
		*plaid.NewAssetReportCreateRequest([]string{accessToken}, 2),
	).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, assetReportCreateResp.GetAssetReportToken())
}
