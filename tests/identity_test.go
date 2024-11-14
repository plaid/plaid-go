package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v31/plaid"
	"github.com/stretchr/testify/assert"
)

func TestIdentityGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	accessToken := createSandboxItem(
		t,
		ctx,
		testClient,
		FIRST_PLATYPUS_BANK,
		[]plaid.Products{plaid.PRODUCTS_IDENTITY},
	)

	identityGetResp, _, err := testClient.PlaidApi.IdentityGet(ctx).IdentityGetRequest(
		*plaid.NewIdentityGetRequest(accessToken),
	).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, identityGetResp.GetAccounts())

	for _, acc := range identityGetResp.GetAccounts() {
		assert.NotEmpty(t, acc.GetOwners())
	}
}
