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
	"fmt"
)

// BaseReportWarningCode The warning code identifies a specific kind of warning. `OWNERS_UNAVAILABLE` indicates that account-owner information is not available. `TRANSACTIONS_UNAVAILABLE` indicates that transactions information associated with Credit and Depository accounts are unavailable.
type BaseReportWarningCode string

var _ = fmt.Printf

// List of BaseReportWarningCode
const (
	BASEREPORTWARNINGCODE_OWNERS_UNAVAILABLE BaseReportWarningCode = "OWNERS_UNAVAILABLE"
	BASEREPORTWARNINGCODE_TRANSACTIONS_UNAVAILABLE BaseReportWarningCode = "TRANSACTIONS_UNAVAILABLE"
)

var allowedBaseReportWarningCodeEnumValues = []BaseReportWarningCode{
	"OWNERS_UNAVAILABLE",
	"TRANSACTIONS_UNAVAILABLE",
}

func (v *BaseReportWarningCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := BaseReportWarningCode(value)


	*v = enumTypeValue
	return nil
}

// NewBaseReportWarningCodeFromValue returns a pointer to a valid BaseReportWarningCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBaseReportWarningCodeFromValue(v string) (*BaseReportWarningCode, error) {
	ev := BaseReportWarningCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BaseReportWarningCode) IsValid() bool {
	for _, existing := range allowedBaseReportWarningCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BaseReportWarningCode value
func (v BaseReportWarningCode) Ptr() *BaseReportWarningCode {
	return &v
}

type NullableBaseReportWarningCode struct {
	value *BaseReportWarningCode
	isSet bool
}

func (v NullableBaseReportWarningCode) Get() *BaseReportWarningCode {
	return v.value
}

func (v *NullableBaseReportWarningCode) Set(val *BaseReportWarningCode) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseReportWarningCode) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseReportWarningCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseReportWarningCode(val *BaseReportWarningCode) *NullableBaseReportWarningCode {
	return &NullableBaseReportWarningCode{value: val, isSet: true}
}

func (v NullableBaseReportWarningCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseReportWarningCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
