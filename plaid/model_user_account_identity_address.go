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

// UserAccountIdentityAddress The user's address.
type UserAccountIdentityAddress struct {
	// The full city name
	City NullableString `json:"city,omitempty"`
	// The region or state. Example: `\"NC\"`
	Region NullableString `json:"region,omitempty"`
	// The full street address Example: `\"564 Main Street, APT 15\"`
	Street NullableString `json:"street,omitempty"`
	// The second line street address
	Street2 NullableString `json:"street2,omitempty"`
	// The postal code. In API versions 2018-05-22 and earlier, this field is called `zip`.
	PostalCode NullableString `json:"postal_code,omitempty"`
	// The ISO 3166-1 alpha-2 country code
	Country NullableString `json:"country,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserAccountIdentityAddress UserAccountIdentityAddress

// NewUserAccountIdentityAddress instantiates a new UserAccountIdentityAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccountIdentityAddress() *UserAccountIdentityAddress {
	this := UserAccountIdentityAddress{}
	return &this
}

// NewUserAccountIdentityAddressWithDefaults instantiates a new UserAccountIdentityAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccountIdentityAddressWithDefaults() *UserAccountIdentityAddress {
	this := UserAccountIdentityAddress{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAccountIdentityAddress) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAccountIdentityAddress) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *UserAccountIdentityAddress) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *UserAccountIdentityAddress) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *UserAccountIdentityAddress) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *UserAccountIdentityAddress) UnsetCity() {
	o.City.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAccountIdentityAddress) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAccountIdentityAddress) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *UserAccountIdentityAddress) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *UserAccountIdentityAddress) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *UserAccountIdentityAddress) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *UserAccountIdentityAddress) UnsetRegion() {
	o.Region.Unset()
}

// GetStreet returns the Street field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAccountIdentityAddress) GetStreet() string {
	if o == nil || o.Street.Get() == nil {
		var ret string
		return ret
	}
	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAccountIdentityAddress) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// HasStreet returns a boolean if a field has been set.
func (o *UserAccountIdentityAddress) HasStreet() bool {
	if o != nil && o.Street.IsSet() {
		return true
	}

	return false
}

// SetStreet gets a reference to the given NullableString and assigns it to the Street field.
func (o *UserAccountIdentityAddress) SetStreet(v string) {
	o.Street.Set(&v)
}
// SetStreetNil sets the value for Street to be an explicit nil
func (o *UserAccountIdentityAddress) SetStreetNil() {
	o.Street.Set(nil)
}

// UnsetStreet ensures that no value is present for Street, not even an explicit nil
func (o *UserAccountIdentityAddress) UnsetStreet() {
	o.Street.Unset()
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAccountIdentityAddress) GetStreet2() string {
	if o == nil || o.Street2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Street2.Get()
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAccountIdentityAddress) GetStreet2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Street2.Get(), o.Street2.IsSet()
}

// HasStreet2 returns a boolean if a field has been set.
func (o *UserAccountIdentityAddress) HasStreet2() bool {
	if o != nil && o.Street2.IsSet() {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given NullableString and assigns it to the Street2 field.
func (o *UserAccountIdentityAddress) SetStreet2(v string) {
	o.Street2.Set(&v)
}
// SetStreet2Nil sets the value for Street2 to be an explicit nil
func (o *UserAccountIdentityAddress) SetStreet2Nil() {
	o.Street2.Set(nil)
}

// UnsetStreet2 ensures that no value is present for Street2, not even an explicit nil
func (o *UserAccountIdentityAddress) UnsetStreet2() {
	o.Street2.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAccountIdentityAddress) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAccountIdentityAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *UserAccountIdentityAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *UserAccountIdentityAddress) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *UserAccountIdentityAddress) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *UserAccountIdentityAddress) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAccountIdentityAddress) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAccountIdentityAddress) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *UserAccountIdentityAddress) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *UserAccountIdentityAddress) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *UserAccountIdentityAddress) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *UserAccountIdentityAddress) UnsetCountry() {
	o.Country.Unset()
}

func (o UserAccountIdentityAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Street.IsSet() {
		toSerialize["street"] = o.Street.Get()
	}
	if o.Street2.IsSet() {
		toSerialize["street2"] = o.Street2.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserAccountIdentityAddress) UnmarshalJSON(bytes []byte) (err error) {
	varUserAccountIdentityAddress := _UserAccountIdentityAddress{}

	if err = json.Unmarshal(bytes, &varUserAccountIdentityAddress); err == nil {
		*o = UserAccountIdentityAddress(varUserAccountIdentityAddress)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "street")
		delete(additionalProperties, "street2")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserAccountIdentityAddress struct {
	value *UserAccountIdentityAddress
	isSet bool
}

func (v NullableUserAccountIdentityAddress) Get() *UserAccountIdentityAddress {
	return v.value
}

func (v *NullableUserAccountIdentityAddress) Set(val *UserAccountIdentityAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccountIdentityAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccountIdentityAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccountIdentityAddress(val *UserAccountIdentityAddress) *NullableUserAccountIdentityAddress {
	return &NullableUserAccountIdentityAddress{value: val, isSet: true}
}

func (v NullableUserAccountIdentityAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccountIdentityAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


