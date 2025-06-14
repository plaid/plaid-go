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

// EmploymentSourceType The types of source employment data that users should be able to share
type EmploymentSourceType string

var _ = fmt.Printf

// List of EmploymentSourceType
const (
	EMPLOYMENTSOURCETYPE_BANK EmploymentSourceType = "bank"
	EMPLOYMENTSOURCETYPE_PAYROLL EmploymentSourceType = "payroll"
)

var allowedEmploymentSourceTypeEnumValues = []EmploymentSourceType{
	"bank",
	"payroll",
}

func (v *EmploymentSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := EmploymentSourceType(value)


	*v = enumTypeValue
	return nil
}

// NewEmploymentSourceTypeFromValue returns a pointer to a valid EmploymentSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmploymentSourceTypeFromValue(v string) (*EmploymentSourceType, error) {
	ev := EmploymentSourceType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmploymentSourceType) IsValid() bool {
	for _, existing := range allowedEmploymentSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmploymentSourceType value
func (v EmploymentSourceType) Ptr() *EmploymentSourceType {
	return &v
}

type NullableEmploymentSourceType struct {
	value *EmploymentSourceType
	isSet bool
}

func (v NullableEmploymentSourceType) Get() *EmploymentSourceType {
	return v.value
}

func (v *NullableEmploymentSourceType) Set(val *EmploymentSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentSourceType(val *EmploymentSourceType) *NullableEmploymentSourceType {
	return &NullableEmploymentSourceType{value: val, isSet: true}
}

func (v NullableEmploymentSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

