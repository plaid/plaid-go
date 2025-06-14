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

// SignalDevice Details about the end user's device. When calling `/signal/evaluate` or `/signal/processor/evaluate`, this field is optional, but strongly recommended to increase the accuracy of Signal results.
type SignalDevice struct {
	// The IP address of the device that initiated the transaction
	IpAddress NullableString `json:"ip_address,omitempty"`
	// The user agent of the device that initiated the transaction (e.g. \"Mozilla/5.0\")
	UserAgent NullableString `json:"user_agent,omitempty"`
}

// NewSignalDevice instantiates a new SignalDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalDevice() *SignalDevice {
	this := SignalDevice{}
	return &this
}

// NewSignalDeviceWithDefaults instantiates a new SignalDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalDeviceWithDefaults() *SignalDevice {
	this := SignalDevice{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalDevice) GetIpAddress() string {
	if o == nil || o.IpAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalDevice) GetIpAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *SignalDevice) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *SignalDevice) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *SignalDevice) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *SignalDevice) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalDevice) GetUserAgent() string {
	if o == nil || o.UserAgent.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserAgent.Get()
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalDevice) GetUserAgentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserAgent.Get(), o.UserAgent.IsSet()
}

// HasUserAgent returns a boolean if a field has been set.
func (o *SignalDevice) HasUserAgent() bool {
	if o != nil && o.UserAgent.IsSet() {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given NullableString and assigns it to the UserAgent field.
func (o *SignalDevice) SetUserAgent(v string) {
	o.UserAgent.Set(&v)
}
// SetUserAgentNil sets the value for UserAgent to be an explicit nil
func (o *SignalDevice) SetUserAgentNil() {
	o.UserAgent.Set(nil)
}

// UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
func (o *SignalDevice) UnsetUserAgent() {
	o.UserAgent.Unset()
}

func (o SignalDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddress.IsSet() {
		toSerialize["ip_address"] = o.IpAddress.Get()
	}
	if o.UserAgent.IsSet() {
		toSerialize["user_agent"] = o.UserAgent.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSignalDevice struct {
	value *SignalDevice
	isSet bool
}

func (v NullableSignalDevice) Get() *SignalDevice {
	return v.value
}

func (v *NullableSignalDevice) Set(val *SignalDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalDevice(val *SignalDevice) *NullableSignalDevice {
	return &NullableSignalDevice{value: val, isSet: true}
}

func (v NullableSignalDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


