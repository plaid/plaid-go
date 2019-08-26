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

	dwollaTokenResp, err := testClient.CreateDwollaToken(tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(dwollaTokenResp.ProcessorToken, "processor-sandbox-"))
}

func TestCreateOcrolusToken(t *testing.T) {
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

	ocrolusTokenResp, err := testClient.CreateOcrolusToken(tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(ocrolusTokenResp.ProcessorToken, "processor-sandbox-"))
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

	t.Skip("CircleCI is not setup to work with stripe")

	stripeTokenResp, err := testClient.CreateStripeToken(tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(stripeTokenResp.StripeBankAccountToken, "btok_"))
}
