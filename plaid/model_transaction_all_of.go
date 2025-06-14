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
	"time"
)

// TransactionAllOf struct for TransactionAllOf
type TransactionAllOf struct {
	// The date that the transaction was authorized. For posted transactions, the `date` field will indicate the posted date, but `authorized_date` will indicate the day the transaction was authorized by the financial institution. If presenting transactions to the user in a UI, the `authorized_date`, when available, is generally preferable to use over the `date` field for posted transactions, as it will generally represent the date the user actually made the transaction. Dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DD` ).
	AuthorizedDate NullableString `json:"authorized_date"`
	// Date and time when a transaction was authorized in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DDTHH:mm:ssZ` ). For posted transactions, the `datetime` field will indicate the posted date, but `authorized_datetime` will indicate the day the transaction was authorized by the financial institution. If presenting transactions to the user in a UI, the `authorized_datetime`, when available, is generally preferable to use over the `datetime` field for posted transactions, as it will generally represent the date the user actually made the transaction.  This field is returned for select financial institutions and comes as provided by the institution. It may contain default time values (such as 00:00:00). This field is only populated in API version 2019-05-29 and later.
	AuthorizedDatetime NullableTime `json:"authorized_datetime"`
	// Date and time when a transaction was posted in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DDTHH:mm:ssZ` ). For the date that the transaction was initiated, rather than posted, see the `authorized_datetime` field.  This field is returned for select financial institutions and comes as provided by the institution. It may contain default time values (such as 00:00:00). This field is only populated in API version 2019-05-29 and later.
	Datetime NullableTime `json:"datetime"`
	// The channel used to make a payment. `online:` transactions that took place online.  `in store:` transactions that were made at a physical location.  `other:` transactions that relate to banks, e.g. fees or deposits.  This field replaces the `transaction_type` field. 
	PaymentChannel string `json:"payment_channel"`
	PersonalFinanceCategory NullablePersonalFinanceCategory `json:"personal_finance_category,omitempty"`
	BusinessFinanceCategory NullableBusinessFinanceCategory `json:"business_finance_category,omitempty"`
	TransactionCode NullableTransactionCode `json:"transaction_code"`
	// The URL of an icon associated with the primary personal finance category. The icon will always be 100×100 pixel PNG file.
	PersonalFinanceCategoryIconUrl *string `json:"personal_finance_category_icon_url,omitempty"`
	// The counterparties present in the transaction. Counterparties, such as the merchant or the financial institution, are extracted by Plaid from the raw description.
	Counterparties *[]TransactionCounterparty `json:"counterparties,omitempty"`
	// A unique, stable, Plaid-generated ID that maps to the merchant. In the case of a merchant with multiple retail locations, this field will map to the broader merchant, not a specific location or store.
	MerchantEntityId NullableString `json:"merchant_entity_id,omitempty"`
}

// NewTransactionAllOf instantiates a new TransactionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionAllOf(authorizedDate NullableString, authorizedDatetime NullableTime, datetime NullableTime, paymentChannel string, transactionCode NullableTransactionCode) *TransactionAllOf {
	this := TransactionAllOf{}
	this.AuthorizedDate = authorizedDate
	this.AuthorizedDatetime = authorizedDatetime
	this.Datetime = datetime
	this.PaymentChannel = paymentChannel
	this.TransactionCode = transactionCode
	return &this
}

// NewTransactionAllOfWithDefaults instantiates a new TransactionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionAllOfWithDefaults() *TransactionAllOf {
	this := TransactionAllOf{}
	return &this
}

// GetAuthorizedDate returns the AuthorizedDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransactionAllOf) GetAuthorizedDate() string {
	if o == nil || o.AuthorizedDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthorizedDate.Get()
}

// GetAuthorizedDateOk returns a tuple with the AuthorizedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetAuthorizedDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthorizedDate.Get(), o.AuthorizedDate.IsSet()
}

// SetAuthorizedDate sets field value
func (o *TransactionAllOf) SetAuthorizedDate(v string) {
	o.AuthorizedDate.Set(&v)
}

// GetAuthorizedDatetime returns the AuthorizedDatetime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TransactionAllOf) GetAuthorizedDatetime() time.Time {
	if o == nil || o.AuthorizedDatetime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.AuthorizedDatetime.Get()
}

// GetAuthorizedDatetimeOk returns a tuple with the AuthorizedDatetime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetAuthorizedDatetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthorizedDatetime.Get(), o.AuthorizedDatetime.IsSet()
}

// SetAuthorizedDatetime sets field value
func (o *TransactionAllOf) SetAuthorizedDatetime(v time.Time) {
	o.AuthorizedDatetime.Set(&v)
}

// GetDatetime returns the Datetime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TransactionAllOf) GetDatetime() time.Time {
	if o == nil || o.Datetime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Datetime.Get()
}

// GetDatetimeOk returns a tuple with the Datetime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetDatetimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Datetime.Get(), o.Datetime.IsSet()
}

