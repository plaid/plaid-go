/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.214.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// PaymentProfileGetResponse PaymentProfileGetResponse defines the response schema for `/payment_profile/get`
type PaymentProfileGetResponse struct {
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the last time the given Payment Profile was updated at
	UpdatedAt time.Time `json:"updated_at"`
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the time the given Payment Profile was created at
	CreatedAt time.Time `json:"created_at"`
	// Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (`YYYY-MM-DDTHH:mm:ssZ`) indicating the time the given Payment Profile was deleted at. Always `null` if the Payment Profile has not been deleted
	DeletedAt NullableTime `json:"deleted_at"`
	Status PaymentProfileStatus `json:"status"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentProfileGetResponse PaymentProfileGetResponse

// NewPaymentProfileGetResponse instantiates a new PaymentProfileGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentProfileGetResponse(updatedAt time.Time, createdAt time.Time, deletedAt NullableTime, status PaymentProfileStatus, requestId string) *PaymentProfileGetResponse {
	this := PaymentProfileGetResponse{}
	this.UpdatedAt = updatedAt
	this.CreatedAt = createdAt
	this.DeletedAt = deletedAt
	this.Status = status
	this.RequestId = requestId
	return &this
}

// NewPaymentProfileGetResponseWithDefaults instantiates a new PaymentProfileGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentProfileGetResponseWithDefaults() *PaymentProfileGetResponse {
	this := PaymentProfileGetResponse{}
	return &this
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PaymentProfileGetResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentProfileGetResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PaymentProfileGetResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentProfileGetResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentProfileGetResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentProfileGetResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PaymentProfileGetResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentProfileGetResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// SetDeletedAt sets field value
func (o *PaymentProfileGetResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// GetStatus returns the Status field value
func (o *PaymentProfileGetResponse) GetStatus() PaymentProfileStatus {
	if o == nil {
		var ret PaymentProfileStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentProfileGetResponse) GetStatusOk() (*PaymentProfileStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentProfileGetResponse) SetStatus(v PaymentProfileStatus) {
	o.Status = v
}

// GetRequestId returns the RequestId field value
func (o *PaymentProfileGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentProfileGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentProfileGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentProfileGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
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

func (o *PaymentProfileGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentProfileGetResponse := _PaymentProfileGetResponse{}

	if err = json.Unmarshal(bytes, &varPaymentProfileGetResponse); err == nil {
		*o = PaymentProfileGetResponse(varPaymentProfileGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "status")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentProfileGetResponse struct {
	value *PaymentProfileGetResponse
	isSet bool
}

func (v NullablePaymentProfileGetResponse) Get() *PaymentProfileGetResponse {
	return v.value
}

func (v *NullablePaymentProfileGetResponse) Set(val *PaymentProfileGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentProfileGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentProfileGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentProfileGetResponse(val *PaymentProfileGetResponse) *NullablePaymentProfileGetResponse {
	return &NullablePaymentProfileGetResponse{value: val, isSet: true}
}

func (v NullablePaymentProfileGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentProfileGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

