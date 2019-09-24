package plaid

import (
	"encoding/json"
	"errors"
)

type Security struct {
	SecurityID             string  `json:"security_id"`
	CUSIP                  string  `json:"cusip"`
	SEDOL                  string  `json:"sedol"`
	ISIN                   string  `json:"isin"`
	InstitutionSecurityID  string  `json:"institution_security_id"`
	InstitutionID          string  `json:"institution_id"`
	ProxySecurityID        string  `json:"proxy_security_id"`
	Name                   string  `json:"name"`
	TickerSymbol           string  `json:"ticker_symbol"`
	IsCashEquivalent       bool    `json:"is_cash_equivalent"`
	Type                   string  `json:"type"`
	ClosePrice             float64 `json:"close_price"`
	ClosePriceAsOf         string  `json:"close_price_as_of"`
	ISOCurrencyCode        string  `json:"iso_currency_code"`
	UnofficialCurrencyCode string  `json:"unofficial_currency_code"`
}

type Holding struct {
	AccountID  string `json:"account_id"`
	SecurityID string `json:"security_id"`

	InstitutionValue     float64 `json:"institution_value"`
	InstitutionPrice     float64 `json:"institution_price"`
	Quantity             float64 `json:"quantity"`
	InstitutionPriceAsOf string  `json:"institution_price_as_of"`
	CostBasis            float64 `json:"cost_basis"`

	ISOCurrencyCode        string `json:"iso_currency_code"`
	UnofficialCurrencyCode string `json:"unofficial_currency_code"`
}

type getHoldingsRequest struct {
	ClientID    string             `json:"client_id"`
	Secret      string             `json:"secret"`
	AccessToken string             `json:"access_token"`
	Options     GetHoldingsOptions `json:"options,omitempty"`
}

type GetHoldingsOptions struct {
	AccountIDs []string `json:"account_ids"`
}

type GetHoldingsResponse struct {
	APIResponse
	Accounts   []Account  `json:"accounts"`
	Item       Item       `json:"item"`
	Securities []Security `json:"securities"`
	Holdings   []Holding  `json:"holdings"`
}

// GetHoldings retrieves various account holdings for investment accounts.
// See https://plaid.com/docs/api/#holdings.
func (c *Client) GetHoldings(accessToken string) (resp GetHoldingsResponse, err error) {
	options := GetHoldingsOptions{
		AccountIDs: []string{},
	}
	return c.GetHoldingsWithOptions(accessToken, options)
}

// GetHoldingsWithOptions retrieves various account holdings for investment accounts.
// See https://plaid.com/docs/api/#holdings.
func (c *Client) GetHoldingsWithOptions(accessToken string, options GetHoldingsOptions) (resp GetHoldingsResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/investments/holdings/get - access token must be specified")
	}
	req := getHoldingsRequest{
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

	err = c.Call("/investments/holdings/get", jsonBody, &resp)
	return resp, err
}
