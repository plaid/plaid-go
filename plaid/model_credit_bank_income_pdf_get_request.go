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

// CreditBankIncomePDFGetRequest CreditBankIncomePDFGetRequest defines the request schema for `/credit/bank_income/pdf/get`
type CreditBankIncomePDFGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken string `json:"user_token"`
}

// NewCreditBankIncomePDFGetRequest instantiates a new CreditBankIncomePDFGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomePDFGetRequest(userToken string) *CreditBankIncomePDFGetRequest {
	this := CreditBankIncomePDFGetRequest{}
	this.UserToken = userToken
	return &this
}

// NewCreditBankIncomePDFGetRequestWithDefaults instantiates a new CreditBankIncomePDFGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomePDFGetRequestWithDefaults() *CreditBankIncomePDFGetRequest {
	this := CreditBankIncomePDFGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditBankIncomePDFGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomePDFGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditBankIncomePDFGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditBankIncomePDFGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditBankIncomePDFGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomePDFGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditBankIncomePDFGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditBankIncomePDFGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value
func (o *CreditBankIncomePDFGetRequest) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *CreditBankIncomePDFGetRequest) GetUserTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *CreditBankIncomePDFGetRequest) SetUserToken(v string) {
	o.UserToken = v
}

func (o CreditBankIncomePDFGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["user_token"] = o.UserToken
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankIncomePDFGetRequest struct {
	value *CreditBankIncomePDFGetRequest
	isSet bool
}

func (v NullableCreditBankIncomePDFGetRequest) Get() *CreditBankIncomePDFGetRequest {
	return v.value
}

func (v *NullableCreditBankIncomePDFGetRequest) Set(val *CreditBankIncomePDFGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomePDFGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomePDFGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomePDFGetRequest(val *CreditBankIncomePDFGetRequest) *NullableCreditBankIncomePDFGetRequest {
	return &NullableCreditBankIncomePDFGetRequest{value: val, isSet: true}
}

func (v NullableCreditBankIncomePDFGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomePDFGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


