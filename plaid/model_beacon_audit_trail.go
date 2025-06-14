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
	"time"
)

// BeaconAuditTrail Information about the last change made to the parent object specifying what caused the change as well as when it occurred.
type BeaconAuditTrail struct {
	Source BeaconAuditTrailSource `json:"source"`
	// ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.
	DashboardUserId NullableString `json:"dashboard_user_id"`
	// An ISO8601 formatted timestamp.
	Timestamp time.Time `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _BeaconAuditTrail BeaconAuditTrail

// NewBeaconAuditTrail instantiates a new BeaconAuditTrail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeaconAuditTrail(source BeaconAuditTrailSource, dashboardUserId NullableString, timestamp time.Time) *BeaconAuditTrail {
	this := BeaconAuditTrail{}
	this.Source = source
	this.DashboardUserId = dashboardUserId
	this.Timestamp = timestamp
	return &this
}

// NewBeaconAuditTrailWithDefaults instantiates a new BeaconAuditTrail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeaconAuditTrailWithDefaults() *BeaconAuditTrail {
	this := BeaconAuditTrail{}
	return &this
}

// GetSource returns the Source field value
func (o *BeaconAuditTrail) GetSource() BeaconAuditTrailSource {
	if o == nil {
		var ret BeaconAuditTrailSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *BeaconAuditTrail) GetSourceOk() (*BeaconAuditTrailSource, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *BeaconAuditTrail) SetSource(v BeaconAuditTrailSource) {
	o.Source = v
}

// GetDashboardUserId returns the DashboardUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BeaconAuditTrail) GetDashboardUserId() string {
	if o == nil || o.DashboardUserId.Get() == nil {
		var ret string
		return ret
	}

	return *o.DashboardUserId.Get()
}

// GetDashboardUserIdOk returns a tuple with the DashboardUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BeaconAuditTrail) GetDashboardUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DashboardUserId.Get(), o.DashboardUserId.IsSet()
}

// SetDashboardUserId sets field value
func (o *BeaconAuditTrail) SetDashboardUserId(v string) {
	o.DashboardUserId.Set(&v)
}

// GetTimestamp returns the Timestamp field value
func (o *BeaconAuditTrail) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *BeaconAuditTrail) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *BeaconAuditTrail) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o BeaconAuditTrail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["dashboard_user_id"] = o.DashboardUserId.Get()
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BeaconAuditTrail) UnmarshalJSON(bytes []byte) (err error) {
	varBeaconAuditTrail := _BeaconAuditTrail{}

	if err = json.Unmarshal(bytes, &varBeaconAuditTrail); err == nil {
		*o = BeaconAuditTrail(varBeaconAuditTrail)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		delete(additionalProperties, "dashboard_user_id")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBeaconAuditTrail struct {
	value *BeaconAuditTrail
	isSet bool
}

func (v NullableBeaconAuditTrail) Get() *BeaconAuditTrail {
	return v.value
}

func (v *NullableBeaconAuditTrail) Set(val *BeaconAuditTrail) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconAuditTrail) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconAuditTrail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconAuditTrail(val *BeaconAuditTrail) *NullableBeaconAuditTrail {
	return &NullableBeaconAuditTrail{value: val, isSet: true}
}

func (v NullableBeaconAuditTrail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconAuditTrail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


