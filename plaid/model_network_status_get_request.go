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

// NetworkStatusGetRequest NetworkStatusGetRequest defines the request schema for `/network/status/get`
type NetworkStatusGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	User NetworkStatusGetUser `json:"user"`
	// The id of a template defined in Plaid Dashboard. This field is used if you have additional criteria that you want to check against (e.g. Layer eligibility).
	TemplateId *string `json:"template_id,omitempty"`
}

// NewNetworkStatusGetRequest instantiates a new NetworkStatusGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkStatusGetRequest(user NetworkStatusGetUser) *NetworkStatusGetRequest {
	this := NetworkStatusGetRequest{}
	this.User = user
	return &this
}

// NewNetworkStatusGetRequestWithDefaults instantiates a new NetworkStatusGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkStatusGetRequestWithDefaults() *NetworkStatusGetRequest {
	this := NetworkStatusGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *NetworkStatusGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkStatusGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *NetworkStatusGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *NetworkStatusGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *NetworkStatusGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkStatusGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *NetworkStatusGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *NetworkStatusGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUser returns the User field value
func (o *NetworkStatusGetRequest) GetUser() NetworkStatusGetUser {
	if o == nil {
		var ret NetworkStatusGetUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *NetworkStatusGetRequest) GetUserOk() (*NetworkStatusGetUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *NetworkStatusGetRequest) SetUser(v NetworkStatusGetUser) {
	o.User = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *NetworkStatusGetRequest) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkStatusGetRequest) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *NetworkStatusGetRequest) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *NetworkStatusGetRequest) SetTemplateId(v string) {
	o.TemplateId = &v
}

func (o NetworkStatusGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.TemplateId != nil {
		toSerialize["template_id"] = o.TemplateId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkStatusGetRequest struct {
	value *NetworkStatusGetRequest
	isSet bool
}

func (v NullableNetworkStatusGetRequest) Get() *NetworkStatusGetRequest {
	return v.value
}

func (v *NullableNetworkStatusGetRequest) Set(val *NetworkStatusGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkStatusGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkStatusGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkStatusGetRequest(val *NetworkStatusGetRequest) *NullableNetworkStatusGetRequest {
	return &NullableNetworkStatusGetRequest{value: val, isSet: true}
}

func (v NullableNetworkStatusGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkStatusGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


