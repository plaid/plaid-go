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

// SandboxUserResetLoginRequest SandboxUserResetLoginRequest defines the request schema for `/sandbox/user/reset_login`
type SandboxUserResetLoginRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken string `json:"user_token"`
	// An array of `item_id`s associated with the User to be reset.  If empty or `null`, this field will default to resetting all Items associated with the User.
	ItemIds []string `json:"item_ids,omitempty"`
}

// NewSandboxUserResetLoginRequest instantiates a new SandboxUserResetLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxUserResetLoginRequest(userToken string) *SandboxUserResetLoginRequest {
	this := SandboxUserResetLoginRequest{}
	this.UserToken = userToken
	return &this
}

// NewSandboxUserResetLoginRequestWithDefaults instantiates a new SandboxUserResetLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxUserResetLoginRequestWithDefaults() *SandboxUserResetLoginRequest {
	this := SandboxUserResetLoginRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxUserResetLoginRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxUserResetLoginRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxUserResetLoginRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxUserResetLoginRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxUserResetLoginRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxUserResetLoginRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxUserResetLoginRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxUserResetLoginRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value
func (o *SandboxUserResetLoginRequest) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *SandboxUserResetLoginRequest) GetUserTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *SandboxUserResetLoginRequest) SetUserToken(v string) {
	o.UserToken = v
}

// GetItemIds returns the ItemIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxUserResetLoginRequest) GetItemIds() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxUserResetLoginRequest) GetItemIdsOk() (*[]string, bool) {
	if o == nil || o.ItemIds == nil {
		return nil, false
	}
	return &o.ItemIds, true
}

// HasItemIds returns a boolean if a field has been set.
func (o *SandboxUserResetLoginRequest) HasItemIds() bool {
	if o != nil && o.ItemIds != nil {
		return true
	}

	return false
}

// SetItemIds gets a reference to the given []string and assigns it to the ItemIds field.
func (o *SandboxUserResetLoginRequest) SetItemIds(v []string) {
	o.ItemIds = v
}

func (o SandboxUserResetLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["user_token"] = o.UserToken
	}
	if o.ItemIds != nil {
		toSerialize["item_ids"] = o.ItemIds
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxUserResetLoginRequest struct {
	value *SandboxUserResetLoginRequest
	isSet bool
}

func (v NullableSandboxUserResetLoginRequest) Get() *SandboxUserResetLoginRequest {
	return v.value
}

func (v *NullableSandboxUserResetLoginRequest) Set(val *SandboxUserResetLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxUserResetLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxUserResetLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxUserResetLoginRequest(val *SandboxUserResetLoginRequest) *NullableSandboxUserResetLoginRequest {
	return &NullableSandboxUserResetLoginRequest{value: val, isSet: true}
}

func (v NullableSandboxUserResetLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxUserResetLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


