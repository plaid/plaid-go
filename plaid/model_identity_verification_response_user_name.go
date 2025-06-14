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

// IdentityVerificationResponseUserName The full name provided by the user. If the user has not submitted their name, this field will be null. Otherwise, both fields are guaranteed to be filled.
type IdentityVerificationResponseUserName struct {
	// A string with at least one non-whitespace character, with a max length of 100 characters.
	GivenName string `json:"given_name"`
	// A string with at least one non-whitespace character, with a max length of 100 characters.
	FamilyName string `json:"family_name"`
	AdditionalProperties map[string]interface{}
}

type _IdentityVerificationResponseUserName IdentityVerificationResponseUserName

// NewIdentityVerificationResponseUserName instantiates a new IdentityVerificationResponseUserName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationResponseUserName(givenName string, familyName string) *IdentityVerificationResponseUserName {
	this := IdentityVerificationResponseUserName{}
	this.GivenName = givenName
	this.FamilyName = familyName
	return &this
}

// NewIdentityVerificationResponseUserNameWithDefaults instantiates a new IdentityVerificationResponseUserName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationResponseUserNameWithDefaults() *IdentityVerificationResponseUserName {
	this := IdentityVerificationResponseUserName{}
	return &this
}

// GetGivenName returns the GivenName field value
func (o *IdentityVerificationResponseUserName) GetGivenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationResponseUserName) GetGivenNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GivenName, true
}

// SetGivenName sets field value
func (o *IdentityVerificationResponseUserName) SetGivenName(v string) {
	o.GivenName = v
}

// GetFamilyName returns the FamilyName field value
func (o *IdentityVerificationResponseUserName) GetFamilyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationResponseUserName) GetFamilyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FamilyName, true
}

// SetFamilyName sets field value
func (o *IdentityVerificationResponseUserName) SetFamilyName(v string) {
	o.FamilyName = v
}

func (o IdentityVerificationResponseUserName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["given_name"] = o.GivenName
	}
	if true {
		toSerialize["family_name"] = o.FamilyName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityVerificationResponseUserName) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityVerificationResponseUserName := _IdentityVerificationResponseUserName{}

	if err = json.Unmarshal(bytes, &varIdentityVerificationResponseUserName); err == nil {
		*o = IdentityVerificationResponseUserName(varIdentityVerificationResponseUserName)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "given_name")
		delete(additionalProperties, "family_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityVerificationResponseUserName struct {
	value *IdentityVerificationResponseUserName
	isSet bool
}

func (v NullableIdentityVerificationResponseUserName) Get() *IdentityVerificationResponseUserName {
	return v.value
}

func (v *NullableIdentityVerificationResponseUserName) Set(val *IdentityVerificationResponseUserName) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationResponseUserName) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationResponseUserName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationResponseUserName(val *IdentityVerificationResponseUserName) *NullableIdentityVerificationResponseUserName {
	return &NullableIdentityVerificationResponseUserName{value: val, isSet: true}
}

func (v NullableIdentityVerificationResponseUserName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationResponseUserName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


