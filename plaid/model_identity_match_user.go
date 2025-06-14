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

// IdentityMatchUser The user's legal name, phone number, email address and address used to perform fuzzy match. If Financial Account Matching is enabled in the Identity Verification product, leave this field empty to automatically match against PII collected from the Identity Verification checks.
type IdentityMatchUser struct {
	// The user's full legal name.
	LegalName NullableString `json:"legal_name,omitempty"`
	// The user's phone number, in E.164 format: +{countrycode}{number}. For example: \"+14157452130\". Phone numbers provided in other formats will be parsed on a best-effort basis. Phone number input is validated against valid number ranges; number strings that do not match a real-world phone numbering scheme may cause the request to fail, even in the Sandbox test environment.
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// The user's email address.
	EmailAddress NullableString `json:"email_address,omitempty"`
	Address NullableAddressDataNullableNoRequiredFields `json:"address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityMatchUser IdentityMatchUser

// NewIdentityMatchUser instantiates a new IdentityMatchUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityMatchUser() *IdentityMatchUser {
	this := IdentityMatchUser{}
	return &this
}

// NewIdentityMatchUserWithDefaults instantiates a new IdentityMatchUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityMatchUserWithDefaults() *IdentityMatchUser {
	this := IdentityMatchUser{}
	return &this
}

// GetLegalName returns the LegalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMatchUser) GetLegalName() string {
	if o == nil || o.LegalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LegalName.Get()
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMatchUser) GetLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LegalName.Get(), o.LegalName.IsSet()
}

// HasLegalName returns a boolean if a field has been set.
func (o *IdentityMatchUser) HasLegalName() bool {
	if o != nil && o.LegalName.IsSet() {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given NullableString and assigns it to the LegalName field.
func (o *IdentityMatchUser) SetLegalName(v string) {
	o.LegalName.Set(&v)
}
// SetLegalNameNil sets the value for LegalName to be an explicit nil
func (o *IdentityMatchUser) SetLegalNameNil() {
	o.LegalName.Set(nil)
}

// UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
func (o *IdentityMatchUser) UnsetLegalName() {
	o.LegalName.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMatchUser) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMatchUser) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *IdentityMatchUser) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *IdentityMatchUser) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *IdentityMatchUser) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *IdentityMatchUser) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMatchUser) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMatchUser) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *IdentityMatchUser) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableString and assigns it to the EmailAddress field.
func (o *IdentityMatchUser) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *IdentityMatchUser) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *IdentityMatchUser) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityMatchUser) GetAddress() AddressDataNullableNoRequiredFields {
	if o == nil || o.Address.Get() == nil {
		var ret AddressDataNullableNoRequiredFields
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityMatchUser) GetAddressOk() (*AddressDataNullableNoRequiredFields, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *IdentityMatchUser) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableAddressDataNullableNoRequiredFields and assigns it to the Address field.
func (o *IdentityMatchUser) SetAddress(v AddressDataNullableNoRequiredFields) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *IdentityMatchUser) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *IdentityMatchUser) UnsetAddress() {
	o.Address.Unset()
}

func (o IdentityMatchUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LegalName.IsSet() {
		toSerialize["legal_name"] = o.LegalName.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.EmailAddress.IsSet() {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityMatchUser) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityMatchUser := _IdentityMatchUser{}

	if err = json.Unmarshal(bytes, &varIdentityMatchUser); err == nil {
		*o = IdentityMatchUser(varIdentityMatchUser)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "legal_name")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "email_address")
		delete(additionalProperties, "address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityMatchUser struct {
	value *IdentityMatchUser
	isSet bool
}

func (v NullableIdentityMatchUser) Get() *IdentityMatchUser {
	return v.value
}

func (v *NullableIdentityMatchUser) Set(val *IdentityMatchUser) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityMatchUser) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityMatchUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityMatchUser(val *IdentityMatchUser) *NullableIdentityMatchUser {
	return &NullableIdentityMatchUser{value: val, isSet: true}
}

func (v NullableIdentityMatchUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityMatchUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


