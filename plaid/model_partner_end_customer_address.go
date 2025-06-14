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

// PartnerEndCustomerAddress The end customer's address.
type PartnerEndCustomerAddress struct {
	City *string `json:"city,omitempty"`
	Street *string `json:"street,omitempty"`
	Region *string `json:"region,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	// ISO-3166-1 alpha-2 country code standard.
	CountryCode *string `json:"country_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerEndCustomerAddress PartnerEndCustomerAddress

// NewPartnerEndCustomerAddress instantiates a new PartnerEndCustomerAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomerAddress() *PartnerEndCustomerAddress {
	this := PartnerEndCustomerAddress{}
	return &this
}

// NewPartnerEndCustomerAddressWithDefaults instantiates a new PartnerEndCustomerAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerAddressWithDefaults() *PartnerEndCustomerAddress {
	this := PartnerEndCustomerAddress{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *PartnerEndCustomerAddress) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerAddress) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *PartnerEndCustomerAddress) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *PartnerEndCustomerAddress) SetCity(v string) {
	o.City = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *PartnerEndCustomerAddress) GetStreet() string {
	if o == nil || o.Street == nil {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerAddress) GetStreetOk() (*string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *PartnerEndCustomerAddress) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *PartnerEndCustomerAddress) SetStreet(v string) {
	o.Street = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PartnerEndCustomerAddress) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerAddress) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PartnerEndCustomerAddress) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PartnerEndCustomerAddress) SetRegion(v string) {
	o.Region = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *PartnerEndCustomerAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PartnerEndCustomerAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *PartnerEndCustomerAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PartnerEndCustomerAddress) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PartnerEndCustomerAddress) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PartnerEndCustomerAddress) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o PartnerEndCustomerAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Street != nil {
		toSerialize["street"] = o.Street
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerEndCustomerAddress) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerEndCustomerAddress := _PartnerEndCustomerAddress{}

	if err = json.Unmarshal(bytes, &varPartnerEndCustomerAddress); err == nil {
		*o = PartnerEndCustomerAddress(varPartnerEndCustomerAddress)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "street")
		delete(additionalProperties, "region")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerEndCustomerAddress struct {
	value *PartnerEndCustomerAddress
	isSet bool
}

func (v NullablePartnerEndCustomerAddress) Get() *PartnerEndCustomerAddress {
	return v.value
}

func (v *NullablePartnerEndCustomerAddress) Set(val *PartnerEndCustomerAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerAddress(val *PartnerEndCustomerAddress) *NullablePartnerEndCustomerAddress {
	return &NullablePartnerEndCustomerAddress{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


