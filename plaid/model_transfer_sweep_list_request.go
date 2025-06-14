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

// TransferSweepListRequest Defines the request schema for `/transfer/sweep/list`
type TransferSweepListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The start `created` datetime of sweeps to return (RFC 3339 format).
	StartDate NullableTime `json:"start_date,omitempty"`
	// The end `created` datetime of sweeps to return (RFC 3339 format).
	EndDate NullableTime `json:"end_date,omitempty"`
	// The maximum number of sweeps to return.
	Count NullableInt32 `json:"count,omitempty"`
	// The number of sweeps to skip before returning results.
	Offset *int32 `json:"offset,omitempty"`
	// Filter sweeps to only those with the specified amount.
	Amount NullableString `json:"amount,omitempty"`
	Status NullableSweepStatus `json:"status,omitempty"`
	// Filter sweeps to only those with the specified originator client.
	OriginatorClientId NullableString `json:"originator_client_id,omitempty"`
	// Filter sweeps to only those with the specified `funding_account_id`.
	FundingAccountId NullableString `json:"funding_account_id,omitempty"`
	// Filter sweeps to only those with the included `transfer_id`.
	TransferId NullableString `json:"transfer_id,omitempty"`
	Trigger NullableSweepTrigger `json:"trigger,omitempty"`
}

// NewTransferSweepListRequest instantiates a new TransferSweepListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferSweepListRequest() *TransferSweepListRequest {
	this := TransferSweepListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// NewTransferSweepListRequestWithDefaults instantiates a new TransferSweepListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferSweepListRequestWithDefaults() *TransferSweepListRequest {
	this := TransferSweepListRequest{}
	var count int32 = 25
	this.Count = *NewNullableInt32(&count)
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransferSweepListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferSweepListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransferSweepListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransferSweepListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferSweepListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransferSweepListRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferSweepListRequest) GetStartDate() time.Time {
	if o == nil || o.StartDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferSweepListRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *TransferSweepListRequest) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *TransferSweepListRequest) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *TransferSweepListRequest) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferSweepListRequest) GetEndDate() time.Time {
	if o == nil || o.EndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferSweepListRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *TransferSweepListRequest) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *TransferSweepListRequest) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *TransferSweepListRequest) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferSweepListRequest) GetCount() int32 {
	if o == nil || o.Count.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferSweepListRequest) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *TransferSweepListRequest) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *TransferSweepListRequest) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *TransferSweepListRequest) UnsetCount() {
	o.Count.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TransferSweepListRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferSweepListRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *TransferSweepListRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferSweepListRequest) GetAmount() string {
	if o == nil || o.Amount.Get() == nil {
		var ret string
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferSweepListRequest) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableString and assigns it to the Amount field.
func (o *TransferSweepListRequest) SetAmount(v string) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *TransferSweepListRequest) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *TransferSweepListRequest) UnsetAmount() {
	o.Amount.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferSweepListRequest) GetStatus() SweepStatus {
	if o == nil || o.Status.Get() == nil {
		var ret SweepStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferSweepListRequest) GetStatusOk() (*SweepStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableSweepStatus and assigns it to the Status field.
func (o *TransferSweepListRequest) SetStatus(v SweepStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TransferSweepListRequest) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TransferSweepListRequest) UnsetStatus() {
	o.Status.Unset()
}

// GetOriginatorClientId returns the OriginatorClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferSweepListRequest) GetOriginatorClientId() string {
	if o == nil || o.OriginatorClientId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginatorClientId.Get()
}

// GetOriginatorClientIdOk returns a tuple with the OriginatorClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferSweepListRequest) GetOriginatorClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginatorClientId.Get(), o.OriginatorClientId.IsSet()
}

// HasOriginatorClientId returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasOriginatorClientId() bool {
	if o != nil && o.OriginatorClientId.IsSet() {
		return true
	}

	return false
}

// SetOriginatorClientId gets a reference to the given NullableString and assigns it to the OriginatorClientId field.
func (o *TransferSweepListRequest) SetOriginatorClientId(v string) {
	o.OriginatorClientId.Set(&v)
}
// SetOriginatorClientIdNil sets the value for OriginatorClientId to be an explicit nil
func (o *TransferSweepListRequest) SetOriginatorClientIdNil() {
	o.OriginatorClientId.Set(nil)
}

