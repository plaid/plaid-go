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

// NumbersBACSNullable Identifying information for transferring money to or from a UK bank account via BACS.
type NumbersBACSNullable struct {
	// The Plaid account ID associated with the account numbers
	AccountId string `json:"account_id"`
	// The BACS account number for the account
	Account string `json:"account"`
	// The BACS sort code for the account
	SortCode string `json:"sort_code"`
}

// NewNumbersBACSNullable instantiates a new NumbersBACSNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbersBACSNullable(accountId string, account string, sortCode string) *NumbersBACSNullable {
	this := NumbersBACSNullable{}
	this.AccountId = accountId
	this.Account = account
	this.SortCode = sortCode
	return &this
}

// NewNumbersBACSNullableWithDefaults instantiates a new NumbersBACSNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersBACSNullableWithDefaults() *NumbersBACSNullable {
	this := NumbersBACSNullable{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NumbersBACSNullable) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NumbersBACSNullable) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NumbersBACSNullable) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccount returns the Account field value
func (o *NumbersBACSNullable) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NumbersBACSNullable) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NumbersBACSNullable) SetAccount(v string) {
	o.Account = v
}

// GetSortCode returns the SortCode field value
func (o *NumbersBACSNullable) GetSortCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value
// and a boolean to check if the value has been set.
func (o *NumbersBACSNullable) GetSortCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SortCode, true
}

// SetSortCode sets field value
func (o *NumbersBACSNullable) SetSortCode(v string) {
	o.SortCode = v
}

func (o NumbersBACSNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["sort_code"] = o.SortCode
	}
	return json.Marshal(toSerialize)
}

type NullableNumbersBACSNullable struct {
	value *NumbersBACSNullable
	isSet bool
}

func (v NullableNumbersBACSNullable) Get() *NumbersBACSNullable {
	return v.value
}

func (v *NullableNumbersBACSNullable) Set(val *NumbersBACSNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersBACSNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersBACSNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersBACSNullable(val *NumbersBACSNullable) *NullableNumbersBACSNullable {
	return &NullableNumbersBACSNullable{value: val, isSet: true}
}

func (v NullableNumbersBACSNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersBACSNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


