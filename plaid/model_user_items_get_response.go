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

// UserItemsGetResponse UserItemsGetResponse defines the response schema for `/user/items/get`
type UserItemsGetResponse struct {
	Items []Item `json:"items"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _UserItemsGetResponse UserItemsGetResponse

// NewUserItemsGetResponse instantiates a new UserItemsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserItemsGetResponse(items []Item, requestId string) *UserItemsGetResponse {
	this := UserItemsGetResponse{}
	this.Items = items
	this.RequestId = requestId
	return &this
}

// NewUserItemsGetResponseWithDefaults instantiates a new UserItemsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserItemsGetResponseWithDefaults() *UserItemsGetResponse {
	this := UserItemsGetResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *UserItemsGetResponse) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *UserItemsGetResponse) GetItemsOk() (*[]Item, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *UserItemsGetResponse) SetItems(v []Item) {
	o.Items = v
}

// GetRequestId returns the RequestId field value
func (o *UserItemsGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *UserItemsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *UserItemsGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o UserItemsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *UserItemsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varUserItemsGetResponse := _UserItemsGetResponse{}

	if err = json.Unmarshal(bytes, &varUserItemsGetResponse); err == nil {
		*o = UserItemsGetResponse(varUserItemsGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserItemsGetResponse struct {
	value *UserItemsGetResponse
	isSet bool
}

func (v NullableUserItemsGetResponse) Get() *UserItemsGetResponse {
	return v.value
}

func (v *NullableUserItemsGetResponse) Set(val *UserItemsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserItemsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserItemsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserItemsGetResponse(val *UserItemsGetResponse) *NullableUserItemsGetResponse {
	return &NullableUserItemsGetResponse{value: val, isSet: true}
}

func (v NullableUserItemsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserItemsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


