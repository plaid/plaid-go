package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/plaid/plaid-go/v16/plaid"
	"github.com/stretchr/testify/assert"
)

func TestAssetReportFullFlow(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()
	clientID := testClient.GetConfig().DefaultHeader["PLAID-CLIENT-ID"]

	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		FIRST_PLATYPUS_BANK,
		[]plaid.Products{plaid.PRODUCTS_ASSETS},
	)

	// Create an asset report
	createRequest := plaid.NewAssetReportCreateRequest(2)
	createRequest.SetAccessTokens([]string{accessToken})
	assetReportCreateResp, _, err := testClient.PlaidApi.AssetReportCreate(ctx).AssetReportCreateRequest(
		*createRequest,
	).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, assetReportCreateResp.GetAssetReportToken())

	assetReportToken := assetReportCreateResp.GetAssetReportToken()

	// Get the asset report
	assetReportGetResponse, err := pollForAssetReport(t, ctx, testClient, assetReportToken)
	assert.NoError(t, err)
	report := assetReportGetResponse.Report
	assert.NotEqual(t, plaid.AssetReport{}, report)

	// Get it as a PDF
	pdfRequest := plaid.NewAssetReportPDFGetRequest(assetReportToken)
	pdfFile, _, err := testClient.PlaidApi.AssetReportPdfGet(ctx).AssetReportPDFGetRequest(*pdfRequest).Execute()
	assert.NoError(t, err)
	pdfFileStat, _ := pdfFile.Stat()
	assert.NotZero(t, pdfFileStat.Size())

	// Get it filtered
	accountIdsToFilter := []string{report.Items[0].Accounts[0].AccountId}
	filterRequest := plaid.NewAssetReportFilterRequest(assetReportToken, accountIdsToFilter)
	filterResponse, _, err := testClient.PlaidApi.AssetReportFilter(ctx).AssetReportFilterRequest(*filterRequest).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, filterResponse.AssetReportToken)

	// Get a refreshed copy of the report
	refreshRequest := plaid.NewAssetReportRefreshRequest(assetReportToken)
	refreshRequest.SetDaysRequested(10)
	refreshResponse, _, err := testClient.PlaidApi.AssetReportRefresh(ctx).AssetReportRefreshRequest(*refreshRequest).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, refreshResponse.AssetReportToken)

	// Create an audit copy
	auditCopyCreateRequest := plaid.NewAssetReportAuditCopyCreateRequest(assetReportToken)
	auditCopyCreateRequest.SetAuditorId(clientID)
	auditCopyCreateResponse, _, err := testClient.PlaidApi.AssetReportAuditCopyCreate(ctx).AssetReportAuditCopyCreateRequest(*auditCopyCreateRequest).Execute()
	assert.NoError(t, err)
	auditCopyToken := auditCopyCreateResponse.AuditCopyToken
	assert.NotEmpty(t, auditCopyToken)

	// Get the audit copy
	acGetRequest := plaid.NewAssetReportAuditCopyGetRequest(auditCopyToken)
	acGetResponse, _, err := testClient.PlaidApi.AssetReportAuditCopyGet(ctx).AssetReportAuditCopyGetRequest(*acGetRequest).Execute()
	assert.NoError(t, err)
	assert.NotEqual(t, plaid.AssetReport{}, acGetResponse.Report)

	// Remove the audit copy token
	acRemoveRequest := plaid.NewAssetReportAuditCopyRemoveRequest(auditCopyToken)
	acRemoveResponse, _, err := testClient.PlaidApi.AssetReportAuditCopyRemove(ctx).AssetReportAuditCopyRemoveRequest(*acRemoveRequest).Execute()
	assert.NoError(t, err)
	assert.True(t, acRemoveResponse.Removed)

	// Remove the asset report
	removeRequest := plaid.NewAssetReportRemoveRequest(assetReportToken)
	removeResponse, _, err := testClient.PlaidApi.AssetReportRemove(ctx).AssetReportRemoveRequest(*removeRequest).Execute()
	assert.NoError(t, err)
	assert.True(t, removeResponse.Removed)
}

const numRetries = 20

func pollForAssetReport(t *testing.T, ctx context.Context, testClient *plaid.APIClient, assetReportToken string) (*plaid.AssetReportGetResponse, error) {
	request := plaid.NewAssetReportGetRequest()
	request.SetAssetReportToken(assetReportToken)
	for i := 0; i < numRetries; i++ {
		response, _, err := testClient.PlaidApi.AssetReportGet(ctx).AssetReportGetRequest(*request).Execute()
		if err != nil {
			plaidErr, err := plaid.ToPlaidError(err)
			assert.NoError(t, err)
			if plaidErr.ErrorCode == "PRODUCT_NOT_READY" {
				time.Sleep(2 * time.Second)
				continue
			} else {
				return nil, err
			}
		} else {
			return &response, nil
		}
	}
	return nil, errors.New("failed to get asset report")
}
