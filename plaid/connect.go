package plaid

import (
	"bytes"
	"encoding/json"
)

// ConnectAddUser (POST /connect) submits a set of user credentials to add a Connect user.
//
// See https://plaid.com/docs/api/#add-user.
func (c *Client) ConnectAddUser(username, password, pin, institutionType string,
	options *ConnectOptions) (postRes *postResponse, mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(connectJson{
		c.clientID,
		c.secret,
		institutionType,
		username,
		password,
		pin,
		options,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/connect", bytes.NewReader(jsonText))
}

// ConnectStepSendMethod (POST /connect/step) specifies a particular send method for MFA,
// e.g. `{"mask":"xxx-xxx-5309"}`.
//
// See https://plaid.com/docs/api/#mfa-authentication.
func (c *Client) ConnectStepSendMethod(accessToken, key, value string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	sendMethod := map[string]string{key: value}
	jsonText, err := json.Marshal(connectStepSendMethodJson{
		c.clientID,
		c.secret,
		accessToken,
		connectStepOptions{sendMethod},
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/connect/step", bytes.NewReader(jsonText))
}

// ConnectStep (POST /connect/step) submits an MFA answer for a given access token.
//
// See https://plaid.com/docs/api/#mfa-authentication.
func (c *Client) ConnectStep(accessToken, answer string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(connectStepJson{
		c.clientID,
		c.secret,
		accessToken,
		answer,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/connect/step", bytes.NewReader(jsonText))
}

// ConnectGet (POST /connect/get) retrieves account and transaction data for a given access token.
//
// See https://plaid.com/docs/api/#get-transactions.
func (c *Client) ConnectGet(accessToken string, options *ConnectGetOptions) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(connectGetJson{
		c.clientID,
		c.secret,
		accessToken,
		options,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/connect/get", bytes.NewReader(jsonText))
}

// ConnectUpdate (PATCH /connect) updates user credentials for a given access token.
//
// See https://plaid.com/docs/api/#update-user.
func (c *Client) ConnectUpdate(username, password, pin, accessToken string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(connectUpdateJson{
		c.clientID,
		c.secret,
		username,
		password,
		pin,
		accessToken,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.patchAndUnmarshal("/connect", bytes.NewReader(jsonText))
}

// ConnectUpdateStep (PATCH /connect/step) updates user credentials and MFA for a given access token.
//
// See https://plaid.com/docs/api/#update-user.
func (c *Client) ConnectUpdateStep(username, password, pin, mfa, accessToken string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(connectUpdateStepJson{
		c.clientID,
		c.secret,
		username,
		password,
		pin,
		mfa,
		accessToken,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.patchAndUnmarshal("/connect/step", bytes.NewReader(jsonText))
}

// ConnectDelete (DELETE /connect) deletes data for a given access token.
//
// See https://plaid.com/docs/api/#delete-user.
func (c *Client) ConnectDelete(accessToken string) (deleteRes *deleteResponse, err error) {
	jsonText, err := json.Marshal(connectDeleteJson{
		c.clientID,
		c.secret,
		accessToken,
	})
	if err != nil {
		return nil, err
	}
	return c.deleteAndUnmarshal("/connect", bytes.NewReader(jsonText))
}

// ConnectOptions represents options associated with adding an Connect user.
//
// See https://plaid.com/docs/api/#add-user.
type ConnectOptions struct {
	Webhook   string `json:"webhook,omitempty"`
	Pending   bool   `json:"pending,omitempty"`
	LoginOnly bool   `json:"login_only,omitempty"`
	List      bool   `json:"list,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}
type connectJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	Type     string `json:"type"`

	Username string          `json:"username"`
	Password string          `json:"password"`
	PIN      string          `json:"pin,omitempty"`
	Options  *ConnectOptions `json:"options,omitempty"`
}

type connectStepOptions struct {
	SendMethod map[string]string `json:"send_method"`
}
type connectStepSendMethodJson struct {
	ClientID    string             `json:"client_id"`
	Secret      string             `json:"secret"`
	AccessToken string             `json:"access_token"`
	Options     connectStepOptions `json:"options"`
}

type connectStepJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`

	MFA string `json:"mfa"`
}

// ConnectGetOptions represents options associated with retrieving a Connect user.
//
// See https://plaid.com/docs/api/#retrieve-transactions.
type ConnectGetOptions struct {
	Pending bool   `json:"pending,omitempty"`
	Account string `json:"account,omitempty"`
	GTE     string `json:"gte,omitempty"`
	LTE     string `json:"lte,omitempty"`
}
type connectGetJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`

	Options *ConnectGetOptions `json:"options,omitempty"`
}

type connectUpdateJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`

	Username    string `json:"username"`
	Password    string `json:"password"`
	PIN         string `json:"pin,omitempty"`
	AccessToken string `json:"access_token"`
}

type connectUpdateStepJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`

	Username    string `json:"username"`
	Password    string `json:"password"`
	PIN         string `json:"pin,omitempty"`
	MFA         string `json:"mfa"`
	AccessToken string `json:"access_token"`
}

type connectDeleteJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}
