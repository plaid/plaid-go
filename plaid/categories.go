package plaid

import (
	"context"
	"encoding/json"
)

type Category struct {
	CategoryID string   `json:"category_id"`
	Group      string   `json:"group"`
	Hierarchy  []string `json:"hierarchy"`
}

type GetCategoriesResponse struct {
	APIResponse
	Categories []Category `json:"categories"`
}

// GetCategories returns information for all categories.
// See https://plaid.com/docs/api/#category-overview.
func (c *Client) GetCategories(ctx context.Context) (resp GetCategoriesResponse, err error) {
	jsonBody, _ := json.Marshal(nil)

	err = c.Call(ctx, "/categories/get", jsonBody, &resp)
	return resp, err
}
