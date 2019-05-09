package plaid

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCreateApexToken(t *testing.T) {
	sandboxResp, err := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	assert.Nil(t, err)

	tokenResp, err := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	assert.Nil(t, err)

	// get test account
	options := GetAccountsOptions{
		AccountIDs: []string{},
	}
	accountsResp, err := testClient.GetAccountsWithOptions(tokenResp.AccessToken, options)
	assert.Nil(t, err)
	accountID := accountsResp.Accounts[0].AccountID

	// Unfortunately, Sandbox does not seem to be setup with the Apex token.
	t.Skip()

	apexTokenResp, err := testClient.CreateApexToken(tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(apexTokenResp.ProcessorToken, "processor-sandbox-"))

}

func TestCreateDwollaToken(t *testing.T) {
	sandboxResp, err := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	assert.Nil(t, err)

	tokenResp, err := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	assert.Nil(t, err)

	// get test account
	options := GetAccountsOptions{
		AccountIDs: []string{},
	}
	accountsResp, err := testClient.GetAccountsWithOptions(tokenResp.AccessToken, options)
	assert.Nil(t, err)
	accountID := accountsResp.Accounts[0].AccountID

	// Unfortunately, Sandbox does not seem to be setup with the Dwolla token.
	t.Skip()

	dwollaTokenResp, err := testClient.CreateDwollaToken(tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(dwollaTokenResp.ProcessorToken, "processor-sandbox-"))
}

func TestCreateStripeToken(t *testing.T) {
	sandboxResp, err := testClient.CreateSandboxPublicToken(sandboxInstitution, testProducts)
	assert.Nil(t, err)

	tokenResp, err := testClient.ExchangePublicToken(sandboxResp.PublicToken)
	assert.Nil(t, err)

	// get test account
	options := GetAccountsOptions{
		AccountIDs: []string{},
	}
	accountsResp, err := testClient.GetAccountsWithOptions(tokenResp.AccessToken, options)
	assert.Nil(t, err)
	accountID := accountsResp.Accounts[0].AccountID

	// Unfortunately, Sandbox does not seem to be setup with the Stripe token.
	t.Skip()

	stripeTokenResp, err := testClient.CreateStripeToken(tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(stripeTokenResp.StripeBankAccountToken, "btok_"))
}
