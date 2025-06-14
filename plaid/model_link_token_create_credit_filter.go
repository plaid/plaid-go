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
)

// LinkTokenCreateCreditFilter A filter to apply to `credit`-type accounts
type LinkTokenCreateCreditFilter struct {
	// An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema). 
	AccountSubtypes *[]CreditAccountSubtype `json:"account_subtypes,omitempty"`
}

// NewLinkTokenCreateCreditFilter instantiates a new LinkTokenCreateCreditFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateCreditFilter() *LinkTokenCreateCreditFilter {
	this := LinkTokenCreateCreditFilter{}
	return &this
}

// NewLinkTokenCreateCreditFilterWithDefaults instantiates a new LinkTokenCreateCreditFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateCreditFilterWithDefaults() *LinkTokenCreateCreditFilter {
	this := LinkTokenCreateCreditFilter{}
	return &this
}

// GetAccountSubtypes returns the AccountSubtypes field value if set, zero value otherwise.
func (o *LinkTokenCreateCreditFilter) GetAccountSubtypes() []CreditAccountSubtype {
	if o == nil || o.AccountSubtypes == nil {
		var ret []CreditAccountSubtype
		return ret
	}
	return *o.AccountSubtypes
}

// GetAccountSubtypesOk returns a tuple with the AccountSubtypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateCreditFilter) GetAccountSubtypesOk() (*[]CreditAccountSubtype, bool) {
	if o == nil || o.AccountSubtypes == nil {
		return nil, false
	}
	return o.AccountSubtypes, true
}

// HasAccountSubtypes returns a boolean if a field has been set.
func (o *LinkTokenCreateCreditFilter) HasAccountSubtypes() bool {
	if o != nil && o.AccountSubtypes != nil {
		return true
	}

	return false
}

// SetAccountSubtypes gets a reference to the given []CreditAccountSubtype and assigns it to the AccountSubtypes field.
func (o *LinkTokenCreateCreditFilter) SetAccountSubtypes(v []CreditAccountSubtype) {
	o.AccountSubtypes = &v
}

func (o LinkTokenCreateCreditFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSubtypes != nil {
		toSerialize["account_subtypes"] = o.AccountSubtypes
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateCreditFilter struct {
	value *LinkTokenCreateCreditFilter
	isSet bool
}

func (v NullableLinkTokenCreateCreditFilter) Get() *LinkTokenCreateCreditFilter {
	return v.value
}

func (v *NullableLinkTokenCreateCreditFilter) Set(val *LinkTokenCreateCreditFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateCreditFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateCreditFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateCreditFilter(val *LinkTokenCreateCreditFilter) *NullableLinkTokenCreateCreditFilter {
	return &NullableLinkTokenCreateCreditFilter{value: val, isSet: true}
}

func (v NullableLinkTokenCreateCreditFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateCreditFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


