/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CraBankIncomeGetRequest CraBankIncomeGetRequest defines the request schema for `/cra/bank_income/get`.
type CraBankIncomeGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken *string `json:"user_token,omitempty"`
}

// NewCraBankIncomeGetRequest instantiates a new CraBankIncomeGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraBankIncomeGetRequest() *CraBankIncomeGetRequest {
	this := CraBankIncomeGetRequest{}
	return &this
}

// NewCraBankIncomeGetRequestWithDefaults instantiates a new CraBankIncomeGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraBankIncomeGetRequestWithDefaults() *CraBankIncomeGetRequest {
	this := CraBankIncomeGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CraBankIncomeGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CraBankIncomeGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CraBankIncomeGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CraBankIncomeGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CraBankIncomeGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CraBankIncomeGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *CraBankIncomeGetRequest) GetUserToken() string {
	if o == nil || o.UserToken == nil {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeGetRequest) GetUserTokenOk() (*string, bool) {
	if o == nil || o.UserToken == nil {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *CraBankIncomeGetRequest) HasUserToken() bool {
	if o != nil && o.UserToken != nil {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *CraBankIncomeGetRequest) SetUserToken(v string) {
	o.UserToken = &v
}

func (o CraBankIncomeGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.UserToken != nil {
		toSerialize["user_token"] = o.UserToken
	}
	return json.Marshal(toSerialize)
}

type NullableCraBankIncomeGetRequest struct {
	value *CraBankIncomeGetRequest
	isSet bool
}

func (v NullableCraBankIncomeGetRequest) Get() *CraBankIncomeGetRequest {
	return v.value
}

func (v *NullableCraBankIncomeGetRequest) Set(val *CraBankIncomeGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCraBankIncomeGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCraBankIncomeGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraBankIncomeGetRequest(val *CraBankIncomeGetRequest) *NullableCraBankIncomeGetRequest {
	return &NullableCraBankIncomeGetRequest{value: val, isSet: true}
}

func (v NullableCraBankIncomeGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraBankIncomeGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

