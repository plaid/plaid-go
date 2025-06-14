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

// DashboardUserGetResponse Account information associated with a team member with access to the Plaid dashboard.
type DashboardUserGetResponse struct {
	// ID of the associated user. To retrieve the email address or other details of the person corresponding to this id, use `/dashboard_user/get`.
	Id string `json:"id"`
	// An ISO8601 formatted timestamp.
	CreatedAt time.Time `json:"created_at"`
	// A valid email address. Must not have leading or trailing spaces and address must be RFC compliant. For more information, see [RFC 3696](https://datatracker.ietf.org/doc/html/rfc3696).
	EmailAddress string `json:"email_address"`
	Status DashboardUserStatus `json:"status"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _DashboardUserGetResponse DashboardUserGetResponse

// NewDashboardUserGetResponse instantiates a new DashboardUserGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardUserGetResponse(id string, createdAt time.Time, emailAddress string, status DashboardUserStatus, requestId string) *DashboardUserGetResponse {
	this := DashboardUserGetResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.EmailAddress = emailAddress
	this.Status = status
	this.RequestId = requestId
	return &this
}

// NewDashboardUserGetResponseWithDefaults instantiates a new DashboardUserGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardUserGetResponseWithDefaults() *DashboardUserGetResponse {
	this := DashboardUserGetResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DashboardUserGetResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DashboardUserGetResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DashboardUserGetResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DashboardUserGetResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DashboardUserGetResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DashboardUserGetResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *DashboardUserGetResponse) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *DashboardUserGetResponse) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *DashboardUserGetResponse) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetStatus returns the Status field value
func (o *DashboardUserGetResponse) GetStatus() DashboardUserStatus {
	if o == nil {
		var ret DashboardUserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DashboardUserGetResponse) GetStatusOk() (*DashboardUserStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DashboardUserGetResponse) SetStatus(v DashboardUserStatus) {
	o.Status = v
}

// GetRequestId returns the RequestId field value
func (o *DashboardUserGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *DashboardUserGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *DashboardUserGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o DashboardUserGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["email_address"] = o.EmailAddress
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DashboardUserGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDashboardUserGetResponse := _DashboardUserGetResponse{}

	if err = json.Unmarshal(bytes, &varDashboardUserGetResponse); err == nil {
		*o = DashboardUserGetResponse(varDashboardUserGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "email_address")
		delete(additionalProperties, "status")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDashboardUserGetResponse struct {
	value *DashboardUserGetResponse
	isSet bool
}

func (v NullableDashboardUserGetResponse) Get() *DashboardUserGetResponse {
	return v.value
}

func (v *NullableDashboardUserGetResponse) Set(val *DashboardUserGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardUserGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardUserGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardUserGetResponse(val *DashboardUserGetResponse) *NullableDashboardUserGetResponse {
	return &NullableDashboardUserGetResponse{value: val, isSet: true}
}

func (v NullableDashboardUserGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardUserGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


