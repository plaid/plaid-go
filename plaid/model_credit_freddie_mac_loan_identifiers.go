/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.503.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CreditFreddieMacLoanIdentifiers Collection of current and previous identifiers for this loan.
type CreditFreddieMacLoanIdentifiers struct {
	LOAN_IDENTIFIER []LoanIdentifier `json:"LOAN_IDENTIFIER"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacLoanIdentifiers CreditFreddieMacLoanIdentifiers

// NewCreditFreddieMacLoanIdentifiers instantiates a new CreditFreddieMacLoanIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacLoanIdentifiers(lOANIDENTIFIER []LoanIdentifier) *CreditFreddieMacLoanIdentifiers {
	this := CreditFreddieMacLoanIdentifiers{}
	this.LOAN_IDENTIFIER = lOANIDENTIFIER
	return &this
}

// NewCreditFreddieMacLoanIdentifiersWithDefaults instantiates a new CreditFreddieMacLoanIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacLoanIdentifiersWithDefaults() *CreditFreddieMacLoanIdentifiers {
	this := CreditFreddieMacLoanIdentifiers{}
	return &this
}

// GetLOAN_IDENTIFIER returns the LOAN_IDENTIFIER field value
func (o *CreditFreddieMacLoanIdentifiers) GetLOAN_IDENTIFIER() []LoanIdentifier {
	if o == nil {
		var ret []LoanIdentifier
		return ret
	}

	return o.LOAN_IDENTIFIER
}

// GetLOAN_IDENTIFIEROk returns a tuple with the LOAN_IDENTIFIER field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacLoanIdentifiers) GetLOAN_IDENTIFIEROk() (*[]LoanIdentifier, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LOAN_IDENTIFIER, true
}

// SetLOAN_IDENTIFIER sets field value
func (o *CreditFreddieMacLoanIdentifiers) SetLOAN_IDENTIFIER(v []LoanIdentifier) {
	o.LOAN_IDENTIFIER = v
}

func (o CreditFreddieMacLoanIdentifiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LOAN_IDENTIFIER"] = o.LOAN_IDENTIFIER
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacLoanIdentifiers) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacLoanIdentifiers := _CreditFreddieMacLoanIdentifiers{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacLoanIdentifiers); err == nil {
		*o = CreditFreddieMacLoanIdentifiers(varCreditFreddieMacLoanIdentifiers)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LOAN_IDENTIFIER")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacLoanIdentifiers struct {
	value *CreditFreddieMacLoanIdentifiers
	isSet bool
}

func (v NullableCreditFreddieMacLoanIdentifiers) Get() *CreditFreddieMacLoanIdentifiers {
	return v.value
}

func (v *NullableCreditFreddieMacLoanIdentifiers) Set(val *CreditFreddieMacLoanIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacLoanIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacLoanIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacLoanIdentifiers(val *CreditFreddieMacLoanIdentifiers) *NullableCreditFreddieMacLoanIdentifiers {
	return &NullableCreditFreddieMacLoanIdentifiers{value: val, isSet: true}
}

func (v NullableCreditFreddieMacLoanIdentifiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacLoanIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

