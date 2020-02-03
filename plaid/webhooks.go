package plaid

import (
	"encoding/json"
	"errors"
)

type WebhookVerificationKey struct {
	Alg       string `json:"alg"`
	CreatedAt int64  `json:"created_at"`
	Crv       string `json:"crv"`
	ExpiredAt int64  `json:"expired_at"`
	Kid       string `json:"kid"`
	Kty       string `json:"kty"`
	Use       string `json:"use"`
	X         string `json:"x"`
	Y         string `json:"y"`
}

type GetWebhookVerificationKeyResponse struct {
	APIResponse
	Key WebhookVerificationKey `json:"key"`
}

type getWebhookVerificationKeyRequest struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	KeyID    string `json:"key_id"`
}

// GetWebhookVerificationKey retrieves the verification key for a given webhook verification key ID
// See https://plaid.com/docs/#webhook-verification.
func (c *Client) GetWebhookVerificationKey(
	keyID string,
) (resp GetWebhookVerificationKeyResponse, err error) {
	if keyID == "" {
		return resp, errors.New("/webhook_verification_key/get - key ID must be specified")
	}

	req := getWebhookVerificationKeyRequest{
		ClientID: c.clientID,
		Secret:   c.secret,
		KeyID:    keyID,
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/webhook_verification_key/get", jsonBody, &resp)
	return resp, err
}
