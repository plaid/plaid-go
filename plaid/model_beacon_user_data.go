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

// BeaconUserData A Beacon User's data and resulting analysis when checked against duplicate records and the Beacon Fraud Network.
type BeaconUserData struct {
	// A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).
	DateOfBirth string `json:"date_of_birth"`
	Name BeaconUserName `json:"name"`
	Address BeaconUserAddress `json:"address"`
	// A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).
	EmailAddress NullableString `json:"email_address"`
	// A phone number in E.164 format.
	PhoneNumber NullableString `json:"phone_number"`
	IdNumber NullableBeaconUserIDNumber `json:"id_number"`
	// An IPv4 or IPV6 address.
	IpAddress NullableString `json:"ip_address"`
	DepositoryAccounts []BeaconUserDepositoryAccount `json:"depository_accounts"`
	AdditionalProperties map[string]interface{}
}

type _BeaconUserData BeaconUserData

// NewBeaconUserData instantiates a new BeaconUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeaconUserData(dateOfBirth string, name BeaconUserName, address BeaconUserAddress, emailAddress NullableString, phoneNumber NullableString, idNumber NullableBeaconUserIDNumber, ipAddress NullableString, depositoryAccounts []BeaconUserDepositoryAccount) *BeaconUserData {
	this := BeaconUserData{}
	this.DateOfBirth = dateOfBirth
	this.Name = name
	this.Address = address
	this.EmailAddress = emailAddress
	this.PhoneNumber = phoneNumber
	this.IdNumber = idNumber
	this.IpAddress = ipAddress
	this.DepositoryAccounts = depositoryAccounts
	return &this
}

// NewBeaconUserDataWithDefaults instantiates a new BeaconUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeaconUserDataWithDefaults() *BeaconUserData {
	this := BeaconUserData{}
	return &this
}

// GetDateOfBirth returns the DateOfBirth field value
func (o *BeaconUserData) GetDateOfBirth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
func (o *BeaconUserData) GetDateOfBirthOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DateOfBirth, true
}

// SetDateOfBirth sets field value
func (o *BeaconUserData) SetDateOfBirth(v string) {
	o.DateOfBirth = v
}

// GetName returns the Name field value
func (o *BeaconUserData) GetName() BeaconUserName {
	if o == nil {
		var ret BeaconUserName
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BeaconUserData) GetNameOk() (*BeaconUserName, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BeaconUserData) SetName(v BeaconUserName) {
	o.Name = v
}

// GetAddress returns the Address field value
func (o *BeaconUserData) GetAddress() BeaconUserAddress {
	if o == nil {
		var ret BeaconUserAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *BeaconUserData) GetAddressOk() (*BeaconUserAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *BeaconUserData) SetAddress(v BeaconUserAddress) {
	o.Address = v
}

// GetEmailAddress returns the EmailAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BeaconUserData) GetEmailAddress() string {
	if o == nil || o.EmailAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmailAddress.Get()
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BeaconUserData) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmailAddress.Get(), o.EmailAddress.IsSet()
}

// SetEmailAddress sets field value
func (o *BeaconUserData) SetEmailAddress(v string) {
	o.EmailAddress.Set(&v)
}

// GetPhoneNumber returns the PhoneNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BeaconUserData) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BeaconUserData) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// SetPhoneNumber sets field value
func (o *BeaconUserData) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}

// GetIdNumber returns the IdNumber field value
// If the value is explicit nil, the zero value for BeaconUserIDNumber will be returned
func (o *BeaconUserData) GetIdNumber() BeaconUserIDNumber {
	if o == nil || o.IdNumber.Get() == nil {
		var ret BeaconUserIDNumber
		return ret
	}

	return *o.IdNumber.Get()
}

// GetIdNumberOk returns a tuple with the IdNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BeaconUserData) GetIdNumberOk() (*BeaconUserIDNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdNumber.Get(), o.IdNumber.IsSet()
}

// SetIdNumber sets field value
func (o *BeaconUserData) SetIdNumber(v BeaconUserIDNumber) {
	o.IdNumber.Set(&v)
}

// GetIpAddress returns the IpAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BeaconUserData) GetIpAddress() string {
	if o == nil || o.IpAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BeaconUserData) GetIpAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// SetIpAddress sets field value
func (o *BeaconUserData) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}

// GetDepositoryAccounts returns the DepositoryAccounts field value
func (o *BeaconUserData) GetDepositoryAccounts() []BeaconUserDepositoryAccount {
	if o == nil {
		var ret []BeaconUserDepositoryAccount
		return ret
	}

	return o.DepositoryAccounts
}

// GetDepositoryAccountsOk returns a tuple with the DepositoryAccounts field value
// and a boolean to check if the value has been set.
func (o *BeaconUserData) GetDepositoryAccountsOk() (*[]BeaconUserDepositoryAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DepositoryAccounts, true
}

// SetDepositoryAccounts sets field value
func (o *BeaconUserData) SetDepositoryAccounts(v []BeaconUserDepositoryAccount) {
	o.DepositoryAccounts = v
}

func (o BeaconUserData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["email_address"] = o.EmailAddress.Get()
	}
	if true {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if true {
		toSerialize["id_number"] = o.IdNumber.Get()
	}
	if true {
		toSerialize["ip_address"] = o.IpAddress.Get()
	}
	if true {
		toSerialize["depository_accounts"] = o.DepositoryAccounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BeaconUserData) UnmarshalJSON(bytes []byte) (err error) {
	varBeaconUserData := _BeaconUserData{}

	if err = json.Unmarshal(bytes, &varBeaconUserData); err == nil {
		*o = BeaconUserData(varBeaconUserData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "date_of_birth")
		delete(additionalProperties, "name")
		delete(additionalProperties, "address")
		delete(additionalProperties, "email_address")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "id_number")
		delete(additionalProperties, "ip_address")
		delete(additionalProperties, "depository_accounts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBeaconUserData struct {
	value *BeaconUserData
	isSet bool
}

func (v NullableBeaconUserData) Get() *BeaconUserData {
	return v.value
}

func (v *NullableBeaconUserData) Set(val *BeaconUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconUserData(val *BeaconUserData) *NullableBeaconUserData {
	return &NullableBeaconUserData{value: val, isSet: true}
}

func (v NullableBeaconUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


