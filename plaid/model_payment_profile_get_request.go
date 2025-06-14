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

// PaymentProfileGetRequest PaymentProfileGetRequest defines the request schema for `/payment_profile/get`
type PaymentProfileGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A payment profile token associated with the Payment Profile data that is being requested.
	PaymentProfileToken string `json:"payment_profile_token"`
}

// NewPaymentProfileGetRequest instantiates a new PaymentProfileGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentProfileGetRequest(paymentProfileToken string) *PaymentProfileGetRequest {
	this := PaymentProfileGetRequest{}
	this.PaymentProfileToken = paymentProfileToken
	return &this
}

// NewPaymentProfileGetRequestWithDefaults instantiates a new PaymentProfileGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentProfileGetRequestWithDefaults() *PaymentProfileGetRequest {
	this := PaymentProfileGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentProfileGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentProfileGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentProfileGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentProfileGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentProfileGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentProfileGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentProfileGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentProfileGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetPaymentProfileToken returns the PaymentProfileToken field value
func (o *PaymentProfileGetRequest) GetPaymentProfileToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentProfileToken
}

// GetPaymentProfileTokenOk returns a tuple with the PaymentProfileToken field value
// and a boolean to check if the value has been set.
func (o *PaymentProfileGetRequest) GetPaymentProfileTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentProfileToken, true
}

// SetPaymentProfileToken sets field value
func (o *PaymentProfileGetRequest) SetPaymentProfileToken(v string) {
	o.PaymentProfileToken = v
}

func (o PaymentProfileGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["payment_profile_token"] = o.PaymentProfileToken
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentProfileGetRequest struct {
	value *PaymentProfileGetRequest
	isSet bool
}

func (v NullablePaymentProfileGetRequest) Get() *PaymentProfileGetRequest {
	return v.value
}

func (v *NullablePaymentProfileGetRequest) Set(val *PaymentProfileGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentProfileGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentProfileGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentProfileGetRequest(val *PaymentProfileGetRequest) *NullablePaymentProfileGetRequest {
	return &NullablePaymentProfileGetRequest{value: val, isSet: true}
}

func (v NullablePaymentProfileGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentProfileGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


