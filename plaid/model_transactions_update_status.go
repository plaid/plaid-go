/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.503.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// TransactionsUpdateStatus A description of the update status for transaction pulls of an item.  `TRANSACTIONS_UPDATE_STATUS_UNKNOWN`: Unable to fetch transactions update status for item. `NOT_READY`: The item is pending transaction pull. `INITIAL_UPDATE_COMPLETE`: Initial pull for the item is complete, pending Historical pull. `HISTORICAL_UPDATE_COMPLETE`: Both Initial and Historical pull for item is complete.
type TransactionsUpdateStatus string

var _ = fmt.Printf

// List of TransactionsUpdateStatus
const (
	TRANSACTIONSUPDATESTATUS_TRANSACTIONS_UPDATE_STATUS_UNKNOWN TransactionsUpdateStatus = "TRANSACTIONS_UPDATE_STATUS_UNKNOWN"
	TRANSACTIONSUPDATESTATUS_NOT_READY TransactionsUpdateStatus = "NOT_READY"
	TRANSACTIONSUPDATESTATUS_INITIAL_UPDATE_COMPLETE TransactionsUpdateStatus = "INITIAL_UPDATE_COMPLETE"
	TRANSACTIONSUPDATESTATUS_HISTORICAL_UPDATE_COMPLETE TransactionsUpdateStatus = "HISTORICAL_UPDATE_COMPLETE"
)

var allowedTransactionsUpdateStatusEnumValues = []TransactionsUpdateStatus{
	"TRANSACTIONS_UPDATE_STATUS_UNKNOWN",
	"NOT_READY",
	"INITIAL_UPDATE_COMPLETE",
	"HISTORICAL_UPDATE_COMPLETE",
}

func (v *TransactionsUpdateStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransactionsUpdateStatus(value)


	*v = enumTypeValue
	return nil
}

// NewTransactionsUpdateStatusFromValue returns a pointer to a valid TransactionsUpdateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionsUpdateStatusFromValue(v string) (*TransactionsUpdateStatus, error) {
	ev := TransactionsUpdateStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionsUpdateStatus) IsValid() bool {
	for _, existing := range allowedTransactionsUpdateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionsUpdateStatus value
func (v TransactionsUpdateStatus) Ptr() *TransactionsUpdateStatus {
	return &v
}

type NullableTransactionsUpdateStatus struct {
	value *TransactionsUpdateStatus
	isSet bool
}

func (v NullableTransactionsUpdateStatus) Get() *TransactionsUpdateStatus {
	return v.value
}

func (v *NullableTransactionsUpdateStatus) Set(val *TransactionsUpdateStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsUpdateStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsUpdateStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsUpdateStatus(val *TransactionsUpdateStatus) *NullableTransactionsUpdateStatus {
	return &NullableTransactionsUpdateStatus{value: val, isSet: true}
}

func (v NullableTransactionsUpdateStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsUpdateStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
