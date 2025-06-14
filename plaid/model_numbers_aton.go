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

// NumbersATON Identifying information for transferring holdings to an investments account via ATON.
type NumbersATON struct {
	// The Plaid account ID associated with the account numbers
	AccountId string `json:"account_id"`
	// The full account number for the account
	Account string `json:"account"`
	AdditionalProperties map[string]interface{}
}

type _NumbersATON NumbersATON

// NewNumbersATON instantiates a new NumbersATON object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbersATON(accountId string, account string) *NumbersATON {
	this := NumbersATON{}
	this.AccountId = accountId
	this.Account = account
	return &this
}

// NewNumbersATONWithDefaults instantiates a new NumbersATON object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersATONWithDefaults() *NumbersATON {
	this := NumbersATON{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NumbersATON) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NumbersATON) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NumbersATON) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccount returns the Account field value
func (o *NumbersATON) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NumbersATON) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NumbersATON) SetAccount(v string) {
	o.Account = v
}

func (o NumbersATON) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NumbersATON) UnmarshalJSON(bytes []byte) (err error) {
	varNumbersATON := _NumbersATON{}

	if err = json.Unmarshal(bytes, &varNumbersATON); err == nil {
		*o = NumbersATON(varNumbersATON)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNumbersATON struct {
	value *NumbersATON
	isSet bool
}

func (v NullableNumbersATON) Get() *NumbersATON {
	return v.value
}

func (v *NullableNumbersATON) Set(val *NumbersATON) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersATON) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersATON) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersATON(val *NumbersATON) *NullableNumbersATON {
	return &NullableNumbersATON{value: val, isSet: true}
}

func (v NullableNumbersATON) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersATON) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


