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

	commonPaymentTestFlows(t, recipientID, true, false)
}

func TestPaymentWithBacs(t *testing.T) {
	params := OptionalRecipientCreateParams{
		BACS: &PaymentRecipientBacs{
			Account:  "26207729",
			SortCode: "560029",
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

	commonPaymentTestFlows(t, recipientID, true, false)
}

func TestPaymentWithBacsAndIban(t *testing.T) {
	iban := "GB33BUKB20201555555555"

	params := OptionalRecipientCreateParams{
		BACS: &PaymentRecipientBacs{
			Account:  "26207729",
			SortCode: "560029",
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

	// testing common flows with the link_token
	commonPaymentTestFlows(t, recipientID, true, false)

	// testing common flows with the legacy payment_token
	commonPaymentTestFlows(t, recipientID, false, false)

	// testing common flows with options
	commonPaymentTestFlows(t, recipientID, false, true)
}

func commonPaymentTestFlows(t *testing.T, recipientID string, useLinkToken bool, useOptions bool) {
	paymentRecipientListResp, err := testClient.ListPaymentRecipients()
	assert.Nil(t, err)
	assert.True(t, len(paymentRecipientListResp.Recipients) > 0)
	var paymentCreateResp CreatePaymentResponse
	var options PaymentOptions
	if useOptions == true {
		requestRefundDetails := true
		options = PaymentOptions{
			BACS: &PaymentRecipientBacs{
				Account:  "1234567890",
				SortCode: "000000",
			},
			RequestRefundDetails: &requestRefundDetails,
		}
		paymentCreateResp, err = testClient.CreatePaymentWithOptions(
			recipientID,
			"TestPayment",
			PaymentAmount{
				Currency: "GBP",
				Value:    100.0,
			},
			nil,
			options,
		)
	} else {
		paymentCreateResp, err = testClient.CreatePayment(
			recipientID,
			"TestPayment",
			PaymentAmount{
				Currency: "GBP",
				Value:    100.0,
			},
			nil,
		)
	}

	assert.Nil(t, err)
	assert.NotNil(t, paymentCreateResp.PaymentID)
	assert.NotNil(t, paymentCreateResp.Status)
	paymentID := paymentCreateResp.PaymentID

	if useLinkToken {
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
	} else {
		paymentTokenCreateResp, err := testClient.CreatePaymentToken(paymentID)
		assert.Nil(t, err)
		assert.NotNil(t, paymentTokenCreateResp.PaymentToken)
		assert.NotNil(t, paymentTokenCreateResp.PaymentTokenExpirationTime)
	}

	paymentGetResp, err := testClient.GetPayment(paymentID)
	assert.Nil(t, err)
	assert.NotNil(t, paymentGetResp.PaymentID)
	assert.NotNil(t, paymentGetResp.Reference)
	assert.NotNil(t, paymentGetResp.Amount)
	assert.Nil(t, paymentGetResp.Schedule)
	assert.NotNil(t, paymentGetResp.Status)
	assert.NotNil(t, paymentGetResp.LastStatusUpdate)
	assert.NotNil(t, paymentGetResp.RecipientID)

	// Verify that we can create a standing order
	currentDate := time.Now()
	startDate := currentDate.Add(7 * 24 * time.Hour).Format("2006-01-02")
	standingOrderPaymentCreateResp, err := testClient.CreatePayment(
		recipientID,
		"TestPayment",
		PaymentAmount{
			Currency: "GBP",
			Value:    100.0,
		},
		&PaymentSchedule{
			Interval:             "MONTHLY",
			IntervalExecutionDay: 1,
			StartDate:            startDate,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, standingOrderPaymentCreateResp.PaymentID)
	assert.NotNil(t, standingOrderPaymentCreateResp.Status)
	standingOrderPaymentID := standingOrderPaymentCreateResp.PaymentID

	staindgOrderPaymentGetResp, err := testClient.GetPayment(standingOrderPaymentID)
	assert.Nil(t, err)
	assert.NotNil(t, staindgOrderPaymentGetResp.PaymentID)
	assert.NotNil(t, staindgOrderPaymentGetResp.Reference)
	assert.NotNil(t, staindgOrderPaymentGetResp.Amount)
	assert.NotNil(t, staindgOrderPaymentGetResp.Schedule)
	assert.NotNil(t, staindgOrderPaymentGetResp.Status)
	assert.NotNil(t, staindgOrderPaymentGetResp.LastStatusUpdate)
	assert.NotNil(t, staindgOrderPaymentGetResp.RecipientID)

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
