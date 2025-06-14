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

// AccountBaseNullable A single account at a financial institution.
type AccountBaseNullable struct {
	// Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  The `account_id` can also change if the `access_token` is deleted and the same credentials that were used to generate that `access_token` are used to generate a new `access_token` on a later date. In that case, the new `account_id` will be different from the old `account_id`.  If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  Like all Plaid identifiers, the `account_id` is case sensitive.
	AccountId string `json:"account_id"`
	Balances AccountBalance `json:"balances"`
	// The last 2-4 alphanumeric characters of either the account’s displayed mask or the account’s official account number. Note that the mask may be non-unique between an Item’s accounts.
	Mask NullableString `json:"mask"`
	// The name of the account, either assigned by the user or by the financial institution itself
	Name string `json:"name"`
	// The official name of the account as given by the financial institution
	OfficialName NullableString `json:"official_name"`
	Type AccountType `json:"type"`
	Subtype NullableAccountSubtype `json:"subtype"`
	// The current verification status of an Auth Item initiated through micro-deposits or database verification. Returned for Auth Items only.  `pending_automatic_verification`: The Item is pending automatic verification  `pending_manual_verification`: The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the micro-deposit.  `automatically_verified`: The Item has successfully been automatically verified   `manually_verified`: The Item has successfully been manually verified  `verification_expired`: Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.  `verification_failed`: The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.    `unsent`: The Item is pending micro-deposit verification, but Plaid has not yet sent the micro-deposit.  `database_matched`: The Item has successfully been verified using Plaid's data sources. Only returned for Auth Items created via Database Match.  `database_insights_pass`: The Item's numbers have been verified using Plaid's data sources: the routing and account number match a routing and account number of an account recognized on the Plaid network, and the account is not known by Plaid to be frozen or closed. Only returned for Auth Items created via Database Auth.  `database_insights_pass_with_caution`:The Item's numbers have been verified using Plaid's data sources and have some signal for being valid: the routing and account number were not recognized on the Plaid network, but the routing number is valid and the account number is a potential valid account number for that routing number. Only returned for Auth Items created via Database Auth.  `database_insights_fail`: The Item's numbers have been verified using Plaid's data sources and have signal for being invalid and/or have no signal for being valid. Typically this indicates that the routing number is invalid, the account number does not match the account number format associated with the routing number, or the account has been reported as closed or frozen. Only returned for Auth Items created via Database Auth.   
	VerificationStatus *string `json:"verification_status,omitempty"`
	// The account holder name that was used for micro-deposit and/or database verification. Only returned for Auth Items created via micro-deposit or database verification. This name was manually-entered by the user during Link, unless it was otherwise provided via the `user.legal_name` request field in `/link/token/create` for the Link session that created the Item.
	VerificationName *string `json:"verification_name,omitempty"`
	VerificationInsights *AccountVerificationInsights `json:"verification_insights,omitempty"`
	// A unique and persistent identifier for accounts that can be used to trace multiple instances of the same account across different Items for depository accounts. This field is currently supported only for Items at institutions that use Tokenized Account Numbers (i.e., Chase and PNC, and in May 2025 US Bank). Because these accounts have a different account number each time they are linked, this field may be used instead of the account number to uniquely identify an account across multiple Items for payments use cases, helping to reduce duplicate Items or attempted fraud. In Sandbox, this field is populated for TAN-based institutions (`ins_56`, `ins_13`) as well as the OAuth Sandbox institution (`ins_127287`); in Production, it will only be populated for accounts at applicable institutions.
	PersistentAccountId *string `json:"persistent_account_id,omitempty"`
	HolderCategory NullableAccountHolderCategory `json:"holder_category,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountBaseNullable AccountBaseNullable

// NewAccountBaseNullable instantiates a new AccountBaseNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBaseNullable(accountId string, balances AccountBalance, mask NullableString, name string, officialName NullableString, type_ AccountType, subtype NullableAccountSubtype) *AccountBaseNullable {
	this := AccountBaseNullable{}
	this.AccountId = accountId
	this.Balances = balances
	this.Mask = mask
	this.Name = name
	this.OfficialName = officialName
	this.Type = type_
	this.Subtype = subtype
	return &this
}

// NewAccountBaseNullableWithDefaults instantiates a new AccountBaseNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBaseNullableWithDefaults() *AccountBaseNullable {
	this := AccountBaseNullable{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AccountBaseNullable) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountBaseNullable) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountBaseNullable) SetAccountId(v string) {
	o.AccountId = v
}

// GetBalances returns the Balances field value
func (o *AccountBaseNullable) GetBalances() AccountBalance {
	if o == nil {
		var ret AccountBalance
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *AccountBaseNullable) GetBalancesOk() (*AccountBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *AccountBaseNullable) SetBalances(v AccountBalance) {
	o.Balances = v
}

// GetMask returns the Mask field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountBaseNullable) GetMask() string {
	if o == nil || o.Mask.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBaseNullable) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// SetMask sets field value
func (o *AccountBaseNullable) SetMask(v string) {
	o.Mask.Set(&v)
}

// GetName returns the Name field value
func (o *AccountBaseNullable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountBaseNullable) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountBaseNullable) SetName(v string) {
	o.Name = v
}

// GetOfficialName returns the OfficialName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountBaseNullable) GetOfficialName() string {
	if o == nil || o.OfficialName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OfficialName.Get()
}

// GetOfficialNameOk returns a tuple with the OfficialName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBaseNullable) GetOfficialNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OfficialName.Get(), o.OfficialName.IsSet()
}

// SetOfficialName sets field value
func (o *AccountBaseNullable) SetOfficialName(v string) {
	o.OfficialName.Set(&v)
}

// GetType returns the Type field value
func (o *AccountBaseNullable) GetType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountBaseNullable) GetTypeOk() (*AccountType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountBaseNullable) SetType(v AccountType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value
// If the value is explicit nil, the zero value for AccountSubtype will be returned
func (o *AccountBaseNullable) GetSubtype() AccountSubtype {
	if o == nil || o.Subtype.Get() == nil {
		var ret AccountSubtype
		return ret
	}

	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBaseNullable) GetSubtypeOk() (*AccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// SetSubtype sets field value
func (o *AccountBaseNullable) SetSubtype(v AccountSubtype) {
	o.Subtype.Set(&v)
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *AccountBaseNullable) GetVerificationStatus() string {
	if o == nil || o.VerificationStatus == nil {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBaseNullable) GetVerificationStatusOk() (*string, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *AccountBaseNullable) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *AccountBaseNullable) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

// GetVerificationName returns the VerificationName field value if set, zero value otherwise.
func (o *AccountBaseNullable) GetVerificationName() string {
	if o == nil || o.VerificationName == nil {
		var ret string
		return ret
	}
	return *o.VerificationName
}

// GetVerificationNameOk returns a tuple with the VerificationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBaseNullable) GetVerificationNameOk() (*string, bool) {
	if o == nil || o.VerificationName == nil {
		return nil, false
	}
	return o.VerificationName, true
}

// HasVerificationName returns a boolean if a field has been set.
func (o *AccountBaseNullable) HasVerificationName() bool {
	if o != nil && o.VerificationName != nil {
		return true
	}

	return false
}

// SetVerificationName gets a reference to the given string and assigns it to the VerificationName field.
func (o *AccountBaseNullable) SetVerificationName(v string) {
	o.VerificationName = &v
}

// GetVerificationInsights returns the VerificationInsights field value if set, zero value otherwise.
func (o *AccountBaseNullable) GetVerificationInsights() AccountVerificationInsights {
	if o == nil || o.VerificationInsights == nil {
		var ret AccountVerificationInsights
		return ret
	}
	return *o.VerificationInsights
}

// GetVerificationInsightsOk returns a tuple with the VerificationInsights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBaseNullable) GetVerificationInsightsOk() (*AccountVerificationInsights, bool) {
	if o == nil || o.VerificationInsights == nil {
		return nil, false
	}
	return o.VerificationInsights, true
}

// HasVerificationInsights returns a boolean if a field has been set.
func (o *AccountBaseNullable) HasVerificationInsights() bool {
	if o != nil && o.VerificationInsights != nil {
		return true
	}

	return false
}

// SetVerificationInsights gets a reference to the given AccountVerificationInsights and assigns it to the VerificationInsights field.
func (o *AccountBaseNullable) SetVerificationInsights(v AccountVerificationInsights) {
	o.VerificationInsights = &v
}

// GetPersistentAccountId returns the PersistentAccountId field value if set, zero value otherwise.
func (o *AccountBaseNullable) GetPersistentAccountId() string {
	if o == nil || o.PersistentAccountId == nil {
		var ret string
		return ret
	}
	return *o.PersistentAccountId
}

// GetPersistentAccountIdOk returns a tuple with the PersistentAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBaseNullable) GetPersistentAccountIdOk() (*string, bool) {
	if o == nil || o.PersistentAccountId == nil {
		return nil, false
	}
	return o.PersistentAccountId, true
}

// HasPersistentAccountId returns a boolean if a field has been set.
func (o *AccountBaseNullable) HasPersistentAccountId() bool {
	if o != nil && o.PersistentAccountId != nil {
		return true
	}

	return false
}

// SetPersistentAccountId gets a reference to the given string and assigns it to the PersistentAccountId field.
func (o *AccountBaseNullable) SetPersistentAccountId(v string) {
	o.PersistentAccountId = &v
}

// GetHolderCategory returns the HolderCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountBaseNullable) GetHolderCategory() AccountHolderCategory {
	if o == nil || o.HolderCategory.Get() == nil {
		var ret AccountHolderCategory
		return ret
	}
	return *o.HolderCategory.Get()
}

// GetHolderCategoryOk returns a tuple with the HolderCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBaseNullable) GetHolderCategoryOk() (*AccountHolderCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HolderCategory.Get(), o.HolderCategory.IsSet()
}

// HasHolderCategory returns a boolean if a field has been set.
func (o *AccountBaseNullable) HasHolderCategory() bool {
	if o != nil && o.HolderCategory.IsSet() {
		return true
	}

	return false
}

// SetHolderCategory gets a reference to the given NullableAccountHolderCategory and assigns it to the HolderCategory field.
func (o *AccountBaseNullable) SetHolderCategory(v AccountHolderCategory) {
	o.HolderCategory.Set(&v)
}
// SetHolderCategoryNil sets the value for HolderCategory to be an explicit nil
func (o *AccountBaseNullable) SetHolderCategoryNil() {
	o.HolderCategory.Set(nil)
}

// UnsetHolderCategory ensures that no value is present for HolderCategory, not even an explicit nil
func (o *AccountBaseNullable) UnsetHolderCategory() {
	o.HolderCategory.Unset()
}

func (o AccountBaseNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["balances"] = o.Balances
	}
	if true {
		toSerialize["mask"] = o.Mask.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["official_name"] = o.OfficialName.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	if o.VerificationStatus != nil {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if o.VerificationName != nil {
		toSerialize["verification_name"] = o.VerificationName
	}
	if o.VerificationInsights != nil {
		toSerialize["verification_insights"] = o.VerificationInsights
	}
	if o.PersistentAccountId != nil {
		toSerialize["persistent_account_id"] = o.PersistentAccountId
	}
	if o.HolderCategory.IsSet() {
		toSerialize["holder_category"] = o.HolderCategory.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountBaseNullable) UnmarshalJSON(bytes []byte) (err error) {
	varAccountBaseNullable := _AccountBaseNullable{}

	if err = json.Unmarshal(bytes, &varAccountBaseNullable); err == nil {
		*o = AccountBaseNullable(varAccountBaseNullable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "balances")
		delete(additionalProperties, "mask")
		delete(additionalProperties, "name")
		delete(additionalProperties, "official_name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "verification_status")
		delete(additionalProperties, "verification_name")
		delete(additionalProperties, "verification_insights")
		delete(additionalProperties, "persistent_account_id")
		delete(additionalProperties, "holder_category")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountBaseNullable struct {
	value *AccountBaseNullable
	isSet bool
}

func (v NullableAccountBaseNullable) Get() *AccountBaseNullable {
	return v.value
}

func (v *NullableAccountBaseNullable) Set(val *AccountBaseNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBaseNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBaseNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBaseNullable(val *AccountBaseNullable) *NullableAccountBaseNullable {
	return &NullableAccountBaseNullable{value: val, isSet: true}
}

func (v NullableAccountBaseNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBaseNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


