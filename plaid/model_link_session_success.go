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

// LinkSessionSuccess An object representing an [onSuccess](https://plaid.com/docs/link/web/#onsuccess) callback from Link. This object has been deprecated in favor of the newer [`results.item_add_result`](https://plaid.com/docs/api/link/#link-token-get-response-link-sessions-results-item-add-results), which can support multiple public tokens in a single Link session.
type LinkSessionSuccess struct {
	// Displayed once a user has successfully linked their Item.
	PublicToken string `json:"public_token"`
	Metadata NullableLinkSessionSuccessMetadata `json:"metadata"`
}

// NewLinkSessionSuccess instantiates a new LinkSessionSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkSessionSuccess(publicToken string, metadata NullableLinkSessionSuccessMetadata) *LinkSessionSuccess {
	this := LinkSessionSuccess{}
	this.PublicToken = publicToken
	this.Metadata = metadata
	return &this
}

// NewLinkSessionSuccessWithDefaults instantiates a new LinkSessionSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkSessionSuccessWithDefaults() *LinkSessionSuccess {
	this := LinkSessionSuccess{}
	return &this
}

// GetPublicToken returns the PublicToken field value
func (o *LinkSessionSuccess) GetPublicToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicToken
}

// GetPublicTokenOk returns a tuple with the PublicToken field value
// and a boolean to check if the value has been set.
func (o *LinkSessionSuccess) GetPublicTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicToken, true
}

// SetPublicToken sets field value
func (o *LinkSessionSuccess) SetPublicToken(v string) {
	o.PublicToken = v
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for LinkSessionSuccessMetadata will be returned
func (o *LinkSessionSuccess) GetMetadata() LinkSessionSuccessMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret LinkSessionSuccessMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkSessionSuccess) GetMetadataOk() (*LinkSessionSuccessMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *LinkSessionSuccess) SetMetadata(v LinkSessionSuccessMetadata) {
	o.Metadata.Set(&v)
}

func (o LinkSessionSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["public_token"] = o.PublicToken
	}
	if true {
		toSerialize["metadata"] = o.Metadata.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLinkSessionSuccess struct {
	value *LinkSessionSuccess
	isSet bool
}

func (v NullableLinkSessionSuccess) Get() *LinkSessionSuccess {
	return v.value
}

func (v *NullableLinkSessionSuccess) Set(val *LinkSessionSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkSessionSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkSessionSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkSessionSuccess(val *LinkSessionSuccess) *NullableLinkSessionSuccess {
	return &NullableLinkSessionSuccess{value: val, isSet: true}
}

func (v NullableLinkSessionSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkSessionSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