// SetDatetime sets field value
func (o *TransactionAllOf) SetDatetime(v time.Time) {
	o.Datetime.Set(&v)
}

// GetPaymentChannel returns the PaymentChannel field value
func (o *TransactionAllOf) GetPaymentChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentChannel
}

// GetPaymentChannelOk returns a tuple with the PaymentChannel field value
// and a boolean to check if the value has been set.
func (o *TransactionAllOf) GetPaymentChannelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentChannel, true
}

// SetPaymentChannel sets field value
func (o *TransactionAllOf) SetPaymentChannel(v string) {
	o.PaymentChannel = v
}

// GetPersonalFinanceCategory returns the PersonalFinanceCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionAllOf) GetPersonalFinanceCategory() PersonalFinanceCategory {
	if o == nil || o.PersonalFinanceCategory.Get() == nil {
		var ret PersonalFinanceCategory
		return ret
	}
	return *o.PersonalFinanceCategory.Get()
}

// GetPersonalFinanceCategoryOk returns a tuple with the PersonalFinanceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetPersonalFinanceCategoryOk() (*PersonalFinanceCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PersonalFinanceCategory.Get(), o.PersonalFinanceCategory.IsSet()
}

// HasPersonalFinanceCategory returns a boolean if a field has been set.
func (o *TransactionAllOf) HasPersonalFinanceCategory() bool {
	if o != nil && o.PersonalFinanceCategory.IsSet() {
		return true
	}

	return false
}

// SetPersonalFinanceCategory gets a reference to the given NullablePersonalFinanceCategory and assigns it to the PersonalFinanceCategory field.
func (o *TransactionAllOf) SetPersonalFinanceCategory(v PersonalFinanceCategory) {
	o.PersonalFinanceCategory.Set(&v)
}
// SetPersonalFinanceCategoryNil sets the value for PersonalFinanceCategory to be an explicit nil
func (o *TransactionAllOf) SetPersonalFinanceCategoryNil() {
	o.PersonalFinanceCategory.Set(nil)
}

// UnsetPersonalFinanceCategory ensures that no value is present for PersonalFinanceCategory, not even an explicit nil
func (o *TransactionAllOf) UnsetPersonalFinanceCategory() {
	o.PersonalFinanceCategory.Unset()
}

// GetBusinessFinanceCategory returns the BusinessFinanceCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionAllOf) GetBusinessFinanceCategory() BusinessFinanceCategory {
	if o == nil || o.BusinessFinanceCategory.Get() == nil {
		var ret BusinessFinanceCategory
		return ret
	}
	return *o.BusinessFinanceCategory.Get()
}

// GetBusinessFinanceCategoryOk returns a tuple with the BusinessFinanceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetBusinessFinanceCategoryOk() (*BusinessFinanceCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BusinessFinanceCategory.Get(), o.BusinessFinanceCategory.IsSet()
}

// HasBusinessFinanceCategory returns a boolean if a field has been set.
func (o *TransactionAllOf) HasBusinessFinanceCategory() bool {
	if o != nil && o.BusinessFinanceCategory.IsSet() {
		return true
	}

	return false
}

// SetBusinessFinanceCategory gets a reference to the given NullableBusinessFinanceCategory and assigns it to the BusinessFinanceCategory field.
func (o *TransactionAllOf) SetBusinessFinanceCategory(v BusinessFinanceCategory) {
	o.BusinessFinanceCategory.Set(&v)
}
// SetBusinessFinanceCategoryNil sets the value for BusinessFinanceCategory to be an explicit nil
func (o *TransactionAllOf) SetBusinessFinanceCategoryNil() {
	o.BusinessFinanceCategory.Set(nil)
}

// UnsetBusinessFinanceCategory ensures that no value is present for BusinessFinanceCategory, not even an explicit nil
func (o *TransactionAllOf) UnsetBusinessFinanceCategory() {
	o.BusinessFinanceCategory.Unset()
}

// GetTransactionCode returns the TransactionCode field value
// If the value is explicit nil, the zero value for TransactionCode will be returned
func (o *TransactionAllOf) GetTransactionCode() TransactionCode {
	if o == nil || o.TransactionCode.Get() == nil {
		var ret TransactionCode
		return ret
	}

	return *o.TransactionCode.Get()
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetTransactionCodeOk() (*TransactionCode, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransactionCode.Get(), o.TransactionCode.IsSet()
}

// SetTransactionCode sets field value
func (o *TransactionAllOf) SetTransactionCode(v TransactionCode) {
	o.TransactionCode.Set(&v)
}

// GetPersonalFinanceCategoryIconUrl returns the PersonalFinanceCategoryIconUrl field value if set, zero value otherwise.
func (o *TransactionAllOf) GetPersonalFinanceCategoryIconUrl() string {
	if o == nil || o.PersonalFinanceCategoryIconUrl == nil {
		var ret string
		return ret
	}
	return *o.PersonalFinanceCategoryIconUrl
}

// GetPersonalFinanceCategoryIconUrlOk returns a tuple with the PersonalFinanceCategoryIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionAllOf) GetPersonalFinanceCategoryIconUrlOk() (*string, bool) {
	if o == nil || o.PersonalFinanceCategoryIconUrl == nil {
		return nil, false
	}
	return o.PersonalFinanceCategoryIconUrl, true
}

