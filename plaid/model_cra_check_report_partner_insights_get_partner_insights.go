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

// CraCheckReportPartnerInsightsGetPartnerInsights Defines configuration to generate Partner Insights
type CraCheckReportPartnerInsightsGetPartnerInsights struct {
	PrismVersions NullablePrismVersions `json:"prism_versions,omitempty"`
}

// NewCraCheckReportPartnerInsightsGetPartnerInsights instantiates a new CraCheckReportPartnerInsightsGetPartnerInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraCheckReportPartnerInsightsGetPartnerInsights() *CraCheckReportPartnerInsightsGetPartnerInsights {
	this := CraCheckReportPartnerInsightsGetPartnerInsights{}
	return &this
}

// NewCraCheckReportPartnerInsightsGetPartnerInsightsWithDefaults instantiates a new CraCheckReportPartnerInsightsGetPartnerInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraCheckReportPartnerInsightsGetPartnerInsightsWithDefaults() *CraCheckReportPartnerInsightsGetPartnerInsights {
	this := CraCheckReportPartnerInsightsGetPartnerInsights{}
	return &this
}

// GetPrismVersions returns the PrismVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CraCheckReportPartnerInsightsGetPartnerInsights) GetPrismVersions() PrismVersions {
	if o == nil || o.PrismVersions.Get() == nil {
		var ret PrismVersions
		return ret
	}
	return *o.PrismVersions.Get()
}

// GetPrismVersionsOk returns a tuple with the PrismVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraCheckReportPartnerInsightsGetPartnerInsights) GetPrismVersionsOk() (*PrismVersions, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrismVersions.Get(), o.PrismVersions.IsSet()
}

// HasPrismVersions returns a boolean if a field has been set.
func (o *CraCheckReportPartnerInsightsGetPartnerInsights) HasPrismVersions() bool {
	if o != nil && o.PrismVersions.IsSet() {
		return true
	}

	return false
}

// SetPrismVersions gets a reference to the given NullablePrismVersions and assigns it to the PrismVersions field.
func (o *CraCheckReportPartnerInsightsGetPartnerInsights) SetPrismVersions(v PrismVersions) {
	o.PrismVersions.Set(&v)
}
// SetPrismVersionsNil sets the value for PrismVersions to be an explicit nil
func (o *CraCheckReportPartnerInsightsGetPartnerInsights) SetPrismVersionsNil() {
	o.PrismVersions.Set(nil)
}

// UnsetPrismVersions ensures that no value is present for PrismVersions, not even an explicit nil
func (o *CraCheckReportPartnerInsightsGetPartnerInsights) UnsetPrismVersions() {
	o.PrismVersions.Unset()
}

func (o CraCheckReportPartnerInsightsGetPartnerInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PrismVersions.IsSet() {
		toSerialize["prism_versions"] = o.PrismVersions.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCraCheckReportPartnerInsightsGetPartnerInsights struct {
	value *CraCheckReportPartnerInsightsGetPartnerInsights
	isSet bool
}

func (v NullableCraCheckReportPartnerInsightsGetPartnerInsights) Get() *CraCheckReportPartnerInsightsGetPartnerInsights {
	return v.value
}

func (v *NullableCraCheckReportPartnerInsightsGetPartnerInsights) Set(val *CraCheckReportPartnerInsightsGetPartnerInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableCraCheckReportPartnerInsightsGetPartnerInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableCraCheckReportPartnerInsightsGetPartnerInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraCheckReportPartnerInsightsGetPartnerInsights(val *CraCheckReportPartnerInsightsGetPartnerInsights) *NullableCraCheckReportPartnerInsightsGetPartnerInsights {
	return &NullableCraCheckReportPartnerInsightsGetPartnerInsights{value: val, isSet: true}
}

func (v NullableCraCheckReportPartnerInsightsGetPartnerInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraCheckReportPartnerInsightsGetPartnerInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


