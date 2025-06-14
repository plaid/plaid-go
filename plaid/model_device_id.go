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

// DeviceId Device Id associated with the device used during the previous link session
type DeviceId struct {
	Type *int32 `json:"type,omitempty"`
	// Identifier for the device
	Id *string `json:"id,omitempty"`
}

// NewDeviceId instantiates a new DeviceId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceId() *DeviceId {
	this := DeviceId{}
	return &this
}

// NewDeviceIdWithDefaults instantiates a new DeviceId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceIdWithDefaults() *DeviceId {
	this := DeviceId{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeviceId) GetType() int32 {
	if o == nil || o.Type == nil {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceId) GetTypeOk() (*int32, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeviceId) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *DeviceId) SetType(v int32) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceId) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceId) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceId) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceId) SetId(v string) {
	o.Id = &v
}

func (o DeviceId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceId struct {
	value *DeviceId
	isSet bool
}

func (v NullableDeviceId) Get() *DeviceId {
	return v.value
}

func (v *NullableDeviceId) Set(val *DeviceId) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceId) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceId(val *DeviceId) *NullableDeviceId {
	return &NullableDeviceId{value: val, isSet: true}
}

func (v NullableDeviceId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


