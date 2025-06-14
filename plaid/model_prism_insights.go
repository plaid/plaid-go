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

// PrismInsights The data from the Insights product returned by Prism Data.
type PrismInsights struct {
	// The version of Prism Data's insights model used.
	Version int32 `json:"version"`
	// The Insights Result object is a map of cash flow attributes, where the key is a string, and the value is a float or string.
	Result *map[string]interface{} `json:"result,omitempty"`
	// The error returned by Prism for this product.
	ErrorReason *string `json:"error_reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrismInsights PrismInsights

// NewPrismInsights instantiates a new PrismInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrismInsights(version int32) *PrismInsights {
	this := PrismInsights{}
	this.Version = version
	return &this
}

// NewPrismInsightsWithDefaults instantiates a new PrismInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrismInsightsWithDefaults() *PrismInsights {
	this := PrismInsights{}
	return &this
}

// GetVersion returns the Version field value
func (o *PrismInsights) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *PrismInsights) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *PrismInsights) SetVersion(v int32) {
	o.Version = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *PrismInsights) GetResult() map[string]interface{} {
	if o == nil || o.Result == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrismInsights) GetResultOk() (*map[string]interface{}, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PrismInsights) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given map[string]interface{} and assigns it to the Result field.
func (o *PrismInsights) SetResult(v map[string]interface{}) {
	o.Result = &v
}

// GetErrorReason returns the ErrorReason field value if set, zero value otherwise.
func (o *PrismInsights) GetErrorReason() string {
	if o == nil || o.ErrorReason == nil {
		var ret string
		return ret
	}
	return *o.ErrorReason
}

// GetErrorReasonOk returns a tuple with the ErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrismInsights) GetErrorReasonOk() (*string, bool) {
	if o == nil || o.ErrorReason == nil {
		return nil, false
	}
	return o.ErrorReason, true
}

// HasErrorReason returns a boolean if a field has been set.
func (o *PrismInsights) HasErrorReason() bool {
	if o != nil && o.ErrorReason != nil {
		return true
	}

	return false
}

// SetErrorReason gets a reference to the given string and assigns it to the ErrorReason field.
func (o *PrismInsights) SetErrorReason(v string) {
	o.ErrorReason = &v
}

func (o PrismInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.ErrorReason != nil {
		toSerialize["error_reason"] = o.ErrorReason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrismInsights) UnmarshalJSON(bytes []byte) (err error) {
	varPrismInsights := _PrismInsights{}

	if err = json.Unmarshal(bytes, &varPrismInsights); err == nil {
		*o = PrismInsights(varPrismInsights)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "result")
		delete(additionalProperties, "error_reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrismInsights struct {
	value *PrismInsights
	isSet bool
}

func (v NullablePrismInsights) Get() *PrismInsights {
	return v.value
}

func (v *NullablePrismInsights) Set(val *PrismInsights) {
	v.value = val
	v.isSet = true
}

func (v NullablePrismInsights) IsSet() bool {
	return v.isSet
}

func (v *NullablePrismInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrismInsights(val *PrismInsights) *NullablePrismInsights {
	return &NullablePrismInsights{value: val, isSet: true}
}

func (v NullablePrismInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrismInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


