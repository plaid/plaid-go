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

// UserCustomPassword Custom test accounts are configured with a JSON configuration object formulated according to the schema below. All top level fields are optional. Sending an empty object as a configuration will result in an account configured with random balances and transaction history.
type UserCustomPassword struct {
	// The version of the password schema to use, possible values are 1 or 2. The default value is 2. You should only specify 1 if you know it is necessary for your test suite.
	Version NullableString `json:"version,omitempty"`
	// A seed, in the form of a string, that will be used to randomly generate account and transaction data, if this data is not specified using the `override_accounts` argument. If no seed is specified, the randomly generated data will be different each time.  Note that transactions data is generated relative to the Item's creation date. Different Items created on different dates with the same seed for transactions data will have different dates for the transactions. The number of days between each transaction and the Item creation will remain constant. For example, an Item created on December 15 might show a transaction on December 14. An Item created on December 20, using the same seed, would show that same transaction occurring on December 19.
	Seed string `json:"seed"`
	// An array of account overrides to configure the accounts for the Item. By default, if no override is specified, transactions and account data will be randomly generated based on the account type and subtype, and other products will have fixed or empty data.
	OverrideAccounts []OverrideAccounts `json:"override_accounts"`
	Mfa MFA `json:"mfa"`
	// You may trigger a reCAPTCHA in Plaid Link in the Sandbox environment by using the recaptcha field. Possible values are `good` or `bad`. A value of `good` will result in successful Item creation and `bad` will result in a `RECAPTCHA_BAD` error to simulate a failed reCAPTCHA. Both values require the reCAPTCHA to be manually solved within Plaid Link.
	Recaptcha string `json:"recaptcha"`
	// An error code to force on Item creation. Possible values are:  `\"INSTITUTION_NOT_RESPONDING\"` `\"INSTITUTION_NO_LONGER_SUPPORTED\"` `\"INVALID_CREDENTIALS\"` `\"INVALID_MFA\"` `\"ITEM_LOCKED\"` `\"ITEM_LOGIN_REQUIRED\"` `\"ITEM_NOT_SUPPORTED\"` `\"INVALID_LINK_TOKEN\"` `\"MFA_NOT_SUPPORTED\"` `\"NO_ACCOUNTS\"` `\"PLAID_ERROR\"` `\"USER_INPUT_TIMEOUT\"` `\"USER_SETUP_REQUIRED\"`
	ForceError string `json:"force_error"`
	AdditionalProperties map[string]interface{}
}

type _UserCustomPassword UserCustomPassword

// NewUserCustomPassword instantiates a new UserCustomPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCustomPassword(seed string, overrideAccounts []OverrideAccounts, mfa MFA, recaptcha string, forceError string) *UserCustomPassword {
	this := UserCustomPassword{}
	this.Seed = seed
	this.OverrideAccounts = overrideAccounts
	this.Mfa = mfa
	this.Recaptcha = recaptcha
	this.ForceError = forceError
	return &this
}

// NewUserCustomPasswordWithDefaults instantiates a new UserCustomPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCustomPasswordWithDefaults() *UserCustomPassword {
	this := UserCustomPassword{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserCustomPassword) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserCustomPassword) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *UserCustomPassword) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *UserCustomPassword) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *UserCustomPassword) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *UserCustomPassword) UnsetVersion() {
	o.Version.Unset()
}

// GetSeed returns the Seed field value
func (o *UserCustomPassword) GetSeed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Seed
}

// GetSeedOk returns a tuple with the Seed field value
// and a boolean to check if the value has been set.
func (o *UserCustomPassword) GetSeedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Seed, true
}

// SetSeed sets field value
func (o *UserCustomPassword) SetSeed(v string) {
	o.Seed = v
}

// GetOverrideAccounts returns the OverrideAccounts field value
func (o *UserCustomPassword) GetOverrideAccounts() []OverrideAccounts {
	if o == nil {
		var ret []OverrideAccounts
		return ret
	}

	return o.OverrideAccounts
}

// GetOverrideAccountsOk returns a tuple with the OverrideAccounts field value
// and a boolean to check if the value has been set.
func (o *UserCustomPassword) GetOverrideAccountsOk() (*[]OverrideAccounts, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OverrideAccounts, true
}

// SetOverrideAccounts sets field value
func (o *UserCustomPassword) SetOverrideAccounts(v []OverrideAccounts) {
	o.OverrideAccounts = v
}

// GetMfa returns the Mfa field value
func (o *UserCustomPassword) GetMfa() MFA {
	if o == nil {
		var ret MFA
		return ret
	}

	return o.Mfa
}

// GetMfaOk returns a tuple with the Mfa field value
// and a boolean to check if the value has been set.
func (o *UserCustomPassword) GetMfaOk() (*MFA, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mfa, true
}

// SetMfa sets field value
func (o *UserCustomPassword) SetMfa(v MFA) {
	o.Mfa = v
}

// GetRecaptcha returns the Recaptcha field value
func (o *UserCustomPassword) GetRecaptcha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recaptcha
}

// GetRecaptchaOk returns a tuple with the Recaptcha field value
// and a boolean to check if the value has been set.
func (o *UserCustomPassword) GetRecaptchaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recaptcha, true
}

// SetRecaptcha sets field value
func (o *UserCustomPassword) SetRecaptcha(v string) {
	o.Recaptcha = v
}

// GetForceError returns the ForceError field value
func (o *UserCustomPassword) GetForceError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForceError
}

// GetForceErrorOk returns a tuple with the ForceError field value
// and a boolean to check if the value has been set.
func (o *UserCustomPassword) GetForceErrorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ForceError, true
}

// SetForceError sets field value
func (o *UserCustomPassword) SetForceError(v string) {
	o.ForceError = v
}

func (o UserCustomPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if true {
		toSerialize["seed"] = o.Seed
	}
	if true {
		toSerialize["override_accounts"] = o.OverrideAccounts
	}
	if true {
		toSerialize["mfa"] = o.Mfa
	}
	if true {
		toSerialize["recaptcha"] = o.Recaptcha
	}
	if true {
		toSerialize["force_error"] = o.ForceError
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserCustomPassword) UnmarshalJSON(bytes []byte) (err error) {
	varUserCustomPassword := _UserCustomPassword{}

	if err = json.Unmarshal(bytes, &varUserCustomPassword); err == nil {
		*o = UserCustomPassword(varUserCustomPassword)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "seed")
		delete(additionalProperties, "override_accounts")
		delete(additionalProperties, "mfa")
		delete(additionalProperties, "recaptcha")
		delete(additionalProperties, "force_error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserCustomPassword struct {
	value *UserCustomPassword
	isSet bool
}

func (v NullableUserCustomPassword) Get() *UserCustomPassword {
	return v.value
}

func (v *NullableUserCustomPassword) Set(val *UserCustomPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCustomPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCustomPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCustomPassword(val *UserCustomPassword) *NullableUserCustomPassword {
	return &NullableUserCustomPassword{value: val, isSet: true}
}

func (v NullableUserCustomPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCustomPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


