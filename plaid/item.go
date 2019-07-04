package plaid

import (
	"encoding/json"
	"errors"
)

type Item struct {
	AvailableProducts []string `json:"available_products"`
	BilledProducts    []string `json:"billed_products"`
	Error             Error    `json:"error"`
	InstitutionID     string   `json:"institution_id"`
	ItemID            string   `json:"item_id"`
	Webhook           string   `json:"webhook"`
}

type getItemRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

type GetItemResponse struct {
	APIResponse
	Item Item `json:"item"`
}

type removeItemRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

type RemoveItemResponse struct {
	APIResponse
	Removed bool `json:"removed"`
}

type updateItemWebhookRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	Webhook     string `json:"webhook"`
}

type UpdateItemWebhookResponse struct {
	APIResponse
	Item Item `json:"item"`
}

type invalidateAccessTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

type InvalidateAccessTokenResponse struct {
	APIResponse
	NewAccessToken string `json:"new_access_token"`
}

type updateAccessTokenVersionRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token_v1"`
}

type UpdateAccessTokenVersionResponse struct {
	APIResponse
	NewAccessToken string `json:"access_token"`
	ItemID         string `json:"item_id"`
}

type createPublicTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

type CreatePublicTokenResponse struct {
	APIResponse
	PublicToken string `json:"public_token"`
}

type exchangePublicTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	PublicToken string `json:"public_token"`
}

type ExchangePublicTokenResponse struct {
	APIResponse
	AccessToken string `json:"access_token"`
	ItemID      string `json:"item_id"`
}

// GetItem retrieves an item associated with an access token.
// See https://plaid.com/docs/api/#retrieve-item.
func (c *Client) GetItem(accessToken string) (resp GetItemResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/item/get - access token must be specified")
	}

	jsonBody, err := json.Marshal(getItemRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/item/get", jsonBody, &resp)
	return resp, err
}

// RemoveItem removes an item associated with an access token.
// See https://plaid.com/docs/api/#remove-an-item.
func (c *Client) RemoveItem(accessToken string) (resp RemoveItemResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/item/remove - access token must be specified")
	}

	jsonBody, err := json.Marshal(removeItemRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/item/remove", jsonBody, &resp)
	return resp, err
}

// UpdateItemWebhook updates the webhook associated with an Item.
// See https://plaid.com/docs/api/#update-webhook.
func (c *Client) UpdateItemWebhook(accessToken, webhook string) (resp UpdateItemWebhookResponse, err error) {
	if accessToken == "" || webhook == "" {
		return resp, errors.New("/item/webhook/update - access token and webhook must be specified")
	}

	jsonBody, err := json.Marshal(updateItemWebhookRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
		Webhook:     webhook,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/item/webhook/update", jsonBody, &resp)
	return resp, err
}

// InvalidateAccessToken invalidates and rotates an access token.
// See https://plaid.com/docs/api/#rotate-access-token.
func (c *Client) InvalidateAccessToken(accessToken string) (resp InvalidateAccessTokenResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/item/access_token/invalidate - access token must be specified")
	}

	jsonBody, err := json.Marshal(invalidateAccessTokenRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/item/access_token/invalidate", jsonBody, &resp)
	return resp, err
}

// UpdateAccessTokenVersion generates an updated access token associated with
// the legacy version of Plaid's API.
// See https://plaid.com/docs/api/#update-access-token-version.
func (c *Client) UpdateAccessTokenVersion(accessToken string) (resp UpdateAccessTokenVersionResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/item/access_token/update_version - access token must be specified")
	}

	jsonBody, err := json.Marshal(updateAccessTokenVersionRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/item/access_token/update_version", jsonBody, &resp)
	return resp, err
}

// CreatePublicToken generates a one-time use public token which expires in
// 30 minutes to update an Item.
// See https://plaid.com/docs/api/#creating-public-tokens.
func (c *Client) CreatePublicToken(accessToken string) (resp CreatePublicTokenResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/item/public_token/create - access token must be specified")
	}

	jsonBody, err := json.Marshal(createPublicTokenRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/item/public_token/create", jsonBody, &resp)
	return resp, err
}

// ExchangePublicToken exchanges a public token for an access token.
// See https://plaid.com/docs/api/#exchange-token-flow.
func (c *Client) ExchangePublicToken(publicToken string) (resp ExchangePublicTokenResponse, err error) {
	if publicToken == "" {
		return resp, errors.New("/item/public_token/exchange - public token must be specified")
	}

	jsonBody, err := json.Marshal(exchangePublicTokenRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		PublicToken: publicToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/item/public_token/exchange", jsonBody, &resp)
	return resp, err
}
