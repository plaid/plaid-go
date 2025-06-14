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

// SandboxPublicTokenCreateRequest SandboxPublicTokenCreateRequest defines the request schema for `/sandbox/public_token/create`
type SandboxPublicTokenCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the institution the Item will be associated with
	InstitutionId string `json:"institution_id"`
	// The products to initially pull for the Item. May be any products that the specified `institution_id`  supports. This array may not be empty.
	InitialProducts []Products `json:"initial_products"`
	Options *SandboxPublicTokenCreateRequestOptions `json:"options,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken *string `json:"user_token,omitempty"`
}

// NewSandboxPublicTokenCreateRequest instantiates a new SandboxPublicTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxPublicTokenCreateRequest(institutionId string, initialProducts []Products) *SandboxPublicTokenCreateRequest {
	this := SandboxPublicTokenCreateRequest{}
	this.InstitutionId = institutionId
	this.InitialProducts = initialProducts
	return &this
}

// NewSandboxPublicTokenCreateRequestWithDefaults instantiates a new SandboxPublicTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxPublicTokenCreateRequestWithDefaults() *SandboxPublicTokenCreateRequest {
	this := SandboxPublicTokenCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxPublicTokenCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxPublicTokenCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxPublicTokenCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxPublicTokenCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetInstitutionId returns the InstitutionId field value
func (o *SandboxPublicTokenCreateRequest) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequest) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *SandboxPublicTokenCreateRequest) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetInitialProducts returns the InitialProducts field value
func (o *SandboxPublicTokenCreateRequest) GetInitialProducts() []Products {
	if o == nil {
		var ret []Products
		return ret
	}

	return o.InitialProducts
}

// GetInitialProductsOk returns a tuple with the InitialProducts field value
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequest) GetInitialProductsOk() (*[]Products, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InitialProducts, true
}

// SetInitialProducts sets field value
func (o *SandboxPublicTokenCreateRequest) SetInitialProducts(v []Products) {
	o.InitialProducts = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SandboxPublicTokenCreateRequest) GetOptions() SandboxPublicTokenCreateRequestOptions {
	if o == nil || o.Options == nil {
		var ret SandboxPublicTokenCreateRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequest) GetOptionsOk() (*SandboxPublicTokenCreateRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given SandboxPublicTokenCreateRequestOptions and assigns it to the Options field.
func (o *SandboxPublicTokenCreateRequest) SetOptions(v SandboxPublicTokenCreateRequestOptions) {
	o.Options = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *SandboxPublicTokenCreateRequest) GetUserToken() string {
	if o == nil || o.UserToken == nil {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequest) GetUserTokenOk() (*string, bool) {
	if o == nil || o.UserToken == nil {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequest) HasUserToken() bool {
	if o != nil && o.UserToken != nil {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *SandboxPublicTokenCreateRequest) SetUserToken(v string) {
	o.UserToken = &v
}

func (o SandboxPublicTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if true {
		toSerialize["initial_products"] = o.InitialProducts
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.UserToken != nil {
		toSerialize["user_token"] = o.UserToken
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxPublicTokenCreateRequest struct {
	value *SandboxPublicTokenCreateRequest
	isSet bool
}

func (v NullableSandboxPublicTokenCreateRequest) Get() *SandboxPublicTokenCreateRequest {
	return v.value
}

func (v *NullableSandboxPublicTokenCreateRequest) Set(val *SandboxPublicTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxPublicTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxPublicTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxPublicTokenCreateRequest(val *SandboxPublicTokenCreateRequest) *NullableSandboxPublicTokenCreateRequest {
	return &NullableSandboxPublicTokenCreateRequest{value: val, isSet: true}
}

func (v NullableSandboxPublicTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxPublicTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


