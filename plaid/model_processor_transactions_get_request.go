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

// ProcessorTransactionsGetRequest ProcessorTransactionsGetRequest defines the request schema for `/processor/transactions/get`
type ProcessorTransactionsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	Options *ProcessorTransactionsGetRequestOptions `json:"options,omitempty"`
	// The processor token obtained from the Plaid integration partner. Processor tokens are in the format: `processor-<environment>-<identifier>`
	ProcessorToken string `json:"processor_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The earliest date for which data should be returned. Dates should be formatted as YYYY-MM-DD.
	StartDate string `json:"start_date"`
	// The latest date for which data should be returned. Dates should be formatted as YYYY-MM-DD.
	EndDate string `json:"end_date"`
}

// NewProcessorTransactionsGetRequest instantiates a new ProcessorTransactionsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorTransactionsGetRequest(processorToken string, startDate string, endDate string) *ProcessorTransactionsGetRequest {
	this := ProcessorTransactionsGetRequest{}
	this.ProcessorToken = processorToken
	this.StartDate = startDate
	this.EndDate = endDate
	return &this
}

// NewProcessorTransactionsGetRequestWithDefaults instantiates a new ProcessorTransactionsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorTransactionsGetRequestWithDefaults() *ProcessorTransactionsGetRequest {
	this := ProcessorTransactionsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorTransactionsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorTransactionsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorTransactionsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProcessorTransactionsGetRequest) GetOptions() ProcessorTransactionsGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret ProcessorTransactionsGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsGetRequest) GetOptionsOk() (*ProcessorTransactionsGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProcessorTransactionsGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ProcessorTransactionsGetRequestOptions and assigns it to the Options field.
func (o *ProcessorTransactionsGetRequest) SetOptions(v ProcessorTransactionsGetRequestOptions) {
	o.Options = &v
}

// GetProcessorToken returns the ProcessorToken field value
func (o *ProcessorTransactionsGetRequest) GetProcessorToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessorToken
}

// GetProcessorTokenOk returns a tuple with the ProcessorToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsGetRequest) GetProcessorTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProcessorToken, true
}

// SetProcessorToken sets field value
func (o *ProcessorTransactionsGetRequest) SetProcessorToken(v string) {
	o.ProcessorToken = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorTransactionsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorTransactionsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorTransactionsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetStartDate returns the StartDate field value
func (o *ProcessorTransactionsGetRequest) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsGetRequest) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *ProcessorTransactionsGetRequest) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *ProcessorTransactionsGetRequest) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *ProcessorTransactionsGetRequest) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *ProcessorTransactionsGetRequest) SetEndDate(v string) {
	o.EndDate = v
}

func (o ProcessorTransactionsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if true {
		toSerialize["processor_token"] = o.ProcessorToken
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if true {
		toSerialize["end_date"] = o.EndDate
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorTransactionsGetRequest struct {
	value *ProcessorTransactionsGetRequest
	isSet bool
}

func (v NullableProcessorTransactionsGetRequest) Get() *ProcessorTransactionsGetRequest {
	return v.value
}

func (v *NullableProcessorTransactionsGetRequest) Set(val *ProcessorTransactionsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorTransactionsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorTransactionsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorTransactionsGetRequest(val *ProcessorTransactionsGetRequest) *NullableProcessorTransactionsGetRequest {
	return &NullableProcessorTransactionsGetRequest{value: val, isSet: true}
}

func (v NullableProcessorTransactionsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorTransactionsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


