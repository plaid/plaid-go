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

// TransferMetricsGetAuthorizationUsage Details regarding authorization usage.
type TransferMetricsGetAuthorizationUsage struct {
	// The daily credit utilization formatted as a decimal.
	DailyCreditUtilization *string `json:"daily_credit_utilization,omitempty"`
	// The daily debit utilization formatted as a decimal.
	DailyDebitUtilization *string `json:"daily_debit_utilization,omitempty"`
	// The monthly credit utilization formatted as a decimal.
	MonthlyCreditUtilization *string `json:"monthly_credit_utilization,omitempty"`
	// The monthly debit utilization formatted as a decimal.
	MonthlyDebitUtilization *string `json:"monthly_debit_utilization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferMetricsGetAuthorizationUsage TransferMetricsGetAuthorizationUsage

// NewTransferMetricsGetAuthorizationUsage instantiates a new TransferMetricsGetAuthorizationUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferMetricsGetAuthorizationUsage() *TransferMetricsGetAuthorizationUsage {
	this := TransferMetricsGetAuthorizationUsage{}
	return &this
}

// NewTransferMetricsGetAuthorizationUsageWithDefaults instantiates a new TransferMetricsGetAuthorizationUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferMetricsGetAuthorizationUsageWithDefaults() *TransferMetricsGetAuthorizationUsage {
	this := TransferMetricsGetAuthorizationUsage{}
	return &this
}

// GetDailyCreditUtilization returns the DailyCreditUtilization field value if set, zero value otherwise.
func (o *TransferMetricsGetAuthorizationUsage) GetDailyCreditUtilization() string {
	if o == nil || o.DailyCreditUtilization == nil {
		var ret string
		return ret
	}
	return *o.DailyCreditUtilization
}

// GetDailyCreditUtilizationOk returns a tuple with the DailyCreditUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetAuthorizationUsage) GetDailyCreditUtilizationOk() (*string, bool) {
	if o == nil || o.DailyCreditUtilization == nil {
		return nil, false
	}
	return o.DailyCreditUtilization, true
}

// HasDailyCreditUtilization returns a boolean if a field has been set.
func (o *TransferMetricsGetAuthorizationUsage) HasDailyCreditUtilization() bool {
	if o != nil && o.DailyCreditUtilization != nil {
		return true
	}

	return false
}

// SetDailyCreditUtilization gets a reference to the given string and assigns it to the DailyCreditUtilization field.
func (o *TransferMetricsGetAuthorizationUsage) SetDailyCreditUtilization(v string) {
	o.DailyCreditUtilization = &v
}

// GetDailyDebitUtilization returns the DailyDebitUtilization field value if set, zero value otherwise.
func (o *TransferMetricsGetAuthorizationUsage) GetDailyDebitUtilization() string {
	if o == nil || o.DailyDebitUtilization == nil {
		var ret string
		return ret
	}
	return *o.DailyDebitUtilization
}

// GetDailyDebitUtilizationOk returns a tuple with the DailyDebitUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetAuthorizationUsage) GetDailyDebitUtilizationOk() (*string, bool) {
	if o == nil || o.DailyDebitUtilization == nil {
		return nil, false
	}
	return o.DailyDebitUtilization, true
}

// HasDailyDebitUtilization returns a boolean if a field has been set.
func (o *TransferMetricsGetAuthorizationUsage) HasDailyDebitUtilization() bool {
	if o != nil && o.DailyDebitUtilization != nil {
		return true
	}

	return false
}

// SetDailyDebitUtilization gets a reference to the given string and assigns it to the DailyDebitUtilization field.
func (o *TransferMetricsGetAuthorizationUsage) SetDailyDebitUtilization(v string) {
	o.DailyDebitUtilization = &v
}

// GetMonthlyCreditUtilization returns the MonthlyCreditUtilization field value if set, zero value otherwise.
func (o *TransferMetricsGetAuthorizationUsage) GetMonthlyCreditUtilization() string {
	if o == nil || o.MonthlyCreditUtilization == nil {
		var ret string
		return ret
	}
	return *o.MonthlyCreditUtilization
}

// GetMonthlyCreditUtilizationOk returns a tuple with the MonthlyCreditUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetAuthorizationUsage) GetMonthlyCreditUtilizationOk() (*string, bool) {
	if o == nil || o.MonthlyCreditUtilization == nil {
		return nil, false
	}
	return o.MonthlyCreditUtilization, true
}

// HasMonthlyCreditUtilization returns a boolean if a field has been set.
func (o *TransferMetricsGetAuthorizationUsage) HasMonthlyCreditUtilization() bool {
	if o != nil && o.MonthlyCreditUtilization != nil {
		return true
	}

	return false
}

// SetMonthlyCreditUtilization gets a reference to the given string and assigns it to the MonthlyCreditUtilization field.
func (o *TransferMetricsGetAuthorizationUsage) SetMonthlyCreditUtilization(v string) {
	o.MonthlyCreditUtilization = &v
}

// GetMonthlyDebitUtilization returns the MonthlyDebitUtilization field value if set, zero value otherwise.
func (o *TransferMetricsGetAuthorizationUsage) GetMonthlyDebitUtilization() string {
	if o == nil || o.MonthlyDebitUtilization == nil {
		var ret string
		return ret
	}
	return *o.MonthlyDebitUtilization
}

// GetMonthlyDebitUtilizationOk returns a tuple with the MonthlyDebitUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMetricsGetAuthorizationUsage) GetMonthlyDebitUtilizationOk() (*string, bool) {
	if o == nil || o.MonthlyDebitUtilization == nil {
		return nil, false
	}
	return o.MonthlyDebitUtilization, true
}

// HasMonthlyDebitUtilization returns a boolean if a field has been set.
func (o *TransferMetricsGetAuthorizationUsage) HasMonthlyDebitUtilization() bool {
	if o != nil && o.MonthlyDebitUtilization != nil {
		return true
	}

	return false
}

// SetMonthlyDebitUtilization gets a reference to the given string and assigns it to the MonthlyDebitUtilization field.
func (o *TransferMetricsGetAuthorizationUsage) SetMonthlyDebitUtilization(v string) {
	o.MonthlyDebitUtilization = &v
}

func (o TransferMetricsGetAuthorizationUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DailyCreditUtilization != nil {
		toSerialize["daily_credit_utilization"] = o.DailyCreditUtilization
	}
	if o.DailyDebitUtilization != nil {
		toSerialize["daily_debit_utilization"] = o.DailyDebitUtilization
	}
	if o.MonthlyCreditUtilization != nil {
		toSerialize["monthly_credit_utilization"] = o.MonthlyCreditUtilization
	}
	if o.MonthlyDebitUtilization != nil {
		toSerialize["monthly_debit_utilization"] = o.MonthlyDebitUtilization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferMetricsGetAuthorizationUsage) UnmarshalJSON(bytes []byte) (err error) {
	varTransferMetricsGetAuthorizationUsage := _TransferMetricsGetAuthorizationUsage{}

	if err = json.Unmarshal(bytes, &varTransferMetricsGetAuthorizationUsage); err == nil {
		*o = TransferMetricsGetAuthorizationUsage(varTransferMetricsGetAuthorizationUsage)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "daily_credit_utilization")
		delete(additionalProperties, "daily_debit_utilization")
		delete(additionalProperties, "monthly_credit_utilization")
		delete(additionalProperties, "monthly_debit_utilization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferMetricsGetAuthorizationUsage struct {
	value *TransferMetricsGetAuthorizationUsage
	isSet bool
}

func (v NullableTransferMetricsGetAuthorizationUsage) Get() *TransferMetricsGetAuthorizationUsage {
	return v.value
}

func (v *NullableTransferMetricsGetAuthorizationUsage) Set(val *TransferMetricsGetAuthorizationUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferMetricsGetAuthorizationUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferMetricsGetAuthorizationUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferMetricsGetAuthorizationUsage(val *TransferMetricsGetAuthorizationUsage) *NullableTransferMetricsGetAuthorizationUsage {
	return &NullableTransferMetricsGetAuthorizationUsage{value: val, isSet: true}
}

func (v NullableTransferMetricsGetAuthorizationUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferMetricsGetAuthorizationUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


