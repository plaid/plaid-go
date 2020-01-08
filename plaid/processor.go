package plaid

import (
	"encoding/json"
	"errors"
)

type processorTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	AccountID   string `json:"account_id"`
	Processor   string `json:"processor"`
}

// CreateProcessorTokenResponse is the response interface for all token requests
type CreateProcessorTokenResponse interface {
	APIResponse
	GetToken() string
}

type createGenericProcessorTokenResponse struct {
	APIResponse
	ProcessorToken string `json:"processor_token"`
}

func (r *createGenericProcessorTokenResponse) GetToken() string {
	return r.ProcessorToken
}

type createStripeTokenResponse struct {
	APIResponse
	StripeBankAccountToken string `json:"stripe_bank_account_token"`
}

func (r *createStripeTokenResponse) GetToken() string {
	return r.StripeBankAccountToken
}

func (c *Client) requestProcessorToken(apiEndpoint, accessToken, accountID string, processor string) (resp CreateProcessorTokenResponse, err error) {
	if accessToken == "" || accountID == "" {
		return resp, errors.New(apiEndpoint + " - access token and account ID must be specified")
	}

	jsonBody, err := json.Marshal(processorTokenRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
		AccountID:   accountID,
		Processor:   processor,
	})
	if err != nil {
		return resp, err
	}

	err = c.Call(apiEndpoint, jsonBody, &resp)
	return resp, err
}

// CreateProcessorToken is used to create a new generic processor token.
func (c *Client) CreateProcessorToken(accessToken, accountID string, processor string) (resp CreateProcessorTokenResponse, err error) {
	if processor == "" {
		return resp, errors.New("you must specify a processor")
	}

	endpoint := "processor/token/create"

	if processor == "stripe" {
		endpoint = "/processor/stripe/bank_account_token/create"
	} else if processor == "apex" {
		endpoint = "/processor/apex/processor_token/create"
	}

	response, err := c.requestProcessorToken(endpoint, accessToken, accountID, processor)
	return CreateProcessorTokenResponse(response), err
}

// CreateApexToken is used to create a new Apex processor token.
func (c *Client) CreateApexToken(accessToken, accountID string) (resp CreateProcessorTokenResponse, err error) {
	response, err := c.requestProcessorToken("/processor/apex/processor_token/create", accessToken, accountID)
	return CreateProcessorTokenResponse(response), err
}

// CreateDwollaToken is used to create a new Dwolla processor token.
func (c *Client) CreateDwollaToken(accessToken, accountID string) (resp CreateProcessorTokenResponse, err error) {
	response, err := c.requestProcessorToken("/processor/dwolla/processor_token/create", accessToken, accountID)
	return CreateProcessorTokenResponse(response), err
}

// CreateOcrolusToken is used to create a new Ocrolus processor token.
func (c *Client) CreateOcrolusToken(accessToken, accountID string) (resp CreateProcessorTokenResponse, err error) {
	response, err := c.requestProcessorToken("/processor/ocrolus/processor_token/create", accessToken, accountID)
	return CreateProcessorTokenResponse(response), err
}

// CreateStripeToken is used to create a new Stripe bank account token.
func (c *Client) CreateStripeToken(accessToken, accountID string) (resp CreateProcessorTokenResponse, err error) {
	response, err := c.requestProcessorToken("/processor/stripe/bank_account_token/create", accessToken, accountID)
	return CreateProcessorTokenResponse(response), err
}
