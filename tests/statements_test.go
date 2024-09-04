package tests

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/plaid/plaid-go/v28/plaid"
)

func TestStatementsFullFlow(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	opts := &plaid.SandboxPublicTokenCreateRequestOptions{}
	opts.Statements.Set(&plaid.SandboxPublicTokenCreateRequestOptionsStatements{
		StartDate: "2023-11-01",
		EndDate:   "2023-12-01",
	})

	accessToken := createSandboxItemWithOptions(
		t,
		ctx,
		testClient,
		FIRST_PLATYPUS_BANK,
		[]plaid.Products{plaid.PRODUCTS_STATEMENTS},
		opts,
	)

	// List statements
	listResponse, err := statementsList(t, ctx, testClient, accessToken)
	require.NoError(t, err)
	require.True(t, len(listResponse.GetAccounts()) > 0)
	require.True(t, len(listResponse.GetAccounts()[0].GetStatements()) > 0)

	// Download statements
	for _, account := range listResponse.GetAccounts() {
		for _, statement := range account.GetStatements() {
			downloadRequest := plaid.NewStatementsDownloadRequest(accessToken, statement.GetStatementId())
			_, httpResponse, err := testClient.PlaidApi.StatementsDownload(ctx).StatementsDownloadRequest(
				*downloadRequest,
			).Execute()
			require.NoError(t, err)

			require.Equal(t, "application/pdf", httpResponse.Header.Get("Content-Type"))
			pdfBytes := getBytesBody(t, httpResponse)
			require.Equal(t, generateSha256Hash(pdfBytes), httpResponse.Header.Get("Plaid-Content-Hash"))
		}
	}

	// Refresh Statements
	startDate, endDate := "2024-01-01", "2024-02-01"
	refreshRequest := plaid.NewStatementsRefreshRequest(accessToken, startDate, endDate)
	refreshResponse, _, err := testClient.PlaidApi.StatementsRefresh(ctx).StatementsRefreshRequest(
		*refreshRequest,
	).Execute()
	require.NoError(t, err)
	require.NotEmpty(t, refreshResponse.GetRequestId())

	listRefreshedResponse, err := statementsList(t, ctx, testClient, accessToken)
	require.NoError(t, err)
	require.True(t, len(listRefreshedResponse.GetAccounts()) > 0)
	require.True(t, len(listRefreshedResponse.GetAccounts()[0].GetStatements()) > 0)
}

func generateSha256Hash(statmentDoc []byte) string {
	hashed := sha256.Sum256(statmentDoc)
	return hex.EncodeToString(hashed[:])
}

// GetBytesBody get the body of an http.response as bytes
func getBytesBody(t *testing.T, resp *http.Response) []byte {
	bytes, err := streamToBytes(resp.Body)
	require.NoError(t, err)
	return bytes
}

// StreamToBytes reads all data from io.Reader and returns bytes
func streamToBytes(stream io.ReadCloser) ([]byte, error) {
	defer func() {
		_ = stream.Close()
	}()

	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(stream); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func statementsList(t *testing.T, ctx context.Context, testClient *plaid.APIClient, accessToken string) (*plaid.StatementsListResponse, error) {
	listRequest := plaid.NewStatementsListRequest(accessToken)
	statementsListResponse, _, err := testClient.PlaidApi.StatementsList(ctx).StatementsListRequest(
		*listRequest,
	).Execute()

	require.NoError(t, err)
	return &statementsListResponse, nil
}
