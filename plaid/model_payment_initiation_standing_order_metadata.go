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

// PaymentInitiationStandingOrderMetadata Metadata specifically related to valid Payment Initiation standing order configurations for the institution.
type PaymentInitiationStandingOrderMetadata struct {
	// Indicates whether the institution supports closed-ended standing orders by providing an end date.
	SupportsStandingOrderEndDate bool `json:"supports_standing_order_end_date"`
	// This is only applicable to `MONTHLY` standing orders. Indicates whether the institution supports negative integers (-1 to -5) for setting up a `MONTHLY` standing order relative to the end of the month.
	SupportsStandingOrderNegativeExecutionDays bool `json:"supports_standing_order_negative_execution_days"`
	// A list of the valid standing order intervals supported by the institution.
	ValidStandingOrderIntervals []PaymentScheduleInterval `json:"valid_standing_order_intervals"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationStandingOrderMetadata PaymentInitiationStandingOrderMetadata

// NewPaymentInitiationStandingOrderMetadata instantiates a new PaymentInitiationStandingOrderMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationStandingOrderMetadata(supportsStandingOrderEndDate bool, supportsStandingOrderNegativeExecutionDays bool, validStandingOrderIntervals []PaymentScheduleInterval) *PaymentInitiationStandingOrderMetadata {
	this := PaymentInitiationStandingOrderMetadata{}
	this.SupportsStandingOrderEndDate = supportsStandingOrderEndDate
	this.SupportsStandingOrderNegativeExecutionDays = supportsStandingOrderNegativeExecutionDays
	this.ValidStandingOrderIntervals = validStandingOrderIntervals
	return &this
}

// NewPaymentInitiationStandingOrderMetadataWithDefaults instantiates a new PaymentInitiationStandingOrderMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationStandingOrderMetadataWithDefaults() *PaymentInitiationStandingOrderMetadata {
	this := PaymentInitiationStandingOrderMetadata{}
	return &this
}

// GetSupportsStandingOrderEndDate returns the SupportsStandingOrderEndDate field value
func (o *PaymentInitiationStandingOrderMetadata) GetSupportsStandingOrderEndDate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportsStandingOrderEndDate
}

// GetSupportsStandingOrderEndDateOk returns a tuple with the SupportsStandingOrderEndDate field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationStandingOrderMetadata) GetSupportsStandingOrderEndDateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupportsStandingOrderEndDate, true
}

// SetSupportsStandingOrderEndDate sets field value
func (o *PaymentInitiationStandingOrderMetadata) SetSupportsStandingOrderEndDate(v bool) {
	o.SupportsStandingOrderEndDate = v
}

// GetSupportsStandingOrderNegativeExecutionDays returns the SupportsStandingOrderNegativeExecutionDays field value
func (o *PaymentInitiationStandingOrderMetadata) GetSupportsStandingOrderNegativeExecutionDays() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportsStandingOrderNegativeExecutionDays
}

// GetSupportsStandingOrderNegativeExecutionDaysOk returns a tuple with the SupportsStandingOrderNegativeExecutionDays field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationStandingOrderMetadata) GetSupportsStandingOrderNegativeExecutionDaysOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupportsStandingOrderNegativeExecutionDays, true
}

// SetSupportsStandingOrderNegativeExecutionDays sets field value
func (o *PaymentInitiationStandingOrderMetadata) SetSupportsStandingOrderNegativeExecutionDays(v bool) {
	o.SupportsStandingOrderNegativeExecutionDays = v
}

// GetValidStandingOrderIntervals returns the ValidStandingOrderIntervals field value
func (o *PaymentInitiationStandingOrderMetadata) GetValidStandingOrderIntervals() []PaymentScheduleInterval {
	if o == nil {
		var ret []PaymentScheduleInterval
		return ret
	}

	return o.ValidStandingOrderIntervals
}

// GetValidStandingOrderIntervalsOk returns a tuple with the ValidStandingOrderIntervals field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationStandingOrderMetadata) GetValidStandingOrderIntervalsOk() (*[]PaymentScheduleInterval, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValidStandingOrderIntervals, true
}

// SetValidStandingOrderIntervals sets field value
func (o *PaymentInitiationStandingOrderMetadata) SetValidStandingOrderIntervals(v []PaymentScheduleInterval) {
	o.ValidStandingOrderIntervals = v
}

func (o PaymentInitiationStandingOrderMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supports_standing_order_end_date"] = o.SupportsStandingOrderEndDate
	}
	if true {
		toSerialize["supports_standing_order_negative_execution_days"] = o.SupportsStandingOrderNegativeExecutionDays
	}
	if true {
		toSerialize["valid_standing_order_intervals"] = o.ValidStandingOrderIntervals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationStandingOrderMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationStandingOrderMetadata := _PaymentInitiationStandingOrderMetadata{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationStandingOrderMetadata); err == nil {
		*o = PaymentInitiationStandingOrderMetadata(varPaymentInitiationStandingOrderMetadata)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "supports_standing_order_end_date")
		delete(additionalProperties, "supports_standing_order_negative_execution_days")
		delete(additionalProperties, "valid_standing_order_intervals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationStandingOrderMetadata struct {
	value *PaymentInitiationStandingOrderMetadata
	isSet bool
}

func (v NullablePaymentInitiationStandingOrderMetadata) Get() *PaymentInitiationStandingOrderMetadata {
	return v.value
}

func (v *NullablePaymentInitiationStandingOrderMetadata) Set(val *PaymentInitiationStandingOrderMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationStandingOrderMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationStandingOrderMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationStandingOrderMetadata(val *PaymentInitiationStandingOrderMetadata) *NullablePaymentInitiationStandingOrderMetadata {
	return &NullablePaymentInitiationStandingOrderMetadata{value: val, isSet: true}
}

func (v NullablePaymentInitiationStandingOrderMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationStandingOrderMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


