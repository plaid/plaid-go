package plaid

import (
	"bytes"
	"encoding/json"
)

// Balance (POST /balance) retrieves real-time balance for a given access token.
//
// See https://plaid.com/docs/api/#balance.
func (c *Client) Balance(accessToken string) (postRes *PostResponse, err error) {
	jsonText, err := json.Marshal(balanceJson{
		c.clientID,
		c.secret,
		accessToken,
	})
	if err != nil {
		return nil, err
	}
	postRes, _, err = c.postAndUnmarshal("/balance", bytes.NewReader(jsonText))
	return postRes, err
}

type balanceJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}
