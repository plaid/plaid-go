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

// BaseReportNumberFlowInsights The number of credits or debits out of the account. This field will only be included for depository accounts.
type BaseReportNumberFlowInsights struct {
	// The start date of this time period. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	StartDate string `json:"start_date"`
	// The end date of this time period. The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	EndDate string `json:"end_date"`
	// The number of credits or debits out of the account for this time period.
	Count int32 `json:"count"`
	AdditionalProperties map[string]interface{}
}

type _BaseReportNumberFlowInsights BaseReportNumberFlowInsights

// NewBaseReportNumberFlowInsights instantiates a new BaseReportNumberFlowInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseReportNumberFlowInsights(startDate string, endDate string, count int32) *BaseReportNumberFlowInsights {
	this := BaseReportNumberFlowInsights{}
	this.StartDate = startDate
	this.EndDate = endDate
	this.Count = count
	return &this
}

// NewBaseReportNumberFlowInsightsWithDefaults instantiates a new BaseReportNumberFlowInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseReportNumberFlowInsightsWithDefaults() *BaseReportNumberFlowInsights {
	this := BaseReportNumberFlowInsights{}
	return &this
}

// GetStartDate returns the StartDate field value
func (o *BaseReportNumberFlowInsights) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *BaseReportNumberFlowInsights) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *BaseReportNumberFlowInsights) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *BaseReportNumberFlowInsights) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *BaseReportNumberFlowInsights) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *BaseReportNumberFlowInsights) SetEndDate(v string) {
	o.EndDate = v
}

// GetCount returns the Count field value
func (o *BaseReportNumberFlowInsights) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *BaseReportNumberFlowInsights) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *BaseReportNumberFlowInsights) SetCount(v int32) {
	o.Count = v
}

func (o BaseReportNumberFlowInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if true {
		toSerialize["end_date"] = o.EndDate
	}
	if true {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseReportNumberFlowInsights) UnmarshalJSON(bytes []byte) (err error) {
	varBaseReportNumberFlowInsights := _BaseReportNumberFlowInsights{}

	if err = json.Unmarshal(bytes, &varBaseReportNumberFlowInsights); err == nil {
		*o = BaseReportNumberFlowInsights(varBaseReportNumberFlowInsights)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseReportNumberFlowInsights struct {
	value *BaseReportNumberFlowInsights
	isSet bool
}

func (v NullableBaseReportNumberFlowInsights) Get() *BaseReportNumberFlowInsights {
	return v.value
}

func (v *NullableBaseReportNumberFlowInsights) Set(val *BaseReportNumberFlowInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseReportNumberFlowInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseReportNumberFlowInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseReportNumberFlowInsights(val *BaseReportNumberFlowInsights) *NullableBaseReportNumberFlowInsights {
	return &NullableBaseReportNumberFlowInsights{value: val, isSet: true}
}

func (v NullableBaseReportNumberFlowInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseReportNumberFlowInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


