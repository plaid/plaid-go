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

// CraLoanApplication Contains loan application data.
type CraLoanApplication struct {
	// The user token for the user associated with the loan.
	UserToken string `json:"user_token"`
	// A unique identifier for the loan application. Personally identifiable information, such as an email address or phone number, should not be used in the `application_id`.
	ApplicationId string `json:"application_id"`
	Type CraLoanType `json:"type"`
	Decision CraLoanApplicationDecision `json:"decision"`
	// The date the user applied for the loan. The date should be in ISO 8601 format (YYYY-MM-DD).
	ApplicationDate *string `json:"application_date,omitempty"`
	// The date when the loan application's decision was made. The date should be in ISO 8601 format (YYYY-MM-DD).
	DecisionDate *string `json:"decision_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CraLoanApplication CraLoanApplication

// NewCraLoanApplication instantiates a new CraLoanApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCraLoanApplication(userToken string, applicationId string, type_ CraLoanType, decision CraLoanApplicationDecision) *CraLoanApplication {
	this := CraLoanApplication{}
	this.UserToken = userToken
	this.ApplicationId = applicationId
	this.Type = type_
	this.Decision = decision
	return &this
}

// NewCraLoanApplicationWithDefaults instantiates a new CraLoanApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCraLoanApplicationWithDefaults() *CraLoanApplication {
	this := CraLoanApplication{}
	return &this
}

// GetUserToken returns the UserToken field value
func (o *CraLoanApplication) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *CraLoanApplication) GetUserTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *CraLoanApplication) SetUserToken(v string) {
	o.UserToken = v
}

// GetApplicationId returns the ApplicationId field value
func (o *CraLoanApplication) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *CraLoanApplication) GetApplicationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *CraLoanApplication) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetType returns the Type field value
func (o *CraLoanApplication) GetType() CraLoanType {
	if o == nil {
		var ret CraLoanType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CraLoanApplication) GetTypeOk() (*CraLoanType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CraLoanApplication) SetType(v CraLoanType) {
	o.Type = v
}

// GetDecision returns the Decision field value
func (o *CraLoanApplication) GetDecision() CraLoanApplicationDecision {
	if o == nil {
		var ret CraLoanApplicationDecision
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *CraLoanApplication) GetDecisionOk() (*CraLoanApplicationDecision, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *CraLoanApplication) SetDecision(v CraLoanApplicationDecision) {
	o.Decision = v
}

// GetApplicationDate returns the ApplicationDate field value if set, zero value otherwise.
func (o *CraLoanApplication) GetApplicationDate() string {
	if o == nil || o.ApplicationDate == nil {
		var ret string
		return ret
	}
	return *o.ApplicationDate
}

// GetApplicationDateOk returns a tuple with the ApplicationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraLoanApplication) GetApplicationDateOk() (*string, bool) {
	if o == nil || o.ApplicationDate == nil {
		return nil, false
	}
	return o.ApplicationDate, true
}

// HasApplicationDate returns a boolean if a field has been set.
func (o *CraLoanApplication) HasApplicationDate() bool {
	if o != nil && o.ApplicationDate != nil {
		return true
	}

	return false
}

// SetApplicationDate gets a reference to the given string and assigns it to the ApplicationDate field.
func (o *CraLoanApplication) SetApplicationDate(v string) {
	o.ApplicationDate = &v
}

// GetDecisionDate returns the DecisionDate field value if set, zero value otherwise.
func (o *CraLoanApplication) GetDecisionDate() string {
	if o == nil || o.DecisionDate == nil {
		var ret string
		return ret
	}
	return *o.DecisionDate
}

// GetDecisionDateOk returns a tuple with the DecisionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CraLoanApplication) GetDecisionDateOk() (*string, bool) {
	if o == nil || o.DecisionDate == nil {
		return nil, false
	}
	return o.DecisionDate, true
}

// HasDecisionDate returns a boolean if a field has been set.
func (o *CraLoanApplication) HasDecisionDate() bool {
	if o != nil && o.DecisionDate != nil {
		return true
	}

	return false
}

// SetDecisionDate gets a reference to the given string and assigns it to the DecisionDate field.
func (o *CraLoanApplication) SetDecisionDate(v string) {
	o.DecisionDate = &v
}

func (o CraLoanApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_token"] = o.UserToken
	}
	if true {
		toSerialize["application_id"] = o.ApplicationId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["decision"] = o.Decision
	}
	if o.ApplicationDate != nil {
		toSerialize["application_date"] = o.ApplicationDate
	}
	if o.DecisionDate != nil {
		toSerialize["decision_date"] = o.DecisionDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CraLoanApplication) UnmarshalJSON(bytes []byte) (err error) {
	varCraLoanApplication := _CraLoanApplication{}

	if err = json.Unmarshal(bytes, &varCraLoanApplication); err == nil {
		*o = CraLoanApplication(varCraLoanApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "user_token")
		delete(additionalProperties, "application_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "application_date")
		delete(additionalProperties, "decision_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCraLoanApplication struct {
	value *CraLoanApplication
	isSet bool
}

func (v NullableCraLoanApplication) Get() *CraLoanApplication {
	return v.value
}

func (v *NullableCraLoanApplication) Set(val *CraLoanApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableCraLoanApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableCraLoanApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCraLoanApplication(val *CraLoanApplication) *NullableCraLoanApplication {
	return &NullableCraLoanApplication{value: val, isSet: true}
}

func (v NullableCraLoanApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCraLoanApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


