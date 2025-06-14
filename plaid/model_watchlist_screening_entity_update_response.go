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

// WatchlistScreeningEntityUpdateResponse The entity screening object allows you to represent an entity in your system, update its profile, and search for it on various watchlists. Note: Rejected entity screenings will not receive new hits, regardless of entity program configuration.
type WatchlistScreeningEntityUpdateResponse struct {
	// ID of the associated entity screening.
	Id string `json:"id"`
	SearchTerms EntityWatchlistScreeningSearchTerms `json:"search_terms"`
	// ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.
	Assignee NullableString `json:"assignee"`
	Status WatchlistScreeningStatus `json:"status"`
	// A unique ID that identifies the end user in your system. This ID can also be used to associate user-specific data from other Plaid products. Financial Account Matching requires this field and the `/link/token/create` `client_user_id` to be consistent. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.
	ClientUserId NullableString `json:"client_user_id"`
	AuditTrail WatchlistScreeningAuditTrail `json:"audit_trail"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningEntityUpdateResponse WatchlistScreeningEntityUpdateResponse

// NewWatchlistScreeningEntityUpdateResponse instantiates a new WatchlistScreeningEntityUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningEntityUpdateResponse(id string, searchTerms EntityWatchlistScreeningSearchTerms, assignee NullableString, status WatchlistScreeningStatus, clientUserId NullableString, auditTrail WatchlistScreeningAuditTrail, requestId string) *WatchlistScreeningEntityUpdateResponse {
	this := WatchlistScreeningEntityUpdateResponse{}
	this.Id = id
	this.SearchTerms = searchTerms
	this.Assignee = assignee
	this.Status = status
	this.ClientUserId = clientUserId
	this.AuditTrail = auditTrail
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningEntityUpdateResponseWithDefaults instantiates a new WatchlistScreeningEntityUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningEntityUpdateResponseWithDefaults() *WatchlistScreeningEntityUpdateResponse {
	this := WatchlistScreeningEntityUpdateResponse{}
	return &this
}

// GetId returns the Id field value
func (o *WatchlistScreeningEntityUpdateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityUpdateResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WatchlistScreeningEntityUpdateResponse) SetId(v string) {
	o.Id = v
}

// GetSearchTerms returns the SearchTerms field value
func (o *WatchlistScreeningEntityUpdateResponse) GetSearchTerms() EntityWatchlistScreeningSearchTerms {
	if o == nil {
		var ret EntityWatchlistScreeningSearchTerms
		return ret
	}

	return o.SearchTerms
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityUpdateResponse) GetSearchTermsOk() (*EntityWatchlistScreeningSearchTerms, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchTerms, true
}

// SetSearchTerms sets field value
func (o *WatchlistScreeningEntityUpdateResponse) SetSearchTerms(v EntityWatchlistScreeningSearchTerms) {
	o.SearchTerms = v
}

// GetAssignee returns the Assignee field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningEntityUpdateResponse) GetAssignee() string {
	if o == nil || o.Assignee.Get() == nil {
		var ret string
		return ret
	}

	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningEntityUpdateResponse) GetAssigneeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// SetAssignee sets field value
func (o *WatchlistScreeningEntityUpdateResponse) SetAssignee(v string) {
	o.Assignee.Set(&v)
}

// GetStatus returns the Status field value
func (o *WatchlistScreeningEntityUpdateResponse) GetStatus() WatchlistScreeningStatus {
	if o == nil {
		var ret WatchlistScreeningStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityUpdateResponse) GetStatusOk() (*WatchlistScreeningStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WatchlistScreeningEntityUpdateResponse) SetStatus(v WatchlistScreeningStatus) {
	o.Status = v
}

// GetClientUserId returns the ClientUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningEntityUpdateResponse) GetClientUserId() string {
	if o == nil || o.ClientUserId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientUserId.Get()
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningEntityUpdateResponse) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientUserId.Get(), o.ClientUserId.IsSet()
}

// SetClientUserId sets field value
func (o *WatchlistScreeningEntityUpdateResponse) SetClientUserId(v string) {
	o.ClientUserId.Set(&v)
}

// GetAuditTrail returns the AuditTrail field value
func (o *WatchlistScreeningEntityUpdateResponse) GetAuditTrail() WatchlistScreeningAuditTrail {
	if o == nil {
		var ret WatchlistScreeningAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityUpdateResponse) GetAuditTrailOk() (*WatchlistScreeningAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *WatchlistScreeningEntityUpdateResponse) SetAuditTrail(v WatchlistScreeningAuditTrail) {
	o.AuditTrail = v
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningEntityUpdateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityUpdateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningEntityUpdateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningEntityUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["search_terms"] = o.SearchTerms
	}
	if true {
		toSerialize["assignee"] = o.Assignee.Get()
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId.Get()
	}
	if true {
		toSerialize["audit_trail"] = o.AuditTrail
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningEntityUpdateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningEntityUpdateResponse := _WatchlistScreeningEntityUpdateResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningEntityUpdateResponse); err == nil {
		*o = WatchlistScreeningEntityUpdateResponse(varWatchlistScreeningEntityUpdateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "search_terms")
		delete(additionalProperties, "assignee")
		delete(additionalProperties, "status")
		delete(additionalProperties, "client_user_id")
		delete(additionalProperties, "audit_trail")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningEntityUpdateResponse struct {
	value *WatchlistScreeningEntityUpdateResponse
	isSet bool
}

func (v NullableWatchlistScreeningEntityUpdateResponse) Get() *WatchlistScreeningEntityUpdateResponse {
	return v.value
}

func (v *NullableWatchlistScreeningEntityUpdateResponse) Set(val *WatchlistScreeningEntityUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningEntityUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningEntityUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningEntityUpdateResponse(val *WatchlistScreeningEntityUpdateResponse) *NullableWatchlistScreeningEntityUpdateResponse {
	return &NullableWatchlistScreeningEntityUpdateResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningEntityUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningEntityUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


