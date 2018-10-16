package plaid

import (
	"encoding/json"
	"errors"
	"time"
)

type Transaction struct {
	AccountID              string      `json:"account_id"`
	Amount                 float64     `json:"amount"`
	ISOCurrencyCode        string      `json:"iso_currency_code"`
	UnofficialCurrencyCode string      `json:"unofficial_currency_code"`
	Category               []string    `json:"category"`
	CategoryID             string      `json:"category_id"`
	Date                   time.Time   `json:"date"`
	Location               Location    `json:"location"`
	Name                   string      `json:"name"`
	PaymentMeta            PaymentMeta `json:"payment_meta"`
	Pending                bool        `json:"pending"`
	PendingTransactionID   string      `json:"pending_transaction_id"`
	AccountOwner           string      `json:"account_owner"`
	ID                     string      `json:"transaction_id"`
	Type                   string      `json:"transaction_type"`
}

func (t Transaction) MarshalJSON() ([]byte, error) {
	type Alias Transaction
	type Aux struct {
		Alias
		Date string `json:"date"`
	}
	aux := Aux{
		Alias: Alias(t),
		Date:  t.Date.Format(DateLayout),
	}
	return json.Marshal(aux)
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction
	type Aux struct {
		Alias
		Date string `json:"date"`
	}
	var aux Aux
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	date, err := time.Parse(DateLayout, aux.Date)
	if err != nil {
		return err
	}
	*t = Transaction(aux.Alias)
	t.Date = date
	return nil
}

type Location struct {
	Address     string  `json:"address"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	State       string  `json:"state"`
	StoreNumber string  `json:"store_number"`
	Zip         string  `json:"zip"`
}

type PaymentMeta struct {
	ByOrderOf        string `json:"by_order_of"`
	Payee            string `json:"payee"`
	Payer            string `json:"payer"`
	PaymentMethod    string `json:"payment_method"`
	PaymentProcessor string `json:"payment_processor"`
	PPDID            string `json:"ppd_id"`
	Reason           string `json:"reason"`
	ReferenceNumber  string `json:"reference_number"`
}

type getTransactionsRequestOptions struct {
	AccountIDs []string `json:"account_ids"`
	Count      int      `json:"count"`
	Offset     int      `json:"offset"`
}

type getTransactionsRequest struct {
	ClientID    string                        `json:"client_id"`
	Secret      string                        `json:"secret"`
	AccessToken string                        `json:"access_token"`
	StartDate   string                        `json:"start_date"`
	EndDate     string                        `json:"end_date"`
	Options     getTransactionsRequestOptions `json:"options,omitempty"`
}

type GetTransactionsResponse struct {
	APIResponse
	Accounts          []Account     `json:"accounts"`
	Item              Item          `json:"item"`
	Transactions      []Transaction `json:"transactions"`
	TotalTransactions int           `json:"total_transactions"`
}

type GetTransactionsOptions struct {
	StartDate  time.Time
	EndDate    time.Time
	AccountIDs []string
	Count      int
	Offset     int
}

// GetTransactionsWithOptions retrieves user-authorized transaction data for credit and depository-type accounts.
// See https://plaid.com/docs/api/#transactions.
func (c *Client) GetTransactionsWithOptions(accessToken string, options GetTransactionsOptions) (resp GetTransactionsResponse, err error) {
	if options.StartDate.IsZero() || options.EndDate.IsZero() {
		return resp, errors.New("/transactions/get - start date and end date must be specified")
	}

	req := getTransactionsRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
		StartDate:   options.StartDate.Format(DateLayout),
		EndDate:     options.EndDate.Format(DateLayout),
		Options: getTransactionsRequestOptions{
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

	err = c.Call("/transactions/get", jsonBody, &resp)
	return resp, err
}

// GetTransactions retrieves user-authorized transaction data for credit and depository-type accounts.
// See https://plaid.com/docs/api/#transactions.
func (c *Client) GetTransactions(accessToken string, startDate, endDate time.Time) (resp GetTransactionsResponse, err error) {
	options := GetTransactionsOptions{
		StartDate:  startDate,
		EndDate:    endDate,
		AccountIDs: []string{},
		Count:      100,
		Offset:     0,
	}
	return c.GetTransactionsWithOptions(accessToken, options)
}
