/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.419.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferRefundFailure The failure reason if the event type for a refund is `\"failed\"` or `\"returned\"`. Null value otherwise.
type TransferRefundFailure struct {
	// The ACH return code, e.g. `R01`.  A return code will be provided if and only if the refund status is `returned`. For a full listing of ACH return codes, see [Transfer errors](https://plaid.com/docs/errors/transfer/#ach-return-codes).
	AchReturnCode NullableString `json:"ach_return_code,omitempty"`
	// A human-readable description of the reason for the failure or reversal.
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferRefundFailure TransferRefundFailure

// NewTransferRefundFailure instantiates a new TransferRefundFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRefundFailure() *TransferRefundFailure {
	this := TransferRefundFailure{}
	return &this
}

// NewTransferRefundFailureWithDefaults instantiates a new TransferRefundFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRefundFailureWithDefaults() *TransferRefundFailure {
	this := TransferRefundFailure{}
	return &this
}

// GetAchReturnCode returns the AchReturnCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferRefundFailure) GetAchReturnCode() string {
	if o == nil || o.AchReturnCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.AchReturnCode.Get()
}

// GetAchReturnCodeOk returns a tuple with the AchReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferRefundFailure) GetAchReturnCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AchReturnCode.Get(), o.AchReturnCode.IsSet()
}

// HasAchReturnCode returns a boolean if a field has been set.
func (o *TransferRefundFailure) HasAchReturnCode() bool {
	if o != nil && o.AchReturnCode.IsSet() {
		return true
	}

	return false
}

// SetAchReturnCode gets a reference to the given NullableString and assigns it to the AchReturnCode field.
func (o *TransferRefundFailure) SetAchReturnCode(v string) {
	o.AchReturnCode.Set(&v)
}
// SetAchReturnCodeNil sets the value for AchReturnCode to be an explicit nil
func (o *TransferRefundFailure) SetAchReturnCodeNil() {
	o.AchReturnCode.Set(nil)
}

// UnsetAchReturnCode ensures that no value is present for AchReturnCode, not even an explicit nil
func (o *TransferRefundFailure) UnsetAchReturnCode() {
	o.AchReturnCode.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransferRefundFailure) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRefundFailure) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransferRefundFailure) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransferRefundFailure) SetDescription(v string) {
	o.Description = &v
}

func (o TransferRefundFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AchReturnCode.IsSet() {
		toSerialize["ach_return_code"] = o.AchReturnCode.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferRefundFailure) UnmarshalJSON(bytes []byte) (err error) {
	varTransferRefundFailure := _TransferRefundFailure{}

	if err = json.Unmarshal(bytes, &varTransferRefundFailure); err == nil {
		*o = TransferRefundFailure(varTransferRefundFailure)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ach_return_code")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferRefundFailure struct {
	value *TransferRefundFailure
	isSet bool
}

func (v NullableTransferRefundFailure) Get() *TransferRefundFailure {
	return v.value
}

func (v *NullableTransferRefundFailure) Set(val *TransferRefundFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRefundFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRefundFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRefundFailure(val *TransferRefundFailure) *NullableTransferRefundFailure {
	return &NullableTransferRefundFailure{value: val, isSet: true}
}

func (v NullableTransferRefundFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRefundFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

