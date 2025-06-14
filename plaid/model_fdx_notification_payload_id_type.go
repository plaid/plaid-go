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

// FDXNotificationPayloadIdType Type of entity causing origination of a notification
type FDXNotificationPayloadIdType string

var _ = fmt.Printf

// List of FDXNotificationPayloadIdType
const (
	FDXNOTIFICATIONPAYLOADIDTYPE_ACCOUNT FDXNotificationPayloadIdType = "ACCOUNT"
	FDXNOTIFICATIONPAYLOADIDTYPE_CUSTOMER FDXNotificationPayloadIdType = "CUSTOMER"
	FDXNOTIFICATIONPAYLOADIDTYPE_PARTY FDXNotificationPayloadIdType = "PARTY"
	FDXNOTIFICATIONPAYLOADIDTYPE_MAINTENANCE FDXNotificationPayloadIdType = "MAINTENANCE"
	FDXNOTIFICATIONPAYLOADIDTYPE_CONSENT FDXNotificationPayloadIdType = "CONSENT"
)

var allowedFDXNotificationPayloadIdTypeEnumValues = []FDXNotificationPayloadIdType{
	"ACCOUNT",
	"CUSTOMER",
	"PARTY",
	"MAINTENANCE",
	"CONSENT",
}

func (v *FDXNotificationPayloadIdType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := FDXNotificationPayloadIdType(value)


	*v = enumTypeValue
	return nil
}

// NewFDXNotificationPayloadIdTypeFromValue returns a pointer to a valid FDXNotificationPayloadIdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFDXNotificationPayloadIdTypeFromValue(v string) (*FDXNotificationPayloadIdType, error) {
	ev := FDXNotificationPayloadIdType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FDXNotificationPayloadIdType) IsValid() bool {
	for _, existing := range allowedFDXNotificationPayloadIdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FDXNotificationPayloadIdType value
func (v FDXNotificationPayloadIdType) Ptr() *FDXNotificationPayloadIdType {
	return &v
}

type NullableFDXNotificationPayloadIdType struct {
	value *FDXNotificationPayloadIdType
	isSet bool
}

func (v NullableFDXNotificationPayloadIdType) Get() *FDXNotificationPayloadIdType {
	return v.value
}

func (v *NullableFDXNotificationPayloadIdType) Set(val *FDXNotificationPayloadIdType) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXNotificationPayloadIdType) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXNotificationPayloadIdType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXNotificationPayloadIdType(val *FDXNotificationPayloadIdType) *NullableFDXNotificationPayloadIdType {
	return &NullableFDXNotificationPayloadIdType{value: val, isSet: true}
}

func (v NullableFDXNotificationPayloadIdType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXNotificationPayloadIdType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

