package plaid

import (
	"bytes"
	"encoding/json"
)

// AuthAddUser (POST /auth) submits a set of user credentials to add an Auth user.
//
// See https://plaid.com/docs/api/#add-auth-user.
func (c *Client) AuthAddUser(username, password, pin, institutionType string,
	options *AuthOptions) (postRes *PostResponse, mfaRes *MFAResponse, err error) {

	jsonText, err := json.Marshal(authJson{
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
	return c.postAndUnmarshal("/auth", bytes.NewReader(jsonText))
}

// AuthStepSendMethod (POST /auth/step) specifies a particular send method for MFA,
// e.g. `{"mask":"xxx-xxx-5309"}`.
//
// See https://plaid.com/docs/api/#auth-mfa.
func (c *Client) AuthStepSendMethod(accessToken, key, value string) (postRes *PostResponse,
	mfaRes *MFAResponse, err error) {

	sendMethod := map[string]string{key: value}
	jsonText, err := json.Marshal(authStepSendMethodJson{
		c.clientID,
		c.secret,
		accessToken,
		authStepOptions{sendMethod},
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/auth/step", bytes.NewReader(jsonText))
}

// AuthStep (POST /auth/step) submits an MFA answer for a given access token.
//
// See https://plaid.com/docs/api/#auth-mfa.
func (c *Client) AuthStep(accessToken, answer string) (postRes *PostResponse,
	mfaRes *MFAResponse, err error) {

	jsonText, err := json.Marshal(authStepJson{
		c.clientID,
		c.secret,
		accessToken,
		answer,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/auth/step", bytes.NewReader(jsonText))
}

// AuthGet (POST /auth/get) retrieves account data for a given access token.
//
// See https://plaid.com/docs/api/#get-auth-data.
func (c *Client) AuthGet(accessToken string) (postRes *PostResponse, err error) {
	jsonText, err := json.Marshal(authGetJson{
		c.clientID,
		c.secret,
		accessToken,
	})
	if err != nil {
		return nil, err
	}
	// /auth/get will never return an MFA response
	postRes, _, err = c.postAndUnmarshal("/auth/get", bytes.NewReader(jsonText))
	return postRes, err
}

// AuthUpdate (PATCH /auth) updates user credentials for a given access token.
//
// See https://plaid.com/docs/api/#update-auth-user.
func (c *Client) AuthUpdate(username, password, pin, accessToken string) (postRes *PostResponse,
	mfaRes *MFAResponse, err error) {

	jsonText, err := json.Marshal(authUpdateJson{
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
	return c.patchAndUnmarshal("/auth", bytes.NewReader(jsonText))
}

// AuthUpdateStep (PATCH /auth/step) updates user credentials and MFA for a given access token.
//
// See https://plaid.com/docs/api/#update-auth-user.
func (c *Client) AuthUpdateStep(username, password, pin, mfa, accessToken string) (postRes *PostResponse,
	mfaRes *MFAResponse, err error) {

	jsonText, err := json.Marshal(authUpdateStepJson{
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
	return c.patchAndUnmarshal("/auth/step", bytes.NewReader(jsonText))
}

// AuthDelete (DELETE /auth) deletes data for a given access token.
//
// See https://plaid.com/docs/api/#delete-auth-user.
func (c *Client) AuthDelete(accessToken string) (deleteRes *DeleteResponse, err error) {
	jsonText, err := json.Marshal(authDeleteJson{
		c.clientID,
		c.secret,
		accessToken,
	})
	if err != nil {
		return nil, err
	}
	return c.deleteAndUnmarshal("/auth", bytes.NewReader(jsonText))
}

// AuthOptions represents options associated with adding an Auth user.
//
// See https://plaid.com/docs/api/#add-auth-user.
type AuthOptions struct {
	List bool `json:"list"`
}
type authJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	Type     string `json:"type"`

	Username string       `json:"username"`
	Password string       `json:"password"`
	PIN      string       `json:"pin,omitempty"`
	Options  *AuthOptions `json:"options,omitempty"`
}

type authStepOptions struct {
	SendMethod map[string]string `json:"send_method"`
}

type authStepSendMethodJson struct {
	ClientID    string          `json:"client_id"`
	Secret      string          `json:"secret"`
	AccessToken string          `json:"access_token"`
	Options     authStepOptions `json:"options"`
}

type authStepJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`

	MFA string `json:"mfa"`
}

type authGetJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

type authUpdateJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`

	Username    string `json:"username"`
	Password    string `json:"password"`
	PIN         string `json:"pin,omitempty"`
	AccessToken string `json:"access_token"`
}

type authUpdateStepJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`

	Username    string `json:"username"`
	Password    string `json:"password"`
	PIN         string `json:"pin,omitempty"`
	MFA         string `json:"mfa"`
	AccessToken string `json:"access_token"`
}

type authDeleteJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}
