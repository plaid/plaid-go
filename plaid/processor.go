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
	Processor   string `json:"processor,omitempty"`
}

// ProcessorTokenResponse defines the generic return format for most processor token requests 
type ProcessorTokenResponse struct {
	APIResponse
	ProcessorToken string `json:"processor_token"`
}
// CreateApexTokenResponse defines the return format for Apex processor token requests
type CreateApexTokenResponse ProcessorTokenResponse
// CreateDwollaTokenResponse defines the return format for Dwolla processor token requests
type CreateDwollaTokenResponse ProcessorTokenResponse
// CreateOcrolusTokenResponse defines the return format for Ocrolus processor token requests 
type CreateOcrolusTokenResponse ProcessorTokenResponse

type createStripeTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	AccountID   string `json:"account_id"`
}

// CreateStripeTokenResponse defines the unique return format for stripe processor token requests 
type CreateStripeTokenResponse struct {
	APIResponse
	StripeBankAccountToken string `json:"stripe_bank_account_token"`
}

func (c *Client) requestProcessorToken(apiEndpoint, accessToken, accountID string, processor string) (resp ProcessorTokenResponse, err error) {
	if accessToken == "" || accountID == "" {
		return resp, errors.New(apiEndpoint + " - access token and account ID must be specified")
	}

	requestBody := processorTokenRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
		AccountID:   accountID,
	}

	if processor != "" {
		requestBody.Processor = processor
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return resp, err
	}

	err = c.Call(apiEndpoint, jsonBody, &resp)
	return resp, err
}

// CreateProcessorToken is used to create a new generic processor token.
func (c *Client) CreateProcessorToken(accessToken, accountID string, processor string) (resp ProcessorTokenResponse, err error) {
	if processor == "" {
		return resp, errors.New("you must specify a processor")
	}

	if processor == "stripe" {
		return resp, errors.New("stripe processor tokens are not compatible with this function, use CreateStripeToken instead")
	} else if processor == "apex" {
		return resp, errors.New("apex processor tokens are not compatible with this function, use CreateStripeToken instead")
	}

	response, err := c.requestProcessorToken("processor/token/create", accessToken, accountID, processor)
	return ProcessorTokenResponse(response), err
}

// CreateApexToken is used to create a new Apex processor token.
func (c *Client) CreateApexToken(accessToken, accountID string) (resp CreateApexTokenResponse, err error) {
	response, err := c.requestProcessorToken("/processor/apex/processor_token/create", accessToken, accountID, "")
	return CreateApexTokenResponse(response), err
}

// CreateDwollaToken is used to create a new Dwolla processor token.
func (c *Client) CreateDwollaToken(accessToken, accountID string) (resp CreateDwollaTokenResponse, err error) {
	response, err := c.requestProcessorToken("/processor/dwolla/processor_token/create", accessToken, accountID, "")
	return CreateDwollaTokenResponse(response), err
}

// CreateOcrolusToken is used to create a new Ocrolus processor token.
func (c *Client) CreateOcrolusToken(accessToken, accountID string) (resp CreateOcrolusTokenResponse, err error) {
	response, err := c.requestProcessorToken("/processor/ocrolus/processor_token/create", accessToken, accountID, "")
	return CreateOcrolusTokenResponse(response), err
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
