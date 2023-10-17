package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v17/plaid"
	"github.com/stretchr/testify/assert"
)

func TestLiabilitiesGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(t, ctx, testClient, FIRST_PLATYPUS_BANK, []plaid.Products{plaid.PRODUCTS_LIABILITIES})

	request := plaid.NewLiabilitiesGetRequest(accessToken)
	resp, _, err := testClient.PlaidApi.LiabilitiesGet(ctx).LiabilitiesGetRequest(*request).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Accounts)
	assert.Len(t, resp.Liabilities.Student, 1)
	assert.Len(t, resp.Liabilities.Credit, 1)
	assert.Len(t, resp.Liabilities.Mortgage, 1)

	accountID := resp.GetAccounts()[7].AccountId
	request.SetOptions(plaid.LiabilitiesGetRequestOptions{
		AccountIds: &[]string{accountID},
	})
	resp, _, err = testClient.PlaidApi.LiabilitiesGet(ctx).LiabilitiesGetRequest(*request).Execute()

	assert.NoError(t, err)
	assert.Len(t, resp.Accounts, 1)
	assert.Len(t, resp.Liabilities.Student, 1)
	assert.Len(t, resp.Liabilities.Mortgage, 0)
}
