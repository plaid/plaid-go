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
	"time"
)

// ItemPublicTokenCreateResponse ItemPublicTokenCreateResponse defines the response schema for `/item/public_token/create`
type ItemPublicTokenCreateResponse struct {
	// A `public_token` for the particular Item corresponding to the specified `access_token`
	PublicToken string `json:"public_token"`
	Expiration *time.Time `json:"expiration,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ItemPublicTokenCreateResponse ItemPublicTokenCreateResponse

// NewItemPublicTokenCreateResponse instantiates a new ItemPublicTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPublicTokenCreateResponse(publicToken string, requestId string) *ItemPublicTokenCreateResponse {
	this := ItemPublicTokenCreateResponse{}
	this.PublicToken = publicToken
	this.RequestId = requestId
	return &this
}

// NewItemPublicTokenCreateResponseWithDefaults instantiates a new ItemPublicTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPublicTokenCreateResponseWithDefaults() *ItemPublicTokenCreateResponse {
	this := ItemPublicTokenCreateResponse{}
	return &this
}

// GetPublicToken returns the PublicToken field value
func (o *ItemPublicTokenCreateResponse) GetPublicToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicToken
}

// GetPublicTokenOk returns a tuple with the PublicToken field value
// and a boolean to check if the value has been set.
func (o *ItemPublicTokenCreateResponse) GetPublicTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicToken, true
}

// SetPublicToken sets field value
func (o *ItemPublicTokenCreateResponse) SetPublicToken(v string) {
	o.PublicToken = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ItemPublicTokenCreateResponse) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPublicTokenCreateResponse) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ItemPublicTokenCreateResponse) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *ItemPublicTokenCreateResponse) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetRequestId returns the RequestId field value
func (o *ItemPublicTokenCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ItemPublicTokenCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ItemPublicTokenCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ItemPublicTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["public_token"] = o.PublicToken
	}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemPublicTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varItemPublicTokenCreateResponse := _ItemPublicTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varItemPublicTokenCreateResponse); err == nil {
		*o = ItemPublicTokenCreateResponse(varItemPublicTokenCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "public_token")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemPublicTokenCreateResponse struct {
	value *ItemPublicTokenCreateResponse
	isSet bool
}

func (v NullableItemPublicTokenCreateResponse) Get() *ItemPublicTokenCreateResponse {
	return v.value
}

func (v *NullableItemPublicTokenCreateResponse) Set(val *ItemPublicTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPublicTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPublicTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPublicTokenCreateResponse(val *ItemPublicTokenCreateResponse) *NullableItemPublicTokenCreateResponse {
	return &NullableItemPublicTokenCreateResponse{value: val, isSet: true}
}

func (v NullableItemPublicTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemPublicTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


