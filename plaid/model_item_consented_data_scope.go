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

// ItemConsentedDataScope A data scope for the products that a user can consent to in [Data Transparency Messaging](/docs/link/data-transparency-messaging-migration-guide)
type ItemConsentedDataScope string

var _ = fmt.Printf

// List of ItemConsentedDataScope
const (
	ITEMCONSENTEDDATASCOPE_ACCOUNT_AND_BALANCE_INFO ItemConsentedDataScope = "account_and_balance_info"
	ITEMCONSENTEDDATASCOPE_CONTACT_INFO ItemConsentedDataScope = "contact_info"
	ITEMCONSENTEDDATASCOPE_ACCOUNT_AND_ROUTING_NUMBERS ItemConsentedDataScope = "account_and_routing_numbers"
	ITEMCONSENTEDDATASCOPE_TRANSACTIONS ItemConsentedDataScope = "transactions"
	ITEMCONSENTEDDATASCOPE_CREDIT_AND_LOANS ItemConsentedDataScope = "credit_and_loans"
	ITEMCONSENTEDDATASCOPE_INVESTMENTS ItemConsentedDataScope = "investments"
	ITEMCONSENTEDDATASCOPE_BANK_STATEMENTS ItemConsentedDataScope = "bank_statements"
	ITEMCONSENTEDDATASCOPE_RISK_INFO ItemConsentedDataScope = "risk_info"
)

var allowedItemConsentedDataScopeEnumValues = []ItemConsentedDataScope{
	"account_and_balance_info",
	"contact_info",
	"account_and_routing_numbers",
	"transactions",
	"credit_and_loans",
	"investments",
	"bank_statements",
	"risk_info",
}

func (v *ItemConsentedDataScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := ItemConsentedDataScope(value)


	*v = enumTypeValue
	return nil
}

// NewItemConsentedDataScopeFromValue returns a pointer to a valid ItemConsentedDataScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewItemConsentedDataScopeFromValue(v string) (*ItemConsentedDataScope, error) {
	ev := ItemConsentedDataScope(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ItemConsentedDataScope) IsValid() bool {
	for _, existing := range allowedItemConsentedDataScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ItemConsentedDataScope value
func (v ItemConsentedDataScope) Ptr() *ItemConsentedDataScope {
	return &v
}

type NullableItemConsentedDataScope struct {
	value *ItemConsentedDataScope
	isSet bool
}

func (v NullableItemConsentedDataScope) Get() *ItemConsentedDataScope {
	return v.value
}

func (v *NullableItemConsentedDataScope) Set(val *ItemConsentedDataScope) {
	v.value = val
	v.isSet = true
}

func (v NullableItemConsentedDataScope) IsSet() bool {
	return v.isSet
}

func (v *NullableItemConsentedDataScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemConsentedDataScope(val *ItemConsentedDataScope) *NullableItemConsentedDataScope {
	return &NullableItemConsentedDataScope{value: val, isSet: true}
}

func (v NullableItemConsentedDataScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemConsentedDataScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
