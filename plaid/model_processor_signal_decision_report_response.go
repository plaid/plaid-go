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

// ProcessorSignalDecisionReportResponse ProcessorSignalDecisionReportResponse defines the response schema for `/processor/signal/decision/report`
type ProcessorSignalDecisionReportResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorSignalDecisionReportResponse ProcessorSignalDecisionReportResponse

// NewProcessorSignalDecisionReportResponse instantiates a new ProcessorSignalDecisionReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorSignalDecisionReportResponse(requestId string) *ProcessorSignalDecisionReportResponse {
	this := ProcessorSignalDecisionReportResponse{}
	this.RequestId = requestId
	return &this
}

// NewProcessorSignalDecisionReportResponseWithDefaults instantiates a new ProcessorSignalDecisionReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorSignalDecisionReportResponseWithDefaults() *ProcessorSignalDecisionReportResponse {
	this := ProcessorSignalDecisionReportResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ProcessorSignalDecisionReportResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProcessorSignalDecisionReportResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProcessorSignalDecisionReportResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ProcessorSignalDecisionReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorSignalDecisionReportResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorSignalDecisionReportResponse := _ProcessorSignalDecisionReportResponse{}

	if err = json.Unmarshal(bytes, &varProcessorSignalDecisionReportResponse); err == nil {
		*o = ProcessorSignalDecisionReportResponse(varProcessorSignalDecisionReportResponse)
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

type NullableProcessorSignalDecisionReportResponse struct {
	value *ProcessorSignalDecisionReportResponse
	isSet bool
}

func (v NullableProcessorSignalDecisionReportResponse) Get() *ProcessorSignalDecisionReportResponse {
	return v.value
}

func (v *NullableProcessorSignalDecisionReportResponse) Set(val *ProcessorSignalDecisionReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorSignalDecisionReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorSignalDecisionReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorSignalDecisionReportResponse(val *ProcessorSignalDecisionReportResponse) *NullableProcessorSignalDecisionReportResponse {
	return &NullableProcessorSignalDecisionReportResponse{value: val, isSet: true}
}

func (v NullableProcessorSignalDecisionReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorSignalDecisionReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


