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

// IdentityUpdateTypes The possible types of identity data that may have changed.
type IdentityUpdateTypes string

var _ = fmt.Printf

// List of IdentityUpdateTypes
const (
	IDENTITYUPDATETYPES_PHONES IdentityUpdateTypes = "PHONES"
	IDENTITYUPDATETYPES_ADDRESSES IdentityUpdateTypes = "ADDRESSES"
	IDENTITYUPDATETYPES_EMAILS IdentityUpdateTypes = "EMAILS"
	IDENTITYUPDATETYPES_NAMES IdentityUpdateTypes = "NAMES"
)

var allowedIdentityUpdateTypesEnumValues = []IdentityUpdateTypes{
	"PHONES",
	"ADDRESSES",
	"EMAILS",
	"NAMES",
}

func (v *IdentityUpdateTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := IdentityUpdateTypes(value)


	*v = enumTypeValue
	return nil
}

// NewIdentityUpdateTypesFromValue returns a pointer to a valid IdentityUpdateTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIdentityUpdateTypesFromValue(v string) (*IdentityUpdateTypes, error) {
	ev := IdentityUpdateTypes(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdentityUpdateTypes) IsValid() bool {
	for _, existing := range allowedIdentityUpdateTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IdentityUpdateTypes value
func (v IdentityUpdateTypes) Ptr() *IdentityUpdateTypes {
	return &v
}

type NullableIdentityUpdateTypes struct {
	value *IdentityUpdateTypes
	isSet bool
}

func (v NullableIdentityUpdateTypes) Get() *IdentityUpdateTypes {
	return v.value
}

func (v *NullableIdentityUpdateTypes) Set(val *IdentityUpdateTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUpdateTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUpdateTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUpdateTypes(val *IdentityUpdateTypes) *NullableIdentityUpdateTypes {
	return &NullableIdentityUpdateTypes{value: val, isSet: true}
}

func (v NullableIdentityUpdateTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUpdateTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

