package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v15/plaid"
	"github.com/stretchr/testify/assert"
)

func TestEnrichGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	outflowDirection, _ := plaid.NewEnrichTransactionDirectionFromValue("OUTFLOW")
	inflowDirection, _ := plaid.NewEnrichTransactionDirectionFromValue("INFLOW")

	sampleTransactionsToEnrich := []plaid.ClientProvidedTransaction{
		{
			Id:              "1",
			Description:     "TST *JETTIES BAGELS",
			Amount:          10.00,
			IsoCurrencyCode: "USD",
			Location: &plaid.ClientProvidedTransactionLocation{
				City:   plaid.PtrString("Ipswich"),
				Region: plaid.PtrString("MA"),
			},
			Direction: *outflowDirection,
		},
		{
			Id:              "2",
			Description:     "AMAZON.COM*MJ3LO9AN2",
			Amount:          10.00,
			IsoCurrencyCode: "USD",
			Direction:       *outflowDirection,
		},
		{
			Id:              "3",
			Description:     "GOOGLE *FRESHBOOKS",
			Amount:          10.00,
			IsoCurrencyCode: "USD",
			Direction:       *outflowDirection,
		},
		{
			Id:              "4",
			Description:     "EARNIN TRANSFER",
			Amount:          100.00,
			IsoCurrencyCode: "USD",
			Direction:       *inflowDirection,
		},
	}

	enrichGetResp, _, err := testClient.PlaidApi.TransactionsEnrich(ctx).TransactionsEnrichGetRequest(
		*plaid.NewTransactionsEnrichGetRequest(
			"depository",
			sampleTransactionsToEnrich,
		),
	).Execute()

	assert.NoError(t, err)

	enrichedTransactions, enrichedTransactionsOk := enrichGetResp.GetEnrichedTransactionsOk()
	assert.True(t, enrichedTransactionsOk)

	assert.Equal(t, len(*enrichedTransactions), len(sampleTransactionsToEnrich))

	for _, item := range *enrichedTransactions {
		assert.NotEmpty(t, item.GetAmount())
		assert.NotEmpty(t, item.GetDescription())
		assert.NotEmpty(t, item.GetDirection())
		assert.NotEmpty(t, item.GetEnrichments())
		assert.NotEmpty(t, item.GetId())
	}
}
