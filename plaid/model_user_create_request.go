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

// UserCreateRequest UserCreateRequest defines the request schema for `/user/create`
type UserCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A unique ID representing the end user. Maximum of 128 characters. Typically this will be a user ID number from your application. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.
	ClientUserId string `json:"client_user_id"`
	// A unique ID representing a CRA reseller's end customer. Maximum of 128 characters.
	EndCustomer *string `json:"end_customer,omitempty"`
	ConsumerReportUserIdentity NullableConsumerReportUserIdentity `json:"consumer_report_user_identity,omitempty"`
}

// NewUserCreateRequest instantiates a new UserCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreateRequest(clientUserId string) *UserCreateRequest {
	this := UserCreateRequest{}
	this.ClientUserId = clientUserId
	return &this
}

// NewUserCreateRequestWithDefaults instantiates a new UserCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateRequestWithDefaults() *UserCreateRequest {
	this := UserCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *UserCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *UserCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *UserCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *UserCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *UserCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *UserCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientUserId returns the ClientUserId field value
func (o *UserCreateRequest) GetClientUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
func (o *UserCreateRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientUserId, true
}

// SetClientUserId sets field value
func (o *UserCreateRequest) SetClientUserId(v string) {
	o.ClientUserId = v
}

// GetEndCustomer returns the EndCustomer field value if set, zero value otherwise.
func (o *UserCreateRequest) GetEndCustomer() string {
	if o == nil || o.EndCustomer == nil {
		var ret string
		return ret
	}
	return *o.EndCustomer
}

// GetEndCustomerOk returns a tuple with the EndCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreateRequest) GetEndCustomerOk() (*string, bool) {
	if o == nil || o.EndCustomer == nil {
		return nil, false
	}
	return o.EndCustomer, true
}

// HasEndCustomer returns a boolean if a field has been set.
func (o *UserCreateRequest) HasEndCustomer() bool {
	if o != nil && o.EndCustomer != nil {
		return true
	}

	return false
}

// SetEndCustomer gets a reference to the given string and assigns it to the EndCustomer field.
func (o *UserCreateRequest) SetEndCustomer(v string) {
	o.EndCustomer = &v
}

// GetConsumerReportUserIdentity returns the ConsumerReportUserIdentity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCreateRequest) GetConsumerReportUserIdentity() ConsumerReportUserIdentity {
	if o == nil || o.ConsumerReportUserIdentity.Get() == nil {
		var ret ConsumerReportUserIdentity
		return ret
	}
	return *o.ConsumerReportUserIdentity.Get()
}

// GetConsumerReportUserIdentityOk returns a tuple with the ConsumerReportUserIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCreateRequest) GetConsumerReportUserIdentityOk() (*ConsumerReportUserIdentity, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConsumerReportUserIdentity.Get(), o.ConsumerReportUserIdentity.IsSet()
}

// HasConsumerReportUserIdentity returns a boolean if a field has been set.
func (o *UserCreateRequest) HasConsumerReportUserIdentity() bool {
	if o != nil && o.ConsumerReportUserIdentity.IsSet() {
		return true
	}

	return false
}

// SetConsumerReportUserIdentity gets a reference to the given NullableConsumerReportUserIdentity and assigns it to the ConsumerReportUserIdentity field.
func (o *UserCreateRequest) SetConsumerReportUserIdentity(v ConsumerReportUserIdentity) {
	o.ConsumerReportUserIdentity.Set(&v)
}
// SetConsumerReportUserIdentityNil sets the value for ConsumerReportUserIdentity to be an explicit nil
func (o *UserCreateRequest) SetConsumerReportUserIdentityNil() {
	o.ConsumerReportUserIdentity.Set(nil)
}

// UnsetConsumerReportUserIdentity ensures that no value is present for ConsumerReportUserIdentity, not even an explicit nil
func (o *UserCreateRequest) UnsetConsumerReportUserIdentity() {
	o.ConsumerReportUserIdentity.Unset()
}

func (o UserCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if o.EndCustomer != nil {
		toSerialize["end_customer"] = o.EndCustomer
	}
	if o.ConsumerReportUserIdentity.IsSet() {
		toSerialize["consumer_report_user_identity"] = o.ConsumerReportUserIdentity.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserCreateRequest struct {
	value *UserCreateRequest
	isSet bool
}

func (v NullableUserCreateRequest) Get() *UserCreateRequest {
	return v.value
}

func (v *NullableUserCreateRequest) Set(val *UserCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreateRequest(val *UserCreateRequest) *NullableUserCreateRequest {
	return &NullableUserCreateRequest{value: val, isSet: true}
}

func (v NullableUserCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


