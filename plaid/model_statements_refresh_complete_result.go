/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.503.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// StatementsRefreshCompleteResult The result of the statement refresh extraction  `SUCCESS`: The statements were successfully extracted and can be listed via `/statements/list/` and downloaded via `/statements/download/`.  `FAILURE`: The statements failed to be extracted.
type StatementsRefreshCompleteResult string

var _ = fmt.Printf

// List of StatementsRefreshCompleteResult
const (
	STATEMENTSREFRESHCOMPLETERESULT_SUCCESS StatementsRefreshCompleteResult = "SUCCESS"
	STATEMENTSREFRESHCOMPLETERESULT_FAILURE StatementsRefreshCompleteResult = "FAILURE"
)

var allowedStatementsRefreshCompleteResultEnumValues = []StatementsRefreshCompleteResult{
	"SUCCESS",
	"FAILURE",
}

func (v *StatementsRefreshCompleteResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := StatementsRefreshCompleteResult(value)


	*v = enumTypeValue
	return nil
}

// NewStatementsRefreshCompleteResultFromValue returns a pointer to a valid StatementsRefreshCompleteResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatementsRefreshCompleteResultFromValue(v string) (*StatementsRefreshCompleteResult, error) {
	ev := StatementsRefreshCompleteResult(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatementsRefreshCompleteResult) IsValid() bool {
	for _, existing := range allowedStatementsRefreshCompleteResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatementsRefreshCompleteResult value
func (v StatementsRefreshCompleteResult) Ptr() *StatementsRefreshCompleteResult {
	return &v
}

type NullableStatementsRefreshCompleteResult struct {
	value *StatementsRefreshCompleteResult
	isSet bool
}

func (v NullableStatementsRefreshCompleteResult) Get() *StatementsRefreshCompleteResult {
	return v.value
}

func (v *NullableStatementsRefreshCompleteResult) Set(val *StatementsRefreshCompleteResult) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementsRefreshCompleteResult) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementsRefreshCompleteResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementsRefreshCompleteResult(val *StatementsRefreshCompleteResult) *NullableStatementsRefreshCompleteResult {
	return &NullableStatementsRefreshCompleteResult{value: val, isSet: true}
}

func (v NullableStatementsRefreshCompleteResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementsRefreshCompleteResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
