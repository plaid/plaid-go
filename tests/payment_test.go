package tests

import (
	"context"
	"testing"
	"time"

	"github.com/plaid/plaid-go/plaid"
	"github.com/stretchr/testify/assert"
)

func TestPaymentRecipientGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	iban := "GB33BUKB20201555555555"
	request := plaid.NewPaymentInitiationRecipientCreateRequest("John Doe")
	request.SetIban(iban)
	request.SetAddress(*plaid.NewPaymentInitiationAddress([]string{"Street Name 999"},
		"City",
		"99999",
		"GB",
	))
	paymentRecipientCreateResp, _, err := testClient.PlaidApi.PaymentInitiationRecipientCreate(ctx).PaymentInitiationRecipientCreateRequest(*request).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, paymentRecipientCreateResp.RecipientId)
	recipientID := paymentRecipientCreateResp.RecipientId

	recipientGetRequest := plaid.NewPaymentInitiationRecipientGetRequest(recipientID)
	paymentRecipientGetResp, _, err := testClient.PlaidApi.PaymentInitiationRecipientGet(ctx).PaymentInitiationRecipientGetRequest(*recipientGetRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, paymentRecipientGetResp.RecipientId)
	assert.NotNil(t, paymentRecipientGetResp.Name)
	assert.NotNil(t, paymentRecipientGetResp.Iban)
	assert.NotNil(t, paymentRecipientGetResp.Address)

	commonPaymentTestFlows(t, ctx, testClient, recipientID, true, false)
}

func TestPaymentWithBacs(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	request := plaid.NewPaymentInitiationRecipientCreateRequest("John Doe")
	request.SetBacs(plaid.RecipientBACSNullable{
		Account:  plaid.PtrString("26207729"),
		SortCode: plaid.PtrString("560029"),
	})
	request.SetAddress(plaid.PaymentInitiationAddress{
		Street:     []string{"Street Name 999"},
		City:       "City",
		PostalCode: "99999",
		Country:    "GB",
	})

	paymentRecipientCreateResp, _, err := testClient.PlaidApi.PaymentInitiationRecipientCreate(ctx).PaymentInitiationRecipientCreateRequest(*request).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, paymentRecipientCreateResp.RecipientId)
	recipientID := paymentRecipientCreateResp.RecipientId

	recipientGetRequest := plaid.NewPaymentInitiationRecipientGetRequest(recipientID)
	paymentRecipientGetResp, _, err := testClient.PlaidApi.PaymentInitiationRecipientGet(ctx).PaymentInitiationRecipientGetRequest(*recipientGetRequest).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, paymentRecipientGetResp.RecipientId)
	assert.NotEmpty(t, paymentRecipientGetResp.Name)
	assert.True(t, paymentRecipientGetResp.Bacs.IsSet())
	assert.True(t, paymentRecipientGetResp.Address.IsSet())

	commonPaymentTestFlows(t, ctx, testClient, recipientID, true, false)
}

func TestPaymentWithBacsAndIban(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	iban := "GB33BUKB20201555555555"

	request := plaid.NewPaymentInitiationRecipientCreateRequest("John Doe")
	request.SetBacs(plaid.RecipientBACSNullable{
		Account:  plaid.PtrString("26207729"),
		SortCode: plaid.PtrString("560029"),
	})
	request.SetAddress(plaid.PaymentInitiationAddress{
		Street:     []string{"Street Name 999"},
		City:       "City",
		PostalCode: "99999",
		Country:    "GB",
	})
	request.SetIban(iban)

	paymentRecipientCreateResp, _, err := testClient.PlaidApi.PaymentInitiationRecipientCreate(ctx).PaymentInitiationRecipientCreateRequest(*request).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, paymentRecipientCreateResp.RecipientId)
	recipientID := paymentRecipientCreateResp.RecipientId

	recipientGetRequest := plaid.NewPaymentInitiationRecipientGetRequest(recipientID)
	paymentRecipientGetResp, _, err := testClient.PlaidApi.PaymentInitiationRecipientGet(ctx).PaymentInitiationRecipientGetRequest(*recipientGetRequest).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, paymentRecipientGetResp.RecipientId)
	assert.NotEmpty(t, paymentRecipientGetResp.Name)
	assert.True(t, paymentRecipientGetResp.Iban.IsSet())
	assert.True(t, paymentRecipientGetResp.Bacs.IsSet())
	assert.True(t, paymentRecipientGetResp.Address.IsSet())

	// testing common flows with the link_token
	commonPaymentTestFlows(t, ctx, testClient, recipientID, true, false)

	// testing common flows with the legacy payment_token
	commonPaymentTestFlows(t, ctx, testClient, recipientID, false, false)

	// testing common flows with options
	commonPaymentTestFlows(t, ctx, testClient, recipientID, false, true)
}

