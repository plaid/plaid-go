package plaid

import (
	"strings"
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func TestCreateLinkTokenRequired(t *testing.T) {
	linkTokenResp, err := testClient.CreateLinkToken(LinkTokenConfigs{
		ClientUserID: time.Now().String(),
		ClientName:   "Plaid Test",
		Products:     []string{"auth"},
	})

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(linkTokenResp.Token, "link-sandbox"))
	assert.NotZero(t, linkTokenResp.Expiration)
}

func TestCreateLinkTokenOptional(t *testing.T) {
	linkTokenResp, err := testClient.CreateLinkToken(LinkTokenConfigs{
		ClientUserID:          time.Now().String(),
		ClientName:            "Plaid Test",
		Products:              []string{"auth"},
		CountryCodes:          []string{"US"},
		Language:              "en",
		Webhook:               "https://webhook-uri.com",
		LinkCustomizationName: "default",
		AccountFilters: map[string]map[string][]string{
			"depository": map[string][]string{
				account_subtypes: []string{"checking", "savings"},
			},
		},
	})

	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(linkTokenResp.Token, "link-sandbox"))
	assert.NotZero(t, linkTokenResp.Expiration)
}
