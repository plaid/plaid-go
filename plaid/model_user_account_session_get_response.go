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

// UserAccountSessionGetResponse UserAccountSessionGetResponse defines the response schema for `/user_account/session/get`
type UserAccountSessionGetResponse struct {
	Identity NullableUserAccountIdentity `json:"identity"`
	Items []UserAccountItem `json:"items"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _UserAccountSessionGetResponse UserAccountSessionGetResponse

// NewUserAccountSessionGetResponse instantiates a new UserAccountSessionGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccountSessionGetResponse(identity NullableUserAccountIdentity, items []UserAccountItem, requestId string) *UserAccountSessionGetResponse {
	this := UserAccountSessionGetResponse{}
	this.Identity = identity
	this.Items = items
	this.RequestId = requestId
	return &this
}

// NewUserAccountSessionGetResponseWithDefaults instantiates a new UserAccountSessionGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccountSessionGetResponseWithDefaults() *UserAccountSessionGetResponse {
	this := UserAccountSessionGetResponse{}
	return &this
}

// GetIdentity returns the Identity field value
// If the value is explicit nil, the zero value for UserAccountIdentity will be returned
func (o *UserAccountSessionGetResponse) GetIdentity() UserAccountIdentity {
	if o == nil || o.Identity.Get() == nil {
		var ret UserAccountIdentity
		return ret
	}

	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAccountSessionGetResponse) GetIdentityOk() (*UserAccountIdentity, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// SetIdentity sets field value
func (o *UserAccountSessionGetResponse) SetIdentity(v UserAccountIdentity) {
	o.Identity.Set(&v)
}

// GetItems returns the Items field value
func (o *UserAccountSessionGetResponse) GetItems() []UserAccountItem {
	if o == nil {
		var ret []UserAccountItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *UserAccountSessionGetResponse) GetItemsOk() (*[]UserAccountItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *UserAccountSessionGetResponse) SetItems(v []UserAccountItem) {
	o.Items = v
}

// GetRequestId returns the RequestId field value
func (o *UserAccountSessionGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *UserAccountSessionGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *UserAccountSessionGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o UserAccountSessionGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity"] = o.Identity.Get()
	}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserAccountSessionGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varUserAccountSessionGetResponse := _UserAccountSessionGetResponse{}

	if err = json.Unmarshal(bytes, &varUserAccountSessionGetResponse); err == nil {
		*o = UserAccountSessionGetResponse(varUserAccountSessionGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "items")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserAccountSessionGetResponse struct {
	value *UserAccountSessionGetResponse
	isSet bool
}

func (v NullableUserAccountSessionGetResponse) Get() *UserAccountSessionGetResponse {
	return v.value
}

func (v *NullableUserAccountSessionGetResponse) Set(val *UserAccountSessionGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccountSessionGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccountSessionGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccountSessionGetResponse(val *UserAccountSessionGetResponse) *NullableUserAccountSessionGetResponse {
	return &NullableUserAccountSessionGetResponse{value: val, isSet: true}
}

func (v NullableUserAccountSessionGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccountSessionGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

