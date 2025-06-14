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

// WatchlistScreeningIndividualHitListResponse Paginated list of individual watchlist screening hits
type WatchlistScreeningIndividualHitListResponse struct {
	// List of individual watchlist screening hits
	WatchlistScreeningHits []WatchlistScreeningHit `json:"watchlist_screening_hits"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningIndividualHitListResponse WatchlistScreeningIndividualHitListResponse

// NewWatchlistScreeningIndividualHitListResponse instantiates a new WatchlistScreeningIndividualHitListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividualHitListResponse(watchlistScreeningHits []WatchlistScreeningHit, nextCursor NullableString, requestId string) *WatchlistScreeningIndividualHitListResponse {
	this := WatchlistScreeningIndividualHitListResponse{}
	this.WatchlistScreeningHits = watchlistScreeningHits
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningIndividualHitListResponseWithDefaults instantiates a new WatchlistScreeningIndividualHitListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualHitListResponseWithDefaults() *WatchlistScreeningIndividualHitListResponse {
	this := WatchlistScreeningIndividualHitListResponse{}
	return &this
}

// GetWatchlistScreeningHits returns the WatchlistScreeningHits field value
func (o *WatchlistScreeningIndividualHitListResponse) GetWatchlistScreeningHits() []WatchlistScreeningHit {
	if o == nil {
		var ret []WatchlistScreeningHit
		return ret
	}

	return o.WatchlistScreeningHits
}

// GetWatchlistScreeningHitsOk returns a tuple with the WatchlistScreeningHits field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualHitListResponse) GetWatchlistScreeningHitsOk() (*[]WatchlistScreeningHit, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistScreeningHits, true
}

// SetWatchlistScreeningHits sets field value
func (o *WatchlistScreeningIndividualHitListResponse) SetWatchlistScreeningHits(v []WatchlistScreeningHit) {
	o.WatchlistScreeningHits = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningIndividualHitListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividualHitListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *WatchlistScreeningIndividualHitListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningIndividualHitListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualHitListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningIndividualHitListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningIndividualHitListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["watchlist_screening_hits"] = o.WatchlistScreeningHits
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

func (o *WatchlistScreeningIndividualHitListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningIndividualHitListResponse := _WatchlistScreeningIndividualHitListResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningIndividualHitListResponse); err == nil {
		*o = WatchlistScreeningIndividualHitListResponse(varWatchlistScreeningIndividualHitListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "watchlist_screening_hits")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningIndividualHitListResponse struct {
	value *WatchlistScreeningIndividualHitListResponse
	isSet bool
}

func (v NullableWatchlistScreeningIndividualHitListResponse) Get() *WatchlistScreeningIndividualHitListResponse {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualHitListResponse) Set(val *WatchlistScreeningIndividualHitListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualHitListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualHitListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualHitListResponse(val *WatchlistScreeningIndividualHitListResponse) *NullableWatchlistScreeningIndividualHitListResponse {
	return &NullableWatchlistScreeningIndividualHitListResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualHitListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualHitListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


