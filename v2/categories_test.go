package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetCategories(t *testing.T) {
	categoriesResp, err := testClient.GetCategories()
	assert.Nil(t, err)
	assert.Equal(t, categoriesResp.Categories[0].CategoryID, "10000000")
	assert.Equal(t, categoriesResp.Categories[0].Group, "special")
}
