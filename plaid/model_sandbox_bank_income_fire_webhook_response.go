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

// SandboxBankIncomeFireWebhookResponse SandboxBankIncomeFireWebhookResponse defines the response schema for `/sandbox/bank_income/fire_webhook`
type SandboxBankIncomeFireWebhookResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxBankIncomeFireWebhookResponse SandboxBankIncomeFireWebhookResponse

// NewSandboxBankIncomeFireWebhookResponse instantiates a new SandboxBankIncomeFireWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxBankIncomeFireWebhookResponse(requestId string) *SandboxBankIncomeFireWebhookResponse {
	this := SandboxBankIncomeFireWebhookResponse{}
	this.RequestId = requestId
	return &this
}

// NewSandboxBankIncomeFireWebhookResponseWithDefaults instantiates a new SandboxBankIncomeFireWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxBankIncomeFireWebhookResponseWithDefaults() *SandboxBankIncomeFireWebhookResponse {
	this := SandboxBankIncomeFireWebhookResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SandboxBankIncomeFireWebhookResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxBankIncomeFireWebhookResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxBankIncomeFireWebhookResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxBankIncomeFireWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxBankIncomeFireWebhookResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxBankIncomeFireWebhookResponse := _SandboxBankIncomeFireWebhookResponse{}

	if err = json.Unmarshal(bytes, &varSandboxBankIncomeFireWebhookResponse); err == nil {
		*o = SandboxBankIncomeFireWebhookResponse(varSandboxBankIncomeFireWebhookResponse)
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

type NullableSandboxBankIncomeFireWebhookResponse struct {
	value *SandboxBankIncomeFireWebhookResponse
	isSet bool
}

func (v NullableSandboxBankIncomeFireWebhookResponse) Get() *SandboxBankIncomeFireWebhookResponse {
	return v.value
}

func (v *NullableSandboxBankIncomeFireWebhookResponse) Set(val *SandboxBankIncomeFireWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxBankIncomeFireWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxBankIncomeFireWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxBankIncomeFireWebhookResponse(val *SandboxBankIncomeFireWebhookResponse) *NullableSandboxBankIncomeFireWebhookResponse {
	return &NullableSandboxBankIncomeFireWebhookResponse{value: val, isSet: true}
}

func (v NullableSandboxBankIncomeFireWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxBankIncomeFireWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


