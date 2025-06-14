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

// AccountHolderCategory Indicates the account's categorization as either a personal or a business account. This field is currently in beta; to request access, contact your account manager.
type AccountHolderCategory string

var _ = fmt.Printf

// List of AccountHolderCategory
const (
	ACCOUNTHOLDERCATEGORY_BUSINESS AccountHolderCategory = "business"
	ACCOUNTHOLDERCATEGORY_PERSONAL AccountHolderCategory = "personal"
	ACCOUNTHOLDERCATEGORY_UNRECOGNIZED AccountHolderCategory = "unrecognized"
)

var allowedAccountHolderCategoryEnumValues = []AccountHolderCategory{
	"business",
	"personal",
	"unrecognized",
}

func (v *AccountHolderCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := AccountHolderCategory(value)


	*v = enumTypeValue
	return nil
}

// NewAccountHolderCategoryFromValue returns a pointer to a valid AccountHolderCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountHolderCategoryFromValue(v string) (*AccountHolderCategory, error) {
	ev := AccountHolderCategory(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountHolderCategory) IsValid() bool {
	for _, existing := range allowedAccountHolderCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountHolderCategory value
func (v AccountHolderCategory) Ptr() *AccountHolderCategory {
	return &v
}

type NullableAccountHolderCategory struct {
	value *AccountHolderCategory
	isSet bool
}

func (v NullableAccountHolderCategory) Get() *AccountHolderCategory {
	return v.value
}

func (v *NullableAccountHolderCategory) Set(val *AccountHolderCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderCategory(val *AccountHolderCategory) *NullableAccountHolderCategory {
	return &NullableAccountHolderCategory{value: val, isSet: true}
}

func (v NullableAccountHolderCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