// HasPersonalFinanceCategoryIconUrl returns a boolean if a field has been set.
func (o *TransactionAllOf) HasPersonalFinanceCategoryIconUrl() bool {
	if o != nil && o.PersonalFinanceCategoryIconUrl != nil {
		return true
	}

	return false
}

// SetPersonalFinanceCategoryIconUrl gets a reference to the given string and assigns it to the PersonalFinanceCategoryIconUrl field.
func (o *TransactionAllOf) SetPersonalFinanceCategoryIconUrl(v string) {
	o.PersonalFinanceCategoryIconUrl = &v
}

// GetCounterparties returns the Counterparties field value if set, zero value otherwise.
func (o *TransactionAllOf) GetCounterparties() []TransactionCounterparty {
	if o == nil || o.Counterparties == nil {
		var ret []TransactionCounterparty
		return ret
	}
	return *o.Counterparties
}

// GetCounterpartiesOk returns a tuple with the Counterparties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionAllOf) GetCounterpartiesOk() (*[]TransactionCounterparty, bool) {
	if o == nil || o.Counterparties == nil {
		return nil, false
	}
	return o.Counterparties, true
}

// HasCounterparties returns a boolean if a field has been set.
func (o *TransactionAllOf) HasCounterparties() bool {
	if o != nil && o.Counterparties != nil {
		return true
	}

	return false
}

// SetCounterparties gets a reference to the given []TransactionCounterparty and assigns it to the Counterparties field.
func (o *TransactionAllOf) SetCounterparties(v []TransactionCounterparty) {
	o.Counterparties = &v
}

// GetMerchantEntityId returns the MerchantEntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionAllOf) GetMerchantEntityId() string {
	if o == nil || o.MerchantEntityId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantEntityId.Get()
}

// GetMerchantEntityIdOk returns a tuple with the MerchantEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionAllOf) GetMerchantEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantEntityId.Get(), o.MerchantEntityId.IsSet()
}

// HasMerchantEntityId returns a boolean if a field has been set.
func (o *TransactionAllOf) HasMerchantEntityId() bool {
	if o != nil && o.MerchantEntityId.IsSet() {
		return true
	}

	return false
}

// SetMerchantEntityId gets a reference to the given NullableString and assigns it to the MerchantEntityId field.
func (o *TransactionAllOf) SetMerchantEntityId(v string) {
	o.MerchantEntityId.Set(&v)
}
// SetMerchantEntityIdNil sets the value for MerchantEntityId to be an explicit nil
func (o *TransactionAllOf) SetMerchantEntityIdNil() {
	o.MerchantEntityId.Set(nil)
}

// UnsetMerchantEntityId ensures that no value is present for MerchantEntityId, not even an explicit nil
func (o *TransactionAllOf) UnsetMerchantEntityId() {
	o.MerchantEntityId.Unset()
}

func (o TransactionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authorized_date"] = o.AuthorizedDate.Get()
	}
	if true {
		toSerialize["authorized_datetime"] = o.AuthorizedDatetime.Get()
	}
	if true {
		toSerialize["datetime"] = o.Datetime.Get()
	}
	if true {
		toSerialize["payment_channel"] = o.PaymentChannel
	}
	if o.PersonalFinanceCategory.IsSet() {
		toSerialize["personal_finance_category"] = o.PersonalFinanceCategory.Get()
	}
	if o.BusinessFinanceCategory.IsSet() {
		toSerialize["business_finance_category"] = o.BusinessFinanceCategory.Get()
	}
	if true {
		toSerialize["transaction_code"] = o.TransactionCode.Get()
	}
	if o.PersonalFinanceCategoryIconUrl != nil {
		toSerialize["personal_finance_category_icon_url"] = o.PersonalFinanceCategoryIconUrl
	}
	if o.Counterparties != nil {
		toSerialize["counterparties"] = o.Counterparties
	}
	if o.MerchantEntityId.IsSet() {
		toSerialize["merchant_entity_id"] = o.MerchantEntityId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionAllOf struct {
	value *TransactionAllOf
	isSet bool
}

func (v NullableTransactionAllOf) Get() *TransactionAllOf {
	return v.value
}

func (v *NullableTransactionAllOf) Set(val *TransactionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionAllOf(val *TransactionAllOf) *NullableTransactionAllOf {
	return &NullableTransactionAllOf{value: val, isSet: true}
}

func (v NullableTransactionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


