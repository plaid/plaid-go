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

// BeaconUserAccountInsightsGetRequest Request input for fetching the risk insights for a Beacon User's Bank Accounts
type BeaconUserAccountInsightsGetRequest struct {
	// ID of the associated Beacon User.
	BeaconUserId string `json:"beacon_user_id"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
}

// NewBeaconUserAccountInsightsGetRequest instantiates a new BeaconUserAccountInsightsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeaconUserAccountInsightsGetRequest(beaconUserId string, accessToken string) *BeaconUserAccountInsightsGetRequest {
	this := BeaconUserAccountInsightsGetRequest{}
	this.BeaconUserId = beaconUserId
	this.AccessToken = accessToken
	return &this
}

// NewBeaconUserAccountInsightsGetRequestWithDefaults instantiates a new BeaconUserAccountInsightsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeaconUserAccountInsightsGetRequestWithDefaults() *BeaconUserAccountInsightsGetRequest {
	this := BeaconUserAccountInsightsGetRequest{}
	return &this
}

// GetBeaconUserId returns the BeaconUserId field value
func (o *BeaconUserAccountInsightsGetRequest) GetBeaconUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeaconUserId
}

// GetBeaconUserIdOk returns a tuple with the BeaconUserId field value
// and a boolean to check if the value has been set.
func (o *BeaconUserAccountInsightsGetRequest) GetBeaconUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BeaconUserId, true
}

// SetBeaconUserId sets field value
func (o *BeaconUserAccountInsightsGetRequest) SetBeaconUserId(v string) {
	o.BeaconUserId = v
}

// GetAccessToken returns the AccessToken field value
func (o *BeaconUserAccountInsightsGetRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *BeaconUserAccountInsightsGetRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *BeaconUserAccountInsightsGetRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BeaconUserAccountInsightsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeaconUserAccountInsightsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BeaconUserAccountInsightsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BeaconUserAccountInsightsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *BeaconUserAccountInsightsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeaconUserAccountInsightsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *BeaconUserAccountInsightsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *BeaconUserAccountInsightsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o BeaconUserAccountInsightsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["beacon_user_id"] = o.BeaconUserId
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableBeaconUserAccountInsightsGetRequest struct {
	value *BeaconUserAccountInsightsGetRequest
	isSet bool
}

func (v NullableBeaconUserAccountInsightsGetRequest) Get() *BeaconUserAccountInsightsGetRequest {
	return v.value
}

func (v *NullableBeaconUserAccountInsightsGetRequest) Set(val *BeaconUserAccountInsightsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconUserAccountInsightsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconUserAccountInsightsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconUserAccountInsightsGetRequest(val *BeaconUserAccountInsightsGetRequest) *NullableBeaconUserAccountInsightsGetRequest {
	return &NullableBeaconUserAccountInsightsGetRequest{value: val, isSet: true}
}

func (v NullableBeaconUserAccountInsightsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconUserAccountInsightsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


