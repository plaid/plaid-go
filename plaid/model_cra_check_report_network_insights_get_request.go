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

// CraCheckReportNetworkInsightsGetRequest CraCheckReportNetworkInsightsGetRequest defines the request schema for `/cra/check_report/network_attributes/get`.
type CraCheckReportNetworkInsightsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken *string `json:"user_token,omitempty"`
	// The third-party user token associated with the requested User data.
	ThirdPartyUserToken *string `json:"third_party_user_token,omitempty"`
}

// NewCraCheckReportNetworkInsightsGetRequest instantiates a new CraCheckReportNetworkInsightsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraCheckReportNetworkInsightsGetRequest() *CraCheckReportNetworkInsightsGetRequest {
	this := CraCheckReportNetworkInsightsGetRequest{}
	return &this
}

// NewCraCheckReportNetworkInsightsGetRequestWithDefaults instantiates a new CraCheckReportNetworkInsightsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraCheckReportNetworkInsightsGetRequestWithDefaults() *CraCheckReportNetworkInsightsGetRequest {
	this := CraCheckReportNetworkInsightsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CraCheckReportNetworkInsightsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraCheckReportNetworkInsightsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CraCheckReportNetworkInsightsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CraCheckReportNetworkInsightsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CraCheckReportNetworkInsightsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraCheckReportNetworkInsightsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CraCheckReportNetworkInsightsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CraCheckReportNetworkInsightsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *CraCheckReportNetworkInsightsGetRequest) GetUserToken() string {
	if o == nil || o.UserToken == nil {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraCheckReportNetworkInsightsGetRequest) GetUserTokenOk() (*string, bool) {
	if o == nil || o.UserToken == nil {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *CraCheckReportNetworkInsightsGetRequest) HasUserToken() bool {
	if o != nil && o.UserToken != nil {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *CraCheckReportNetworkInsightsGetRequest) SetUserToken(v string) {
	o.UserToken = &v
}

// GetThirdPartyUserToken returns the ThirdPartyUserToken field value if set, zero value otherwise.
func (o *CraCheckReportNetworkInsightsGetRequest) GetThirdPartyUserToken() string {
	if o == nil || o.ThirdPartyUserToken == nil {
		var ret string
		return ret
	}
	return *o.ThirdPartyUserToken
}

// GetThirdPartyUserTokenOk returns a tuple with the ThirdPartyUserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraCheckReportNetworkInsightsGetRequest) GetThirdPartyUserTokenOk() (*string, bool) {
	if o == nil || o.ThirdPartyUserToken == nil {
		return nil, false
	}
	return o.ThirdPartyUserToken, true
}

// HasThirdPartyUserToken returns a boolean if a field has been set.
func (o *CraCheckReportNetworkInsightsGetRequest) HasThirdPartyUserToken() bool {
	if o != nil && o.ThirdPartyUserToken != nil {
		return true
	}

	return false
}

// SetThirdPartyUserToken gets a reference to the given string and assigns it to the ThirdPartyUserToken field.
func (o *CraCheckReportNetworkInsightsGetRequest) SetThirdPartyUserToken(v string) {
	o.ThirdPartyUserToken = &v
}

func (o CraCheckReportNetworkInsightsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.UserToken != nil {
		toSerialize["user_token"] = o.UserToken
	}
	if o.ThirdPartyUserToken != nil {
		toSerialize["third_party_user_token"] = o.ThirdPartyUserToken
	}
	return json.Marshal(toSerialize)
}

type NullableCraCheckReportNetworkInsightsGetRequest struct {
	value *CraCheckReportNetworkInsightsGetRequest
	isSet bool
}

func (v NullableCraCheckReportNetworkInsightsGetRequest) Get() *CraCheckReportNetworkInsightsGetRequest {
	return v.value
}

func (v *NullableCraCheckReportNetworkInsightsGetRequest) Set(val *CraCheckReportNetworkInsightsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCraCheckReportNetworkInsightsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCraCheckReportNetworkInsightsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraCheckReportNetworkInsightsGetRequest(val *CraCheckReportNetworkInsightsGetRequest) *NullableCraCheckReportNetworkInsightsGetRequest {
	return &NullableCraCheckReportNetworkInsightsGetRequest{value: val, isSet: true}
}

func (v NullableCraCheckReportNetworkInsightsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraCheckReportNetworkInsightsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


