/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.19.10
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SandboxPublicTokenCreateResponse SandboxPublicTokenCreateResponse defines the response schema for `/sandbox/public_token/create`
type SandboxPublicTokenCreateResponse struct {
	// A public token that can be exchanged for an access token using `/item/public_token/exchange`
	PublicToken string `json:"public_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxPublicTokenCreateResponse SandboxPublicTokenCreateResponse

// NewSandboxPublicTokenCreateResponse instantiates a new SandboxPublicTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxPublicTokenCreateResponse(publicToken string, requestId string) *SandboxPublicTokenCreateResponse {
	this := SandboxPublicTokenCreateResponse{}
	this.PublicToken = publicToken
	this.RequestId = requestId
	return &this
}

// NewSandboxPublicTokenCreateResponseWithDefaults instantiates a new SandboxPublicTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxPublicTokenCreateResponseWithDefaults() *SandboxPublicTokenCreateResponse {
	this := SandboxPublicTokenCreateResponse{}
	return &this
}

// GetPublicToken returns the PublicToken field value
func (o *SandboxPublicTokenCreateResponse) GetPublicToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicToken
}

// GetPublicTokenOk returns a tuple with the PublicToken field value
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateResponse) GetPublicTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicToken, true
}

// SetPublicToken sets field value
func (o *SandboxPublicTokenCreateResponse) SetPublicToken(v string) {
	o.PublicToken = v
}

// GetRequestId returns the RequestId field value
func (o *SandboxPublicTokenCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxPublicTokenCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxPublicTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["public_token"] = o.PublicToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxPublicTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxPublicTokenCreateResponse := _SandboxPublicTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varSandboxPublicTokenCreateResponse); err == nil {
		*o = SandboxPublicTokenCreateResponse(varSandboxPublicTokenCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "public_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxPublicTokenCreateResponse struct {
	value *SandboxPublicTokenCreateResponse
	isSet bool
}

func (v NullableSandboxPublicTokenCreateResponse) Get() *SandboxPublicTokenCreateResponse {
	return v.value
}

func (v *NullableSandboxPublicTokenCreateResponse) Set(val *SandboxPublicTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxPublicTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxPublicTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxPublicTokenCreateResponse(val *SandboxPublicTokenCreateResponse) *NullableSandboxPublicTokenCreateResponse {
	return &NullableSandboxPublicTokenCreateResponse{value: val, isSet: true}
}

func (v NullableSandboxPublicTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxPublicTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

