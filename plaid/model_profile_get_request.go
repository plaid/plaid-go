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

// ProfileGetRequest ProfileGetRequest defines the request schema for `/profile/get`
type ProfileGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The profile token generated by the end user authorization session.
	ProfileToken string `json:"profile_token"`
}

// NewProfileGetRequest instantiates a new ProfileGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileGetRequest(profileToken string) *ProfileGetRequest {
	this := ProfileGetRequest{}
	this.ProfileToken = profileToken
	return &this
}

// NewProfileGetRequestWithDefaults instantiates a new ProfileGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileGetRequestWithDefaults() *ProfileGetRequest {
	this := ProfileGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProfileGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProfileGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProfileGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProfileGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProfileGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProfileGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetProfileToken returns the ProfileToken field value
func (o *ProfileGetRequest) GetProfileToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileToken
}

// GetProfileTokenOk returns a tuple with the ProfileToken field value
// and a boolean to check if the value has been set.
func (o *ProfileGetRequest) GetProfileTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProfileToken, true
}

// SetProfileToken sets field value
func (o *ProfileGetRequest) SetProfileToken(v string) {
	o.ProfileToken = v
}

func (o ProfileGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["profile_token"] = o.ProfileToken
	}
	return json.Marshal(toSerialize)
}

type NullableProfileGetRequest struct {
	value *ProfileGetRequest
	isSet bool
}

func (v NullableProfileGetRequest) Get() *ProfileGetRequest {
	return v.value
}

func (v *NullableProfileGetRequest) Set(val *ProfileGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileGetRequest(val *ProfileGetRequest) *NullableProfileGetRequest {
	return &NullableProfileGetRequest{value: val, isSet: true}
}

func (v NullableProfileGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


