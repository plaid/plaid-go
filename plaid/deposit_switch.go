package plaid

import (
	"encoding/json"
	"errors"
)

// GetDepositSwitchResponse details the response from /deposit_switch/get
type GetDepositSwitchResponse struct {
	DepositSwitchID               string  `json:"deposit_switch_id"`
	TargetItemID                  string  `json:"target_item_id"`
	TargetAccountID               string  `json:"target_account_id"`
	State                         string  `json:"state"`
	RequestID                     string  `json:"request_id"`
	CreatedDate                   string  `json:"created_date,omitempty"`
	IsAllocatedRemainder          bool    `json:"is_allocated_remainder,omitempty"`
	AccountHasMultipleAllocations bool    `json:"account_has_multiple_allocations,omitempty"`
	PercentAllocated              float64 `json:"percent_allocated,omitempty"`
	CompletedDate                 string  `json:"completed_date,omitempty"`
}

type getDepositSwitchRequest struct {
	ClientID        string `json:"client_id"`
	Secret          string `json:"secret"`
	DepositSwitchID string `json:"deposit_switch_id"`
}

// GetDepositSwitch retrieves deposit switch data.
func (c *Client) GetDepositSwitch(
	depositSwitchID string,
) (resp GetDepositSwitchResponse, err error) {
	if depositSwitchID == "" {
		return resp, errors.New("/deposit_switch/get - deposit switch id must be specified")
	}
	req := getDepositSwitchRequest{
		ClientID:        c.clientID,
		Secret:          c.secret,
		DepositSwitchID: depositSwitchID,
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/deposit_switch/get", jsonBody, &resp)

	return resp, err
}

// this is exactly the same as getDepositSwitchRequest
type createDepositSwitchTokenRequest struct {
	ClientID        string `json:"client_id"`
	Secret          string `json:"secret"`
	DepositSwitchID string `json:"deposit_switch_id"`
}

type createDepositSwitchTokenResponse struct {
	DepositSwitchToken string `json:"deposit_switch_token"`
	Expires            string `json:"deposit_switch_token_expiration_time"`
}

func (c *Client) CreateDepositSwitchToken(
	depositSwitchID string,
) (resp createDepositSwitchTokenResponse, err error) {
	if depositSwitchID == "" {
		return resp, errors.New("/deposit_switch/token/create - deposit switch id must be specified")
	}
	req := createDepositSwitchTokenRequest{
		ClientID:        c.clientID,
		Secret:          c.secret,
		DepositSwitchID: depositSwitchID,
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/deposit_switch/token/create", jsonBody, &resp)

	return resp, err
}

type createDepositSwitchRequest struct {
	ClientID          string `json:"client_id"`
	Secret            string `json:"secret"`
	TargetAccountID   string `json:"target_account_id"`
	TargetAccessToken string `json:"target_access_token"`
}

type createDepositSwitchResponse struct {
	DepositSwitchID string `json:"deposit_switch_id"`
}

func (c *Client) CreateDepositSwitch(targetAccountID string, targetAccessToken string) (resp createDepositSwitchResponse, err error) {
	if targetAccountID == "" {
		return resp, errors.New("/deposit_switch/create - target account id must be specified")
	}
	if targetAccessToken == "" {
		return resp, errors.New("/deposit_switch/create - target access token must be specified")
	}
	req := createDepositSwitchRequest{
		ClientID:          c.clientID,
		Secret:            c.secret,
		TargetAccountID:   targetAccountID,
		TargetAccessToken: targetAccessToken,
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/deposit_switch/create", jsonBody, &resp)

	return resp, err
}
