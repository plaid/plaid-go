package plaid

import (
	"encoding/json"
	"errors"
)

type Account struct {
	AccountID          string          `json:"account_id"`
	Balances           AccountBalances `json:"balances"`
	Mask               string          `json:"mask"`
	Name               string          `json:"name"`
	OfficialName       string          `json:"official_name"`
	Subtype            string          `json:"subtype"`
	Type               string          `json:"type"`
	VerificationStatus string          `json:"verification_status"`
}

type AccountBalances struct {
	Available              float64 `json:"available"`
	Current                float64 `json:"current"`
	Limit                  float64 `json:"limit"`
	ISOCurrencyCode        string  `json:"iso_currency_code"`
	UnofficialCurrencyCode string  `json:"unofficial_currency_code"`
}

type ACHNumber struct {
	Account     string `json:"account"`
	AccountID   string `json:"account_id"`
	Routing     string `json:"routing"`
	WireRouting string `json:"wire_routing"`
}

type EFTNumber struct {
	Account     string `json:"account"`
	AccountID   string `json:"account_id"`
	Institution string `json:"institution"`
	Branch      string `json:"branch"`
}

type IBANNumber struct {
	AccountID string `json:"account_id"`
	IBAN      string `json:"iban"`
	BIC       string `json:"bic"`
}

type BACSNumber struct {
	AccountID string `json:"account_id"`
	Account   string `json:"account"`
	SortCode  string `json:"sort_code"`
}

type getBalancesRequestOptions struct {
	AccountIDs []string `json:"account_ids,omitempty"`
}

type getBalancesRequest struct {
	ClientID    string                    `json:"client_id"`
	Secret      string                    `json:"secret"`
	AccessToken string                    `json:"access_token"`
	Options     getBalancesRequestOptions `json:"options,omitempty"`
}

type GetBalancesResponse struct {
	APIResponse
	Accounts []Account `json:"accounts"`
}

type getAccountsRequestOptions struct {
	AccountIDs []string `json:"account_ids,omitempty"`
}

type getAccountsRequest struct {
	ClientID    string                    `json:"client_id"`
	Secret      string                    `json:"secret"`
	AccessToken string                    `json:"access_token"`
	Options     getAccountsRequestOptions `json:"options,omitempty"`
}

type GetAccountsResponse struct {
	APIResponse
	Accounts []Account `json:"accounts"`
	Item Item `json:"item"`
}

type GetAccountsOptions struct {
	AccountIDs []string
}

type GetBalancesOptions struct {
	AccountIDs []string
}

// GetBalances returns the real-time balance for each of an Item's accounts.
// See https://plaid.com/docs/api/#balance.
func (c *Client) GetBalances(accessToken string) (resp GetBalancesResponse, err error) {
	options := GetBalancesOptions{
		AccountIDs: []string{},
	}
	return c.GetBalancesWithOptions(accessToken, options)
}

// GetBalancesWithOptions returns the real-time balance for each of an Item's accounts.
// See https://plaid.com/docs/api/#balance.
func (c *Client) GetBalancesWithOptions(accessToken string, options GetBalancesOptions) (resp GetBalancesResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/accounts/balance/get - access token must be specified")
	}
	req := getBalancesRequest{
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

	err = c.Call("/accounts/balance/get", jsonBody, &resp)
	return resp, err
}

// GetAccountsWithOptions retrieves accounts associated with an Item.
// See https://plaid.com/docs/api/#accounts.
func (c *Client) GetAccountsWithOptions(accessToken string, options GetAccountsOptions) (resp GetAccountsResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/accounts/get - access token must be specified")
	}

	req := getAccountsRequest{
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

	err = c.Call("/accounts/get", jsonBody, &resp)
	return resp, err
}

// GetAccounts retrieves accounts associated with an Item.
// See https://plaid.com/docs/api/#accounts.
func (c *Client) GetAccounts(accessToken string) (resp GetAccountsResponse, err error) {
	options := GetAccountsOptions{
		AccountIDs: []string{},
	}
	return c.GetAccountsWithOptions(accessToken, options)
}
