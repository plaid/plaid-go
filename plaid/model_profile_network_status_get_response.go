/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.586.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ProfileNetworkStatusGetResponse ProfileNetworkStatusGetResponse defines the response schema for `/profile/network_status/get`
type ProfileNetworkStatusGetResponse struct {
	NetworkStatus ProfileNetworkStatusGetNetworkStatus `json:"network_status"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ProfileNetworkStatusGetResponse ProfileNetworkStatusGetResponse

// NewProfileNetworkStatusGetResponse instantiates a new ProfileNetworkStatusGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileNetworkStatusGetResponse(networkStatus ProfileNetworkStatusGetNetworkStatus, requestId string) *ProfileNetworkStatusGetResponse {
	this := ProfileNetworkStatusGetResponse{}
	this.NetworkStatus = networkStatus
	this.RequestId = requestId
	return &this
}

// NewProfileNetworkStatusGetResponseWithDefaults instantiates a new ProfileNetworkStatusGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileNetworkStatusGetResponseWithDefaults() *ProfileNetworkStatusGetResponse {
	this := ProfileNetworkStatusGetResponse{}
	return &this
}

// GetNetworkStatus returns the NetworkStatus field value
func (o *ProfileNetworkStatusGetResponse) GetNetworkStatus() ProfileNetworkStatusGetNetworkStatus {
	if o == nil {
		var ret ProfileNetworkStatusGetNetworkStatus
		return ret
	}

	return o.NetworkStatus
}

// GetNetworkStatusOk returns a tuple with the NetworkStatus field value
// and a boolean to check if the value has been set.
func (o *ProfileNetworkStatusGetResponse) GetNetworkStatusOk() (*ProfileNetworkStatusGetNetworkStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetworkStatus, true
}

// SetNetworkStatus sets field value
func (o *ProfileNetworkStatusGetResponse) SetNetworkStatus(v ProfileNetworkStatusGetNetworkStatus) {
	o.NetworkStatus = v
}

// GetRequestId returns the RequestId field value
func (o *ProfileNetworkStatusGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProfileNetworkStatusGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProfileNetworkStatusGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ProfileNetworkStatusGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["network_status"] = o.NetworkStatus
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProfileNetworkStatusGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProfileNetworkStatusGetResponse := _ProfileNetworkStatusGetResponse{}

	if err = json.Unmarshal(bytes, &varProfileNetworkStatusGetResponse); err == nil {
		*o = ProfileNetworkStatusGetResponse(varProfileNetworkStatusGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "network_status")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProfileNetworkStatusGetResponse struct {
	value *ProfileNetworkStatusGetResponse
	isSet bool
}

func (v NullableProfileNetworkStatusGetResponse) Get() *ProfileNetworkStatusGetResponse {
	return v.value
}

func (v *NullableProfileNetworkStatusGetResponse) Set(val *ProfileNetworkStatusGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileNetworkStatusGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileNetworkStatusGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileNetworkStatusGetResponse(val *ProfileNetworkStatusGetResponse) *NullableProfileNetworkStatusGetResponse {
	return &NullableProfileNetworkStatusGetResponse{value: val, isSet: true}
}

func (v NullableProfileNetworkStatusGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileNetworkStatusGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

