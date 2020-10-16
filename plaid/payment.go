package plaid

import (
	"encoding/json"
	"time"
)

type PaymentRecipientAddress struct {
	// Street is an array with length in range [1, 4].
	Street     []string `json:"street"`
	City       string   `json:"city"`
	PostalCode string   `json:"postal_code"`
	// Country is an uppercase ISO 3166-1 alpha-2 country code.
	Country string `json:"country"`
}

type createPaymentRecipientRequest struct {
	ClientID string                   `json:"client_id"`
	Secret   string                   `json:"secret"`
	Name     string                   `json:"name"`
	IBAN     *string                  `json:"iban",omitempty`
	Address  *PaymentRecipientAddress `json:"address",omitempty`
	BACS     *PaymentRecipientBacs    `json:"bacs",omitempty`
}

type OptionalRecipientCreateParams struct {
	Address *PaymentRecipientAddress
	BACS    *PaymentRecipientBacs
	IBAN    *string
}

type PaymentRecipientBacs struct {
	Account  string `json:"account"`
	SortCode string `json:"sort_code"`
}

type CreatePaymentRecipientResponse struct {
	APIResponse
	RecipientID string `json:"recipient_id"`
}

func (c *Client) CreatePaymentRecipient(
	name string,
	params OptionalRecipientCreateParams,
) (resp CreatePaymentRecipientResponse, err error) {
	jsonBody, err := json.Marshal(createPaymentRecipientRequest{
		ClientID: c.clientID,
		Secret:   c.secret,
		Name:     name,
		Address:  params.Address,
		IBAN:     params.IBAN,
		BACS:     params.BACS,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/payment_initiation/recipient/create", jsonBody, &resp)
	return resp, err
}

type getPaymentRecipientRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	RecipientID string `json:"recipient_id"`
}

type Recipient struct {
	RecipientID string                   `json:"recipient_id"`
	Name        string                   `json:"name"`
	IBAN        *string                  `json:"iban",omitempty`
	Address     *PaymentRecipientAddress `json:"address"`
	BACS        *PaymentRecipientBacs    `json:"bacs",omitempty`
}

type GetPaymentRecipientResponse struct {
	APIResponse
	Recipient
}

func (c *Client) GetPaymentRecipient(recipientID string) (resp GetPaymentRecipientResponse, err error) {
	jsonBody, err := json.Marshal(getPaymentRecipientRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		RecipientID: recipientID,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/payment_initiation/recipient/get", jsonBody, &resp)
	return resp, err
}

type listPaymentRecipientsRequest struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
}

type ListPaymentRecipientsResponse struct {
	APIResponse
	Recipients []Recipient `json:"recipients"`
}

func (c *Client) ListPaymentRecipients() (resp ListPaymentRecipientsResponse, err error) {
	jsonBody, err := json.Marshal(listPaymentRecipientsRequest{
		ClientID: c.clientID,
		Secret:   c.secret,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/payment_initiation/recipient/list", jsonBody, &resp)
	return resp, err
}

type PaymentAmount struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

type createPaymentRequest struct {
	ClientID    string        `json:"client_id"`
	Secret      string        `json:"secret"`
	RecipientID string        `json:"recipient_id"`
	Reference   string        `json:"reference"`
	Amount      PaymentAmount `json:"amount"`
}

type CreatePaymentResponse struct {
	APIResponse
	PaymentID string `json:"payment_id"`
	Status    string `json:"status"`
}

func (c *Client) CreatePayment(
	recipientID string,
	reference string,
	amount PaymentAmount,
) (resp CreatePaymentResponse, err error) {
	jsonBody, err := json.Marshal(createPaymentRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		RecipientID: recipientID,
		Reference:   reference,
		Amount:      amount,
	})
	if err != nil {
		return resp, err
	}

	err = c.Call("/payment_initiation/payment/create", jsonBody, &resp)
	return resp, err
}

type getPaymentRequest struct {
	ClientID  string `json:"client_id"`
	Secret    string `json:"secret"`
	PaymentID string `json:"payment_id"`
}

type Payment struct {
	PaymentID        string        `json:"payment_id"`
	Reference        string        `json:"reference"`
	Amount           PaymentAmount `json:"amount"`
	Status           string        `json:"status"`
	LastStatusUpdate time.Time     `json:"last_status_update"`
	RecipientID      string        `json:"recipient_id"`
}

type GetPaymentResponse struct {
	APIResponse
	Payment
}

func (c *Client) GetPayment(paymentID string) (resp GetPaymentResponse, err error) {
	jsonBody, err := json.Marshal(getPaymentRequest{
		ClientID:  c.clientID,
		Secret:    c.secret,
		PaymentID: paymentID,
	})
	if err != nil {
		return resp, err
	}

	err = c.Call("/payment_initiation/payment/get", jsonBody, &resp)
	return resp, err
}

type listPaymentsRequest struct {
	ClientID string  `json:"client_id"`
	Secret   string  `json:"secret"`
	Count    *int    `json:"count"`
	Cursor   *string `json:"cursor"`
}

type ListPaymentsResponse struct {
	APIResponse
	Payments   []Payment `json:"payments"`
	NextCursor string    `json:"next_cursor"`
}

type ListPaymentsOptions struct {
	Count  *int
	Cursor *string
}

func (c *Client) ListPayments(options ListPaymentsOptions) (resp ListPaymentsResponse, err error) {
	jsonBody, err := json.Marshal(listPaymentsRequest{
		ClientID: c.clientID,
		Secret:   c.secret,
		Count:    options.Count,
		Cursor:   options.Cursor,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/payment_initiation/payment/list", jsonBody, &resp)
	return resp, err
}
