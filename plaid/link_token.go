package plaid

import (
	"encoding/json"
	"time"
)

type LinkTokenUser struct {
	ClientUserID             string    `json:"client_user_id"`
	LegalName                string    `json:"legal_name,omitempty"`
	PhoneNumber              string    `json:"phone_number,omitempty"`
	EmailAddress             string    `json:"email_address,omitempty"`
	PhoneNumberVerifiedTime  time.Time `json:"phone_number_verified_time,omitempty"`
	EmailAddressVerifiedTime time.Time `json:"email_address_verified_time,omitempty"`
}

type CrossAppItemAdd struct {
	TargetApplicationToken string `json:"target_application_token"`
	ForeignID              string `json:"foreign_id,omitempty"`
}

type PaymentInitiation struct {
	PaymentID string `json:"payment_id"`
}

type LinkTokenConfigs struct {
	User                  *LinkTokenUser                  `json:"user"`
	ClientName            string                          `json:"client_name"`
	Products              []string                        `json:"products,omitempty"`
	AccessToken           string                          `json:"access_token,omitempty"`
	CountryCodes          []string                        `json:"country_codes,omitempty"`
	Webhook               string                          `json:"webhook,omitempty"`
	AccountFilters        *map[string]map[string][]string `json:"account_filters,omitempty"`
	CrossAppItemAdd       *CrossAppItemAdd                `json:"cross_app_item_add,omitempty"`
	PaymentInitiation     *PaymentInitiation              `json:"payment_initiation,omitempty"`
	Language              string                          `json:"language,omitempty"`
	LinkCustomizationName string                          `json:"link_customization_name,omitempty"`
	RedirectUri           string                          `json:"redirect_uri,omitempty"`
	AndroidPackageName    string                          `json:"android_package_name,omitempty"`
}

type createLinkTokenRequest struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	LinkTokenConfigs
}

type CreateLinkTokenResponse struct {
	LinkToken  string    `json:"link_token"`
	Expiration time.Time `json:"expiration"`
}

func (c *Client) CreateLinkToken(configs LinkTokenConfigs) (resp CreateLinkTokenResponse, err error) {
	jsonBody, err := json.Marshal(createLinkTokenRequest{
		ClientID:         c.clientID,
		Secret:           c.secret,
		LinkTokenConfigs: configs,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/link/token/create", jsonBody, &resp)
	return resp, err
}