func commonPaymentTestFlows(t *testing.T, ctx context.Context, testClient *plaid.APIClient, recipientID string, useLinkToken bool, useOptions bool) {
	paymentRecipientListResp, _, err := testClient.PlaidApi.PaymentInitiationRecipientList(ctx).PaymentInitiationRecipientListRequest(*plaid.NewPaymentInitiationRecipientListRequest()).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, paymentRecipientListResp.Recipients)

	request := plaid.NewPaymentInitiationPaymentCreateRequest(
		recipientID,
		"TestPayment",
		*plaid.NewPaymentAmount("GBP", 100.0),
	)

	if useOptions == true {
		requestRefundDetails := true
		options := *plaid.NewExternalPaymentOptions()
		options.SetBacs(plaid.PaymentInitiationOptionalRestrictionBacs{
			Account:  plaid.PtrString("1234567890"),
			SortCode: plaid.PtrString("000000"),
		})
		options.SetRequestRefundDetails(requestRefundDetails)

		request.SetOptions(options)
	}
	paymentCreateResp, _, err := testClient.PlaidApi.PaymentInitiationPaymentCreate(ctx).PaymentInitiationPaymentCreateRequest(*request).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, paymentCreateResp.PaymentId)
	assert.NotEmpty(t, paymentCreateResp.Status)
	paymentID := paymentCreateResp.PaymentId

	if useLinkToken {
		user := plaid.NewLinkTokenCreateRequestUser(time.Now().String())
		request := plaid.LinkTokenCreateRequest{
			ClientName:        "Plaid Test",
			Products:          &[]plaid.Products{plaid.PRODUCTS_PAYMENT_INITIATION},
			CountryCodes:      []plaid.CountryCode{plaid.COUNTRYCODE_US},
			Language:          "en",
			PaymentInitiation: plaid.NewLinkTokenCreateRequestPaymentInitiation(paymentID),
			User:              *user,
		}
		linkTokenCreateResp, _, err := testClient.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(request).Execute()

		assert.NoError(t, err)
		assert.NotNil(t, linkTokenCreateResp.LinkToken)
		assert.NotNil(t, linkTokenCreateResp.Expiration)
	} else {
		request := plaid.NewPaymentInitiationPaymentTokenCreateRequest(paymentID)
		paymentTokenCreateResp, _, err := testClient.PlaidApi.CreatePaymentToken(ctx).PaymentInitiationPaymentTokenCreateRequest(*request).Execute()
		assert.NoError(t, err)
		assert.NotNil(t, paymentTokenCreateResp.PaymentToken)
		assert.NotNil(t, paymentTokenCreateResp.PaymentTokenExpirationTime)
	}

	paymentGetResp, _, err := testClient.PlaidApi.PaymentInitiationPaymentGet(ctx).PaymentInitiationPaymentGetRequest(*plaid.NewPaymentInitiationPaymentGetRequest(paymentID)).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, paymentGetResp.PaymentId)
	assert.NotEmpty(t, paymentGetResp.Reference)
	assert.NotEmpty(t, paymentGetResp.Status)
	assert.NotZero(t, paymentGetResp.LastStatusUpdate)
	assert.NotEmpty(t, paymentGetResp.RecipientId)

	// Verify that we can create a standing order
	currentDate := time.Now()
	startDate := currentDate.Add(7 * 24 * time.Hour).Format("2006-01-02")
	amount := plaid.NewPaymentAmount("GBP", 100.0)

	request = plaid.NewPaymentInitiationPaymentCreateRequest(
		recipientID,
		"TestPayment",
		*amount,
	)
	request.SetSchedule(*plaid.NewExternalPaymentScheduleRequest(
		plaid.PAYMENTSCHEDULEINTERVAL_MONTHLY,
		1,
		startDate,
	))
	standingOrderPaymentCreateResp, _, err := testClient.PlaidApi.PaymentInitiationPaymentCreate(ctx).PaymentInitiationPaymentCreateRequest(*request).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, standingOrderPaymentCreateResp.PaymentId)
	assert.NotEmpty(t, standingOrderPaymentCreateResp.Status)
	standingOrderPaymentID := standingOrderPaymentCreateResp.PaymentId

	standingOrderPaymentGetResp, _, err := testClient.PlaidApi.PaymentInitiationPaymentGet(ctx).PaymentInitiationPaymentGetRequest(*plaid.NewPaymentInitiationPaymentGetRequest(standingOrderPaymentID)).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, standingOrderPaymentGetResp.PaymentId)
	assert.NotEmpty(t, standingOrderPaymentGetResp.Reference)
	assert.True(t, standingOrderPaymentGetResp.Schedule.IsSet())
	assert.NotEmpty(t, standingOrderPaymentGetResp.Status)
	assert.NotZero(t, standingOrderPaymentGetResp.LastStatusUpdate)
	assert.NotEmpty(t, standingOrderPaymentGetResp.RecipientId)

	count := int32(10)
	listRequest := plaid.NewPaymentInitiationPaymentListRequest()
	listRequest.SetCount(count)
	paymentListResp, _, err := testClient.PlaidApi.PaymentInitiationPaymentList(ctx).PaymentInitiationPaymentListRequest(*listRequest).Execute()
	assert.NoError(t, err)
	assert.NotEmpty(t, paymentListResp.Payments)

	cursor := paymentListResp.GetNextCursor()
	if !cursor.IsZero() {
		listRequest = plaid.NewPaymentInitiationPaymentListRequest()
		listRequest.SetCursor(cursor)
		_, _, err := testClient.PlaidApi.PaymentInitiationPaymentList(ctx).PaymentInitiationPaymentListRequest(*listRequest).Execute()
		assert.NoError(t, err)
	}
}
