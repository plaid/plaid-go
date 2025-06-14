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

// CraVoaReportTransactionsInsights Transaction data associated with the account.
type CraVoaReportTransactionsInsights struct {
	// Transaction history associated with the account.
	AllTransactions []BaseReportTransaction `json:"all_transactions"`
	// The latest timeframe provided by the FI, in an ISO 8601 format (YYYY-MM-DD).
	EndDate NullableString `json:"end_date"`
	// The earliest timeframe provided by the FI, in an ISO 8601 format (YYYY-MM-DD).
	StartDate NullableString `json:"start_date"`
	AdditionalProperties map[string]interface{}
}

type _CraVoaReportTransactionsInsights CraVoaReportTransactionsInsights

// NewCraVoaReportTransactionsInsights instantiates a new CraVoaReportTransactionsInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraVoaReportTransactionsInsights(allTransactions []BaseReportTransaction, endDate NullableString, startDate NullableString) *CraVoaReportTransactionsInsights {
	this := CraVoaReportTransactionsInsights{}
	this.AllTransactions = allTransactions
	this.EndDate = endDate
	this.StartDate = startDate
	return &this
}

// NewCraVoaReportTransactionsInsightsWithDefaults instantiates a new CraVoaReportTransactionsInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraVoaReportTransactionsInsightsWithDefaults() *CraVoaReportTransactionsInsights {
	this := CraVoaReportTransactionsInsights{}
	return &this
}

// GetAllTransactions returns the AllTransactions field value
func (o *CraVoaReportTransactionsInsights) GetAllTransactions() []BaseReportTransaction {
	if o == nil {
		var ret []BaseReportTransaction
		return ret
	}

	return o.AllTransactions
}

// GetAllTransactionsOk returns a tuple with the AllTransactions field value
// and a boolean to check if the value has been set.
func (o *CraVoaReportTransactionsInsights) GetAllTransactionsOk() (*[]BaseReportTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AllTransactions, true
}

// SetAllTransactions sets field value
func (o *CraVoaReportTransactionsInsights) SetAllTransactions(v []BaseReportTransaction) {
	o.AllTransactions = v
}

// GetEndDate returns the EndDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CraVoaReportTransactionsInsights) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraVoaReportTransactionsInsights) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// SetEndDate sets field value
func (o *CraVoaReportTransactionsInsights) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// GetStartDate returns the StartDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CraVoaReportTransactionsInsights) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraVoaReportTransactionsInsights) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// SetStartDate sets field value
func (o *CraVoaReportTransactionsInsights) SetStartDate(v string) {
	o.StartDate.Set(&v)
}

func (o CraVoaReportTransactionsInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["all_transactions"] = o.AllTransactions
	}
	if true {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if true {
		toSerialize["start_date"] = o.StartDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraVoaReportTransactionsInsights) UnmarshalJSON(bytes []byte) (err error) {
	varCraVoaReportTransactionsInsights := _CraVoaReportTransactionsInsights{}

	if err = json.Unmarshal(bytes, &varCraVoaReportTransactionsInsights); err == nil {
		*o = CraVoaReportTransactionsInsights(varCraVoaReportTransactionsInsights)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "all_transactions")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "start_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraVoaReportTransactionsInsights struct {
	value *CraVoaReportTransactionsInsights
	isSet bool
}

func (v NullableCraVoaReportTransactionsInsights) Get() *CraVoaReportTransactionsInsights {
	return v.value
}

func (v *NullableCraVoaReportTransactionsInsights) Set(val *CraVoaReportTransactionsInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableCraVoaReportTransactionsInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableCraVoaReportTransactionsInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraVoaReportTransactionsInsights(val *CraVoaReportTransactionsInsights) *NullableCraVoaReportTransactionsInsights {
	return &NullableCraVoaReportTransactionsInsights{value: val, isSet: true}
}

func (v NullableCraVoaReportTransactionsInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraVoaReportTransactionsInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


