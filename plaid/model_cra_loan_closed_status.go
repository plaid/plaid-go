/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.586.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// CraLoanClosedStatus Contains the status and date information of the loan when unregistering.
type CraLoanClosedStatus struct {
	Status CraLoanStatus `json:"status"`
	// The effective date for the status of the loan. The date should be in ISO 8601 format (YYYY-MM-DD).
	Date string `json:"date"`
	AdditionalProperties map[string]interface{}
}

type _CraLoanClosedStatus CraLoanClosedStatus

// NewCraLoanClosedStatus instantiates a new CraLoanClosedStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraLoanClosedStatus(status CraLoanStatus, date string) *CraLoanClosedStatus {
	this := CraLoanClosedStatus{}
	this.Status = status
	this.Date = date
	return &this
}

// NewCraLoanClosedStatusWithDefaults instantiates a new CraLoanClosedStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraLoanClosedStatusWithDefaults() *CraLoanClosedStatus {
	this := CraLoanClosedStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *CraLoanClosedStatus) GetStatus() CraLoanStatus {
	if o == nil {
		var ret CraLoanStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CraLoanClosedStatus) GetStatusOk() (*CraLoanStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CraLoanClosedStatus) SetStatus(v CraLoanStatus) {
	o.Status = v
}

// GetDate returns the Date field value
func (o *CraLoanClosedStatus) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *CraLoanClosedStatus) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *CraLoanClosedStatus) SetDate(v string) {
	o.Date = v
}

func (o CraLoanClosedStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["date"] = o.Date
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraLoanClosedStatus) UnmarshalJSON(bytes []byte) (err error) {
	varCraLoanClosedStatus := _CraLoanClosedStatus{}

	if err = json.Unmarshal(bytes, &varCraLoanClosedStatus); err == nil {
		*o = CraLoanClosedStatus(varCraLoanClosedStatus)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraLoanClosedStatus struct {
	value *CraLoanClosedStatus
	isSet bool
}

func (v NullableCraLoanClosedStatus) Get() *CraLoanClosedStatus {
	return v.value
}

func (v *NullableCraLoanClosedStatus) Set(val *CraLoanClosedStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCraLoanClosedStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCraLoanClosedStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraLoanClosedStatus(val *CraLoanClosedStatus) *NullableCraLoanClosedStatus {
	return &NullableCraLoanClosedStatus{value: val, isSet: true}
}

func (v NullableCraLoanClosedStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraLoanClosedStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

