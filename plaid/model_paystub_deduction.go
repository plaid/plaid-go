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
)

// PaystubDeduction Deduction on the paystub
type PaystubDeduction struct {
	// The description of the deduction, as provided on the paystub. For example: `\"401(k)\"`, `\"FICA MED TAX\"`.
	Type NullableString `json:"type"`
	// `true` if the deduction is pre-tax; `false` otherwise.
	IsPretax NullableBool `json:"is_pretax"`
	// The amount of the deduction.
	Total NullableFloat64 `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _PaystubDeduction PaystubDeduction

// NewPaystubDeduction instantiates a new PaystubDeduction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystubDeduction(type_ NullableString, isPretax NullableBool, total NullableFloat64) *PaystubDeduction {
	this := PaystubDeduction{}
	this.Type = type_
	this.IsPretax = isPretax
	this.Total = total
	return &this
}

// NewPaystubDeductionWithDefaults instantiates a new PaystubDeduction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubDeductionWithDefaults() *PaystubDeduction {
	this := PaystubDeduction{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaystubDeduction) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubDeduction) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *PaystubDeduction) SetType(v string) {
	o.Type.Set(&v)
}

// GetIsPretax returns the IsPretax field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *PaystubDeduction) GetIsPretax() bool {
	if o == nil || o.IsPretax.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsPretax.Get()
}

// GetIsPretaxOk returns a tuple with the IsPretax field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubDeduction) GetIsPretaxOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsPretax.Get(), o.IsPretax.IsSet()
}

// SetIsPretax sets field value
func (o *PaystubDeduction) SetIsPretax(v bool) {
	o.IsPretax.Set(&v)
}

// GetTotal returns the Total field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PaystubDeduction) GetTotal() float64 {
	if o == nil || o.Total.Get() == nil {
		var ret float64
		return ret
	}

	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubDeduction) GetTotalOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// SetTotal sets field value
func (o *PaystubDeduction) SetTotal(v float64) {
	o.Total.Set(&v)
}

func (o PaystubDeduction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if true {
		toSerialize["is_pretax"] = o.IsPretax.Get()
	}
	if true {
		toSerialize["total"] = o.Total.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaystubDeduction) UnmarshalJSON(bytes []byte) (err error) {
	varPaystubDeduction := _PaystubDeduction{}

	if err = json.Unmarshal(bytes, &varPaystubDeduction); err == nil {
		*o = PaystubDeduction(varPaystubDeduction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "is_pretax")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaystubDeduction struct {
	value *PaystubDeduction
	isSet bool
}

func (v NullablePaystubDeduction) Get() *PaystubDeduction {
	return v.value
}

func (v *NullablePaystubDeduction) Set(val *PaystubDeduction) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubDeduction) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubDeduction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubDeduction(val *PaystubDeduction) *NullablePaystubDeduction {
	return &NullablePaystubDeduction{value: val, isSet: true}
}

func (v NullablePaystubDeduction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubDeduction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


