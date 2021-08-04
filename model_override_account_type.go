/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.19.12
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// OverrideAccountType `investment:` Investment account  `credit:` Credit card  `depository:` Depository account  `loan:` Loan account  `other:` Non-specified account type  See the [Account type schema](https://plaid.com/docs/api/accounts#account-type-schema) for a full listing of account types and corresponding subtypes.
type OverrideAccountType string

// List of OverrideAccountType
const (
	OVERRIDEACCOUNTTYPE_INVESTMENT OverrideAccountType = "investment"
	OVERRIDEACCOUNTTYPE_CREDIT OverrideAccountType = "credit"
	OVERRIDEACCOUNTTYPE_DEPOSITORY OverrideAccountType = "depository"
	OVERRIDEACCOUNTTYPE_LOAN OverrideAccountType = "loan"
	OVERRIDEACCOUNTTYPE_OTHER OverrideAccountType = "other"
)

func (v *OverrideAccountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OverrideAccountType(value)
	for _, existing := range []OverrideAccountType{ "investment", "credit", "depository", "loan", "other",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OverrideAccountType", value)
}

// Ptr returns reference to OverrideAccountType value
func (v OverrideAccountType) Ptr() *OverrideAccountType {
	return &v
}

type NullableOverrideAccountType struct {
	value *OverrideAccountType
	isSet bool
}

func (v NullableOverrideAccountType) Get() *OverrideAccountType {
	return v.value
}

func (v *NullableOverrideAccountType) Set(val *OverrideAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableOverrideAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableOverrideAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOverrideAccountType(val *OverrideAccountType) *NullableOverrideAccountType {
	return &NullableOverrideAccountType{value: val, isSet: true}
}

func (v NullableOverrideAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOverrideAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
