/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// MatchSummaryCode An enum indicating the match type between data provided by user and data checked against an external data source.   `match` indicates that the provided input data was a strong match against external data.  `partial_match` indicates the data approximately matched against external data. For example, \"Knope\" vs. \"Knope-Wyatt\" for last name.  `no_match` indicates that Plaid was able to perform a check against an external data source and it did not match the provided input data.  `no_data` indicates that Plaid was unable to find external data to compare against the provided input data.  `no_input` indicates that Plaid was unable to perform a check because no information was provided for this field by the end user.
type MatchSummaryCode string

// List of MatchSummaryCode
const (
	MATCHSUMMARYCODE_MATCH MatchSummaryCode = "match"
	MATCHSUMMARYCODE_PARTIAL_MATCH MatchSummaryCode = "partial_match"
	MATCHSUMMARYCODE_NO_MATCH MatchSummaryCode = "no_match"
	MATCHSUMMARYCODE_NO_DATA MatchSummaryCode = "no_data"
	MATCHSUMMARYCODE_NO_INPUT MatchSummaryCode = "no_input"
)

var allowedMatchSummaryCodeEnumValues = []MatchSummaryCode{
	"match",
	"partial_match",
	"no_match",
	"no_data",
	"no_input",
}

func (v *MatchSummaryCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MatchSummaryCode(value)
	for _, existing := range allowedMatchSummaryCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MatchSummaryCode", value)
}

// NewMatchSummaryCodeFromValue returns a pointer to a valid MatchSummaryCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMatchSummaryCodeFromValue(v string) (*MatchSummaryCode, error) {
	ev := MatchSummaryCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MatchSummaryCode: valid values are %v", v, allowedMatchSummaryCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MatchSummaryCode) IsValid() bool {
	for _, existing := range allowedMatchSummaryCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MatchSummaryCode value
func (v MatchSummaryCode) Ptr() *MatchSummaryCode {
	return &v
}

type NullableMatchSummaryCode struct {
	value *MatchSummaryCode
	isSet bool
}

func (v NullableMatchSummaryCode) Get() *MatchSummaryCode {
	return v.value
}

func (v *NullableMatchSummaryCode) Set(val *MatchSummaryCode) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchSummaryCode) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchSummaryCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchSummaryCode(val *MatchSummaryCode) *NullableMatchSummaryCode {
	return &NullableMatchSummaryCode{value: val, isSet: true}
}

func (v NullableMatchSummaryCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchSummaryCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
