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
)

// IssuesSearchResponse IssuesSearchResponse defines the response schema for `/issues/search`.
type IssuesSearchResponse struct {
	// A list of issues affecting the Item, session, or request passed in, conforming to the Issues data model. An empty list indicates that no matching issues were found.
	Issues *[]Issue `json:"issues,omitempty"`
	// A unique identifier for the API request.
	RequestId *string `json:"request_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IssuesSearchResponse IssuesSearchResponse

// NewIssuesSearchResponse instantiates a new IssuesSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuesSearchResponse() *IssuesSearchResponse {
	this := IssuesSearchResponse{}
	return &this
}

// NewIssuesSearchResponseWithDefaults instantiates a new IssuesSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuesSearchResponseWithDefaults() *IssuesSearchResponse {
	this := IssuesSearchResponse{}
	return &this
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *IssuesSearchResponse) GetIssues() []Issue {
	if o == nil || o.Issues == nil {
		var ret []Issue
		return ret
	}
	return *o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesSearchResponse) GetIssuesOk() (*[]Issue, bool) {
	if o == nil || o.Issues == nil {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *IssuesSearchResponse) HasIssues() bool {
	if o != nil && o.Issues != nil {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []Issue and assigns it to the Issues field.
func (o *IssuesSearchResponse) SetIssues(v []Issue) {
	o.Issues = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *IssuesSearchResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesSearchResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *IssuesSearchResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *IssuesSearchResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o IssuesSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Issues != nil {
		toSerialize["issues"] = o.Issues
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IssuesSearchResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIssuesSearchResponse := _IssuesSearchResponse{}

	if err = json.Unmarshal(bytes, &varIssuesSearchResponse); err == nil {
		*o = IssuesSearchResponse(varIssuesSearchResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "issues")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIssuesSearchResponse struct {
	value *IssuesSearchResponse
	isSet bool
}

func (v NullableIssuesSearchResponse) Get() *IssuesSearchResponse {
	return v.value
}

func (v *NullableIssuesSearchResponse) Set(val *IssuesSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuesSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuesSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuesSearchResponse(val *IssuesSearchResponse) *NullableIssuesSearchResponse {
	return &NullableIssuesSearchResponse{value: val, isSet: true}
}

func (v NullableIssuesSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuesSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

