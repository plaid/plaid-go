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

// WalletTransactionGetRequest WalletTransactionGetRequest defines the request schema for `/wallet/transaction/get`
type WalletTransactionGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the transaction to fetch
	TransactionId string `json:"transaction_id"`
}

// NewWalletTransactionGetRequest instantiates a new WalletTransactionGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionGetRequest(transactionId string) *WalletTransactionGetRequest {
	this := WalletTransactionGetRequest{}
	this.TransactionId = transactionId
	return &this
}

// NewWalletTransactionGetRequestWithDefaults instantiates a new WalletTransactionGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionGetRequestWithDefaults() *WalletTransactionGetRequest {
	this := WalletTransactionGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WalletTransactionGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WalletTransactionGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WalletTransactionGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WalletTransactionGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WalletTransactionGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WalletTransactionGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetTransactionId returns the TransactionId field value
func (o *WalletTransactionGetRequest) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetRequest) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *WalletTransactionGetRequest) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o WalletTransactionGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableWalletTransactionGetRequest struct {
	value *WalletTransactionGetRequest
	isSet bool
}

func (v NullableWalletTransactionGetRequest) Get() *WalletTransactionGetRequest {
	return v.value
}

func (v *NullableWalletTransactionGetRequest) Set(val *WalletTransactionGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionGetRequest(val *WalletTransactionGetRequest) *NullableWalletTransactionGetRequest {
	return &NullableWalletTransactionGetRequest{value: val, isSet: true}
}

func (v NullableWalletTransactionGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


