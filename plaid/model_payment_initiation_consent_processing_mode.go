/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.586.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// PaymentInitiationConsentProcessingMode Decides the mode under which the payment processing should be performed, using `IMMEDIATE` as default.  `IMMEDIATE`: Will immediately execute the payment, waiting for a response from the ASPSP before returning the result of the payment initiation. This is ideal for user-present flows.   `ASYNC`: Will accept a payment execution request and schedule it for processing, immediately returning the new `payment_id`. Listen for webhooks or use the [`/payment_initiation/payment/get`](https://plaid.com/docs/api/products/payment-initiation/#payment_initiationpaymentget) endpoint to obtain updates on the payment status. This is ideal for non user-present flows.
type PaymentInitiationConsentProcessingMode string

var _ = fmt.Printf

// List of PaymentInitiationConsentProcessingMode
const (
	PAYMENTINITIATIONCONSENTPROCESSINGMODE_ASYNC PaymentInitiationConsentProcessingMode = "ASYNC"
	PAYMENTINITIATIONCONSENTPROCESSINGMODE_IMMEDIATE PaymentInitiationConsentProcessingMode = "IMMEDIATE"
)

var allowedPaymentInitiationConsentProcessingModeEnumValues = []PaymentInitiationConsentProcessingMode{
	"ASYNC",
	"IMMEDIATE",
}

func (v *PaymentInitiationConsentProcessingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PaymentInitiationConsentProcessingMode(value)


	*v = enumTypeValue
	return nil
}

// NewPaymentInitiationConsentProcessingModeFromValue returns a pointer to a valid PaymentInitiationConsentProcessingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentInitiationConsentProcessingModeFromValue(v string) (*PaymentInitiationConsentProcessingMode, error) {
	ev := PaymentInitiationConsentProcessingMode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentInitiationConsentProcessingMode) IsValid() bool {
	for _, existing := range allowedPaymentInitiationConsentProcessingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentInitiationConsentProcessingMode value
func (v PaymentInitiationConsentProcessingMode) Ptr() *PaymentInitiationConsentProcessingMode {
	return &v
}

type NullablePaymentInitiationConsentProcessingMode struct {
	value *PaymentInitiationConsentProcessingMode
	isSet bool
}

func (v NullablePaymentInitiationConsentProcessingMode) Get() *PaymentInitiationConsentProcessingMode {
	return v.value
}

func (v *NullablePaymentInitiationConsentProcessingMode) Set(val *PaymentInitiationConsentProcessingMode) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsentProcessingMode) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsentProcessingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsentProcessingMode(val *PaymentInitiationConsentProcessingMode) *NullablePaymentInitiationConsentProcessingMode {
	return &NullablePaymentInitiationConsentProcessingMode{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsentProcessingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsentProcessingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
