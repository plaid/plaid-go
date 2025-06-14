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

// LinkTokenEUConfig Configuration parameters for EU flows
type LinkTokenEUConfig struct {
	// If `true`, open Link without an initial UI. Defaults to `false`.
	Headless *bool `json:"headless,omitempty"`
}

// NewLinkTokenEUConfig instantiates a new LinkTokenEUConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenEUConfig() *LinkTokenEUConfig {
	this := LinkTokenEUConfig{}
	return &this
}

// NewLinkTokenEUConfigWithDefaults instantiates a new LinkTokenEUConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenEUConfigWithDefaults() *LinkTokenEUConfig {
	this := LinkTokenEUConfig{}
	return &this
}

// GetHeadless returns the Headless field value if set, zero value otherwise.
func (o *LinkTokenEUConfig) GetHeadless() bool {
	if o == nil || o.Headless == nil {
		var ret bool
		return ret
	}
	return *o.Headless
}

// GetHeadlessOk returns a tuple with the Headless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenEUConfig) GetHeadlessOk() (*bool, bool) {
	if o == nil || o.Headless == nil {
		return nil, false
	}
	return o.Headless, true
}

// HasHeadless returns a boolean if a field has been set.
func (o *LinkTokenEUConfig) HasHeadless() bool {
	if o != nil && o.Headless != nil {
		return true
	}

	return false
}

// SetHeadless gets a reference to the given bool and assigns it to the Headless field.
func (o *LinkTokenEUConfig) SetHeadless(v bool) {
	o.Headless = &v
}

func (o LinkTokenEUConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Headless != nil {
		toSerialize["headless"] = o.Headless
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenEUConfig struct {
	value *LinkTokenEUConfig
	isSet bool
}

func (v NullableLinkTokenEUConfig) Get() *LinkTokenEUConfig {
	return v.value
}

func (v *NullableLinkTokenEUConfig) Set(val *LinkTokenEUConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenEUConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenEUConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenEUConfig(val *LinkTokenEUConfig) *NullableLinkTokenEUConfig {
	return &NullableLinkTokenEUConfig{value: val, isSet: true}
}

func (v NullableLinkTokenEUConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenEUConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


