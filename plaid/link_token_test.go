package plaid

import (
	"strings"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func TestCreateLinkTokenRequired(t *testing.T) {
	linkTokenResp, err := testClient.CreateLinkToken(LinkTokenConfigs{
		User: &LinkTokenUser{
			ClientUserID: time.Now().String(),
		},
		ClientName:   "Plaid Test",
		Products:     []string{"auth"},
		CountryCodes: []string{"US"},
		Language:     "en",
	})

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(linkTokenResp.LinkToken, "link-sandbox"))
	assert.NotZero(t, linkTokenResp.Expiration)
	assert.NotZero(t, linkTokenResp.RequestID)
}

func TestCreateLinkTokenOptional(t *testing.T) {
	linkTokenResp, err := testClient.CreateLinkToken(LinkTokenConfigs{
		User: &LinkTokenUser{
			ClientUserID:             time.Now().String(),
			LegalName:                "Legal Name",
			PhoneNumber:              "2025550165",
			EmailAddress:             "test@email.com",
			PhoneNumberVerifiedTime:  time.Now(),
			EmailAddressVerifiedTime: time.Now(),
		},
		ClientName:            "Plaid Test",
		Products:              []string{"auth"},
		CountryCodes:          []string{"US"},
		Language:              "en",
		Webhook:               "https://webhook-uri.com",
		LinkCustomizationName: "default",
		AccountFilters: &map[string]map[string][]string{
			"depository": {
				"account_subtypes": {"checking", "savings"},
			},
		},
	})

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(linkTokenResp.LinkToken, "link-sandbox"))
	assert.NotZero(t, linkTokenResp.Expiration)
	assert.NotZero(t, linkTokenResp.RequestID)
}

func TestCreateLinkTokenThenGet(t *testing.T) {
	createLinkTokenResp, err := testClient.CreateLinkToken(LinkTokenConfigs{
		User: &LinkTokenUser{
			ClientUserID:             time.Now().String(),
			LegalName:                "Legal Name",
			PhoneNumber:              "2025550165",
			EmailAddress:             "test@email.com",
			PhoneNumberVerifiedTime:  time.Now(),
			EmailAddressVerifiedTime: time.Now(),
		},
		ClientName:   "Plaid Test",
		Products:     []string{"auth"},
		CountryCodes: []string{"US"},
		Webhook:      "https://webhook-uri.com",
		AccountFilters: &map[string]map[string][]string{
			"depository": {
				"account_subtypes": {"checking", "savings"},
			},
		},
		Language:              "en",
		LinkCustomizationName: "default",
	})

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(createLinkTokenResp.LinkToken, "link-sandbox"))
	assert.NotZero(t, createLinkTokenResp.Expiration)
	assert.NotZero(t, createLinkTokenResp.RequestID)
	getLinkTokenResp, err := testClient.GetLinkToken(createLinkTokenResp.LinkToken)
	assert.Nil(t, err)
	assert.Equal(t, createLinkTokenResp.LinkToken, getLinkTokenResp.LinkToken)
	assert.Equal(t, []string{"auth"}, getLinkTokenResp.Metadata.InitialProducts)
	assert.Equal(t, "https://webhook-uri.com", getLinkTokenResp.Metadata.Webhook)
	assert.Equal(t, []string{"US"}, getLinkTokenResp.Metadata.CountryCodes)
	assert.Equal(t, "en", getLinkTokenResp.Metadata.Language)
	assert.Equal(t, &map[string]map[string][]string{
		"depository": {
			"account_subtypes": {"checking", "savings"},
		},
	}, getLinkTokenResp.Metadata.AccountFilters)
	assert.Equal(t, "Plaid Test", getLinkTokenResp.Metadata.ClientName)
	assert.NotZero(t, getLinkTokenResp.Expiration)
	assert.NotZero(t, getLinkTokenResp.CreatedAt)
}
