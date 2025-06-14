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

// BankTransferEventListBankTransferType The type of bank transfer. This will be either `debit` or `credit`.  A `debit` indicates a transfer of money into your origination account; a `credit` indicates a transfer of money out of your origination account.
type BankTransferEventListBankTransferType string

var _ = fmt.Printf

// List of BankTransferEventListBankTransferType
const (
	BANKTRANSFEREVENTLISTBANKTRANSFERTYPE_DEBIT BankTransferEventListBankTransferType = "debit"
	BANKTRANSFEREVENTLISTBANKTRANSFERTYPE_CREDIT BankTransferEventListBankTransferType = "credit"
	BANKTRANSFEREVENTLISTBANKTRANSFERTYPE_NULL BankTransferEventListBankTransferType = "null"
)

var allowedBankTransferEventListBankTransferTypeEnumValues = []BankTransferEventListBankTransferType{
	"debit",
	"credit",
	"null",
}

func (v *BankTransferEventListBankTransferType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := BankTransferEventListBankTransferType(value)


	*v = enumTypeValue
	return nil
}

// NewBankTransferEventListBankTransferTypeFromValue returns a pointer to a valid BankTransferEventListBankTransferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBankTransferEventListBankTransferTypeFromValue(v string) (*BankTransferEventListBankTransferType, error) {
	ev := BankTransferEventListBankTransferType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankTransferEventListBankTransferType) IsValid() bool {
	for _, existing := range allowedBankTransferEventListBankTransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BankTransferEventListBankTransferType value
func (v BankTransferEventListBankTransferType) Ptr() *BankTransferEventListBankTransferType {
	return &v
}

type NullableBankTransferEventListBankTransferType struct {
	value *BankTransferEventListBankTransferType
	isSet bool
}

func (v NullableBankTransferEventListBankTransferType) Get() *BankTransferEventListBankTransferType {
	return v.value
}

func (v *NullableBankTransferEventListBankTransferType) Set(val *BankTransferEventListBankTransferType) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferEventListBankTransferType) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferEventListBankTransferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferEventListBankTransferType(val *BankTransferEventListBankTransferType) *NullableBankTransferEventListBankTransferType {
	return &NullableBankTransferEventListBankTransferType{value: val, isSet: true}
}

func (v NullableBankTransferEventListBankTransferType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferEventListBankTransferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

