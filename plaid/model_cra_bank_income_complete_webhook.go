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

// CraBankIncomeCompleteWebhook Fired when a bank income report has finished generating or failed to generate, triggered by calling `/cra/bank_income/get`.
type CraBankIncomeCompleteWebhook struct {
	// `CRA_INCOME`
	WebhookType string `json:"webhook_type"`
	// `BANK_INCOME_COMPLETE`
	WebhookCode string `json:"webhook_code"`
	// The `user_id` corresponding to the user the webhook has fired for.
	UserId string `json:"user_id"`
	Result *CraBankIncomeCompleteResult `json:"result,omitempty"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _CraBankIncomeCompleteWebhook CraBankIncomeCompleteWebhook

// NewCraBankIncomeCompleteWebhook instantiates a new CraBankIncomeCompleteWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraBankIncomeCompleteWebhook(webhookType string, webhookCode string, userId string, environment WebhookEnvironmentValues) *CraBankIncomeCompleteWebhook {
	this := CraBankIncomeCompleteWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.UserId = userId
	this.Environment = environment
	return &this
}

// NewCraBankIncomeCompleteWebhookWithDefaults instantiates a new CraBankIncomeCompleteWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraBankIncomeCompleteWebhookWithDefaults() *CraBankIncomeCompleteWebhook {
	this := CraBankIncomeCompleteWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *CraBankIncomeCompleteWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeCompleteWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *CraBankIncomeCompleteWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *CraBankIncomeCompleteWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeCompleteWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *CraBankIncomeCompleteWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetUserId returns the UserId field value
func (o *CraBankIncomeCompleteWebhook) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeCompleteWebhook) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CraBankIncomeCompleteWebhook) SetUserId(v string) {
	o.UserId = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CraBankIncomeCompleteWebhook) GetResult() CraBankIncomeCompleteResult {
	if o == nil || o.Result == nil {
		var ret CraBankIncomeCompleteResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraBankIncomeCompleteWebhook) GetResultOk() (*CraBankIncomeCompleteResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CraBankIncomeCompleteWebhook) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given CraBankIncomeCompleteResult and assigns it to the Result field.
func (o *CraBankIncomeCompleteWebhook) SetResult(v CraBankIncomeCompleteResult) {
	o.Result = &v
}

// GetEnvironment returns the Environment field value
func (o *CraBankIncomeCompleteWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *CraBankIncomeCompleteWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *CraBankIncomeCompleteWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o CraBankIncomeCompleteWebhook) MarshalJSON() ([]byte, error) {
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
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraBankIncomeCompleteWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varCraBankIncomeCompleteWebhook := _CraBankIncomeCompleteWebhook{}

	if err = json.Unmarshal(bytes, &varCraBankIncomeCompleteWebhook); err == nil {
		*o = CraBankIncomeCompleteWebhook(varCraBankIncomeCompleteWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "result")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraBankIncomeCompleteWebhook struct {
	value *CraBankIncomeCompleteWebhook
	isSet bool
}

func (v NullableCraBankIncomeCompleteWebhook) Get() *CraBankIncomeCompleteWebhook {
	return v.value
}

func (v *NullableCraBankIncomeCompleteWebhook) Set(val *CraBankIncomeCompleteWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableCraBankIncomeCompleteWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableCraBankIncomeCompleteWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraBankIncomeCompleteWebhook(val *CraBankIncomeCompleteWebhook) *NullableCraBankIncomeCompleteWebhook {
	return &NullableCraBankIncomeCompleteWebhook{value: val, isSet: true}
}

func (v NullableCraBankIncomeCompleteWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraBankIncomeCompleteWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


