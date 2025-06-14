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

// RiskCheckSyntheticIdentity Field containing the data used in determining the outcome of the synthetic identity risk check.  Contains the following fields:  `score` - A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.
type RiskCheckSyntheticIdentity struct {
	// A score from 0 to 100 indicating the likelihood that the user is a synthetic identity.
	Score int32 `json:"score"`
	RiskLevel *RiskLevel `json:"risk_level,omitempty"`
	FirstPartySyntheticFraud NullableSyntheticFraud `json:"first_party_synthetic_fraud,omitempty"`
	ThirdPartySyntheticFraud NullableSyntheticFraud `json:"third_party_synthetic_fraud,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RiskCheckSyntheticIdentity RiskCheckSyntheticIdentity

// NewRiskCheckSyntheticIdentity instantiates a new RiskCheckSyntheticIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskCheckSyntheticIdentity(score int32) *RiskCheckSyntheticIdentity {
	this := RiskCheckSyntheticIdentity{}
	this.Score = score
	return &this
}

// NewRiskCheckSyntheticIdentityWithDefaults instantiates a new RiskCheckSyntheticIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskCheckSyntheticIdentityWithDefaults() *RiskCheckSyntheticIdentity {
	this := RiskCheckSyntheticIdentity{}
	return &this
}

// GetScore returns the Score field value
func (o *RiskCheckSyntheticIdentity) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *RiskCheckSyntheticIdentity) GetScoreOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *RiskCheckSyntheticIdentity) SetScore(v int32) {
	o.Score = v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *RiskCheckSyntheticIdentity) GetRiskLevel() RiskLevel {
	if o == nil || o.RiskLevel == nil {
		var ret RiskLevel
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskCheckSyntheticIdentity) GetRiskLevelOk() (*RiskLevel, bool) {
	if o == nil || o.RiskLevel == nil {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *RiskCheckSyntheticIdentity) HasRiskLevel() bool {
	if o != nil && o.RiskLevel != nil {
		return true
	}

	return false
}

// SetRiskLevel gets a reference to the given RiskLevel and assigns it to the RiskLevel field.
func (o *RiskCheckSyntheticIdentity) SetRiskLevel(v RiskLevel) {
	o.RiskLevel = &v
}

// GetFirstPartySyntheticFraud returns the FirstPartySyntheticFraud field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskCheckSyntheticIdentity) GetFirstPartySyntheticFraud() SyntheticFraud {
	if o == nil || o.FirstPartySyntheticFraud.Get() == nil {
		var ret SyntheticFraud
		return ret
	}
	return *o.FirstPartySyntheticFraud.Get()
}

// GetFirstPartySyntheticFraudOk returns a tuple with the FirstPartySyntheticFraud field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskCheckSyntheticIdentity) GetFirstPartySyntheticFraudOk() (*SyntheticFraud, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstPartySyntheticFraud.Get(), o.FirstPartySyntheticFraud.IsSet()
}

// HasFirstPartySyntheticFraud returns a boolean if a field has been set.
func (o *RiskCheckSyntheticIdentity) HasFirstPartySyntheticFraud() bool {
	if o != nil && o.FirstPartySyntheticFraud.IsSet() {
		return true
	}

	return false
}

// SetFirstPartySyntheticFraud gets a reference to the given NullableSyntheticFraud and assigns it to the FirstPartySyntheticFraud field.
func (o *RiskCheckSyntheticIdentity) SetFirstPartySyntheticFraud(v SyntheticFraud) {
	o.FirstPartySyntheticFraud.Set(&v)
}
// SetFirstPartySyntheticFraudNil sets the value for FirstPartySyntheticFraud to be an explicit nil
func (o *RiskCheckSyntheticIdentity) SetFirstPartySyntheticFraudNil() {
	o.FirstPartySyntheticFraud.Set(nil)
}

// UnsetFirstPartySyntheticFraud ensures that no value is present for FirstPartySyntheticFraud, not even an explicit nil
func (o *RiskCheckSyntheticIdentity) UnsetFirstPartySyntheticFraud() {
	o.FirstPartySyntheticFraud.Unset()
}

// GetThirdPartySyntheticFraud returns the ThirdPartySyntheticFraud field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskCheckSyntheticIdentity) GetThirdPartySyntheticFraud() SyntheticFraud {
	if o == nil || o.ThirdPartySyntheticFraud.Get() == nil {
		var ret SyntheticFraud
		return ret
	}
	return *o.ThirdPartySyntheticFraud.Get()
}

// GetThirdPartySyntheticFraudOk returns a tuple with the ThirdPartySyntheticFraud field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskCheckSyntheticIdentity) GetThirdPartySyntheticFraudOk() (*SyntheticFraud, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ThirdPartySyntheticFraud.Get(), o.ThirdPartySyntheticFraud.IsSet()
}

// HasThirdPartySyntheticFraud returns a boolean if a field has been set.
func (o *RiskCheckSyntheticIdentity) HasThirdPartySyntheticFraud() bool {
	if o != nil && o.ThirdPartySyntheticFraud.IsSet() {
		return true
	}

	return false
}

// SetThirdPartySyntheticFraud gets a reference to the given NullableSyntheticFraud and assigns it to the ThirdPartySyntheticFraud field.
func (o *RiskCheckSyntheticIdentity) SetThirdPartySyntheticFraud(v SyntheticFraud) {
	o.ThirdPartySyntheticFraud.Set(&v)
}
// SetThirdPartySyntheticFraudNil sets the value for ThirdPartySyntheticFraud to be an explicit nil
func (o *RiskCheckSyntheticIdentity) SetThirdPartySyntheticFraudNil() {
	o.ThirdPartySyntheticFraud.Set(nil)
}

// UnsetThirdPartySyntheticFraud ensures that no value is present for ThirdPartySyntheticFraud, not even an explicit nil
func (o *RiskCheckSyntheticIdentity) UnsetThirdPartySyntheticFraud() {
	o.ThirdPartySyntheticFraud.Unset()
}

func (o RiskCheckSyntheticIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["score"] = o.Score
	}
	if o.RiskLevel != nil {
		toSerialize["risk_level"] = o.RiskLevel
	}
	if o.FirstPartySyntheticFraud.IsSet() {
		toSerialize["first_party_synthetic_fraud"] = o.FirstPartySyntheticFraud.Get()
	}
	if o.ThirdPartySyntheticFraud.IsSet() {
		toSerialize["third_party_synthetic_fraud"] = o.ThirdPartySyntheticFraud.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskCheckSyntheticIdentity) UnmarshalJSON(bytes []byte) (err error) {
	varRiskCheckSyntheticIdentity := _RiskCheckSyntheticIdentity{}

	if err = json.Unmarshal(bytes, &varRiskCheckSyntheticIdentity); err == nil {
		*o = RiskCheckSyntheticIdentity(varRiskCheckSyntheticIdentity)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "score")
		delete(additionalProperties, "risk_level")
		delete(additionalProperties, "first_party_synthetic_fraud")
		delete(additionalProperties, "third_party_synthetic_fraud")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskCheckSyntheticIdentity struct {
	value *RiskCheckSyntheticIdentity
	isSet bool
}

func (v NullableRiskCheckSyntheticIdentity) Get() *RiskCheckSyntheticIdentity {
	return v.value
}

func (v *NullableRiskCheckSyntheticIdentity) Set(val *RiskCheckSyntheticIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckSyntheticIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckSyntheticIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckSyntheticIdentity(val *RiskCheckSyntheticIdentity) *NullableRiskCheckSyntheticIdentity {
	return &NullableRiskCheckSyntheticIdentity{value: val, isSet: true}
}

func (v NullableRiskCheckSyntheticIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckSyntheticIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


