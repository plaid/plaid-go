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

// BaseReportCreateResponse BaseReportCreateResponse defines the response schema for `cra/base_report/create`
type BaseReportCreateResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BaseReportCreateResponse BaseReportCreateResponse

// NewBaseReportCreateResponse instantiates a new BaseReportCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseReportCreateResponse(requestId string) *BaseReportCreateResponse {
	this := BaseReportCreateResponse{}
	this.RequestId = requestId
	return &this
}

// NewBaseReportCreateResponseWithDefaults instantiates a new BaseReportCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseReportCreateResponseWithDefaults() *BaseReportCreateResponse {
	this := BaseReportCreateResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *BaseReportCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BaseReportCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BaseReportCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BaseReportCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseReportCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBaseReportCreateResponse := _BaseReportCreateResponse{}

	if err = json.Unmarshal(bytes, &varBaseReportCreateResponse); err == nil {
		*o = BaseReportCreateResponse(varBaseReportCreateResponse)
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

type NullableBaseReportCreateResponse struct {
	value *BaseReportCreateResponse
	isSet bool
}

func (v NullableBaseReportCreateResponse) Get() *BaseReportCreateResponse {
	return v.value
}

func (v *NullableBaseReportCreateResponse) Set(val *BaseReportCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseReportCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseReportCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseReportCreateResponse(val *BaseReportCreateResponse) *NullableBaseReportCreateResponse {
	return &NullableBaseReportCreateResponse{value: val, isSet: true}
}

func (v NullableBaseReportCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseReportCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

