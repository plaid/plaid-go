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

// UserStatedIncomeSourcePayType The pay type - `GROSS`, `NET`, or `UNKNOWN` for a specified income source
type UserStatedIncomeSourcePayType string

var _ = fmt.Printf

// List of UserStatedIncomeSourcePayType
const (
	USERSTATEDINCOMESOURCEPAYTYPE_UNKNOWN UserStatedIncomeSourcePayType = "UNKNOWN"
	USERSTATEDINCOMESOURCEPAYTYPE_GROSS UserStatedIncomeSourcePayType = "GROSS"
	USERSTATEDINCOMESOURCEPAYTYPE_NET UserStatedIncomeSourcePayType = "NET"
)

var allowedUserStatedIncomeSourcePayTypeEnumValues = []UserStatedIncomeSourcePayType{
	"UNKNOWN",
	"GROSS",
	"NET",
}

func (v *UserStatedIncomeSourcePayType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := UserStatedIncomeSourcePayType(value)


	*v = enumTypeValue
	return nil
}

// NewUserStatedIncomeSourcePayTypeFromValue returns a pointer to a valid UserStatedIncomeSourcePayType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserStatedIncomeSourcePayTypeFromValue(v string) (*UserStatedIncomeSourcePayType, error) {
	ev := UserStatedIncomeSourcePayType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserStatedIncomeSourcePayType) IsValid() bool {
	for _, existing := range allowedUserStatedIncomeSourcePayTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserStatedIncomeSourcePayType value
func (v UserStatedIncomeSourcePayType) Ptr() *UserStatedIncomeSourcePayType {
	return &v
}

type NullableUserStatedIncomeSourcePayType struct {
	value *UserStatedIncomeSourcePayType
	isSet bool
}

func (v NullableUserStatedIncomeSourcePayType) Get() *UserStatedIncomeSourcePayType {
	return v.value
}

func (v *NullableUserStatedIncomeSourcePayType) Set(val *UserStatedIncomeSourcePayType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatedIncomeSourcePayType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatedIncomeSourcePayType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatedIncomeSourcePayType(val *UserStatedIncomeSourcePayType) *NullableUserStatedIncomeSourcePayType {
	return &NullableUserStatedIncomeSourcePayType{value: val, isSet: true}
}

func (v NullableUserStatedIncomeSourcePayType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatedIncomeSourcePayType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

