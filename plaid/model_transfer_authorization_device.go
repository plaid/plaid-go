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

// TransferAuthorizationDevice Information about the device being used to initiate the authorization. These fields are not currently incorporated into the risk check.
type TransferAuthorizationDevice struct {
	// The IP address of the device being used to initiate the authorization.
	IpAddress *string `json:"ip_address,omitempty"`
	// The user agent of the device being used to initiate the authorization.
	UserAgent *string `json:"user_agent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferAuthorizationDevice TransferAuthorizationDevice

// NewTransferAuthorizationDevice instantiates a new TransferAuthorizationDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAuthorizationDevice() *TransferAuthorizationDevice {
	this := TransferAuthorizationDevice{}
	return &this
}

// NewTransferAuthorizationDeviceWithDefaults instantiates a new TransferAuthorizationDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAuthorizationDeviceWithDefaults() *TransferAuthorizationDevice {
	this := TransferAuthorizationDevice{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *TransferAuthorizationDevice) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationDevice) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *TransferAuthorizationDevice) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *TransferAuthorizationDevice) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *TransferAuthorizationDevice) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAuthorizationDevice) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *TransferAuthorizationDevice) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *TransferAuthorizationDevice) SetUserAgent(v string) {
	o.UserAgent = &v
}

func (o TransferAuthorizationDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddress != nil {
		toSerialize["ip_address"] = o.IpAddress
	}
	if o.UserAgent != nil {
		toSerialize["user_agent"] = o.UserAgent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferAuthorizationDevice) UnmarshalJSON(bytes []byte) (err error) {
	varTransferAuthorizationDevice := _TransferAuthorizationDevice{}

	if err = json.Unmarshal(bytes, &varTransferAuthorizationDevice); err == nil {
		*o = TransferAuthorizationDevice(varTransferAuthorizationDevice)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ip_address")
		delete(additionalProperties, "user_agent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferAuthorizationDevice struct {
	value *TransferAuthorizationDevice
	isSet bool
}

func (v NullableTransferAuthorizationDevice) Get() *TransferAuthorizationDevice {
	return v.value
}

func (v *NullableTransferAuthorizationDevice) Set(val *TransferAuthorizationDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorizationDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorizationDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorizationDevice(val *TransferAuthorizationDevice) *NullableTransferAuthorizationDevice {
	return &NullableTransferAuthorizationDevice{value: val, isSet: true}
}

func (v NullableTransferAuthorizationDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorizationDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


