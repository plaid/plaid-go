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

// CraNetworkInsightsItem Contains data about the connected Item.
type CraNetworkInsightsItem struct {
	// The ID for the institution the user linked.
	InstitutionId string `json:"institution_id"`
	// The name of the institution the user linked.
	InstitutionName string `json:"institution_name"`
	// The identifier for the Item.
	ItemId string `json:"item_id"`
}

// NewCraNetworkInsightsItem instantiates a new CraNetworkInsightsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraNetworkInsightsItem(institutionId string, institutionName string, itemId string) *CraNetworkInsightsItem {
	this := CraNetworkInsightsItem{}
	this.InstitutionId = institutionId
	this.InstitutionName = institutionName
	this.ItemId = itemId
	return &this
}

// NewCraNetworkInsightsItemWithDefaults instantiates a new CraNetworkInsightsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraNetworkInsightsItemWithDefaults() *CraNetworkInsightsItem {
	this := CraNetworkInsightsItem{}
	return &this
}

// GetInstitutionId returns the InstitutionId field value
func (o *CraNetworkInsightsItem) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *CraNetworkInsightsItem) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *CraNetworkInsightsItem) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetInstitutionName returns the InstitutionName field value
func (o *CraNetworkInsightsItem) GetInstitutionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value
// and a boolean to check if the value has been set.
func (o *CraNetworkInsightsItem) GetInstitutionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionName, true
}

// SetInstitutionName sets field value
func (o *CraNetworkInsightsItem) SetInstitutionName(v string) {
	o.InstitutionName = v
}

// GetItemId returns the ItemId field value
func (o *CraNetworkInsightsItem) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *CraNetworkInsightsItem) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *CraNetworkInsightsItem) SetItemId(v string) {
	o.ItemId = v
}

func (o CraNetworkInsightsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if true {
		toSerialize["institution_name"] = o.InstitutionName
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	return json.Marshal(toSerialize)
}

type NullableCraNetworkInsightsItem struct {
	value *CraNetworkInsightsItem
	isSet bool
}

func (v NullableCraNetworkInsightsItem) Get() *CraNetworkInsightsItem {
	return v.value
}

func (v *NullableCraNetworkInsightsItem) Set(val *CraNetworkInsightsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCraNetworkInsightsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCraNetworkInsightsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraNetworkInsightsItem(val *CraNetworkInsightsItem) *NullableCraNetworkInsightsItem {
	return &NullableCraNetworkInsightsItem{value: val, isSet: true}
}

func (v NullableCraNetworkInsightsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraNetworkInsightsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


