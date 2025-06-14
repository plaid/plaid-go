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

// InvestmentsAuthAccountDetails401k Additional account fee and contribution information for 401k type accounts.
type InvestmentsAuthAccountDetails401k struct {
	// The ID of the 401k account.
	AccountId *string `json:"account_id,omitempty"`
	FeeDetails *InvestmentsAuth401kFeeDetails `json:"fee_details,omitempty"`
	ContributionDetails *InvestmentsAuth401kContributionDetails `json:"contribution_details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InvestmentsAuthAccountDetails401k InvestmentsAuthAccountDetails401k

// NewInvestmentsAuthAccountDetails401k instantiates a new InvestmentsAuthAccountDetails401k object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsAuthAccountDetails401k() *InvestmentsAuthAccountDetails401k {
	this := InvestmentsAuthAccountDetails401k{}
	return &this
}

// NewInvestmentsAuthAccountDetails401kWithDefaults instantiates a new InvestmentsAuthAccountDetails401k object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsAuthAccountDetails401kWithDefaults() *InvestmentsAuthAccountDetails401k {
	this := InvestmentsAuthAccountDetails401k{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *InvestmentsAuthAccountDetails401k) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthAccountDetails401k) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *InvestmentsAuthAccountDetails401k) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *InvestmentsAuthAccountDetails401k) SetAccountId(v string) {
	o.AccountId = &v
}

// GetFeeDetails returns the FeeDetails field value if set, zero value otherwise.
func (o *InvestmentsAuthAccountDetails401k) GetFeeDetails() InvestmentsAuth401kFeeDetails {
	if o == nil || o.FeeDetails == nil {
		var ret InvestmentsAuth401kFeeDetails
		return ret
	}
	return *o.FeeDetails
}

// GetFeeDetailsOk returns a tuple with the FeeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthAccountDetails401k) GetFeeDetailsOk() (*InvestmentsAuth401kFeeDetails, bool) {
	if o == nil || o.FeeDetails == nil {
		return nil, false
	}
	return o.FeeDetails, true
}

// HasFeeDetails returns a boolean if a field has been set.
func (o *InvestmentsAuthAccountDetails401k) HasFeeDetails() bool {
	if o != nil && o.FeeDetails != nil {
		return true
	}

	return false
}

// SetFeeDetails gets a reference to the given InvestmentsAuth401kFeeDetails and assigns it to the FeeDetails field.
func (o *InvestmentsAuthAccountDetails401k) SetFeeDetails(v InvestmentsAuth401kFeeDetails) {
	o.FeeDetails = &v
}

// GetContributionDetails returns the ContributionDetails field value if set, zero value otherwise.
func (o *InvestmentsAuthAccountDetails401k) GetContributionDetails() InvestmentsAuth401kContributionDetails {
	if o == nil || o.ContributionDetails == nil {
		var ret InvestmentsAuth401kContributionDetails
		return ret
	}
	return *o.ContributionDetails
}

// GetContributionDetailsOk returns a tuple with the ContributionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsAuthAccountDetails401k) GetContributionDetailsOk() (*InvestmentsAuth401kContributionDetails, bool) {
	if o == nil || o.ContributionDetails == nil {
		return nil, false
	}
	return o.ContributionDetails, true
}

// HasContributionDetails returns a boolean if a field has been set.
func (o *InvestmentsAuthAccountDetails401k) HasContributionDetails() bool {
	if o != nil && o.ContributionDetails != nil {
		return true
	}

	return false
}

// SetContributionDetails gets a reference to the given InvestmentsAuth401kContributionDetails and assigns it to the ContributionDetails field.
func (o *InvestmentsAuthAccountDetails401k) SetContributionDetails(v InvestmentsAuth401kContributionDetails) {
	o.ContributionDetails = &v
}

func (o InvestmentsAuthAccountDetails401k) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.FeeDetails != nil {
		toSerialize["fee_details"] = o.FeeDetails
	}
	if o.ContributionDetails != nil {
		toSerialize["contribution_details"] = o.ContributionDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InvestmentsAuthAccountDetails401k) UnmarshalJSON(bytes []byte) (err error) {
	varInvestmentsAuthAccountDetails401k := _InvestmentsAuthAccountDetails401k{}

	if err = json.Unmarshal(bytes, &varInvestmentsAuthAccountDetails401k); err == nil {
		*o = InvestmentsAuthAccountDetails401k(varInvestmentsAuthAccountDetails401k)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "fee_details")
		delete(additionalProperties, "contribution_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvestmentsAuthAccountDetails401k struct {
	value *InvestmentsAuthAccountDetails401k
	isSet bool
}

func (v NullableInvestmentsAuthAccountDetails401k) Get() *InvestmentsAuthAccountDetails401k {
	return v.value
}

func (v *NullableInvestmentsAuthAccountDetails401k) Set(val *InvestmentsAuthAccountDetails401k) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsAuthAccountDetails401k) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsAuthAccountDetails401k) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsAuthAccountDetails401k(val *InvestmentsAuthAccountDetails401k) *NullableInvestmentsAuthAccountDetails401k {
	return &NullableInvestmentsAuthAccountDetails401k{value: val, isSet: true}
}

func (v NullableInvestmentsAuthAccountDetails401k) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsAuthAccountDetails401k) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


