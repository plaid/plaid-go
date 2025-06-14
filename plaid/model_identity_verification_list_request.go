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

// IdentityVerificationListRequest Request input for listing Identity Verifications
type IdentityVerificationListRequest struct {
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// ID of the associated Identity Verification template.
	TemplateId string `json:"template_id"`
	// A unique ID that identifies the end user in your system. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.
	ClientUserId string `json:"client_user_id"`
	// An identifier that determines which page of results you receive.
	Cursor NullableString `json:"cursor,omitempty"`
}

// NewIdentityVerificationListRequest instantiates a new IdentityVerificationListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationListRequest(templateId string, clientUserId string) *IdentityVerificationListRequest {
	this := IdentityVerificationListRequest{}
	this.TemplateId = templateId
	this.ClientUserId = clientUserId
	return &this
}

// NewIdentityVerificationListRequestWithDefaults instantiates a new IdentityVerificationListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationListRequestWithDefaults() *IdentityVerificationListRequest {
	this := IdentityVerificationListRequest{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IdentityVerificationListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IdentityVerificationListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IdentityVerificationListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IdentityVerificationListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IdentityVerificationListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IdentityVerificationListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetTemplateId returns the TemplateId field value
func (o *IdentityVerificationListRequest) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationListRequest) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *IdentityVerificationListRequest) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetClientUserId returns the ClientUserId field value
func (o *IdentityVerificationListRequest) GetClientUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationListRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientUserId, true
}

// SetClientUserId sets field value
func (o *IdentityVerificationListRequest) SetClientUserId(v string) {
	o.ClientUserId = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationListRequest) GetCursor() string {
	if o == nil || o.Cursor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationListRequest) GetCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *IdentityVerificationListRequest) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *IdentityVerificationListRequest) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *IdentityVerificationListRequest) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *IdentityVerificationListRequest) UnsetCursor() {
	o.Cursor.Unset()
}

func (o IdentityVerificationListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["template_id"] = o.TemplateId
	}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityVerificationListRequest struct {
	value *IdentityVerificationListRequest
	isSet bool
}

func (v NullableIdentityVerificationListRequest) Get() *IdentityVerificationListRequest {
	return v.value
}

func (v *NullableIdentityVerificationListRequest) Set(val *IdentityVerificationListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationListRequest(val *IdentityVerificationListRequest) *NullableIdentityVerificationListRequest {
	return &NullableIdentityVerificationListRequest{value: val, isSet: true}
}

func (v NullableIdentityVerificationListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


