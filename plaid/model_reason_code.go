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
	"fmt"
)

// ReasonCode Specifies the reason for cancelling transfer. This is required for RfP transfers, and will be ignored for other networks.  `\"AC03\"` - Invalid Creditor Account Number  `\"AM09\"` - Incorrect Amount  `\"CUST\"` - Requested By Customer - Cancellation requested  `\"DUPL\"` - Duplicate Payment  `\"FRAD\"` - Fraudulent Payment - Unauthorized or fraudulently induced  `\"TECH\"` - Technical Problem - Cancellation due to system issues  `\"UPAY\"` - Undue Payment - Payment was made through another channel  `\"AC14\"` - Invalid or Missing Creditor Account Type  `\"AM06\"` - Amount Too Low  `\"BE05\"` - Unrecognized Initiating Party  `\"FOCR\"` - Following Refund Request  `\"MS02\"` - No Specified Reason - Customer  `\"MS03\"` - No Specified Reason - Agent  `\"RR04\"` - Regulatory Reason  `\"RUTA\"` - Return Upon Unable To Apply
type ReasonCode string

var _ = fmt.Printf

// List of ReasonCode
const (
	REASONCODE_AC03 ReasonCode = "AC03"
	REASONCODE_AM09 ReasonCode = "AM09"
	REASONCODE_CUST ReasonCode = "CUST"
	REASONCODE_DUPL ReasonCode = "DUPL"
	REASONCODE_FRAD ReasonCode = "FRAD"
	REASONCODE_TECH ReasonCode = "TECH"
	REASONCODE_UPAY ReasonCode = "UPAY"
	REASONCODE_AC14 ReasonCode = "AC14"
	REASONCODE_AM06 ReasonCode = "AM06"
	REASONCODE_BE05 ReasonCode = "BE05"
	REASONCODE_FOCR ReasonCode = "FOCR"
	REASONCODE_MS02 ReasonCode = "MS02"
	REASONCODE_MS03 ReasonCode = "MS03"
	REASONCODE_RR04 ReasonCode = "RR04"
	REASONCODE_RUTA ReasonCode = "RUTA"
)

var allowedReasonCodeEnumValues = []ReasonCode{
	"AC03",
	"AM09",
	"CUST",
	"DUPL",
	"FRAD",
	"TECH",
	"UPAY",
	"AC14",
	"AM06",
	"BE05",
	"FOCR",
	"MS02",
	"MS03",
	"RR04",
	"RUTA",
}

func (v *ReasonCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := ReasonCode(value)


	*v = enumTypeValue
	return nil
}

// NewReasonCodeFromValue returns a pointer to a valid ReasonCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReasonCodeFromValue(v string) (*ReasonCode, error) {
	ev := ReasonCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReasonCode) IsValid() bool {
	for _, existing := range allowedReasonCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReasonCode value
func (v ReasonCode) Ptr() *ReasonCode {
	return &v
}

type NullableReasonCode struct {
	value *ReasonCode
	isSet bool
}

func (v NullableReasonCode) Get() *ReasonCode {
	return v.value
}

func (v *NullableReasonCode) Set(val *ReasonCode) {
	v.value = val
	v.isSet = true
}

func (v NullableReasonCode) IsSet() bool {
	return v.isSet
}

func (v *NullableReasonCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReasonCode(val *ReasonCode) *NullableReasonCode {
	return &NullableReasonCode{value: val, isSet: true}
}

func (v NullableReasonCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReasonCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

