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

// PaymentConsentPeriodicAmount Defines consent payments limitations per period.
type PaymentConsentPeriodicAmount struct {
	Amount PaymentConsentPeriodicAmountAmount `json:"amount"`
	Interval PaymentConsentPeriodicInterval `json:"interval"`
	Alignment PaymentConsentPeriodicAlignment `json:"alignment"`
}

// NewPaymentConsentPeriodicAmount instantiates a new PaymentConsentPeriodicAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentConsentPeriodicAmount(amount PaymentConsentPeriodicAmountAmount, interval PaymentConsentPeriodicInterval, alignment PaymentConsentPeriodicAlignment) *PaymentConsentPeriodicAmount {
	this := PaymentConsentPeriodicAmount{}
	this.Amount = amount
	this.Interval = interval
	this.Alignment = alignment
	return &this
}

// NewPaymentConsentPeriodicAmountWithDefaults instantiates a new PaymentConsentPeriodicAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentConsentPeriodicAmountWithDefaults() *PaymentConsentPeriodicAmount {
	this := PaymentConsentPeriodicAmount{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentConsentPeriodicAmount) GetAmount() PaymentConsentPeriodicAmountAmount {
	if o == nil {
		var ret PaymentConsentPeriodicAmountAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentConsentPeriodicAmount) GetAmountOk() (*PaymentConsentPeriodicAmountAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentConsentPeriodicAmount) SetAmount(v PaymentConsentPeriodicAmountAmount) {
	o.Amount = v
}

// GetInterval returns the Interval field value
func (o *PaymentConsentPeriodicAmount) GetInterval() PaymentConsentPeriodicInterval {
	if o == nil {
		var ret PaymentConsentPeriodicInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *PaymentConsentPeriodicAmount) GetIntervalOk() (*PaymentConsentPeriodicInterval, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *PaymentConsentPeriodicAmount) SetInterval(v PaymentConsentPeriodicInterval) {
	o.Interval = v
}

// GetAlignment returns the Alignment field value
func (o *PaymentConsentPeriodicAmount) GetAlignment() PaymentConsentPeriodicAlignment {
	if o == nil {
		var ret PaymentConsentPeriodicAlignment
		return ret
	}

	return o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value
// and a boolean to check if the value has been set.
func (o *PaymentConsentPeriodicAmount) GetAlignmentOk() (*PaymentConsentPeriodicAlignment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Alignment, true
}

// SetAlignment sets field value
func (o *PaymentConsentPeriodicAmount) SetAlignment(v PaymentConsentPeriodicAlignment) {
	o.Alignment = v
}

func (o PaymentConsentPeriodicAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["interval"] = o.Interval
	}
	if true {
		toSerialize["alignment"] = o.Alignment
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentConsentPeriodicAmount struct {
	value *PaymentConsentPeriodicAmount
	isSet bool
}

func (v NullablePaymentConsentPeriodicAmount) Get() *PaymentConsentPeriodicAmount {
	return v.value
}

func (v *NullablePaymentConsentPeriodicAmount) Set(val *PaymentConsentPeriodicAmount) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentConsentPeriodicAmount) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentConsentPeriodicAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentConsentPeriodicAmount(val *PaymentConsentPeriodicAmount) *NullablePaymentConsentPeriodicAmount {
	return &NullablePaymentConsentPeriodicAmount{value: val, isSet: true}
}

func (v NullablePaymentConsentPeriodicAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentConsentPeriodicAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


