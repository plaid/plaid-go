package plaid

import (
	"context"
	"encoding/json"
	"errors"
)

type createProcessorTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	AccountID   string `json:"account_id"`
}

type createProcessorTokenResponse struct {
	APIResponse
	ProcessorToken string `json:"processor_token"`
}

type CreateApexTokenResponse createProcessorTokenResponse
type CreateDwollaTokenResponse createProcessorTokenResponse
type CreateOcrolusTokenResponse createProcessorTokenResponse

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

func (c *Client) createProcessorToken(ctx context.Context, apiEndpoint, accessToken, accountID string) (resp createProcessorTokenResponse, err error) {
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

	err = c.Call(ctx, apiEndpoint, jsonBody, &resp)
	return resp, err
}

// CreateApexToken is used to create a new Apex processor token.
func (c *Client) CreateApexToken(ctx context.Context, accessToken, accountID string) (resp CreateApexTokenResponse, err error) {
	response, err := c.createProcessorToken(ctx, "/processor/apex/processor_token/create", accessToken, accountID)
	return CreateApexTokenResponse(response), err
}

// CreateDwollaToken is used to create a new Dwolla processor token.
func (c *Client) CreateDwollaToken(ctx context.Context, accessToken, accountID string) (resp CreateDwollaTokenResponse, err error) {
	response, err := c.createProcessorToken(ctx, "/processor/dwolla/processor_token/create", accessToken, accountID)
	return CreateDwollaTokenResponse(response), err
}

// CreateOcrolusToken is used to create a new Ocrolus processor token.
func (c *Client) CreateOcrolusToken(ctx context.Context, accessToken, accountID string) (resp CreateOcrolusTokenResponse, err error) {
	response, err := c.createProcessorToken(ctx, "/processor/ocrolus/processor_token/create", accessToken, accountID)
	return CreateOcrolusTokenResponse(response), err
}

// CreateStripeToken is used to create a new Stripe bank account token.
func (c *Client) CreateStripeToken(ctx context.Context, accessToken, accountID string) (resp CreateStripeTokenResponse, err error) {
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

	err = c.Call(ctx, "/processor/stripe/bank_account_token/create", jsonBody, &resp)
	return resp, err
}
