/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.370.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ProcessorIdentityMatchResponse ProcessorIdentityMatchResponse defines the response schema for `/processor/identity/match`
type ProcessorIdentityMatchResponse struct {
	Account AccountIdentityMatchScore `json:"account"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorIdentityMatchResponse ProcessorIdentityMatchResponse

// NewProcessorIdentityMatchResponse instantiates a new ProcessorIdentityMatchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorIdentityMatchResponse(account AccountIdentityMatchScore, requestId string) *ProcessorIdentityMatchResponse {
	this := ProcessorIdentityMatchResponse{}
	this.Account = account
	this.RequestId = requestId
	return &this
}

// NewProcessorIdentityMatchResponseWithDefaults instantiates a new ProcessorIdentityMatchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorIdentityMatchResponseWithDefaults() *ProcessorIdentityMatchResponse {
	this := ProcessorIdentityMatchResponse{}
	return &this
}

// GetAccount returns the Account field value
func (o *ProcessorIdentityMatchResponse) GetAccount() AccountIdentityMatchScore {
	if o == nil {
		var ret AccountIdentityMatchScore
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *ProcessorIdentityMatchResponse) GetAccountOk() (*AccountIdentityMatchScore, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *ProcessorIdentityMatchResponse) SetAccount(v AccountIdentityMatchScore) {
	o.Account = v
}

// GetRequestId returns the RequestId field value
func (o *ProcessorIdentityMatchResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProcessorIdentityMatchResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProcessorIdentityMatchResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ProcessorIdentityMatchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorIdentityMatchResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorIdentityMatchResponse := _ProcessorIdentityMatchResponse{}

	if err = json.Unmarshal(bytes, &varProcessorIdentityMatchResponse); err == nil {
		*o = ProcessorIdentityMatchResponse(varProcessorIdentityMatchResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorIdentityMatchResponse struct {
	value *ProcessorIdentityMatchResponse
	isSet bool
}

func (v NullableProcessorIdentityMatchResponse) Get() *ProcessorIdentityMatchResponse {
	return v.value
}

func (v *NullableProcessorIdentityMatchResponse) Set(val *ProcessorIdentityMatchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorIdentityMatchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorIdentityMatchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorIdentityMatchResponse(val *ProcessorIdentityMatchResponse) *NullableProcessorIdentityMatchResponse {
	return &NullableProcessorIdentityMatchResponse{value: val, isSet: true}
}

func (v NullableProcessorIdentityMatchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorIdentityMatchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

