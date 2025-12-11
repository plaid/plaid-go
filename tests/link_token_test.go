package tests

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/plaid/plaid-go/v41/plaid"
	"github.com/stretchr/testify/assert"
)

func TestLinkTokenCreateRequired(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	request := plaid.NewLinkTokenCreateRequest("Plaid Test", "en", []plaid.CountryCode{plaid.COUNTRYCODE_US})
	request.SetUser(*plaid.NewLinkTokenCreateRequestUser(time.Now().String()))
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
	resp, _, err := testClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(resp.GetLinkToken(), fmt.Sprintf("link-%s", PlaidEnv())))
	assert.NotZero(t, resp.GetExpiration())
	assert.NotZero(t, resp.GetRequestId())
}

func TestLinkTokenCreateOptional(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()
	now := time.Now()
	nullableNow := *plaid.NewNullableTime(&now)

	request := plaid.NewLinkTokenCreateRequest(
		"Plaid Test",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
	)
	request.SetUser(plaid.LinkTokenCreateRequestUser{
		ClientUserId:             time.Now().String(),
		LegalName:                plaid.PtrString("Legal Name"),
		PhoneNumber:              plaid.PtrString("2025550165"),
		PhoneNumberVerifiedTime:  nullableNow,
		EmailAddress:             plaid.PtrString("test@email.com"),
		EmailAddressVerifiedTime: nullableNow,
		Ssn:                      plaid.PtrString("123-22-1234"),
	})
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
	request.SetLinkCustomizationName("default")
	request.SetWebhook("https://webhook-uri.com")
	request.SetAccountFilters(plaid.LinkTokenAccountFilters{
		Depository: &plaid.DepositoryFilter{
			AccountSubtypes: []plaid.DepositoryAccountSubtype{plaid.DEPOSITORYACCOUNTSUBTYPE_CHECKING, plaid.DEPOSITORYACCOUNTSUBTYPE_SAVINGS},
		},
	})
	resp, _, err := testClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(resp.LinkToken, fmt.Sprintf("link-%s", PlaidEnv())))
	assert.NotZero(t, resp.Expiration)
	assert.NotZero(t, resp.RequestId)
}

func TestLinkTokenCreateThenGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()
	now := time.Now()
	nullableNow := *plaid.NewNullableTime(&now)

	request := plaid.NewLinkTokenCreateRequest(
		"Plaid Test",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
	)
	request.SetUser(plaid.LinkTokenCreateRequestUser{
		ClientUserId:             time.Now().String(),
		LegalName:                plaid.PtrString("Legal Name"),
		PhoneNumber:              plaid.PtrString("2025550165"),
		PhoneNumberVerifiedTime:  nullableNow,
		EmailAddress:             plaid.PtrString("test@email.com"),
		EmailAddressVerifiedTime: nullableNow,
		Ssn:                      plaid.PtrString("123-22-1234"),
	})
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
	request.SetLinkCustomizationName("default")
	request.SetWebhook("https://webhook-uri.com")
	request.SetAccountFilters(plaid.LinkTokenAccountFilters{
		Depository: &plaid.DepositoryFilter{
			AccountSubtypes: []plaid.DepositoryAccountSubtype{plaid.DEPOSITORYACCOUNTSUBTYPE_CHECKING, plaid.DEPOSITORYACCOUNTSUBTYPE_SAVINGS},
		},
	})
	createLinkTokenResp, _, err := testClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(createLinkTokenResp.LinkToken, fmt.Sprintf("link-%s", PlaidEnv())))
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

func TestLinkTokenCreateThenGet_ExtendedAuth(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()
	now := time.Now()

	request := plaid.NewLinkTokenCreateRequest(
		"Plaid Test",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
	)
	request.SetUser(plaid.LinkTokenCreateRequestUser{ClientUserId: now.String()})
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_AUTH})
	enabled := true
	request.SetAuth(plaid.LinkTokenCreateRequestAuth{
		AuthTypeSelectEnabled:         &enabled,
		AutomatedMicrodepositsEnabled: &enabled,
		InstantMatchEnabled:           &enabled,
		SameDayMicrodepositsEnabled:   &enabled,
	})
	createLinkTokenResp, _, err := testClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	assert.NoError(t, err)
	assert.True(t, strings.HasPrefix(createLinkTokenResp.LinkToken, fmt.Sprintf("link-%s", PlaidEnv())))
	assert.NotZero(t, createLinkTokenResp.Expiration)
	assert.NotZero(t, createLinkTokenResp.RequestId)
}
