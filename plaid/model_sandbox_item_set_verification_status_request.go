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

// SandboxItemSetVerificationStatusRequest SandboxItemSetVerificationStatusRequest defines the request schema for `/sandbox/item/set_verification_status`
type SandboxItemSetVerificationStatusRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// The `account_id` of the account whose verification status is to be modified
	AccountId string `json:"account_id"`
	// The verification status to set the account to.
	VerificationStatus string `json:"verification_status"`
}

// NewSandboxItemSetVerificationStatusRequest instantiates a new SandboxItemSetVerificationStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxItemSetVerificationStatusRequest(accessToken string, accountId string, verificationStatus string) *SandboxItemSetVerificationStatusRequest {
	this := SandboxItemSetVerificationStatusRequest{}
	this.AccessToken = accessToken
	this.AccountId = accountId
	this.VerificationStatus = verificationStatus
	return &this
}

// NewSandboxItemSetVerificationStatusRequestWithDefaults instantiates a new SandboxItemSetVerificationStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxItemSetVerificationStatusRequestWithDefaults() *SandboxItemSetVerificationStatusRequest {
	this := SandboxItemSetVerificationStatusRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxItemSetVerificationStatusRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxItemSetVerificationStatusRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxItemSetVerificationStatusRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxItemSetVerificationStatusRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxItemSetVerificationStatusRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxItemSetVerificationStatusRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxItemSetVerificationStatusRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxItemSetVerificationStatusRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *SandboxItemSetVerificationStatusRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *SandboxItemSetVerificationStatusRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *SandboxItemSetVerificationStatusRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetAccountId returns the AccountId field value
func (o *SandboxItemSetVerificationStatusRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *SandboxItemSetVerificationStatusRequest) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *SandboxItemSetVerificationStatusRequest) SetAccountId(v string) {
	o.AccountId = v
}

// GetVerificationStatus returns the VerificationStatus field value
func (o *SandboxItemSetVerificationStatusRequest) GetVerificationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value
// and a boolean to check if the value has been set.
func (o *SandboxItemSetVerificationStatusRequest) GetVerificationStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationStatus, true
}

// SetVerificationStatus sets field value
func (o *SandboxItemSetVerificationStatusRequest) SetVerificationStatus(v string) {
	o.VerificationStatus = v
}

func (o SandboxItemSetVerificationStatusRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxItemSetVerificationStatusRequest struct {
	value *SandboxItemSetVerificationStatusRequest
	isSet bool
}

func (v NullableSandboxItemSetVerificationStatusRequest) Get() *SandboxItemSetVerificationStatusRequest {
	return v.value
}

func (v *NullableSandboxItemSetVerificationStatusRequest) Set(val *SandboxItemSetVerificationStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxItemSetVerificationStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxItemSetVerificationStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxItemSetVerificationStatusRequest(val *SandboxItemSetVerificationStatusRequest) *NullableSandboxItemSetVerificationStatusRequest {
	return &NullableSandboxItemSetVerificationStatusRequest{value: val, isSet: true}
}

func (v NullableSandboxItemSetVerificationStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxItemSetVerificationStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


