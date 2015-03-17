package plaid

import (
	"bytes"
	"encoding/json"
)

// POST /balance
// Retrieves real-time balance for an access token
func (c client) Balance(accessToken string) (postRes *postResponse, err error) {
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
