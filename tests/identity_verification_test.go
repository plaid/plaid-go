package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"

	"github.com/plaid/plaid-go/v13/plaid"
	"github.com/stretchr/testify/assert"
)

const TEMPLATE_ID = "flwtmp_aWogUuKsL6NEHU"

func TestIdentityVerification(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()
	CLIENT_USER_ID := "idv-user-" + strconv.FormatInt(time.Now().Unix(), 10)
	user := plaid.NewIdentityVerificationRequestUser(CLIENT_USER_ID)
	user.SetEmailAddress(CLIENT_USER_ID + "@example.com")

	// POST /identity_verification/create
	createRequest := plaid.NewIdentityVerificationCreateRequest(true, TEMPLATE_ID, true, *user)
	createResponse, _, err := testClient.PlaidApi.IdentityVerificationCreate(ctx).IdentityVerificationCreateRequest(
		*createRequest,
	).Execute()
	require.NoError(t, err)
	assert.NotNil(t, createResponse.ShareableUrl)
	assert.Equal(t, plaid.IdentityVerificationStatus("active"), createResponse.Status, "Response status should be active")

	// POST /identity_verification/retry
	retryRequest := plaid.NewIdentityVerificationRetryRequest(CLIENT_USER_ID, TEMPLATE_ID, plaid.Strategy("reset"))
	retryResponse, _, err := testClient.PlaidApi.IdentityVerificationRetry(ctx).IdentityVerificationRetryRequest(
		*retryRequest,
	).Execute()
	require.NoError(t, err)
	assert.NotNil(t, retryResponse.ShareableUrl)
	assert.Equal(t, retryResponse.Status, plaid.IdentityVerificationStatus("active"), "Response status should be active")

	// POST /identity_verification/list
	listRequest := plaid.NewIdentityVerificationListRequest(TEMPLATE_ID, CLIENT_USER_ID)
	listResponse, _, err := testClient.PlaidApi.IdentityVerificationList(ctx).IdentityVerificationListRequest(
		*listRequest,
	).Execute()
	require.NoError(t, err)
	assert.NotNil(t, listResponse.IdentityVerifications[0])
	assert.Equal(t, CLIENT_USER_ID, listResponse.IdentityVerifications[0].ClientUserId, "Client User ID should match create request")

	// POST /identity_verification/get
	getRequest := plaid.NewIdentityVerificationGetRequest(listResponse.IdentityVerifications[0].Id)
	getResponse, _, err := testClient.PlaidApi.IdentityVerificationGet(ctx).IdentityVerificationGetRequest(
		*getRequest,
	).Execute()
	require.NoError(t, err)
	assert.Equal(t, listResponse.IdentityVerifications[0].Id, getResponse.Id, "Requested id should match with response")
}
