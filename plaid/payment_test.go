package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPayment(t *testing.T) {
	paymentRecipientCreateResp, err := testClient.CreatePaymentRecipient("John Doe", "GB33BUKB20201555555555", PaymentRecipientAddress{
		Street:     []string{"Street Name 999"},
		City:       "City",
		PostalCode: "99999",
		Country:    "GB",
	})
	assert.Nil(t, err)
	assert.NotNil(t, paymentRecipientCreateResp.RecipientID)
	recipientID := paymentRecipientCreateResp.RecipientID

	paymentRecipientGetResp, err := testClient.GetPaymentRecipient(paymentRecipientCreateResp.RecipientID)
	assert.Nil(t, err)
	assert.NotNil(t, paymentRecipientGetResp.RecipientID)
	assert.NotNil(t, paymentRecipientGetResp.Name)
	assert.NotNil(t, paymentRecipientGetResp.IBAN)
	assert.NotNil(t, paymentRecipientGetResp.Address)

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

	paymentTokenCreateResp, err := testClient.CreatePaymentToken(paymentID)
	assert.Nil(t, err)
	assert.NotNil(t, paymentTokenCreateResp.PaymentToken)
	assert.NotNil(t, paymentTokenCreateResp.PaymentTokenExpirationTime)

	paymentGetResp, err := testClient.GetPayment(paymentCreateResp.PaymentID)
	assert.Nil(t, err)
	assert.NotNil(t, paymentGetResp.PaymentID)
	assert.NotNil(t, paymentGetResp.PaymentToken)
	assert.NotNil(t, paymentGetResp.Reference)
	assert.NotNil(t, paymentGetResp.Amount)
	assert.NotNil(t, paymentGetResp.Status)
	assert.NotNil(t, paymentGetResp.LastStatusUpdate)
	assert.NotNil(t, paymentGetResp.PaymentTokenExpirationTime)
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
