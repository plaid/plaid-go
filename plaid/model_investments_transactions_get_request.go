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

// InvestmentsTransactionsGetRequest InvestmentsTransactionsGetRequest defines the request schema for `/investments/transactions/get`
type InvestmentsTransactionsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// The earliest date for which to fetch transaction history. Dates should be formatted as YYYY-MM-DD.
	StartDate string `json:"start_date"`
	// The most recent date for which to fetch transaction history. Dates should be formatted as YYYY-MM-DD.
	EndDate string `json:"end_date"`
	Options *InvestmentsTransactionsGetRequestOptions `json:"options,omitempty"`
}

// NewInvestmentsTransactionsGetRequest instantiates a new InvestmentsTransactionsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsTransactionsGetRequest(accessToken string, startDate string, endDate string) *InvestmentsTransactionsGetRequest {
	this := InvestmentsTransactionsGetRequest{}
	this.AccessToken = accessToken
	this.StartDate = startDate
	this.EndDate = endDate
	return &this
}

// NewInvestmentsTransactionsGetRequestWithDefaults instantiates a new InvestmentsTransactionsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsTransactionsGetRequestWithDefaults() *InvestmentsTransactionsGetRequest {
	this := InvestmentsTransactionsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InvestmentsTransactionsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InvestmentsTransactionsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InvestmentsTransactionsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *InvestmentsTransactionsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *InvestmentsTransactionsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *InvestmentsTransactionsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *InvestmentsTransactionsGetRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *InvestmentsTransactionsGetRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetStartDate returns the StartDate field value
func (o *InvestmentsTransactionsGetRequest) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetRequest) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *InvestmentsTransactionsGetRequest) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *InvestmentsTransactionsGetRequest) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetRequest) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *InvestmentsTransactionsGetRequest) SetEndDate(v string) {
	o.EndDate = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InvestmentsTransactionsGetRequest) GetOptions() InvestmentsTransactionsGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret InvestmentsTransactionsGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetRequest) GetOptionsOk() (*InvestmentsTransactionsGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InvestmentsTransactionsGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given InvestmentsTransactionsGetRequestOptions and assigns it to the Options field.
func (o *InvestmentsTransactionsGetRequest) SetOptions(v InvestmentsTransactionsGetRequestOptions) {
	o.Options = &v
}

func (o InvestmentsTransactionsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if true {
		toSerialize["end_date"] = o.EndDate
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableInvestmentsTransactionsGetRequest struct {
	value *InvestmentsTransactionsGetRequest
	isSet bool
}

func (v NullableInvestmentsTransactionsGetRequest) Get() *InvestmentsTransactionsGetRequest {
	return v.value
}

func (v *NullableInvestmentsTransactionsGetRequest) Set(val *InvestmentsTransactionsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsTransactionsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsTransactionsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsTransactionsGetRequest(val *InvestmentsTransactionsGetRequest) *NullableInvestmentsTransactionsGetRequest {
	return &NullableInvestmentsTransactionsGetRequest{value: val, isSet: true}
}

func (v NullableInvestmentsTransactionsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsTransactionsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


