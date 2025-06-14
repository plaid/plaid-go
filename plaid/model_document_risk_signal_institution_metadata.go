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

// DocumentRiskSignalInstitutionMetadata An object which contains additional metadata about the institution used to compute the verification attribute
type DocumentRiskSignalInstitutionMetadata struct {
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	AdditionalProperties map[string]interface{}
}

type _DocumentRiskSignalInstitutionMetadata DocumentRiskSignalInstitutionMetadata

// NewDocumentRiskSignalInstitutionMetadata instantiates a new DocumentRiskSignalInstitutionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentRiskSignalInstitutionMetadata(itemId string) *DocumentRiskSignalInstitutionMetadata {
	this := DocumentRiskSignalInstitutionMetadata{}
	this.ItemId = itemId
	return &this
}

// NewDocumentRiskSignalInstitutionMetadataWithDefaults instantiates a new DocumentRiskSignalInstitutionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentRiskSignalInstitutionMetadataWithDefaults() *DocumentRiskSignalInstitutionMetadata {
	this := DocumentRiskSignalInstitutionMetadata{}
	return &this
}

// GetItemId returns the ItemId field value
func (o *DocumentRiskSignalInstitutionMetadata) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *DocumentRiskSignalInstitutionMetadata) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *DocumentRiskSignalInstitutionMetadata) SetItemId(v string) {
	o.ItemId = v
}

func (o DocumentRiskSignalInstitutionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item_id"] = o.ItemId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DocumentRiskSignalInstitutionMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varDocumentRiskSignalInstitutionMetadata := _DocumentRiskSignalInstitutionMetadata{}

	if err = json.Unmarshal(bytes, &varDocumentRiskSignalInstitutionMetadata); err == nil {
		*o = DocumentRiskSignalInstitutionMetadata(varDocumentRiskSignalInstitutionMetadata)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentRiskSignalInstitutionMetadata struct {
	value *DocumentRiskSignalInstitutionMetadata
	isSet bool
}

func (v NullableDocumentRiskSignalInstitutionMetadata) Get() *DocumentRiskSignalInstitutionMetadata {
	return v.value
}

func (v *NullableDocumentRiskSignalInstitutionMetadata) Set(val *DocumentRiskSignalInstitutionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentRiskSignalInstitutionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentRiskSignalInstitutionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentRiskSignalInstitutionMetadata(val *DocumentRiskSignalInstitutionMetadata) *NullableDocumentRiskSignalInstitutionMetadata {
	return &NullableDocumentRiskSignalInstitutionMetadata{value: val, isSet: true}
}

func (v NullableDocumentRiskSignalInstitutionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentRiskSignalInstitutionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


