/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.84.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// IncomeVerificationSourceType The types of source income data that users should be able to share
type IncomeVerificationSourceType string

// List of IncomeVerificationSourceType
const (
	INCOMEVERIFICATIONSOURCETYPE_BANK IncomeVerificationSourceType = "bank"
	INCOMEVERIFICATIONSOURCETYPE_PAYROLL IncomeVerificationSourceType = "payroll"
)

var allowedIncomeVerificationSourceTypeEnumValues = []IncomeVerificationSourceType{
	"bank",
	"payroll",
}

func (v *IncomeVerificationSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IncomeVerificationSourceType(value)
	for _, existing := range allowedIncomeVerificationSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IncomeVerificationSourceType", value)
}

// NewIncomeVerificationSourceTypeFromValue returns a pointer to a valid IncomeVerificationSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIncomeVerificationSourceTypeFromValue(v string) (*IncomeVerificationSourceType, error) {
	ev := IncomeVerificationSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IncomeVerificationSourceType: valid values are %v", v, allowedIncomeVerificationSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IncomeVerificationSourceType) IsValid() bool {
	for _, existing := range allowedIncomeVerificationSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IncomeVerificationSourceType value
func (v IncomeVerificationSourceType) Ptr() *IncomeVerificationSourceType {
	return &v
}

type NullableIncomeVerificationSourceType struct {
	value *IncomeVerificationSourceType
	isSet bool
}

func (v NullableIncomeVerificationSourceType) Get() *IncomeVerificationSourceType {
	return v.value
}

func (v *NullableIncomeVerificationSourceType) Set(val *IncomeVerificationSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationSourceType(val *IncomeVerificationSourceType) *NullableIncomeVerificationSourceType {
	return &NullableIncomeVerificationSourceType{value: val, isSet: true}
}

func (v NullableIncomeVerificationSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
