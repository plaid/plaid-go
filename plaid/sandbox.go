package plaid

import (
	"encoding/json"
	"errors"
)

type createSandboxPublicTokenRequest struct {
	InstitutionID   string   `json:"institution_id"`
	InitialProducts []string `json:"initial_products"`
	PublicKey       string   `json:"public_key"`
}

type CreateSandboxPublicTokenResponse struct {
	APIResponse
	PublicToken string `json:"public_token"`
}

type resetSandboxItemRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

type ResetSandboxItemResponse struct {
	APIResponse
	ResetLogin bool `json:"reset_login"`
}

type setSandboxItemVerificationStatusRequest struct {
	ClientID           string `json:"client_id"`
	Secret             string `json:"secret"`
	AccessToken        string `json:"access_token"`
	AccountID          string `json:"account_id"`
	VerificationStatus string `json:"verification_status"`
}

type SetSandboxItemVerificationStatusResponse struct {
	APIResponse
}

func (c *Client) CreateSandboxPublicToken(institutionID string, initialProducts []string) (resp CreateSandboxPublicTokenResponse, err error) {
	if institutionID == "" || len(initialProducts) == 0 {
		return resp, errors.New("/sandbox/public_token/create - institution id and initial products must be specified")
	}

	jsonBody, err := json.Marshal(createSandboxPublicTokenRequest{
		InstitutionID:   institutionID,
		InitialProducts: initialProducts,
		PublicKey:       c.publicKey,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/sandbox/public_token/create", jsonBody, &resp)
	return resp, err
}

func (c *Client) ResetSandboxItem(accessToken string) (resp ResetSandboxItemResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/sandbox/item/reset_login - access token must be specified")
	}

	jsonBody, err := json.Marshal(resetSandboxItemRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/sandbox/item/reset_login", jsonBody, &resp)
	return resp, err
}

func (c *Client) SetSandboxItemVerificationStatus(accessToken string, accountID string, verificationStatus string) (resp SetSandboxItemVerificationStatusResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/sandbox/item/set_verification_status - access token must be specified")
	}

	if accountID == "" {
		return resp, errors.New("/sandbox/item/set_verification_status - accountID must be specified")
	}

	if verificationStatus == "" {
		return resp, errors.New("/sandbox/item/set_verification_status - verification status must be specified")
	}

	jsonBody, err := json.Marshal(setSandboxItemVerificationStatusRequest{
		ClientID:           c.clientID,
		Secret:             c.secret,
		AccessToken:        accessToken,
		AccountID:          accountID,
		VerificationStatus: verificationStatus,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/sandbox/item/set_verification_status", jsonBody, &resp)
	return resp, err
}
