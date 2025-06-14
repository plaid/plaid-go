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

// FDXFiAttribute Financial Institution provider-specific attribute
type FDXFiAttribute struct {
	// Name of attribute
	Name string `json:"name"`
	// Value of attribute
	Value string `json:"value"`
}

// NewFDXFiAttribute instantiates a new FDXFiAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFDXFiAttribute(name string, value string) *FDXFiAttribute {
	this := FDXFiAttribute{}
	this.Name = name
	this.Value = value
	return &this
}

// NewFDXFiAttributeWithDefaults instantiates a new FDXFiAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFDXFiAttributeWithDefaults() *FDXFiAttribute {
	this := FDXFiAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *FDXFiAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FDXFiAttribute) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FDXFiAttribute) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *FDXFiAttribute) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FDXFiAttribute) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FDXFiAttribute) SetValue(v string) {
	o.Value = v
}

func (o FDXFiAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFDXFiAttribute struct {
	value *FDXFiAttribute
	isSet bool
}

func (v NullableFDXFiAttribute) Get() *FDXFiAttribute {
	return v.value
}

func (v *NullableFDXFiAttribute) Set(val *FDXFiAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXFiAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXFiAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXFiAttribute(val *FDXFiAttribute) *NullableFDXFiAttribute {
	return &NullableFDXFiAttribute{value: val, isSet: true}
}

func (v NullableFDXFiAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXFiAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


