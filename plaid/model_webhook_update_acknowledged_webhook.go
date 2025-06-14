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

// WebhookUpdateAcknowledgedWebhook Fired when an Item's webhook is updated. This will be sent to the newly specified webhook.
type WebhookUpdateAcknowledgedWebhook struct {
	// `ITEM`
	WebhookType string `json:"webhook_type"`
	// `WEBHOOK_UPDATE_ACKNOWLEDGED`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	// The new webhook URL
	NewWebhookUrl string `json:"new_webhook_url"`
	Error NullablePlaidError `json:"error,omitempty"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _WebhookUpdateAcknowledgedWebhook WebhookUpdateAcknowledgedWebhook

// NewWebhookUpdateAcknowledgedWebhook instantiates a new WebhookUpdateAcknowledgedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookUpdateAcknowledgedWebhook(webhookType string, webhookCode string, itemId string, newWebhookUrl string, environment WebhookEnvironmentValues) *WebhookUpdateAcknowledgedWebhook {
	this := WebhookUpdateAcknowledgedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.NewWebhookUrl = newWebhookUrl
	this.Environment = environment
	return &this
}

// NewWebhookUpdateAcknowledgedWebhookWithDefaults instantiates a new WebhookUpdateAcknowledgedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookUpdateAcknowledgedWebhookWithDefaults() *WebhookUpdateAcknowledgedWebhook {
	this := WebhookUpdateAcknowledgedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *WebhookUpdateAcknowledgedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *WebhookUpdateAcknowledgedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *WebhookUpdateAcknowledgedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *WebhookUpdateAcknowledgedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *WebhookUpdateAcknowledgedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *WebhookUpdateAcknowledgedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *WebhookUpdateAcknowledgedWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *WebhookUpdateAcknowledgedWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *WebhookUpdateAcknowledgedWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetNewWebhookUrl returns the NewWebhookUrl field value
func (o *WebhookUpdateAcknowledgedWebhook) GetNewWebhookUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewWebhookUrl
}

// GetNewWebhookUrlOk returns a tuple with the NewWebhookUrl field value
// and a boolean to check if the value has been set.
func (o *WebhookUpdateAcknowledgedWebhook) GetNewWebhookUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewWebhookUrl, true
}

// SetNewWebhookUrl sets field value
func (o *WebhookUpdateAcknowledgedWebhook) SetNewWebhookUrl(v string) {
	o.NewWebhookUrl = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookUpdateAcknowledgedWebhook) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookUpdateAcknowledgedWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *WebhookUpdateAcknowledgedWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullablePlaidError and assigns it to the Error field.
func (o *WebhookUpdateAcknowledgedWebhook) SetError(v PlaidError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *WebhookUpdateAcknowledgedWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *WebhookUpdateAcknowledgedWebhook) UnsetError() {
	o.Error.Unset()
}

// GetEnvironment returns the Environment field value
func (o *WebhookUpdateAcknowledgedWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *WebhookUpdateAcknowledgedWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *WebhookUpdateAcknowledgedWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o WebhookUpdateAcknowledgedWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["new_webhook_url"] = o.NewWebhookUrl
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WebhookUpdateAcknowledgedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varWebhookUpdateAcknowledgedWebhook := _WebhookUpdateAcknowledgedWebhook{}

	if err = json.Unmarshal(bytes, &varWebhookUpdateAcknowledgedWebhook); err == nil {
		*o = WebhookUpdateAcknowledgedWebhook(varWebhookUpdateAcknowledgedWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "new_webhook_url")
		delete(additionalProperties, "error")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookUpdateAcknowledgedWebhook struct {
	value *WebhookUpdateAcknowledgedWebhook
	isSet bool
}

func (v NullableWebhookUpdateAcknowledgedWebhook) Get() *WebhookUpdateAcknowledgedWebhook {
	return v.value
}

func (v *NullableWebhookUpdateAcknowledgedWebhook) Set(val *WebhookUpdateAcknowledgedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookUpdateAcknowledgedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookUpdateAcknowledgedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookUpdateAcknowledgedWebhook(val *WebhookUpdateAcknowledgedWebhook) *NullableWebhookUpdateAcknowledgedWebhook {
	return &NullableWebhookUpdateAcknowledgedWebhook{value: val, isSet: true}
}

func (v NullableWebhookUpdateAcknowledgedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookUpdateAcknowledgedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


