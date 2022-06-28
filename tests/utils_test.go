package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/plaid/plaid-go/v8/plaid"
)

// TestJsonUnmarshal tests the case where a null response unmarshalled to a NullableString
// sets the isSet property to true.
func TestJsonUnmarshal(t *testing.T) {
	var testValue plaid.Location
	testJson := `{
        "address": null,
        "city": "hello",
        "country": null,
        "lat": null,
        "lon": null,
        "postal_code": null,
        "store_number": null
      }`

	json.Unmarshal([]byte(testJson), &testValue)

	// String should unmarshal and set isSet to true
	assert.True(t, testValue.City.IsSet())
	assert.NotEmpty(t, testValue.City.Get())
	// Null should not be set and contain an empty value
	assert.False(t, testValue.Address.IsSet())
	assert.Empty(t, testValue.Address.Get())
	// Field not present should not be set and contain an empty value
	assert.False(t, testValue.Region.IsSet())
	assert.Empty(t, testValue.Region.Get())
}
