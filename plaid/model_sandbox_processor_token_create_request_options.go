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

// SandboxProcessorTokenCreateRequestOptions An optional set of options to be used when configuring the Item. If specified, must not be `null`.
type SandboxProcessorTokenCreateRequestOptions struct {
	// Test username to use for the creation of the Sandbox Item. Default value is `user_good`.
	OverrideUsername NullableString `json:"override_username,omitempty"`
	// Test password to use for the creation of the Sandbox Item. Default value is `pass_good`.
	OverridePassword NullableString `json:"override_password,omitempty"`
}

// NewSandboxProcessorTokenCreateRequestOptions instantiates a new SandboxProcessorTokenCreateRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxProcessorTokenCreateRequestOptions() *SandboxProcessorTokenCreateRequestOptions {
	this := SandboxProcessorTokenCreateRequestOptions{}
	var overrideUsername string = "user_good"
	this.OverrideUsername = *NewNullableString(&overrideUsername)
	var overridePassword string = "pass_good"
	this.OverridePassword = *NewNullableString(&overridePassword)
	return &this
}

// NewSandboxProcessorTokenCreateRequestOptionsWithDefaults instantiates a new SandboxProcessorTokenCreateRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxProcessorTokenCreateRequestOptionsWithDefaults() *SandboxProcessorTokenCreateRequestOptions {
	this := SandboxProcessorTokenCreateRequestOptions{}
	var overrideUsername string = "user_good"
	this.OverrideUsername = *NewNullableString(&overrideUsername)
	var overridePassword string = "pass_good"
	this.OverridePassword = *NewNullableString(&overridePassword)
	return &this
}

// GetOverrideUsername returns the OverrideUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxProcessorTokenCreateRequestOptions) GetOverrideUsername() string {
	if o == nil || o.OverrideUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideUsername.Get()
}

// GetOverrideUsernameOk returns a tuple with the OverrideUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxProcessorTokenCreateRequestOptions) GetOverrideUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverrideUsername.Get(), o.OverrideUsername.IsSet()
}

// HasOverrideUsername returns a boolean if a field has been set.
func (o *SandboxProcessorTokenCreateRequestOptions) HasOverrideUsername() bool {
	if o != nil && o.OverrideUsername.IsSet() {
		return true
	}

	return false
}

// SetOverrideUsername gets a reference to the given NullableString and assigns it to the OverrideUsername field.
func (o *SandboxProcessorTokenCreateRequestOptions) SetOverrideUsername(v string) {
	o.OverrideUsername.Set(&v)
}
// SetOverrideUsernameNil sets the value for OverrideUsername to be an explicit nil
func (o *SandboxProcessorTokenCreateRequestOptions) SetOverrideUsernameNil() {
	o.OverrideUsername.Set(nil)
}

// UnsetOverrideUsername ensures that no value is present for OverrideUsername, not even an explicit nil
func (o *SandboxProcessorTokenCreateRequestOptions) UnsetOverrideUsername() {
	o.OverrideUsername.Unset()
}

// GetOverridePassword returns the OverridePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxProcessorTokenCreateRequestOptions) GetOverridePassword() string {
	if o == nil || o.OverridePassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverridePassword.Get()
}

// GetOverridePasswordOk returns a tuple with the OverridePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxProcessorTokenCreateRequestOptions) GetOverridePasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverridePassword.Get(), o.OverridePassword.IsSet()
}

// HasOverridePassword returns a boolean if a field has been set.
func (o *SandboxProcessorTokenCreateRequestOptions) HasOverridePassword() bool {
	if o != nil && o.OverridePassword.IsSet() {
		return true
	}

	return false
}

// SetOverridePassword gets a reference to the given NullableString and assigns it to the OverridePassword field.
func (o *SandboxProcessorTokenCreateRequestOptions) SetOverridePassword(v string) {
	o.OverridePassword.Set(&v)
}
// SetOverridePasswordNil sets the value for OverridePassword to be an explicit nil
func (o *SandboxProcessorTokenCreateRequestOptions) SetOverridePasswordNil() {
	o.OverridePassword.Set(nil)
}

// UnsetOverridePassword ensures that no value is present for OverridePassword, not even an explicit nil
func (o *SandboxProcessorTokenCreateRequestOptions) UnsetOverridePassword() {
	o.OverridePassword.Unset()
}

func (o SandboxProcessorTokenCreateRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OverrideUsername.IsSet() {
		toSerialize["override_username"] = o.OverrideUsername.Get()
	}
	if o.OverridePassword.IsSet() {
		toSerialize["override_password"] = o.OverridePassword.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxProcessorTokenCreateRequestOptions struct {
	value *SandboxProcessorTokenCreateRequestOptions
	isSet bool
}

func (v NullableSandboxProcessorTokenCreateRequestOptions) Get() *SandboxProcessorTokenCreateRequestOptions {
	return v.value
}

func (v *NullableSandboxProcessorTokenCreateRequestOptions) Set(val *SandboxProcessorTokenCreateRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxProcessorTokenCreateRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxProcessorTokenCreateRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxProcessorTokenCreateRequestOptions(val *SandboxProcessorTokenCreateRequestOptions) *NullableSandboxProcessorTokenCreateRequestOptions {
	return &NullableSandboxProcessorTokenCreateRequestOptions{value: val, isSet: true}
}

func (v NullableSandboxProcessorTokenCreateRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxProcessorTokenCreateRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


