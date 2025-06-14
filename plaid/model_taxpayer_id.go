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

// TaxpayerID Taxpayer ID of the individual receiving the paystub.
type TaxpayerID struct {
	// Type of ID, e.g. 'SSN'
	IdType NullableString `json:"id_type,omitempty"`
	// ID mask; i.e. last 4 digits of the taxpayer ID
	IdMask NullableString `json:"id_mask,omitempty"`
	// Last 4 digits of unique number of ID.
	Last4Digits NullableString `json:"last_4_digits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaxpayerID TaxpayerID

// NewTaxpayerID instantiates a new TaxpayerID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxpayerID() *TaxpayerID {
	this := TaxpayerID{}
	return &this
}

// NewTaxpayerIDWithDefaults instantiates a new TaxpayerID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxpayerIDWithDefaults() *TaxpayerID {
	this := TaxpayerID{}
	return &this
}

// GetIdType returns the IdType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxpayerID) GetIdType() string {
	if o == nil || o.IdType.Get() == nil {
		var ret string
		return ret
	}
	return *o.IdType.Get()
}

// GetIdTypeOk returns a tuple with the IdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxpayerID) GetIdTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdType.Get(), o.IdType.IsSet()
}

// HasIdType returns a boolean if a field has been set.
func (o *TaxpayerID) HasIdType() bool {
	if o != nil && o.IdType.IsSet() {
		return true
	}

	return false
}

// SetIdType gets a reference to the given NullableString and assigns it to the IdType field.
func (o *TaxpayerID) SetIdType(v string) {
	o.IdType.Set(&v)
}
// SetIdTypeNil sets the value for IdType to be an explicit nil
func (o *TaxpayerID) SetIdTypeNil() {
	o.IdType.Set(nil)
}

// UnsetIdType ensures that no value is present for IdType, not even an explicit nil
func (o *TaxpayerID) UnsetIdType() {
	o.IdType.Unset()
}

// GetIdMask returns the IdMask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxpayerID) GetIdMask() string {
	if o == nil || o.IdMask.Get() == nil {
		var ret string
		return ret
	}
	return *o.IdMask.Get()
}

// GetIdMaskOk returns a tuple with the IdMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxpayerID) GetIdMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdMask.Get(), o.IdMask.IsSet()
}

// HasIdMask returns a boolean if a field has been set.
func (o *TaxpayerID) HasIdMask() bool {
	if o != nil && o.IdMask.IsSet() {
		return true
	}

	return false
}

// SetIdMask gets a reference to the given NullableString and assigns it to the IdMask field.
func (o *TaxpayerID) SetIdMask(v string) {
	o.IdMask.Set(&v)
}
// SetIdMaskNil sets the value for IdMask to be an explicit nil
func (o *TaxpayerID) SetIdMaskNil() {
	o.IdMask.Set(nil)
}

// UnsetIdMask ensures that no value is present for IdMask, not even an explicit nil
func (o *TaxpayerID) UnsetIdMask() {
	o.IdMask.Unset()
}

// GetLast4Digits returns the Last4Digits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaxpayerID) GetLast4Digits() string {
	if o == nil || o.Last4Digits.Get() == nil {
		var ret string
		return ret
	}
	return *o.Last4Digits.Get()
}

// GetLast4DigitsOk returns a tuple with the Last4Digits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxpayerID) GetLast4DigitsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Last4Digits.Get(), o.Last4Digits.IsSet()
}

// HasLast4Digits returns a boolean if a field has been set.
func (o *TaxpayerID) HasLast4Digits() bool {
	if o != nil && o.Last4Digits.IsSet() {
		return true
	}

	return false
}

// SetLast4Digits gets a reference to the given NullableString and assigns it to the Last4Digits field.
func (o *TaxpayerID) SetLast4Digits(v string) {
	o.Last4Digits.Set(&v)
}
// SetLast4DigitsNil sets the value for Last4Digits to be an explicit nil
func (o *TaxpayerID) SetLast4DigitsNil() {
	o.Last4Digits.Set(nil)
}

// UnsetLast4Digits ensures that no value is present for Last4Digits, not even an explicit nil
func (o *TaxpayerID) UnsetLast4Digits() {
	o.Last4Digits.Unset()
}

func (o TaxpayerID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IdType.IsSet() {
		toSerialize["id_type"] = o.IdType.Get()
	}
	if o.IdMask.IsSet() {
		toSerialize["id_mask"] = o.IdMask.Get()
	}
	if o.Last4Digits.IsSet() {
		toSerialize["last_4_digits"] = o.Last4Digits.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TaxpayerID) UnmarshalJSON(bytes []byte) (err error) {
	varTaxpayerID := _TaxpayerID{}

	if err = json.Unmarshal(bytes, &varTaxpayerID); err == nil {
		*o = TaxpayerID(varTaxpayerID)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id_type")
		delete(additionalProperties, "id_mask")
		delete(additionalProperties, "last_4_digits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaxpayerID struct {
	value *TaxpayerID
	isSet bool
}

func (v NullableTaxpayerID) Get() *TaxpayerID {
	return v.value
}

func (v *NullableTaxpayerID) Set(val *TaxpayerID) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxpayerID) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxpayerID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxpayerID(val *TaxpayerID) *NullableTaxpayerID {
	return &NullableTaxpayerID{value: val, isSet: true}
}

func (v NullableTaxpayerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxpayerID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


