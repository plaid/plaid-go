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

// EntityDocument An official document, usually issued by a governing body or institution, with an associated identifier.
type EntityDocument struct {
	Type EntityDocumentType `json:"type"`
	// The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.
	Number string `json:"number"`
	AdditionalProperties map[string]interface{}
}

type _EntityDocument EntityDocument

// NewEntityDocument instantiates a new EntityDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityDocument(type_ EntityDocumentType, number string) *EntityDocument {
	this := EntityDocument{}
	this.Type = type_
	this.Number = number
	return &this
}

// NewEntityDocumentWithDefaults instantiates a new EntityDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityDocumentWithDefaults() *EntityDocument {
	this := EntityDocument{}
	return &this
}

// GetType returns the Type field value
func (o *EntityDocument) GetType() EntityDocumentType {
	if o == nil {
		var ret EntityDocumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntityDocument) GetTypeOk() (*EntityDocumentType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EntityDocument) SetType(v EntityDocumentType) {
	o.Type = v
}

// GetNumber returns the Number field value
func (o *EntityDocument) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *EntityDocument) GetNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *EntityDocument) SetNumber(v string) {
	o.Number = v
}

func (o EntityDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["number"] = o.Number
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityDocument) UnmarshalJSON(bytes []byte) (err error) {
	varEntityDocument := _EntityDocument{}

	if err = json.Unmarshal(bytes, &varEntityDocument); err == nil {
		*o = EntityDocument(varEntityDocument)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityDocument struct {
	value *EntityDocument
	isSet bool
}

func (v NullableEntityDocument) Get() *EntityDocument {
	return v.value
}

func (v *NullableEntityDocument) Set(val *EntityDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityDocument(val *EntityDocument) *NullableEntityDocument {
	return &NullableEntityDocument{value: val, isSet: true}
}

func (v NullableEntityDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


