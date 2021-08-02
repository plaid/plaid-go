package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoriesGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	categoriesGetResp, _, err := testClient.PlaidApi.CategoriesGet(ctx).Body(nil).Execute()

	assert.NoError(t, err)
	assert.NotEmpty(t, categoriesGetResp.GetCategories())

	category := categoriesGetResp.GetCategories()[0]
	assert.NotEmpty(t, category.GetCategoryId())
	assert.NotEmpty(t, category.GetGroup())
}
