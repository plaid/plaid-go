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

// VerificationRefreshStatus The verification refresh status. One of the following:  `\"VERIFICATION_REFRESH_STATUS_USER_PRESENCE_REQUIRED\"` User presence is required to refresh an income verification. `\"VERIFICATION_REFRESH_SUCCESSFUL\"` The income verification refresh was successful. `\"VERIFICATION_REFRESH_NOT_FOUND\"` No new data was found after the income verification refresh.
type VerificationRefreshStatus string

var _ = fmt.Printf

// List of VerificationRefreshStatus
const (
	VERIFICATIONREFRESHSTATUS_STATUS_USER_PRESENCE_REQUIRED VerificationRefreshStatus = "VERIFICATION_REFRESH_STATUS_USER_PRESENCE_REQUIRED"
	VERIFICATIONREFRESHSTATUS_SUCCESSFUL VerificationRefreshStatus = "VERIFICATION_REFRESH_SUCCESSFUL"
	VERIFICATIONREFRESHSTATUS_NOT_FOUND VerificationRefreshStatus = "VERIFICATION_REFRESH_NOT_FOUND"
)

var allowedVerificationRefreshStatusEnumValues = []VerificationRefreshStatus{
	"VERIFICATION_REFRESH_STATUS_USER_PRESENCE_REQUIRED",
	"VERIFICATION_REFRESH_SUCCESSFUL",
	"VERIFICATION_REFRESH_NOT_FOUND",
}

func (v *VerificationRefreshStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := VerificationRefreshStatus(value)


	*v = enumTypeValue
	return nil
}

// NewVerificationRefreshStatusFromValue returns a pointer to a valid VerificationRefreshStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerificationRefreshStatusFromValue(v string) (*VerificationRefreshStatus, error) {
	ev := VerificationRefreshStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerificationRefreshStatus) IsValid() bool {
	for _, existing := range allowedVerificationRefreshStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VerificationRefreshStatus value
func (v VerificationRefreshStatus) Ptr() *VerificationRefreshStatus {
	return &v
}

type NullableVerificationRefreshStatus struct {
	value *VerificationRefreshStatus
	isSet bool
}

func (v NullableVerificationRefreshStatus) Get() *VerificationRefreshStatus {
	return v.value
}

func (v *NullableVerificationRefreshStatus) Set(val *VerificationRefreshStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationRefreshStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationRefreshStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationRefreshStatus(val *VerificationRefreshStatus) *NullableVerificationRefreshStatus {
	return &NullableVerificationRefreshStatus{value: val, isSet: true}
}

func (v NullableVerificationRefreshStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationRefreshStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

