/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.370.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferDebitUsageConfiguration Specifies the originator's expected usage of debits. For all dollar amounts, use a decimal string with two digits of precision e.g. \"10.00\". This field is required if the originator is expected to process debit transfers.
type TransferDebitUsageConfiguration struct {
	ExpectedFrequency OriginatorExpectedTransferFrequency `json:"expected_frequency"`
	// The originator’s expected highest amount for a single debit transfer.
	ExpectedHighestAmount string `json:"expected_highest_amount"`
	// The originator’s expected average amount per debit.
	ExpectedAverageAmount string `json:"expected_average_amount"`
	// The originator’s monthly expected ACH debit processing amount for the next 6-12 months.
	ExpectedMonthlyAmount string `json:"expected_monthly_amount"`
	// Specifies the expected use cases for the originator’s debit transfers. This should be a list that contains one or more of the following codes:  `\"ccd\"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  `\"ppd\"` - Prearranged Payment or Deposit - the transfer is part of a pre-existing relationship with a consumer, eg. bill payment  `\"tel\"` - Telephone-Initiated Entry  `\"web\"` - Internet-Initiated Entry - debits from a consumer’s account where their authorization is obtained over the Internet
	SecCodes []ACHClass `json:"sec_codes"`
}

// NewTransferDebitUsageConfiguration instantiates a new TransferDebitUsageConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferDebitUsageConfiguration(expectedFrequency OriginatorExpectedTransferFrequency, expectedHighestAmount string, expectedAverageAmount string, expectedMonthlyAmount string, secCodes []ACHClass) *TransferDebitUsageConfiguration {
	this := TransferDebitUsageConfiguration{}
	this.ExpectedFrequency = expectedFrequency
	this.ExpectedHighestAmount = expectedHighestAmount
	this.ExpectedAverageAmount = expectedAverageAmount
	this.ExpectedMonthlyAmount = expectedMonthlyAmount
	this.SecCodes = secCodes
	return &this
}

// NewTransferDebitUsageConfigurationWithDefaults instantiates a new TransferDebitUsageConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferDebitUsageConfigurationWithDefaults() *TransferDebitUsageConfiguration {
	this := TransferDebitUsageConfiguration{}
	return &this
}

// GetExpectedFrequency returns the ExpectedFrequency field value
func (o *TransferDebitUsageConfiguration) GetExpectedFrequency() OriginatorExpectedTransferFrequency {
	if o == nil {
		var ret OriginatorExpectedTransferFrequency
		return ret
	}

	return o.ExpectedFrequency
}

// GetExpectedFrequencyOk returns a tuple with the ExpectedFrequency field value
// and a boolean to check if the value has been set.
func (o *TransferDebitUsageConfiguration) GetExpectedFrequencyOk() (*OriginatorExpectedTransferFrequency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpectedFrequency, true
}

// SetExpectedFrequency sets field value
func (o *TransferDebitUsageConfiguration) SetExpectedFrequency(v OriginatorExpectedTransferFrequency) {
	o.ExpectedFrequency = v
}

// GetExpectedHighestAmount returns the ExpectedHighestAmount field value
func (o *TransferDebitUsageConfiguration) GetExpectedHighestAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpectedHighestAmount
}

// GetExpectedHighestAmountOk returns a tuple with the ExpectedHighestAmount field value
// and a boolean to check if the value has been set.
func (o *TransferDebitUsageConfiguration) GetExpectedHighestAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpectedHighestAmount, true
}

// SetExpectedHighestAmount sets field value
func (o *TransferDebitUsageConfiguration) SetExpectedHighestAmount(v string) {
	o.ExpectedHighestAmount = v
}

// GetExpectedAverageAmount returns the ExpectedAverageAmount field value
func (o *TransferDebitUsageConfiguration) GetExpectedAverageAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpectedAverageAmount
}

// GetExpectedAverageAmountOk returns a tuple with the ExpectedAverageAmount field value
// and a boolean to check if the value has been set.
func (o *TransferDebitUsageConfiguration) GetExpectedAverageAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpectedAverageAmount, true
}

// SetExpectedAverageAmount sets field value
func (o *TransferDebitUsageConfiguration) SetExpectedAverageAmount(v string) {
	o.ExpectedAverageAmount = v
}

// GetExpectedMonthlyAmount returns the ExpectedMonthlyAmount field value
func (o *TransferDebitUsageConfiguration) GetExpectedMonthlyAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpectedMonthlyAmount
}

// GetExpectedMonthlyAmountOk returns a tuple with the ExpectedMonthlyAmount field value
// and a boolean to check if the value has been set.
func (o *TransferDebitUsageConfiguration) GetExpectedMonthlyAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpectedMonthlyAmount, true
}

// SetExpectedMonthlyAmount sets field value
func (o *TransferDebitUsageConfiguration) SetExpectedMonthlyAmount(v string) {
	o.ExpectedMonthlyAmount = v
}

// GetSecCodes returns the SecCodes field value
func (o *TransferDebitUsageConfiguration) GetSecCodes() []ACHClass {
	if o == nil {
		var ret []ACHClass
		return ret
	}

	return o.SecCodes
}

// GetSecCodesOk returns a tuple with the SecCodes field value
// and a boolean to check if the value has been set.
func (o *TransferDebitUsageConfiguration) GetSecCodesOk() (*[]ACHClass, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecCodes, true
}

// SetSecCodes sets field value
func (o *TransferDebitUsageConfiguration) SetSecCodes(v []ACHClass) {
	o.SecCodes = v
}

func (o TransferDebitUsageConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expected_frequency"] = o.ExpectedFrequency
	}
	if true {
		toSerialize["expected_highest_amount"] = o.ExpectedHighestAmount
	}
	if true {
		toSerialize["expected_average_amount"] = o.ExpectedAverageAmount
	}
	if true {
		toSerialize["expected_monthly_amount"] = o.ExpectedMonthlyAmount
	}
	if true {
		toSerialize["sec_codes"] = o.SecCodes
	}
	return json.Marshal(toSerialize)
}

type NullableTransferDebitUsageConfiguration struct {
	value *TransferDebitUsageConfiguration
	isSet bool
}

func (v NullableTransferDebitUsageConfiguration) Get() *TransferDebitUsageConfiguration {
	return v.value
}

func (v *NullableTransferDebitUsageConfiguration) Set(val *TransferDebitUsageConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDebitUsageConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDebitUsageConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDebitUsageConfiguration(val *TransferDebitUsageConfiguration) *NullableTransferDebitUsageConfiguration {
	return &NullableTransferDebitUsageConfiguration{value: val, isSet: true}
}

func (v NullableTransferDebitUsageConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDebitUsageConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

