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

// TransferOriginatorFundingAccountUpdateRequest Defines the request schema for `/transfer/originator/funding_account/update`
type TransferOriginatorFundingAccountUpdateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The Plaid client ID of the transfer originator.
	OriginatorClientId string `json:"originator_client_id"`
	FundingAccount TransferFundingAccount `json:"funding_account"`
}

// NewTransferOriginatorFundingAccountUpdateRequest instantiates a new TransferOriginatorFundingAccountUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferOriginatorFundingAccountUpdateRequest(originatorClientId string, fundingAccount TransferFundingAccount) *TransferOriginatorFundingAccountUpdateRequest {
	this := TransferOriginatorFundingAccountUpdateRequest{}
	this.OriginatorClientId = originatorClientId
	this.FundingAccount = fundingAccount
	return &this
}

// NewTransferOriginatorFundingAccountUpdateRequestWithDefaults instantiates a new TransferOriginatorFundingAccountUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferOriginatorFundingAccountUpdateRequestWithDefaults() *TransferOriginatorFundingAccountUpdateRequest {
	this := TransferOriginatorFundingAccountUpdateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferOriginatorFundingAccountUpdateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOriginatorFundingAccountUpdateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferOriginatorFundingAccountUpdateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferOriginatorFundingAccountUpdateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferOriginatorFundingAccountUpdateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferOriginatorFundingAccountUpdateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferOriginatorFundingAccountUpdateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferOriginatorFundingAccountUpdateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetOriginatorClientId returns the OriginatorClientId field value
func (o *TransferOriginatorFundingAccountUpdateRequest) GetOriginatorClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatorClientId
}

// GetOriginatorClientIdOk returns a tuple with the OriginatorClientId field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorFundingAccountUpdateRequest) GetOriginatorClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginatorClientId, true
}

// SetOriginatorClientId sets field value
func (o *TransferOriginatorFundingAccountUpdateRequest) SetOriginatorClientId(v string) {
	o.OriginatorClientId = v
}

// GetFundingAccount returns the FundingAccount field value
func (o *TransferOriginatorFundingAccountUpdateRequest) GetFundingAccount() TransferFundingAccount {
	if o == nil {
		var ret TransferFundingAccount
		return ret
	}

	return o.FundingAccount
}

// GetFundingAccountOk returns a tuple with the FundingAccount field value
// and a boolean to check if the value has been set.
func (o *TransferOriginatorFundingAccountUpdateRequest) GetFundingAccountOk() (*TransferFundingAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FundingAccount, true
}

// SetFundingAccount sets field value
func (o *TransferOriginatorFundingAccountUpdateRequest) SetFundingAccount(v TransferFundingAccount) {
	o.FundingAccount = v
}

func (o TransferOriginatorFundingAccountUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["originator_client_id"] = o.OriginatorClientId
	}
	if true {
		toSerialize["funding_account"] = o.FundingAccount
	}
	return json.Marshal(toSerialize)
}

type NullableTransferOriginatorFundingAccountUpdateRequest struct {
	value *TransferOriginatorFundingAccountUpdateRequest
	isSet bool
}

func (v NullableTransferOriginatorFundingAccountUpdateRequest) Get() *TransferOriginatorFundingAccountUpdateRequest {
	return v.value
}

func (v *NullableTransferOriginatorFundingAccountUpdateRequest) Set(val *TransferOriginatorFundingAccountUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferOriginatorFundingAccountUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferOriginatorFundingAccountUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferOriginatorFundingAccountUpdateRequest(val *TransferOriginatorFundingAccountUpdateRequest) *NullableTransferOriginatorFundingAccountUpdateRequest {
	return &NullableTransferOriginatorFundingAccountUpdateRequest{value: val, isSet: true}
}

func (v NullableTransferOriginatorFundingAccountUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferOriginatorFundingAccountUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


