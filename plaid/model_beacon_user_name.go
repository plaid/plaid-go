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

// BeaconUserName The full name for a given Beacon User.
type BeaconUserName struct {
	// A string with at least one non-whitespace character, with a max length of 100 characters.
	GivenName string `json:"given_name"`
	// A string with at least one non-whitespace character, with a max length of 100 characters.
	FamilyName string `json:"family_name"`
	AdditionalProperties map[string]interface{}
}

type _BeaconUserName BeaconUserName

// NewBeaconUserName instantiates a new BeaconUserName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeaconUserName(givenName string, familyName string) *BeaconUserName {
	this := BeaconUserName{}
	this.GivenName = givenName
	this.FamilyName = familyName
	return &this
}

// NewBeaconUserNameWithDefaults instantiates a new BeaconUserName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeaconUserNameWithDefaults() *BeaconUserName {
	this := BeaconUserName{}
	return &this
}

// GetGivenName returns the GivenName field value
func (o *BeaconUserName) GetGivenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value
// and a boolean to check if the value has been set.
func (o *BeaconUserName) GetGivenNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GivenName, true
}

// SetGivenName sets field value
func (o *BeaconUserName) SetGivenName(v string) {
	o.GivenName = v
}

// GetFamilyName returns the FamilyName field value
func (o *BeaconUserName) GetFamilyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value
// and a boolean to check if the value has been set.
func (o *BeaconUserName) GetFamilyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FamilyName, true
}

// SetFamilyName sets field value
func (o *BeaconUserName) SetFamilyName(v string) {
	o.FamilyName = v
}

func (o BeaconUserName) MarshalJSON() ([]byte, error) {
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

func (o *BeaconUserName) UnmarshalJSON(bytes []byte) (err error) {
	varBeaconUserName := _BeaconUserName{}

	if err = json.Unmarshal(bytes, &varBeaconUserName); err == nil {
		*o = BeaconUserName(varBeaconUserName)
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

type NullableBeaconUserName struct {
	value *BeaconUserName
	isSet bool
}

func (v NullableBeaconUserName) Get() *BeaconUserName {
	return v.value
}

func (v *NullableBeaconUserName) Set(val *BeaconUserName) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconUserName) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconUserName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconUserName(val *BeaconUserName) *NullableBeaconUserName {
	return &NullableBeaconUserName{value: val, isSet: true}
}

func (v NullableBeaconUserName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconUserName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


