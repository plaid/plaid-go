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

// TransactionsCategoryRule A representation of a transactions category rule.
type TransactionsCategoryRule struct {
	// A unique identifier of the rule created
	Id *string `json:"id,omitempty"`
	// A unique identifier of the Item the rule was created for.
	ItemId *string `json:"item_id,omitempty"`
	// Date and time when a rule was created in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DDTHH:mm:ssZ` ). 
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Personal finance category unique identifier.  In the personal finance category taxonomy, this field is represented by the detailed category field. 
	PersonalFinanceCategory *string `json:"personal_finance_category,omitempty"`
	RuleDetails *TransactionsRuleDetails `json:"rule_details,omitempty"`
}

// NewTransactionsCategoryRule instantiates a new TransactionsCategoryRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsCategoryRule() *TransactionsCategoryRule {
	this := TransactionsCategoryRule{}
	return &this
}

// NewTransactionsCategoryRuleWithDefaults instantiates a new TransactionsCategoryRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsCategoryRuleWithDefaults() *TransactionsCategoryRule {
	this := TransactionsCategoryRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionsCategoryRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsCategoryRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionsCategoryRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransactionsCategoryRule) SetId(v string) {
	o.Id = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *TransactionsCategoryRule) GetItemId() string {
	if o == nil || o.ItemId == nil {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsCategoryRule) GetItemIdOk() (*string, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *TransactionsCategoryRule) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *TransactionsCategoryRule) SetItemId(v string) {
	o.ItemId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TransactionsCategoryRule) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsCategoryRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TransactionsCategoryRule) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TransactionsCategoryRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetPersonalFinanceCategory returns the PersonalFinanceCategory field value if set, zero value otherwise.
func (o *TransactionsCategoryRule) GetPersonalFinanceCategory() string {
	if o == nil || o.PersonalFinanceCategory == nil {
		var ret string
		return ret
	}
	return *o.PersonalFinanceCategory
}

// GetPersonalFinanceCategoryOk returns a tuple with the PersonalFinanceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsCategoryRule) GetPersonalFinanceCategoryOk() (*string, bool) {
	if o == nil || o.PersonalFinanceCategory == nil {
		return nil, false
	}
	return o.PersonalFinanceCategory, true
}

// HasPersonalFinanceCategory returns a boolean if a field has been set.
func (o *TransactionsCategoryRule) HasPersonalFinanceCategory() bool {
	if o != nil && o.PersonalFinanceCategory != nil {
		return true
	}

	return false
}

// SetPersonalFinanceCategory gets a reference to the given string and assigns it to the PersonalFinanceCategory field.
func (o *TransactionsCategoryRule) SetPersonalFinanceCategory(v string) {
	o.PersonalFinanceCategory = &v
}

// GetRuleDetails returns the RuleDetails field value if set, zero value otherwise.
func (o *TransactionsCategoryRule) GetRuleDetails() TransactionsRuleDetails {
	if o == nil || o.RuleDetails == nil {
		var ret TransactionsRuleDetails
		return ret
	}
	return *o.RuleDetails
}

// GetRuleDetailsOk returns a tuple with the RuleDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsCategoryRule) GetRuleDetailsOk() (*TransactionsRuleDetails, bool) {
	if o == nil || o.RuleDetails == nil {
		return nil, false
	}
	return o.RuleDetails, true
}

// HasRuleDetails returns a boolean if a field has been set.
func (o *TransactionsCategoryRule) HasRuleDetails() bool {
	if o != nil && o.RuleDetails != nil {
		return true
	}

	return false
}

// SetRuleDetails gets a reference to the given TransactionsRuleDetails and assigns it to the RuleDetails field.
func (o *TransactionsCategoryRule) SetRuleDetails(v TransactionsRuleDetails) {
	o.RuleDetails = &v
}

func (o TransactionsCategoryRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ItemId != nil {
		toSerialize["item_id"] = o.ItemId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.PersonalFinanceCategory != nil {
		toSerialize["personal_finance_category"] = o.PersonalFinanceCategory
	}
	if o.RuleDetails != nil {
		toSerialize["rule_details"] = o.RuleDetails
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsCategoryRule struct {
	value *TransactionsCategoryRule
	isSet bool
}

func (v NullableTransactionsCategoryRule) Get() *TransactionsCategoryRule {
	return v.value
}

func (v *NullableTransactionsCategoryRule) Set(val *TransactionsCategoryRule) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsCategoryRule) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsCategoryRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsCategoryRule(val *TransactionsCategoryRule) *NullableTransactionsCategoryRule {
	return &NullableTransactionsCategoryRule{value: val, isSet: true}
}

func (v NullableTransactionsCategoryRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsCategoryRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


