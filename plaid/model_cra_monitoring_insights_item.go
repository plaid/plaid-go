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
	"time"
)

// CraMonitoringInsightsItem An object representing a Monitoring Insights Item
type CraMonitoringInsightsItem struct {
	// The date and time when the specific insights were generated (per-item), in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (e.g. \"2018-04-12T03:32:11Z\").
	DateGenerated time.Time `json:"date_generated"`
	// The `item_id` of the Item associated with the insights
	ItemId string `json:"item_id"`
	Status MonitoringInsightsItemStatus `json:"status"`
	Insights NullableMonitoringInsights `json:"insights"`
	AdditionalProperties map[string]interface{}
}

type _CraMonitoringInsightsItem CraMonitoringInsightsItem

// NewCraMonitoringInsightsItem instantiates a new CraMonitoringInsightsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraMonitoringInsightsItem(dateGenerated time.Time, itemId string, status MonitoringInsightsItemStatus, insights NullableMonitoringInsights) *CraMonitoringInsightsItem {
	this := CraMonitoringInsightsItem{}
	this.DateGenerated = dateGenerated
	this.ItemId = itemId
	this.Status = status
	this.Insights = insights
	return &this
}

// NewCraMonitoringInsightsItemWithDefaults instantiates a new CraMonitoringInsightsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraMonitoringInsightsItemWithDefaults() *CraMonitoringInsightsItem {
	this := CraMonitoringInsightsItem{}
	return &this
}

// GetDateGenerated returns the DateGenerated field value
func (o *CraMonitoringInsightsItem) GetDateGenerated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateGenerated
}

// GetDateGeneratedOk returns a tuple with the DateGenerated field value
// and a boolean to check if the value has been set.
func (o *CraMonitoringInsightsItem) GetDateGeneratedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DateGenerated, true
}

// SetDateGenerated sets field value
func (o *CraMonitoringInsightsItem) SetDateGenerated(v time.Time) {
	o.DateGenerated = v
}

// GetItemId returns the ItemId field value
func (o *CraMonitoringInsightsItem) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *CraMonitoringInsightsItem) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *CraMonitoringInsightsItem) SetItemId(v string) {
	o.ItemId = v
}

// GetStatus returns the Status field value
func (o *CraMonitoringInsightsItem) GetStatus() MonitoringInsightsItemStatus {
	if o == nil {
		var ret MonitoringInsightsItemStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CraMonitoringInsightsItem) GetStatusOk() (*MonitoringInsightsItemStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CraMonitoringInsightsItem) SetStatus(v MonitoringInsightsItemStatus) {
	o.Status = v
}

// GetInsights returns the Insights field value
// If the value is explicit nil, the zero value for MonitoringInsights will be returned
func (o *CraMonitoringInsightsItem) GetInsights() MonitoringInsights {
	if o == nil || o.Insights.Get() == nil {
		var ret MonitoringInsights
		return ret
	}

	return *o.Insights.Get()
}

// GetInsightsOk returns a tuple with the Insights field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CraMonitoringInsightsItem) GetInsightsOk() (*MonitoringInsights, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Insights.Get(), o.Insights.IsSet()
}

// SetInsights sets field value
func (o *CraMonitoringInsightsItem) SetInsights(v MonitoringInsights) {
	o.Insights.Set(&v)
}

func (o CraMonitoringInsightsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date_generated"] = o.DateGenerated
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["insights"] = o.Insights.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraMonitoringInsightsItem) UnmarshalJSON(bytes []byte) (err error) {
	varCraMonitoringInsightsItem := _CraMonitoringInsightsItem{}

	if err = json.Unmarshal(bytes, &varCraMonitoringInsightsItem); err == nil {
		*o = CraMonitoringInsightsItem(varCraMonitoringInsightsItem)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "date_generated")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "insights")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraMonitoringInsightsItem struct {
	value *CraMonitoringInsightsItem
	isSet bool
}

func (v NullableCraMonitoringInsightsItem) Get() *CraMonitoringInsightsItem {
	return v.value
}

func (v *NullableCraMonitoringInsightsItem) Set(val *CraMonitoringInsightsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCraMonitoringInsightsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCraMonitoringInsightsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraMonitoringInsightsItem(val *CraMonitoringInsightsItem) *NullableCraMonitoringInsightsItem {
	return &NullableCraMonitoringInsightsItem{value: val, isSet: true}
}

func (v NullableCraMonitoringInsightsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraMonitoringInsightsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

