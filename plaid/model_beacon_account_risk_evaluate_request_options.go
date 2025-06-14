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

// BeaconAccountRiskEvaluateRequestOptions An optional object to filter `/beacon/account_risk/v1/evaluate` results to a subset of the accounts on the linked Item.
type BeaconAccountRiskEvaluateRequestOptions struct {
	// An array of `account_ids` for the specific accounts to evaluate.
	AccountIds *[]string `json:"account_ids,omitempty"`
}

// NewBeaconAccountRiskEvaluateRequestOptions instantiates a new BeaconAccountRiskEvaluateRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeaconAccountRiskEvaluateRequestOptions() *BeaconAccountRiskEvaluateRequestOptions {
	this := BeaconAccountRiskEvaluateRequestOptions{}
	return &this
}

// NewBeaconAccountRiskEvaluateRequestOptionsWithDefaults instantiates a new BeaconAccountRiskEvaluateRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeaconAccountRiskEvaluateRequestOptionsWithDefaults() *BeaconAccountRiskEvaluateRequestOptions {
	this := BeaconAccountRiskEvaluateRequestOptions{}
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *BeaconAccountRiskEvaluateRequestOptions) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeaconAccountRiskEvaluateRequestOptions) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *BeaconAccountRiskEvaluateRequestOptions) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *BeaconAccountRiskEvaluateRequestOptions) SetAccountIds(v []string) {
	o.AccountIds = &v
}

func (o BeaconAccountRiskEvaluateRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}
	return json.Marshal(toSerialize)
}

type NullableBeaconAccountRiskEvaluateRequestOptions struct {
	value *BeaconAccountRiskEvaluateRequestOptions
	isSet bool
}

func (v NullableBeaconAccountRiskEvaluateRequestOptions) Get() *BeaconAccountRiskEvaluateRequestOptions {
	return v.value
}

func (v *NullableBeaconAccountRiskEvaluateRequestOptions) Set(val *BeaconAccountRiskEvaluateRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconAccountRiskEvaluateRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconAccountRiskEvaluateRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconAccountRiskEvaluateRequestOptions(val *BeaconAccountRiskEvaluateRequestOptions) *NullableBeaconAccountRiskEvaluateRequestOptions {
	return &NullableBeaconAccountRiskEvaluateRequestOptions{value: val, isSet: true}
}

func (v NullableBeaconAccountRiskEvaluateRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconAccountRiskEvaluateRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


