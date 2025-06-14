package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/plaid/plaid-go/v37/plaid"
)

func TestPayStubsGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()
	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		"ins_135842",
		[]plaid.Products{plaid.PRODUCTS_INCOME_VERIFICATION},
	)

	request := plaid.NewIncomeVerificationPaystubsGetRequest()
	request.SetAccessToken(accessToken)

	resp, _, err := testClient.PlaidApi.IncomeVerificationPaystubsGet(ctx).IncomeVerificationPaystubsGetRequest(*request).Execute()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(resp.GetPaystubs()))
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
