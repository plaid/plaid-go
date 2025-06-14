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

// ProcessorTransactionsRecurringGetRequest ProcessorTransactionsRecurringGetRequest defines the request schema for `/processor/transactions/recurring/get`
type ProcessorTransactionsRecurringGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// The processor token obtained from the Plaid integration partner. Processor tokens are in the format: `processor-<environment>-<identifier>`
	ProcessorToken string `json:"processor_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	Options *TransactionsRecurringGetRequestOptions `json:"options,omitempty"`
}

// NewProcessorTransactionsRecurringGetRequest instantiates a new ProcessorTransactionsRecurringGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorTransactionsRecurringGetRequest(processorToken string) *ProcessorTransactionsRecurringGetRequest {
	this := ProcessorTransactionsRecurringGetRequest{}
	this.ProcessorToken = processorToken
	return &this
}

// NewProcessorTransactionsRecurringGetRequestWithDefaults instantiates a new ProcessorTransactionsRecurringGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorTransactionsRecurringGetRequestWithDefaults() *ProcessorTransactionsRecurringGetRequest {
	this := ProcessorTransactionsRecurringGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorTransactionsRecurringGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsRecurringGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorTransactionsRecurringGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorTransactionsRecurringGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetProcessorToken returns the ProcessorToken field value
func (o *ProcessorTransactionsRecurringGetRequest) GetProcessorToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessorToken
}

// GetProcessorTokenOk returns a tuple with the ProcessorToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsRecurringGetRequest) GetProcessorTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProcessorToken, true
}

// SetProcessorToken sets field value
func (o *ProcessorTransactionsRecurringGetRequest) SetProcessorToken(v string) {
	o.ProcessorToken = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorTransactionsRecurringGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsRecurringGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorTransactionsRecurringGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorTransactionsRecurringGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProcessorTransactionsRecurringGetRequest) GetOptions() TransactionsRecurringGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret TransactionsRecurringGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsRecurringGetRequest) GetOptionsOk() (*TransactionsRecurringGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProcessorTransactionsRecurringGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given TransactionsRecurringGetRequestOptions and assigns it to the Options field.
func (o *ProcessorTransactionsRecurringGetRequest) SetOptions(v TransactionsRecurringGetRequestOptions) {
	o.Options = &v
}

func (o ProcessorTransactionsRecurringGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["processor_token"] = o.ProcessorToken
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorTransactionsRecurringGetRequest struct {
	value *ProcessorTransactionsRecurringGetRequest
	isSet bool
}

func (v NullableProcessorTransactionsRecurringGetRequest) Get() *ProcessorTransactionsRecurringGetRequest {
	return v.value
}

func (v *NullableProcessorTransactionsRecurringGetRequest) Set(val *ProcessorTransactionsRecurringGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorTransactionsRecurringGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorTransactionsRecurringGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorTransactionsRecurringGetRequest(val *ProcessorTransactionsRecurringGetRequest) *NullableProcessorTransactionsRecurringGetRequest {
	return &NullableProcessorTransactionsRecurringGetRequest{value: val, isSet: true}
}

func (v NullableProcessorTransactionsRecurringGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorTransactionsRecurringGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


