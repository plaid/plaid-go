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

// EntityWatchlistSearchTerms Search inputs for creating an entity watchlist screening
type EntityWatchlistSearchTerms struct {
	// ID of the associated entity program.
	EntityWatchlistProgramId string `json:"entity_watchlist_program_id"`
	// The name of the organization being screened. Must have at least one alphabetical character, have a maximum length of 100 characters, and not include leading or trailing spaces.
	LegalName string `json:"legal_name"`
	// The numeric or alphanumeric identifier associated with this document. Must be between 4 and 32 characters long, and cannot have leading or trailing spaces.
	DocumentNumber NullableString `json:"document_number,omitempty"`
	// A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).
	EmailAddress NullableString `json:"email_address,omitempty"`
	// Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.
	Country NullableString `json:"country,omitempty"`
	// A phone number in E.164 format.
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// An 'http' or 'https' URL (must begin with either of those).
	Url NullableString `json:"url,omitempty"`
}

// NewEntityWatchlistSearchTerms instantiates a new EntityWatchlistSearchTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityWatchlistSearchTerms(entityWatchlistProgramId string, legalName string) *EntityWatchlistSearchTerms {
	this := EntityWatchlistSearchTerms{}
	this.EntityWatchlistProgramId = entityWatchlistProgramId
	this.LegalName = legalName
	return &this
}

// NewEntityWatchlistSearchTermsWithDefaults instantiates a new EntityWatchlistSearchTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityWatchlistSearchTermsWithDefaults() *EntityWatchlistSearchTerms {
	this := EntityWatchlistSearchTerms{}
	return &this
}

// GetEntityWatchlistProgramId returns the EntityWatchlistProgramId field value
func (o *EntityWatchlistSearchTerms) GetEntityWatchlistProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityWatchlistProgramId
}

// GetEntityWatchlistProgramIdOk returns a tuple with the EntityWatchlistProgramId field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistSearchTerms) GetEntityWatchlistProgramIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistProgramId, true
}

// SetEntityWatchlistProgramId sets field value
func (o *EntityWatchlistSearchTerms) SetEntityWatchlistProgramId(v string) {
	o.EntityWatchlistProgramId = v
}

// GetLegalName returns the LegalName field value
func (o *EntityWatchlistSearchTerms) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistSearchTerms) GetLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *EntityWatchlistSearchTerms) SetLegalName(v string) {
	o.LegalName = v
}

// GetDocumentNumber returns the DocumentNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntityWatchlistSearchTerms) GetDocumentNumber() string {
	if o == nil || o.DocumentNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.DocumentNumber.Get()
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistSearchTerms) GetDocumentNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentNumber.Get(), o.DocumentNumber.IsSet()
}

// HasDocumentNumber returns a boolean if a field has been set.
func (o *EntityWatchlistSearchTerms) HasDocumentNumber() bool {
	if o != nil && o.DocumentNumber.IsSet() {
		return true
	}

	return false
}

// SetDocumentNumber gets a reference to the given NullableString and assigns it to the DocumentNumber field.
func (o *EntityWatchlistSearchTerms) SetDocumentNumber(v string) {
	o.DocumentNumber.Set(&v)
}
// SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil
func (o *EntityWatchlistSearchTerms) SetDocumentNumberNil() {
	o.DocumentNumber.Set(nil)
}

// UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
func (o *EntityWatchlistSearchTerms) UnsetDocumentNumber() {
	o.DocumentNumber.Unset()
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntityWatchlistSearchTerms) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistSearchTerms) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *EntityWatchlistSearchTerms) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableString and assigns it to the EmailAddress field.
func (o *EntityWatchlistSearchTerms) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *EntityWatchlistSearchTerms) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *EntityWatchlistSearchTerms) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntityWatchlistSearchTerms) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistSearchTerms) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *EntityWatchlistSearchTerms) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *EntityWatchlistSearchTerms) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *EntityWatchlistSearchTerms) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *EntityWatchlistSearchTerms) UnsetCountry() {
	o.Country.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntityWatchlistSearchTerms) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistSearchTerms) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *EntityWatchlistSearchTerms) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *EntityWatchlistSearchTerms) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *EntityWatchlistSearchTerms) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *EntityWatchlistSearchTerms) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntityWatchlistSearchTerms) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntityWatchlistSearchTerms) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *EntityWatchlistSearchTerms) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *EntityWatchlistSearchTerms) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *EntityWatchlistSearchTerms) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *EntityWatchlistSearchTerms) UnsetUrl() {
	o.Url.Unset()
}

func (o EntityWatchlistSearchTerms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_watchlist_program_id"] = o.EntityWatchlistProgramId
	}
	if true {
		toSerialize["legal_name"] = o.LegalName
	}
	if o.DocumentNumber.IsSet() {
		toSerialize["document_number"] = o.DocumentNumber.Get()
	}
	if o.EmailAddress.IsSet() {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEntityWatchlistSearchTerms struct {
	value *EntityWatchlistSearchTerms
	isSet bool
}

func (v NullableEntityWatchlistSearchTerms) Get() *EntityWatchlistSearchTerms {
	return v.value
}

func (v *NullableEntityWatchlistSearchTerms) Set(val *EntityWatchlistSearchTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityWatchlistSearchTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityWatchlistSearchTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityWatchlistSearchTerms(val *EntityWatchlistSearchTerms) *NullableEntityWatchlistSearchTerms {
	return &NullableEntityWatchlistSearchTerms{value: val, isSet: true}
}

func (v NullableEntityWatchlistSearchTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityWatchlistSearchTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


