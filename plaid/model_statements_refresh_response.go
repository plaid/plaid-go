/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.503.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// StatementsRefreshResponse StatementsRefreshResponse defines the response schema for `/statements/refresh`
type StatementsRefreshResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _StatementsRefreshResponse StatementsRefreshResponse

// NewStatementsRefreshResponse instantiates a new StatementsRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementsRefreshResponse(requestId string) *StatementsRefreshResponse {
	this := StatementsRefreshResponse{}
	this.RequestId = requestId
	return &this
}

// NewStatementsRefreshResponseWithDefaults instantiates a new StatementsRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementsRefreshResponseWithDefaults() *StatementsRefreshResponse {
	this := StatementsRefreshResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *StatementsRefreshResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *StatementsRefreshResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *StatementsRefreshResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o StatementsRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StatementsRefreshResponse) UnmarshalJSON(bytes []byte) (err error) {
	varStatementsRefreshResponse := _StatementsRefreshResponse{}

	if err = json.Unmarshal(bytes, &varStatementsRefreshResponse); err == nil {
		*o = StatementsRefreshResponse(varStatementsRefreshResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStatementsRefreshResponse struct {
	value *StatementsRefreshResponse
	isSet bool
}

func (v NullableStatementsRefreshResponse) Get() *StatementsRefreshResponse {
	return v.value
}

func (v *NullableStatementsRefreshResponse) Set(val *StatementsRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementsRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementsRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementsRefreshResponse(val *StatementsRefreshResponse) *NullableStatementsRefreshResponse {
	return &NullableStatementsRefreshResponse{value: val, isSet: true}
}

func (v NullableStatementsRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementsRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

