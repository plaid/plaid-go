/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.128.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaginatedIndividualWatchlistProgramListResponse Paginated list of individual watchlist screening programs
type PaginatedIndividualWatchlistProgramListResponse struct {
	// List of individual watchlist screening programs
	WatchlistPrograms []IndividualWatchlistProgram `json:"watchlist_programs"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedIndividualWatchlistProgramListResponse PaginatedIndividualWatchlistProgramListResponse

// NewPaginatedIndividualWatchlistProgramListResponse instantiates a new PaginatedIndividualWatchlistProgramListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedIndividualWatchlistProgramListResponse(watchlistPrograms []IndividualWatchlistProgram, nextCursor NullableString, requestId string) *PaginatedIndividualWatchlistProgramListResponse {
	this := PaginatedIndividualWatchlistProgramListResponse{}
	this.WatchlistPrograms = watchlistPrograms
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewPaginatedIndividualWatchlistProgramListResponseWithDefaults instantiates a new PaginatedIndividualWatchlistProgramListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedIndividualWatchlistProgramListResponseWithDefaults() *PaginatedIndividualWatchlistProgramListResponse {
	this := PaginatedIndividualWatchlistProgramListResponse{}
	return &this
}

// GetWatchlistPrograms returns the WatchlistPrograms field value
func (o *PaginatedIndividualWatchlistProgramListResponse) GetWatchlistPrograms() []IndividualWatchlistProgram {
	if o == nil {
		var ret []IndividualWatchlistProgram
		return ret
	}

	return o.WatchlistPrograms
}

// GetWatchlistProgramsOk returns a tuple with the WatchlistPrograms field value
// and a boolean to check if the value has been set.
func (o *PaginatedIndividualWatchlistProgramListResponse) GetWatchlistProgramsOk() (*[]IndividualWatchlistProgram, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistPrograms, true
}

// SetWatchlistPrograms sets field value
func (o *PaginatedIndividualWatchlistProgramListResponse) SetWatchlistPrograms(v []IndividualWatchlistProgram) {
	o.WatchlistPrograms = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginatedIndividualWatchlistProgramListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedIndividualWatchlistProgramListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *PaginatedIndividualWatchlistProgramListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *PaginatedIndividualWatchlistProgramListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaginatedIndividualWatchlistProgramListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaginatedIndividualWatchlistProgramListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaginatedIndividualWatchlistProgramListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["watchlist_programs"] = o.WatchlistPrograms
	}
	if true {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaginatedIndividualWatchlistProgramListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaginatedIndividualWatchlistProgramListResponse := _PaginatedIndividualWatchlistProgramListResponse{}

	if err = json.Unmarshal(bytes, &varPaginatedIndividualWatchlistProgramListResponse); err == nil {
		*o = PaginatedIndividualWatchlistProgramListResponse(varPaginatedIndividualWatchlistProgramListResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "watchlist_programs")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedIndividualWatchlistProgramListResponse struct {
	value *PaginatedIndividualWatchlistProgramListResponse
	isSet bool
}

func (v NullablePaginatedIndividualWatchlistProgramListResponse) Get() *PaginatedIndividualWatchlistProgramListResponse {
	return v.value
}

func (v *NullablePaginatedIndividualWatchlistProgramListResponse) Set(val *PaginatedIndividualWatchlistProgramListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedIndividualWatchlistProgramListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedIndividualWatchlistProgramListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedIndividualWatchlistProgramListResponse(val *PaginatedIndividualWatchlistProgramListResponse) *NullablePaginatedIndividualWatchlistProgramListResponse {
	return &NullablePaginatedIndividualWatchlistProgramListResponse{value: val, isSet: true}
}

func (v NullablePaginatedIndividualWatchlistProgramListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedIndividualWatchlistProgramListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

