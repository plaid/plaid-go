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

// AssetReportRefreshResponse AssetReportRefreshResponse defines the response schema for `/asset_report/refresh`
type AssetReportRefreshResponse struct {
	// A unique ID identifying an Asset Report. Like all Plaid identifiers, this ID is case sensitive.
	AssetReportId string `json:"asset_report_id"`
	// A token that can be provided to endpoints such as `/asset_report/get` or `/asset_report/pdf/get` to fetch or update an Asset Report.
	AssetReportToken string `json:"asset_report_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportRefreshResponse AssetReportRefreshResponse

// NewAssetReportRefreshResponse instantiates a new AssetReportRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportRefreshResponse(assetReportId string, assetReportToken string, requestId string) *AssetReportRefreshResponse {
	this := AssetReportRefreshResponse{}
	this.AssetReportId = assetReportId
	this.AssetReportToken = assetReportToken
	this.RequestId = requestId
	return &this
}

// NewAssetReportRefreshResponseWithDefaults instantiates a new AssetReportRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportRefreshResponseWithDefaults() *AssetReportRefreshResponse {
	this := AssetReportRefreshResponse{}
	return &this
}

// GetAssetReportId returns the AssetReportId field value
func (o *AssetReportRefreshResponse) GetAssetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportId
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshResponse) GetAssetReportIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportId, true
}

// SetAssetReportId sets field value
func (o *AssetReportRefreshResponse) SetAssetReportId(v string) {
	o.AssetReportId = v
}

// GetAssetReportToken returns the AssetReportToken field value
func (o *AssetReportRefreshResponse) GetAssetReportToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetReportToken
}

// GetAssetReportTokenOk returns a tuple with the AssetReportToken field value
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshResponse) GetAssetReportTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetReportToken, true
}

// SetAssetReportToken sets field value
func (o *AssetReportRefreshResponse) SetAssetReportToken(v string) {
	o.AssetReportToken = v
}

// GetRequestId returns the RequestId field value
func (o *AssetReportRefreshResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AssetReportRefreshResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AssetReportRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asset_report_id"] = o.AssetReportId
	}
	if true {
		toSerialize["asset_report_token"] = o.AssetReportToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportRefreshResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportRefreshResponse := _AssetReportRefreshResponse{}

	if err = json.Unmarshal(bytes, &varAssetReportRefreshResponse); err == nil {
		*o = AssetReportRefreshResponse(varAssetReportRefreshResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "asset_report_id")
		delete(additionalProperties, "asset_report_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportRefreshResponse struct {
	value *AssetReportRefreshResponse
	isSet bool
}

func (v NullableAssetReportRefreshResponse) Get() *AssetReportRefreshResponse {
	return v.value
}

func (v *NullableAssetReportRefreshResponse) Set(val *AssetReportRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportRefreshResponse(val *AssetReportRefreshResponse) *NullableAssetReportRefreshResponse {
	return &NullableAssetReportRefreshResponse{value: val, isSet: true}
}

func (v NullableAssetReportRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


