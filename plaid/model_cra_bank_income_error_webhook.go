/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.503.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CraBankIncomeErrorWebhook Fired when a bank income report has failed to generate
type CraBankIncomeErrorWebhook struct {
	// `CRA_INCOME`
	WebhookType string `json:"webhook_type"`
	// `ERROR`
	WebhookCode string `json:"webhook_code"`
	// The `user_id` corresponding to the user the webhook has fired for.
	UserId string `json:"user_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _CraBankIncomeErrorWebhook CraBankIncomeErrorWebhook

// NewCraBankIncomeErrorWebhook instantiates a new CraBankIncomeErrorWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraBankIncomeErrorWebhook(webhookType string, webhookCode string, userId string, environment WebhookEnvironmentValues) *CraBankIncomeErrorWebhook {
	this := CraBankIncomeErrorWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.UserId = userId
	this.Environment = environment
	return &this
}

// NewCraBankIncomeErrorWebhookWithDefaults instantiates a new CraBankIncomeErrorWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraBankIncomeErrorWebhookWithDefaults() *CraBankIncomeErrorWebhook {
	this := CraBankIncomeErrorWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *CraBankIncomeErrorWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeErrorWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *CraBankIncomeErrorWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *CraBankIncomeErrorWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeErrorWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *CraBankIncomeErrorWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetUserId returns the UserId field value
func (o *CraBankIncomeErrorWebhook) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeErrorWebhook) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CraBankIncomeErrorWebhook) SetUserId(v string) {
	o.UserId = v
}

// GetEnvironment returns the Environment field value
func (o *CraBankIncomeErrorWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeErrorWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *CraBankIncomeErrorWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o CraBankIncomeErrorWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraBankIncomeErrorWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varCraBankIncomeErrorWebhook := _CraBankIncomeErrorWebhook{}

	if err = json.Unmarshal(bytes, &varCraBankIncomeErrorWebhook); err == nil {
		*o = CraBankIncomeErrorWebhook(varCraBankIncomeErrorWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraBankIncomeErrorWebhook struct {
	value *CraBankIncomeErrorWebhook
	isSet bool
}

func (v NullableCraBankIncomeErrorWebhook) Get() *CraBankIncomeErrorWebhook {
	return v.value
}

func (v *NullableCraBankIncomeErrorWebhook) Set(val *CraBankIncomeErrorWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableCraBankIncomeErrorWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableCraBankIncomeErrorWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraBankIncomeErrorWebhook(val *CraBankIncomeErrorWebhook) *NullableCraBankIncomeErrorWebhook {
	return &NullableCraBankIncomeErrorWebhook{value: val, isSet: true}
}

func (v NullableCraBankIncomeErrorWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraBankIncomeErrorWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

