package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCreateAssetReport(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)

	// create asset report
	createAssetReportResp, err := testClient.CreateAssetReport([]string{tokenResp.AccessToken}, 30, nil)
	assert.NoError(t, err)

	assert.NotEmpty(t, createAssetReportResp.AssetReportToken)
	assert.NotEmpty(t, createAssetReportResp.AssetReportID)
	assert.NotEmpty(t, createAssetReportResp.RequestID)
}

func TestRemoveAuditCopy(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(sandboxResp.PublicToken)

	// create asset report
	createAssetReportResp, err := testClient.CreateAssetReport([]string{tokenResp.AccessToken}, 30, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, createAssetReportResp.AssetReportToken)

	// create asset report audit copy
	createAuditCopyResp, err := testClient.CreateAuditCopy(createAssetReportResp.AssetReportToken, "ocrolus")
	assert.NoError(t, err)

	assert.NotEmpty(t, createAuditCopyResp.AuditCopyToken)

	// remove asset report audit copy
	removeAuditCopyResp, err := testClient.RemoveAuditCopy(createAuditCopyResp.AuditCopyToken)
	assert.NoError(t, err)

	assert.True(t, removeAuditCopyResp.Removed)
}
