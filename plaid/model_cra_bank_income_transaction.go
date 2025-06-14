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

// CraBankIncomeTransaction The transactions data for the end user's income source(s).
type CraBankIncomeTransaction struct {
	// The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.
	TransactionId string `json:"transaction_id"`
	// The settled value of the transaction, denominated in the transaction's currency as stated in `iso_currency_code` or `unofficial_currency_code`. Positive values when money moves out of the account; negative values when money moves in. For example, credit card purchases are positive; credit card payment, direct deposits, and refunds are negative.
	Amount float32 `json:"amount"`
	// For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted. Both dates are returned in an ISO 8601 format (YYYY-MM-DD).
	Date string `json:"date"`
	// The merchant name or transaction description.
	Name *string `json:"name,omitempty"`
	// The string returned by the financial institution to describe the transaction.
	OriginalDescription NullableString `json:"original_description"`
	// When true, identifies the transaction as pending or unsettled. Pending transaction details (name, type, amount, category ID) may change before they are settled.
	Pending bool `json:"pending"`
	// The check number of the transaction. This field is only populated for check transactions.
	CheckNumber NullableString `json:"check_number,omitempty"`
	// The ISO 4217 currency code of the amount or balance.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the amount or balance. Always `null` if `iso_currency_code` is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	BonusType NullableCraBankIncomeBonusType `json:"bonus_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CraBankIncomeTransaction CraBankIncomeTransaction

// NewCraBankIncomeTransaction instantiates a new CraBankIncomeTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraBankIncomeTransaction(transactionId string, amount float32, date string, originalDescription NullableString, pending bool, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString) *CraBankIncomeTransaction {
	this := CraBankIncomeTransaction{}
	this.TransactionId = transactionId
	this.Amount = amount
	this.Date = date
	this.OriginalDescription = originalDescription
	this.Pending = pending
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	return &this
}

// NewCraBankIncomeTransactionWithDefaults instantiates a new CraBankIncomeTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraBankIncomeTransactionWithDefaults() *CraBankIncomeTransaction {
	this := CraBankIncomeTransaction{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *CraBankIncomeTransaction) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeTransaction) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *CraBankIncomeTransaction) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetAmount returns the Amount field value
func (o *CraBankIncomeTransaction) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeTransaction) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CraBankIncomeTransaction) SetAmount(v float32) {
	o.Amount = v
}

// GetDate returns the Date field value
func (o *CraBankIncomeTransaction) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeTransaction) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *CraBankIncomeTransaction) SetDate(v string) {
	o.Date = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CraBankIncomeTransaction) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeTransaction) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CraBankIncomeTransaction) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CraBankIncomeTransaction) SetName(v string) {
	o.Name = &v
}

// GetOriginalDescription returns the OriginalDescription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CraBankIncomeTransaction) GetOriginalDescription() string {
	if o == nil || o.OriginalDescription.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginalDescription.Get()
}

// GetOriginalDescriptionOk returns a tuple with the OriginalDescription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeTransaction) GetOriginalDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalDescription.Get(), o.OriginalDescription.IsSet()
}

// SetOriginalDescription sets field value
func (o *CraBankIncomeTransaction) SetOriginalDescription(v string) {
	o.OriginalDescription.Set(&v)
}

// GetPending returns the Pending field value
func (o *CraBankIncomeTransaction) GetPending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeTransaction) GetPendingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *CraBankIncomeTransaction) SetPending(v bool) {
	o.Pending = v
}

// GetCheckNumber returns the CheckNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CraBankIncomeTransaction) GetCheckNumber() string {
	if o == nil || o.CheckNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.CheckNumber.Get()
}

// GetCheckNumberOk returns a tuple with the CheckNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeTransaction) GetCheckNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CheckNumber.Get(), o.CheckNumber.IsSet()
}

// HasCheckNumber returns a boolean if a field has been set.
func (o *CraBankIncomeTransaction) HasCheckNumber() bool {
	if o != nil && o.CheckNumber.IsSet() {
		return true
	}

	return false
}

// SetCheckNumber gets a reference to the given NullableString and assigns it to the CheckNumber field.
func (o *CraBankIncomeTransaction) SetCheckNumber(v string) {
	o.CheckNumber.Set(&v)
}
// SetCheckNumberNil sets the value for CheckNumber to be an explicit nil
func (o *CraBankIncomeTransaction) SetCheckNumberNil() {
	o.CheckNumber.Set(nil)
}

// UnsetCheckNumber ensures that no value is present for CheckNumber, not even an explicit nil
func (o *CraBankIncomeTransaction) UnsetCheckNumber() {
	o.CheckNumber.Unset()
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CraBankIncomeTransaction) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeTransaction) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *CraBankIncomeTransaction) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CraBankIncomeTransaction) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeTransaction) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *CraBankIncomeTransaction) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

// GetBonusType returns the BonusType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CraBankIncomeTransaction) GetBonusType() CraBankIncomeBonusType {
	if o == nil || o.BonusType.Get() == nil {
		var ret CraBankIncomeBonusType
		return ret
	}
	return *o.BonusType.Get()
}

// GetBonusTypeOk returns a tuple with the BonusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraBankIncomeTransaction) GetBonusTypeOk() (*CraBankIncomeBonusType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BonusType.Get(), o.BonusType.IsSet()
}

// HasBonusType returns a boolean if a field has been set.
func (o *CraBankIncomeTransaction) HasBonusType() bool {
	if o != nil && o.BonusType.IsSet() {
		return true
	}

	return false
}

// SetBonusType gets a reference to the given NullableCraBankIncomeBonusType and assigns it to the BonusType field.
func (o *CraBankIncomeTransaction) SetBonusType(v CraBankIncomeBonusType) {
	o.BonusType.Set(&v)
}
// SetBonusTypeNil sets the value for BonusType to be an explicit nil
func (o *CraBankIncomeTransaction) SetBonusTypeNil() {
	o.BonusType.Set(nil)
}

// UnsetBonusType ensures that no value is present for BonusType, not even an explicit nil
func (o *CraBankIncomeTransaction) UnsetBonusType() {
	o.BonusType.Unset()
}

func (o CraBankIncomeTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["original_description"] = o.OriginalDescription.Get()
	}
	if true {
		toSerialize["pending"] = o.Pending
	}
	if o.CheckNumber.IsSet() {
		toSerialize["check_number"] = o.CheckNumber.Get()
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if o.BonusType.IsSet() {
		toSerialize["bonus_type"] = o.BonusType.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraBankIncomeTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varCraBankIncomeTransaction := _CraBankIncomeTransaction{}

	if err = json.Unmarshal(bytes, &varCraBankIncomeTransaction); err == nil {
		*o = CraBankIncomeTransaction(varCraBankIncomeTransaction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "date")
		delete(additionalProperties, "name")
		delete(additionalProperties, "original_description")
		delete(additionalProperties, "pending")
		delete(additionalProperties, "check_number")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "bonus_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraBankIncomeTransaction struct {
	value *CraBankIncomeTransaction
	isSet bool
}

func (v NullableCraBankIncomeTransaction) Get() *CraBankIncomeTransaction {
	return v.value
}

func (v *NullableCraBankIncomeTransaction) Set(val *CraBankIncomeTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableCraBankIncomeTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableCraBankIncomeTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraBankIncomeTransaction(val *CraBankIncomeTransaction) *NullableCraBankIncomeTransaction {
	return &NullableCraBankIncomeTransaction{value: val, isSet: true}
}

func (v NullableCraBankIncomeTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraBankIncomeTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


