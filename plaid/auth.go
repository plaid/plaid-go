package plaid

import (
	"encoding/json"
	"errors"
)

type getAuthRequestOptions struct {
	AccountIDs []string `json:"account_ids"`
}

type getAuthRequest struct {
	ClientID    string                `json:"client_id"`
	Secret      string                `json:"secret"`
	AccessToken string                `json:"access_token"`
	Options     getAuthRequestOptions `json:"options,omitempty"`
}

type AccountNumberCollection struct {
	ACH           []ACHNumber  `json:"ach"`
	EFT           []EFTNumber  `json:"eft"`
	International []IBANNumber `json:"international"`
	BACS          []BACSNumber `json:"bacs"`
}

type GetAuthResponse struct {
	APIResponse
	Accounts []Account               `json:"accounts"`
	Numbers  AccountNumberCollection `json:"numbers"`
}

type GetAuthOptions struct {
	AccountIDs []string
}

// GetAuthWithOptions retrieves bank account and routing numbers associated with an Item's
// checking and savings accounts, along with other information.
// See https://plaid.com/docs/api/#auth.
func (c *Client) GetAuthWithOptions(accessToken string, options GetAuthOptions) (resp GetAuthResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/auth/get - access token must be specified")
	}

	req := getAuthRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	}
	if len(options.AccountIDs) > 0 {
		req.Options.AccountIDs = options.AccountIDs
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/auth/get", jsonBody, &resp)
	return resp, err
}

// GetAuth retrieves bank account and routing numbers associated with an Item's
// checking and savings accounts, along with other information.
// See https://plaid.com/docs/api/#auth.
func (c *Client) GetAuth(accessToken string) (resp GetAuthResponse, err error) {
	options := GetAuthOptions{
		AccountIDs: []string{},
	}
	return c.GetAuthWithOptions(accessToken, options)
}
