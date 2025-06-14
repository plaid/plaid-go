/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.645.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// CraBankIncomeCompleteResult The result of the bank income report generation  `SUCCESS`: The bank income report was successfully generated and can be retrieved via `/cra/bank_income/get`.  `FAILURE`: The bank income report failed to be generated
type CraBankIncomeCompleteResult string

var _ = fmt.Printf

// List of CraBankIncomeCompleteResult
const (
	CRABANKINCOMECOMPLETERESULT_SUCCESS CraBankIncomeCompleteResult = "SUCCESS"
	CRABANKINCOMECOMPLETERESULT_FAILURE CraBankIncomeCompleteResult = "FAILURE"
)

var allowedCraBankIncomeCompleteResultEnumValues = []CraBankIncomeCompleteResult{
	"SUCCESS",
	"FAILURE",
}

func (v *CraBankIncomeCompleteResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CraBankIncomeCompleteResult(value)


	*v = enumTypeValue
	return nil
}

// NewCraBankIncomeCompleteResultFromValue returns a pointer to a valid CraBankIncomeCompleteResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCraBankIncomeCompleteResultFromValue(v string) (*CraBankIncomeCompleteResult, error) {
	ev := CraBankIncomeCompleteResult(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CraBankIncomeCompleteResult) IsValid() bool {
	for _, existing := range allowedCraBankIncomeCompleteResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CraBankIncomeCompleteResult value
func (v CraBankIncomeCompleteResult) Ptr() *CraBankIncomeCompleteResult {
	return &v
}

type NullableCraBankIncomeCompleteResult struct {
	value *CraBankIncomeCompleteResult
	isSet bool
}

func (v NullableCraBankIncomeCompleteResult) Get() *CraBankIncomeCompleteResult {
	return v.value
}

func (v *NullableCraBankIncomeCompleteResult) Set(val *CraBankIncomeCompleteResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCraBankIncomeCompleteResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCraBankIncomeCompleteResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraBankIncomeCompleteResult(val *CraBankIncomeCompleteResult) *NullableCraBankIncomeCompleteResult {
	return &NullableCraBankIncomeCompleteResult{value: val, isSet: true}
}

func (v NullableCraBankIncomeCompleteResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraBankIncomeCompleteResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

