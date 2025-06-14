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

// CreditPayStubPayBasisType The explicit pay basis on the paystub (if present).
type CreditPayStubPayBasisType string

var _ = fmt.Printf

// List of CreditPayStubPayBasisType
const (
	CREDITPAYSTUBPAYBASISTYPE_SALARY CreditPayStubPayBasisType = "SALARY"
	CREDITPAYSTUBPAYBASISTYPE_HOURLY CreditPayStubPayBasisType = "HOURLY"
	CREDITPAYSTUBPAYBASISTYPE_COMMISSION CreditPayStubPayBasisType = "COMMISSION"
)

var allowedCreditPayStubPayBasisTypeEnumValues = []CreditPayStubPayBasisType{
	"SALARY",
	"HOURLY",
	"COMMISSION",
}

func (v *CreditPayStubPayBasisType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CreditPayStubPayBasisType(value)


	*v = enumTypeValue
	return nil
}

// NewCreditPayStubPayBasisTypeFromValue returns a pointer to a valid CreditPayStubPayBasisType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditPayStubPayBasisTypeFromValue(v string) (*CreditPayStubPayBasisType, error) {
	ev := CreditPayStubPayBasisType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditPayStubPayBasisType) IsValid() bool {
	for _, existing := range allowedCreditPayStubPayBasisTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditPayStubPayBasisType value
func (v CreditPayStubPayBasisType) Ptr() *CreditPayStubPayBasisType {
	return &v
}

type NullableCreditPayStubPayBasisType struct {
	value *CreditPayStubPayBasisType
	isSet bool
}

func (v NullableCreditPayStubPayBasisType) Get() *CreditPayStubPayBasisType {
	return v.value
}

func (v *NullableCreditPayStubPayBasisType) Set(val *CreditPayStubPayBasisType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayStubPayBasisType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayStubPayBasisType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayStubPayBasisType(val *CreditPayStubPayBasisType) *NullableCreditPayStubPayBasisType {
	return &NullableCreditPayStubPayBasisType{value: val, isSet: true}
}

func (v NullableCreditPayStubPayBasisType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayStubPayBasisType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

