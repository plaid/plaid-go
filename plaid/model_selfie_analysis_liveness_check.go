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

// SelfieAnalysisLivenessCheck Assessment of whether the selfie capture is of a real human being, as opposed to a picture of a human on a screen, a picture of a paper cut out, someone wearing a mask, or a deepfake.
type SelfieAnalysisLivenessCheck string

var _ = fmt.Printf

// List of SelfieAnalysisLivenessCheck
const (
	SELFIEANALYSISLIVENESSCHECK_SUCCESS SelfieAnalysisLivenessCheck = "success"
	SELFIEANALYSISLIVENESSCHECK_FAILED SelfieAnalysisLivenessCheck = "failed"
)

var allowedSelfieAnalysisLivenessCheckEnumValues = []SelfieAnalysisLivenessCheck{
	"success",
	"failed",
}

func (v *SelfieAnalysisLivenessCheck) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := SelfieAnalysisLivenessCheck(value)


	*v = enumTypeValue
	return nil
}

// NewSelfieAnalysisLivenessCheckFromValue returns a pointer to a valid SelfieAnalysisLivenessCheck
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSelfieAnalysisLivenessCheckFromValue(v string) (*SelfieAnalysisLivenessCheck, error) {
	ev := SelfieAnalysisLivenessCheck(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelfieAnalysisLivenessCheck) IsValid() bool {
	for _, existing := range allowedSelfieAnalysisLivenessCheckEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SelfieAnalysisLivenessCheck value
func (v SelfieAnalysisLivenessCheck) Ptr() *SelfieAnalysisLivenessCheck {
	return &v
}

type NullableSelfieAnalysisLivenessCheck struct {
	value *SelfieAnalysisLivenessCheck
	isSet bool
}

func (v NullableSelfieAnalysisLivenessCheck) Get() *SelfieAnalysisLivenessCheck {
	return v.value
}

func (v *NullableSelfieAnalysisLivenessCheck) Set(val *SelfieAnalysisLivenessCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfieAnalysisLivenessCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfieAnalysisLivenessCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfieAnalysisLivenessCheck(val *SelfieAnalysisLivenessCheck) *NullableSelfieAnalysisLivenessCheck {
	return &NullableSelfieAnalysisLivenessCheck{value: val, isSet: true}
}

func (v NullableSelfieAnalysisLivenessCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfieAnalysisLivenessCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

