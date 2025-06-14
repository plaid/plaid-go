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

// LinkDeliveryWebhookCommunicationMethod The communication method used to deliver the Hosted Link session
type LinkDeliveryWebhookCommunicationMethod string

var _ = fmt.Printf

// List of LinkDeliveryWebhookCommunicationMethod
const (
	LINKDELIVERYWEBHOOKCOMMUNICATIONMETHOD_SMS LinkDeliveryWebhookCommunicationMethod = "SMS"
	LINKDELIVERYWEBHOOKCOMMUNICATIONMETHOD_EMAIL LinkDeliveryWebhookCommunicationMethod = "EMAIL"
)

var allowedLinkDeliveryWebhookCommunicationMethodEnumValues = []LinkDeliveryWebhookCommunicationMethod{
	"SMS",
	"EMAIL",
}

func (v *LinkDeliveryWebhookCommunicationMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := LinkDeliveryWebhookCommunicationMethod(value)


	*v = enumTypeValue
	return nil
}

// NewLinkDeliveryWebhookCommunicationMethodFromValue returns a pointer to a valid LinkDeliveryWebhookCommunicationMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkDeliveryWebhookCommunicationMethodFromValue(v string) (*LinkDeliveryWebhookCommunicationMethod, error) {
	ev := LinkDeliveryWebhookCommunicationMethod(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkDeliveryWebhookCommunicationMethod) IsValid() bool {
	for _, existing := range allowedLinkDeliveryWebhookCommunicationMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinkDeliveryWebhookCommunicationMethod value
func (v LinkDeliveryWebhookCommunicationMethod) Ptr() *LinkDeliveryWebhookCommunicationMethod {
	return &v
}

type NullableLinkDeliveryWebhookCommunicationMethod struct {
	value *LinkDeliveryWebhookCommunicationMethod
	isSet bool
}

func (v NullableLinkDeliveryWebhookCommunicationMethod) Get() *LinkDeliveryWebhookCommunicationMethod {
	return v.value
}

func (v *NullableLinkDeliveryWebhookCommunicationMethod) Set(val *LinkDeliveryWebhookCommunicationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryWebhookCommunicationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryWebhookCommunicationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryWebhookCommunicationMethod(val *LinkDeliveryWebhookCommunicationMethod) *NullableLinkDeliveryWebhookCommunicationMethod {
	return &NullableLinkDeliveryWebhookCommunicationMethod{value: val, isSet: true}
}

func (v NullableLinkDeliveryWebhookCommunicationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryWebhookCommunicationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

