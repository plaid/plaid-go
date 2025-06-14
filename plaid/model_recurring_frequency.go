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
	"fmt"
)

// RecurringFrequency Describes the frequency of the transaction stream.  `WEEKLY`: Assigned to a transaction stream that occurs approximately every week.  `BIWEEKLY`: Assigned to a transaction stream that occurs approximately every 2 weeks.  `SEMI_MONTHLY`: Assigned to a transaction stream that occurs approximately twice per month. This frequency is typically seen for inflow transaction streams.  `MONTHLY`: Assigned to a transaction stream that occurs approximately every month.  `ANNUALLY`: Assigned to a transaction stream that occurs approximately every year.  `DAILY`: Assigned to a transaction stream that occurs approximately every day.  `DYNAMIC`: Assigned to a transaction stream that varies in recurrence frequency. This frequency is typically seen for inflow streams in the gig economy.  `UNKNOWN`: Assigned to a transaction stream that isn't recurring in nature.
type RecurringFrequency string

var _ = fmt.Printf

// List of RecurringFrequency
const (
	RECURRINGFREQUENCY_UNKNOWN RecurringFrequency = "UNKNOWN"
	RECURRINGFREQUENCY_WEEKLY RecurringFrequency = "WEEKLY"
	RECURRINGFREQUENCY_BIWEEKLY RecurringFrequency = "BIWEEKLY"
	RECURRINGFREQUENCY_SEMI_MONTHLY RecurringFrequency = "SEMI_MONTHLY"
	RECURRINGFREQUENCY_MONTHLY RecurringFrequency = "MONTHLY"
	RECURRINGFREQUENCY_ANNUALLY RecurringFrequency = "ANNUALLY"
	RECURRINGFREQUENCY_DAILY RecurringFrequency = "DAILY"
	RECURRINGFREQUENCY_DYNAMIC RecurringFrequency = "DYNAMIC"
	RECURRINGFREQUENCY_NULL RecurringFrequency = "null"
)

var allowedRecurringFrequencyEnumValues = []RecurringFrequency{
	"UNKNOWN",
	"WEEKLY",
	"BIWEEKLY",
	"SEMI_MONTHLY",
	"MONTHLY",
	"ANNUALLY",
	"DAILY",
	"DYNAMIC",
	"null",
}

func (v *RecurringFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := RecurringFrequency(value)


	*v = enumTypeValue
	return nil
}

// NewRecurringFrequencyFromValue returns a pointer to a valid RecurringFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecurringFrequencyFromValue(v string) (*RecurringFrequency, error) {
	ev := RecurringFrequency(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecurringFrequency) IsValid() bool {
	for _, existing := range allowedRecurringFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecurringFrequency value
func (v RecurringFrequency) Ptr() *RecurringFrequency {
	return &v
}

type NullableRecurringFrequency struct {
	value *RecurringFrequency
	isSet bool
}

func (v NullableRecurringFrequency) Get() *RecurringFrequency {
	return v.value
}

func (v *NullableRecurringFrequency) Set(val *RecurringFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringFrequency(val *RecurringFrequency) *NullableRecurringFrequency {
	return &NullableRecurringFrequency{value: val, isSet: true}
}

func (v NullableRecurringFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

