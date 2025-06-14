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

// MonitoringInsights An object representing the Monitoring Insights for the given Item
type MonitoringInsights struct {
	Income MonitoringIncomeInsights `json:"income"`
	Loans MonitoringLoanInsights `json:"loans"`
	AdditionalProperties map[string]interface{}
}

type _MonitoringInsights MonitoringInsights

// NewMonitoringInsights instantiates a new MonitoringInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringInsights(income MonitoringIncomeInsights, loans MonitoringLoanInsights) *MonitoringInsights {
	this := MonitoringInsights{}
	this.Income = income
	this.Loans = loans
	return &this
}

// NewMonitoringInsightsWithDefaults instantiates a new MonitoringInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringInsightsWithDefaults() *MonitoringInsights {
	this := MonitoringInsights{}
	return &this
}

// GetIncome returns the Income field value
func (o *MonitoringInsights) GetIncome() MonitoringIncomeInsights {
	if o == nil {
		var ret MonitoringIncomeInsights
		return ret
	}

	return o.Income
}

// GetIncomeOk returns a tuple with the Income field value
// and a boolean to check if the value has been set.
func (o *MonitoringInsights) GetIncomeOk() (*MonitoringIncomeInsights, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Income, true
}

// SetIncome sets field value
func (o *MonitoringInsights) SetIncome(v MonitoringIncomeInsights) {
	o.Income = v
}

// GetLoans returns the Loans field value
func (o *MonitoringInsights) GetLoans() MonitoringLoanInsights {
	if o == nil {
		var ret MonitoringLoanInsights
		return ret
	}

	return o.Loans
}

// GetLoansOk returns a tuple with the Loans field value
// and a boolean to check if the value has been set.
func (o *MonitoringInsights) GetLoansOk() (*MonitoringLoanInsights, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Loans, true
}

// SetLoans sets field value
func (o *MonitoringInsights) SetLoans(v MonitoringLoanInsights) {
	o.Loans = v
}

func (o MonitoringInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["income"] = o.Income
	}
	if true {
		toSerialize["loans"] = o.Loans
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MonitoringInsights) UnmarshalJSON(bytes []byte) (err error) {
	varMonitoringInsights := _MonitoringInsights{}

	if err = json.Unmarshal(bytes, &varMonitoringInsights); err == nil {
		*o = MonitoringInsights(varMonitoringInsights)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "income")
		delete(additionalProperties, "loans")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMonitoringInsights struct {
	value *MonitoringInsights
	isSet bool
}

func (v NullableMonitoringInsights) Get() *MonitoringInsights {
	return v.value
}

func (v *NullableMonitoringInsights) Set(val *MonitoringInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringInsights(val *MonitoringInsights) *NullableMonitoringInsights {
	return &NullableMonitoringInsights{value: val, isSet: true}
}

func (v NullableMonitoringInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


