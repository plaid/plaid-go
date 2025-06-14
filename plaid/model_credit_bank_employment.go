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

// CreditBankEmployment Detailed information for the bank employment.
type CreditBankEmployment struct {
	// A unique identifier for the bank employment.
	BankEmploymentId string `json:"bank_employment_id"`
	// Plaid's unique identifier for the account.
	AccountId string `json:"account_id"`
	Employer CreditBankEmployer `json:"employer"`
	// The date of the most recent deposit from this employer.
	LatestDepositDate string `json:"latest_deposit_date"`
	// The date of the earliest deposit from this employer from within the period of the days requested.
	EarliestDepositDate string `json:"earliest_deposit_date"`
}

// NewCreditBankEmployment instantiates a new CreditBankEmployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankEmployment(bankEmploymentId string, accountId string, employer CreditBankEmployer, latestDepositDate string, earliestDepositDate string) *CreditBankEmployment {
	this := CreditBankEmployment{}
	this.BankEmploymentId = bankEmploymentId
	this.AccountId = accountId
	this.Employer = employer
	this.LatestDepositDate = latestDepositDate
	this.EarliestDepositDate = earliestDepositDate
	return &this
}

// NewCreditBankEmploymentWithDefaults instantiates a new CreditBankEmployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankEmploymentWithDefaults() *CreditBankEmployment {
	this := CreditBankEmployment{}
	return &this
}

// GetBankEmploymentId returns the BankEmploymentId field value
func (o *CreditBankEmployment) GetBankEmploymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankEmploymentId
}

// GetBankEmploymentIdOk returns a tuple with the BankEmploymentId field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmployment) GetBankEmploymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankEmploymentId, true
}

// SetBankEmploymentId sets field value
func (o *CreditBankEmployment) SetBankEmploymentId(v string) {
	o.BankEmploymentId = v
}

// GetAccountId returns the AccountId field value
func (o *CreditBankEmployment) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmployment) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CreditBankEmployment) SetAccountId(v string) {
	o.AccountId = v
}

// GetEmployer returns the Employer field value
func (o *CreditBankEmployment) GetEmployer() CreditBankEmployer {
	if o == nil {
		var ret CreditBankEmployer
		return ret
	}

	return o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmployment) GetEmployerOk() (*CreditBankEmployer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employer, true
}

// SetEmployer sets field value
func (o *CreditBankEmployment) SetEmployer(v CreditBankEmployer) {
	o.Employer = v
}

// GetLatestDepositDate returns the LatestDepositDate field value
func (o *CreditBankEmployment) GetLatestDepositDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LatestDepositDate
}

// GetLatestDepositDateOk returns a tuple with the LatestDepositDate field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmployment) GetLatestDepositDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LatestDepositDate, true
}

// SetLatestDepositDate sets field value
func (o *CreditBankEmployment) SetLatestDepositDate(v string) {
	o.LatestDepositDate = v
}

// GetEarliestDepositDate returns the EarliestDepositDate field value
func (o *CreditBankEmployment) GetEarliestDepositDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EarliestDepositDate
}

// GetEarliestDepositDateOk returns a tuple with the EarliestDepositDate field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmployment) GetEarliestDepositDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EarliestDepositDate, true
}

// SetEarliestDepositDate sets field value
func (o *CreditBankEmployment) SetEarliestDepositDate(v string) {
	o.EarliestDepositDate = v
}

func (o CreditBankEmployment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bank_employment_id"] = o.BankEmploymentId
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["employer"] = o.Employer
	}
	if true {
		toSerialize["latest_deposit_date"] = o.LatestDepositDate
	}
	if true {
		toSerialize["earliest_deposit_date"] = o.EarliestDepositDate
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankEmployment struct {
	value *CreditBankEmployment
	isSet bool
}

func (v NullableCreditBankEmployment) Get() *CreditBankEmployment {
	return v.value
}

func (v *NullableCreditBankEmployment) Set(val *CreditBankEmployment) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankEmployment) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankEmployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankEmployment(val *CreditBankEmployment) *NullableCreditBankEmployment {
	return &NullableCreditBankEmployment{value: val, isSet: true}
}

func (v NullableCreditBankEmployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankEmployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


