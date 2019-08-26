package plaid

import (
	"encoding/json"
	"errors"
)

type createProcessorTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	AccountID   string `json:"account_id"`
}

type CreateProcessorTokenResponse struct {
	APIResponse
	ProcessorToken string `json:"processor_token"`
}

type createStripeTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	AccountID   string `json:"account_id"`
}

type CreateStripeTokenResponse struct {
	APIResponse
	StripeBankAccountToken string `json:"stripe_bank_account_token"`
}

// CreateProcessorToken is a helper function used to create processor tokens for a given api endpoint.
func (c *Client) CreateProcessorToken(apiEndpoint, accessToken, accountID string) (resp CreateProcessorTokenResponse, err error) {
	if accessToken == "" || accountID == "" {
		return resp, errors.New(apiEndpoint + " - access token and account ID must be specified")
	}

	jsonBody, err := json.Marshal(createProcessorTokenRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
		AccountID:   accountID,
	})
	if err != nil {
		return resp, err
	}

	err = c.Call(apiEndpoint, jsonBody, &resp)
	return resp, err
}

// CreateApexToken is used to create a new Apex processor token.
func (c *Client) CreateApexToken(accessToken, accountID string) (resp CreateProcessorTokenResponse, err error) {
	return c.CreateProcessorToken("/processor/apex/processor_token/create", accessToken, accountID)
}

// CreateDwollaToken is used to create a new Dwolla processor token.
func (c *Client) CreateDwollaToken(accessToken, accountID string) (resp CreateProcessorTokenResponse, err error) {
	return c.CreateProcessorToken("/processor/dwolla/processor_token/create", accessToken, accountID)
}

// CreateOcrolusToken is used to create a new Ocrolus processor token.
func (c *Client) CreateOcrolusToken(accessToken, accountID string) (resp CreateProcessorTokenResponse, err error) {
	return c.CreateProcessorToken("/processor/ocrolus/processor_token/create", accessToken, accountID)
}

// CreateStripeToken is used to create a new Stripe bank account token.
func (c *Client) CreateStripeToken(accessToken, accountID string) (resp CreateStripeTokenResponse, err error) {
	if accessToken == "" || accountID == "" {
		return resp, errors.New("/processor/stripe/bank_account_token/create - access token and account ID must be specified")
	}

	jsonBody, err := json.Marshal(createStripeTokenRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
		AccountID:   accountID,
	})
	if err != nil {
		return resp, err
	}

	err = c.Call("/processor/stripe/bank_account_token/create", jsonBody, &resp)
	return resp, err
}
