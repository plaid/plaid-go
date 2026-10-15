package plaid_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	plaid "github.com/plaid/plaid-go/v49"
	"github.com/plaid/plaid-go/v49/client"
	"github.com/plaid/plaid-go/v49/option"
	"github.com/plaid/plaid-go/v49/transfer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Plaid's multipart handler compares purpose directly with DUE_DILIGENCE;
// JSON-encoding the enum adds quotes and causes INVALID_FIELD.
func TestTransferDiligenceUploadMultipartFields(t *testing.T) {
	const document = "%PDF-1.7\ndiligence test document"
	var requests int
	httpClient := doerFunc(func(r *http.Request) (*http.Response, error) {
		requests++
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/transfer/diligence/document/upload", r.URL.Path)
		require.NoError(t, r.ParseMultipartForm(1<<20))
		defer func() { require.NoError(t, r.MultipartForm.RemoveAll()) }()
		assert.Equal(t, []string{"originator-test"}, r.MultipartForm.Value["originator_client_id"])
		assert.Equal(t, []string{"DUE_DILIGENCE"}, r.MultipartForm.Value["purpose"])
		file, header, err := r.FormFile("file")
		require.NoError(t, err)
		defer func() { require.NoError(t, file.Close()) }()
		data, err := io.ReadAll(file)
		require.NoError(t, err)
		assert.Equal(t, document, string(data))
		assert.Equal(t, "diligence.pdf", header.Filename)
		assert.Equal(t, "application/pdf", header.Header.Get("Content-Type"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(strings.NewReader(`{"request_id":"upload-test"}`)),
		}, nil
	})
	c := client.NewClient(
		option.WithCustomEndpoint("https://upload.example.test", httpClient),
		option.WithClientID(plaid.String("dummy-client-id")),
		option.WithSecret("dummy-secret"),
	)
	response, err := c.Transfer.Diligence.UploadDocument(context.Background(), &transfer.TransferDiligenceDocumentUploadRequest{
		OriginatorClientID: "originator-test",
		Purpose:            plaid.TransferDocumentPurposeDueDiligence,
		File:               plaid.NewFileParam(strings.NewReader(document), "diligence.pdf", "application/pdf"),
	})
	require.NoError(t, err)
	assert.Equal(t, "upload-test", response.RequestID)
	assert.Equal(t, 1, requests)
}
