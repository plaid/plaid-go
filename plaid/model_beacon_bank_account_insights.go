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

// BeaconBankAccountInsights Bank Account Insights encapsulate the risk insights for a single Bank Account linked to an Item that is assocaited with a Beacon User.
type BeaconBankAccountInsights struct {
	// The Plaid `account_id`
	AccountId string `json:"account_id"`
	Type AccountType `json:"type"`
	Subtype NullableAccountSubtype `json:"subtype"`
	Attributes BeaconAccountRiskAttributes `json:"attributes"`
	AdditionalProperties map[string]interface{}
}

type _BeaconBankAccountInsights BeaconBankAccountInsights

// NewBeaconBankAccountInsights instantiates a new BeaconBankAccountInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeaconBankAccountInsights(accountId string, type_ AccountType, subtype NullableAccountSubtype, attributes BeaconAccountRiskAttributes) *BeaconBankAccountInsights {
	this := BeaconBankAccountInsights{}
	this.AccountId = accountId
	this.Type = type_
	this.Subtype = subtype
	this.Attributes = attributes
	return &this
}

// NewBeaconBankAccountInsightsWithDefaults instantiates a new BeaconBankAccountInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeaconBankAccountInsightsWithDefaults() *BeaconBankAccountInsights {
	this := BeaconBankAccountInsights{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *BeaconBankAccountInsights) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *BeaconBankAccountInsights) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *BeaconBankAccountInsights) SetAccountId(v string) {
	o.AccountId = v
}

// GetType returns the Type field value
func (o *BeaconBankAccountInsights) GetType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BeaconBankAccountInsights) GetTypeOk() (*AccountType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BeaconBankAccountInsights) SetType(v AccountType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value
// If the value is explicit nil, the zero value for AccountSubtype will be returned
func (o *BeaconBankAccountInsights) GetSubtype() AccountSubtype {
	if o == nil || o.Subtype.Get() == nil {
		var ret AccountSubtype
		return ret
	}

	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BeaconBankAccountInsights) GetSubtypeOk() (*AccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// SetSubtype sets field value
func (o *BeaconBankAccountInsights) SetSubtype(v AccountSubtype) {
	o.Subtype.Set(&v)
}

// GetAttributes returns the Attributes field value
func (o *BeaconBankAccountInsights) GetAttributes() BeaconAccountRiskAttributes {
	if o == nil {
		var ret BeaconAccountRiskAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BeaconBankAccountInsights) GetAttributesOk() (*BeaconAccountRiskAttributes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BeaconBankAccountInsights) SetAttributes(v BeaconAccountRiskAttributes) {
	o.Attributes = v
}

func (o BeaconBankAccountInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BeaconBankAccountInsights) UnmarshalJSON(bytes []byte) (err error) {
	varBeaconBankAccountInsights := _BeaconBankAccountInsights{}

	if err = json.Unmarshal(bytes, &varBeaconBankAccountInsights); err == nil {
		*o = BeaconBankAccountInsights(varBeaconBankAccountInsights)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBeaconBankAccountInsights struct {
	value *BeaconBankAccountInsights
	isSet bool
}

func (v NullableBeaconBankAccountInsights) Get() *BeaconBankAccountInsights {
	return v.value
}

func (v *NullableBeaconBankAccountInsights) Set(val *BeaconBankAccountInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconBankAccountInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconBankAccountInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconBankAccountInsights(val *BeaconBankAccountInsights) *NullableBeaconBankAccountInsights {
	return &NullableBeaconBankAccountInsights{value: val, isSet: true}
}

func (v NullableBeaconBankAccountInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconBankAccountInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

