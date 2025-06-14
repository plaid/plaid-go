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

// PartnerEndCustomerCustomerSupportInfo This information is public. Users of your app will see this information when managing connections between your app and their bank accounts in Plaid Portal. Defaults to partner's customer support info if omitted. This field is mandatory for partners whose Plaid accounts were created after November 26, 2024 and will be mandatory for all partners by the 1033 compliance deadline.
type PartnerEndCustomerCustomerSupportInfo struct {
	// This field is mandatory for partners whose Plaid accounts were created after November 26, 2024 and will be mandatory for all partners by the 1033 compliance deadline.
	Email *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	ContactUrl *string `json:"contact_url,omitempty"`
	LinkUpdateUrl *string `json:"link_update_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerEndCustomerCustomerSupportInfo PartnerEndCustomerCustomerSupportInfo

// NewPartnerEndCustomerCustomerSupportInfo instantiates a new PartnerEndCustomerCustomerSupportInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomerCustomerSupportInfo() *PartnerEndCustomerCustomerSupportInfo {
	this := PartnerEndCustomerCustomerSupportInfo{}
	return &this
}

// NewPartnerEndCustomerCustomerSupportInfoWithDefaults instantiates a new PartnerEndCustomerCustomerSupportInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerCustomerSupportInfoWithDefaults() *PartnerEndCustomerCustomerSupportInfo {
	this := PartnerEndCustomerCustomerSupportInfo{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PartnerEndCustomerCustomerSupportInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerCustomerSupportInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PartnerEndCustomerCustomerSupportInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PartnerEndCustomerCustomerSupportInfo) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *PartnerEndCustomerCustomerSupportInfo) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerCustomerSupportInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PartnerEndCustomerCustomerSupportInfo) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *PartnerEndCustomerCustomerSupportInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetContactUrl returns the ContactUrl field value if set, zero value otherwise.
func (o *PartnerEndCustomerCustomerSupportInfo) GetContactUrl() string {
	if o == nil || o.ContactUrl == nil {
		var ret string
		return ret
	}
	return *o.ContactUrl
}

// GetContactUrlOk returns a tuple with the ContactUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerCustomerSupportInfo) GetContactUrlOk() (*string, bool) {
	if o == nil || o.ContactUrl == nil {
		return nil, false
	}
	return o.ContactUrl, true
}

// HasContactUrl returns a boolean if a field has been set.
func (o *PartnerEndCustomerCustomerSupportInfo) HasContactUrl() bool {
	if o != nil && o.ContactUrl != nil {
		return true
	}

	return false
}

// SetContactUrl gets a reference to the given string and assigns it to the ContactUrl field.
func (o *PartnerEndCustomerCustomerSupportInfo) SetContactUrl(v string) {
	o.ContactUrl = &v
}

// GetLinkUpdateUrl returns the LinkUpdateUrl field value if set, zero value otherwise.
func (o *PartnerEndCustomerCustomerSupportInfo) GetLinkUpdateUrl() string {
	if o == nil || o.LinkUpdateUrl == nil {
		var ret string
		return ret
	}
	return *o.LinkUpdateUrl
}

// GetLinkUpdateUrlOk returns a tuple with the LinkUpdateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerCustomerSupportInfo) GetLinkUpdateUrlOk() (*string, bool) {
	if o == nil || o.LinkUpdateUrl == nil {
		return nil, false
	}
	return o.LinkUpdateUrl, true
}

// HasLinkUpdateUrl returns a boolean if a field has been set.
func (o *PartnerEndCustomerCustomerSupportInfo) HasLinkUpdateUrl() bool {
	if o != nil && o.LinkUpdateUrl != nil {
		return true
	}

	return false
}

// SetLinkUpdateUrl gets a reference to the given string and assigns it to the LinkUpdateUrl field.
func (o *PartnerEndCustomerCustomerSupportInfo) SetLinkUpdateUrl(v string) {
	o.LinkUpdateUrl = &v
}

func (o PartnerEndCustomerCustomerSupportInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.ContactUrl != nil {
		toSerialize["contact_url"] = o.ContactUrl
	}
	if o.LinkUpdateUrl != nil {
		toSerialize["link_update_url"] = o.LinkUpdateUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerEndCustomerCustomerSupportInfo) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerEndCustomerCustomerSupportInfo := _PartnerEndCustomerCustomerSupportInfo{}

	if err = json.Unmarshal(bytes, &varPartnerEndCustomerCustomerSupportInfo); err == nil {
		*o = PartnerEndCustomerCustomerSupportInfo(varPartnerEndCustomerCustomerSupportInfo)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "contact_url")
		delete(additionalProperties, "link_update_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerEndCustomerCustomerSupportInfo struct {
	value *PartnerEndCustomerCustomerSupportInfo
	isSet bool
}

func (v NullablePartnerEndCustomerCustomerSupportInfo) Get() *PartnerEndCustomerCustomerSupportInfo {
	return v.value
}

func (v *NullablePartnerEndCustomerCustomerSupportInfo) Set(val *PartnerEndCustomerCustomerSupportInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerCustomerSupportInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerCustomerSupportInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerCustomerSupportInfo(val *PartnerEndCustomerCustomerSupportInfo) *NullablePartnerEndCustomerCustomerSupportInfo {
	return &NullablePartnerEndCustomerCustomerSupportInfo{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerCustomerSupportInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerCustomerSupportInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


