package plaid

import (
	"context"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func TestGetIncome(t *testing.T) {
	sandboxResp, _ := testClient.CreateSandboxPublicToken(context.Background(), sandboxInstitution, testProducts)
	tokenResp, _ := testClient.ExchangePublicToken(context.Background(), sandboxResp.PublicToken)
	incomeResp, err := testClient.GetIncome(context.Background(), tokenResp.AccessToken)

	if plaidErr, ok := err.(Error); ok {
		for ok && plaidErr.ErrorCode == "PRODUCT_NOT_READY" {
			time.Sleep(5 * time.Second)
			incomeResp, err = testClient.GetIncome(context.Background(), tokenResp.AccessToken)
			plaidErr, ok = err.(Error)
		}
	}

	assert.Nil(t, err)
	assert.NotNil(t, incomeResp.Income)
}
