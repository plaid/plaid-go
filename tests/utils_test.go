package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/plaid/plaid-go/v32/plaid"
)

// TestJsonUnmarshal tests the case where a null response unmarshalled to a Nullable type
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
	// This should also work for non-String types
	assert.False(t, testValue.Lat.IsSet())
	assert.Empty(t, testValue.Lat.Get())
	// Field not present should not be set and contain an empty value
	assert.False(t, testValue.Region.IsSet())
	assert.Empty(t, testValue.Region.Get())
}

func TestMakeGenericOpenAPIError(t *testing.T) {
	plaidError := *plaid.NewPlaidError(plaid.PLAIDERRORTYPE_ITEM_ERROR, "PRODUCT_NOT_READY", "", plaid.NullableString{})
	genericOpenAPIError := plaid.MakeGenericOpenAPIError([]byte{}, "400 Bad Request", plaidError)

	derivedPlaidError, err := plaid.ToPlaidError(genericOpenAPIError)
	assert.Nil(t, err)
	assert.Equal(t, plaidError, derivedPlaidError)
}
