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

// WatchlistScreeningEntityHitListResponse Paginated list of entity watchlist screening hits
type WatchlistScreeningEntityHitListResponse struct {
	// List of entity watchlist screening hits
	EntityWatchlistScreeningHits []EntityWatchlistScreeningHit `json:"entity_watchlist_screening_hits"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningEntityHitListResponse WatchlistScreeningEntityHitListResponse

// NewWatchlistScreeningEntityHitListResponse instantiates a new WatchlistScreeningEntityHitListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningEntityHitListResponse(entityWatchlistScreeningHits []EntityWatchlistScreeningHit, nextCursor NullableString, requestId string) *WatchlistScreeningEntityHitListResponse {
	this := WatchlistScreeningEntityHitListResponse{}
	this.EntityWatchlistScreeningHits = entityWatchlistScreeningHits
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningEntityHitListResponseWithDefaults instantiates a new WatchlistScreeningEntityHitListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningEntityHitListResponseWithDefaults() *WatchlistScreeningEntityHitListResponse {
	this := WatchlistScreeningEntityHitListResponse{}
	return &this
}

// GetEntityWatchlistScreeningHits returns the EntityWatchlistScreeningHits field value
func (o *WatchlistScreeningEntityHitListResponse) GetEntityWatchlistScreeningHits() []EntityWatchlistScreeningHit {
	if o == nil {
		var ret []EntityWatchlistScreeningHit
		return ret
	}

	return o.EntityWatchlistScreeningHits
}

// GetEntityWatchlistScreeningHitsOk returns a tuple with the EntityWatchlistScreeningHits field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityHitListResponse) GetEntityWatchlistScreeningHitsOk() (*[]EntityWatchlistScreeningHit, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistScreeningHits, true
}

// SetEntityWatchlistScreeningHits sets field value
func (o *WatchlistScreeningEntityHitListResponse) SetEntityWatchlistScreeningHits(v []EntityWatchlistScreeningHit) {
	o.EntityWatchlistScreeningHits = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningEntityHitListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningEntityHitListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *WatchlistScreeningEntityHitListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningEntityHitListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityHitListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningEntityHitListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningEntityHitListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_watchlist_screening_hits"] = o.EntityWatchlistScreeningHits
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

func (o *WatchlistScreeningEntityHitListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningEntityHitListResponse := _WatchlistScreeningEntityHitListResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningEntityHitListResponse); err == nil {
		*o = WatchlistScreeningEntityHitListResponse(varWatchlistScreeningEntityHitListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_watchlist_screening_hits")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningEntityHitListResponse struct {
	value *WatchlistScreeningEntityHitListResponse
	isSet bool
}

func (v NullableWatchlistScreeningEntityHitListResponse) Get() *WatchlistScreeningEntityHitListResponse {
	return v.value
}

func (v *NullableWatchlistScreeningEntityHitListResponse) Set(val *WatchlistScreeningEntityHitListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningEntityHitListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningEntityHitListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningEntityHitListResponse(val *WatchlistScreeningEntityHitListResponse) *NullableWatchlistScreeningEntityHitListResponse {
	return &NullableWatchlistScreeningEntityHitListResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningEntityHitListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningEntityHitListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


