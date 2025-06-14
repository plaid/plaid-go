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

// TransferEventListTransferType The type of transfer. This will be either `debit` or `credit`.  A `debit` indicates a transfer of money into your origination account; a `credit` indicates a transfer of money out of your origination account.
type TransferEventListTransferType string

var _ = fmt.Printf

// List of TransferEventListTransferType
const (
	TRANSFEREVENTLISTTRANSFERTYPE_DEBIT TransferEventListTransferType = "debit"
	TRANSFEREVENTLISTTRANSFERTYPE_CREDIT TransferEventListTransferType = "credit"
	TRANSFEREVENTLISTTRANSFERTYPE_NULL TransferEventListTransferType = "null"
)

var allowedTransferEventListTransferTypeEnumValues = []TransferEventListTransferType{
	"debit",
	"credit",
	"null",
}

func (v *TransferEventListTransferType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferEventListTransferType(value)


	*v = enumTypeValue
	return nil
}

// NewTransferEventListTransferTypeFromValue returns a pointer to a valid TransferEventListTransferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferEventListTransferTypeFromValue(v string) (*TransferEventListTransferType, error) {
	ev := TransferEventListTransferType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferEventListTransferType) IsValid() bool {
	for _, existing := range allowedTransferEventListTransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferEventListTransferType value
func (v TransferEventListTransferType) Ptr() *TransferEventListTransferType {
	return &v
}

type NullableTransferEventListTransferType struct {
	value *TransferEventListTransferType
	isSet bool
}

func (v NullableTransferEventListTransferType) Get() *TransferEventListTransferType {
	return v.value
}

func (v *NullableTransferEventListTransferType) Set(val *TransferEventListTransferType) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEventListTransferType) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEventListTransferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEventListTransferType(val *TransferEventListTransferType) *NullableTransferEventListTransferType {
	return &NullableTransferEventListTransferType{value: val, isSet: true}
}

func (v NullableTransferEventListTransferType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEventListTransferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

