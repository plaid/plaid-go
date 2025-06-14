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

// BankTransferSweepGetResponse BankTransferSweepGetResponse defines the response schema for `/bank_transfer/sweep/get`
type BankTransferSweepGetResponse struct {
	Sweep BankTransferSweep `json:"sweep"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferSweepGetResponse BankTransferSweepGetResponse

// NewBankTransferSweepGetResponse instantiates a new BankTransferSweepGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferSweepGetResponse(sweep BankTransferSweep, requestId string) *BankTransferSweepGetResponse {
	this := BankTransferSweepGetResponse{}
	this.Sweep = sweep
	this.RequestId = requestId
	return &this
}

// NewBankTransferSweepGetResponseWithDefaults instantiates a new BankTransferSweepGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferSweepGetResponseWithDefaults() *BankTransferSweepGetResponse {
	this := BankTransferSweepGetResponse{}
	return &this
}

// GetSweep returns the Sweep field value
func (o *BankTransferSweepGetResponse) GetSweep() BankTransferSweep {
	if o == nil {
		var ret BankTransferSweep
		return ret
	}

	return o.Sweep
}

// GetSweepOk returns a tuple with the Sweep field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweepGetResponse) GetSweepOk() (*BankTransferSweep, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sweep, true
}

// SetSweep sets field value
func (o *BankTransferSweepGetResponse) SetSweep(v BankTransferSweep) {
	o.Sweep = v
}

// GetRequestId returns the RequestId field value
func (o *BankTransferSweepGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweepGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BankTransferSweepGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BankTransferSweepGetResponse) MarshalJSON() ([]byte, error) {
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

func (o *BankTransferSweepGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferSweepGetResponse := _BankTransferSweepGetResponse{}

	if err = json.Unmarshal(bytes, &varBankTransferSweepGetResponse); err == nil {
		*o = BankTransferSweepGetResponse(varBankTransferSweepGetResponse)
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

type NullableBankTransferSweepGetResponse struct {
	value *BankTransferSweepGetResponse
	isSet bool
}

func (v NullableBankTransferSweepGetResponse) Get() *BankTransferSweepGetResponse {
	return v.value
}

func (v *NullableBankTransferSweepGetResponse) Set(val *BankTransferSweepGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferSweepGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferSweepGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferSweepGetResponse(val *BankTransferSweepGetResponse) *NullableBankTransferSweepGetResponse {
	return &NullableBankTransferSweepGetResponse{value: val, isSet: true}
}

func (v NullableBankTransferSweepGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferSweepGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


