package plaid

import (
	"bytes"
	"encoding/json"
)

// POST /auth
// Submits a new user given a set of credentials
func (c client) AuthAddUser(username, password, pin, institutionType string,
	options *AuthOptions) (postRes *postResponse, mfaRes *mfaResponse, err error) {

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
	return postAndUnmarshal(c.environment, "/auth",
		bytes.NewReader(jsonText))
}

// POST /auth/step
// Submits an mfa "send_method", e.g. {"mask":"xxx-xxx-5309"}
func (c client) AuthStepSendMethod(accessToken, key, value string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

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
	return postAndUnmarshal(c.environment, "/auth/step",
		bytes.NewReader(jsonText))
}

// POST /auth/step
// Submits an mfa answer
func (c client) AuthStep(accessToken, answer string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(authStepJson{
		c.clientID,
		c.secret,
		accessToken,
		answer,
	})
	if err != nil {
		return nil, nil, err
	}
	return postAndUnmarshal(c.environment, "/auth/step",
		bytes.NewReader(jsonText))
}

// POST /auth/get
// Retrieves account data for an access token
func (c client) AuthGet(accessToken string) (postRes *postResponse, err error) {
	jsonText, err := json.Marshal(authGetJson{
		c.clientID,
		c.secret,
		accessToken,
	})
	if err != nil {
		return nil, err
	}
	// /auth/get will never return an MFA response
	postRes, _, err = postAndUnmarshal(c.environment, "/auth/get",
		bytes.NewReader(jsonText))
	return postRes, err
}

// DELETE /auth
// Deletes data associated with an access token
func (c client) AuthDelete(accessToken string) (deleteRes *deleteResponse, err error) {
	jsonText, err := json.Marshal(authDeleteJson{
		c.clientID,
		c.secret,
		accessToken,
	})
	if err != nil {
		return nil, err
	}
	return deleteAndUnmarshal(c.environment, "/auth",
		bytes.NewReader(jsonText))
}

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

type authDeleteJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}
