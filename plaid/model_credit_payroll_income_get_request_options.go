/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.503.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditPayrollIncomeGetRequestOptions An optional object for `/credit/payroll_income/get` request options.
type CreditPayrollIncomeGetRequestOptions struct {
	// An array of `item_id`s whose payroll information is returned. Each `item_id` should uniquely identify a payroll income item. If this field is not provided, all `item_id`s associated with the `user_token` will returned in the response.
	ItemIds *[]string `json:"item_ids,omitempty"`
}

// NewCreditPayrollIncomeGetRequestOptions instantiates a new CreditPayrollIncomeGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayrollIncomeGetRequestOptions() *CreditPayrollIncomeGetRequestOptions {
	this := CreditPayrollIncomeGetRequestOptions{}
	return &this
}

// NewCreditPayrollIncomeGetRequestOptionsWithDefaults instantiates a new CreditPayrollIncomeGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayrollIncomeGetRequestOptionsWithDefaults() *CreditPayrollIncomeGetRequestOptions {
	this := CreditPayrollIncomeGetRequestOptions{}
	return &this
}

// GetItemIds returns the ItemIds field value if set, zero value otherwise.
func (o *CreditPayrollIncomeGetRequestOptions) GetItemIds() []string {
	if o == nil || o.ItemIds == nil {
		var ret []string
		return ret
	}
	return *o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeGetRequestOptions) GetItemIdsOk() (*[]string, bool) {
	if o == nil || o.ItemIds == nil {
		return nil, false
	}
	return o.ItemIds, true
}

// HasItemIds returns a boolean if a field has been set.
func (o *CreditPayrollIncomeGetRequestOptions) HasItemIds() bool {
	if o != nil && o.ItemIds != nil {
		return true
	}

	return false
}

// SetItemIds gets a reference to the given []string and assigns it to the ItemIds field.
func (o *CreditPayrollIncomeGetRequestOptions) SetItemIds(v []string) {
	o.ItemIds = &v
}

func (o CreditPayrollIncomeGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemIds != nil {
		toSerialize["item_ids"] = o.ItemIds
	}
	return json.Marshal(toSerialize)
}

type NullableCreditPayrollIncomeGetRequestOptions struct {
	value *CreditPayrollIncomeGetRequestOptions
	isSet bool
}

func (v NullableCreditPayrollIncomeGetRequestOptions) Get() *CreditPayrollIncomeGetRequestOptions {
	return v.value
}

func (v *NullableCreditPayrollIncomeGetRequestOptions) Set(val *CreditPayrollIncomeGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayrollIncomeGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayrollIncomeGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayrollIncomeGetRequestOptions(val *CreditPayrollIncomeGetRequestOptions) *NullableCreditPayrollIncomeGetRequestOptions {
	return &NullableCreditPayrollIncomeGetRequestOptions{value: val, isSet: true}
}

func (v NullableCreditPayrollIncomeGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayrollIncomeGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

