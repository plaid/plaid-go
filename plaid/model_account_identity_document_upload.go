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
)

// AccountIdentityDocumentUpload Identity information about an account
type AccountIdentityDocumentUpload struct {
	// Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  The `account_id` can also change if the `access_token` is deleted and the same credentials that were used to generate that `access_token` are used to generate a new `access_token` on a later date. In that case, the new `account_id` will be different from the old `account_id`.  If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  Like all Plaid identifiers, the `account_id` is case sensitive.
	AccountId string `json:"account_id"`
	Balances AccountBalance `json:"balances"`
	// The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.
	Mask NullableString `json:"mask"`
	// The name of the account, either assigned by the user or by the financial institution itself
	Name string `json:"name"`
	// The official name of the account as given by the financial institution
	OfficialName NullableString `json:"official_name"`
	Type AccountType `json:"type"`
	Subtype NullableAccountSubtype `json:"subtype"`
	// The current verification status of an Auth Item initiated through micro-deposits or database verification. Returned for Auth Items only.  `pending_automatic_verification`: The Item is pending automatic verification  `pending_manual_verification`: The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the micro-deposit.  `automatically_verified`: The Item has successfully been automatically verified   `manually_verified`: The Item has successfully been manually verified  `verification_expired`: Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.  `verification_failed`: The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.  `database_matched`: The Item has successfully been verified using Plaid's data sources. Only returned for Auth Items created via Database Match.  `database_insights_pass`: The Item's numbers have been verified using Plaid's data sources and have strong signal for being valid. Only returned for Auth Items created via Database Insights. Note: Database Insights is currently a beta feature, please contact your account manager for more information.  `database_insights_pass_with_caution`: The Item's numbers have been verified using Plaid's data sources and have some signal for being valid. Only returned for Auth Items created via Database Insights. Note: Database Insights is currently a beta feature, please contact your account manager for more information.  `database_insights_fail`:  The Item's numbers have been verified using Plaid's data sources and have signal for being invalid and/or have no signal for being valid. Only returned for Auth Items created via Database Insights. Note: Database Insights is currently a beta feature, please contact your account manager for more information.   
	VerificationStatus *string `json:"verification_status,omitempty"`
	VerificationInsights *AccountVerificationInsights `json:"verification_insights,omitempty"`
	// A unique and persistent identifier for accounts that can be used to trace multiple instances of the same account across different Items for depository accounts. This field is currently supported only for Items at institutions that use Tokenized Account Numbers (i.e., Chase and PNC). Because these accounts have a different account number each time they are linked, this field may be used instead of the account number to uniquely identify an account across multiple Items for payments use cases, helping to reduce duplicate Items or attempted fraud. In Sandbox, this field may be populated for any account; in Production, it will only be populated for accounts at applicable institutions.
	PersistentAccountId *string `json:"persistent_account_id,omitempty"`
	HolderCategory NullableAccountHolderCategory `json:"holder_category,omitempty"`
	// Data returned by the financial institution about the account owner or owners. Only returned by Identity or Assets endpoints. For business accounts, the name reported may be either the name of the individual or the name of the business, depending on the institution; detecting whether the linked account is a business account is not currently supported. Multiple owners on a single account will be represented in the same `owner` object, not in multiple owner objects within the array. In API versions 2018-05-22 and earlier, the `owners` object is not returned, and instead identity information is returned in the top level `identity` object. For more details, see [Plaid API versioning](https://plaid.com/docs/api/versioning/#version-2019-05-29)
	Owners []Owner `json:"owners"`
	// Data about the documents that were uploaded as proof of account ownership.
	Documents []IdentityDocumentUpload `json:"documents,omitempty"`
}

// NewAccountIdentityDocumentUpload instantiates a new AccountIdentityDocumentUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountIdentityDocumentUpload(accountId string, balances AccountBalance, mask NullableString, name string, officialName NullableString, type_ AccountType, subtype NullableAccountSubtype, owners []Owner) *AccountIdentityDocumentUpload {
	this := AccountIdentityDocumentUpload{}
	this.AccountId = accountId
	this.Balances = balances
	this.Mask = mask
	this.Name = name
	this.OfficialName = officialName
	this.Type = type_
	this.Subtype = subtype
	this.Owners = owners
	return &this
}

