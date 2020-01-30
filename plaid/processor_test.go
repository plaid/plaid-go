package plaid

import (
	"context"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCreateApexToken(t *testing.T) {
	sandboxResp, err := testClient.CreateSandboxPublicToken(context.Background(), sandboxInstitution, testProducts)
	assert.Nil(t, err)

	tokenResp, err := testClient.ExchangePublicToken(context.Background(), sandboxResp.PublicToken)
	assert.Nil(t, err)

	// get test account
	options := GetAccountsOptions{
		AccountIDs: []string{},
	}
	accountsResp, err := testClient.GetAccountsWithOptions(context.Background(), tokenResp.AccessToken, options)
	assert.Nil(t, err)
	accountID := accountsResp.Accounts[0].AccountID

	apexTokenResp, err := testClient.CreateApexToken(context.Background(), tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(apexTokenResp.ProcessorToken, "processor-sandbox-"))

}

func TestCreateDwollaToken(t *testing.T) {
	sandboxResp, err := testClient.CreateSandboxPublicToken(context.Background(), sandboxInstitution, testProducts)
	assert.Nil(t, err)

	tokenResp, err := testClient.ExchangePublicToken(context.Background(), sandboxResp.PublicToken)
	assert.Nil(t, err)

	// get test account
	options := GetAccountsOptions{
		AccountIDs: []string{},
	}
	accountsResp, err := testClient.GetAccountsWithOptions(context.Background(), tokenResp.AccessToken, options)
	assert.Nil(t, err)
	accountID := accountsResp.Accounts[0].AccountID

	dwollaTokenResp, err := testClient.CreateDwollaToken(context.Background(), tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(dwollaTokenResp.ProcessorToken, "processor-sandbox-"))
}

func TestCreateOcrolusToken(t *testing.T) {
	sandboxResp, err := testClient.CreateSandboxPublicToken(context.Background(), sandboxInstitution, testProducts)
	assert.Nil(t, err)

	tokenResp, err := testClient.ExchangePublicToken(context.Background(), sandboxResp.PublicToken)
	assert.Nil(t, err)

	// get test account
	options := GetAccountsOptions{
		AccountIDs: []string{},
	}
	accountsResp, err := testClient.GetAccountsWithOptions(context.Background(), tokenResp.AccessToken, options)
	assert.Nil(t, err)
	accountID := accountsResp.Accounts[0].AccountID

	ocrolusTokenResp, err := testClient.CreateOcrolusToken(context.Background(), tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(ocrolusTokenResp.ProcessorToken, "processor-sandbox-"))
}

func TestCreateStripeToken(t *testing.T) {
	sandboxResp, err := testClient.CreateSandboxPublicToken(context.Background(), sandboxInstitution, testProducts)
	assert.Nil(t, err)

	tokenResp, err := testClient.ExchangePublicToken(context.Background(), sandboxResp.PublicToken)
	assert.Nil(t, err)

	// get test account
	options := GetAccountsOptions{
		AccountIDs: []string{},
	}
	accountsResp, err := testClient.GetAccountsWithOptions(context.Background(), tokenResp.AccessToken, options)
	assert.Nil(t, err)
	accountID := accountsResp.Accounts[0].AccountID

	t.Skip("CircleCI is not setup to work with stripe")

	stripeTokenResp, err := testClient.CreateStripeToken(context.Background(), tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(stripeTokenResp.StripeBankAccountToken, "btok_"))
}
