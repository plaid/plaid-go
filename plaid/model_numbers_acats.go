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

// NumbersACATS Identifying information for transferring holdings to an investments account via ACATS.
type NumbersACATS struct {
	// The Plaid account ID associated with the account numbers
	AccountId string `json:"account_id"`
	// The full account number for the account
	Account string `json:"account"`
	// Identifiers for the clearinghouses that are associated with the account in order of relevance. If this array is empty, call `/institutions/get_by_id` with the `item.institution_id` to get the DTC number.
	DtcNumbers []string `json:"dtc_numbers"`
	AdditionalProperties map[string]interface{}
}

type _NumbersACATS NumbersACATS

// NewNumbersACATS instantiates a new NumbersACATS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbersACATS(accountId string, account string, dtcNumbers []string) *NumbersACATS {
	this := NumbersACATS{}
	this.AccountId = accountId
	this.Account = account
	this.DtcNumbers = dtcNumbers
	return &this
}

// NewNumbersACATSWithDefaults instantiates a new NumbersACATS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersACATSWithDefaults() *NumbersACATS {
	this := NumbersACATS{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NumbersACATS) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NumbersACATS) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NumbersACATS) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccount returns the Account field value
func (o *NumbersACATS) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NumbersACATS) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NumbersACATS) SetAccount(v string) {
	o.Account = v
}

// GetDtcNumbers returns the DtcNumbers field value
func (o *NumbersACATS) GetDtcNumbers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DtcNumbers
}

// GetDtcNumbersOk returns a tuple with the DtcNumbers field value
// and a boolean to check if the value has been set.
func (o *NumbersACATS) GetDtcNumbersOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtcNumbers, true
}

// SetDtcNumbers sets field value
func (o *NumbersACATS) SetDtcNumbers(v []string) {
	o.DtcNumbers = v
}

func (o NumbersACATS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["dtc_numbers"] = o.DtcNumbers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NumbersACATS) UnmarshalJSON(bytes []byte) (err error) {
	varNumbersACATS := _NumbersACATS{}

	if err = json.Unmarshal(bytes, &varNumbersACATS); err == nil {
		*o = NumbersACATS(varNumbersACATS)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "account")
		delete(additionalProperties, "dtc_numbers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNumbersACATS struct {
	value *NumbersACATS
	isSet bool
}

func (v NullableNumbersACATS) Get() *NumbersACATS {
	return v.value
}

func (v *NullableNumbersACATS) Set(val *NumbersACATS) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersACATS) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersACATS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersACATS(val *NumbersACATS) *NullableNumbersACATS {
	return &NullableNumbersACATS{value: val, isSet: true}
}

func (v NullableNumbersACATS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersACATS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


