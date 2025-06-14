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

// CraMonitoringInsightsGetRequest CraMonitoringInsightsGetRequest defines the request schema for `/cra/monitoring_insights/get`
type CraMonitoringInsightsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken string `json:"user_token"`
	ConsumerReportPermissiblePurpose MonitoringConsumerReportPermissiblePurpose `json:"consumer_report_permissible_purpose"`
}

// NewCraMonitoringInsightsGetRequest instantiates a new CraMonitoringInsightsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraMonitoringInsightsGetRequest(userToken string, consumerReportPermissiblePurpose MonitoringConsumerReportPermissiblePurpose) *CraMonitoringInsightsGetRequest {
	this := CraMonitoringInsightsGetRequest{}
	this.UserToken = userToken
	this.ConsumerReportPermissiblePurpose = consumerReportPermissiblePurpose
	return &this
}

// NewCraMonitoringInsightsGetRequestWithDefaults instantiates a new CraMonitoringInsightsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraMonitoringInsightsGetRequestWithDefaults() *CraMonitoringInsightsGetRequest {
	this := CraMonitoringInsightsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CraMonitoringInsightsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraMonitoringInsightsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CraMonitoringInsightsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CraMonitoringInsightsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CraMonitoringInsightsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraMonitoringInsightsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CraMonitoringInsightsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CraMonitoringInsightsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value
func (o *CraMonitoringInsightsGetRequest) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *CraMonitoringInsightsGetRequest) GetUserTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *CraMonitoringInsightsGetRequest) SetUserToken(v string) {
	o.UserToken = v
}

// GetConsumerReportPermissiblePurpose returns the ConsumerReportPermissiblePurpose field value
func (o *CraMonitoringInsightsGetRequest) GetConsumerReportPermissiblePurpose() MonitoringConsumerReportPermissiblePurpose {
	if o == nil {
		var ret MonitoringConsumerReportPermissiblePurpose
		return ret
	}

	return o.ConsumerReportPermissiblePurpose
}

// GetConsumerReportPermissiblePurposeOk returns a tuple with the ConsumerReportPermissiblePurpose field value
// and a boolean to check if the value has been set.
func (o *CraMonitoringInsightsGetRequest) GetConsumerReportPermissiblePurposeOk() (*MonitoringConsumerReportPermissiblePurpose, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsumerReportPermissiblePurpose, true
}

// SetConsumerReportPermissiblePurpose sets field value
func (o *CraMonitoringInsightsGetRequest) SetConsumerReportPermissiblePurpose(v MonitoringConsumerReportPermissiblePurpose) {
	o.ConsumerReportPermissiblePurpose = v
}

func (o CraMonitoringInsightsGetRequest) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["consumer_report_permissible_purpose"] = o.ConsumerReportPermissiblePurpose
	}
	return json.Marshal(toSerialize)
}

type NullableCraMonitoringInsightsGetRequest struct {
	value *CraMonitoringInsightsGetRequest
	isSet bool
}

func (v NullableCraMonitoringInsightsGetRequest) Get() *CraMonitoringInsightsGetRequest {
	return v.value
}

func (v *NullableCraMonitoringInsightsGetRequest) Set(val *CraMonitoringInsightsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCraMonitoringInsightsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCraMonitoringInsightsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraMonitoringInsightsGetRequest(val *CraMonitoringInsightsGetRequest) *NullableCraMonitoringInsightsGetRequest {
	return &NullableCraMonitoringInsightsGetRequest{value: val, isSet: true}
}

func (v NullableCraMonitoringInsightsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraMonitoringInsightsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


