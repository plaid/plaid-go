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

// InvestmentsAuthGetNumbers Identifying information for transferring holdings to an investments account.
type InvestmentsAuthGetNumbers struct {
	Acats *[]NumbersACATS `json:"acats,omitempty"`
	Aton *[]NumbersATON `json:"aton,omitempty"`
	Retirement401k *[]NumbersRetirement401k `json:"retirement_401k,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InvestmentsAuthGetNumbers InvestmentsAuthGetNumbers

// NewInvestmentsAuthGetNumbers instantiates a new InvestmentsAuthGetNumbers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsAuthGetNumbers() *InvestmentsAuthGetNumbers {
	this := InvestmentsAuthGetNumbers{}
	return &this
}

// NewInvestmentsAuthGetNumbersWithDefaults instantiates a new InvestmentsAuthGetNumbers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsAuthGetNumbersWithDefaults() *InvestmentsAuthGetNumbers {
	this := InvestmentsAuthGetNumbers{}
	return &this
}

// GetAcats returns the Acats field value if set, zero value otherwise.
func (o *InvestmentsAuthGetNumbers) GetAcats() []NumbersACATS {
	if o == nil || o.Acats == nil {
		var ret []NumbersACATS
		return ret
	}
	return *o.Acats
}

// GetAcatsOk returns a tuple with the Acats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthGetNumbers) GetAcatsOk() (*[]NumbersACATS, bool) {
	if o == nil || o.Acats == nil {
		return nil, false
	}
	return o.Acats, true
}

// HasAcats returns a boolean if a field has been set.
func (o *InvestmentsAuthGetNumbers) HasAcats() bool {
	if o != nil && o.Acats != nil {
		return true
	}

	return false
}

// SetAcats gets a reference to the given []NumbersACATS and assigns it to the Acats field.
func (o *InvestmentsAuthGetNumbers) SetAcats(v []NumbersACATS) {
	o.Acats = &v
}

// GetAton returns the Aton field value if set, zero value otherwise.
func (o *InvestmentsAuthGetNumbers) GetAton() []NumbersATON {
	if o == nil || o.Aton == nil {
		var ret []NumbersATON
		return ret
	}
	return *o.Aton
}

// GetAtonOk returns a tuple with the Aton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthGetNumbers) GetAtonOk() (*[]NumbersATON, bool) {
	if o == nil || o.Aton == nil {
		return nil, false
	}
	return o.Aton, true
}

// HasAton returns a boolean if a field has been set.
func (o *InvestmentsAuthGetNumbers) HasAton() bool {
	if o != nil && o.Aton != nil {
		return true
	}

	return false
}

// SetAton gets a reference to the given []NumbersATON and assigns it to the Aton field.
func (o *InvestmentsAuthGetNumbers) SetAton(v []NumbersATON) {
	o.Aton = &v
}

// GetRetirement401k returns the Retirement401k field value if set, zero value otherwise.
func (o *InvestmentsAuthGetNumbers) GetRetirement401k() []NumbersRetirement401k {
	if o == nil || o.Retirement401k == nil {
		var ret []NumbersRetirement401k
		return ret
	}
	return *o.Retirement401k
}

// GetRetirement401kOk returns a tuple with the Retirement401k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthGetNumbers) GetRetirement401kOk() (*[]NumbersRetirement401k, bool) {
	if o == nil || o.Retirement401k == nil {
		return nil, false
	}
	return o.Retirement401k, true
}

// HasRetirement401k returns a boolean if a field has been set.
func (o *InvestmentsAuthGetNumbers) HasRetirement401k() bool {
	if o != nil && o.Retirement401k != nil {
		return true
	}

	return false
}

// SetRetirement401k gets a reference to the given []NumbersRetirement401k and assigns it to the Retirement401k field.
func (o *InvestmentsAuthGetNumbers) SetRetirement401k(v []NumbersRetirement401k) {
	o.Retirement401k = &v
}

func (o InvestmentsAuthGetNumbers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Acats != nil {
		toSerialize["acats"] = o.Acats
	}
	if o.Aton != nil {
		toSerialize["aton"] = o.Aton
	}
	if o.Retirement401k != nil {
		toSerialize["retirement_401k"] = o.Retirement401k
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InvestmentsAuthGetNumbers) UnmarshalJSON(bytes []byte) (err error) {
	varInvestmentsAuthGetNumbers := _InvestmentsAuthGetNumbers{}

	if err = json.Unmarshal(bytes, &varInvestmentsAuthGetNumbers); err == nil {
		*o = InvestmentsAuthGetNumbers(varInvestmentsAuthGetNumbers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "acats")
		delete(additionalProperties, "aton")
		delete(additionalProperties, "retirement_401k")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvestmentsAuthGetNumbers struct {
	value *InvestmentsAuthGetNumbers
	isSet bool
}

func (v NullableInvestmentsAuthGetNumbers) Get() *InvestmentsAuthGetNumbers {
	return v.value
}

func (v *NullableInvestmentsAuthGetNumbers) Set(val *InvestmentsAuthGetNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsAuthGetNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsAuthGetNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsAuthGetNumbers(val *InvestmentsAuthGetNumbers) *NullableInvestmentsAuthGetNumbers {
	return &NullableInvestmentsAuthGetNumbers{value: val, isSet: true}
}

func (v NullableInvestmentsAuthGetNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsAuthGetNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


