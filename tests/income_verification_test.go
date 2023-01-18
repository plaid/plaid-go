package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v10/plaid"
	"github.com/stretchr/testify/assert"
)

func TestPayStubsGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()
	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		"ins_129618",
		[]plaid.Products{plaid.PRODUCTS_INCOME_VERIFICATION},
	)

	request := plaid.NewIncomeVerificationPaystubsGetRequest()
	request.SetAccessToken(accessToken)

	resp, _, err := testClient.PlaidApi.IncomeVerificationPaystubsGet(ctx).IncomeVerificationPaystubsGetRequest(*request).Execute()

	assert.NoError(t, err)
	assert.Equal(t, 3, len(resp.GetPaystubs()))
}

func TestDocumentsDownload(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		"ins_129618",
		[]plaid.Products{plaid.PRODUCTS_INCOME_VERIFICATION},
	)

	request := plaid.NewIncomeVerificationDocumentsDownloadRequest()
	request.SetAccessToken(accessToken)

	resp, _, err := testClient.PlaidApi.IncomeVerificationDocumentsDownload(ctx).IncomeVerificationDocumentsDownloadRequest(*request).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestPreCheck(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	precheckEmployer := plaid.IncomeVerificationPrecheckEmployer{}
	precheckEmployer.SetName("Plaid")

	precheckUser := plaid.IncomeVerificationPrecheckUser{}
	precheckUser.SetFirstName("Jane")
	precheckUser.SetLastName("Doe")

	request := plaid.NewIncomeVerificationPrecheckRequest()
	request.SetEmployer(precheckEmployer)
	request.SetUser(precheckUser)

	resp, _, err := testClient.PlaidApi.IncomeVerificationPrecheck(ctx).IncomeVerificationPrecheckRequest(*request).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, plaid.IncomeVerificationPrecheckConfidence("HIGH"), resp.GetConfidence())
}
