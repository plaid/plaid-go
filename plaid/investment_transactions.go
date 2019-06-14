package plaid

import (
	"encoding/json"
	"errors"
)

type InvestmentTransaction struct {
	InvestmentTransactionID string `json:"investment_transaction_id"`
	AccountID               string `json:"account_id"`
	SecurityID              string `json:"security_id"`
	CancelTransactionID     string `json:"cancel_transaction_id"`

	Date                   string  `json:"date"`
	Name                   string  `json:"name"`
	Quantity               float64 `json:"quantity"`
	Amount                 float64 `json:"amount"`
	Price                  float64 `json:"price"`
	Fees                   float64 `json:"fees"`
	Type                   string  `json:"type"`
	ISOCurrencyCode        string  `json:"iso_currency_code"`
	UnofficialCurrencyCode string  `json:"unofficial_currency_code"`
}

type GetInvestmentTransactionsResponse struct {
	APIResponse
	Item                        Item                    `json:"item"`
	Accounts                    []Account               `json:"accounts"`
	InvestmentTransactions      []InvestmentTransaction `json:"investment_transactions"`
	Securities                  []Security              `json:"securities"`
	TotalInvestmentTransactions int                     `json:"total_investment_transactions"`
}

type GetInvestmentTransactionsOptions struct {
	StartDate  string
	EndDate    string
	AccountIDs []string
	Count      int
	Offset     int
}

type getInvestmentTransactionsRequest struct {
	ClientID    string                                  `json:"client_id"`
	Secret      string                                  `json:"secret"`
	AccessToken string                                  `json:"access_token"`
	StartDate   string                                  `json:"start_date"`
	EndDate     string                                  `json:"end_date"`
	Options     getInvestmentTransactionsRequestOptions `json:"options,omitempty"`
}

type getInvestmentTransactionsRequestOptions struct {
	AccountIDs []string `json:"account_ids"`
	Count      int      `json:"count"`
	Offset     int      `json:"offset"`
}

// GetInvestmentTransactionsWithOptions retrieves user-authorized investment transaction data for investment-type accounts.
// See https://plaid.com/docs/api/#investment-transactions.
func (c *Client) GetInvestmentTransactionsWithOptions(accessToken string, options GetInvestmentTransactionsOptions) (resp GetInvestmentTransactionsResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/investments/transacstions/get - access token must be specified")
	}
	if options.StartDate == "" || options.EndDate == "" {
		return resp, errors.New("/investment/transactions/get - start date and end date must be specified")
	}

	req := getInvestmentTransactionsRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
		StartDate:   options.StartDate,
		EndDate:     options.EndDate,
		Options: getInvestmentTransactionsRequestOptions{
			Count:  options.Count,
			Offset: options.Offset,
		},
	}
	if len(options.AccountIDs) > 0 {
		req.Options.AccountIDs = options.AccountIDs
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/investments/transactions/get", jsonBody, &resp)
	return resp, err
}

// GetInvestmentTransactions retrieves user-authorized transaction data for investment-type accounts.
// See https://plaid.com/docs/api/#investment-transactions.
func (c *Client) GetInvestmentTransactions(accessToken, startDate, endDate string) (resp GetInvestmentTransactionsResponse, err error) {
	options := GetInvestmentTransactionsOptions{
		StartDate:  startDate,
		EndDate:    endDate,
		AccountIDs: []string{},
		Count:      100,
		Offset:     0,
	}
	return c.GetInvestmentTransactionsWithOptions(accessToken, options)
}
