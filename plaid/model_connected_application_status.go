/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.84.5
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// ConnectedApplicationStatus Describes if the connected item is active (i.e. has not been revoked or unlinked)
type ConnectedApplicationStatus string

// List of ConnectedApplicationStatus
const (
	CONNECTEDAPPLICATIONSTATUS_ACTIVE ConnectedApplicationStatus = "ACTIVE"
	CONNECTEDAPPLICATIONSTATUS_INACTIVE ConnectedApplicationStatus = "INACTIVE"
)

var allowedConnectedApplicationStatusEnumValues = []ConnectedApplicationStatus{
	"ACTIVE",
	"INACTIVE",
}

func (v *ConnectedApplicationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectedApplicationStatus(value)
	for _, existing := range allowedConnectedApplicationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectedApplicationStatus", value)
}

// NewConnectedApplicationStatusFromValue returns a pointer to a valid ConnectedApplicationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectedApplicationStatusFromValue(v string) (*ConnectedApplicationStatus, error) {
	ev := ConnectedApplicationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectedApplicationStatus: valid values are %v", v, allowedConnectedApplicationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectedApplicationStatus) IsValid() bool {
	for _, existing := range allowedConnectedApplicationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectedApplicationStatus value
func (v ConnectedApplicationStatus) Ptr() *ConnectedApplicationStatus {
	return &v
}

type NullableConnectedApplicationStatus struct {
	value *ConnectedApplicationStatus
	isSet bool
}

func (v NullableConnectedApplicationStatus) Get() *ConnectedApplicationStatus {
	return v.value
}

func (v *NullableConnectedApplicationStatus) Set(val *ConnectedApplicationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedApplicationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedApplicationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedApplicationStatus(val *ConnectedApplicationStatus) *NullableConnectedApplicationStatus {
	return &NullableConnectedApplicationStatus{value: val, isSet: true}
}

func (v NullableConnectedApplicationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedApplicationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
