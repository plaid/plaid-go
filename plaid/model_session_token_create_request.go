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

// SessionTokenCreateRequest SessionTokenCreateRequest defines the request schema for `/session/token/create`
type SessionTokenCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The id of a template defined in Plaid Dashboard
	TemplateId string `json:"template_id"`
	User SessionTokenCreateRequestUser `json:"user"`
	// A URI indicating the destination where a user should be forwarded after completing the Link flow; used to support OAuth authentication flows when launching Link in the browser or another app. The `redirect_uri` should not contain any query parameters. When used in Production, must be an https URI. To specify any subdomain, use `*` as a wildcard character, e.g. `https://_*.example.com/oauth.html`. Note that any redirect URI must also be added to the Allowed redirect URIs list in the [developer dashboard](https://dashboard.plaid.com/team/api). If initializing on Android, `android_package_name` must be specified instead and `redirect_uri` should be left blank.
	RedirectUri *string `json:"redirect_uri,omitempty"`
	// The name of your app's Android package. Required if using the session token to initialize Layer on Android. Any package name specified here must also be added to the Allowed Android package names setting on the [developer dashboard](https://dashboard.plaid.com/team/api). When creating a session token for initializing Layer on other platforms, `android_package_name` must be left blank and `redirect_uri` should be used instead.
	AndroidPackageName *string `json:"android_package_name,omitempty"`
	// The destination URL to which any webhooks should be sent. If you use the same webhook listener for all Sandbox or all Production activity, set this value in the Layer template editor in the Dashboard instead. Only provide a value in this field if you need to use multiple webhook URLs per environment (an uncommon use case). If provided, a value in this field will take priority over webhook values set in the Layer template editor.
	Webhook *string `json:"webhook,omitempty"`
}

// NewSessionTokenCreateRequest instantiates a new SessionTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionTokenCreateRequest(templateId string, user SessionTokenCreateRequestUser) *SessionTokenCreateRequest {
	this := SessionTokenCreateRequest{}
	this.TemplateId = templateId
	this.User = user
	return &this
}

// NewSessionTokenCreateRequestWithDefaults instantiates a new SessionTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionTokenCreateRequestWithDefaults() *SessionTokenCreateRequest {
	this := SessionTokenCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SessionTokenCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionTokenCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SessionTokenCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SessionTokenCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SessionTokenCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionTokenCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SessionTokenCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SessionTokenCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetTemplateId returns the TemplateId field value
func (o *SessionTokenCreateRequest) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *SessionTokenCreateRequest) GetTemplateIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *SessionTokenCreateRequest) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetUser returns the User field value
func (o *SessionTokenCreateRequest) GetUser() SessionTokenCreateRequestUser {
	if o == nil {
		var ret SessionTokenCreateRequestUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *SessionTokenCreateRequest) GetUserOk() (*SessionTokenCreateRequestUser, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *SessionTokenCreateRequest) SetUser(v SessionTokenCreateRequestUser) {
	o.User = v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *SessionTokenCreateRequest) GetRedirectUri() string {
	if o == nil || o.RedirectUri == nil {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionTokenCreateRequest) GetRedirectUriOk() (*string, bool) {
	if o == nil || o.RedirectUri == nil {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *SessionTokenCreateRequest) HasRedirectUri() bool {
	if o != nil && o.RedirectUri != nil {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *SessionTokenCreateRequest) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetAndroidPackageName returns the AndroidPackageName field value if set, zero value otherwise.
func (o *SessionTokenCreateRequest) GetAndroidPackageName() string {
	if o == nil || o.AndroidPackageName == nil {
		var ret string
		return ret
	}
	return *o.AndroidPackageName
}

// GetAndroidPackageNameOk returns a tuple with the AndroidPackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionTokenCreateRequest) GetAndroidPackageNameOk() (*string, bool) {
	if o == nil || o.AndroidPackageName == nil {
		return nil, false
	}
	return o.AndroidPackageName, true
}

// HasAndroidPackageName returns a boolean if a field has been set.
func (o *SessionTokenCreateRequest) HasAndroidPackageName() bool {
	if o != nil && o.AndroidPackageName != nil {
		return true
	}

	return false
}

// SetAndroidPackageName gets a reference to the given string and assigns it to the AndroidPackageName field.
func (o *SessionTokenCreateRequest) SetAndroidPackageName(v string) {
	o.AndroidPackageName = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *SessionTokenCreateRequest) GetWebhook() string {
	if o == nil || o.Webhook == nil {
		var ret string
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionTokenCreateRequest) GetWebhookOk() (*string, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *SessionTokenCreateRequest) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given string and assigns it to the Webhook field.
func (o *SessionTokenCreateRequest) SetWebhook(v string) {
	o.Webhook = &v
}

func (o SessionTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["template_id"] = o.TemplateId
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.RedirectUri != nil {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if o.AndroidPackageName != nil {
		toSerialize["android_package_name"] = o.AndroidPackageName
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableSessionTokenCreateRequest struct {
	value *SessionTokenCreateRequest
	isSet bool
}

func (v NullableSessionTokenCreateRequest) Get() *SessionTokenCreateRequest {
	return v.value
}

func (v *NullableSessionTokenCreateRequest) Set(val *SessionTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionTokenCreateRequest(val *SessionTokenCreateRequest) *NullableSessionTokenCreateRequest {
	return &NullableSessionTokenCreateRequest{value: val, isSet: true}
}

func (v NullableSessionTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


