package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v10/plaid"
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
			Description:     "DDA PURCHASE *XXXX XXXXXXXX FAMILY DOLLAR",
			Amount:          10.00,
			IsoCurrencyCode: "USD",
			Location: &plaid.ClientProvidedTransactionLocation{
				City:   plaid.PtrString("Philadelphia"),
				Region: plaid.PtrString("PA"),
			},
			Direction: *outflowDirection,
		},
		{
			Id:              "2",
			Description:     "PURCHASE JUNIATA SUPE PHILADELPHIA PA CARDXXXX",
			Amount:          10.00,
			IsoCurrencyCode: "USD",
			Direction:       *outflowDirection,
		},
		{
			Id:              "3",
			Description:     "DEBIT CARD PURCHASE",
			Amount:          10.00,
			IsoCurrencyCode: "USD",
			Direction:       *outflowDirection,
		},
		{
			Id:              "4",
			Description:     "DIRECT DEP* UBER.COM",
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
