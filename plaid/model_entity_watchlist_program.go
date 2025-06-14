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

// EntityWatchlistProgram A program that configures the active lists, search parameters, and other behavior for initial and ongoing screening of entities.
type EntityWatchlistProgram struct {
	// ID of the associated entity program.
	Id string `json:"id"`
	// An ISO8601 formatted timestamp.
	CreatedAt time.Time `json:"created_at"`
	// Indicator specifying whether the program is enabled and will perform daily rescans.
	IsRescanningEnabled bool `json:"is_rescanning_enabled"`
	// Watchlists enabled for the associated program
	ListsEnabled []EntityWatchlistCode `json:"lists_enabled"`
	// A name for the entity program to define its purpose. For example, \"High Risk Organizations\" or \"Applicants\".
	Name string `json:"name"`
	NameSensitivity ProgramNameSensitivity `json:"name_sensitivity"`
	AuditTrail WatchlistScreeningAuditTrail `json:"audit_trail"`
	// Archived programs are read-only and cannot screen new customers nor participate in ongoing monitoring.
	IsArchived bool `json:"is_archived"`
	AdditionalProperties map[string]interface{}
}

type _EntityWatchlistProgram EntityWatchlistProgram

// NewEntityWatchlistProgram instantiates a new EntityWatchlistProgram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityWatchlistProgram(id string, createdAt time.Time, isRescanningEnabled bool, listsEnabled []EntityWatchlistCode, name string, nameSensitivity ProgramNameSensitivity, auditTrail WatchlistScreeningAuditTrail, isArchived bool) *EntityWatchlistProgram {
	this := EntityWatchlistProgram{}
	this.Id = id
	this.CreatedAt = createdAt
	this.IsRescanningEnabled = isRescanningEnabled
	this.ListsEnabled = listsEnabled
	this.Name = name
	this.NameSensitivity = nameSensitivity
	this.AuditTrail = auditTrail
	this.IsArchived = isArchived
	return &this
}

// NewEntityWatchlistProgramWithDefaults instantiates a new EntityWatchlistProgram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityWatchlistProgramWithDefaults() *EntityWatchlistProgram {
	this := EntityWatchlistProgram{}
	return &this
}

// GetId returns the Id field value
func (o *EntityWatchlistProgram) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistProgram) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntityWatchlistProgram) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EntityWatchlistProgram) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistProgram) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EntityWatchlistProgram) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetIsRescanningEnabled returns the IsRescanningEnabled field value
func (o *EntityWatchlistProgram) GetIsRescanningEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRescanningEnabled
}

// GetIsRescanningEnabledOk returns a tuple with the IsRescanningEnabled field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistProgram) GetIsRescanningEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsRescanningEnabled, true
}

// SetIsRescanningEnabled sets field value
func (o *EntityWatchlistProgram) SetIsRescanningEnabled(v bool) {
	o.IsRescanningEnabled = v
}

// GetListsEnabled returns the ListsEnabled field value
func (o *EntityWatchlistProgram) GetListsEnabled() []EntityWatchlistCode {
	if o == nil {
		var ret []EntityWatchlistCode
		return ret
	}

	return o.ListsEnabled
}

// GetListsEnabledOk returns a tuple with the ListsEnabled field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistProgram) GetListsEnabledOk() (*[]EntityWatchlistCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ListsEnabled, true
}

// SetListsEnabled sets field value
func (o *EntityWatchlistProgram) SetListsEnabled(v []EntityWatchlistCode) {
	o.ListsEnabled = v
}

// GetName returns the Name field value
func (o *EntityWatchlistProgram) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistProgram) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntityWatchlistProgram) SetName(v string) {
	o.Name = v
}

// GetNameSensitivity returns the NameSensitivity field value
func (o *EntityWatchlistProgram) GetNameSensitivity() ProgramNameSensitivity {
	if o == nil {
		var ret ProgramNameSensitivity
		return ret
	}

	return o.NameSensitivity
}

// GetNameSensitivityOk returns a tuple with the NameSensitivity field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistProgram) GetNameSensitivityOk() (*ProgramNameSensitivity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NameSensitivity, true
}

// SetNameSensitivity sets field value
func (o *EntityWatchlistProgram) SetNameSensitivity(v ProgramNameSensitivity) {
	o.NameSensitivity = v
}

// GetAuditTrail returns the AuditTrail field value
func (o *EntityWatchlistProgram) GetAuditTrail() WatchlistScreeningAuditTrail {
	if o == nil {
		var ret WatchlistScreeningAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistProgram) GetAuditTrailOk() (*WatchlistScreeningAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *EntityWatchlistProgram) SetAuditTrail(v WatchlistScreeningAuditTrail) {
	o.AuditTrail = v
}

// GetIsArchived returns the IsArchived field value
func (o *EntityWatchlistProgram) GetIsArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value
// and a boolean to check if the value has been set.
func (o *EntityWatchlistProgram) GetIsArchivedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsArchived, true
}

// SetIsArchived sets field value
func (o *EntityWatchlistProgram) SetIsArchived(v bool) {
	o.IsArchived = v
}

func (o EntityWatchlistProgram) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["is_rescanning_enabled"] = o.IsRescanningEnabled
	}
	if true {
		toSerialize["lists_enabled"] = o.ListsEnabled
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["name_sensitivity"] = o.NameSensitivity
	}
	if true {
		toSerialize["audit_trail"] = o.AuditTrail
	}
	if true {
		toSerialize["is_archived"] = o.IsArchived
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityWatchlistProgram) UnmarshalJSON(bytes []byte) (err error) {
	varEntityWatchlistProgram := _EntityWatchlistProgram{}

	if err = json.Unmarshal(bytes, &varEntityWatchlistProgram); err == nil {
		*o = EntityWatchlistProgram(varEntityWatchlistProgram)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "is_rescanning_enabled")
		delete(additionalProperties, "lists_enabled")
		delete(additionalProperties, "name")
		delete(additionalProperties, "name_sensitivity")
		delete(additionalProperties, "audit_trail")
		delete(additionalProperties, "is_archived")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityWatchlistProgram struct {
	value *EntityWatchlistProgram
	isSet bool
}

func (v NullableEntityWatchlistProgram) Get() *EntityWatchlistProgram {
	return v.value
}

func (v *NullableEntityWatchlistProgram) Set(val *EntityWatchlistProgram) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityWatchlistProgram) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityWatchlistProgram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityWatchlistProgram(val *EntityWatchlistProgram) *NullableEntityWatchlistProgram {
	return &NullableEntityWatchlistProgram{value: val, isSet: true}
}

func (v NullableEntityWatchlistProgram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityWatchlistProgram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