// NewAccountIdentityDocumentUploadWithDefaults instantiates a new AccountIdentityDocumentUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountIdentityDocumentUploadWithDefaults() *AccountIdentityDocumentUpload {
	this := AccountIdentityDocumentUpload{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AccountIdentityDocumentUpload) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountIdentityDocumentUpload) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountIdentityDocumentUpload) SetAccountId(v string) {
	o.AccountId = v
}

// GetBalances returns the Balances field value
func (o *AccountIdentityDocumentUpload) GetBalances() AccountBalance {
	if o == nil {
		var ret AccountBalance
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *AccountIdentityDocumentUpload) GetBalancesOk() (*AccountBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *AccountIdentityDocumentUpload) SetBalances(v AccountBalance) {
	o.Balances = v
}

// GetMask returns the Mask field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountIdentityDocumentUpload) GetMask() string {
	if o == nil || o.Mask.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityDocumentUpload) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// SetMask sets field value
func (o *AccountIdentityDocumentUpload) SetMask(v string) {
	o.Mask.Set(&v)
}

// GetName returns the Name field value
func (o *AccountIdentityDocumentUpload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountIdentityDocumentUpload) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountIdentityDocumentUpload) SetName(v string) {
	o.Name = v
}

// GetOfficialName returns the OfficialName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountIdentityDocumentUpload) GetOfficialName() string {
	if o == nil || o.OfficialName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OfficialName.Get()
}

// GetOfficialNameOk returns a tuple with the OfficialName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityDocumentUpload) GetOfficialNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OfficialName.Get(), o.OfficialName.IsSet()
}

// SetOfficialName sets field value
func (o *AccountIdentityDocumentUpload) SetOfficialName(v string) {
	o.OfficialName.Set(&v)
}

// GetType returns the Type field value
func (o *AccountIdentityDocumentUpload) GetType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountIdentityDocumentUpload) GetTypeOk() (*AccountType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountIdentityDocumentUpload) SetType(v AccountType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value
// If the value is explicit nil, the zero value for AccountSubtype will be returned
func (o *AccountIdentityDocumentUpload) GetSubtype() AccountSubtype {
	if o == nil || o.Subtype.Get() == nil {
		var ret AccountSubtype
		return ret
	}

	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityDocumentUpload) GetSubtypeOk() (*AccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// SetSubtype sets field value
func (o *AccountIdentityDocumentUpload) SetSubtype(v AccountSubtype) {
	o.Subtype.Set(&v)
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *AccountIdentityDocumentUpload) GetVerificationStatus() string {
	if o == nil || o.VerificationStatus == nil {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountIdentityDocumentUpload) GetVerificationStatusOk() (*string, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *AccountIdentityDocumentUpload) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *AccountIdentityDocumentUpload) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

// GetVerificationInsights returns the VerificationInsights field value if set, zero value otherwise.
func (o *AccountIdentityDocumentUpload) GetVerificationInsights() AccountVerificationInsights {
	if o == nil || o.VerificationInsights == nil {
		var ret AccountVerificationInsights
		return ret
	}
	return *o.VerificationInsights
}

// GetVerificationInsightsOk returns a tuple with the VerificationInsights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountIdentityDocumentUpload) GetVerificationInsightsOk() (*AccountVerificationInsights, bool) {
	if o == nil || o.VerificationInsights == nil {
		return nil, false
	}
	return o.VerificationInsights, true
}

// HasVerificationInsights returns a boolean if a field has been set.
func (o *AccountIdentityDocumentUpload) HasVerificationInsights() bool {
	if o != nil && o.VerificationInsights != nil {
		return true
	}

	return false
}

// SetVerificationInsights gets a reference to the given AccountVerificationInsights and assigns it to the VerificationInsights field.
func (o *AccountIdentityDocumentUpload) SetVerificationInsights(v AccountVerificationInsights) {
	o.VerificationInsights = &v
}

// GetPersistentAccountId returns the PersistentAccountId field value if set, zero value otherwise.
func (o *AccountIdentityDocumentUpload) GetPersistentAccountId() string {
	if o == nil || o.PersistentAccountId == nil {
		var ret string
		return ret
	}
	return *o.PersistentAccountId
}

// GetPersistentAccountIdOk returns a tuple with the PersistentAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountIdentityDocumentUpload) GetPersistentAccountIdOk() (*string, bool) {
	if o == nil || o.PersistentAccountId == nil {
		return nil, false
	}
	return o.PersistentAccountId, true
}

