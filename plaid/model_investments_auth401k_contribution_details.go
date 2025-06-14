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

// InvestmentsAuth401kContributionDetails Object containing information on contribution transactions for the 401k account. Note that the sum fields in this object represent the total of absolute contribution values.
type InvestmentsAuth401kContributionDetails struct {
	// A list of the most recent contribution transactions for the 401k account. Includes all contributions made on the same day.
	LastContributionTransactions []InvestmentTransaction `json:"last_contribution_transactions"`
	// Number of contribution transactions on this account, for the past month.
	ContributionCount1m int32 `json:"contribution_count_1m"`
	// Sum of the contribution transactions on this account, for the past month.
	ContributionAmount1m float32 `json:"contribution_amount_1m"`
	// Number of contribution transactions on this account, for the past 6 months.
	ContributionCount6m int32 `json:"contribution_count_6m"`
	// Sum of the contribution transactions on this account, for the past 6 months.
	ContributionAmount6m float32 `json:"contribution_amount_6m"`
	// Number of contribution transactions on this account, for the past 12 months.
	ContributionCount12m int32 `json:"contribution_count_12m"`
	// Sum of the contribution transactions on this account, for the past 12 months.
	ContributionAmount12m float32 `json:"contribution_amount_12m"`
	AdditionalProperties map[string]interface{}
}

type _InvestmentsAuth401kContributionDetails InvestmentsAuth401kContributionDetails

// NewInvestmentsAuth401kContributionDetails instantiates a new InvestmentsAuth401kContributionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsAuth401kContributionDetails(lastContributionTransactions []InvestmentTransaction, contributionCount1m int32, contributionAmount1m float32, contributionCount6m int32, contributionAmount6m float32, contributionCount12m int32, contributionAmount12m float32) *InvestmentsAuth401kContributionDetails {
	this := InvestmentsAuth401kContributionDetails{}
	this.LastContributionTransactions = lastContributionTransactions
	this.ContributionCount1m = contributionCount1m
	this.ContributionAmount1m = contributionAmount1m
	this.ContributionCount6m = contributionCount6m
	this.ContributionAmount6m = contributionAmount6m
	this.ContributionCount12m = contributionCount12m
	this.ContributionAmount12m = contributionAmount12m
	return &this
}

// NewInvestmentsAuth401kContributionDetailsWithDefaults instantiates a new InvestmentsAuth401kContributionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsAuth401kContributionDetailsWithDefaults() *InvestmentsAuth401kContributionDetails {
	this := InvestmentsAuth401kContributionDetails{}
	return &this
}

// GetLastContributionTransactions returns the LastContributionTransactions field value
func (o *InvestmentsAuth401kContributionDetails) GetLastContributionTransactions() []InvestmentTransaction {
	if o == nil {
		var ret []InvestmentTransaction
		return ret
	}

	return o.LastContributionTransactions
}

// GetLastContributionTransactionsOk returns a tuple with the LastContributionTransactions field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuth401kContributionDetails) GetLastContributionTransactionsOk() (*[]InvestmentTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastContributionTransactions, true
}

// SetLastContributionTransactions sets field value
func (o *InvestmentsAuth401kContributionDetails) SetLastContributionTransactions(v []InvestmentTransaction) {
	o.LastContributionTransactions = v
}

// GetContributionCount1m returns the ContributionCount1m field value
func (o *InvestmentsAuth401kContributionDetails) GetContributionCount1m() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContributionCount1m
}

// GetContributionCount1mOk returns a tuple with the ContributionCount1m field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuth401kContributionDetails) GetContributionCount1mOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContributionCount1m, true
}

// SetContributionCount1m sets field value
func (o *InvestmentsAuth401kContributionDetails) SetContributionCount1m(v int32) {
	o.ContributionCount1m = v
}

// GetContributionAmount1m returns the ContributionAmount1m field value
func (o *InvestmentsAuth401kContributionDetails) GetContributionAmount1m() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ContributionAmount1m
}

// GetContributionAmount1mOk returns a tuple with the ContributionAmount1m field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuth401kContributionDetails) GetContributionAmount1mOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContributionAmount1m, true
}

// SetContributionAmount1m sets field value
func (o *InvestmentsAuth401kContributionDetails) SetContributionAmount1m(v float32) {
	o.ContributionAmount1m = v
}

