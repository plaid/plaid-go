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

// PaymentInitiationConsent PaymentInitiationConsent defines a payment initiation consent.
type PaymentInitiationConsent struct {
	// The consent ID.
	ConsentId string `json:"consent_id"`
	Status PaymentInitiationConsentStatus `json:"status"`
	// Consent creation timestamp, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the recipient the payment consent is for.
	RecipientId string `json:"recipient_id"`
	// A reference for the payment consent.
	Reference string `json:"reference"`
	Constraints PaymentInitiationConsentConstraints `json:"constraints"`
	// Deprecated, use the 'type' field instead.
	Scopes *[]PaymentInitiationConsentScope `json:"scopes,omitempty"`
	Type *PaymentInitiationConsentType `json:"type,omitempty"`
	PayerDetails NullableExternalPaymentRefundDetails `json:"payer_details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationConsent PaymentInitiationConsent

// NewPaymentInitiationConsent instantiates a new PaymentInitiationConsent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationConsent(consentId string, status PaymentInitiationConsentStatus, createdAt time.Time, recipientId string, reference string, constraints PaymentInitiationConsentConstraints) *PaymentInitiationConsent {
	this := PaymentInitiationConsent{}
	this.ConsentId = consentId
	this.Status = status
	this.CreatedAt = createdAt
	this.RecipientId = recipientId
	this.Reference = reference
	this.Constraints = constraints
	return &this
}

// NewPaymentInitiationConsentWithDefaults instantiates a new PaymentInitiationConsent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationConsentWithDefaults() *PaymentInitiationConsent {
	this := PaymentInitiationConsent{}
	return &this
}

// GetConsentId returns the ConsentId field value
func (o *PaymentInitiationConsent) GetConsentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsentId
}

// GetConsentIdOk returns a tuple with the ConsentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsent) GetConsentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsentId, true
}

// SetConsentId sets field value
func (o *PaymentInitiationConsent) SetConsentId(v string) {
	o.ConsentId = v
}

// GetStatus returns the Status field value
func (o *PaymentInitiationConsent) GetStatus() PaymentInitiationConsentStatus {
	if o == nil {
		var ret PaymentInitiationConsentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsent) GetStatusOk() (*PaymentInitiationConsentStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentInitiationConsent) SetStatus(v PaymentInitiationConsentStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentInitiationConsent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentInitiationConsent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetRecipientId returns the RecipientId field value
func (o *PaymentInitiationConsent) GetRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsent) GetRecipientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientId, true
}

// SetRecipientId sets field value
func (o *PaymentInitiationConsent) SetRecipientId(v string) {
	o.RecipientId = v
}

// GetReference returns the Reference field value
func (o *PaymentInitiationConsent) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsent) GetReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *PaymentInitiationConsent) SetReference(v string) {
	o.Reference = v
}

// GetConstraints returns the Constraints field value
func (o *PaymentInitiationConsent) GetConstraints() PaymentInitiationConsentConstraints {
	if o == nil {
		var ret PaymentInitiationConsentConstraints
		return ret
	}

	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsent) GetConstraintsOk() (*PaymentInitiationConsentConstraints, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Constraints, true
}

// SetConstraints sets field value
func (o *PaymentInitiationConsent) SetConstraints(v PaymentInitiationConsentConstraints) {
	o.Constraints = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *PaymentInitiationConsent) GetScopes() []PaymentInitiationConsentScope {
	if o == nil || o.Scopes == nil {
		var ret []PaymentInitiationConsentScope
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsent) GetScopesOk() (*[]PaymentInitiationConsentScope, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *PaymentInitiationConsent) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []PaymentInitiationConsentScope and assigns it to the Scopes field.
func (o *PaymentInitiationConsent) SetScopes(v []PaymentInitiationConsentScope) {
	o.Scopes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentInitiationConsent) GetType() PaymentInitiationConsentType {
	if o == nil || o.Type == nil {
		var ret PaymentInitiationConsentType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsent) GetTypeOk() (*PaymentInitiationConsentType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentInitiationConsent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given PaymentInitiationConsentType and assigns it to the Type field.
func (o *PaymentInitiationConsent) SetType(v PaymentInitiationConsentType) {
	o.Type = &v
}

// GetPayerDetails returns the PayerDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationConsent) GetPayerDetails() ExternalPaymentRefundDetails {
	if o == nil || o.PayerDetails.Get() == nil {
		var ret ExternalPaymentRefundDetails
		return ret
	}
	return *o.PayerDetails.Get()
}

// GetPayerDetailsOk returns a tuple with the PayerDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationConsent) GetPayerDetailsOk() (*ExternalPaymentRefundDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayerDetails.Get(), o.PayerDetails.IsSet()
}

// HasPayerDetails returns a boolean if a field has been set.
func (o *PaymentInitiationConsent) HasPayerDetails() bool {
	if o != nil && o.PayerDetails.IsSet() {
		return true
	}

	return false
}

// SetPayerDetails gets a reference to the given NullableExternalPaymentRefundDetails and assigns it to the PayerDetails field.
func (o *PaymentInitiationConsent) SetPayerDetails(v ExternalPaymentRefundDetails) {
	o.PayerDetails.Set(&v)
}
// SetPayerDetailsNil sets the value for PayerDetails to be an explicit nil
func (o *PaymentInitiationConsent) SetPayerDetailsNil() {
	o.PayerDetails.Set(nil)
}

// UnsetPayerDetails ensures that no value is present for PayerDetails, not even an explicit nil
func (o *PaymentInitiationConsent) UnsetPayerDetails() {
	o.PayerDetails.Unset()
}

func (o PaymentInitiationConsent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["consent_id"] = o.ConsentId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["recipient_id"] = o.RecipientId
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["constraints"] = o.Constraints
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.PayerDetails.IsSet() {
		toSerialize["payer_details"] = o.PayerDetails.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationConsent) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationConsent := _PaymentInitiationConsent{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationConsent); err == nil {
		*o = PaymentInitiationConsent(varPaymentInitiationConsent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "consent_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "recipient_id")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "type")
		delete(additionalProperties, "payer_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationConsent struct {
	value *PaymentInitiationConsent
	isSet bool
}

func (v NullablePaymentInitiationConsent) Get() *PaymentInitiationConsent {
	return v.value
}

func (v *NullablePaymentInitiationConsent) Set(val *PaymentInitiationConsent) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsent) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsent(val *PaymentInitiationConsent) *NullablePaymentInitiationConsent {
	return &NullablePaymentInitiationConsent{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


