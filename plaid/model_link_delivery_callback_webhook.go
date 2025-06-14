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

// LinkDeliveryCallbackWebhook Webhook containing metadata proxied over from Link callback e.g `onEvent`, `onExit`, `onSuccess`.
type LinkDeliveryCallbackWebhook struct {
	// `LINK_DELIVERY`
	WebhookType string `json:"webhook_type"`
	// `LINK_CALLBACK`
	WebhookCode string `json:"webhook_code"`
	// The ID of the Hosted Link session.
	LinkDeliverySessionId string `json:"link_delivery_session_id"`
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
	Timestamp string `json:"timestamp"`
	Error NullablePlaidError `json:"error,omitempty"`
	LinkCallbackMetadata LinkCallbackMetadata `json:"link_callback_metadata"`
	AdditionalProperties map[string]interface{}
}

type _LinkDeliveryCallbackWebhook LinkDeliveryCallbackWebhook

// NewLinkDeliveryCallbackWebhook instantiates a new LinkDeliveryCallbackWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDeliveryCallbackWebhook(webhookType string, webhookCode string, linkDeliverySessionId string, timestamp string, linkCallbackMetadata LinkCallbackMetadata) *LinkDeliveryCallbackWebhook {
	this := LinkDeliveryCallbackWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.LinkDeliverySessionId = linkDeliverySessionId
	this.Timestamp = timestamp
	this.LinkCallbackMetadata = linkCallbackMetadata
	return &this
}

// NewLinkDeliveryCallbackWebhookWithDefaults instantiates a new LinkDeliveryCallbackWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDeliveryCallbackWebhookWithDefaults() *LinkDeliveryCallbackWebhook {
	this := LinkDeliveryCallbackWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *LinkDeliveryCallbackWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCallbackWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *LinkDeliveryCallbackWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *LinkDeliveryCallbackWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCallbackWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *LinkDeliveryCallbackWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetLinkDeliverySessionId returns the LinkDeliverySessionId field value
func (o *LinkDeliveryCallbackWebhook) GetLinkDeliverySessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkDeliverySessionId
}

// GetLinkDeliverySessionIdOk returns a tuple with the LinkDeliverySessionId field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCallbackWebhook) GetLinkDeliverySessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkDeliverySessionId, true
}

// SetLinkDeliverySessionId sets field value
func (o *LinkDeliveryCallbackWebhook) SetLinkDeliverySessionId(v string) {
	o.LinkDeliverySessionId = v
}

// GetTimestamp returns the Timestamp field value
func (o *LinkDeliveryCallbackWebhook) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCallbackWebhook) GetTimestampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *LinkDeliveryCallbackWebhook) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkDeliveryCallbackWebhook) GetError() PlaidError {
	if o == nil || o.Error.Get() == nil {
		var ret PlaidError
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkDeliveryCallbackWebhook) GetErrorOk() (*PlaidError, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *LinkDeliveryCallbackWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullablePlaidError and assigns it to the Error field.
func (o *LinkDeliveryCallbackWebhook) SetError(v PlaidError) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *LinkDeliveryCallbackWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *LinkDeliveryCallbackWebhook) UnsetError() {
	o.Error.Unset()
}

// GetLinkCallbackMetadata returns the LinkCallbackMetadata field value
func (o *LinkDeliveryCallbackWebhook) GetLinkCallbackMetadata() LinkCallbackMetadata {
	if o == nil {
		var ret LinkCallbackMetadata
		return ret
	}

	return o.LinkCallbackMetadata
}

// GetLinkCallbackMetadataOk returns a tuple with the LinkCallbackMetadata field value
// and a boolean to check if the value has been set.
func (o *LinkDeliveryCallbackWebhook) GetLinkCallbackMetadataOk() (*LinkCallbackMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkCallbackMetadata, true
}

// SetLinkCallbackMetadata sets field value
func (o *LinkDeliveryCallbackWebhook) SetLinkCallbackMetadata(v LinkCallbackMetadata) {
	o.LinkCallbackMetadata = v
}

func (o LinkDeliveryCallbackWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["link_delivery_session_id"] = o.LinkDeliverySessionId
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["link_callback_metadata"] = o.LinkCallbackMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkDeliveryCallbackWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varLinkDeliveryCallbackWebhook := _LinkDeliveryCallbackWebhook{}

	if err = json.Unmarshal(bytes, &varLinkDeliveryCallbackWebhook); err == nil {
		*o = LinkDeliveryCallbackWebhook(varLinkDeliveryCallbackWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "link_delivery_session_id")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "error")
		delete(additionalProperties, "link_callback_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkDeliveryCallbackWebhook struct {
	value *LinkDeliveryCallbackWebhook
	isSet bool
}

func (v NullableLinkDeliveryCallbackWebhook) Get() *LinkDeliveryCallbackWebhook {
	return v.value
}

func (v *NullableLinkDeliveryCallbackWebhook) Set(val *LinkDeliveryCallbackWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryCallbackWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryCallbackWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryCallbackWebhook(val *LinkDeliveryCallbackWebhook) *NullableLinkDeliveryCallbackWebhook {
	return &NullableLinkDeliveryCallbackWebhook{value: val, isSet: true}
}

func (v NullableLinkDeliveryCallbackWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryCallbackWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