// GetContributionCount6m returns the ContributionCount6m field value
func (o *InvestmentsAuth401kContributionDetails) GetContributionCount6m() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContributionCount6m
}

// GetContributionCount6mOk returns a tuple with the ContributionCount6m field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuth401kContributionDetails) GetContributionCount6mOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContributionCount6m, true
}

// SetContributionCount6m sets field value
func (o *InvestmentsAuth401kContributionDetails) SetContributionCount6m(v int32) {
	o.ContributionCount6m = v
}

// GetContributionAmount6m returns the ContributionAmount6m field value
func (o *InvestmentsAuth401kContributionDetails) GetContributionAmount6m() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ContributionAmount6m
}

// GetContributionAmount6mOk returns a tuple with the ContributionAmount6m field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuth401kContributionDetails) GetContributionAmount6mOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContributionAmount6m, true
}

// SetContributionAmount6m sets field value
func (o *InvestmentsAuth401kContributionDetails) SetContributionAmount6m(v float32) {
	o.ContributionAmount6m = v
}

// GetContributionCount12m returns the ContributionCount12m field value
func (o *InvestmentsAuth401kContributionDetails) GetContributionCount12m() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ContributionCount12m
}

// GetContributionCount12mOk returns a tuple with the ContributionCount12m field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuth401kContributionDetails) GetContributionCount12mOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContributionCount12m, true
}

// SetContributionCount12m sets field value
func (o *InvestmentsAuth401kContributionDetails) SetContributionCount12m(v int32) {
	o.ContributionCount12m = v
}

// GetContributionAmount12m returns the ContributionAmount12m field value
func (o *InvestmentsAuth401kContributionDetails) GetContributionAmount12m() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ContributionAmount12m
}

// GetContributionAmount12mOk returns a tuple with the ContributionAmount12m field value
// and a boolean to check if the value has been set.
func (o *InvestmentsAuth401kContributionDetails) GetContributionAmount12mOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContributionAmount12m, true
}

// SetContributionAmount12m sets field value
func (o *InvestmentsAuth401kContributionDetails) SetContributionAmount12m(v float32) {
	o.ContributionAmount12m = v
}

func (o InvestmentsAuth401kContributionDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["last_contribution_transactions"] = o.LastContributionTransactions
	}
	if true {
		toSerialize["contribution_count_1m"] = o.ContributionCount1m
	}
	if true {
		toSerialize["contribution_amount_1m"] = o.ContributionAmount1m
	}
	if true {
		toSerialize["contribution_count_6m"] = o.ContributionCount6m
	}
	if true {
		toSerialize["contribution_amount_6m"] = o.ContributionAmount6m
	}
	if true {
		toSerialize["contribution_count_12m"] = o.ContributionCount12m
	}
	if true {
		toSerialize["contribution_amount_12m"] = o.ContributionAmount12m
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InvestmentsAuth401kContributionDetails) UnmarshalJSON(bytes []byte) (err error) {
	varInvestmentsAuth401kContributionDetails := _InvestmentsAuth401kContributionDetails{}

	if err = json.Unmarshal(bytes, &varInvestmentsAuth401kContributionDetails); err == nil {
		*o = InvestmentsAuth401kContributionDetails(varInvestmentsAuth401kContributionDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "last_contribution_transactions")
		delete(additionalProperties, "contribution_count_1m")
		delete(additionalProperties, "contribution_amount_1m")
		delete(additionalProperties, "contribution_count_6m")
		delete(additionalProperties, "contribution_amount_6m")
		delete(additionalProperties, "contribution_count_12m")
		delete(additionalProperties, "contribution_amount_12m")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvestmentsAuth401kContributionDetails struct {
	value *InvestmentsAuth401kContributionDetails
	isSet bool
}

func (v NullableInvestmentsAuth401kContributionDetails) Get() *InvestmentsAuth401kContributionDetails {
	return v.value
}

func (v *NullableInvestmentsAuth401kContributionDetails) Set(val *InvestmentsAuth401kContributionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsAuth401kContributionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsAuth401kContributionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsAuth401kContributionDetails(val *InvestmentsAuth401kContributionDetails) *NullableInvestmentsAuth401kContributionDetails {
	return &NullableInvestmentsAuth401kContributionDetails{value: val, isSet: true}
}

func (v NullableInvestmentsAuth401kContributionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsAuth401kContributionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


