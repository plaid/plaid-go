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

// TotalInflowAmount30d Total amount of debit transactions into the account in the last 30 days. This field will be empty for non-depository accounts. This field only takes into account USD transactions from the account.
type TotalInflowAmount30d struct {
	// Value of amount with up to 2 decimal places.
	Amount float32 `json:"amount"`
	// The ISO 4217 currency code of the amount or balance.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _TotalInflowAmount30d TotalInflowAmount30d

// NewTotalInflowAmount30d instantiates a new TotalInflowAmount30d object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotalInflowAmount30d(amount float32, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString) *TotalInflowAmount30d {
	this := TotalInflowAmount30d{}
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	return &this
}

// NewTotalInflowAmount30dWithDefaults instantiates a new TotalInflowAmount30d object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotalInflowAmount30dWithDefaults() *TotalInflowAmount30d {
	this := TotalInflowAmount30d{}
	return &this
}

// GetAmount returns the Amount field value
func (o *TotalInflowAmount30d) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TotalInflowAmount30d) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TotalInflowAmount30d) SetAmount(v float32) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TotalInflowAmount30d) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TotalInflowAmount30d) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *TotalInflowAmount30d) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TotalInflowAmount30d) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TotalInflowAmount30d) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *TotalInflowAmount30d) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

func (o TotalInflowAmount30d) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TotalInflowAmount30d) UnmarshalJSON(bytes []byte) (err error) {
	varTotalInflowAmount30d := _TotalInflowAmount30d{}

	if err = json.Unmarshal(bytes, &varTotalInflowAmount30d); err == nil {
		*o = TotalInflowAmount30d(varTotalInflowAmount30d)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTotalInflowAmount30d struct {
	value *TotalInflowAmount30d
	isSet bool
}

func (v NullableTotalInflowAmount30d) Get() *TotalInflowAmount30d {
	return v.value
}

func (v *NullableTotalInflowAmount30d) Set(val *TotalInflowAmount30d) {
	v.value = val
	v.isSet = true
}

func (v NullableTotalInflowAmount30d) IsSet() bool {
	return v.isSet
}

func (v *NullableTotalInflowAmount30d) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotalInflowAmount30d(val *TotalInflowAmount30d) *NullableTotalInflowAmount30d {
	return &NullableTotalInflowAmount30d{value: val, isSet: true}
}

func (v NullableTotalInflowAmount30d) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotalInflowAmount30d) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


