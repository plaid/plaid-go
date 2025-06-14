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

// CraPartnerInsightsCompleteWebhook Fired when a partner insights report has finished generating and results are available
type CraPartnerInsightsCompleteWebhook struct {
	// `CRA_INSIGHTS`
	WebhookType string `json:"webhook_type"`
	// `PARTNER_INSIGHTS_COMPLETE`
	WebhookCode string `json:"webhook_code"`
	// The `user_id` corresponding to the user the webhook has fired for.
	UserId string `json:"user_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _CraPartnerInsightsCompleteWebhook CraPartnerInsightsCompleteWebhook

// NewCraPartnerInsightsCompleteWebhook instantiates a new CraPartnerInsightsCompleteWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraPartnerInsightsCompleteWebhook(webhookType string, webhookCode string, userId string, environment WebhookEnvironmentValues) *CraPartnerInsightsCompleteWebhook {
	this := CraPartnerInsightsCompleteWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.UserId = userId
	this.Environment = environment
	return &this
}

// NewCraPartnerInsightsCompleteWebhookWithDefaults instantiates a new CraPartnerInsightsCompleteWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraPartnerInsightsCompleteWebhookWithDefaults() *CraPartnerInsightsCompleteWebhook {
	this := CraPartnerInsightsCompleteWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *CraPartnerInsightsCompleteWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *CraPartnerInsightsCompleteWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *CraPartnerInsightsCompleteWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *CraPartnerInsightsCompleteWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *CraPartnerInsightsCompleteWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *CraPartnerInsightsCompleteWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetUserId returns the UserId field value
func (o *CraPartnerInsightsCompleteWebhook) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CraPartnerInsightsCompleteWebhook) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CraPartnerInsightsCompleteWebhook) SetUserId(v string) {
	o.UserId = v
}

// GetEnvironment returns the Environment field value
func (o *CraPartnerInsightsCompleteWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *CraPartnerInsightsCompleteWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *CraPartnerInsightsCompleteWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o CraPartnerInsightsCompleteWebhook) MarshalJSON() ([]byte, error) {
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

func (o *CraPartnerInsightsCompleteWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varCraPartnerInsightsCompleteWebhook := _CraPartnerInsightsCompleteWebhook{}

	if err = json.Unmarshal(bytes, &varCraPartnerInsightsCompleteWebhook); err == nil {
		*o = CraPartnerInsightsCompleteWebhook(varCraPartnerInsightsCompleteWebhook)
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

type NullableCraPartnerInsightsCompleteWebhook struct {
	value *CraPartnerInsightsCompleteWebhook
	isSet bool
}

func (v NullableCraPartnerInsightsCompleteWebhook) Get() *CraPartnerInsightsCompleteWebhook {
	return v.value
}

func (v *NullableCraPartnerInsightsCompleteWebhook) Set(val *CraPartnerInsightsCompleteWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableCraPartnerInsightsCompleteWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableCraPartnerInsightsCompleteWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraPartnerInsightsCompleteWebhook(val *CraPartnerInsightsCompleteWebhook) *NullableCraPartnerInsightsCompleteWebhook {
	return &NullableCraPartnerInsightsCompleteWebhook{value: val, isSet: true}
}

func (v NullableCraPartnerInsightsCompleteWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraPartnerInsightsCompleteWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


