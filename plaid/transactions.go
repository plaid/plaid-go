package plaid

import (
  "bytes"
  "encoding/json"
)

// TransactionsGet (POST /transactions/get) retrieves account and transaction data for a given access token.
//
// See https://plaid.com/docs/api/#get-transactions.
func (c *Client) TransactionsGet(accessToken string, options *TransactionsGetOptions) (postRes *postResponse,
  mfaRes *mfaResponse, err error) {

  jsonText, err := json.Marshal(transactionsGetJson{
    ClientID:    c.clientID,
    Secret:      c.secret,
    AccessToken: accessToken,
    Options:     options,
  })
  if err != nil {
    return nil, nil, err
  }
  return c.postAndUnmarshal("/transactions/get", bytes.NewReader(jsonText))
}

// TransactionsGetOptions represents options associated with retrieving a Connect user.
//
// See https://plaid.com/docs/api/#retrieve-transactions.
type TransactionsGetOptions struct {
  Pending                    bool   `json:"pending,omitempty"`
  Account                    string `json:"account,omitempty"`
  GTE                        string `json:"gte,omitempty"`
  LTE                        string `json:"lte,omitempty"`
  IncludeOriginalDescription bool   `json:"include_original_description,omitempty"`
  Count                      int    `json:"count,omitempty"`
  Offset                     int    `json:"offset,omitempty"`
}
type transactionsGetJson struct {
  ClientID    string `json:"client_id"`
  Secret      string `json:"secret"`
  AccessToken string `json:"access_token"`

  Options *TransactionsGetOptions `json:"options,omitempty"`
}
