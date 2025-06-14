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

// CreditEmploymentItem The object containing employment items.
type CreditEmploymentItem struct {
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	Employments []CreditEmploymentVerification `json:"employments"`
	// Token to represent the underlying Employment data
	EmploymentReportToken *string `json:"employment_report_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreditEmploymentItem CreditEmploymentItem

// NewCreditEmploymentItem instantiates a new CreditEmploymentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditEmploymentItem(itemId string, employments []CreditEmploymentVerification) *CreditEmploymentItem {
	this := CreditEmploymentItem{}
	this.ItemId = itemId
	this.Employments = employments
	return &this
}

// NewCreditEmploymentItemWithDefaults instantiates a new CreditEmploymentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditEmploymentItemWithDefaults() *CreditEmploymentItem {
	this := CreditEmploymentItem{}
	return &this
}

// GetItemId returns the ItemId field value
func (o *CreditEmploymentItem) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *CreditEmploymentItem) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *CreditEmploymentItem) SetItemId(v string) {
	o.ItemId = v
}

// GetEmployments returns the Employments field value
func (o *CreditEmploymentItem) GetEmployments() []CreditEmploymentVerification {
	if o == nil {
		var ret []CreditEmploymentVerification
		return ret
	}

	return o.Employments
}

// GetEmploymentsOk returns a tuple with the Employments field value
// and a boolean to check if the value has been set.
func (o *CreditEmploymentItem) GetEmploymentsOk() (*[]CreditEmploymentVerification, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Employments, true
}

// SetEmployments sets field value
func (o *CreditEmploymentItem) SetEmployments(v []CreditEmploymentVerification) {
	o.Employments = v
}

// GetEmploymentReportToken returns the EmploymentReportToken field value if set, zero value otherwise.
func (o *CreditEmploymentItem) GetEmploymentReportToken() string {
	if o == nil || o.EmploymentReportToken == nil {
		var ret string
		return ret
	}
	return *o.EmploymentReportToken
}

// GetEmploymentReportTokenOk returns a tuple with the EmploymentReportToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditEmploymentItem) GetEmploymentReportTokenOk() (*string, bool) {
	if o == nil || o.EmploymentReportToken == nil {
		return nil, false
	}
	return o.EmploymentReportToken, true
}

// HasEmploymentReportToken returns a boolean if a field has been set.
func (o *CreditEmploymentItem) HasEmploymentReportToken() bool {
	if o != nil && o.EmploymentReportToken != nil {
		return true
	}

	return false
}

// SetEmploymentReportToken gets a reference to the given string and assigns it to the EmploymentReportToken field.
func (o *CreditEmploymentItem) SetEmploymentReportToken(v string) {
	o.EmploymentReportToken = &v
}

func (o CreditEmploymentItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["employments"] = o.Employments
	}
	if o.EmploymentReportToken != nil {
		toSerialize["employment_report_token"] = o.EmploymentReportToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditEmploymentItem) UnmarshalJSON(bytes []byte) (err error) {
	varCreditEmploymentItem := _CreditEmploymentItem{}

	if err = json.Unmarshal(bytes, &varCreditEmploymentItem); err == nil {
		*o = CreditEmploymentItem(varCreditEmploymentItem)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "employments")
		delete(additionalProperties, "employment_report_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditEmploymentItem struct {
	value *CreditEmploymentItem
	isSet bool
}

func (v NullableCreditEmploymentItem) Get() *CreditEmploymentItem {
	return v.value
}

func (v *NullableCreditEmploymentItem) Set(val *CreditEmploymentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditEmploymentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditEmploymentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditEmploymentItem(val *CreditEmploymentItem) *NullableCreditEmploymentItem {
	return &NullableCreditEmploymentItem{value: val, isSet: true}
}

func (v NullableCreditEmploymentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditEmploymentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


