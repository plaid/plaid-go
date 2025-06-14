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
	"time"
)

// WalletTransactionStatusUpdateWebhook Fired when the status of a wallet transaction has changed.
type WalletTransactionStatusUpdateWebhook struct {
	// `WALLET`
	WebhookType string `json:"webhook_type"`
	// `WALLET_TRANSACTION_STATUS_UPDATE`
	WebhookCode string `json:"webhook_code"`
	// The `transaction_id` for the wallet transaction being updated
	TransactionId string `json:"transaction_id"`
	// The `payment_id` associated with the transaction. This will be present in case of `REFUND` and `PIS_PAY_IN`.
	PaymentId NullableString `json:"payment_id,omitempty"`
	// The EMI (E-Money Institution) wallet that this payment is associated with. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.
	WalletId *string `json:"wallet_id,omitempty"`
	NewStatus WalletTransactionStatus `json:"new_status"`
	OldStatus WalletTransactionStatus `json:"old_status"`
	// The timestamp of the update, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format, e.g. `\"2017-09-14T14:42:19.350Z\"`
	Timestamp time.Time `json:"timestamp"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransactionStatusUpdateWebhook WalletTransactionStatusUpdateWebhook

// NewWalletTransactionStatusUpdateWebhook instantiates a new WalletTransactionStatusUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionStatusUpdateWebhook(webhookType string, webhookCode string, transactionId string, newStatus WalletTransactionStatus, oldStatus WalletTransactionStatus, timestamp time.Time, environment WebhookEnvironmentValues) *WalletTransactionStatusUpdateWebhook {
	this := WalletTransactionStatusUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.TransactionId = transactionId
	this.NewStatus = newStatus
	this.OldStatus = oldStatus
	this.Timestamp = timestamp
	this.Environment = environment
	return &this
}

// NewWalletTransactionStatusUpdateWebhookWithDefaults instantiates a new WalletTransactionStatusUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionStatusUpdateWebhookWithDefaults() *WalletTransactionStatusUpdateWebhook {
	this := WalletTransactionStatusUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *WalletTransactionStatusUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionStatusUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *WalletTransactionStatusUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *WalletTransactionStatusUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionStatusUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *WalletTransactionStatusUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetTransactionId returns the TransactionId field value
func (o *WalletTransactionStatusUpdateWebhook) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionStatusUpdateWebhook) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *WalletTransactionStatusUpdateWebhook) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletTransactionStatusUpdateWebhook) GetPaymentId() string {
	if o == nil || o.PaymentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletTransactionStatusUpdateWebhook) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// HasPaymentId returns a boolean if a field has been set.
func (o *WalletTransactionStatusUpdateWebhook) HasPaymentId() bool {
	if o != nil && o.PaymentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given NullableString and assigns it to the PaymentId field.
func (o *WalletTransactionStatusUpdateWebhook) SetPaymentId(v string) {
	o.PaymentId.Set(&v)
}
// SetPaymentIdNil sets the value for PaymentId to be an explicit nil
func (o *WalletTransactionStatusUpdateWebhook) SetPaymentIdNil() {
	o.PaymentId.Set(nil)
}

// UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
func (o *WalletTransactionStatusUpdateWebhook) UnsetPaymentId() {
	o.PaymentId.Unset()
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *WalletTransactionStatusUpdateWebhook) GetWalletId() string {
	if o == nil || o.WalletId == nil {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletTransactionStatusUpdateWebhook) GetWalletIdOk() (*string, bool) {
	if o == nil || o.WalletId == nil {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *WalletTransactionStatusUpdateWebhook) HasWalletId() bool {
	if o != nil && o.WalletId != nil {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *WalletTransactionStatusUpdateWebhook) SetWalletId(v string) {
	o.WalletId = &v
}

// GetNewStatus returns the NewStatus field value
func (o *WalletTransactionStatusUpdateWebhook) GetNewStatus() WalletTransactionStatus {
	if o == nil {
		var ret WalletTransactionStatus
		return ret
	}

	return o.NewStatus
}

// GetNewStatusOk returns a tuple with the NewStatus field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionStatusUpdateWebhook) GetNewStatusOk() (*WalletTransactionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewStatus, true
}

// SetNewStatus sets field value
func (o *WalletTransactionStatusUpdateWebhook) SetNewStatus(v WalletTransactionStatus) {
	o.NewStatus = v
}

// GetOldStatus returns the OldStatus field value
func (o *WalletTransactionStatusUpdateWebhook) GetOldStatus() WalletTransactionStatus {
	if o == nil {
		var ret WalletTransactionStatus
		return ret
	}

	return o.OldStatus
}

// GetOldStatusOk returns a tuple with the OldStatus field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionStatusUpdateWebhook) GetOldStatusOk() (*WalletTransactionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OldStatus, true
}

// SetOldStatus sets field value
func (o *WalletTransactionStatusUpdateWebhook) SetOldStatus(v WalletTransactionStatus) {
	o.OldStatus = v
}

// GetTimestamp returns the Timestamp field value
func (o *WalletTransactionStatusUpdateWebhook) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionStatusUpdateWebhook) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *WalletTransactionStatusUpdateWebhook) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetEnvironment returns the Environment field value
func (o *WalletTransactionStatusUpdateWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionStatusUpdateWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *WalletTransactionStatusUpdateWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o WalletTransactionStatusUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if o.PaymentId.IsSet() {
		toSerialize["payment_id"] = o.PaymentId.Get()
	}
	if o.WalletId != nil {
		toSerialize["wallet_id"] = o.WalletId
	}
	if true {
		toSerialize["new_status"] = o.NewStatus
	}
	if true {
		toSerialize["old_status"] = o.OldStatus
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletTransactionStatusUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransactionStatusUpdateWebhook := _WalletTransactionStatusUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varWalletTransactionStatusUpdateWebhook); err == nil {
		*o = WalletTransactionStatusUpdateWebhook(varWalletTransactionStatusUpdateWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "wallet_id")
		delete(additionalProperties, "new_status")
		delete(additionalProperties, "old_status")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransactionStatusUpdateWebhook struct {
	value *WalletTransactionStatusUpdateWebhook
	isSet bool
}

func (v NullableWalletTransactionStatusUpdateWebhook) Get() *WalletTransactionStatusUpdateWebhook {
	return v.value
}

func (v *NullableWalletTransactionStatusUpdateWebhook) Set(val *WalletTransactionStatusUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionStatusUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionStatusUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionStatusUpdateWebhook(val *WalletTransactionStatusUpdateWebhook) *NullableWalletTransactionStatusUpdateWebhook {
	return &NullableWalletTransactionStatusUpdateWebhook{value: val, isSet: true}
}

func (v NullableWalletTransactionStatusUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionStatusUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


