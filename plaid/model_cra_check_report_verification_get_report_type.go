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

// CraCheckReportVerificationGetReportType Type of verification report.
type CraCheckReportVerificationGetReportType string

var _ = fmt.Printf

// List of CraCheckReportVerificationGetReportType
const (
	CRACHECKREPORTVERIFICATIONGETREPORTTYPE_VOA CraCheckReportVerificationGetReportType = "VOA"
	CRACHECKREPORTVERIFICATIONGETREPORTTYPE_VOE CraCheckReportVerificationGetReportType = "VOE"
)

var allowedCraCheckReportVerificationGetReportTypeEnumValues = []CraCheckReportVerificationGetReportType{
	"VOA",
	"VOE",
}

func (v *CraCheckReportVerificationGetReportType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CraCheckReportVerificationGetReportType(value)


	*v = enumTypeValue
	return nil
}

// NewCraCheckReportVerificationGetReportTypeFromValue returns a pointer to a valid CraCheckReportVerificationGetReportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCraCheckReportVerificationGetReportTypeFromValue(v string) (*CraCheckReportVerificationGetReportType, error) {
	ev := CraCheckReportVerificationGetReportType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CraCheckReportVerificationGetReportType) IsValid() bool {
	for _, existing := range allowedCraCheckReportVerificationGetReportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CraCheckReportVerificationGetReportType value
func (v CraCheckReportVerificationGetReportType) Ptr() *CraCheckReportVerificationGetReportType {
	return &v
}

type NullableCraCheckReportVerificationGetReportType struct {
	value *CraCheckReportVerificationGetReportType
	isSet bool
}

func (v NullableCraCheckReportVerificationGetReportType) Get() *CraCheckReportVerificationGetReportType {
	return v.value
}

func (v *NullableCraCheckReportVerificationGetReportType) Set(val *CraCheckReportVerificationGetReportType) {
	v.value = val
	v.isSet = true
}

func (v NullableCraCheckReportVerificationGetReportType) IsSet() bool {
	return v.isSet
}

func (v *NullableCraCheckReportVerificationGetReportType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraCheckReportVerificationGetReportType(val *CraCheckReportVerificationGetReportType) *NullableCraCheckReportVerificationGetReportType {
	return &NullableCraCheckReportVerificationGetReportType{value: val, isSet: true}
}

func (v NullableCraCheckReportVerificationGetReportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraCheckReportVerificationGetReportType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

