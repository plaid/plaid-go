package plaid

import (
	"encoding/json"
	"errors"
)

type Identity struct {
	Addresses    []Address     `json:"addresses"`
	Emails       []Email       `json:"emails"`
	Names        []string      `json:"names"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
}

type Address struct {
	Data    AddressData `json:"data"`
	Primary bool        `json:"primary"`
}

type AddressData struct {
	City       string `json:"city"`
	Region     string `json:"region"`
	Street     string `json:"street"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type Email struct {
	Data    string `json:"data"`
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
}

type PhoneNumber struct {
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
	Data    string `json:"data"`
}

type getIdentityRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

type AccountWithOwners struct {
	Owners []Identity `json:"owners"`
	Account
}

type GetIdentityResponse struct {
	APIResponse
	Accounts []AccountWithOwners `json:"accounts"`
	Item     Item                `json:"item"`
}

// GetIdentity retrieves various account holder information on file with an
// associated financial institution.
// See https://plaid.com/docs/api/#identity.
func (c *Client) GetIdentity(accessToken string) (resp GetIdentityResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/identity/get - access token must be specified")
	}

	jsonBody, err := json.Marshal(getIdentityRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/identity/get", jsonBody, &resp)
	return resp, err
}
