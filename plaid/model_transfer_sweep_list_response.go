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

// TransferSweepListResponse Defines the response schema for `/transfer/sweep/list`
type TransferSweepListResponse struct {
	Sweeps []TransferSweep `json:"sweeps"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferSweepListResponse TransferSweepListResponse

// NewTransferSweepListResponse instantiates a new TransferSweepListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferSweepListResponse(sweeps []TransferSweep, requestId string) *TransferSweepListResponse {
	this := TransferSweepListResponse{}
	this.Sweeps = sweeps
	this.RequestId = requestId
	return &this
}

// NewTransferSweepListResponseWithDefaults instantiates a new TransferSweepListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferSweepListResponseWithDefaults() *TransferSweepListResponse {
	this := TransferSweepListResponse{}
	return &this
}

// GetSweeps returns the Sweeps field value
func (o *TransferSweepListResponse) GetSweeps() []TransferSweep {
	if o == nil {
		var ret []TransferSweep
		return ret
	}

	return o.Sweeps
}

// GetSweepsOk returns a tuple with the Sweeps field value
// and a boolean to check if the value has been set.
func (o *TransferSweepListResponse) GetSweepsOk() (*[]TransferSweep, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sweeps, true
}

// SetSweeps sets field value
func (o *TransferSweepListResponse) SetSweeps(v []TransferSweep) {
	o.Sweeps = v
}

// GetRequestId returns the RequestId field value
func (o *TransferSweepListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferSweepListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferSweepListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferSweepListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sweeps"] = o.Sweeps
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferSweepListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferSweepListResponse := _TransferSweepListResponse{}

	if err = json.Unmarshal(bytes, &varTransferSweepListResponse); err == nil {
		*o = TransferSweepListResponse(varTransferSweepListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sweeps")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferSweepListResponse struct {
	value *TransferSweepListResponse
	isSet bool
}

func (v NullableTransferSweepListResponse) Get() *TransferSweepListResponse {
	return v.value
}

func (v *NullableTransferSweepListResponse) Set(val *TransferSweepListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferSweepListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferSweepListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferSweepListResponse(val *TransferSweepListResponse) *NullableTransferSweepListResponse {
	return &NullableTransferSweepListResponse{value: val, isSet: true}
}

func (v NullableTransferSweepListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferSweepListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


