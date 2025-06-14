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

// CraBankIncomeEmployer The object containing employer data.
type CraBankIncomeEmployer struct {
	// The name of the employer.
	Name NullableString `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CraBankIncomeEmployer CraBankIncomeEmployer

// NewCraBankIncomeEmployer instantiates a new CraBankIncomeEmployer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraBankIncomeEmployer(name NullableString) *CraBankIncomeEmployer {
	this := CraBankIncomeEmployer{}
	this.Name = name
	return &this
}

// NewCraBankIncomeEmployerWithDefaults instantiates a new CraBankIncomeEmployer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraBankIncomeEmployerWithDefaults() *CraBankIncomeEmployer {
	this := CraBankIncomeEmployer{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CraBankIncomeEmployer) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeEmployer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CraBankIncomeEmployer) SetName(v string) {
	o.Name.Set(&v)
}

func (o CraBankIncomeEmployer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraBankIncomeEmployer) UnmarshalJSON(bytes []byte) (err error) {
	varCraBankIncomeEmployer := _CraBankIncomeEmployer{}

	if err = json.Unmarshal(bytes, &varCraBankIncomeEmployer); err == nil {
		*o = CraBankIncomeEmployer(varCraBankIncomeEmployer)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraBankIncomeEmployer struct {
	value *CraBankIncomeEmployer
	isSet bool
}

func (v NullableCraBankIncomeEmployer) Get() *CraBankIncomeEmployer {
	return v.value
}

func (v *NullableCraBankIncomeEmployer) Set(val *CraBankIncomeEmployer) {
	v.value = val
	v.isSet = true
}

func (v NullableCraBankIncomeEmployer) IsSet() bool {
	return v.isSet
}

func (v *NullableCraBankIncomeEmployer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraBankIncomeEmployer(val *CraBankIncomeEmployer) *NullableCraBankIncomeEmployer {
	return &NullableCraBankIncomeEmployer{value: val, isSet: true}
}

func (v NullableCraBankIncomeEmployer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraBankIncomeEmployer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


