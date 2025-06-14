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

// PaymentAmountToRefund The amount and currency of a payment
type PaymentAmountToRefund struct {
	Currency PaymentAmountCurrency `json:"currency"`
	// The amount of the payment. Must contain at most two digits of precision e.g. `1.23`.
	Value float64 `json:"value"`
}

// NewPaymentAmountToRefund instantiates a new PaymentAmountToRefund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAmountToRefund(currency PaymentAmountCurrency, value float64) *PaymentAmountToRefund {
	this := PaymentAmountToRefund{}
	this.Currency = currency
	this.Value = value
	return &this
}

// NewPaymentAmountToRefundWithDefaults instantiates a new PaymentAmountToRefund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAmountToRefundWithDefaults() *PaymentAmountToRefund {
	this := PaymentAmountToRefund{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *PaymentAmountToRefund) GetCurrency() PaymentAmountCurrency {
	if o == nil {
		var ret PaymentAmountCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentAmountToRefund) GetCurrencyOk() (*PaymentAmountCurrency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentAmountToRefund) SetCurrency(v PaymentAmountCurrency) {
	o.Currency = v
}

// GetValue returns the Value field value
func (o *PaymentAmountToRefund) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PaymentAmountToRefund) GetValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PaymentAmountToRefund) SetValue(v float64) {
	o.Value = v
}

func (o PaymentAmountToRefund) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentAmountToRefund struct {
	value *PaymentAmountToRefund
	isSet bool
}

func (v NullablePaymentAmountToRefund) Get() *PaymentAmountToRefund {
	return v.value
}

func (v *NullablePaymentAmountToRefund) Set(val *PaymentAmountToRefund) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAmountToRefund) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAmountToRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAmountToRefund(val *PaymentAmountToRefund) *NullablePaymentAmountToRefund {
	return &NullablePaymentAmountToRefund{value: val, isSet: true}
}

func (v NullablePaymentAmountToRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAmountToRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


