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

// ScreeningHitData Information associated with the watchlist hit
type ScreeningHitData struct {
	// Dates of birth associated with the watchlist hit
	DatesOfBirth *[]ScreeningHitDateOfBirthItem `json:"dates_of_birth,omitempty"`
	// Documents associated with the watchlist hit
	Documents *[]ScreeningHitDocumentsItems `json:"documents,omitempty"`
	// Locations associated with the watchlist hit
	Locations *[]GenericScreeningHitLocationItems `json:"locations,omitempty"`
	// Names associated with the watchlist hit
	Names *[]ScreeningHitNamesItems `json:"names,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScreeningHitData ScreeningHitData

// NewScreeningHitData instantiates a new ScreeningHitData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScreeningHitData() *ScreeningHitData {
	this := ScreeningHitData{}
	return &this
}

// NewScreeningHitDataWithDefaults instantiates a new ScreeningHitData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScreeningHitDataWithDefaults() *ScreeningHitData {
	this := ScreeningHitData{}
	return &this
}

// GetDatesOfBirth returns the DatesOfBirth field value if set, zero value otherwise.
func (o *ScreeningHitData) GetDatesOfBirth() []ScreeningHitDateOfBirthItem {
	if o == nil || o.DatesOfBirth == nil {
		var ret []ScreeningHitDateOfBirthItem
		return ret
	}
	return *o.DatesOfBirth
}

// GetDatesOfBirthOk returns a tuple with the DatesOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitData) GetDatesOfBirthOk() (*[]ScreeningHitDateOfBirthItem, bool) {
	if o == nil || o.DatesOfBirth == nil {
		return nil, false
	}
	return o.DatesOfBirth, true
}

// HasDatesOfBirth returns a boolean if a field has been set.
func (o *ScreeningHitData) HasDatesOfBirth() bool {
	if o != nil && o.DatesOfBirth != nil {
		return true
	}

	return false
}

// SetDatesOfBirth gets a reference to the given []ScreeningHitDateOfBirthItem and assigns it to the DatesOfBirth field.
func (o *ScreeningHitData) SetDatesOfBirth(v []ScreeningHitDateOfBirthItem) {
	o.DatesOfBirth = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *ScreeningHitData) GetDocuments() []ScreeningHitDocumentsItems {
	if o == nil || o.Documents == nil {
		var ret []ScreeningHitDocumentsItems
		return ret
	}
	return *o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitData) GetDocumentsOk() (*[]ScreeningHitDocumentsItems, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *ScreeningHitData) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []ScreeningHitDocumentsItems and assigns it to the Documents field.
func (o *ScreeningHitData) SetDocuments(v []ScreeningHitDocumentsItems) {
	o.Documents = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *ScreeningHitData) GetLocations() []GenericScreeningHitLocationItems {
	if o == nil || o.Locations == nil {
		var ret []GenericScreeningHitLocationItems
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitData) GetLocationsOk() (*[]GenericScreeningHitLocationItems, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ScreeningHitData) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []GenericScreeningHitLocationItems and assigns it to the Locations field.
func (o *ScreeningHitData) SetLocations(v []GenericScreeningHitLocationItems) {
	o.Locations = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *ScreeningHitData) GetNames() []ScreeningHitNamesItems {
	if o == nil || o.Names == nil {
		var ret []ScreeningHitNamesItems
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScreeningHitData) GetNamesOk() (*[]ScreeningHitNamesItems, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *ScreeningHitData) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []ScreeningHitNamesItems and assigns it to the Names field.
func (o *ScreeningHitData) SetNames(v []ScreeningHitNamesItems) {
	o.Names = &v
}

func (o ScreeningHitData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DatesOfBirth != nil {
		toSerialize["dates_of_birth"] = o.DatesOfBirth
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ScreeningHitData) UnmarshalJSON(bytes []byte) (err error) {
	varScreeningHitData := _ScreeningHitData{}

	if err = json.Unmarshal(bytes, &varScreeningHitData); err == nil {
		*o = ScreeningHitData(varScreeningHitData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dates_of_birth")
		delete(additionalProperties, "documents")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "names")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScreeningHitData struct {
	value *ScreeningHitData
	isSet bool
}

func (v NullableScreeningHitData) Get() *ScreeningHitData {
	return v.value
}

func (v *NullableScreeningHitData) Set(val *ScreeningHitData) {
	v.value = val
	v.isSet = true
}

func (v NullableScreeningHitData) IsSet() bool {
	return v.isSet
}

func (v *NullableScreeningHitData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScreeningHitData(val *ScreeningHitData) *NullableScreeningHitData {
	return &NullableScreeningHitData{value: val, isSet: true}
}

func (v NullableScreeningHitData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScreeningHitData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


