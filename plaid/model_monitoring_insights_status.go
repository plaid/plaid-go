/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.586.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// MonitoringInsightsStatus Enum for the status of the insights
type MonitoringInsightsStatus string

var _ = fmt.Printf

// List of MonitoringInsightsStatus
const (
	MONITORINGINSIGHTSSTATUS_AVAILABLE MonitoringInsightsStatus = "AVAILABLE"
	MONITORINGINSIGHTSSTATUS_FAILED MonitoringInsightsStatus = "FAILED"
	MONITORINGINSIGHTSSTATUS_PENDING MonitoringInsightsStatus = "PENDING"
	MONITORINGINSIGHTSSTATUS_UNSUPPORTED MonitoringInsightsStatus = "UNSUPPORTED"
	MONITORINGINSIGHTSSTATUS_UNHEALTHY MonitoringInsightsStatus = "UNHEALTHY"
)

var allowedMonitoringInsightsStatusEnumValues = []MonitoringInsightsStatus{
	"AVAILABLE",
	"FAILED",
	"PENDING",
	"UNSUPPORTED",
	"UNHEALTHY",
}

func (v *MonitoringInsightsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := MonitoringInsightsStatus(value)


	*v = enumTypeValue
	return nil
}

// NewMonitoringInsightsStatusFromValue returns a pointer to a valid MonitoringInsightsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitoringInsightsStatusFromValue(v string) (*MonitoringInsightsStatus, error) {
	ev := MonitoringInsightsStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitoringInsightsStatus) IsValid() bool {
	for _, existing := range allowedMonitoringInsightsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitoringInsightsStatus value
func (v MonitoringInsightsStatus) Ptr() *MonitoringInsightsStatus {
	return &v
}

type NullableMonitoringInsightsStatus struct {
	value *MonitoringInsightsStatus
	isSet bool
}

func (v NullableMonitoringInsightsStatus) Get() *MonitoringInsightsStatus {
	return v.value
}

func (v *NullableMonitoringInsightsStatus) Set(val *MonitoringInsightsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringInsightsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringInsightsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringInsightsStatus(val *MonitoringInsightsStatus) *NullableMonitoringInsightsStatus {
	return &NullableMonitoringInsightsStatus{value: val, isSet: true}
}

func (v NullableMonitoringInsightsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringInsightsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
