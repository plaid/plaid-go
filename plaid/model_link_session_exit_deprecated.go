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

// LinkSessionExitDeprecated An object representing an [onExit](https://plaid.com/docs/link/web/#onexit) callback from Link. This field has been deprecated in favor of [`exit`](https://plaid.com/docs/api/link/#link-token-get-response-link-sessions-exit), for improved naming consistency.
type LinkSessionExitDeprecated struct {
	Error NullablePlaidError `json:"error"`
	Metadata NullableLinkSessionExitMetadata `json:"metadata"`
	AdditionalProperties map[string]interface{}
}

type _LinkSessionExitDeprecated LinkSessionExitDeprecated

// NewLinkSessionExitDeprecated instantiates a new LinkSessionExitDeprecated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkSessionExitDeprecated(error_ NullablePlaidError, metadata NullableLinkSessionExitMetadata) *LinkSessionExitDeprecated {
	this := LinkSessionExitDeprecated{}
	this.Error = error_
	this.Metadata = metadata
	return &this
}

// NewLinkSessionExitDeprecatedWithDefaults instantiates a new LinkSessionExitDeprecated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkSessionExitDeprecatedWithDefaults() *LinkSessionExitDeprecated {
	this := LinkSessionExitDeprecated{}
	return &this
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for PlaidError will be returned
func (o *LinkSessionExitDeprecated) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkSessionExitDeprecated) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *LinkSessionExitDeprecated) SetError(v PlaidError) {
	o.Error.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for LinkSessionExitMetadata will be returned
func (o *LinkSessionExitDeprecated) GetMetadata() LinkSessionExitMetadata {
	if o == nil || o.Metadata.Get() == nil {
		var ret LinkSessionExitMetadata
		return ret
	}

	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkSessionExitDeprecated) GetMetadataOk() (*LinkSessionExitMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// SetMetadata sets field value
func (o *LinkSessionExitDeprecated) SetMetadata(v LinkSessionExitMetadata) {
	o.Metadata.Set(&v)
}

func (o LinkSessionExitDeprecated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["metadata"] = o.Metadata.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkSessionExitDeprecated) UnmarshalJSON(bytes []byte) (err error) {
	varLinkSessionExitDeprecated := _LinkSessionExitDeprecated{}

	if err = json.Unmarshal(bytes, &varLinkSessionExitDeprecated); err == nil {
		*o = LinkSessionExitDeprecated(varLinkSessionExitDeprecated)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkSessionExitDeprecated struct {
	value *LinkSessionExitDeprecated
	isSet bool
}

func (v NullableLinkSessionExitDeprecated) Get() *LinkSessionExitDeprecated {
	return v.value
}

func (v *NullableLinkSessionExitDeprecated) Set(val *LinkSessionExitDeprecated) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkSessionExitDeprecated) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkSessionExitDeprecated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkSessionExitDeprecated(val *LinkSessionExitDeprecated) *NullableLinkSessionExitDeprecated {
	return &NullableLinkSessionExitDeprecated{value: val, isSet: true}
}

func (v NullableLinkSessionExitDeprecated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkSessionExitDeprecated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


