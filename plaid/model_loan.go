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

// Loan Information specific to a mortgage loan agreement between one or more borrowers and a mortgage lender.
type Loan struct {
	LOAN_IDENTIFIERS LoanIdentifiers `json:"LOAN_IDENTIFIERS"`
	AdditionalProperties map[string]interface{}
}

type _Loan Loan

// NewLoan instantiates a new Loan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoan(lOANIDENTIFIERS LoanIdentifiers) *Loan {
	this := Loan{}
	this.LOAN_IDENTIFIERS = lOANIDENTIFIERS
	return &this
}

// NewLoanWithDefaults instantiates a new Loan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoanWithDefaults() *Loan {
	this := Loan{}
	return &this
}

// GetLOAN_IDENTIFIERS returns the LOAN_IDENTIFIERS field value
func (o *Loan) GetLOAN_IDENTIFIERS() LoanIdentifiers {
	if o == nil {
		var ret LoanIdentifiers
		return ret
	}

	return o.LOAN_IDENTIFIERS
}

// GetLOAN_IDENTIFIERSOk returns a tuple with the LOAN_IDENTIFIERS field value
// and a boolean to check if the value has been set.
func (o *Loan) GetLOAN_IDENTIFIERSOk() (*LoanIdentifiers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LOAN_IDENTIFIERS, true
}

// SetLOAN_IDENTIFIERS sets field value
func (o *Loan) SetLOAN_IDENTIFIERS(v LoanIdentifiers) {
	o.LOAN_IDENTIFIERS = v
}

func (o Loan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LOAN_IDENTIFIERS"] = o.LOAN_IDENTIFIERS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Loan) UnmarshalJSON(bytes []byte) (err error) {
	varLoan := _Loan{}

	if err = json.Unmarshal(bytes, &varLoan); err == nil {
		*o = Loan(varLoan)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LOAN_IDENTIFIERS")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoan struct {
	value *Loan
	isSet bool
}

func (v NullableLoan) Get() *Loan {
	return v.value
}

func (v *NullableLoan) Set(val *Loan) {
	v.value = val
	v.isSet = true
}

func (v NullableLoan) IsSet() bool {
	return v.isSet
}

func (v *NullableLoan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoan(val *Loan) *NullableLoan {
	return &NullableLoan{value: val, isSet: true}
}

func (v NullableLoan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


