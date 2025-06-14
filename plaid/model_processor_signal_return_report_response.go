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

// ProcessorSignalReturnReportResponse ProcessorSignalReturnReportResponse defines the response schema for `/processor/signal/return/report`
type ProcessorSignalReturnReportResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorSignalReturnReportResponse ProcessorSignalReturnReportResponse

// NewProcessorSignalReturnReportResponse instantiates a new ProcessorSignalReturnReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorSignalReturnReportResponse(requestId string) *ProcessorSignalReturnReportResponse {
	this := ProcessorSignalReturnReportResponse{}
	this.RequestId = requestId
	return &this
}

// NewProcessorSignalReturnReportResponseWithDefaults instantiates a new ProcessorSignalReturnReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorSignalReturnReportResponseWithDefaults() *ProcessorSignalReturnReportResponse {
	this := ProcessorSignalReturnReportResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ProcessorSignalReturnReportResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProcessorSignalReturnReportResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProcessorSignalReturnReportResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ProcessorSignalReturnReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorSignalReturnReportResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorSignalReturnReportResponse := _ProcessorSignalReturnReportResponse{}

	if err = json.Unmarshal(bytes, &varProcessorSignalReturnReportResponse); err == nil {
		*o = ProcessorSignalReturnReportResponse(varProcessorSignalReturnReportResponse)
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

type NullableProcessorSignalReturnReportResponse struct {
	value *ProcessorSignalReturnReportResponse
	isSet bool
}

func (v NullableProcessorSignalReturnReportResponse) Get() *ProcessorSignalReturnReportResponse {
	return v.value
}

func (v *NullableProcessorSignalReturnReportResponse) Set(val *ProcessorSignalReturnReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorSignalReturnReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorSignalReturnReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorSignalReturnReportResponse(val *ProcessorSignalReturnReportResponse) *NullableProcessorSignalReturnReportResponse {
	return &NullableProcessorSignalReturnReportResponse{value: val, isSet: true}
}

func (v NullableProcessorSignalReturnReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorSignalReturnReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


