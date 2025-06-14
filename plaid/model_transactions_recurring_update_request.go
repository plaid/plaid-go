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

// TransactionsRecurringUpdateRequest TransactionsRecurringUpdateRequest defined the request schema for `/transactions/recurring/streams/update` endpoint.
type TransactionsRecurringUpdateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A list of all the operations to be performed. This will either all succeed or all fail.
	Inputs []TransactionsRecurringUpdateInput `json:"inputs"`
}

// NewTransactionsRecurringUpdateRequest instantiates a new TransactionsRecurringUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRecurringUpdateRequest(accessToken string, inputs []TransactionsRecurringUpdateInput) *TransactionsRecurringUpdateRequest {
	this := TransactionsRecurringUpdateRequest{}
	this.AccessToken = accessToken
	this.Inputs = inputs
	return &this
}

// NewTransactionsRecurringUpdateRequestWithDefaults instantiates a new TransactionsRecurringUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRecurringUpdateRequestWithDefaults() *TransactionsRecurringUpdateRequest {
	this := TransactionsRecurringUpdateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransactionsRecurringUpdateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringUpdateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransactionsRecurringUpdateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransactionsRecurringUpdateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetAccessToken returns the AccessToken field value
func (o *TransactionsRecurringUpdateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringUpdateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TransactionsRecurringUpdateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransactionsRecurringUpdateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringUpdateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransactionsRecurringUpdateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransactionsRecurringUpdateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetInputs returns the Inputs field value
func (o *TransactionsRecurringUpdateRequest) GetInputs() []TransactionsRecurringUpdateInput {
	if o == nil {
		var ret []TransactionsRecurringUpdateInput
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *TransactionsRecurringUpdateRequest) GetInputsOk() (*[]TransactionsRecurringUpdateInput, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value
func (o *TransactionsRecurringUpdateRequest) SetInputs(v []TransactionsRecurringUpdateInput) {
	o.Inputs = v
}

func (o TransactionsRecurringUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsRecurringUpdateRequest struct {
	value *TransactionsRecurringUpdateRequest
	isSet bool
}

func (v NullableTransactionsRecurringUpdateRequest) Get() *TransactionsRecurringUpdateRequest {
	return v.value
}

func (v *NullableTransactionsRecurringUpdateRequest) Set(val *TransactionsRecurringUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRecurringUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRecurringUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRecurringUpdateRequest(val *TransactionsRecurringUpdateRequest) *NullableTransactionsRecurringUpdateRequest {
	return &NullableTransactionsRecurringUpdateRequest{value: val, isSet: true}
}

func (v NullableTransactionsRecurringUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRecurringUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


