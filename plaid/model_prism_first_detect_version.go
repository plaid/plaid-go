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

// PrismFirstDetectVersion The version of Prism FirstDetect. If not specified, will default to v3.
type PrismFirstDetectVersion string

var _ = fmt.Printf

// List of PrismFirstDetectVersion
const (
	PRISMFIRSTDETECTVERSION__3 PrismFirstDetectVersion = "3"
	PRISMFIRSTDETECTVERSION_NULL PrismFirstDetectVersion = "null"
)

var allowedPrismFirstDetectVersionEnumValues = []PrismFirstDetectVersion{
	"3",
	"null",
}

func (v *PrismFirstDetectVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PrismFirstDetectVersion(value)


	*v = enumTypeValue
	return nil
}

// NewPrismFirstDetectVersionFromValue returns a pointer to a valid PrismFirstDetectVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrismFirstDetectVersionFromValue(v string) (*PrismFirstDetectVersion, error) {
	ev := PrismFirstDetectVersion(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrismFirstDetectVersion) IsValid() bool {
	for _, existing := range allowedPrismFirstDetectVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrismFirstDetectVersion value
func (v PrismFirstDetectVersion) Ptr() *PrismFirstDetectVersion {
	return &v
}

type NullablePrismFirstDetectVersion struct {
	value *PrismFirstDetectVersion
	isSet bool
}

func (v NullablePrismFirstDetectVersion) Get() *PrismFirstDetectVersion {
	return v.value
}

func (v *NullablePrismFirstDetectVersion) Set(val *PrismFirstDetectVersion) {
	v.value = val
	v.isSet = true
}

func (v NullablePrismFirstDetectVersion) IsSet() bool {
	return v.isSet
}

func (v *NullablePrismFirstDetectVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrismFirstDetectVersion(val *PrismFirstDetectVersion) *NullablePrismFirstDetectVersion {
	return &NullablePrismFirstDetectVersion{value: val, isSet: true}
}

func (v NullablePrismFirstDetectVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrismFirstDetectVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

