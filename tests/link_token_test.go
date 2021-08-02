package tests

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/plaid/plaid-go"
	"github.com/stretchr/testify/assert"
)

func TestLinkTokenCreateRequired(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	request := plaid.NewLinkTokenCreateRequest("Plaid Test", "en", []plaid.CountryCode{plaid.COUNTRYCODE_US}, *plaid.NewLinkTokenCreateRequestUser(time.Now().String()))
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
	resp, _, err := testClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(resp.GetLinkToken(), "link-sandbox"))
	assert.NotZero(t, resp.GetExpiration())
	assert.NotZero(t, resp.GetRequestId())
}

func TestLinkTokenCreateOptional(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId:             time.Now().String(),
		LegalName:                plaid.PtrString("Legal Name"),
		PhoneNumber:              plaid.PtrString("2025550165"),
		PhoneNumberVerifiedTime:  plaid.PtrString(time.Now().Format(RFC3339Nano)),
		EmailAddress:             plaid.PtrString("test@email.com"),
		EmailAddressVerifiedTime: plaid.PtrString(time.Now().Format(RFC3339Nano)),
		Ssn:                      plaid.PtrString("123-22-1234"),
	}
	request := plaid.NewLinkTokenCreateRequest(
		"Plaid Test",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
	request.SetLinkCustomizationName("default")
	request.SetWebhook("https://webhook-uri.com")
	request.SetAccountFilters(plaid.LinkTokenAccountFilters{
		Depository: &plaid.DepositoryFilter{
			AccountSubtypes: []plaid.AccountSubtype{plaid.ACCOUNTSUBTYPE_CHECKING, plaid.ACCOUNTSUBTYPE_SAVINGS},
		},
	})
	resp, _, err := testClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(resp.LinkToken, "link-sandbox"))
	assert.NotZero(t, resp.Expiration)
	assert.NotZero(t, resp.RequestId)
}

func TestLinkTokenCreateThenGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId:             time.Now().String(),
		LegalName:                plaid.PtrString("Legal Name"),
		PhoneNumber:              plaid.PtrString("2025550165"),
		PhoneNumberVerifiedTime:  plaid.PtrString(time.Now().Format(RFC3339Nano)),
		EmailAddress:             plaid.PtrString("test@email.com"),
		EmailAddressVerifiedTime: plaid.PtrString(time.Now().Format(RFC3339Nano)),
		Ssn:                      plaid.PtrString("123-22-1234"),
	}

	request := plaid.NewLinkTokenCreateRequest(
		"Plaid Test",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
	request.SetLinkCustomizationName("default")
	request.SetWebhook("https://webhook-uri.com")
	request.SetAccountFilters(plaid.LinkTokenAccountFilters{
		Depository: &plaid.DepositoryFilter{
			AccountSubtypes: []plaid.AccountSubtype{plaid.ACCOUNTSUBTYPE_CHECKING, plaid.ACCOUNTSUBTYPE_SAVINGS},
		},
	})
	createLinkTokenResp, _, err := testClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(createLinkTokenResp.LinkToken, "link-sandbox"))
	assert.NotZero(t, createLinkTokenResp.Expiration)
	assert.NotZero(t, createLinkTokenResp.RequestId)

	linkTokenGetRequest := plaid.NewLinkTokenGetRequest(createLinkTokenResp.LinkToken)
	linkTokenGetResp, _, err := testClient.PlaidApi.LinkTokenGet(ctx).LinkTokenGetRequest(*linkTokenGetRequest).Execute()
	assert.NoError(t, err)
	assert.Equal(t, createLinkTokenResp.LinkToken, linkTokenGetResp.LinkToken)
	assert.Equal(t, []plaid.Products{plaid.PRODUCTS_AUTH}, linkTokenGetResp.Metadata.InitialProducts)
	assert.Equal(t, "https://webhook-uri.com", linkTokenGetResp.Metadata.GetWebhook())
	assert.Equal(t, []plaid.CountryCode{plaid.COUNTRYCODE_US}, linkTokenGetResp.Metadata.CountryCodes)
	assert.Equal(t, "en", linkTokenGetResp.Metadata.GetLanguage())
	assert.Equal(t, "Plaid Test", linkTokenGetResp.Metadata.GetClientName())
	assert.NotZero(t, linkTokenGetResp.GetExpiration())
	assert.NotZero(t, linkTokenGetResp.GetCreatedAt())
}
