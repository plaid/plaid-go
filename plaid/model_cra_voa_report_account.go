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

// CraVoaReportAccount VOA Report information about an account.
type CraVoaReportAccount struct {
	// Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  Like all Plaid identifiers, the `account_id` is case sensitive.
	AccountId string `json:"account_id"`
	Balances CraVoaReportAccountBalances `json:"balances"`
	// The information about previously submitted valid dispute statements by the consumer
	ConsumerDisputes []ConsumerDispute `json:"consumer_disputes"`
	// The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.
	Mask NullableString `json:"mask"`
	// The name of the account, either assigned by the user or by the financial institution itself.
	Name string `json:"name"`
	// The official name of the account as given by the financial institution.
	OfficialName NullableString `json:"official_name"`
	Type AccountType `json:"type"`
	Subtype NullableAccountSubtype `json:"subtype"`
	// The duration of transaction history available within this report for this Item, typically defined as the time since the date of the earliest transaction in that account.
	DaysAvailable float32 `json:"days_available"`
	TransactionsInsights CraVoaReportTransactionsInsights `json:"transactions_insights"`
	// Data returned by the financial institution about the account owner or owners.
	Owners []Owner `json:"owners"`
	OwnershipType NullableOwnershipType `json:"ownership_type"`
	AdditionalProperties map[string]interface{}
}

type _CraVoaReportAccount CraVoaReportAccount

// NewCraVoaReportAccount instantiates a new CraVoaReportAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraVoaReportAccount(accountId string, balances CraVoaReportAccountBalances, consumerDisputes []ConsumerDispute, mask NullableString, name string, officialName NullableString, type_ AccountType, subtype NullableAccountSubtype, daysAvailable float32, transactionsInsights CraVoaReportTransactionsInsights, owners []Owner, ownershipType NullableOwnershipType) *CraVoaReportAccount {
	this := CraVoaReportAccount{}
	this.AccountId = accountId
	this.Balances = balances
	this.ConsumerDisputes = consumerDisputes
	this.Mask = mask
	this.Name = name
	this.OfficialName = officialName
	this.Type = type_
	this.Subtype = subtype
	this.DaysAvailable = daysAvailable
	this.TransactionsInsights = transactionsInsights
	this.Owners = owners
	this.OwnershipType = ownershipType
	return &this
}

// NewCraVoaReportAccountWithDefaults instantiates a new CraVoaReportAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraVoaReportAccountWithDefaults() *CraVoaReportAccount {
	this := CraVoaReportAccount{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *CraVoaReportAccount) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CraVoaReportAccount) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CraVoaReportAccount) SetAccountId(v string) {
	o.AccountId = v
}

// GetBalances returns the Balances field value
func (o *CraVoaReportAccount) GetBalances() CraVoaReportAccountBalances {
	if o == nil {
		var ret CraVoaReportAccountBalances
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *CraVoaReportAccount) GetBalancesOk() (*CraVoaReportAccountBalances, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *CraVoaReportAccount) SetBalances(v CraVoaReportAccountBalances) {
	o.Balances = v
}

// GetConsumerDisputes returns the ConsumerDisputes field value
func (o *CraVoaReportAccount) GetConsumerDisputes() []ConsumerDispute {
	if o == nil {
		var ret []ConsumerDispute
		return ret
	}

	return o.ConsumerDisputes
}

// GetConsumerDisputesOk returns a tuple with the ConsumerDisputes field value
// and a boolean to check if the value has been set.
func (o *CraVoaReportAccount) GetConsumerDisputesOk() (*[]ConsumerDispute, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsumerDisputes, true
}

// SetConsumerDisputes sets field value
func (o *CraVoaReportAccount) SetConsumerDisputes(v []ConsumerDispute) {
	o.ConsumerDisputes = v
}

// GetMask returns the Mask field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CraVoaReportAccount) GetMask() string {
	if o == nil || o.Mask.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraVoaReportAccount) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// SetMask sets field value
func (o *CraVoaReportAccount) SetMask(v string) {
	o.Mask.Set(&v)
}

// GetName returns the Name field value
func (o *CraVoaReportAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CraVoaReportAccount) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CraVoaReportAccount) SetName(v string) {
	o.Name = v
}

// GetOfficialName returns the OfficialName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CraVoaReportAccount) GetOfficialName() string {
	if o == nil || o.OfficialName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OfficialName.Get()
}

// GetOfficialNameOk returns a tuple with the OfficialName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraVoaReportAccount) GetOfficialNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OfficialName.Get(), o.OfficialName.IsSet()
}

// SetOfficialName sets field value
func (o *CraVoaReportAccount) SetOfficialName(v string) {
	o.OfficialName.Set(&v)
}

