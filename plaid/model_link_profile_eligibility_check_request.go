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

// LinkProfileEligibilityCheckRequest LinkProfileEligibilityCheckRequest defines the request schema for `/link/profile/eligibility/check`
type LinkProfileEligibilityCheckRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The unique ID for the user's Link session
	LinkSessionId string `json:"link_session_id"`
	User LinkProfileEligibilityCheckUser `json:"user"`
}

// NewLinkProfileEligibilityCheckRequest instantiates a new LinkProfileEligibilityCheckRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkProfileEligibilityCheckRequest(linkSessionId string, user LinkProfileEligibilityCheckUser) *LinkProfileEligibilityCheckRequest {
	this := LinkProfileEligibilityCheckRequest{}
	this.LinkSessionId = linkSessionId
	this.User = user
	return &this
}

// NewLinkProfileEligibilityCheckRequestWithDefaults instantiates a new LinkProfileEligibilityCheckRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkProfileEligibilityCheckRequestWithDefaults() *LinkProfileEligibilityCheckRequest {
	this := LinkProfileEligibilityCheckRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *LinkProfileEligibilityCheckRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkProfileEligibilityCheckRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *LinkProfileEligibilityCheckRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *LinkProfileEligibilityCheckRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *LinkProfileEligibilityCheckRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkProfileEligibilityCheckRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *LinkProfileEligibilityCheckRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *LinkProfileEligibilityCheckRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetLinkSessionId returns the LinkSessionId field value
func (o *LinkProfileEligibilityCheckRequest) GetLinkSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkSessionId
}

// GetLinkSessionIdOk returns a tuple with the LinkSessionId field value
// and a boolean to check if the value has been set.
func (o *LinkProfileEligibilityCheckRequest) GetLinkSessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkSessionId, true
}

// SetLinkSessionId sets field value
func (o *LinkProfileEligibilityCheckRequest) SetLinkSessionId(v string) {
	o.LinkSessionId = v
}

// GetUser returns the User field value
func (o *LinkProfileEligibilityCheckRequest) GetUser() LinkProfileEligibilityCheckUser {
	if o == nil {
		var ret LinkProfileEligibilityCheckUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *LinkProfileEligibilityCheckRequest) GetUserOk() (*LinkProfileEligibilityCheckUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *LinkProfileEligibilityCheckRequest) SetUser(v LinkProfileEligibilityCheckUser) {
	o.User = v
}

func (o LinkProfileEligibilityCheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["link_session_id"] = o.LinkSessionId
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableLinkProfileEligibilityCheckRequest struct {
	value *LinkProfileEligibilityCheckRequest
	isSet bool
}

func (v NullableLinkProfileEligibilityCheckRequest) Get() *LinkProfileEligibilityCheckRequest {
	return v.value
}

func (v *NullableLinkProfileEligibilityCheckRequest) Set(val *LinkProfileEligibilityCheckRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkProfileEligibilityCheckRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkProfileEligibilityCheckRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkProfileEligibilityCheckRequest(val *LinkProfileEligibilityCheckRequest) *NullableLinkProfileEligibilityCheckRequest {
	return &NullableLinkProfileEligibilityCheckRequest{value: val, isSet: true}
}

func (v NullableLinkProfileEligibilityCheckRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkProfileEligibilityCheckRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


