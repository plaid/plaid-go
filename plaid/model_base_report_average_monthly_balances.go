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

// BaseReportAverageMonthlyBalances Average balance in dollar amount per month
type BaseReportAverageMonthlyBalances struct {
	// The start date of this time period. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	StartDate *string `json:"start_date,omitempty"`
	// The end date of this time period. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	EndDate *string `json:"end_date,omitempty"`
	AverageBalance *CreditAmountWithCurrency `json:"average_balance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseReportAverageMonthlyBalances BaseReportAverageMonthlyBalances

// NewBaseReportAverageMonthlyBalances instantiates a new BaseReportAverageMonthlyBalances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseReportAverageMonthlyBalances() *BaseReportAverageMonthlyBalances {
	this := BaseReportAverageMonthlyBalances{}
	return &this
}

// NewBaseReportAverageMonthlyBalancesWithDefaults instantiates a new BaseReportAverageMonthlyBalances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseReportAverageMonthlyBalancesWithDefaults() *BaseReportAverageMonthlyBalances {
	this := BaseReportAverageMonthlyBalances{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BaseReportAverageMonthlyBalances) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportAverageMonthlyBalances) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BaseReportAverageMonthlyBalances) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *BaseReportAverageMonthlyBalances) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BaseReportAverageMonthlyBalances) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportAverageMonthlyBalances) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BaseReportAverageMonthlyBalances) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *BaseReportAverageMonthlyBalances) SetEndDate(v string) {
	o.EndDate = &v
}

// GetAverageBalance returns the AverageBalance field value if set, zero value otherwise.
func (o *BaseReportAverageMonthlyBalances) GetAverageBalance() CreditAmountWithCurrency {
	if o == nil || o.AverageBalance == nil {
		var ret CreditAmountWithCurrency
		return ret
	}
	return *o.AverageBalance
}

// GetAverageBalanceOk returns a tuple with the AverageBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseReportAverageMonthlyBalances) GetAverageBalanceOk() (*CreditAmountWithCurrency, bool) {
	if o == nil || o.AverageBalance == nil {
		return nil, false
	}
	return o.AverageBalance, true
}

// HasAverageBalance returns a boolean if a field has been set.
func (o *BaseReportAverageMonthlyBalances) HasAverageBalance() bool {
	if o != nil && o.AverageBalance != nil {
		return true
	}

	return false
}

// SetAverageBalance gets a reference to the given CreditAmountWithCurrency and assigns it to the AverageBalance field.
func (o *BaseReportAverageMonthlyBalances) SetAverageBalance(v CreditAmountWithCurrency) {
	o.AverageBalance = &v
}

func (o BaseReportAverageMonthlyBalances) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.AverageBalance != nil {
		toSerialize["average_balance"] = o.AverageBalance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseReportAverageMonthlyBalances) UnmarshalJSON(bytes []byte) (err error) {
	varBaseReportAverageMonthlyBalances := _BaseReportAverageMonthlyBalances{}

	if err = json.Unmarshal(bytes, &varBaseReportAverageMonthlyBalances); err == nil {
		*o = BaseReportAverageMonthlyBalances(varBaseReportAverageMonthlyBalances)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "average_balance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseReportAverageMonthlyBalances struct {
	value *BaseReportAverageMonthlyBalances
	isSet bool
}

func (v NullableBaseReportAverageMonthlyBalances) Get() *BaseReportAverageMonthlyBalances {
	return v.value
}

func (v *NullableBaseReportAverageMonthlyBalances) Set(val *BaseReportAverageMonthlyBalances) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseReportAverageMonthlyBalances) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseReportAverageMonthlyBalances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseReportAverageMonthlyBalances(val *BaseReportAverageMonthlyBalances) *NullableBaseReportAverageMonthlyBalances {
	return &NullableBaseReportAverageMonthlyBalances{value: val, isSet: true}
}

func (v NullableBaseReportAverageMonthlyBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseReportAverageMonthlyBalances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

