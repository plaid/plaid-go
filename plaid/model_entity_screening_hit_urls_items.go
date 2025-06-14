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

// EntityScreeningHitUrlsItems Analyzed URLs for the associated hit
type EntityScreeningHitUrlsItems struct {
	Analysis *MatchSummary `json:"analysis,omitempty"`
	Data *EntityScreeningHitUrls `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityScreeningHitUrlsItems EntityScreeningHitUrlsItems

// NewEntityScreeningHitUrlsItems instantiates a new EntityScreeningHitUrlsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningHitUrlsItems() *EntityScreeningHitUrlsItems {
	this := EntityScreeningHitUrlsItems{}
	return &this
}

// NewEntityScreeningHitUrlsItemsWithDefaults instantiates a new EntityScreeningHitUrlsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningHitUrlsItemsWithDefaults() *EntityScreeningHitUrlsItems {
	this := EntityScreeningHitUrlsItems{}
	return &this
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *EntityScreeningHitUrlsItems) GetAnalysis() MatchSummary {
	if o == nil || o.Analysis == nil {
		var ret MatchSummary
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitUrlsItems) GetAnalysisOk() (*MatchSummary, bool) {
	if o == nil || o.Analysis == nil {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *EntityScreeningHitUrlsItems) HasAnalysis() bool {
	if o != nil && o.Analysis != nil {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given MatchSummary and assigns it to the Analysis field.
func (o *EntityScreeningHitUrlsItems) SetAnalysis(v MatchSummary) {
	o.Analysis = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EntityScreeningHitUrlsItems) GetData() EntityScreeningHitUrls {
	if o == nil || o.Data == nil {
		var ret EntityScreeningHitUrls
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitUrlsItems) GetDataOk() (*EntityScreeningHitUrls, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EntityScreeningHitUrlsItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given EntityScreeningHitUrls and assigns it to the Data field.
func (o *EntityScreeningHitUrlsItems) SetData(v EntityScreeningHitUrls) {
	o.Data = &v
}

func (o EntityScreeningHitUrlsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Analysis != nil {
		toSerialize["analysis"] = o.Analysis
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityScreeningHitUrlsItems) UnmarshalJSON(bytes []byte) (err error) {
	varEntityScreeningHitUrlsItems := _EntityScreeningHitUrlsItems{}

	if err = json.Unmarshal(bytes, &varEntityScreeningHitUrlsItems); err == nil {
		*o = EntityScreeningHitUrlsItems(varEntityScreeningHitUrlsItems)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "analysis")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityScreeningHitUrlsItems struct {
	value *EntityScreeningHitUrlsItems
	isSet bool
}

func (v NullableEntityScreeningHitUrlsItems) Get() *EntityScreeningHitUrlsItems {
	return v.value
}

func (v *NullableEntityScreeningHitUrlsItems) Set(val *EntityScreeningHitUrlsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningHitUrlsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningHitUrlsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningHitUrlsItems(val *EntityScreeningHitUrlsItems) *NullableEntityScreeningHitUrlsItems {
	return &NullableEntityScreeningHitUrlsItems{value: val, isSet: true}
}

func (v NullableEntityScreeningHitUrlsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningHitUrlsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


