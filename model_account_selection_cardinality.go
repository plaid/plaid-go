/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.19.10
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// AccountSelectionCardinality The application requires that accounts be limited to a specific cardinality.  `MULTI_SELECT`: indicates that the user should be allowed to pick multiple accounts. `SINGLE_SELECT`: indicates that the user should be allowed to pick only a single account.  `ALL`: indicates that the user must share all of their accounts and should not be given the opportunity to de-select
type AccountSelectionCardinality string

// List of AccountSelectionCardinality
const (
	ACCOUNTSELECTIONCARDINALITY_SINGLE_SELECT AccountSelectionCardinality = "SINGLE_SELECT"
	ACCOUNTSELECTIONCARDINALITY_MULTI_SELECT AccountSelectionCardinality = "MULTI_SELECT"
	ACCOUNTSELECTIONCARDINALITY_ALL AccountSelectionCardinality = "ALL"
)

func (v *AccountSelectionCardinality) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountSelectionCardinality(value)
	for _, existing := range []AccountSelectionCardinality{ "SINGLE_SELECT", "MULTI_SELECT", "ALL",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountSelectionCardinality", value)
}

// Ptr returns reference to AccountSelectionCardinality value
func (v AccountSelectionCardinality) Ptr() *AccountSelectionCardinality {
	return &v
}

type NullableAccountSelectionCardinality struct {
	value *AccountSelectionCardinality
	isSet bool
}

func (v NullableAccountSelectionCardinality) Get() *AccountSelectionCardinality {
	return v.value
}

func (v *NullableAccountSelectionCardinality) Set(val *AccountSelectionCardinality) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSelectionCardinality) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSelectionCardinality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSelectionCardinality(val *AccountSelectionCardinality) *NullableAccountSelectionCardinality {
	return &NullableAccountSelectionCardinality{value: val, isSet: true}
}

func (v NullableAccountSelectionCardinality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSelectionCardinality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
