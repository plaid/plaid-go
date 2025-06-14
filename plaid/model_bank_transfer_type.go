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

// BankTransferType The type of bank transfer. This will be either `debit` or `credit`.  A `debit` indicates a transfer of money into the origination account; a `credit` indicates a transfer of money out of the origination account.
type BankTransferType string

var _ = fmt.Printf

// List of BankTransferType
const (
	BANKTRANSFERTYPE_DEBIT BankTransferType = "debit"
	BANKTRANSFERTYPE_CREDIT BankTransferType = "credit"
)

var allowedBankTransferTypeEnumValues = []BankTransferType{
	"debit",
	"credit",
}

func (v *BankTransferType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := BankTransferType(value)


	*v = enumTypeValue
	return nil
}

// NewBankTransferTypeFromValue returns a pointer to a valid BankTransferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBankTransferTypeFromValue(v string) (*BankTransferType, error) {
	ev := BankTransferType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankTransferType) IsValid() bool {
	for _, existing := range allowedBankTransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BankTransferType value
func (v BankTransferType) Ptr() *BankTransferType {
	return &v
}

type NullableBankTransferType struct {
	value *BankTransferType
	isSet bool
}

func (v NullableBankTransferType) Get() *BankTransferType {
	return v.value
}

func (v *NullableBankTransferType) Set(val *BankTransferType) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferType) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferType(val *BankTransferType) *NullableBankTransferType {
	return &NullableBankTransferType{value: val, isSet: true}
}

func (v NullableBankTransferType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

