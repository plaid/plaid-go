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

// HistoricalAnnualIncome An object representing the historical annual income amount.
type HistoricalAnnualIncome struct {
	// The historical annual income at the time of subscription
	BaselineAmount float32 `json:"baseline_amount"`
	// The current historical annual income
	CurrentAmount float32 `json:"current_amount"`
	AdditionalProperties map[string]interface{}
}

type _HistoricalAnnualIncome HistoricalAnnualIncome

// NewHistoricalAnnualIncome instantiates a new HistoricalAnnualIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalAnnualIncome(baselineAmount float32, currentAmount float32) *HistoricalAnnualIncome {
	this := HistoricalAnnualIncome{}
	this.BaselineAmount = baselineAmount
	this.CurrentAmount = currentAmount
	return &this
}

// NewHistoricalAnnualIncomeWithDefaults instantiates a new HistoricalAnnualIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalAnnualIncomeWithDefaults() *HistoricalAnnualIncome {
	this := HistoricalAnnualIncome{}
	return &this
}

// GetBaselineAmount returns the BaselineAmount field value
func (o *HistoricalAnnualIncome) GetBaselineAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BaselineAmount
}

// GetBaselineAmountOk returns a tuple with the BaselineAmount field value
// and a boolean to check if the value has been set.
func (o *HistoricalAnnualIncome) GetBaselineAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BaselineAmount, true
}

// SetBaselineAmount sets field value
func (o *HistoricalAnnualIncome) SetBaselineAmount(v float32) {
	o.BaselineAmount = v
}

// GetCurrentAmount returns the CurrentAmount field value
func (o *HistoricalAnnualIncome) GetCurrentAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentAmount
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value
// and a boolean to check if the value has been set.
func (o *HistoricalAnnualIncome) GetCurrentAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentAmount, true
}

// SetCurrentAmount sets field value
func (o *HistoricalAnnualIncome) SetCurrentAmount(v float32) {
	o.CurrentAmount = v
}

func (o HistoricalAnnualIncome) MarshalJSON() ([]byte, error) {
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

func (o *HistoricalAnnualIncome) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalAnnualIncome := _HistoricalAnnualIncome{}

	if err = json.Unmarshal(bytes, &varHistoricalAnnualIncome); err == nil {
		*o = HistoricalAnnualIncome(varHistoricalAnnualIncome)
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

type NullableHistoricalAnnualIncome struct {
	value *HistoricalAnnualIncome
	isSet bool
}

func (v NullableHistoricalAnnualIncome) Get() *HistoricalAnnualIncome {
	return v.value
}

func (v *NullableHistoricalAnnualIncome) Set(val *HistoricalAnnualIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalAnnualIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalAnnualIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalAnnualIncome(val *HistoricalAnnualIncome) *NullableHistoricalAnnualIncome {
	return &NullableHistoricalAnnualIncome{value: val, isSet: true}
}

func (v NullableHistoricalAnnualIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalAnnualIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


