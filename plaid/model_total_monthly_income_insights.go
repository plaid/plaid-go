/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.586.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TotalMonthlyIncomeInsights Details about about the total monthly income
type TotalMonthlyIncomeInsights struct {
	// The aggregated income for the 30 days prior to subscription date
	BaselineAmount float32 `json:"baseline_amount"`
	// The aggregated income of the last 30 days
	CurrentAmount float32 `json:"current_amount"`
	AdditionalProperties map[string]interface{}
}

type _TotalMonthlyIncomeInsights TotalMonthlyIncomeInsights

// NewTotalMonthlyIncomeInsights instantiates a new TotalMonthlyIncomeInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotalMonthlyIncomeInsights(baselineAmount float32, currentAmount float32) *TotalMonthlyIncomeInsights {
	this := TotalMonthlyIncomeInsights{}
	this.BaselineAmount = baselineAmount
	this.CurrentAmount = currentAmount
	return &this
}

// NewTotalMonthlyIncomeInsightsWithDefaults instantiates a new TotalMonthlyIncomeInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotalMonthlyIncomeInsightsWithDefaults() *TotalMonthlyIncomeInsights {
	this := TotalMonthlyIncomeInsights{}
	return &this
}

// GetBaselineAmount returns the BaselineAmount field value
func (o *TotalMonthlyIncomeInsights) GetBaselineAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BaselineAmount
}

// GetBaselineAmountOk returns a tuple with the BaselineAmount field value
// and a boolean to check if the value has been set.
func (o *TotalMonthlyIncomeInsights) GetBaselineAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BaselineAmount, true
}

// SetBaselineAmount sets field value
func (o *TotalMonthlyIncomeInsights) SetBaselineAmount(v float32) {
	o.BaselineAmount = v
}

// GetCurrentAmount returns the CurrentAmount field value
func (o *TotalMonthlyIncomeInsights) GetCurrentAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentAmount
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value
// and a boolean to check if the value has been set.
func (o *TotalMonthlyIncomeInsights) GetCurrentAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentAmount, true
}

// SetCurrentAmount sets field value
func (o *TotalMonthlyIncomeInsights) SetCurrentAmount(v float32) {
	o.CurrentAmount = v
}

func (o TotalMonthlyIncomeInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["baseline_amount"] = o.BaselineAmount
	}
	if true {
		toSerialize["current_amount"] = o.CurrentAmount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TotalMonthlyIncomeInsights) UnmarshalJSON(bytes []byte) (err error) {
	varTotalMonthlyIncomeInsights := _TotalMonthlyIncomeInsights{}

	if err = json.Unmarshal(bytes, &varTotalMonthlyIncomeInsights); err == nil {
		*o = TotalMonthlyIncomeInsights(varTotalMonthlyIncomeInsights)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "baseline_amount")
		delete(additionalProperties, "current_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTotalMonthlyIncomeInsights struct {
	value *TotalMonthlyIncomeInsights
	isSet bool
}

func (v NullableTotalMonthlyIncomeInsights) Get() *TotalMonthlyIncomeInsights {
	return v.value
}

func (v *NullableTotalMonthlyIncomeInsights) Set(val *TotalMonthlyIncomeInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableTotalMonthlyIncomeInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableTotalMonthlyIncomeInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotalMonthlyIncomeInsights(val *TotalMonthlyIncomeInsights) *NullableTotalMonthlyIncomeInsights {
	return &NullableTotalMonthlyIncomeInsights{value: val, isSet: true}
}

func (v NullableTotalMonthlyIncomeInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotalMonthlyIncomeInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

