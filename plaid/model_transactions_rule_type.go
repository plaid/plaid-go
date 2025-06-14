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

// TransactionsRuleType Transaction rule's match type. For TRANSACTION_ID field, EXACT_MATCH is available. Matches are case sensitive. 
type TransactionsRuleType string

var _ = fmt.Printf

// List of TransactionsRuleType
const (
	TRANSACTIONSRULETYPE_EXACT_MATCH TransactionsRuleType = "EXACT_MATCH"
	TRANSACTIONSRULETYPE_SUBSTRING_MATCH TransactionsRuleType = "SUBSTRING_MATCH"
)

var allowedTransactionsRuleTypeEnumValues = []TransactionsRuleType{
	"EXACT_MATCH",
	"SUBSTRING_MATCH",
}

func (v *TransactionsRuleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransactionsRuleType(value)


	*v = enumTypeValue
	return nil
}

// NewTransactionsRuleTypeFromValue returns a pointer to a valid TransactionsRuleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionsRuleTypeFromValue(v string) (*TransactionsRuleType, error) {
	ev := TransactionsRuleType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionsRuleType) IsValid() bool {
	for _, existing := range allowedTransactionsRuleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionsRuleType value
func (v TransactionsRuleType) Ptr() *TransactionsRuleType {
	return &v
}

type NullableTransactionsRuleType struct {
	value *TransactionsRuleType
	isSet bool
}

func (v NullableTransactionsRuleType) Get() *TransactionsRuleType {
	return v.value
}

func (v *NullableTransactionsRuleType) Set(val *TransactionsRuleType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRuleType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRuleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRuleType(val *TransactionsRuleType) *NullableTransactionsRuleType {
	return &NullableTransactionsRuleType{value: val, isSet: true}
}

func (v NullableTransactionsRuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRuleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

