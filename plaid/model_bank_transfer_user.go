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

// BankTransferUser The legal name and other information for the account holder.
type BankTransferUser struct {
	// The account holder’s full legal name. If the transfer `ach_class` is `ccd`, this should be the business name of the account holder.
	LegalName string `json:"legal_name"`
	// The account holder’s email.
	EmailAddress NullableString `json:"email_address,omitempty"`
	// The account holder's routing number. This field is only used in response data. Do not provide this field when making requests.
	RoutingNumber *string `json:"routing_number,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferUser BankTransferUser

// NewBankTransferUser instantiates a new BankTransferUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferUser(legalName string) *BankTransferUser {
	this := BankTransferUser{}
	this.LegalName = legalName
	return &this
}

// NewBankTransferUserWithDefaults instantiates a new BankTransferUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferUserWithDefaults() *BankTransferUser {
	this := BankTransferUser{}
	return &this
}

// GetLegalName returns the LegalName field value
func (o *BankTransferUser) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *BankTransferUser) GetLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *BankTransferUser) SetLegalName(v string) {
	o.LegalName = v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankTransferUser) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferUser) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *BankTransferUser) HasEmailAddress() bool {
	if o != nil && o.EmailAddress.IsSet() {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given NullableString and assigns it to the EmailAddress field.
func (o *BankTransferUser) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}
// SetEmailAddressNil sets the value for EmailAddress to be an explicit nil
func (o *BankTransferUser) SetEmailAddressNil() {
	o.EmailAddress.Set(nil)
}

// UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
func (o *BankTransferUser) UnsetEmailAddress() {
	o.EmailAddress.Unset()
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *BankTransferUser) GetRoutingNumber() string {
	if o == nil || o.RoutingNumber == nil {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferUser) GetRoutingNumberOk() (*string, bool) {
	if o == nil || o.RoutingNumber == nil {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *BankTransferUser) HasRoutingNumber() bool {
	if o != nil && o.RoutingNumber != nil {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *BankTransferUser) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

func (o BankTransferUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["legal_name"] = o.LegalName
	}
	if o.EmailAddress.IsSet() {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if o.RoutingNumber != nil {
		toSerialize["routing_number"] = o.RoutingNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferUser) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferUser := _BankTransferUser{}

	if err = json.Unmarshal(bytes, &varBankTransferUser); err == nil {
		*o = BankTransferUser(varBankTransferUser)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "legal_name")
		delete(additionalProperties, "email_address")
		delete(additionalProperties, "routing_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferUser struct {
	value *BankTransferUser
	isSet bool
}

func (v NullableBankTransferUser) Get() *BankTransferUser {
	return v.value
}

func (v *NullableBankTransferUser) Set(val *BankTransferUser) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferUser) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferUser(val *BankTransferUser) *NullableBankTransferUser {
	return &NullableBankTransferUser{value: val, isSet: true}
}

func (v NullableBankTransferUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


