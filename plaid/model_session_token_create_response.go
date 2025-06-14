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

// SessionTokenCreateResponse SessionTokenCreateResponse defines the response schema for `/session/token/create`
type SessionTokenCreateResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	Link *SessionTokenCreateResponseLink `json:"link,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionTokenCreateResponse SessionTokenCreateResponse

// NewSessionTokenCreateResponse instantiates a new SessionTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionTokenCreateResponse(requestId string) *SessionTokenCreateResponse {
	this := SessionTokenCreateResponse{}
	this.RequestId = requestId
	return &this
}

// NewSessionTokenCreateResponseWithDefaults instantiates a new SessionTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionTokenCreateResponseWithDefaults() *SessionTokenCreateResponse {
	this := SessionTokenCreateResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *SessionTokenCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SessionTokenCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SessionTokenCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *SessionTokenCreateResponse) GetLink() SessionTokenCreateResponseLink {
	if o == nil || o.Link == nil {
		var ret SessionTokenCreateResponseLink
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionTokenCreateResponse) GetLinkOk() (*SessionTokenCreateResponseLink, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *SessionTokenCreateResponse) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given SessionTokenCreateResponseLink and assigns it to the Link field.
func (o *SessionTokenCreateResponse) SetLink(v SessionTokenCreateResponseLink) {
	o.Link = &v
}

func (o SessionTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SessionTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSessionTokenCreateResponse := _SessionTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varSessionTokenCreateResponse); err == nil {
		*o = SessionTokenCreateResponse(varSessionTokenCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "link")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionTokenCreateResponse struct {
	value *SessionTokenCreateResponse
	isSet bool
}

func (v NullableSessionTokenCreateResponse) Get() *SessionTokenCreateResponse {
	return v.value
}

func (v *NullableSessionTokenCreateResponse) Set(val *SessionTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionTokenCreateResponse(val *SessionTokenCreateResponse) *NullableSessionTokenCreateResponse {
	return &NullableSessionTokenCreateResponse{value: val, isSet: true}
}

func (v NullableSessionTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


