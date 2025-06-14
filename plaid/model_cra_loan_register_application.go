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

// CraLoanRegisterApplication Contains loan application data to register.
type CraLoanRegisterApplication struct {
	// A unique identifier for the loan application. Personally identifiable information, such as an email address or phone number, should not be used in the `application_id`.
	ApplicationId *string `json:"application_id,omitempty"`
	// The date the user applied for the loan. The date should be in ISO 8601 format (YYYY-MM-DD).
	ApplicationDate *string `json:"application_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CraLoanRegisterApplication CraLoanRegisterApplication

// NewCraLoanRegisterApplication instantiates a new CraLoanRegisterApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraLoanRegisterApplication() *CraLoanRegisterApplication {
	this := CraLoanRegisterApplication{}
	return &this
}

// NewCraLoanRegisterApplicationWithDefaults instantiates a new CraLoanRegisterApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraLoanRegisterApplicationWithDefaults() *CraLoanRegisterApplication {
	this := CraLoanRegisterApplication{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *CraLoanRegisterApplication) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraLoanRegisterApplication) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *CraLoanRegisterApplication) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *CraLoanRegisterApplication) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetApplicationDate returns the ApplicationDate field value if set, zero value otherwise.
func (o *CraLoanRegisterApplication) GetApplicationDate() string {
	if o == nil || o.ApplicationDate == nil {
		var ret string
		return ret
	}
	return *o.ApplicationDate
}

// GetApplicationDateOk returns a tuple with the ApplicationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraLoanRegisterApplication) GetApplicationDateOk() (*string, bool) {
	if o == nil || o.ApplicationDate == nil {
		return nil, false
	}
	return o.ApplicationDate, true
}

// HasApplicationDate returns a boolean if a field has been set.
func (o *CraLoanRegisterApplication) HasApplicationDate() bool {
	if o != nil && o.ApplicationDate != nil {
		return true
	}

	return false
}

// SetApplicationDate gets a reference to the given string and assigns it to the ApplicationDate field.
func (o *CraLoanRegisterApplication) SetApplicationDate(v string) {
	o.ApplicationDate = &v
}

func (o CraLoanRegisterApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.ApplicationDate != nil {
		toSerialize["application_date"] = o.ApplicationDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraLoanRegisterApplication) UnmarshalJSON(bytes []byte) (err error) {
	varCraLoanRegisterApplication := _CraLoanRegisterApplication{}

	if err = json.Unmarshal(bytes, &varCraLoanRegisterApplication); err == nil {
		*o = CraLoanRegisterApplication(varCraLoanRegisterApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "application_id")
		delete(additionalProperties, "application_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraLoanRegisterApplication struct {
	value *CraLoanRegisterApplication
	isSet bool
}

func (v NullableCraLoanRegisterApplication) Get() *CraLoanRegisterApplication {
	return v.value
}

func (v *NullableCraLoanRegisterApplication) Set(val *CraLoanRegisterApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableCraLoanRegisterApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableCraLoanRegisterApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraLoanRegisterApplication(val *CraLoanRegisterApplication) *NullableCraLoanRegisterApplication {
	return &NullableCraLoanRegisterApplication{value: val, isSet: true}
}

func (v NullableCraLoanRegisterApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraLoanRegisterApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


