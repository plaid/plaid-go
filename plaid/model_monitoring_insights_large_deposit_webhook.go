/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.620.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// MonitoringInsightsLargeDepositWebhook For each user enabled for Cash Flow Updates, this webhook will fire when an update detects a large deposit. Upon receiving the webhook, call `/cra/monitoring_insights/get` to retrieve the updated insights.
type MonitoringInsightsLargeDepositWebhook struct {
	// `CRA_MONITORING`
	WebhookType string `json:"webhook_type"`
	// `LARGE_DEPOSIT_DETECTED`
	WebhookCode string `json:"webhook_code"`
	Status MonitoringInsightsStatus `json:"status"`
	// The `user_id` that the report is associated with
	UserId string `json:"user_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _MonitoringInsightsLargeDepositWebhook MonitoringInsightsLargeDepositWebhook

// NewMonitoringInsightsLargeDepositWebhook instantiates a new MonitoringInsightsLargeDepositWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringInsightsLargeDepositWebhook(webhookType string, webhookCode string, status MonitoringInsightsStatus, userId string, environment WebhookEnvironmentValues) *MonitoringInsightsLargeDepositWebhook {
	this := MonitoringInsightsLargeDepositWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.Status = status
	this.UserId = userId
	this.Environment = environment
	return &this
}

// NewMonitoringInsightsLargeDepositWebhookWithDefaults instantiates a new MonitoringInsightsLargeDepositWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringInsightsLargeDepositWebhookWithDefaults() *MonitoringInsightsLargeDepositWebhook {
	this := MonitoringInsightsLargeDepositWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *MonitoringInsightsLargeDepositWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *MonitoringInsightsLargeDepositWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *MonitoringInsightsLargeDepositWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *MonitoringInsightsLargeDepositWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *MonitoringInsightsLargeDepositWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *MonitoringInsightsLargeDepositWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetStatus returns the Status field value
func (o *MonitoringInsightsLargeDepositWebhook) GetStatus() MonitoringInsightsStatus {
	if o == nil {
		var ret MonitoringInsightsStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MonitoringInsightsLargeDepositWebhook) GetStatusOk() (*MonitoringInsightsStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MonitoringInsightsLargeDepositWebhook) SetStatus(v MonitoringInsightsStatus) {
	o.Status = v
}

// GetUserId returns the UserId field value
func (o *MonitoringInsightsLargeDepositWebhook) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *MonitoringInsightsLargeDepositWebhook) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *MonitoringInsightsLargeDepositWebhook) SetUserId(v string) {
	o.UserId = v
}

// GetEnvironment returns the Environment field value
func (o *MonitoringInsightsLargeDepositWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *MonitoringInsightsLargeDepositWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *MonitoringInsightsLargeDepositWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o MonitoringInsightsLargeDepositWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["status"] = o.Status
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

func (o *MonitoringInsightsLargeDepositWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varMonitoringInsightsLargeDepositWebhook := _MonitoringInsightsLargeDepositWebhook{}

	if err = json.Unmarshal(bytes, &varMonitoringInsightsLargeDepositWebhook); err == nil {
		*o = MonitoringInsightsLargeDepositWebhook(varMonitoringInsightsLargeDepositWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "status")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMonitoringInsightsLargeDepositWebhook struct {
	value *MonitoringInsightsLargeDepositWebhook
	isSet bool
}

func (v NullableMonitoringInsightsLargeDepositWebhook) Get() *MonitoringInsightsLargeDepositWebhook {
	return v.value
}

func (v *NullableMonitoringInsightsLargeDepositWebhook) Set(val *MonitoringInsightsLargeDepositWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringInsightsLargeDepositWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringInsightsLargeDepositWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringInsightsLargeDepositWebhook(val *MonitoringInsightsLargeDepositWebhook) *NullableMonitoringInsightsLargeDepositWebhook {
	return &NullableMonitoringInsightsLargeDepositWebhook{value: val, isSet: true}
}

func (v NullableMonitoringInsightsLargeDepositWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringInsightsLargeDepositWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


