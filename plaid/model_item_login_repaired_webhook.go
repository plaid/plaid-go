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

// ItemLoginRepairedWebhook Fired when an Item has exited the `ITEM_LOGIN_REQUIRED` state without the user having gone through the update mode flow in your app (this can happen if the user completed the update mode in a different app). If you have messaging that tells the user to complete the update mode flow, you should silence this messaging upon receiving the `LOGIN_REPAIRED` webhook.
type ItemLoginRepairedWebhook struct {
	// `ITEM`
	WebhookType string `json:"webhook_type"`
	// `LOGIN_REPAIRED`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _ItemLoginRepairedWebhook ItemLoginRepairedWebhook

// NewItemLoginRepairedWebhook instantiates a new ItemLoginRepairedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemLoginRepairedWebhook(webhookType string, webhookCode string, itemId string, environment WebhookEnvironmentValues) *ItemLoginRepairedWebhook {
	this := ItemLoginRepairedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.Environment = environment
	return &this
}

// NewItemLoginRepairedWebhookWithDefaults instantiates a new ItemLoginRepairedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemLoginRepairedWebhookWithDefaults() *ItemLoginRepairedWebhook {
	this := ItemLoginRepairedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *ItemLoginRepairedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *ItemLoginRepairedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *ItemLoginRepairedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *ItemLoginRepairedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *ItemLoginRepairedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *ItemLoginRepairedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *ItemLoginRepairedWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *ItemLoginRepairedWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *ItemLoginRepairedWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetEnvironment returns the Environment field value
func (o *ItemLoginRepairedWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ItemLoginRepairedWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ItemLoginRepairedWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o ItemLoginRepairedWebhook) MarshalJSON() ([]byte, error) {
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
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemLoginRepairedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varItemLoginRepairedWebhook := _ItemLoginRepairedWebhook{}

	if err = json.Unmarshal(bytes, &varItemLoginRepairedWebhook); err == nil {
		*o = ItemLoginRepairedWebhook(varItemLoginRepairedWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemLoginRepairedWebhook struct {
	value *ItemLoginRepairedWebhook
	isSet bool
}

func (v NullableItemLoginRepairedWebhook) Get() *ItemLoginRepairedWebhook {
	return v.value
}

func (v *NullableItemLoginRepairedWebhook) Set(val *ItemLoginRepairedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableItemLoginRepairedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableItemLoginRepairedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemLoginRepairedWebhook(val *ItemLoginRepairedWebhook) *NullableItemLoginRepairedWebhook {
	return &NullableItemLoginRepairedWebhook{value: val, isSet: true}
}

func (v NullableItemLoginRepairedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemLoginRepairedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


