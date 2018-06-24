package plaid

import (
	"encoding/json"
	"errors"
)

type createApexTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	AccountID   string `json:"account_id"`
}

type CreateApexTokenResponse struct {
	APIResponse
	ProcessorToken string `json:"processor_token"`
}

type createDwollaTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	AccountID   string `json:"account_id"`
}

type CreateDwollaTokenResponse struct {
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

func (c *Client) CreateApexToken(accessToken, accountID string) (resp CreateApexTokenResponse, err error) {
	if accessToken == "" || accountID == "" {
		return resp, errors.New("/processor/apex/processor_token/create - access token and account ID must be specified")
	}

	jsonBody, err := json.Marshal(createApexTokenRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
		AccountID:   accountID,
	})
	if err != nil {
		return resp, err
	}

	err = c.Call("/processor/apex/processor_token/create", jsonBody, &resp)
	return resp, err

}

func (c *Client) CreateDwollaToken(accessToken, accountID string) (resp CreateDwollaTokenResponse, err error) {
	if accessToken == "" || accountID == "" {
		return resp, errors.New("/processor/dwolla/processor_token/create - access token and account ID must be specified")
	}

	jsonBody, err := json.Marshal(createDwollaTokenRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
		AccountID:   accountID,
	})
	if err != nil {
		return resp, err
	}

	err = c.Call("/processor/dwolla/processor_token/create", jsonBody, &resp)
	return resp, err

}

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
