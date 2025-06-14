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

// DateRange A date range with a start and end date
type DateRange struct {
	// A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).
	Beginning string `json:"beginning"`
	// A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).
	Ending string `json:"ending"`
	AdditionalProperties map[string]interface{}
}

type _DateRange DateRange

// NewDateRange instantiates a new DateRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateRange(beginning string, ending string) *DateRange {
	this := DateRange{}
	this.Beginning = beginning
	this.Ending = ending
	return &this
}

// NewDateRangeWithDefaults instantiates a new DateRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateRangeWithDefaults() *DateRange {
	this := DateRange{}
	return &this
}

// GetBeginning returns the Beginning field value
func (o *DateRange) GetBeginning() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Beginning
}

// GetBeginningOk returns a tuple with the Beginning field value
// and a boolean to check if the value has been set.
func (o *DateRange) GetBeginningOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Beginning, true
}

// SetBeginning sets field value
func (o *DateRange) SetBeginning(v string) {
	o.Beginning = v
}

// GetEnding returns the Ending field value
func (o *DateRange) GetEnding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ending
}

// GetEndingOk returns a tuple with the Ending field value
// and a boolean to check if the value has been set.
func (o *DateRange) GetEndingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ending, true
}

// SetEnding sets field value
func (o *DateRange) SetEnding(v string) {
	o.Ending = v
}

func (o DateRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["beginning"] = o.Beginning
	}
	if true {
		toSerialize["ending"] = o.Ending
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DateRange) UnmarshalJSON(bytes []byte) (err error) {
	varDateRange := _DateRange{}

	if err = json.Unmarshal(bytes, &varDateRange); err == nil {
		*o = DateRange(varDateRange)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "beginning")
		delete(additionalProperties, "ending")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDateRange struct {
	value *DateRange
	isSet bool
}

func (v NullableDateRange) Get() *DateRange {
	return v.value
}

func (v *NullableDateRange) Set(val *DateRange) {
	v.value = val
	v.isSet = true
}

func (v NullableDateRange) IsSet() bool {
	return v.isSet
}

func (v *NullableDateRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateRange(val *DateRange) *NullableDateRange {
	return &NullableDateRange{value: val, isSet: true}
}

func (v NullableDateRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


