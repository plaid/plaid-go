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

// Loans A collection of loans that are part of a single deal.
type Loans struct {
	LOAN Loan `json:"LOAN"`
	AdditionalProperties map[string]interface{}
}

type _Loans Loans

// NewLoans instantiates a new Loans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoans(lOAN Loan) *Loans {
	this := Loans{}
	this.LOAN = lOAN
	return &this
}

// NewLoansWithDefaults instantiates a new Loans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoansWithDefaults() *Loans {
	this := Loans{}
	return &this
}

// GetLOAN returns the LOAN field value
func (o *Loans) GetLOAN() Loan {
	if o == nil {
		var ret Loan
		return ret
	}

	return o.LOAN
}

// GetLOANOk returns a tuple with the LOAN field value
// and a boolean to check if the value has been set.
func (o *Loans) GetLOANOk() (*Loan, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LOAN, true
}

// SetLOAN sets field value
func (o *Loans) SetLOAN(v Loan) {
	o.LOAN = v
}

func (o Loans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LOAN"] = o.LOAN
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Loans) UnmarshalJSON(bytes []byte) (err error) {
	varLoans := _Loans{}

	if err = json.Unmarshal(bytes, &varLoans); err == nil {
		*o = Loans(varLoans)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LOAN")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoans struct {
	value *Loans
	isSet bool
}

func (v NullableLoans) Get() *Loans {
	return v.value
}

func (v *NullableLoans) Set(val *Loans) {
	v.value = val
	v.isSet = true
}

func (v NullableLoans) IsSet() bool {
	return v.isSet
}

func (v *NullableLoans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoans(val *Loans) *NullableLoans {
	return &NullableLoans{value: val, isSet: true}
}

func (v NullableLoans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


