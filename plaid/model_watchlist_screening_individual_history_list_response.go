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

// WatchlistScreeningIndividualHistoryListResponse Paginated list of individual watchlist screenings.
type WatchlistScreeningIndividualHistoryListResponse struct {
	// List of individual watchlist screenings
	WatchlistScreenings []WatchlistScreeningIndividual `json:"watchlist_screenings"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningIndividualHistoryListResponse WatchlistScreeningIndividualHistoryListResponse

// NewWatchlistScreeningIndividualHistoryListResponse instantiates a new WatchlistScreeningIndividualHistoryListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividualHistoryListResponse(watchlistScreenings []WatchlistScreeningIndividual, nextCursor NullableString, requestId string) *WatchlistScreeningIndividualHistoryListResponse {
	this := WatchlistScreeningIndividualHistoryListResponse{}
	this.WatchlistScreenings = watchlistScreenings
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningIndividualHistoryListResponseWithDefaults instantiates a new WatchlistScreeningIndividualHistoryListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualHistoryListResponseWithDefaults() *WatchlistScreeningIndividualHistoryListResponse {
	this := WatchlistScreeningIndividualHistoryListResponse{}
	return &this
}

// GetWatchlistScreenings returns the WatchlistScreenings field value
func (o *WatchlistScreeningIndividualHistoryListResponse) GetWatchlistScreenings() []WatchlistScreeningIndividual {
	if o == nil {
		var ret []WatchlistScreeningIndividual
		return ret
	}

	return o.WatchlistScreenings
}

// GetWatchlistScreeningsOk returns a tuple with the WatchlistScreenings field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualHistoryListResponse) GetWatchlistScreeningsOk() (*[]WatchlistScreeningIndividual, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistScreenings, true
}

// SetWatchlistScreenings sets field value
func (o *WatchlistScreeningIndividualHistoryListResponse) SetWatchlistScreenings(v []WatchlistScreeningIndividual) {
	o.WatchlistScreenings = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningIndividualHistoryListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividualHistoryListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *WatchlistScreeningIndividualHistoryListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningIndividualHistoryListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualHistoryListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningIndividualHistoryListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningIndividualHistoryListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["watchlist_screenings"] = o.WatchlistScreenings
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

func (o *WatchlistScreeningIndividualHistoryListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningIndividualHistoryListResponse := _WatchlistScreeningIndividualHistoryListResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningIndividualHistoryListResponse); err == nil {
		*o = WatchlistScreeningIndividualHistoryListResponse(varWatchlistScreeningIndividualHistoryListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "watchlist_screenings")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningIndividualHistoryListResponse struct {
	value *WatchlistScreeningIndividualHistoryListResponse
	isSet bool
}

func (v NullableWatchlistScreeningIndividualHistoryListResponse) Get() *WatchlistScreeningIndividualHistoryListResponse {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualHistoryListResponse) Set(val *WatchlistScreeningIndividualHistoryListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualHistoryListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualHistoryListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualHistoryListResponse(val *WatchlistScreeningIndividualHistoryListResponse) *NullableWatchlistScreeningIndividualHistoryListResponse {
	return &NullableWatchlistScreeningIndividualHistoryListResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualHistoryListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualHistoryListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


