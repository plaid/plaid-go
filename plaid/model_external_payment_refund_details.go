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

// ExternalPaymentRefundDetails Details about external payment refund
type ExternalPaymentRefundDetails struct {
	// The name of the account holder.
	Name string `json:"name"`
	// The International Bank Account Number (IBAN) for the account.
	Iban NullableString `json:"iban"`
	Bacs NullableRecipientBACSNullable `json:"bacs"`
}

// NewExternalPaymentRefundDetails instantiates a new ExternalPaymentRefundDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentRefundDetails(name string, iban NullableString, bacs NullableRecipientBACSNullable) *ExternalPaymentRefundDetails {
	this := ExternalPaymentRefundDetails{}
	this.Name = name
	this.Iban = iban
	this.Bacs = bacs
	return &this
}

// NewExternalPaymentRefundDetailsWithDefaults instantiates a new ExternalPaymentRefundDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentRefundDetailsWithDefaults() *ExternalPaymentRefundDetails {
	this := ExternalPaymentRefundDetails{}
	return &this
}

// GetName returns the Name field value
func (o *ExternalPaymentRefundDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalPaymentRefundDetails) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalPaymentRefundDetails) SetName(v string) {
	o.Name = v
}

// GetIban returns the Iban field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExternalPaymentRefundDetails) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}

	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentRefundDetails) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// SetIban sets field value
func (o *ExternalPaymentRefundDetails) SetIban(v string) {
	o.Iban.Set(&v)
}

// GetBacs returns the Bacs field value
// If the value is explicit nil, the zero value for RecipientBACSNullable will be returned
func (o *ExternalPaymentRefundDetails) GetBacs() RecipientBACSNullable {
	if o == nil || o.Bacs.Get() == nil {
		var ret RecipientBACSNullable
		return ret
	}

	return *o.Bacs.Get()
}

// GetBacsOk returns a tuple with the Bacs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalPaymentRefundDetails) GetBacsOk() (*RecipientBACSNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bacs.Get(), o.Bacs.IsSet()
}

// SetBacs sets field value
func (o *ExternalPaymentRefundDetails) SetBacs(v RecipientBACSNullable) {
	o.Bacs.Set(&v)
}

func (o ExternalPaymentRefundDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["iban"] = o.Iban.Get()
	}
	if true {
		toSerialize["bacs"] = o.Bacs.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPaymentRefundDetails struct {
	value *ExternalPaymentRefundDetails
	isSet bool
}

func (v NullableExternalPaymentRefundDetails) Get() *ExternalPaymentRefundDetails {
	return v.value
}

func (v *NullableExternalPaymentRefundDetails) Set(val *ExternalPaymentRefundDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentRefundDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentRefundDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentRefundDetails(val *ExternalPaymentRefundDetails) *NullableExternalPaymentRefundDetails {
	return &NullableExternalPaymentRefundDetails{value: val, isSet: true}
}

func (v NullableExternalPaymentRefundDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentRefundDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


