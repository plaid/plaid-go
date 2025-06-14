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

// IdentityMatchRequest IdentityMatchRequest defines the request schema for `/identity/match`
type IdentityMatchRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	User *IdentityMatchUser `json:"user,omitempty"`
	Options *IdentityMatchRequestOptions `json:"options,omitempty"`
}

// NewIdentityMatchRequest instantiates a new IdentityMatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityMatchRequest(accessToken string) *IdentityMatchRequest {
	this := IdentityMatchRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewIdentityMatchRequestWithDefaults instantiates a new IdentityMatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityMatchRequestWithDefaults() *IdentityMatchRequest {
	this := IdentityMatchRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IdentityMatchRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMatchRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IdentityMatchRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IdentityMatchRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IdentityMatchRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMatchRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IdentityMatchRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IdentityMatchRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *IdentityMatchRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *IdentityMatchRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *IdentityMatchRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IdentityMatchRequest) GetUser() IdentityMatchUser {
	if o == nil || o.User == nil {
		var ret IdentityMatchUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMatchRequest) GetUserOk() (*IdentityMatchUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IdentityMatchRequest) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IdentityMatchUser and assigns it to the User field.
func (o *IdentityMatchRequest) SetUser(v IdentityMatchUser) {
	o.User = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *IdentityMatchRequest) GetOptions() IdentityMatchRequestOptions {
	if o == nil || o.Options == nil {
		var ret IdentityMatchRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMatchRequest) GetOptionsOk() (*IdentityMatchRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *IdentityMatchRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given IdentityMatchRequestOptions and assigns it to the Options field.
func (o *IdentityMatchRequest) SetOptions(v IdentityMatchRequestOptions) {
	o.Options = &v
}

func (o IdentityMatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityMatchRequest struct {
	value *IdentityMatchRequest
	isSet bool
}

func (v NullableIdentityMatchRequest) Get() *IdentityMatchRequest {
	return v.value
}

func (v *NullableIdentityMatchRequest) Set(val *IdentityMatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityMatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityMatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityMatchRequest(val *IdentityMatchRequest) *NullableIdentityMatchRequest {
	return &NullableIdentityMatchRequest{value: val, isSet: true}
}

func (v NullableIdentityMatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityMatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


