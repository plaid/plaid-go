package plaid

import (
	"fmt"
)

type Error struct {
	APIResponse

	// List of all errors: https://github.com/plaid/support/blob/master/errors.md
	ErrorType      string `json:"error_type"`
	ErrorCode      string `json:"error_code"`
	ErrorMessage   string `json:"error_message"`
	DisplayMessage string `json:"display_message"`

	// StatusCode needs to be manually set from the response
	StatusCode int
}

func (e Error) Error() string {
	return fmt.Sprintf("Plaid Error - request ID: %s, http status: %d, type: %s, code: %s, message: %s",
		e.RequestID, e.StatusCode, e.ErrorType, e.ErrorCode, e.ErrorMessage)
}
