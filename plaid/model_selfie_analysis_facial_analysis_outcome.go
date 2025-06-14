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
	"fmt"
)

// SelfieAnalysisFacialAnalysisOutcome Outcome of the facial analysis for a specific facial feature.
type SelfieAnalysisFacialAnalysisOutcome string

var _ = fmt.Printf

// List of SelfieAnalysisFacialAnalysisOutcome
const (
	SELFIEANALYSISFACIALANALYSISOUTCOME_SUCCESS SelfieAnalysisFacialAnalysisOutcome = "success"
	SELFIEANALYSISFACIALANALYSISOUTCOME_FAILED SelfieAnalysisFacialAnalysisOutcome = "failed"
)

var allowedSelfieAnalysisFacialAnalysisOutcomeEnumValues = []SelfieAnalysisFacialAnalysisOutcome{
	"success",
	"failed",
}

func (v *SelfieAnalysisFacialAnalysisOutcome) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := SelfieAnalysisFacialAnalysisOutcome(value)


	*v = enumTypeValue
	return nil
}

// NewSelfieAnalysisFacialAnalysisOutcomeFromValue returns a pointer to a valid SelfieAnalysisFacialAnalysisOutcome
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelfieAnalysisFacialAnalysisOutcomeFromValue(v string) (*SelfieAnalysisFacialAnalysisOutcome, error) {
	ev := SelfieAnalysisFacialAnalysisOutcome(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelfieAnalysisFacialAnalysisOutcome) IsValid() bool {
	for _, existing := range allowedSelfieAnalysisFacialAnalysisOutcomeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelfieAnalysisFacialAnalysisOutcome value
func (v SelfieAnalysisFacialAnalysisOutcome) Ptr() *SelfieAnalysisFacialAnalysisOutcome {
	return &v
}

type NullableSelfieAnalysisFacialAnalysisOutcome struct {
	value *SelfieAnalysisFacialAnalysisOutcome
	isSet bool
}

func (v NullableSelfieAnalysisFacialAnalysisOutcome) Get() *SelfieAnalysisFacialAnalysisOutcome {
	return v.value
}

func (v *NullableSelfieAnalysisFacialAnalysisOutcome) Set(val *SelfieAnalysisFacialAnalysisOutcome) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfieAnalysisFacialAnalysisOutcome) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfieAnalysisFacialAnalysisOutcome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfieAnalysisFacialAnalysisOutcome(val *SelfieAnalysisFacialAnalysisOutcome) *NullableSelfieAnalysisFacialAnalysisOutcome {
	return &NullableSelfieAnalysisFacialAnalysisOutcome{value: val, isSet: true}
}

func (v NullableSelfieAnalysisFacialAnalysisOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfieAnalysisFacialAnalysisOutcome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

