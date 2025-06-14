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

// AddressNullable A physical mailing address.
type AddressNullable struct {
	Data AddressData `json:"data"`
	// When `true`, identifies the address as the primary address on an account.
	Primary *bool `json:"primary,omitempty"`
}

// NewAddressNullable instantiates a new AddressNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressNullable(data AddressData) *AddressNullable {
	this := AddressNullable{}
	this.Data = data
	return &this
}

// NewAddressNullableWithDefaults instantiates a new AddressNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressNullableWithDefaults() *AddressNullable {
	this := AddressNullable{}
	return &this
}

// GetData returns the Data field value
func (o *AddressNullable) GetData() AddressData {
	if o == nil {
		var ret AddressData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddressNullable) GetDataOk() (*AddressData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddressNullable) SetData(v AddressData) {
	o.Data = v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *AddressNullable) GetPrimary() bool {
	if o == nil || o.Primary == nil {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressNullable) GetPrimaryOk() (*bool, bool) {
	if o == nil || o.Primary == nil {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *AddressNullable) HasPrimary() bool {
	if o != nil && o.Primary != nil {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *AddressNullable) SetPrimary(v bool) {
	o.Primary = &v
}

func (o AddressNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Primary != nil {
		toSerialize["primary"] = o.Primary
	}
	return json.Marshal(toSerialize)
}

type NullableAddressNullable struct {
	value *AddressNullable
	isSet bool
}

func (v NullableAddressNullable) Get() *AddressNullable {
	return v.value
}

func (v *NullableAddressNullable) Set(val *AddressNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressNullable(val *AddressNullable) *NullableAddressNullable {
	return &NullableAddressNullable{value: val, isSet: true}
}

func (v NullableAddressNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


