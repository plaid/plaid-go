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

// RiskCheckEmailTopLevelDomainIsSuspicious Indicates whether the email address top level domain, which is the last part of the domain, is fraudulent or risky if known. In most cases, a suspicious top level domain is also associated with a disposable or high-risk domain.
type RiskCheckEmailTopLevelDomainIsSuspicious string

var _ = fmt.Printf

// List of RiskCheckEmailTopLevelDomainIsSuspicious
const (
	RISKCHECKEMAILTOPLEVELDOMAINISSUSPICIOUS_YES RiskCheckEmailTopLevelDomainIsSuspicious = "yes"
	RISKCHECKEMAILTOPLEVELDOMAINISSUSPICIOUS_NO RiskCheckEmailTopLevelDomainIsSuspicious = "no"
	RISKCHECKEMAILTOPLEVELDOMAINISSUSPICIOUS_NO_DATA RiskCheckEmailTopLevelDomainIsSuspicious = "no_data"
)

var allowedRiskCheckEmailTopLevelDomainIsSuspiciousEnumValues = []RiskCheckEmailTopLevelDomainIsSuspicious{
	"yes",
	"no",
	"no_data",
}

func (v *RiskCheckEmailTopLevelDomainIsSuspicious) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := RiskCheckEmailTopLevelDomainIsSuspicious(value)


	*v = enumTypeValue
	return nil
}

// NewRiskCheckEmailTopLevelDomainIsSuspiciousFromValue returns a pointer to a valid RiskCheckEmailTopLevelDomainIsSuspicious
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRiskCheckEmailTopLevelDomainIsSuspiciousFromValue(v string) (*RiskCheckEmailTopLevelDomainIsSuspicious, error) {
	ev := RiskCheckEmailTopLevelDomainIsSuspicious(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RiskCheckEmailTopLevelDomainIsSuspicious) IsValid() bool {
	for _, existing := range allowedRiskCheckEmailTopLevelDomainIsSuspiciousEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RiskCheckEmailTopLevelDomainIsSuspicious value
func (v RiskCheckEmailTopLevelDomainIsSuspicious) Ptr() *RiskCheckEmailTopLevelDomainIsSuspicious {
	return &v
}

type NullableRiskCheckEmailTopLevelDomainIsSuspicious struct {
	value *RiskCheckEmailTopLevelDomainIsSuspicious
	isSet bool
}

func (v NullableRiskCheckEmailTopLevelDomainIsSuspicious) Get() *RiskCheckEmailTopLevelDomainIsSuspicious {
	return v.value
}

func (v *NullableRiskCheckEmailTopLevelDomainIsSuspicious) Set(val *RiskCheckEmailTopLevelDomainIsSuspicious) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckEmailTopLevelDomainIsSuspicious) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckEmailTopLevelDomainIsSuspicious) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckEmailTopLevelDomainIsSuspicious(val *RiskCheckEmailTopLevelDomainIsSuspicious) *NullableRiskCheckEmailTopLevelDomainIsSuspicious {
	return &NullableRiskCheckEmailTopLevelDomainIsSuspicious{value: val, isSet: true}
}

func (v NullableRiskCheckEmailTopLevelDomainIsSuspicious) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckEmailTopLevelDomainIsSuspicious) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

