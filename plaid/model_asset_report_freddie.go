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

// AssetReportFreddie An object representing an Asset Report with Freddie Mac schema.
type AssetReportFreddie struct {
	LOANS Loans `json:"LOANS"`
	PARTIES Parties `json:"PARTIES"`
	SERVICES Services `json:"SERVICES"`
	AdditionalProperties map[string]interface{}
}

type _AssetReportFreddie AssetReportFreddie

// NewAssetReportFreddie instantiates a new AssetReportFreddie object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportFreddie(lOANS Loans, pARTIES Parties, sERVICES Services) *AssetReportFreddie {
	this := AssetReportFreddie{}
	this.LOANS = lOANS
	this.PARTIES = pARTIES
	this.SERVICES = sERVICES
	return &this
}

// NewAssetReportFreddieWithDefaults instantiates a new AssetReportFreddie object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportFreddieWithDefaults() *AssetReportFreddie {
	this := AssetReportFreddie{}
	return &this
}

// GetLOANS returns the LOANS field value
func (o *AssetReportFreddie) GetLOANS() Loans {
	if o == nil {
		var ret Loans
		return ret
	}

	return o.LOANS
}

// GetLOANSOk returns a tuple with the LOANS field value
// and a boolean to check if the value has been set.
func (o *AssetReportFreddie) GetLOANSOk() (*Loans, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LOANS, true
}

// SetLOANS sets field value
func (o *AssetReportFreddie) SetLOANS(v Loans) {
	o.LOANS = v
}

// GetPARTIES returns the PARTIES field value
func (o *AssetReportFreddie) GetPARTIES() Parties {
	if o == nil {
		var ret Parties
		return ret
	}

	return o.PARTIES
}

// GetPARTIESOk returns a tuple with the PARTIES field value
// and a boolean to check if the value has been set.
func (o *AssetReportFreddie) GetPARTIESOk() (*Parties, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PARTIES, true
}

// SetPARTIES sets field value
func (o *AssetReportFreddie) SetPARTIES(v Parties) {
	o.PARTIES = v
}

// GetSERVICES returns the SERVICES field value
func (o *AssetReportFreddie) GetSERVICES() Services {
	if o == nil {
		var ret Services
		return ret
	}

	return o.SERVICES
}

// GetSERVICESOk returns a tuple with the SERVICES field value
// and a boolean to check if the value has been set.
func (o *AssetReportFreddie) GetSERVICESOk() (*Services, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SERVICES, true
}

// SetSERVICES sets field value
func (o *AssetReportFreddie) SetSERVICES(v Services) {
	o.SERVICES = v
}

func (o AssetReportFreddie) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["LOANS"] = o.LOANS
	}
	if true {
		toSerialize["PARTIES"] = o.PARTIES
	}
	if true {
		toSerialize["SERVICES"] = o.SERVICES
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetReportFreddie) UnmarshalJSON(bytes []byte) (err error) {
	varAssetReportFreddie := _AssetReportFreddie{}

	if err = json.Unmarshal(bytes, &varAssetReportFreddie); err == nil {
		*o = AssetReportFreddie(varAssetReportFreddie)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "LOANS")
		delete(additionalProperties, "PARTIES")
		delete(additionalProperties, "SERVICES")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetReportFreddie struct {
	value *AssetReportFreddie
	isSet bool
}

func (v NullableAssetReportFreddie) Get() *AssetReportFreddie {
	return v.value
}

func (v *NullableAssetReportFreddie) Set(val *AssetReportFreddie) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportFreddie) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportFreddie) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportFreddie(val *AssetReportFreddie) *NullableAssetReportFreddie {
	return &NullableAssetReportFreddie{value: val, isSet: true}
}

func (v NullableAssetReportFreddie) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportFreddie) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


