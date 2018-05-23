package plaid

// APIResponse is the base struct for all responses from the Plaid API.
type APIResponse struct {
	RequestID string `json:"request_id"`
}
