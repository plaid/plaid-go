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

// TransferRepaymentReturnListRequest Defines the request schema for `/transfer/repayment/return/list`
type TransferRepaymentReturnListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Identifier of the repayment to query.
	RepaymentId string `json:"repayment_id"`
	// The maximum number of repayments to return.
	Count NullableInt32 `json:"count,omitempty"`
	// The number of repayments to skip before returning results.
	Offset *int32 `json:"offset,omitempty"`
}

// NewTransferRepaymentReturnListRequest instantiates a new TransferRepaymentReturnListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRepaymentReturnListRequest(repaymentId string) *TransferRepaymentReturnListRequest {
	this := TransferRepaymentReturnListRequest{}
	this.RepaymentId = repaymentId
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// NewTransferRepaymentReturnListRequestWithDefaults instantiates a new TransferRepaymentReturnListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRepaymentReturnListRequestWithDefaults() *TransferRepaymentReturnListRequest {
	this := TransferRepaymentReturnListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferRepaymentReturnListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRepaymentReturnListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferRepaymentReturnListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferRepaymentReturnListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferRepaymentReturnListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRepaymentReturnListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferRepaymentReturnListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferRepaymentReturnListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetRepaymentId returns the RepaymentId field value
func (o *TransferRepaymentReturnListRequest) GetRepaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepaymentId
}

// GetRepaymentIdOk returns a tuple with the RepaymentId field value
// and a boolean to check if the value has been set.
func (o *TransferRepaymentReturnListRequest) GetRepaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RepaymentId, true
}

// SetRepaymentId sets field value
func (o *TransferRepaymentReturnListRequest) SetRepaymentId(v string) {
	o.RepaymentId = v
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferRepaymentReturnListRequest) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferRepaymentReturnListRequest) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *TransferRepaymentReturnListRequest) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *TransferRepaymentReturnListRequest) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *TransferRepaymentReturnListRequest) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *TransferRepaymentReturnListRequest) UnsetCount() {
	o.Count.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TransferRepaymentReturnListRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRepaymentReturnListRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TransferRepaymentReturnListRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *TransferRepaymentReturnListRequest) SetOffset(v int32) {
	o.Offset = &v
}

func (o TransferRepaymentReturnListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["repayment_id"] = o.RepaymentId
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	return json.Marshal(toSerialize)
}

type NullableTransferRepaymentReturnListRequest struct {
	value *TransferRepaymentReturnListRequest
	isSet bool
}

func (v NullableTransferRepaymentReturnListRequest) Get() *TransferRepaymentReturnListRequest {
	return v.value
}

func (v *NullableTransferRepaymentReturnListRequest) Set(val *TransferRepaymentReturnListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRepaymentReturnListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRepaymentReturnListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRepaymentReturnListRequest(val *TransferRepaymentReturnListRequest) *NullableTransferRepaymentReturnListRequest {
	return &NullableTransferRepaymentReturnListRequest{value: val, isSet: true}
}

func (v NullableTransferRepaymentReturnListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRepaymentReturnListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


