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

// PrismProduct The Prism products that can be returned by the Plaid API
type PrismProduct string

var _ = fmt.Printf

// List of PrismProduct
const (
	PRISMPRODUCT_INSIGHTS PrismProduct = "insights"
	PRISMPRODUCT_SCORES PrismProduct = "scores"
)

var allowedPrismProductEnumValues = []PrismProduct{
	"insights",
	"scores",
}

func (v *PrismProduct) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PrismProduct(value)


	*v = enumTypeValue
	return nil
}

// NewPrismProductFromValue returns a pointer to a valid PrismProduct
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrismProductFromValue(v string) (*PrismProduct, error) {
	ev := PrismProduct(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrismProduct) IsValid() bool {
	for _, existing := range allowedPrismProductEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrismProduct value
func (v PrismProduct) Ptr() *PrismProduct {
	return &v
}

type NullablePrismProduct struct {
	value *PrismProduct
	isSet bool
}

func (v NullablePrismProduct) Get() *PrismProduct {
	return v.value
}

func (v *NullablePrismProduct) Set(val *PrismProduct) {
	v.value = val
	v.isSet = true
}

func (v NullablePrismProduct) IsSet() bool {
	return v.isSet
}

func (v *NullablePrismProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrismProduct(val *PrismProduct) *NullablePrismProduct {
	return &NullablePrismProduct{value: val, isSet: true}
}

func (v NullablePrismProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrismProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

