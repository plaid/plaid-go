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

// SandboxIncomeWebhookFireRequestWebhookCode The webhook codes that can be fired by this test endpoint.
type SandboxIncomeWebhookFireRequestWebhookCode string

var _ = fmt.Printf

// List of SandboxIncomeWebhookFireRequestWebhookCode
const (
	SANDBOXINCOMEWEBHOOKFIREREQUESTWEBHOOKCODE_VERIFICATION SandboxIncomeWebhookFireRequestWebhookCode = "INCOME_VERIFICATION"
	SANDBOXINCOMEWEBHOOKFIREREQUESTWEBHOOKCODE_VERIFICATION_RISK_SIGNALS SandboxIncomeWebhookFireRequestWebhookCode = "INCOME_VERIFICATION_RISK_SIGNALS"
)

var allowedSandboxIncomeWebhookFireRequestWebhookCodeEnumValues = []SandboxIncomeWebhookFireRequestWebhookCode{
	"INCOME_VERIFICATION",
	"INCOME_VERIFICATION_RISK_SIGNALS",
}

func (v *SandboxIncomeWebhookFireRequestWebhookCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := SandboxIncomeWebhookFireRequestWebhookCode(value)


	*v = enumTypeValue
	return nil
}

// NewSandboxIncomeWebhookFireRequestWebhookCodeFromValue returns a pointer to a valid SandboxIncomeWebhookFireRequestWebhookCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSandboxIncomeWebhookFireRequestWebhookCodeFromValue(v string) (*SandboxIncomeWebhookFireRequestWebhookCode, error) {
	ev := SandboxIncomeWebhookFireRequestWebhookCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SandboxIncomeWebhookFireRequestWebhookCode) IsValid() bool {
	for _, existing := range allowedSandboxIncomeWebhookFireRequestWebhookCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SandboxIncomeWebhookFireRequestWebhookCode value
func (v SandboxIncomeWebhookFireRequestWebhookCode) Ptr() *SandboxIncomeWebhookFireRequestWebhookCode {
	return &v
}

type NullableSandboxIncomeWebhookFireRequestWebhookCode struct {
	value *SandboxIncomeWebhookFireRequestWebhookCode
	isSet bool
}

func (v NullableSandboxIncomeWebhookFireRequestWebhookCode) Get() *SandboxIncomeWebhookFireRequestWebhookCode {
	return v.value
}

func (v *NullableSandboxIncomeWebhookFireRequestWebhookCode) Set(val *SandboxIncomeWebhookFireRequestWebhookCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxIncomeWebhookFireRequestWebhookCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxIncomeWebhookFireRequestWebhookCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxIncomeWebhookFireRequestWebhookCode(val *SandboxIncomeWebhookFireRequestWebhookCode) *NullableSandboxIncomeWebhookFireRequestWebhookCode {
	return &NullableSandboxIncomeWebhookFireRequestWebhookCode{value: val, isSet: true}
}

func (v NullableSandboxIncomeWebhookFireRequestWebhookCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxIncomeWebhookFireRequestWebhookCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
