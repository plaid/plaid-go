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

// UserAccountSessionGetRequest UserAccountSessionGetRequest defines the request schema for `/user_account/session/get`
type UserAccountSessionGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The public token generated by the end user Layer session.
	PublicToken string `json:"public_token"`
}

// NewUserAccountSessionGetRequest instantiates a new UserAccountSessionGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccountSessionGetRequest(publicToken string) *UserAccountSessionGetRequest {
	this := UserAccountSessionGetRequest{}
	this.PublicToken = publicToken
	return &this
}

// NewUserAccountSessionGetRequestWithDefaults instantiates a new UserAccountSessionGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccountSessionGetRequestWithDefaults() *UserAccountSessionGetRequest {
	this := UserAccountSessionGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *UserAccountSessionGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccountSessionGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *UserAccountSessionGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *UserAccountSessionGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *UserAccountSessionGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccountSessionGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *UserAccountSessionGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *UserAccountSessionGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetPublicToken returns the PublicToken field value
func (o *UserAccountSessionGetRequest) GetPublicToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicToken
}

// GetPublicTokenOk returns a tuple with the PublicToken field value
// and a boolean to check if the value has been set.
func (o *UserAccountSessionGetRequest) GetPublicTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicToken, true
}

// SetPublicToken sets field value
func (o *UserAccountSessionGetRequest) SetPublicToken(v string) {
	o.PublicToken = v
}

func (o UserAccountSessionGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["public_token"] = o.PublicToken
	}
	return json.Marshal(toSerialize)
}

type NullableUserAccountSessionGetRequest struct {
	value *UserAccountSessionGetRequest
	isSet bool
}

func (v NullableUserAccountSessionGetRequest) Get() *UserAccountSessionGetRequest {
	return v.value
}

func (v *NullableUserAccountSessionGetRequest) Set(val *UserAccountSessionGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccountSessionGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccountSessionGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccountSessionGetRequest(val *UserAccountSessionGetRequest) *NullableUserAccountSessionGetRequest {
	return &NullableUserAccountSessionGetRequest{value: val, isSet: true}
}

func (v NullableUserAccountSessionGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccountSessionGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

