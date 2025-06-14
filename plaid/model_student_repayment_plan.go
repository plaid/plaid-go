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

// StudentRepaymentPlan An object representing the repayment plan for the student loan
type StudentRepaymentPlan struct {
	// The description of the repayment plan as provided by the servicer.
	Description NullableString `json:"description"`
	// The type of the repayment plan.
	Type NullableString `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _StudentRepaymentPlan StudentRepaymentPlan

// NewStudentRepaymentPlan instantiates a new StudentRepaymentPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudentRepaymentPlan(description NullableString, type_ NullableString) *StudentRepaymentPlan {
	this := StudentRepaymentPlan{}
	this.Description = description
	this.Type = type_
	return &this
}

// NewStudentRepaymentPlanWithDefaults instantiates a new StudentRepaymentPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudentRepaymentPlanWithDefaults() *StudentRepaymentPlan {
	this := StudentRepaymentPlan{}
	return &this
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentRepaymentPlan) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentRepaymentPlan) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *StudentRepaymentPlan) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StudentRepaymentPlan) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentRepaymentPlan) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *StudentRepaymentPlan) SetType(v string) {
	o.Type.Set(&v)
}

func (o StudentRepaymentPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StudentRepaymentPlan) UnmarshalJSON(bytes []byte) (err error) {
	varStudentRepaymentPlan := _StudentRepaymentPlan{}

	if err = json.Unmarshal(bytes, &varStudentRepaymentPlan); err == nil {
		*o = StudentRepaymentPlan(varStudentRepaymentPlan)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudentRepaymentPlan struct {
	value *StudentRepaymentPlan
	isSet bool
}

func (v NullableStudentRepaymentPlan) Get() *StudentRepaymentPlan {
	return v.value
}

func (v *NullableStudentRepaymentPlan) Set(val *StudentRepaymentPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableStudentRepaymentPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableStudentRepaymentPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudentRepaymentPlan(val *StudentRepaymentPlan) *NullableStudentRepaymentPlan {
	return &NullableStudentRepaymentPlan{value: val, isSet: true}
}

func (v NullableStudentRepaymentPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudentRepaymentPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


