/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.586.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferPlatformPersonIDNumber ID number of the person
type TransferPlatformPersonIDNumber struct {
	// Value of the person's ID Number. Alpha-numeric, with all formatting characters stripped.
	Value string `json:"value"`
	Type IDNumberType `json:"type"`
}

// NewTransferPlatformPersonIDNumber instantiates a new TransferPlatformPersonIDNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferPlatformPersonIDNumber(value string, type_ IDNumberType) *TransferPlatformPersonIDNumber {
	this := TransferPlatformPersonIDNumber{}
	this.Value = value
	this.Type = type_
	return &this
}

// NewTransferPlatformPersonIDNumberWithDefaults instantiates a new TransferPlatformPersonIDNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferPlatformPersonIDNumberWithDefaults() *TransferPlatformPersonIDNumber {
	this := TransferPlatformPersonIDNumber{}
	return &this
}

// GetValue returns the Value field value
func (o *TransferPlatformPersonIDNumber) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TransferPlatformPersonIDNumber) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TransferPlatformPersonIDNumber) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value
func (o *TransferPlatformPersonIDNumber) GetType() IDNumberType {
	if o == nil {
		var ret IDNumberType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferPlatformPersonIDNumber) GetTypeOk() (*IDNumberType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferPlatformPersonIDNumber) SetType(v IDNumberType) {
	o.Type = v
}

func (o TransferPlatformPersonIDNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTransferPlatformPersonIDNumber struct {
	value *TransferPlatformPersonIDNumber
	isSet bool
}

func (v NullableTransferPlatformPersonIDNumber) Get() *TransferPlatformPersonIDNumber {
	return v.value
}

func (v *NullableTransferPlatformPersonIDNumber) Set(val *TransferPlatformPersonIDNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferPlatformPersonIDNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferPlatformPersonIDNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferPlatformPersonIDNumber(val *TransferPlatformPersonIDNumber) *NullableTransferPlatformPersonIDNumber {
	return &NullableTransferPlatformPersonIDNumber{value: val, isSet: true}
}

func (v NullableTransferPlatformPersonIDNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferPlatformPersonIDNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

