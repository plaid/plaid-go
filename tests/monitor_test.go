package tests

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/plaid/plaid-go/v40/plaid"
	"github.com/stretchr/testify/assert"
)

const (
	WATCHLIST_PROGRAM_ID = "prg_egdF5fGjY8fWqk"
	LEGAL_NAME           = "Domingo Huang"
	DOCUMENT_NUMBER      = "123456789"
	YYYYMMDD             = "2006-01-02"
)

func TestMonitor(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()
	CLIENT_USER_ID := "monitor-user-" + strconv.FormatInt(time.Now().Unix(), 10)
	country := "US"
	dateOfBirth := time.Now().UTC().Format(YYYYMMDD)
	searchTerms := plaid.NewWatchlistScreeningRequestSearchTerms(WATCHLIST_PROGRAM_ID, LEGAL_NAME)
	searchTerms.SetCountry(country)
	searchTerms.SetDateOfBirth(dateOfBirth)
	searchTerms.SetDocumentNumber(DOCUMENT_NUMBER)

	// POST /watchlist_screening/individual/create
	createRequest := plaid.NewWatchlistScreeningIndividualCreateRequest(*searchTerms)
	createRequest.SetClientUserId(CLIENT_USER_ID)
	createResponse, _, err := testClient.PlaidApi.WatchlistScreeningIndividualCreate(ctx).WatchlistScreeningIndividualCreateRequest(
		*createRequest,
	).Execute()
	assert.NoError(t, err)
	assert.Equal(t, LEGAL_NAME, createResponse.SearchTerms.LegalName, "Request name should match response")
	assert.Equal(t, DOCUMENT_NUMBER, *createResponse.SearchTerms.DocumentNumber.Get(), "Request document number should match response")
	assert.Equal(t, dateOfBirth, *createResponse.SearchTerms.DateOfBirth.Get(), "Request dob should match response")
	assert.Equal(t, country, *createResponse.SearchTerms.Country.Get(), "Request country should match response")

	// POST /watchlist_screening/individual/update
	updateRequest := plaid.NewWatchlistScreeningIndividualUpdateRequest(createResponse.Id)
	updateRequest.SetStatus(plaid.WatchlistScreeningStatus("pending_review"))
	updateResponse, _, err := testClient.PlaidApi.WatchlistScreeningIndividualUpdate(ctx).WatchlistScreeningIndividualUpdateRequest(
		*updateRequest,
	).Execute()
	assert.NoError(t, err)
	assert.Equal(t, plaid.WatchlistScreeningStatus("pending_review"), updateResponse.Status, "Request status should match response")

	// POST /watchlist_screening/individual/list
	listRequest := plaid.NewWatchlistScreeningIndividualListRequest(WATCHLIST_PROGRAM_ID)
	listRequest.SetClientUserId(CLIENT_USER_ID)
	listResponse, _, err := testClient.PlaidApi.WatchlistScreeningIndividualList(ctx).WatchlistScreeningIndividualListRequest(
		*listRequest,
	).Execute()
	assert.NoError(t, err)
	assert.Len(t, listResponse.WatchlistScreenings, 1)

	// POST /watchlist_screening/individual/get
	getRequest := plaid.NewWatchlistScreeningIndividualGetRequest(listResponse.WatchlistScreenings[0].Id)
	getResponse, _, err := testClient.PlaidApi.WatchlistScreeningIndividualGet(ctx).WatchlistScreeningIndividualGetRequest(
		*getRequest,
	).Execute()
	assert.NoError(t, err)
	assert.Equal(t, LEGAL_NAME, getResponse.SearchTerms.LegalName, "Response legal name should match legal name used in prior create request")
}
