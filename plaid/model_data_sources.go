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
	"fmt"
)

// DataSources A description of the source of data for a given product/data type.  `INSTITUTION`: The institution supports this product, and the data was provided by the institution. `INSTITUTION_MASK`: The user manually provided the full account number, which was matched to the account mask provided by the institution. Only applicable to the `numbers` data type. `USER`: The institution does not support this product, and the data was manually provided by the user.
type DataSources string

var _ = fmt.Printf

// List of DataSources
const (
	DATASOURCES_INSTITUTION DataSources = "INSTITUTION"
	DATASOURCES_INSTITUTION_MASK DataSources = "INSTITUTION_MASK"
	DATASOURCES_USER DataSources = "USER"
)

var allowedDataSourcesEnumValues = []DataSources{
	"INSTITUTION",
	"INSTITUTION_MASK",
	"USER",
}

func (v *DataSources) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := DataSources(value)


	*v = enumTypeValue
	return nil
}

// NewDataSourcesFromValue returns a pointer to a valid DataSources
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataSourcesFromValue(v string) (*DataSources, error) {
	ev := DataSources(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataSources) IsValid() bool {
	for _, existing := range allowedDataSourcesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataSources value
func (v DataSources) Ptr() *DataSources {
	return &v
}

type NullableDataSources struct {
	value *DataSources
	isSet bool
}

func (v NullableDataSources) Get() *DataSources {
	return v.value
}

func (v *NullableDataSources) Set(val *DataSources) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSources) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSources(val *DataSources) *NullableDataSources {
	return &NullableDataSources{value: val, isSet: true}
}

func (v NullableDataSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
