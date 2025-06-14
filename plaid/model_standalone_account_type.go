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

// StandaloneAccountType The schema below describes the various `types` and corresponding `subtypes` that Plaid recognizes and reports for financial institution accounts. For a mapping of supported types and subtypes to Plaid products, see the [Account type / product support matrix](https://plaid.com/docs/api/accounts/#account-type--product-support-matrix).
type StandaloneAccountType struct {
	// An account type holding cash, in which funds are deposited.
	Depository string `json:"depository"`
	// A credit card type account.
	Credit string `json:"credit"`
	// A loan type account.
	Loan string `json:"loan"`
	// An investment account. In API versions 2018-05-22 and earlier, this type is called `brokerage`.
	Investment string `json:"investment"`
	// Other or unknown account type.
	Other string `json:"other"`
	AdditionalProperties map[string]interface{}
}

type _StandaloneAccountType StandaloneAccountType

// NewStandaloneAccountType instantiates a new StandaloneAccountType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandaloneAccountType(depository string, credit string, loan string, investment string, other string) *StandaloneAccountType {
	this := StandaloneAccountType{}
	this.Depository = depository
	this.Credit = credit
	this.Loan = loan
	this.Investment = investment
	this.Other = other
	return &this
}

// NewStandaloneAccountTypeWithDefaults instantiates a new StandaloneAccountType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandaloneAccountTypeWithDefaults() *StandaloneAccountType {
	this := StandaloneAccountType{}
	return &this
}

// GetDepository returns the Depository field value
func (o *StandaloneAccountType) GetDepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Depository
}

// GetDepositoryOk returns a tuple with the Depository field value
// and a boolean to check if the value has been set.
func (o *StandaloneAccountType) GetDepositoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Depository, true
}

// SetDepository sets field value
func (o *StandaloneAccountType) SetDepository(v string) {
	o.Depository = v
}

// GetCredit returns the Credit field value
func (o *StandaloneAccountType) GetCredit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credit
}

// GetCreditOk returns a tuple with the Credit field value
// and a boolean to check if the value has been set.
func (o *StandaloneAccountType) GetCreditOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credit, true
}

// SetCredit sets field value
func (o *StandaloneAccountType) SetCredit(v string) {
	o.Credit = v
}

// GetLoan returns the Loan field value
func (o *StandaloneAccountType) GetLoan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Loan
}

// GetLoanOk returns a tuple with the Loan field value
// and a boolean to check if the value has been set.
func (o *StandaloneAccountType) GetLoanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Loan, true
}

// SetLoan sets field value
func (o *StandaloneAccountType) SetLoan(v string) {
	o.Loan = v
}

// GetInvestment returns the Investment field value
func (o *StandaloneAccountType) GetInvestment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Investment
}

// GetInvestmentOk returns a tuple with the Investment field value
// and a boolean to check if the value has been set.
func (o *StandaloneAccountType) GetInvestmentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Investment, true
}

// SetInvestment sets field value
func (o *StandaloneAccountType) SetInvestment(v string) {
	o.Investment = v
}

// GetOther returns the Other field value
func (o *StandaloneAccountType) GetOther() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Other
}

// GetOtherOk returns a tuple with the Other field value
// and a boolean to check if the value has been set.
func (o *StandaloneAccountType) GetOtherOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Other, true
}

// SetOther sets field value
func (o *StandaloneAccountType) SetOther(v string) {
	o.Other = v
}

func (o StandaloneAccountType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["depository"] = o.Depository
	}
	if true {
		toSerialize["credit"] = o.Credit
	}
	if true {
		toSerialize["loan"] = o.Loan
	}
	if true {
		toSerialize["investment"] = o.Investment
	}
	if true {
		toSerialize["other"] = o.Other
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StandaloneAccountType) UnmarshalJSON(bytes []byte) (err error) {
	varStandaloneAccountType := _StandaloneAccountType{}

	if err = json.Unmarshal(bytes, &varStandaloneAccountType); err == nil {
		*o = StandaloneAccountType(varStandaloneAccountType)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "depository")
		delete(additionalProperties, "credit")
		delete(additionalProperties, "loan")
		delete(additionalProperties, "investment")
		delete(additionalProperties, "other")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStandaloneAccountType struct {
	value *StandaloneAccountType
	isSet bool
}

func (v NullableStandaloneAccountType) Get() *StandaloneAccountType {
	return v.value
}

func (v *NullableStandaloneAccountType) Set(val *StandaloneAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableStandaloneAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableStandaloneAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandaloneAccountType(val *StandaloneAccountType) *NullableStandaloneAccountType {
	return &NullableStandaloneAccountType{value: val, isSet: true}
}

func (v NullableStandaloneAccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandaloneAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


