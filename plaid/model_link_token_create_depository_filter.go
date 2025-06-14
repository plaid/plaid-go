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

// LinkTokenCreateDepositoryFilter A filter to apply to `depository`-type accounts
type LinkTokenCreateDepositoryFilter struct {
	// An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#account-type-schema). 
	AccountSubtypes *[]DepositoryAccountSubtype `json:"account_subtypes,omitempty"`
}

// NewLinkTokenCreateDepositoryFilter instantiates a new LinkTokenCreateDepositoryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateDepositoryFilter() *LinkTokenCreateDepositoryFilter {
	this := LinkTokenCreateDepositoryFilter{}
	return &this
}

// NewLinkTokenCreateDepositoryFilterWithDefaults instantiates a new LinkTokenCreateDepositoryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateDepositoryFilterWithDefaults() *LinkTokenCreateDepositoryFilter {
	this := LinkTokenCreateDepositoryFilter{}
	return &this
}

// GetAccountSubtypes returns the AccountSubtypes field value if set, zero value otherwise.
func (o *LinkTokenCreateDepositoryFilter) GetAccountSubtypes() []DepositoryAccountSubtype {
	if o == nil || o.AccountSubtypes == nil {
		var ret []DepositoryAccountSubtype
		return ret
	}
	return *o.AccountSubtypes
}

// GetAccountSubtypesOk returns a tuple with the AccountSubtypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateDepositoryFilter) GetAccountSubtypesOk() (*[]DepositoryAccountSubtype, bool) {
	if o == nil || o.AccountSubtypes == nil {
		return nil, false
	}
	return o.AccountSubtypes, true
}

// HasAccountSubtypes returns a boolean if a field has been set.
func (o *LinkTokenCreateDepositoryFilter) HasAccountSubtypes() bool {
	if o != nil && o.AccountSubtypes != nil {
		return true
	}

	return false
}

// SetAccountSubtypes gets a reference to the given []DepositoryAccountSubtype and assigns it to the AccountSubtypes field.
func (o *LinkTokenCreateDepositoryFilter) SetAccountSubtypes(v []DepositoryAccountSubtype) {
	o.AccountSubtypes = &v
}

func (o LinkTokenCreateDepositoryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSubtypes != nil {
		toSerialize["account_subtypes"] = o.AccountSubtypes
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateDepositoryFilter struct {
	value *LinkTokenCreateDepositoryFilter
	isSet bool
}

func (v NullableLinkTokenCreateDepositoryFilter) Get() *LinkTokenCreateDepositoryFilter {
	return v.value
}

func (v *NullableLinkTokenCreateDepositoryFilter) Set(val *LinkTokenCreateDepositoryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateDepositoryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateDepositoryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateDepositoryFilter(val *LinkTokenCreateDepositoryFilter) *NullableLinkTokenCreateDepositoryFilter {
	return &NullableLinkTokenCreateDepositoryFilter{value: val, isSet: true}
}

func (v NullableLinkTokenCreateDepositoryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateDepositoryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


