/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.503.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IdentityVerificationAutofillUserData User information that was autofilled. All this information should be confirmed by the user before using.
type IdentityVerificationAutofillUserData struct {
	Name NullableIdentityVerificationResponseUserName `json:"name"`
	Address NullableIdentityVerificationAutofillAddress `json:"address"`
	IdNumber NullableUserIDNumber `json:"id_number"`
	AdditionalProperties map[string]interface{}
}

type _IdentityVerificationAutofillUserData IdentityVerificationAutofillUserData

// NewIdentityVerificationAutofillUserData instantiates a new IdentityVerificationAutofillUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationAutofillUserData(name NullableIdentityVerificationResponseUserName, address NullableIdentityVerificationAutofillAddress, idNumber NullableUserIDNumber) *IdentityVerificationAutofillUserData {
	this := IdentityVerificationAutofillUserData{}
	this.Name = name
	this.Address = address
	this.IdNumber = idNumber
	return &this
}

// NewIdentityVerificationAutofillUserDataWithDefaults instantiates a new IdentityVerificationAutofillUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationAutofillUserDataWithDefaults() *IdentityVerificationAutofillUserData {
	this := IdentityVerificationAutofillUserData{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for IdentityVerificationResponseUserName will be returned
func (o *IdentityVerificationAutofillUserData) GetName() IdentityVerificationResponseUserName {
	if o == nil || o.Name.Get() == nil {
		var ret IdentityVerificationResponseUserName
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationAutofillUserData) GetNameOk() (*IdentityVerificationResponseUserName, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *IdentityVerificationAutofillUserData) SetName(v IdentityVerificationResponseUserName) {
	o.Name.Set(&v)
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for IdentityVerificationAutofillAddress will be returned
func (o *IdentityVerificationAutofillUserData) GetAddress() IdentityVerificationAutofillAddress {
	if o == nil || o.Address.Get() == nil {
		var ret IdentityVerificationAutofillAddress
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationAutofillUserData) GetAddressOk() (*IdentityVerificationAutofillAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *IdentityVerificationAutofillUserData) SetAddress(v IdentityVerificationAutofillAddress) {
	o.Address.Set(&v)
}

// GetIdNumber returns the IdNumber field value
// If the value is explicit nil, the zero value for UserIDNumber will be returned
func (o *IdentityVerificationAutofillUserData) GetIdNumber() UserIDNumber {
	if o == nil || o.IdNumber.Get() == nil {
		var ret UserIDNumber
		return ret
	}

	return *o.IdNumber.Get()
}

// GetIdNumberOk returns a tuple with the IdNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationAutofillUserData) GetIdNumberOk() (*UserIDNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdNumber.Get(), o.IdNumber.IsSet()
}

// SetIdNumber sets field value
func (o *IdentityVerificationAutofillUserData) SetIdNumber(v UserIDNumber) {
	o.IdNumber.Set(&v)
}

func (o IdentityVerificationAutofillUserData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["address"] = o.Address.Get()
	}
	if true {
		toSerialize["id_number"] = o.IdNumber.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityVerificationAutofillUserData) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityVerificationAutofillUserData := _IdentityVerificationAutofillUserData{}

	if err = json.Unmarshal(bytes, &varIdentityVerificationAutofillUserData); err == nil {
		*o = IdentityVerificationAutofillUserData(varIdentityVerificationAutofillUserData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "address")
		delete(additionalProperties, "id_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityVerificationAutofillUserData struct {
	value *IdentityVerificationAutofillUserData
	isSet bool
}

func (v NullableIdentityVerificationAutofillUserData) Get() *IdentityVerificationAutofillUserData {
	return v.value
}

func (v *NullableIdentityVerificationAutofillUserData) Set(val *IdentityVerificationAutofillUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationAutofillUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationAutofillUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationAutofillUserData(val *IdentityVerificationAutofillUserData) *NullableIdentityVerificationAutofillUserData {
	return &NullableIdentityVerificationAutofillUserData{value: val, isSet: true}
}

func (v NullableIdentityVerificationAutofillUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationAutofillUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

