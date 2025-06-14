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

// BusinessFinanceCategory Information describing the intent of the transaction. Most relevant for business finance use cases, but not limited to such use cases.
type BusinessFinanceCategory struct {
	// A high level category that communicates the broad category of the transaction.
	Primary string `json:"primary"`
	// A granular category conveying the transaction's intent. This field can also be used as a unique identifier for the category.
	Detailed string `json:"detailed"`
	// A description of how confident we are that the provided categories accurately describe the transaction intent.  `VERY_HIGH`: We are more than 98% confident that this category reflects the intent of the transaction. `HIGH`: We are more than 90% confident that this category reflects the intent of the transaction. `MEDIUM`: We are moderately confident that this category reflects the intent of the transaction. `LOW`: This category may reflect the intent, but there may be other categories that are more accurate. `UNKNOWN`: We don’t know the confidence level for this category.
	ConfidenceLevel NullableString `json:"confidence_level,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BusinessFinanceCategory BusinessFinanceCategory

// NewBusinessFinanceCategory instantiates a new BusinessFinanceCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessFinanceCategory(primary string, detailed string) *BusinessFinanceCategory {
	this := BusinessFinanceCategory{}
	this.Primary = primary
	this.Detailed = detailed
	return &this
}

// NewBusinessFinanceCategoryWithDefaults instantiates a new BusinessFinanceCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessFinanceCategoryWithDefaults() *BusinessFinanceCategory {
	this := BusinessFinanceCategory{}
	return &this
}

// GetPrimary returns the Primary field value
func (o *BusinessFinanceCategory) GetPrimary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *BusinessFinanceCategory) GetPrimaryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *BusinessFinanceCategory) SetPrimary(v string) {
	o.Primary = v
}

// GetDetailed returns the Detailed field value
func (o *BusinessFinanceCategory) GetDetailed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detailed
}

// GetDetailedOk returns a tuple with the Detailed field value
// and a boolean to check if the value has been set.
func (o *BusinessFinanceCategory) GetDetailedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Detailed, true
}

// SetDetailed sets field value
func (o *BusinessFinanceCategory) SetDetailed(v string) {
	o.Detailed = v
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessFinanceCategory) GetConfidenceLevel() string {
	if o == nil || o.ConfidenceLevel.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConfidenceLevel.Get()
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessFinanceCategory) GetConfidenceLevelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConfidenceLevel.Get(), o.ConfidenceLevel.IsSet()
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *BusinessFinanceCategory) HasConfidenceLevel() bool {
	if o != nil && o.ConfidenceLevel.IsSet() {
		return true
	}

	return false
}

// SetConfidenceLevel gets a reference to the given NullableString and assigns it to the ConfidenceLevel field.
func (o *BusinessFinanceCategory) SetConfidenceLevel(v string) {
	o.ConfidenceLevel.Set(&v)
}
// SetConfidenceLevelNil sets the value for ConfidenceLevel to be an explicit nil
func (o *BusinessFinanceCategory) SetConfidenceLevelNil() {
	o.ConfidenceLevel.Set(nil)
}

// UnsetConfidenceLevel ensures that no value is present for ConfidenceLevel, not even an explicit nil
func (o *BusinessFinanceCategory) UnsetConfidenceLevel() {
	o.ConfidenceLevel.Unset()
}

func (o BusinessFinanceCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["primary"] = o.Primary
	}
	if true {
		toSerialize["detailed"] = o.Detailed
	}
	if o.ConfidenceLevel.IsSet() {
		toSerialize["confidence_level"] = o.ConfidenceLevel.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BusinessFinanceCategory) UnmarshalJSON(bytes []byte) (err error) {
	varBusinessFinanceCategory := _BusinessFinanceCategory{}

	if err = json.Unmarshal(bytes, &varBusinessFinanceCategory); err == nil {
		*o = BusinessFinanceCategory(varBusinessFinanceCategory)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "primary")
		delete(additionalProperties, "detailed")
		delete(additionalProperties, "confidence_level")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBusinessFinanceCategory struct {
	value *BusinessFinanceCategory
	isSet bool
}

func (v NullableBusinessFinanceCategory) Get() *BusinessFinanceCategory {
	return v.value
}

func (v *NullableBusinessFinanceCategory) Set(val *BusinessFinanceCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessFinanceCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessFinanceCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessFinanceCategory(val *BusinessFinanceCategory) *NullableBusinessFinanceCategory {
	return &NullableBusinessFinanceCategory{value: val, isSet: true}
}

func (v NullableBusinessFinanceCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessFinanceCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


