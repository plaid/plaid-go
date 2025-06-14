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

// CreditBankIncomeAccount The Item's bank accounts that have the selected data.
type CreditBankIncomeAccount struct {
	// Plaid's unique identifier for the account.
	AccountId string `json:"account_id"`
	// The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.
	Mask NullableString `json:"mask"`
	// The name of the bank account.
	Name string `json:"name"`
	// The official name of the bank account.
	OfficialName NullableString `json:"official_name"`
	Subtype DepositoryAccountSubtype `json:"subtype"`
	Type CreditBankIncomeAccountType `json:"type"`
	// Data returned by the financial institution about the account owner or owners. Identity information is optional, so field may return an empty array.
	Owners []Owner `json:"owners"`
}

// NewCreditBankIncomeAccount instantiates a new CreditBankIncomeAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeAccount(accountId string, mask NullableString, name string, officialName NullableString, subtype DepositoryAccountSubtype, type_ CreditBankIncomeAccountType, owners []Owner) *CreditBankIncomeAccount {
	this := CreditBankIncomeAccount{}
	this.AccountId = accountId
	this.Mask = mask
	this.Name = name
	this.OfficialName = officialName
	this.Subtype = subtype
	this.Type = type_
	this.Owners = owners
	return &this
}

// NewCreditBankIncomeAccountWithDefaults instantiates a new CreditBankIncomeAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeAccountWithDefaults() *CreditBankIncomeAccount {
	this := CreditBankIncomeAccount{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *CreditBankIncomeAccount) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeAccount) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CreditBankIncomeAccount) SetAccountId(v string) {
	o.AccountId = v
}

// GetMask returns the Mask field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditBankIncomeAccount) GetMask() string {
	if o == nil || o.Mask.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankIncomeAccount) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// SetMask sets field value
func (o *CreditBankIncomeAccount) SetMask(v string) {
	o.Mask.Set(&v)
}

// GetName returns the Name field value
func (o *CreditBankIncomeAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeAccount) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreditBankIncomeAccount) SetName(v string) {
	o.Name = v
}

// GetOfficialName returns the OfficialName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditBankIncomeAccount) GetOfficialName() string {
	if o == nil || o.OfficialName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OfficialName.Get()
}

// GetOfficialNameOk returns a tuple with the OfficialName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditBankIncomeAccount) GetOfficialNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OfficialName.Get(), o.OfficialName.IsSet()
}

// SetOfficialName sets field value
func (o *CreditBankIncomeAccount) SetOfficialName(v string) {
	o.OfficialName.Set(&v)
}

// GetSubtype returns the Subtype field value
func (o *CreditBankIncomeAccount) GetSubtype() DepositoryAccountSubtype {
	if o == nil {
		var ret DepositoryAccountSubtype
		return ret
	}

	return o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeAccount) GetSubtypeOk() (*DepositoryAccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subtype, true
}

// SetSubtype sets field value
func (o *CreditBankIncomeAccount) SetSubtype(v DepositoryAccountSubtype) {
	o.Subtype = v
}

// GetType returns the Type field value
func (o *CreditBankIncomeAccount) GetType() CreditBankIncomeAccountType {
	if o == nil {
		var ret CreditBankIncomeAccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeAccount) GetTypeOk() (*CreditBankIncomeAccountType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreditBankIncomeAccount) SetType(v CreditBankIncomeAccountType) {
	o.Type = v
}

// GetOwners returns the Owners field value
func (o *CreditBankIncomeAccount) GetOwners() []Owner {
	if o == nil {
		var ret []Owner
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeAccount) GetOwnersOk() (*[]Owner, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value
func (o *CreditBankIncomeAccount) SetOwners(v []Owner) {
	o.Owners = v
}

func (o CreditBankIncomeAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["mask"] = o.Mask.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["official_name"] = o.OfficialName.Get()
	}
	if true {
		toSerialize["subtype"] = o.Subtype
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["owners"] = o.Owners
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankIncomeAccount struct {
	value *CreditBankIncomeAccount
	isSet bool
}

func (v NullableCreditBankIncomeAccount) Get() *CreditBankIncomeAccount {
	return v.value
}

func (v *NullableCreditBankIncomeAccount) Set(val *CreditBankIncomeAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeAccount(val *CreditBankIncomeAccount) *NullableCreditBankIncomeAccount {
	return &NullableCreditBankIncomeAccount{value: val, isSet: true}
}

func (v NullableCreditBankIncomeAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