// GetType returns the Type field value
func (o *CraVoaReportAccount) GetType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CraVoaReportAccount) GetTypeOk() (*AccountType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CraVoaReportAccount) SetType(v AccountType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value
// If the value is explicit nil, the zero value for AccountSubtype will be returned
func (o *CraVoaReportAccount) GetSubtype() AccountSubtype {
	if o == nil || o.Subtype.Get() == nil {
		var ret AccountSubtype
		return ret
	}

	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraVoaReportAccount) GetSubtypeOk() (*AccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// SetSubtype sets field value
func (o *CraVoaReportAccount) SetSubtype(v AccountSubtype) {
	o.Subtype.Set(&v)
}

// GetDaysAvailable returns the DaysAvailable field value
func (o *CraVoaReportAccount) GetDaysAvailable() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DaysAvailable
}

// GetDaysAvailableOk returns a tuple with the DaysAvailable field value
// and a boolean to check if the value has been set.
func (o *CraVoaReportAccount) GetDaysAvailableOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DaysAvailable, true
}

// SetDaysAvailable sets field value
func (o *CraVoaReportAccount) SetDaysAvailable(v float32) {
	o.DaysAvailable = v
}

// GetTransactionsInsights returns the TransactionsInsights field value
func (o *CraVoaReportAccount) GetTransactionsInsights() CraVoaReportTransactionsInsights {
	if o == nil {
		var ret CraVoaReportTransactionsInsights
		return ret
	}

	return o.TransactionsInsights
}

// GetTransactionsInsightsOk returns a tuple with the TransactionsInsights field value
// and a boolean to check if the value has been set.
func (o *CraVoaReportAccount) GetTransactionsInsightsOk() (*CraVoaReportTransactionsInsights, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionsInsights, true
}

// SetTransactionsInsights sets field value
func (o *CraVoaReportAccount) SetTransactionsInsights(v CraVoaReportTransactionsInsights) {
	o.TransactionsInsights = v
}

// GetOwners returns the Owners field value
func (o *CraVoaReportAccount) GetOwners() []Owner {
	if o == nil {
		var ret []Owner
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *CraVoaReportAccount) GetOwnersOk() (*[]Owner, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value
func (o *CraVoaReportAccount) SetOwners(v []Owner) {
	o.Owners = v
}

// GetOwnershipType returns the OwnershipType field value
// If the value is explicit nil, the zero value for OwnershipType will be returned
func (o *CraVoaReportAccount) GetOwnershipType() OwnershipType {
	if o == nil || o.OwnershipType.Get() == nil {
		var ret OwnershipType
		return ret
	}

	return *o.OwnershipType.Get()
}

// GetOwnershipTypeOk returns a tuple with the OwnershipType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraVoaReportAccount) GetOwnershipTypeOk() (*OwnershipType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OwnershipType.Get(), o.OwnershipType.IsSet()
}

// SetOwnershipType sets field value
func (o *CraVoaReportAccount) SetOwnershipType(v OwnershipType) {
	o.OwnershipType.Set(&v)
}

func (o CraVoaReportAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["balances"] = o.Balances
	}
	if true {
		toSerialize["consumer_disputes"] = o.ConsumerDisputes
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
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	if true {
		toSerialize["days_available"] = o.DaysAvailable
	}
	if true {
		toSerialize["transactions_insights"] = o.TransactionsInsights
	}
	if true {
		toSerialize["owners"] = o.Owners
	}
	if true {
		toSerialize["ownership_type"] = o.OwnershipType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraVoaReportAccount) UnmarshalJSON(bytes []byte) (err error) {
	varCraVoaReportAccount := _CraVoaReportAccount{}

	if err = json.Unmarshal(bytes, &varCraVoaReportAccount); err == nil {
		*o = CraVoaReportAccount(varCraVoaReportAccount)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "balances")
		delete(additionalProperties, "consumer_disputes")
		delete(additionalProperties, "mask")
		delete(additionalProperties, "name")
		delete(additionalProperties, "official_name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "days_available")
		delete(additionalProperties, "transactions_insights")
		delete(additionalProperties, "owners")
		delete(additionalProperties, "ownership_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraVoaReportAccount struct {
	value *CraVoaReportAccount
	isSet bool
}

func (v NullableCraVoaReportAccount) Get() *CraVoaReportAccount {
	return v.value
}

func (v *NullableCraVoaReportAccount) Set(val *CraVoaReportAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCraVoaReportAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCraVoaReportAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraVoaReportAccount(val *CraVoaReportAccount) *NullableCraVoaReportAccount {
	return &NullableCraVoaReportAccount{value: val, isSet: true}
}

func (v NullableCraVoaReportAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraVoaReportAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


