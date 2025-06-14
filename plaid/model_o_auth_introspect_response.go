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

// OAuthIntrospectResponse OAuth token introspect response
type OAuthIntrospectResponse struct {
	// Boolean indicator of whether or not the presented token is currently active. A `true` value indicates that the token has been issued, has not been revoked, and is within the time window of validity.
	Active bool `json:"active"`
	// A JSON string containing a space-separated list of scopes associated with this token, in the format described in [https://datatracker.ietf.org/doc/html/rfc6749#section-3.3](https://datatracker.ietf.org/doc/html/rfc6749#section-3.3). Currently accepted values are:  - `user:read` allows reading user data. - `user:write` allows writing user data. - `exchange` allows exchanging a token using the `urn:plaid:params:oauth:user-token` grant type. - `mcp:dashboard` allows access to the MCP dashboard server.
	Scope *string `json:"scope,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Expiration time as UNIX timestamp since January 1 1970 UTC
	Exp *int32 `json:"exp,omitempty"`
	// Issued at time as UNIX timestamp since January 1 1970 UTC
	Iat *int32 `json:"iat,omitempty"`
	// Subject of the token
	Sub *string `json:"sub,omitempty"`
	// Audience of the token
	Aud *string `json:"aud,omitempty"`
	// Issuer of the token
	Iss *string `json:"iss,omitempty"`
	// Type of the token
	TokenType *string `json:"token_type,omitempty"`
	// User ID of the token
	UserId *string `json:"user_id,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _OAuthIntrospectResponse OAuthIntrospectResponse

// NewOAuthIntrospectResponse instantiates a new OAuthIntrospectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthIntrospectResponse(active bool, requestId string) *OAuthIntrospectResponse {
	this := OAuthIntrospectResponse{}
	this.Active = active
	this.RequestId = requestId
	return &this
}

// NewOAuthIntrospectResponseWithDefaults instantiates a new OAuthIntrospectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthIntrospectResponseWithDefaults() *OAuthIntrospectResponse {
	this := OAuthIntrospectResponse{}
	return &this
}

// GetActive returns the Active field value
func (o *OAuthIntrospectResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *OAuthIntrospectResponse) SetActive(v bool) {
	o.Active = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OAuthIntrospectResponse) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OAuthIntrospectResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *OAuthIntrospectResponse) SetScope(v string) {
	o.Scope = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *OAuthIntrospectResponse) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *OAuthIntrospectResponse) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *OAuthIntrospectResponse) SetClientId(v string) {
	o.ClientId = &v
}

// GetExp returns the Exp field value if set, zero value otherwise.
func (o *OAuthIntrospectResponse) GetExp() int32 {
	if o == nil || o.Exp == nil {
		var ret int32
		return ret
	}
	return *o.Exp
}

// GetExpOk returns a tuple with the Exp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetExpOk() (*int32, bool) {
	if o == nil || o.Exp == nil {
		return nil, false
	}
	return o.Exp, true
}

// HasExp returns a boolean if a field has been set.
func (o *OAuthIntrospectResponse) HasExp() bool {
	if o != nil && o.Exp != nil {
		return true
	}

	return false
}

// SetExp gets a reference to the given int32 and assigns it to the Exp field.
func (o *OAuthIntrospectResponse) SetExp(v int32) {
	o.Exp = &v
}

// GetIat returns the Iat field value if set, zero value otherwise.
func (o *OAuthIntrospectResponse) GetIat() int32 {
	if o == nil || o.Iat == nil {
		var ret int32
		return ret
	}
	return *o.Iat
}

// GetIatOk returns a tuple with the Iat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetIatOk() (*int32, bool) {
	if o == nil || o.Iat == nil {
		return nil, false
	}
	return o.Iat, true
}

// HasIat returns a boolean if a field has been set.
func (o *OAuthIntrospectResponse) HasIat() bool {
	if o != nil && o.Iat != nil {
		return true
	}

	return false
}

// SetIat gets a reference to the given int32 and assigns it to the Iat field.
func (o *OAuthIntrospectResponse) SetIat(v int32) {
	o.Iat = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *OAuthIntrospectResponse) GetSub() string {
	if o == nil || o.Sub == nil {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetSubOk() (*string, bool) {
	if o == nil || o.Sub == nil {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *OAuthIntrospectResponse) HasSub() bool {
	if o != nil && o.Sub != nil {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *OAuthIntrospectResponse) SetSub(v string) {
	o.Sub = &v
}

// GetAud returns the Aud field value if set, zero value otherwise.
func (o *OAuthIntrospectResponse) GetAud() string {
	if o == nil || o.Aud == nil {
		var ret string
		return ret
	}
	return *o.Aud
}

// GetAudOk returns a tuple with the Aud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetAudOk() (*string, bool) {
	if o == nil || o.Aud == nil {
		return nil, false
	}
	return o.Aud, true
}

// HasAud returns a boolean if a field has been set.
func (o *OAuthIntrospectResponse) HasAud() bool {
	if o != nil && o.Aud != nil {
		return true
	}

	return false
}

// SetAud gets a reference to the given string and assigns it to the Aud field.
func (o *OAuthIntrospectResponse) SetAud(v string) {
	o.Aud = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *OAuthIntrospectResponse) GetIss() string {
	if o == nil || o.Iss == nil {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetIssOk() (*string, bool) {
	if o == nil || o.Iss == nil {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *OAuthIntrospectResponse) HasIss() bool {
	if o != nil && o.Iss != nil {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *OAuthIntrospectResponse) SetIss(v string) {
	o.Iss = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *OAuthIntrospectResponse) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *OAuthIntrospectResponse) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *OAuthIntrospectResponse) SetTokenType(v string) {
	o.TokenType = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OAuthIntrospectResponse) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OAuthIntrospectResponse) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *OAuthIntrospectResponse) SetUserId(v string) {
	o.UserId = &v
}

// GetRequestId returns the RequestId field value
func (o *OAuthIntrospectResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *OAuthIntrospectResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *OAuthIntrospectResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o OAuthIntrospectResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["active"] = o.Active
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Exp != nil {
		toSerialize["exp"] = o.Exp
	}
	if o.Iat != nil {
		toSerialize["iat"] = o.Iat
	}
	if o.Sub != nil {
		toSerialize["sub"] = o.Sub
	}
	if o.Aud != nil {
		toSerialize["aud"] = o.Aud
	}
	if o.Iss != nil {
		toSerialize["iss"] = o.Iss
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuthIntrospectResponse) UnmarshalJSON(bytes []byte) (err error) {
	varOAuthIntrospectResponse := _OAuthIntrospectResponse{}

	if err = json.Unmarshal(bytes, &varOAuthIntrospectResponse); err == nil {
		*o = OAuthIntrospectResponse(varOAuthIntrospectResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "exp")
		delete(additionalProperties, "iat")
		delete(additionalProperties, "sub")
		delete(additionalProperties, "aud")
		delete(additionalProperties, "iss")
		delete(additionalProperties, "token_type")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuthIntrospectResponse struct {
	value *OAuthIntrospectResponse
	isSet bool
}

func (v NullableOAuthIntrospectResponse) Get() *OAuthIntrospectResponse {
	return v.value
}

func (v *NullableOAuthIntrospectResponse) Set(val *OAuthIntrospectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthIntrospectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthIntrospectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthIntrospectResponse(val *OAuthIntrospectResponse) *NullableOAuthIntrospectResponse {
	return &NullableOAuthIntrospectResponse{value: val, isSet: true}
}

func (v NullableOAuthIntrospectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthIntrospectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


