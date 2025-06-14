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

// TransferBalanceGetRequest Defines the request schema for `/transfer/balance/get`
type TransferBalanceGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Client ID of the end customer.
	OriginatorClientId NullableString `json:"originator_client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	Type *TransferBalanceType `json:"type,omitempty"`
}

// NewTransferBalanceGetRequest instantiates a new TransferBalanceGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferBalanceGetRequest() *TransferBalanceGetRequest {
	this := TransferBalanceGetRequest{}
	return &this
}

// NewTransferBalanceGetRequestWithDefaults instantiates a new TransferBalanceGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferBalanceGetRequestWithDefaults() *TransferBalanceGetRequest {
	this := TransferBalanceGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferBalanceGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferBalanceGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferBalanceGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferBalanceGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetOriginatorClientId returns the OriginatorClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferBalanceGetRequest) GetOriginatorClientId() string {
	if o == nil || o.OriginatorClientId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginatorClientId.Get()
}

// GetOriginatorClientIdOk returns a tuple with the OriginatorClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferBalanceGetRequest) GetOriginatorClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginatorClientId.Get(), o.OriginatorClientId.IsSet()
}

// HasOriginatorClientId returns a boolean if a field has been set.
func (o *TransferBalanceGetRequest) HasOriginatorClientId() bool {
	if o != nil && o.OriginatorClientId.IsSet() {
		return true
	}

	return false
}

// SetOriginatorClientId gets a reference to the given NullableString and assigns it to the OriginatorClientId field.
func (o *TransferBalanceGetRequest) SetOriginatorClientId(v string) {
	o.OriginatorClientId.Set(&v)
}
// SetOriginatorClientIdNil sets the value for OriginatorClientId to be an explicit nil
func (o *TransferBalanceGetRequest) SetOriginatorClientIdNil() {
	o.OriginatorClientId.Set(nil)
}

// UnsetOriginatorClientId ensures that no value is present for OriginatorClientId, not even an explicit nil
func (o *TransferBalanceGetRequest) UnsetOriginatorClientId() {
	o.OriginatorClientId.Unset()
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferBalanceGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferBalanceGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferBalanceGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferBalanceGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransferBalanceGetRequest) GetType() TransferBalanceType {
	if o == nil || o.Type == nil {
		var ret TransferBalanceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferBalanceGetRequest) GetTypeOk() (*TransferBalanceType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransferBalanceGetRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TransferBalanceType and assigns it to the Type field.
func (o *TransferBalanceGetRequest) SetType(v TransferBalanceType) {
	o.Type = &v
}

func (o TransferBalanceGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.OriginatorClientId.IsSet() {
		toSerialize["originator_client_id"] = o.OriginatorClientId.Get()
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTransferBalanceGetRequest struct {
	value *TransferBalanceGetRequest
	isSet bool
}

func (v NullableTransferBalanceGetRequest) Get() *TransferBalanceGetRequest {
	return v.value
}

func (v *NullableTransferBalanceGetRequest) Set(val *TransferBalanceGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferBalanceGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferBalanceGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferBalanceGetRequest(val *TransferBalanceGetRequest) *NullableTransferBalanceGetRequest {
	return &NullableTransferBalanceGetRequest{value: val, isSet: true}
}

func (v NullableTransferBalanceGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferBalanceGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


