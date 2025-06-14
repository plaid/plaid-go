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

// DepositoryFilter A filter to apply to `depository`-type accounts
type DepositoryFilter struct {
	// An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema). 
	AccountSubtypes []DepositoryAccountSubtype `json:"account_subtypes"`
	AdditionalProperties map[string]interface{}
}

type _DepositoryFilter DepositoryFilter

// NewDepositoryFilter instantiates a new DepositoryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositoryFilter(accountSubtypes []DepositoryAccountSubtype) *DepositoryFilter {
	this := DepositoryFilter{}
	this.AccountSubtypes = accountSubtypes
	return &this
}

// NewDepositoryFilterWithDefaults instantiates a new DepositoryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositoryFilterWithDefaults() *DepositoryFilter {
	this := DepositoryFilter{}
	return &this
}

// GetAccountSubtypes returns the AccountSubtypes field value
func (o *DepositoryFilter) GetAccountSubtypes() []DepositoryAccountSubtype {
	if o == nil {
		var ret []DepositoryAccountSubtype
		return ret
	}

	return o.AccountSubtypes
}

// GetAccountSubtypesOk returns a tuple with the AccountSubtypes field value
// and a boolean to check if the value has been set.
func (o *DepositoryFilter) GetAccountSubtypesOk() (*[]DepositoryAccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountSubtypes, true
}

// SetAccountSubtypes sets field value
func (o *DepositoryFilter) SetAccountSubtypes(v []DepositoryAccountSubtype) {
	o.AccountSubtypes = v
}

func (o DepositoryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_subtypes"] = o.AccountSubtypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DepositoryFilter) UnmarshalJSON(bytes []byte) (err error) {
	varDepositoryFilter := _DepositoryFilter{}

	if err = json.Unmarshal(bytes, &varDepositoryFilter); err == nil {
		*o = DepositoryFilter(varDepositoryFilter)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_subtypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositoryFilter struct {
	value *DepositoryFilter
	isSet bool
}

func (v NullableDepositoryFilter) Get() *DepositoryFilter {
	return v.value
}

func (v *NullableDepositoryFilter) Set(val *DepositoryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositoryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositoryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositoryFilter(val *DepositoryFilter) *NullableDepositoryFilter {
	return &NullableDepositoryFilter{value: val, isSet: true}
}

func (v NullableDepositoryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositoryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


