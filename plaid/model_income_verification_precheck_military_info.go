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

// IncomeVerificationPrecheckMilitaryInfo Data about military info in the income verification precheck.
type IncomeVerificationPrecheckMilitaryInfo struct {
	// Is the user currently active duty in the US military
	IsActiveDuty NullableBool `json:"is_active_duty,omitempty"`
	// If the user is currently serving in the US military, the branch of the military in which they are serving Valid values: 'AIR FORCE', 'ARMY', 'COAST GUARD', 'MARINES', 'NAVY', 'UNKNOWN'
	Branch NullableString `json:"branch,omitempty"`
}

// NewIncomeVerificationPrecheckMilitaryInfo instantiates a new IncomeVerificationPrecheckMilitaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationPrecheckMilitaryInfo() *IncomeVerificationPrecheckMilitaryInfo {
	this := IncomeVerificationPrecheckMilitaryInfo{}
	return &this
}

// NewIncomeVerificationPrecheckMilitaryInfoWithDefaults instantiates a new IncomeVerificationPrecheckMilitaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationPrecheckMilitaryInfoWithDefaults() *IncomeVerificationPrecheckMilitaryInfo {
	this := IncomeVerificationPrecheckMilitaryInfo{}
	return &this
}

// GetIsActiveDuty returns the IsActiveDuty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckMilitaryInfo) GetIsActiveDuty() bool {
	if o == nil || o.IsActiveDuty.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsActiveDuty.Get()
}

// GetIsActiveDutyOk returns a tuple with the IsActiveDuty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckMilitaryInfo) GetIsActiveDutyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsActiveDuty.Get(), o.IsActiveDuty.IsSet()
}

// HasIsActiveDuty returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckMilitaryInfo) HasIsActiveDuty() bool {
	if o != nil && o.IsActiveDuty.IsSet() {
		return true
	}

	return false
}

// SetIsActiveDuty gets a reference to the given NullableBool and assigns it to the IsActiveDuty field.
func (o *IncomeVerificationPrecheckMilitaryInfo) SetIsActiveDuty(v bool) {
	o.IsActiveDuty.Set(&v)
}
// SetIsActiveDutyNil sets the value for IsActiveDuty to be an explicit nil
func (o *IncomeVerificationPrecheckMilitaryInfo) SetIsActiveDutyNil() {
	o.IsActiveDuty.Set(nil)
}

// UnsetIsActiveDuty ensures that no value is present for IsActiveDuty, not even an explicit nil
func (o *IncomeVerificationPrecheckMilitaryInfo) UnsetIsActiveDuty() {
	o.IsActiveDuty.Unset()
}

// GetBranch returns the Branch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeVerificationPrecheckMilitaryInfo) GetBranch() string {
	if o == nil || o.Branch.Get() == nil {
		var ret string
		return ret
	}
	return *o.Branch.Get()
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeVerificationPrecheckMilitaryInfo) GetBranchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Branch.Get(), o.Branch.IsSet()
}

// HasBranch returns a boolean if a field has been set.
func (o *IncomeVerificationPrecheckMilitaryInfo) HasBranch() bool {
	if o != nil && o.Branch.IsSet() {
		return true
	}

	return false
}

// SetBranch gets a reference to the given NullableString and assigns it to the Branch field.
func (o *IncomeVerificationPrecheckMilitaryInfo) SetBranch(v string) {
	o.Branch.Set(&v)
}
// SetBranchNil sets the value for Branch to be an explicit nil
func (o *IncomeVerificationPrecheckMilitaryInfo) SetBranchNil() {
	o.Branch.Set(nil)
}

// UnsetBranch ensures that no value is present for Branch, not even an explicit nil
func (o *IncomeVerificationPrecheckMilitaryInfo) UnsetBranch() {
	o.Branch.Unset()
}

func (o IncomeVerificationPrecheckMilitaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsActiveDuty.IsSet() {
		toSerialize["is_active_duty"] = o.IsActiveDuty.Get()
	}
	if o.Branch.IsSet() {
		toSerialize["branch"] = o.Branch.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationPrecheckMilitaryInfo struct {
	value *IncomeVerificationPrecheckMilitaryInfo
	isSet bool
}

func (v NullableIncomeVerificationPrecheckMilitaryInfo) Get() *IncomeVerificationPrecheckMilitaryInfo {
	return v.value
}

func (v *NullableIncomeVerificationPrecheckMilitaryInfo) Set(val *IncomeVerificationPrecheckMilitaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationPrecheckMilitaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationPrecheckMilitaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationPrecheckMilitaryInfo(val *IncomeVerificationPrecheckMilitaryInfo) *NullableIncomeVerificationPrecheckMilitaryInfo {
	return &NullableIncomeVerificationPrecheckMilitaryInfo{value: val, isSet: true}
}

func (v NullableIncomeVerificationPrecheckMilitaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationPrecheckMilitaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