// HasPersistentAccountId returns a boolean if a field has been set.
func (o *AccountIdentityDocumentUpload) HasPersistentAccountId() bool {
	if o != nil && o.PersistentAccountId != nil {
		return true
	}

	return false
}

// SetPersistentAccountId gets a reference to the given string and assigns it to the PersistentAccountId field.
func (o *AccountIdentityDocumentUpload) SetPersistentAccountId(v string) {
	o.PersistentAccountId = &v
}

// GetHolderCategory returns the HolderCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentityDocumentUpload) GetHolderCategory() AccountHolderCategory {
	if o == nil || o.HolderCategory.Get() == nil {
		var ret AccountHolderCategory
		return ret
	}
	return *o.HolderCategory.Get()
}

// GetHolderCategoryOk returns a tuple with the HolderCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityDocumentUpload) GetHolderCategoryOk() (*AccountHolderCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HolderCategory.Get(), o.HolderCategory.IsSet()
}

// HasHolderCategory returns a boolean if a field has been set.
func (o *AccountIdentityDocumentUpload) HasHolderCategory() bool {
	if o != nil && o.HolderCategory.IsSet() {
		return true
	}

	return false
}

// SetHolderCategory gets a reference to the given NullableAccountHolderCategory and assigns it to the HolderCategory field.
func (o *AccountIdentityDocumentUpload) SetHolderCategory(v AccountHolderCategory) {
	o.HolderCategory.Set(&v)
}
// SetHolderCategoryNil sets the value for HolderCategory to be an explicit nil
func (o *AccountIdentityDocumentUpload) SetHolderCategoryNil() {
	o.HolderCategory.Set(nil)
}

// UnsetHolderCategory ensures that no value is present for HolderCategory, not even an explicit nil
func (o *AccountIdentityDocumentUpload) UnsetHolderCategory() {
	o.HolderCategory.Unset()
}

// GetOwners returns the Owners field value
func (o *AccountIdentityDocumentUpload) GetOwners() []Owner {
	if o == nil {
		var ret []Owner
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *AccountIdentityDocumentUpload) GetOwnersOk() (*[]Owner, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value
func (o *AccountIdentityDocumentUpload) SetOwners(v []Owner) {
	o.Owners = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentityDocumentUpload) GetDocuments() []IdentityDocumentUpload {
	if o == nil  {
		var ret []IdentityDocumentUpload
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentityDocumentUpload) GetDocumentsOk() (*[]IdentityDocumentUpload, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return &o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *AccountIdentityDocumentUpload) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []IdentityDocumentUpload and assigns it to the Documents field.
func (o *AccountIdentityDocumentUpload) SetDocuments(v []IdentityDocumentUpload) {
	o.Documents = v
}

func (o AccountIdentityDocumentUpload) MarshalJSON() ([]byte, error) {
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
	if o.VerificationInsights != nil {
		toSerialize["verification_insights"] = o.VerificationInsights
	}
	if o.PersistentAccountId != nil {
		toSerialize["persistent_account_id"] = o.PersistentAccountId
	}
	if o.HolderCategory.IsSet() {
		toSerialize["holder_category"] = o.HolderCategory.Get()
	}
	if true {
		toSerialize["owners"] = o.Owners
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	return json.Marshal(toSerialize)
}

type NullableAccountIdentityDocumentUpload struct {
	value *AccountIdentityDocumentUpload
	isSet bool
}

func (v NullableAccountIdentityDocumentUpload) Get() *AccountIdentityDocumentUpload {
	return v.value
}

func (v *NullableAccountIdentityDocumentUpload) Set(val *AccountIdentityDocumentUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountIdentityDocumentUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountIdentityDocumentUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountIdentityDocumentUpload(val *AccountIdentityDocumentUpload) *NullableAccountIdentityDocumentUpload {
	return &NullableAccountIdentityDocumentUpload{value: val, isSet: true}
}

func (v NullableAccountIdentityDocumentUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountIdentityDocumentUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

