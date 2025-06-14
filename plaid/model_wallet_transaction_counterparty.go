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

// WalletTransactionCounterparty An object representing the e-wallet transaction's counterparty
type WalletTransactionCounterparty struct {
	// The name of the counterparty
	Name string `json:"name"`
	Numbers WalletTransactionCounterpartyNumbers `json:"numbers"`
	Address NullablePaymentInitiationAddress `json:"address,omitempty"`
	// The counterparty's birthdate, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format.
	DateOfBirth NullableString `json:"date_of_birth,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransactionCounterparty WalletTransactionCounterparty

// NewWalletTransactionCounterparty instantiates a new WalletTransactionCounterparty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionCounterparty(name string, numbers WalletTransactionCounterpartyNumbers) *WalletTransactionCounterparty {
	this := WalletTransactionCounterparty{}
	this.Name = name
	this.Numbers = numbers
	return &this
}

// NewWalletTransactionCounterpartyWithDefaults instantiates a new WalletTransactionCounterparty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionCounterpartyWithDefaults() *WalletTransactionCounterparty {
	this := WalletTransactionCounterparty{}
	return &this
}

// GetName returns the Name field value
func (o *WalletTransactionCounterparty) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionCounterparty) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WalletTransactionCounterparty) SetName(v string) {
	o.Name = v
}

// GetNumbers returns the Numbers field value
func (o *WalletTransactionCounterparty) GetNumbers() WalletTransactionCounterpartyNumbers {
	if o == nil {
		var ret WalletTransactionCounterpartyNumbers
		return ret
	}

	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionCounterparty) GetNumbersOk() (*WalletTransactionCounterpartyNumbers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Numbers, true
}

// SetNumbers sets field value
func (o *WalletTransactionCounterparty) SetNumbers(v WalletTransactionCounterpartyNumbers) {
	o.Numbers = v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletTransactionCounterparty) GetAddress() PaymentInitiationAddress {
	if o == nil || o.Address.Get() == nil {
		var ret PaymentInitiationAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletTransactionCounterparty) GetAddressOk() (*PaymentInitiationAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *WalletTransactionCounterparty) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullablePaymentInitiationAddress and assigns it to the Address field.
func (o *WalletTransactionCounterparty) SetAddress(v PaymentInitiationAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *WalletTransactionCounterparty) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *WalletTransactionCounterparty) UnsetAddress() {
	o.Address.Unset()
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletTransactionCounterparty) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateOfBirth.Get()
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletTransactionCounterparty) GetDateOfBirthOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateOfBirth.Get(), o.DateOfBirth.IsSet()
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *WalletTransactionCounterparty) HasDateOfBirth() bool {
	if o != nil && o.DateOfBirth.IsSet() {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given NullableString and assigns it to the DateOfBirth field.
func (o *WalletTransactionCounterparty) SetDateOfBirth(v string) {
	o.DateOfBirth.Set(&v)
}
// SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil
func (o *WalletTransactionCounterparty) SetDateOfBirthNil() {
	o.DateOfBirth.Set(nil)
}

// UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
func (o *WalletTransactionCounterparty) UnsetDateOfBirth() {
	o.DateOfBirth.Unset()
}

func (o WalletTransactionCounterparty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["numbers"] = o.Numbers
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.DateOfBirth.IsSet() {
		toSerialize["date_of_birth"] = o.DateOfBirth.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletTransactionCounterparty) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransactionCounterparty := _WalletTransactionCounterparty{}

	if err = json.Unmarshal(bytes, &varWalletTransactionCounterparty); err == nil {
		*o = WalletTransactionCounterparty(varWalletTransactionCounterparty)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "numbers")
		delete(additionalProperties, "address")
		delete(additionalProperties, "date_of_birth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransactionCounterparty struct {
	value *WalletTransactionCounterparty
	isSet bool
}

func (v NullableWalletTransactionCounterparty) Get() *WalletTransactionCounterparty {
	return v.value
}

func (v *NullableWalletTransactionCounterparty) Set(val *WalletTransactionCounterparty) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionCounterparty) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionCounterparty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionCounterparty(val *WalletTransactionCounterparty) *NullableWalletTransactionCounterparty {
	return &NullableWalletTransactionCounterparty{value: val, isSet: true}
}

func (v NullableWalletTransactionCounterparty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionCounterparty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


