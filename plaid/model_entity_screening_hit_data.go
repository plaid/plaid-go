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

// EntityScreeningHitData Information associated with the entity watchlist hit
type EntityScreeningHitData struct {
	// Documents associated with the watchlist hit
	Documents *[]EntityScreeningHitDocumentsItems `json:"documents,omitempty"`
	// Email addresses associated with the watchlist hit
	EmailAddresses *[]EntityScreeningHitEmailsItems `json:"email_addresses,omitempty"`
	// Locations associated with the watchlist hit
	Locations *[]GenericScreeningHitLocationItems `json:"locations,omitempty"`
	// Names associated with the watchlist hit
	Names *[]EntityScreeningHitNamesItems `json:"names,omitempty"`
	// Phone numbers associated with the watchlist hit
	PhoneNumbers *[]EntityScreeningHitsPhoneNumberItems `json:"phone_numbers,omitempty"`
	// URLs associated with the watchlist hit
	Urls *[]EntityScreeningHitUrlsItems `json:"urls,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityScreeningHitData EntityScreeningHitData

// NewEntityScreeningHitData instantiates a new EntityScreeningHitData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningHitData() *EntityScreeningHitData {
	this := EntityScreeningHitData{}
	return &this
}

// NewEntityScreeningHitDataWithDefaults instantiates a new EntityScreeningHitData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningHitDataWithDefaults() *EntityScreeningHitData {
	this := EntityScreeningHitData{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *EntityScreeningHitData) GetDocuments() []EntityScreeningHitDocumentsItems {
	if o == nil || o.Documents == nil {
		var ret []EntityScreeningHitDocumentsItems
		return ret
	}
	return *o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitData) GetDocumentsOk() (*[]EntityScreeningHitDocumentsItems, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *EntityScreeningHitData) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []EntityScreeningHitDocumentsItems and assigns it to the Documents field.
func (o *EntityScreeningHitData) SetDocuments(v []EntityScreeningHitDocumentsItems) {
	o.Documents = &v
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise.
func (o *EntityScreeningHitData) GetEmailAddresses() []EntityScreeningHitEmailsItems {
	if o == nil || o.EmailAddresses == nil {
		var ret []EntityScreeningHitEmailsItems
		return ret
	}
	return *o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitData) GetEmailAddressesOk() (*[]EntityScreeningHitEmailsItems, bool) {
	if o == nil || o.EmailAddresses == nil {
		return nil, false
	}
	return o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *EntityScreeningHitData) HasEmailAddresses() bool {
	if o != nil && o.EmailAddresses != nil {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given []EntityScreeningHitEmailsItems and assigns it to the EmailAddresses field.
func (o *EntityScreeningHitData) SetEmailAddresses(v []EntityScreeningHitEmailsItems) {
	o.EmailAddresses = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *EntityScreeningHitData) GetLocations() []GenericScreeningHitLocationItems {
	if o == nil || o.Locations == nil {
		var ret []GenericScreeningHitLocationItems
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitData) GetLocationsOk() (*[]GenericScreeningHitLocationItems, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *EntityScreeningHitData) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []GenericScreeningHitLocationItems and assigns it to the Locations field.
func (o *EntityScreeningHitData) SetLocations(v []GenericScreeningHitLocationItems) {
	o.Locations = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *EntityScreeningHitData) GetNames() []EntityScreeningHitNamesItems {
	if o == nil || o.Names == nil {
		var ret []EntityScreeningHitNamesItems
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitData) GetNamesOk() (*[]EntityScreeningHitNamesItems, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *EntityScreeningHitData) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []EntityScreeningHitNamesItems and assigns it to the Names field.
func (o *EntityScreeningHitData) SetNames(v []EntityScreeningHitNamesItems) {
	o.Names = &v
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *EntityScreeningHitData) GetPhoneNumbers() []EntityScreeningHitsPhoneNumberItems {
	if o == nil || o.PhoneNumbers == nil {
		var ret []EntityScreeningHitsPhoneNumberItems
		return ret
	}
	return *o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitData) GetPhoneNumbersOk() (*[]EntityScreeningHitsPhoneNumberItems, bool) {
	if o == nil || o.PhoneNumbers == nil {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *EntityScreeningHitData) HasPhoneNumbers() bool {
	if o != nil && o.PhoneNumbers != nil {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given []EntityScreeningHitsPhoneNumberItems and assigns it to the PhoneNumbers field.
func (o *EntityScreeningHitData) SetPhoneNumbers(v []EntityScreeningHitsPhoneNumberItems) {
	o.PhoneNumbers = &v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *EntityScreeningHitData) GetUrls() []EntityScreeningHitUrlsItems {
	if o == nil || o.Urls == nil {
		var ret []EntityScreeningHitUrlsItems
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitData) GetUrlsOk() (*[]EntityScreeningHitUrlsItems, bool) {
	if o == nil || o.Urls == nil {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *EntityScreeningHitData) HasUrls() bool {
	if o != nil && o.Urls != nil {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []EntityScreeningHitUrlsItems and assigns it to the Urls field.
func (o *EntityScreeningHitData) SetUrls(v []EntityScreeningHitUrlsItems) {
	o.Urls = &v
}

func (o EntityScreeningHitData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.EmailAddresses != nil {
		toSerialize["email_addresses"] = o.EmailAddresses
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.PhoneNumbers != nil {
		toSerialize["phone_numbers"] = o.PhoneNumbers
	}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityScreeningHitData) UnmarshalJSON(bytes []byte) (err error) {
	varEntityScreeningHitData := _EntityScreeningHitData{}

	if err = json.Unmarshal(bytes, &varEntityScreeningHitData); err == nil {
		*o = EntityScreeningHitData(varEntityScreeningHitData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "documents")
		delete(additionalProperties, "email_addresses")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "names")
		delete(additionalProperties, "phone_numbers")
		delete(additionalProperties, "urls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityScreeningHitData struct {
	value *EntityScreeningHitData
	isSet bool
}

func (v NullableEntityScreeningHitData) Get() *EntityScreeningHitData {
	return v.value
}

func (v *NullableEntityScreeningHitData) Set(val *EntityScreeningHitData) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningHitData) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningHitData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningHitData(val *EntityScreeningHitData) *NullableEntityScreeningHitData {
	return &NullableEntityScreeningHitData{value: val, isSet: true}
}

func (v NullableEntityScreeningHitData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningHitData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


