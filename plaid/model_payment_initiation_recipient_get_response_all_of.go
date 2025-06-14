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

// PaymentInitiationRecipientGetResponseAllOf struct for PaymentInitiationRecipientGetResponseAllOf
type PaymentInitiationRecipientGetResponseAllOf struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId *string `json:"request_id,omitempty"`
}

// NewPaymentInitiationRecipientGetResponseAllOf instantiates a new PaymentInitiationRecipientGetResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationRecipientGetResponseAllOf() *PaymentInitiationRecipientGetResponseAllOf {
	this := PaymentInitiationRecipientGetResponseAllOf{}
	return &this
}

// NewPaymentInitiationRecipientGetResponseAllOfWithDefaults instantiates a new PaymentInitiationRecipientGetResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationRecipientGetResponseAllOfWithDefaults() *PaymentInitiationRecipientGetResponseAllOf {
	this := PaymentInitiationRecipientGetResponseAllOf{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *PaymentInitiationRecipientGetResponseAllOf) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientGetResponseAllOf) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientGetResponseAllOf) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *PaymentInitiationRecipientGetResponseAllOf) SetRequestId(v string) {
	o.RequestId = &v
}

func (o PaymentInitiationRecipientGetResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationRecipientGetResponseAllOf struct {
	value *PaymentInitiationRecipientGetResponseAllOf
	isSet bool
}

func (v NullablePaymentInitiationRecipientGetResponseAllOf) Get() *PaymentInitiationRecipientGetResponseAllOf {
	return v.value
}

func (v *NullablePaymentInitiationRecipientGetResponseAllOf) Set(val *PaymentInitiationRecipientGetResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRecipientGetResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRecipientGetResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRecipientGetResponseAllOf(val *PaymentInitiationRecipientGetResponseAllOf) *NullablePaymentInitiationRecipientGetResponseAllOf {
	return &NullablePaymentInitiationRecipientGetResponseAllOf{value: val, isSet: true}
}

func (v NullablePaymentInitiationRecipientGetResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRecipientGetResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


