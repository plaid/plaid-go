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

// TransferLedgerBalance Information about the balance of the ledger held with Plaid.
type TransferLedgerBalance struct {
	// The amount of this balance available for use (decimal string with two digits of precision e.g. \"10.00\").
	Available string `json:"available"`
	// The amount of pending funds that are in processing (decimal string with two digits of precision e.g. \"10.00\").
	Pending string `json:"pending"`
	AdditionalProperties map[string]interface{}
}

type _TransferLedgerBalance TransferLedgerBalance

// NewTransferLedgerBalance instantiates a new TransferLedgerBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferLedgerBalance(available string, pending string) *TransferLedgerBalance {
	this := TransferLedgerBalance{}
	this.Available = available
	this.Pending = pending
	return &this
}

// NewTransferLedgerBalanceWithDefaults instantiates a new TransferLedgerBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferLedgerBalanceWithDefaults() *TransferLedgerBalance {
	this := TransferLedgerBalance{}
	return &this
}

// GetAvailable returns the Available field value
func (o *TransferLedgerBalance) GetAvailable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *TransferLedgerBalance) GetAvailableOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *TransferLedgerBalance) SetAvailable(v string) {
	o.Available = v
}

// GetPending returns the Pending field value
func (o *TransferLedgerBalance) GetPending() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *TransferLedgerBalance) GetPendingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *TransferLedgerBalance) SetPending(v string) {
	o.Pending = v
}

func (o TransferLedgerBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["available"] = o.Available
	}
	if true {
		toSerialize["pending"] = o.Pending
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferLedgerBalance) UnmarshalJSON(bytes []byte) (err error) {
	varTransferLedgerBalance := _TransferLedgerBalance{}

	if err = json.Unmarshal(bytes, &varTransferLedgerBalance); err == nil {
		*o = TransferLedgerBalance(varTransferLedgerBalance)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "available")
		delete(additionalProperties, "pending")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferLedgerBalance struct {
	value *TransferLedgerBalance
	isSet bool
}

func (v NullableTransferLedgerBalance) Get() *TransferLedgerBalance {
	return v.value
}

func (v *NullableTransferLedgerBalance) Set(val *TransferLedgerBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferLedgerBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferLedgerBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferLedgerBalance(val *TransferLedgerBalance) *NullableTransferLedgerBalance {
	return &NullableTransferLedgerBalance{value: val, isSet: true}
}

func (v NullableTransferLedgerBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferLedgerBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


