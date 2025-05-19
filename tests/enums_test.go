package tests

import (
	"testing"

	"github.com/plaid/plaid-go/v35/plaid"
	"github.com/stretchr/testify/assert"
)

// We want to be able to accept new enum values as they're added to the API
// without needing clients to update their client library versions. The API
// will error if it doesn't recognize the enum so we still have validation.
func TestUnknownEnums(t *testing.T) {
	val, err := plaid.NewAccountSelectionCardinalityFromValue("UNKNOWN_VALUE")

	assert.NoError(t, err)
	assert.Equal(t, *val, plaid.AccountSelectionCardinality("UNKNOWN_VALUE"))
}
