package plaid

import (
	"context"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetCategories(t *testing.T) {
	categoriesResp, err := testClient.GetCategories(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, categoriesResp.Categories[0].CategoryID, "10000000")
	assert.Equal(t, categoriesResp.Categories[0].Group, "special")
}
