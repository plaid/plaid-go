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

// TransferLedgerWithdrawResponse Defines the response schema for `/transfer/ledger/withdraw`
type TransferLedgerWithdrawResponse struct {
	Sweep TransferSweep `json:"sweep"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferLedgerWithdrawResponse TransferLedgerWithdrawResponse

// NewTransferLedgerWithdrawResponse instantiates a new TransferLedgerWithdrawResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferLedgerWithdrawResponse(sweep TransferSweep, requestId string) *TransferLedgerWithdrawResponse {
	this := TransferLedgerWithdrawResponse{}
	this.Sweep = sweep
	this.RequestId = requestId
	return &this
}

// NewTransferLedgerWithdrawResponseWithDefaults instantiates a new TransferLedgerWithdrawResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferLedgerWithdrawResponseWithDefaults() *TransferLedgerWithdrawResponse {
	this := TransferLedgerWithdrawResponse{}
	return &this
}

// GetSweep returns the Sweep field value
func (o *TransferLedgerWithdrawResponse) GetSweep() TransferSweep {
	if o == nil {
		var ret TransferSweep
		return ret
	}

	return o.Sweep
}

// GetSweepOk returns a tuple with the Sweep field value
// and a boolean to check if the value has been set.
func (o *TransferLedgerWithdrawResponse) GetSweepOk() (*TransferSweep, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sweep, true
}

// SetSweep sets field value
func (o *TransferLedgerWithdrawResponse) SetSweep(v TransferSweep) {
	o.Sweep = v
}

// GetRequestId returns the RequestId field value
func (o *TransferLedgerWithdrawResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferLedgerWithdrawResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferLedgerWithdrawResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferLedgerWithdrawResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sweep"] = o.Sweep
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferLedgerWithdrawResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferLedgerWithdrawResponse := _TransferLedgerWithdrawResponse{}

	if err = json.Unmarshal(bytes, &varTransferLedgerWithdrawResponse); err == nil {
		*o = TransferLedgerWithdrawResponse(varTransferLedgerWithdrawResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sweep")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferLedgerWithdrawResponse struct {
	value *TransferLedgerWithdrawResponse
	isSet bool
}

func (v NullableTransferLedgerWithdrawResponse) Get() *TransferLedgerWithdrawResponse {
	return v.value
}

func (v *NullableTransferLedgerWithdrawResponse) Set(val *TransferLedgerWithdrawResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferLedgerWithdrawResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferLedgerWithdrawResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferLedgerWithdrawResponse(val *TransferLedgerWithdrawResponse) *NullableTransferLedgerWithdrawResponse {
	return &NullableTransferLedgerWithdrawResponse{value: val, isSet: true}
}

func (v NullableTransferLedgerWithdrawResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferLedgerWithdrawResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

