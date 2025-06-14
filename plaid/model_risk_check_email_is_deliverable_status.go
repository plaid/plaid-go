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

// RiskCheckEmailIsDeliverableStatus SMTP-MX check to confirm the email address exists if known.
type RiskCheckEmailIsDeliverableStatus string

var _ = fmt.Printf

// List of RiskCheckEmailIsDeliverableStatus
const (
	RISKCHECKEMAILISDELIVERABLESTATUS_YES RiskCheckEmailIsDeliverableStatus = "yes"
	RISKCHECKEMAILISDELIVERABLESTATUS_NO RiskCheckEmailIsDeliverableStatus = "no"
	RISKCHECKEMAILISDELIVERABLESTATUS_NO_DATA RiskCheckEmailIsDeliverableStatus = "no_data"
)

var allowedRiskCheckEmailIsDeliverableStatusEnumValues = []RiskCheckEmailIsDeliverableStatus{
	"yes",
	"no",
	"no_data",
}

func (v *RiskCheckEmailIsDeliverableStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := RiskCheckEmailIsDeliverableStatus(value)


	*v = enumTypeValue
	return nil
}

// NewRiskCheckEmailIsDeliverableStatusFromValue returns a pointer to a valid RiskCheckEmailIsDeliverableStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRiskCheckEmailIsDeliverableStatusFromValue(v string) (*RiskCheckEmailIsDeliverableStatus, error) {
	ev := RiskCheckEmailIsDeliverableStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RiskCheckEmailIsDeliverableStatus) IsValid() bool {
	for _, existing := range allowedRiskCheckEmailIsDeliverableStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RiskCheckEmailIsDeliverableStatus value
func (v RiskCheckEmailIsDeliverableStatus) Ptr() *RiskCheckEmailIsDeliverableStatus {
	return &v
}

type NullableRiskCheckEmailIsDeliverableStatus struct {
	value *RiskCheckEmailIsDeliverableStatus
	isSet bool
}

func (v NullableRiskCheckEmailIsDeliverableStatus) Get() *RiskCheckEmailIsDeliverableStatus {
	return v.value
}

func (v *NullableRiskCheckEmailIsDeliverableStatus) Set(val *RiskCheckEmailIsDeliverableStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckEmailIsDeliverableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckEmailIsDeliverableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckEmailIsDeliverableStatus(val *RiskCheckEmailIsDeliverableStatus) *NullableRiskCheckEmailIsDeliverableStatus {
	return &NullableRiskCheckEmailIsDeliverableStatus{value: val, isSet: true}
}

func (v NullableRiskCheckEmailIsDeliverableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckEmailIsDeliverableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