// UnsetOriginatorClientId ensures that no value is present for OriginatorClientId, not even an explicit nil
func (o *TransferSweepListRequest) UnsetOriginatorClientId() {
	o.OriginatorClientId.Unset()
}

// GetFundingAccountId returns the FundingAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferSweepListRequest) GetFundingAccountId() string {
	if o == nil || o.FundingAccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.FundingAccountId.Get()
}

// GetFundingAccountIdOk returns a tuple with the FundingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferSweepListRequest) GetFundingAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FundingAccountId.Get(), o.FundingAccountId.IsSet()
}

// HasFundingAccountId returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasFundingAccountId() bool {
	if o != nil && o.FundingAccountId.IsSet() {
		return true
	}

	return false
}

// SetFundingAccountId gets a reference to the given NullableString and assigns it to the FundingAccountId field.
func (o *TransferSweepListRequest) SetFundingAccountId(v string) {
	o.FundingAccountId.Set(&v)
}
// SetFundingAccountIdNil sets the value for FundingAccountId to be an explicit nil
func (o *TransferSweepListRequest) SetFundingAccountIdNil() {
	o.FundingAccountId.Set(nil)
}

// UnsetFundingAccountId ensures that no value is present for FundingAccountId, not even an explicit nil
func (o *TransferSweepListRequest) UnsetFundingAccountId() {
	o.FundingAccountId.Unset()
}

// GetTransferId returns the TransferId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferSweepListRequest) GetTransferId() string {
	if o == nil || o.TransferId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransferId.Get()
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferSweepListRequest) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransferId.Get(), o.TransferId.IsSet()
}

// HasTransferId returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasTransferId() bool {
	if o != nil && o.TransferId.IsSet() {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given NullableString and assigns it to the TransferId field.
func (o *TransferSweepListRequest) SetTransferId(v string) {
	o.TransferId.Set(&v)
}
// SetTransferIdNil sets the value for TransferId to be an explicit nil
func (o *TransferSweepListRequest) SetTransferIdNil() {
	o.TransferId.Set(nil)
}

// UnsetTransferId ensures that no value is present for TransferId, not even an explicit nil
func (o *TransferSweepListRequest) UnsetTransferId() {
	o.TransferId.Unset()
}

// GetTrigger returns the Trigger field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferSweepListRequest) GetTrigger() SweepTrigger {
	if o == nil || o.Trigger.Get() == nil {
		var ret SweepTrigger
		return ret
	}
	return *o.Trigger.Get()
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferSweepListRequest) GetTriggerOk() (*SweepTrigger, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Trigger.Get(), o.Trigger.IsSet()
}

// HasTrigger returns a boolean if a field has been set.
func (o *TransferSweepListRequest) HasTrigger() bool {
	if o != nil && o.Trigger.IsSet() {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given NullableSweepTrigger and assigns it to the Trigger field.
func (o *TransferSweepListRequest) SetTrigger(v SweepTrigger) {
	o.Trigger.Set(&v)
}
// SetTriggerNil sets the value for Trigger to be an explicit nil
func (o *TransferSweepListRequest) SetTriggerNil() {
	o.Trigger.Set(nil)
}

// UnsetTrigger ensures that no value is present for Trigger, not even an explicit nil
func (o *TransferSweepListRequest) UnsetTrigger() {
	o.Trigger.Unset()
}

func (o TransferSweepListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.OriginatorClientId.IsSet() {
		toSerialize["originator_client_id"] = o.OriginatorClientId.Get()
	}
	if o.FundingAccountId.IsSet() {
		toSerialize["funding_account_id"] = o.FundingAccountId.Get()
	}
	if o.TransferId.IsSet() {
		toSerialize["transfer_id"] = o.TransferId.Get()
	}
	if o.Trigger.IsSet() {
		toSerialize["trigger"] = o.Trigger.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTransferSweepListRequest struct {
	value *TransferSweepListRequest
	isSet bool
}

func (v NullableTransferSweepListRequest) Get() *TransferSweepListRequest {
	return v.value
}

func (v *NullableTransferSweepListRequest) Set(val *TransferSweepListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferSweepListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferSweepListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferSweepListRequest(val *TransferSweepListRequest) *NullableTransferSweepListRequest {
	return &NullableTransferSweepListRequest{value: val, isSet: true}
}

func (v NullableTransferSweepListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferSweepListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


