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
)

// CreditFreddieMacIndividualName Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacIndividualName struct {
	// The first name of the individual represented by the parent object.
	FirstName string `json:"FirstName"`
	// The last name of the individual represented by the parent object.
	LastName string `json:"LastName"`
	// The middle name of the individual represented by the parent object.
	MiddleName string `json:"MiddleName"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacIndividualName CreditFreddieMacIndividualName

// NewCreditFreddieMacIndividualName instantiates a new CreditFreddieMacIndividualName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacIndividualName(firstName string, lastName string, middleName string) *CreditFreddieMacIndividualName {
	this := CreditFreddieMacIndividualName{}
	this.FirstName = firstName
	this.LastName = lastName
	this.MiddleName = middleName
	return &this
}

// NewCreditFreddieMacIndividualNameWithDefaults instantiates a new CreditFreddieMacIndividualName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacIndividualNameWithDefaults() *CreditFreddieMacIndividualName {
	this := CreditFreddieMacIndividualName{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *CreditFreddieMacIndividualName) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacIndividualName) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *CreditFreddieMacIndividualName) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *CreditFreddieMacIndividualName) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacIndividualName) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *CreditFreddieMacIndividualName) SetLastName(v string) {
	o.LastName = v
}

// GetMiddleName returns the MiddleName field value
func (o *CreditFreddieMacIndividualName) GetMiddleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacIndividualName) GetMiddleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MiddleName, true
}

// SetMiddleName sets field value
func (o *CreditFreddieMacIndividualName) SetMiddleName(v string) {
	o.MiddleName = v
}

func (o CreditFreddieMacIndividualName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["FirstName"] = o.FirstName
	}
	if true {
		toSerialize["LastName"] = o.LastName
	}
	if true {
		toSerialize["MiddleName"] = o.MiddleName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacIndividualName) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacIndividualName := _CreditFreddieMacIndividualName{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacIndividualName); err == nil {
		*o = CreditFreddieMacIndividualName(varCreditFreddieMacIndividualName)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "FirstName")
		delete(additionalProperties, "LastName")
		delete(additionalProperties, "MiddleName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacIndividualName struct {
	value *CreditFreddieMacIndividualName
	isSet bool
}

func (v NullableCreditFreddieMacIndividualName) Get() *CreditFreddieMacIndividualName {
	return v.value
}

func (v *NullableCreditFreddieMacIndividualName) Set(val *CreditFreddieMacIndividualName) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacIndividualName) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacIndividualName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacIndividualName(val *CreditFreddieMacIndividualName) *NullableCreditFreddieMacIndividualName {
	return &NullableCreditFreddieMacIndividualName{value: val, isSet: true}
}

func (v NullableCreditFreddieMacIndividualName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacIndividualName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

