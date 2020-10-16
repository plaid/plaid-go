package plaid

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
)

func TestPaymentWithIban(t *testing.T) {
	iban := "GB33BUKB20201555555555"
	params := OptionalRecipientCreateParams{
		IBAN: &iban,
		Address: &PaymentRecipientAddress{
			Street:     []string{"Street Name 999"},
			City:       "City",
			PostalCode: "99999",
			Country:    "GB",
		},
	}
	paymentRecipientCreateResp, err := testClient.CreatePaymentRecipient("John Doe", params)
	assert.Nil(t, err)
	assert.NotNil(t, paymentRecipientCreateResp.RecipientID)
	recipientID := paymentRecipientCreateResp.RecipientID

	paymentRecipientGetResp, err := testClient.GetPaymentRecipient(paymentRecipientCreateResp.RecipientID)
	assert.Nil(t, err)
	assert.NotNil(t, paymentRecipientGetResp.RecipientID)
	assert.NotNil(t, paymentRecipientGetResp.Name)
	assert.NotNil(t, paymentRecipientGetResp.IBAN)
	assert.NotNil(t, paymentRecipientGetResp.Address)

	commonPaymentTestFlows(t, recipientID)
}

func TestPaymentWithBacs(t *testing.T) {
	params := OptionalRecipientCreateParams{
		BACS: &PaymentRecipientBacs{
			Account:  "12345678",
			SortCode: "010203",
		},
		Address: &PaymentRecipientAddress{
			Street:     []string{"Street Name 999"},
			City:       "City",
			PostalCode: "99999",
			Country:    "GB",
		},
	}
	paymentRecipientCreateResp, err := testClient.CreatePaymentRecipient("John Doe", params)
	assert.Nil(t, err)
	assert.NotNil(t, paymentRecipientCreateResp.RecipientID)
	recipientID := paymentRecipientCreateResp.RecipientID

	paymentRecipientGetResp, err := testClient.GetPaymentRecipient(paymentRecipientCreateResp.RecipientID)
	assert.Nil(t, err)
	assert.NotNil(t, paymentRecipientGetResp.RecipientID)
	assert.NotNil(t, paymentRecipientGetResp.Name)
	assert.NotNil(t, paymentRecipientGetResp.BACS)
	assert.NotNil(t, paymentRecipientGetResp.Address)

	commonPaymentTestFlows(t, recipientID)
}

func TestPaymentWithBacsAndIban(t *testing.T) {
	iban := "GB33BUKB20201555555555"

	params := OptionalRecipientCreateParams{
		BACS: &PaymentRecipientBacs{
			Account:  "12345678",
			SortCode: "010203",
		},
		Address: &PaymentRecipientAddress{
			Street:     []string{"Street Name 999"},
			City:       "City",
			PostalCode: "99999",
			Country:    "GB",
		},
		IBAN: &iban,
	}
	paymentRecipientCreateResp, err := testClient.CreatePaymentRecipient("John Doe", params)
	assert.Nil(t, err)
	assert.NotNil(t, paymentRecipientCreateResp.RecipientID)
	recipientID := paymentRecipientCreateResp.RecipientID

	paymentRecipientGetResp, err := testClient.GetPaymentRecipient(paymentRecipientCreateResp.RecipientID)
	assert.Nil(t, err)
	assert.NotNil(t, paymentRecipientGetResp.RecipientID)
	assert.NotNil(t, paymentRecipientGetResp.Name)
	assert.NotNil(t, paymentRecipientGetResp.IBAN)
	assert.NotNil(t, paymentRecipientGetResp.BACS)
	assert.NotNil(t, paymentRecipientGetResp.Address)

	commonPaymentTestFlows(t, recipientID)
}

func commonPaymentTestFlows(t *testing.T, recipientID string) {
	paymentRecipientListResp, err := testClient.ListPaymentRecipients()
	assert.Nil(t, err)
	assert.True(t, len(paymentRecipientListResp.Recipients) > 0)

	paymentCreateResp, err := testClient.CreatePayment(recipientID, "TestPayment", PaymentAmount{
		Currency: "GBP",
		Value:    100.0,
	})
	assert.Nil(t, err)
	assert.NotNil(t, paymentCreateResp.PaymentID)
	assert.NotNil(t, paymentCreateResp.Status)
	paymentID := paymentCreateResp.PaymentID

	linkTokenCreateResp, err := testClient.CreateLinkToken(LinkTokenConfigs{
		User: &LinkTokenUser{
			ClientUserID: time.Now().String(),
		},
		ClientName:   "Plaid Test",
		Products:     []string{"payment_initiation"},
		CountryCodes: []string{"US"},
		Language:     "en",
		PaymentInitiation: &PaymentInitiation{
			PaymentID: paymentID,
		},
	})

	assert.Nil(t, err)
	assert.NotNil(t, linkTokenCreateResp.LinkToken)
	assert.NotNil(t, linkTokenCreateResp.Expiration)

	paymentGetResp, err := testClient.GetPayment(paymentCreateResp.PaymentID)
	assert.Nil(t, err)
	assert.NotNil(t, paymentGetResp.PaymentID)
	assert.NotNil(t, paymentGetResp.Reference)
	assert.NotNil(t, paymentGetResp.Amount)
	assert.NotNil(t, paymentGetResp.Status)
	assert.NotNil(t, paymentGetResp.LastStatusUpdate)
	assert.NotNil(t, paymentGetResp.RecipientID)

	count := 10
	paymentListResp, err := testClient.ListPayments(ListPaymentsOptions{Count: &count})
	assert.Nil(t, err)
	assert.True(t, len(paymentListResp.Payments) > 0)

	cursor := paymentListResp.NextCursor
	if cursor != "" {
		_, err = testClient.ListPayments(ListPaymentsOptions{Cursor: &cursor})
		assert.Nil(t, err)
	}
}
