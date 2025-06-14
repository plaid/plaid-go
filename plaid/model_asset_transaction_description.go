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

// AssetTransactionDescription Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type AssetTransactionDescription struct {
	// Asset Transaction Description String up to 3 occurances 1 required.
	AssetTransactionDescription string `json:"AssetTransactionDescription"`
	AdditionalProperties map[string]interface{}
}

type _AssetTransactionDescription AssetTransactionDescription

// NewAssetTransactionDescription instantiates a new AssetTransactionDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTransactionDescription(assetTransactionDescription string) *AssetTransactionDescription {
	this := AssetTransactionDescription{}
	this.AssetTransactionDescription = assetTransactionDescription
	return &this
}

// NewAssetTransactionDescriptionWithDefaults instantiates a new AssetTransactionDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTransactionDescriptionWithDefaults() *AssetTransactionDescription {
	this := AssetTransactionDescription{}
	return &this
}

// GetAssetTransactionDescription returns the AssetTransactionDescription field value
func (o *AssetTransactionDescription) GetAssetTransactionDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetTransactionDescription
}

// GetAssetTransactionDescriptionOk returns a tuple with the AssetTransactionDescription field value
// and a boolean to check if the value has been set.
func (o *AssetTransactionDescription) GetAssetTransactionDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetTransactionDescription, true
}

// SetAssetTransactionDescription sets field value
func (o *AssetTransactionDescription) SetAssetTransactionDescription(v string) {
	o.AssetTransactionDescription = v
}

func (o AssetTransactionDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AssetTransactionDescription"] = o.AssetTransactionDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetTransactionDescription) UnmarshalJSON(bytes []byte) (err error) {
	varAssetTransactionDescription := _AssetTransactionDescription{}

	if err = json.Unmarshal(bytes, &varAssetTransactionDescription); err == nil {
		*o = AssetTransactionDescription(varAssetTransactionDescription)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AssetTransactionDescription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetTransactionDescription struct {
	value *AssetTransactionDescription
	isSet bool
}

func (v NullableAssetTransactionDescription) Get() *AssetTransactionDescription {
	return v.value
}

func (v *NullableAssetTransactionDescription) Set(val *AssetTransactionDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTransactionDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTransactionDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTransactionDescription(val *AssetTransactionDescription) *NullableAssetTransactionDescription {
	return &NullableAssetTransactionDescription{value: val, isSet: true}
}

func (v NullableAssetTransactionDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTransactionDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


