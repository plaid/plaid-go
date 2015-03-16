package plaid

import (
	"bytes"
	"encoding/json"
)

// POST /upgrade
// Upgrade an access token to an additional product
func (c client) Upgrade(accessToken, upgradeTo string,
	options *UpgradeOptions) (postRes *postResponse, mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(upgradeJson{
		c.clientID,
		c.secret,
		accessToken,
		upgradeTo,
		options,
	})
	if err != nil {
		return nil, nil, err
	}
	return postAndUnmarshal(c.environment, "/upgrade",
		bytes.NewReader(jsonText))
}

// POST /upgrade/step
// Submits an mfa "send_method", e.g. {"mask":"xxx-xxx-5309"}
func (c client) UpgradeStepSendMethod(accessToken, key, value string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	sendMethod := map[string]string{key: value}
	jsonText, err := json.Marshal(upgradeStepSendMethodJson{
		c.clientID,
		c.secret,
		accessToken,
		upgradeStepOptions{sendMethod},
	})
	if err != nil {
		return nil, nil, err
	}
	return postAndUnmarshal(c.environment, "/upgrade/step",
		bytes.NewReader(jsonText))
}

// POST /upgrade/step
// Submits an mfa answer
func (c client) UpgradeStep(accessToken, answer string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(upgradeStepJson{
		c.clientID,
		c.secret,
		accessToken,
		answer,
	})
	if err != nil {
		return nil, nil, err
	}
	return postAndUnmarshal(c.environment, "/upgrade/step",
		bytes.NewReader(jsonText))
}

type UpgradeOptions struct {
	Webhook string `json:"webhook,omitempty"`
}

type upgradeJson struct {
	ClientID    string          `json:"client_id"`
	Secret      string          `json:"secret"`
	AccessToken string          `json:"access_token"`
	UpgradeTo   string          `json:"upgrade_to"`
	Options     *UpgradeOptions `json:"options,omitempty"`
}

type upgradeStepOptions struct {
	SendMethod map[string]string `json:"send_method"`
}

type upgradeStepSendMethodJson struct {
	ClientID    string             `json:"client_id"`
	Secret      string             `json:"secret"`
	AccessToken string             `json:"access_token"`
	Options     upgradeStepOptions `json:"options"`
}

type upgradeStepJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
	MFA         string `json:"mfa"`
}
