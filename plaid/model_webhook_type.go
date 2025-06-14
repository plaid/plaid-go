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

// WebhookType The webhook types that can be fired by this test endpoint.
type WebhookType string

var _ = fmt.Printf

// List of WebhookType
const (
	WEBHOOKTYPE_AUTH WebhookType = "AUTH"
	WEBHOOKTYPE_HOLDINGS WebhookType = "HOLDINGS"
	WEBHOOKTYPE_INVESTMENTS_TRANSACTIONS WebhookType = "INVESTMENTS_TRANSACTIONS"
	WEBHOOKTYPE_ITEM WebhookType = "ITEM"
	WEBHOOKTYPE_LIABILITIES WebhookType = "LIABILITIES"
	WEBHOOKTYPE_TRANSACTIONS WebhookType = "TRANSACTIONS"
	WEBHOOKTYPE_ASSETS WebhookType = "ASSETS"
)

var allowedWebhookTypeEnumValues = []WebhookType{
	"AUTH",
	"HOLDINGS",
	"INVESTMENTS_TRANSACTIONS",
	"ITEM",
	"LIABILITIES",
	"TRANSACTIONS",
	"ASSETS",
}

func (v *WebhookType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := WebhookType(value)


	*v = enumTypeValue
	return nil
}

// NewWebhookTypeFromValue returns a pointer to a valid WebhookType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookTypeFromValue(v string) (*WebhookType, error) {
	ev := WebhookType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookType) IsValid() bool {
	for _, existing := range allowedWebhookTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookType value
func (v WebhookType) Ptr() *WebhookType {
	return &v
}

type NullableWebhookType struct {
	value *WebhookType
	isSet bool
}

func (v NullableWebhookType) Get() *WebhookType {
	return v.value
}

func (v *NullableWebhookType) Set(val *WebhookType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookType(val *WebhookType) *NullableWebhookType {
	return &NullableWebhookType{value: val, isSet: true}
}

func (v NullableWebhookType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

