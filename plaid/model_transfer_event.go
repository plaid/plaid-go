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

// TransferEvent Represents an event in the Transfers API.
type TransferEvent struct {
	// Plaid’s unique identifier for this event. IDs are sequential unsigned 64-bit integers.
	EventId int32 `json:"event_id"`
	// The datetime when this event occurred. This will be of the form `2006-01-02T15:04:05Z`.
	Timestamp time.Time `json:"timestamp"`
	EventType TransferEventType `json:"event_type"`
	// The account ID associated with the transfer. This field is omitted for Plaid Ledger Sweep events.
	AccountId *string `json:"account_id,omitempty"`
	// The id of the associated funding account, available in the Plaid Dashboard. If present, this indicates which of your business checking accounts will be credited or debited.
	FundingAccountId NullableString `json:"funding_account_id"`
	// Plaid’s unique identifier for a Plaid Ledger Balance.
	LedgerId NullableString `json:"ledger_id,omitempty"`
	// Plaid's unique identifier for a transfer. This field is an empty string for Plaid Ledger Sweep events.
	TransferId string `json:"transfer_id"`
	// The ID of the origination account that this balance belongs to.
	OriginationAccountId NullableString `json:"origination_account_id"`
	TransferType *OmittableTransferType `json:"transfer_type,omitempty"`
	// The amount of the transfer (decimal string with two digits of precision e.g. \"10.00\"). This field is omitted for Plaid Ledger Sweep events.
	TransferAmount *string `json:"transfer_amount,omitempty"`
	FailureReason NullableTransferFailure `json:"failure_reason"`
	// Plaid’s unique identifier for a sweep.
	SweepId NullableString `json:"sweep_id"`
	// A signed amount of how much was `swept` or `return_swept` for this transfer (decimal string with two digits of precision e.g. \"-5.50\").
	SweepAmount NullableString `json:"sweep_amount"`
	// Plaid’s unique identifier for a refund. A non-null value indicates the event is for the associated refund of the transfer.
	RefundId NullableString `json:"refund_id"`
	// The Plaid client ID that is the originator of the transfer that this event applies to. Only present if the transfer was created on behalf of another client as a third-party sender (TPS).
	OriginatorClientId NullableString `json:"originator_client_id"`
	// The `id` returned by the /transfer/intent/create endpoint, for transfers created via Transfer UI. For transfers not created by Transfer UI, the value is `null`. This will currently only be populated for RfP transfers.
	IntentId NullableString `json:"intent_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferEvent TransferEvent

// NewTransferEvent instantiates a new TransferEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferEvent(eventId int32, timestamp time.Time, eventType TransferEventType, fundingAccountId NullableString, transferId string, originationAccountId NullableString, failureReason NullableTransferFailure, sweepId NullableString, sweepAmount NullableString, refundId NullableString, originatorClientId NullableString) *TransferEvent {
	this := TransferEvent{}
	this.EventId = eventId
	this.Timestamp = timestamp
	this.EventType = eventType
	this.FundingAccountId = fundingAccountId
	this.TransferId = transferId
	this.OriginationAccountId = originationAccountId
	this.FailureReason = failureReason
	this.SweepId = sweepId
	this.SweepAmount = sweepAmount
	this.RefundId = refundId
	this.OriginatorClientId = originatorClientId
	return &this
}

// NewTransferEventWithDefaults instantiates a new TransferEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferEventWithDefaults() *TransferEvent {
	this := TransferEvent{}
	return &this
}

// GetEventId returns the EventId field value
func (o *TransferEvent) GetEventId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetEventIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *TransferEvent) SetEventId(v int32) {
	o.EventId = v
}

// GetTimestamp returns the Timestamp field value
func (o *TransferEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *TransferEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetEventType returns the EventType field value
func (o *TransferEvent) GetEventType() TransferEventType {
	if o == nil {
		var ret TransferEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetEventTypeOk() (*TransferEventType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *TransferEvent) SetEventType(v TransferEventType) {
	o.EventType = v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TransferEvent) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TransferEvent) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *TransferEvent) SetAccountId(v string) {
	o.AccountId = &v
}

// GetFundingAccountId returns the FundingAccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferEvent) GetFundingAccountId() string {
	if o == nil || o.FundingAccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.FundingAccountId.Get()
}

// GetFundingAccountIdOk returns a tuple with the FundingAccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetFundingAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FundingAccountId.Get(), o.FundingAccountId.IsSet()
}

// SetFundingAccountId sets field value
func (o *TransferEvent) SetFundingAccountId(v string) {
	o.FundingAccountId.Set(&v)
}

// GetLedgerId returns the LedgerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEvent) GetLedgerId() string {
	if o == nil || o.LedgerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LedgerId.Get()
}

// GetLedgerIdOk returns a tuple with the LedgerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetLedgerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LedgerId.Get(), o.LedgerId.IsSet()
}

// HasLedgerId returns a boolean if a field has been set.
func (o *TransferEvent) HasLedgerId() bool {
	if o != nil && o.LedgerId.IsSet() {
		return true
	}

	return false
}

// SetLedgerId gets a reference to the given NullableString and assigns it to the LedgerId field.
func (o *TransferEvent) SetLedgerId(v string) {
	o.LedgerId.Set(&v)
}
// SetLedgerIdNil sets the value for LedgerId to be an explicit nil
func (o *TransferEvent) SetLedgerIdNil() {
	o.LedgerId.Set(nil)
}

// UnsetLedgerId ensures that no value is present for LedgerId, not even an explicit nil
func (o *TransferEvent) UnsetLedgerId() {
	o.LedgerId.Unset()
}

// GetTransferId returns the TransferId field value
func (o *TransferEvent) GetTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferId, true
}

// SetTransferId sets field value
func (o *TransferEvent) SetTransferId(v string) {
	o.TransferId = v
}

// GetOriginationAccountId returns the OriginationAccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferEvent) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// SetOriginationAccountId sets field value
func (o *TransferEvent) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}

// GetTransferType returns the TransferType field value if set, zero value otherwise.
func (o *TransferEvent) GetTransferType() OmittableTransferType {
	if o == nil || o.TransferType == nil {
		var ret OmittableTransferType
		return ret
	}
	return *o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTransferTypeOk() (*OmittableTransferType, bool) {
	if o == nil || o.TransferType == nil {
		return nil, false
	}
	return o.TransferType, true
}

// HasTransferType returns a boolean if a field has been set.
func (o *TransferEvent) HasTransferType() bool {
	if o != nil && o.TransferType != nil {
		return true
	}

	return false
}

// SetTransferType gets a reference to the given OmittableTransferType and assigns it to the TransferType field.
func (o *TransferEvent) SetTransferType(v OmittableTransferType) {
	o.TransferType = &v
}

// GetTransferAmount returns the TransferAmount field value if set, zero value otherwise.
func (o *TransferEvent) GetTransferAmount() string {
	if o == nil || o.TransferAmount == nil {
		var ret string
		return ret
	}
	return *o.TransferAmount
}

// GetTransferAmountOk returns a tuple with the TransferAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTransferAmountOk() (*string, bool) {
	if o == nil || o.TransferAmount == nil {
		return nil, false
	}
	return o.TransferAmount, true
}

// HasTransferAmount returns a boolean if a field has been set.
func (o *TransferEvent) HasTransferAmount() bool {
	if o != nil && o.TransferAmount != nil {
		return true
	}

	return false
}

// SetTransferAmount gets a reference to the given string and assigns it to the TransferAmount field.
func (o *TransferEvent) SetTransferAmount(v string) {
	o.TransferAmount = &v
}

// GetFailureReason returns the FailureReason field value
// If the value is explicit nil, the zero value for TransferFailure will be returned
func (o *TransferEvent) GetFailureReason() TransferFailure {
	if o == nil || o.FailureReason.Get() == nil {
		var ret TransferFailure
		return ret
	}

	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetFailureReasonOk() (*TransferFailure, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// SetFailureReason sets field value
func (o *TransferEvent) SetFailureReason(v TransferFailure) {
	o.FailureReason.Set(&v)
}

// GetSweepId returns the SweepId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferEvent) GetSweepId() string {
	if o == nil || o.SweepId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SweepId.Get()
}

// GetSweepIdOk returns a tuple with the SweepId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetSweepIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SweepId.Get(), o.SweepId.IsSet()
}

// SetSweepId sets field value
func (o *TransferEvent) SetSweepId(v string) {
	o.SweepId.Set(&v)
}

// GetSweepAmount returns the SweepAmount field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferEvent) GetSweepAmount() string {
	if o == nil || o.SweepAmount.Get() == nil {
		var ret string
		return ret
	}

	return *o.SweepAmount.Get()
}

// GetSweepAmountOk returns a tuple with the SweepAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetSweepAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SweepAmount.Get(), o.SweepAmount.IsSet()
}

// SetSweepAmount sets field value
func (o *TransferEvent) SetSweepAmount(v string) {
	o.SweepAmount.Set(&v)
}

// GetRefundId returns the RefundId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferEvent) GetRefundId() string {
	if o == nil || o.RefundId.Get() == nil {
		var ret string
		return ret
	}

	return *o.RefundId.Get()
}

// GetRefundIdOk returns a tuple with the RefundId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetRefundIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefundId.Get(), o.RefundId.IsSet()
}

// SetRefundId sets field value
func (o *TransferEvent) SetRefundId(v string) {
	o.RefundId.Set(&v)
}

// GetOriginatorClientId returns the OriginatorClientId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferEvent) GetOriginatorClientId() string {
	if o == nil || o.OriginatorClientId.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginatorClientId.Get()
}

// GetOriginatorClientIdOk returns a tuple with the OriginatorClientId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetOriginatorClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginatorClientId.Get(), o.OriginatorClientId.IsSet()
}

// SetOriginatorClientId sets field value
func (o *TransferEvent) SetOriginatorClientId(v string) {
	o.OriginatorClientId.Set(&v)
}

// GetIntentId returns the IntentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferEvent) GetIntentId() string {
	if o == nil || o.IntentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.IntentId.Get()
}

// GetIntentIdOk returns a tuple with the IntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetIntentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IntentId.Get(), o.IntentId.IsSet()
}

// HasIntentId returns a boolean if a field has been set.
func (o *TransferEvent) HasIntentId() bool {
	if o != nil && o.IntentId.IsSet() {
		return true
	}

	return false
}

// SetIntentId gets a reference to the given NullableString and assigns it to the IntentId field.
func (o *TransferEvent) SetIntentId(v string) {
	o.IntentId.Set(&v)
}
// SetIntentIdNil sets the value for IntentId to be an explicit nil
func (o *TransferEvent) SetIntentIdNil() {
	o.IntentId.Set(nil)
}

// UnsetIntentId ensures that no value is present for IntentId, not even an explicit nil
func (o *TransferEvent) UnsetIntentId() {
	o.IntentId.Unset()
}

func (o TransferEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_id"] = o.EventId
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["event_type"] = o.EventType
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["funding_account_id"] = o.FundingAccountId.Get()
	}
	if o.LedgerId.IsSet() {
		toSerialize["ledger_id"] = o.LedgerId.Get()
	}
	if true {
		toSerialize["transfer_id"] = o.TransferId
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	if o.TransferType != nil {
		toSerialize["transfer_type"] = o.TransferType
	}
	if o.TransferAmount != nil {
		toSerialize["transfer_amount"] = o.TransferAmount
	}
	if true {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}
	if true {
		toSerialize["sweep_id"] = o.SweepId.Get()
	}
	if true {
		toSerialize["sweep_amount"] = o.SweepAmount.Get()
	}
	if true {
		toSerialize["refund_id"] = o.RefundId.Get()
	}
	if true {
		toSerialize["originator_client_id"] = o.OriginatorClientId.Get()
	}
	if o.IntentId.IsSet() {
		toSerialize["intent_id"] = o.IntentId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferEvent) UnmarshalJSON(bytes []byte) (err error) {
	varTransferEvent := _TransferEvent{}

	if err = json.Unmarshal(bytes, &varTransferEvent); err == nil {
		*o = TransferEvent(varTransferEvent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "event_id")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "event_type")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "funding_account_id")
		delete(additionalProperties, "ledger_id")
		delete(additionalProperties, "transfer_id")
		delete(additionalProperties, "origination_account_id")
		delete(additionalProperties, "transfer_type")
		delete(additionalProperties, "transfer_amount")
		delete(additionalProperties, "failure_reason")
		delete(additionalProperties, "sweep_id")
		delete(additionalProperties, "sweep_amount")
		delete(additionalProperties, "refund_id")
		delete(additionalProperties, "originator_client_id")
		delete(additionalProperties, "intent_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferEvent struct {
	value *TransferEvent
	isSet bool
}

func (v NullableTransferEvent) Get() *TransferEvent {
	return v.value
}

func (v *NullableTransferEvent) Set(val *TransferEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEvent(val *TransferEvent) *NullableTransferEvent {
	return &NullableTransferEvent{value: val, isSet: true}
}

func (v NullableTransferEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


