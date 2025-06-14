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

// ProcessorSyncUpdatesAvailableWebhook This webhook is only sent to [Plaid processor partners](https://plaid.com/docs/auth/partnerships/).  Fired when an Item's transactions change. This can be due to any event resulting in new changes, such as an initial 30-day transactions fetch upon the initialization of an Item with transactions, the backfill of historical transactions that occurs shortly after, or when changes are populated from a regularly-scheduled transactions update job. It is recommended to listen for the `SYNC_UPDATES_AVAILABLE` webhook when using the `/processor/transactions/sync` endpoint. Note that when using `/processor/transactions/sync` the older webhooks `INITIAL_UPDATE`, `HISTORICAL_UPDATE`, `DEFAULT_UPDATE`, and `TRANSACTIONS_REMOVED`, which are intended for use with `/processor/transactions/get`, will also continue to be sent in order to maintain backwards compatibility. It is not necessary to listen for and respond to those webhooks when using `/processor/transactions/sync`.  After receipt of this webhook, the new changes can be fetched for the Item from `/processor/transactions/sync`.  Note that to receive this webhook for an Item, `/processor/transactions/sync` must have been called at least once on that Item. This means that, unlike the `INITIAL_UPDATE` and `HISTORICAL_UPDATE` webhooks, it will not fire immediately upon Item creation. If `/transactions/sync` is called on an Item that was *not* initialized with Transactions, the webhook will fire twice: once the first 30 days of transactions data has been fetched, and a second time when all available historical transactions data has been fetched.  This webhook will typically not fire in the Sandbox environment, due to the lack of dynamic transactions data. To test this webhook in Sandbox, call `/sandbox/item/fire_webhook`.
type ProcessorSyncUpdatesAvailableWebhook struct {
	// `TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `SYNC_UPDATES_AVAILABLE`
	WebhookCode string `json:"webhook_code"`
	// The ID of the account.
	AccountId string `json:"account_id"`
	// Indicates if initial pull information is available.
	InitialUpdateComplete bool `json:"initial_update_complete"`
	// Indicates if historical pull information is available.
	HistoricalUpdateComplete bool `json:"historical_update_complete"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorSyncUpdatesAvailableWebhook ProcessorSyncUpdatesAvailableWebhook

// NewProcessorSyncUpdatesAvailableWebhook instantiates a new ProcessorSyncUpdatesAvailableWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorSyncUpdatesAvailableWebhook(webhookType string, webhookCode string, accountId string, initialUpdateComplete bool, historicalUpdateComplete bool, environment WebhookEnvironmentValues) *ProcessorSyncUpdatesAvailableWebhook {
	this := ProcessorSyncUpdatesAvailableWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.AccountId = accountId
	this.InitialUpdateComplete = initialUpdateComplete
	this.HistoricalUpdateComplete = historicalUpdateComplete
	this.Environment = environment
	return &this
}

// NewProcessorSyncUpdatesAvailableWebhookWithDefaults instantiates a new ProcessorSyncUpdatesAvailableWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorSyncUpdatesAvailableWebhookWithDefaults() *ProcessorSyncUpdatesAvailableWebhook {
	this := ProcessorSyncUpdatesAvailableWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *ProcessorSyncUpdatesAvailableWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *ProcessorSyncUpdatesAvailableWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *ProcessorSyncUpdatesAvailableWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *ProcessorSyncUpdatesAvailableWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *ProcessorSyncUpdatesAvailableWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *ProcessorSyncUpdatesAvailableWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetAccountId returns the AccountId field value
func (o *ProcessorSyncUpdatesAvailableWebhook) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ProcessorSyncUpdatesAvailableWebhook) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ProcessorSyncUpdatesAvailableWebhook) SetAccountId(v string) {
	o.AccountId = v
}

// GetInitialUpdateComplete returns the InitialUpdateComplete field value
func (o *ProcessorSyncUpdatesAvailableWebhook) GetInitialUpdateComplete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InitialUpdateComplete
}

// GetInitialUpdateCompleteOk returns a tuple with the InitialUpdateComplete field value
// and a boolean to check if the value has been set.
func (o *ProcessorSyncUpdatesAvailableWebhook) GetInitialUpdateCompleteOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InitialUpdateComplete, true
}

// SetInitialUpdateComplete sets field value
func (o *ProcessorSyncUpdatesAvailableWebhook) SetInitialUpdateComplete(v bool) {
	o.InitialUpdateComplete = v
}

// GetHistoricalUpdateComplete returns the HistoricalUpdateComplete field value
func (o *ProcessorSyncUpdatesAvailableWebhook) GetHistoricalUpdateComplete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HistoricalUpdateComplete
}

// GetHistoricalUpdateCompleteOk returns a tuple with the HistoricalUpdateComplete field value
// and a boolean to check if the value has been set.
func (o *ProcessorSyncUpdatesAvailableWebhook) GetHistoricalUpdateCompleteOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HistoricalUpdateComplete, true
}

// SetHistoricalUpdateComplete sets field value
func (o *ProcessorSyncUpdatesAvailableWebhook) SetHistoricalUpdateComplete(v bool) {
	o.HistoricalUpdateComplete = v
}

// GetEnvironment returns the Environment field value
func (o *ProcessorSyncUpdatesAvailableWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ProcessorSyncUpdatesAvailableWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ProcessorSyncUpdatesAvailableWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o ProcessorSyncUpdatesAvailableWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["initial_update_complete"] = o.InitialUpdateComplete
	}
	if true {
		toSerialize["historical_update_complete"] = o.HistoricalUpdateComplete
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorSyncUpdatesAvailableWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorSyncUpdatesAvailableWebhook := _ProcessorSyncUpdatesAvailableWebhook{}

	if err = json.Unmarshal(bytes, &varProcessorSyncUpdatesAvailableWebhook); err == nil {
		*o = ProcessorSyncUpdatesAvailableWebhook(varProcessorSyncUpdatesAvailableWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "initial_update_complete")
		delete(additionalProperties, "historical_update_complete")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorSyncUpdatesAvailableWebhook struct {
	value *ProcessorSyncUpdatesAvailableWebhook
	isSet bool
}

func (v NullableProcessorSyncUpdatesAvailableWebhook) Get() *ProcessorSyncUpdatesAvailableWebhook {
	return v.value
}

func (v *NullableProcessorSyncUpdatesAvailableWebhook) Set(val *ProcessorSyncUpdatesAvailableWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorSyncUpdatesAvailableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorSyncUpdatesAvailableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorSyncUpdatesAvailableWebhook(val *ProcessorSyncUpdatesAvailableWebhook) *NullableProcessorSyncUpdatesAvailableWebhook {
	return &NullableProcessorSyncUpdatesAvailableWebhook{value: val, isSet: true}
}

func (v NullableProcessorSyncUpdatesAvailableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorSyncUpdatesAvailableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


