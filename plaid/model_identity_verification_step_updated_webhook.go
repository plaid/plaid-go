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

// IdentityVerificationStepUpdatedWebhook Fired when an end user has completed a step of the Identity Verification process.
type IdentityVerificationStepUpdatedWebhook struct {
	// `IDENTITY_VERIFICATION`
	WebhookType string `json:"webhook_type"`
	// `STEP_UPDATED`
	WebhookCode string `json:"webhook_code"`
	// The ID of the associated Identity Verification attempt.
	IdentityVerificationId string `json:"identity_verification_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _IdentityVerificationStepUpdatedWebhook IdentityVerificationStepUpdatedWebhook

// NewIdentityVerificationStepUpdatedWebhook instantiates a new IdentityVerificationStepUpdatedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationStepUpdatedWebhook(webhookType string, webhookCode string, identityVerificationId string, environment WebhookEnvironmentValues) *IdentityVerificationStepUpdatedWebhook {
	this := IdentityVerificationStepUpdatedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.IdentityVerificationId = identityVerificationId
	this.Environment = environment
	return &this
}

// NewIdentityVerificationStepUpdatedWebhookWithDefaults instantiates a new IdentityVerificationStepUpdatedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationStepUpdatedWebhookWithDefaults() *IdentityVerificationStepUpdatedWebhook {
	this := IdentityVerificationStepUpdatedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *IdentityVerificationStepUpdatedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepUpdatedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *IdentityVerificationStepUpdatedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *IdentityVerificationStepUpdatedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepUpdatedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *IdentityVerificationStepUpdatedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetIdentityVerificationId returns the IdentityVerificationId field value
func (o *IdentityVerificationStepUpdatedWebhook) GetIdentityVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityVerificationId
}

// GetIdentityVerificationIdOk returns a tuple with the IdentityVerificationId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepUpdatedWebhook) GetIdentityVerificationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IdentityVerificationId, true
}

// SetIdentityVerificationId sets field value
func (o *IdentityVerificationStepUpdatedWebhook) SetIdentityVerificationId(v string) {
	o.IdentityVerificationId = v
}

// GetEnvironment returns the Environment field value
func (o *IdentityVerificationStepUpdatedWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationStepUpdatedWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *IdentityVerificationStepUpdatedWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o IdentityVerificationStepUpdatedWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["identity_verification_id"] = o.IdentityVerificationId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityVerificationStepUpdatedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityVerificationStepUpdatedWebhook := _IdentityVerificationStepUpdatedWebhook{}

	if err = json.Unmarshal(bytes, &varIdentityVerificationStepUpdatedWebhook); err == nil {
		*o = IdentityVerificationStepUpdatedWebhook(varIdentityVerificationStepUpdatedWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "identity_verification_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityVerificationStepUpdatedWebhook struct {
	value *IdentityVerificationStepUpdatedWebhook
	isSet bool
}

func (v NullableIdentityVerificationStepUpdatedWebhook) Get() *IdentityVerificationStepUpdatedWebhook {
	return v.value
}

func (v *NullableIdentityVerificationStepUpdatedWebhook) Set(val *IdentityVerificationStepUpdatedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationStepUpdatedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationStepUpdatedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationStepUpdatedWebhook(val *IdentityVerificationStepUpdatedWebhook) *NullableIdentityVerificationStepUpdatedWebhook {
	return &NullableIdentityVerificationStepUpdatedWebhook{value: val, isSet: true}
}

func (v NullableIdentityVerificationStepUpdatedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationStepUpdatedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


