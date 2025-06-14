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

// PendingDisconnectWebhook Fired when an Item is expected to be disconnected. The webhook will currently be fired 7 days before the existing Item is scheduled for disconnection. This can be resolved by having the user go through Link’s [update mode](http://plaid.com/docs/link/update-mode). Currently, this webhook is fired only for US or Canadian institutions; in the UK or EU, you should continue to listed for the [`PENDING_EXPIRATION`](https://plaid.com/docs/api/items/#pending_expiration) webhook instead.
type PendingDisconnectWebhook struct {
	// `ITEM`
	WebhookType string `json:"webhook_type"`
	// `PENDING_DISCONNECT`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	Reason PendingDisconnectWebhookReason `json:"reason"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _PendingDisconnectWebhook PendingDisconnectWebhook

// NewPendingDisconnectWebhook instantiates a new PendingDisconnectWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingDisconnectWebhook(webhookType string, webhookCode string, itemId string, reason PendingDisconnectWebhookReason, environment WebhookEnvironmentValues) *PendingDisconnectWebhook {
	this := PendingDisconnectWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.Reason = reason
	this.Environment = environment
	return &this
}

// NewPendingDisconnectWebhookWithDefaults instantiates a new PendingDisconnectWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingDisconnectWebhookWithDefaults() *PendingDisconnectWebhook {
	this := PendingDisconnectWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *PendingDisconnectWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *PendingDisconnectWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *PendingDisconnectWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *PendingDisconnectWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *PendingDisconnectWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *PendingDisconnectWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *PendingDisconnectWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *PendingDisconnectWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *PendingDisconnectWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetReason returns the Reason field value
func (o *PendingDisconnectWebhook) GetReason() PendingDisconnectWebhookReason {
	if o == nil {
		var ret PendingDisconnectWebhookReason
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PendingDisconnectWebhook) GetReasonOk() (*PendingDisconnectWebhookReason, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PendingDisconnectWebhook) SetReason(v PendingDisconnectWebhookReason) {
	o.Reason = v
}

// GetEnvironment returns the Environment field value
func (o *PendingDisconnectWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *PendingDisconnectWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *PendingDisconnectWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o PendingDisconnectWebhook) MarshalJSON() ([]byte, error) {
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
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PendingDisconnectWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varPendingDisconnectWebhook := _PendingDisconnectWebhook{}

	if err = json.Unmarshal(bytes, &varPendingDisconnectWebhook); err == nil {
		*o = PendingDisconnectWebhook(varPendingDisconnectWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePendingDisconnectWebhook struct {
	value *PendingDisconnectWebhook
	isSet bool
}

func (v NullablePendingDisconnectWebhook) Get() *PendingDisconnectWebhook {
	return v.value
}

func (v *NullablePendingDisconnectWebhook) Set(val *PendingDisconnectWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingDisconnectWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingDisconnectWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingDisconnectWebhook(val *PendingDisconnectWebhook) *NullablePendingDisconnectWebhook {
	return &NullablePendingDisconnectWebhook{value: val, isSet: true}
}

func (v NullablePendingDisconnectWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingDisconnectWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


