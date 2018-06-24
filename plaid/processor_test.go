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

	stripeTokenResp, err := testClient.CreateStripeToken(tokenResp.AccessToken, accountID)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(stripeTokenResp.StripeBankAccountToken, "btok_"))
}
