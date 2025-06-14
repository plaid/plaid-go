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

// OwnerOverride Data about the owner or owners of an account. Any fields not specified will be filled in with default Sandbox information.
type OwnerOverride struct {
	// A list of names associated with the account by the financial institution. These should always be the names of individuals, even for business accounts. Note that the same name data will be used for all accounts associated with an Item.
	Names []string `json:"names"`
	// A list of phone numbers associated with the account.
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
	// A list of email addresses associated with the account.
	Emails []Email `json:"emails"`
	// Data about the various addresses associated with the account.
	Addresses []Address `json:"addresses"`
	AdditionalProperties map[string]interface{}
}

type _OwnerOverride OwnerOverride

// NewOwnerOverride instantiates a new OwnerOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerOverride(names []string, phoneNumbers []PhoneNumber, emails []Email, addresses []Address) *OwnerOverride {
	this := OwnerOverride{}
	this.Names = names
	this.PhoneNumbers = phoneNumbers
	this.Emails = emails
	this.Addresses = addresses
	return &this
}

// NewOwnerOverrideWithDefaults instantiates a new OwnerOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerOverrideWithDefaults() *OwnerOverride {
	this := OwnerOverride{}
	return &this
}

// GetNames returns the Names field value
func (o *OwnerOverride) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *OwnerOverride) GetNamesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Names, true
}

// SetNames sets field value
func (o *OwnerOverride) SetNames(v []string) {
	o.Names = v
}

// GetPhoneNumbers returns the PhoneNumbers field value
func (o *OwnerOverride) GetPhoneNumbers() []PhoneNumber {
	if o == nil {
		var ret []PhoneNumber
		return ret
	}

	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value
// and a boolean to check if the value has been set.
func (o *OwnerOverride) GetPhoneNumbersOk() (*[]PhoneNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PhoneNumbers, true
}

// SetPhoneNumbers sets field value
func (o *OwnerOverride) SetPhoneNumbers(v []PhoneNumber) {
	o.PhoneNumbers = v
}

// GetEmails returns the Emails field value
func (o *OwnerOverride) GetEmails() []Email {
	if o == nil {
		var ret []Email
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *OwnerOverride) GetEmailsOk() (*[]Email, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Emails, true
}

// SetEmails sets field value
func (o *OwnerOverride) SetEmails(v []Email) {
	o.Emails = v
}

// GetAddresses returns the Addresses field value
func (o *OwnerOverride) GetAddresses() []Address {
	if o == nil {
		var ret []Address
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *OwnerOverride) GetAddressesOk() (*[]Address, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *OwnerOverride) SetAddresses(v []Address) {
	o.Addresses = v
}

func (o OwnerOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["names"] = o.Names
	}
	if true {
		toSerialize["phone_numbers"] = o.PhoneNumbers
	}
	if true {
		toSerialize["emails"] = o.Emails
	}
	if true {
		toSerialize["addresses"] = o.Addresses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OwnerOverride) UnmarshalJSON(bytes []byte) (err error) {
	varOwnerOverride := _OwnerOverride{}

	if err = json.Unmarshal(bytes, &varOwnerOverride); err == nil {
		*o = OwnerOverride(varOwnerOverride)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "names")
		delete(additionalProperties, "phone_numbers")
		delete(additionalProperties, "emails")
		delete(additionalProperties, "addresses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOwnerOverride struct {
	value *OwnerOverride
	isSet bool
}

func (v NullableOwnerOverride) Get() *OwnerOverride {
	return v.value
}

func (v *NullableOwnerOverride) Set(val *OwnerOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerOverride(val *OwnerOverride) *NullableOwnerOverride {
	return &NullableOwnerOverride{value: val, isSet: true}
}

func (v NullableOwnerOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


