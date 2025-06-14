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

// CreditPlatformIds The object containing a set of ids related to an employee.
type CreditPlatformIds struct {
	// The ID of an employee as given by their employer.
	EmployeeId NullableString `json:"employee_id"`
	// The ID of an employee as given by their payroll.
	PayrollId NullableString `json:"payroll_id"`
	// The ID of the position of the employee.
	PositionId NullableString `json:"position_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditPlatformIds CreditPlatformIds

// NewCreditPlatformIds instantiates a new CreditPlatformIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPlatformIds(employeeId NullableString, payrollId NullableString, positionId NullableString) *CreditPlatformIds {
	this := CreditPlatformIds{}
	this.EmployeeId = employeeId
	this.PayrollId = payrollId
	this.PositionId = positionId
	return &this
}

// NewCreditPlatformIdsWithDefaults instantiates a new CreditPlatformIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPlatformIdsWithDefaults() *CreditPlatformIds {
	this := CreditPlatformIds{}
	return &this
}

// GetEmployeeId returns the EmployeeId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPlatformIds) GetEmployeeId() string {
	if o == nil || o.EmployeeId.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmployeeId.Get()
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPlatformIds) GetEmployeeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeeId.Get(), o.EmployeeId.IsSet()
}

// SetEmployeeId sets field value
func (o *CreditPlatformIds) SetEmployeeId(v string) {
	o.EmployeeId.Set(&v)
}

// GetPayrollId returns the PayrollId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPlatformIds) GetPayrollId() string {
	if o == nil || o.PayrollId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PayrollId.Get()
}

// GetPayrollIdOk returns a tuple with the PayrollId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPlatformIds) GetPayrollIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayrollId.Get(), o.PayrollId.IsSet()
}

// SetPayrollId sets field value
func (o *CreditPlatformIds) SetPayrollId(v string) {
	o.PayrollId.Set(&v)
}

// GetPositionId returns the PositionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditPlatformIds) GetPositionId() string {
	if o == nil || o.PositionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PositionId.Get()
}

// GetPositionIdOk returns a tuple with the PositionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditPlatformIds) GetPositionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PositionId.Get(), o.PositionId.IsSet()
}

// SetPositionId sets field value
func (o *CreditPlatformIds) SetPositionId(v string) {
	o.PositionId.Set(&v)
}

func (o CreditPlatformIds) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["employee_id"] = o.EmployeeId.Get()
	}
	if true {
		toSerialize["payroll_id"] = o.PayrollId.Get()
	}
	if true {
		toSerialize["position_id"] = o.PositionId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditPlatformIds) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPlatformIds := _CreditPlatformIds{}

	if err = json.Unmarshal(bytes, &varCreditPlatformIds); err == nil {
		*o = CreditPlatformIds(varCreditPlatformIds)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "employee_id")
		delete(additionalProperties, "payroll_id")
		delete(additionalProperties, "position_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPlatformIds struct {
	value *CreditPlatformIds
	isSet bool
}

func (v NullableCreditPlatformIds) Get() *CreditPlatformIds {
	return v.value
}

func (v *NullableCreditPlatformIds) Set(val *CreditPlatformIds) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPlatformIds) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPlatformIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPlatformIds(val *CreditPlatformIds) *NullableCreditPlatformIds {
	return &NullableCreditPlatformIds{value: val, isSet: true}
}

func (v NullableCreditPlatformIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPlatformIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


