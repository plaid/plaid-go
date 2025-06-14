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

// LinkTokenGetMetadataResponse An object specifying the arguments originally provided to the `/link/token/create` call.
type LinkTokenGetMetadataResponse struct {
	// The `products` specified in the `/link/token/create` call.
	InitialProducts []Products `json:"initial_products"`
	// The `webhook` specified in the `/link/token/create` call.
	Webhook NullableString `json:"webhook"`
	// The `country_codes` specified in the `/link/token/create` call.
	CountryCodes []CountryCode `json:"country_codes"`
	// The `language` specified in the `/link/token/create` call.
	Language NullableString `json:"language"`
	InstitutionData *LinkTokenCreateInstitutionData `json:"institution_data,omitempty"`
	AccountFilters *AccountFiltersResponse `json:"account_filters,omitempty"`
	// The `redirect_uri` specified in the `/link/token/create` call.
	RedirectUri NullableString `json:"redirect_uri"`
	// The `client_name` specified in the `/link/token/create` call.
	ClientName NullableString `json:"client_name"`
	AdditionalProperties map[string]interface{}
}

type _LinkTokenGetMetadataResponse LinkTokenGetMetadataResponse

// NewLinkTokenGetMetadataResponse instantiates a new LinkTokenGetMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenGetMetadataResponse(initialProducts []Products, webhook NullableString, countryCodes []CountryCode, language NullableString, redirectUri NullableString, clientName NullableString) *LinkTokenGetMetadataResponse {
	this := LinkTokenGetMetadataResponse{}
	this.InitialProducts = initialProducts
	this.Webhook = webhook
	this.CountryCodes = countryCodes
	this.Language = language
	this.RedirectUri = redirectUri
	this.ClientName = clientName
	return &this
}

// NewLinkTokenGetMetadataResponseWithDefaults instantiates a new LinkTokenGetMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenGetMetadataResponseWithDefaults() *LinkTokenGetMetadataResponse {
	this := LinkTokenGetMetadataResponse{}
	return &this
}

// GetInitialProducts returns the InitialProducts field value
func (o *LinkTokenGetMetadataResponse) GetInitialProducts() []Products {
	if o == nil {
		var ret []Products
		return ret
	}

	return o.InitialProducts
}

// GetInitialProductsOk returns a tuple with the InitialProducts field value
// and a boolean to check if the value has been set.
func (o *LinkTokenGetMetadataResponse) GetInitialProductsOk() (*[]Products, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InitialProducts, true
}

// SetInitialProducts sets field value
func (o *LinkTokenGetMetadataResponse) SetInitialProducts(v []Products) {
	o.InitialProducts = v
}

// GetWebhook returns the Webhook field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LinkTokenGetMetadataResponse) GetWebhook() string {
	if o == nil || o.Webhook.Get() == nil {
		var ret string
		return ret
	}

	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkTokenGetMetadataResponse) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// SetWebhook sets field value
func (o *LinkTokenGetMetadataResponse) SetWebhook(v string) {
	o.Webhook.Set(&v)
}

// GetCountryCodes returns the CountryCodes field value
func (o *LinkTokenGetMetadataResponse) GetCountryCodes() []CountryCode {
	if o == nil {
		var ret []CountryCode
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *LinkTokenGetMetadataResponse) GetCountryCodesOk() (*[]CountryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *LinkTokenGetMetadataResponse) SetCountryCodes(v []CountryCode) {
	o.CountryCodes = v
}

// GetLanguage returns the Language field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LinkTokenGetMetadataResponse) GetLanguage() string {
	if o == nil || o.Language.Get() == nil {
		var ret string
		return ret
	}

	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkTokenGetMetadataResponse) GetLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// SetLanguage sets field value
func (o *LinkTokenGetMetadataResponse) SetLanguage(v string) {
	o.Language.Set(&v)
}

// GetInstitutionData returns the InstitutionData field value if set, zero value otherwise.
func (o *LinkTokenGetMetadataResponse) GetInstitutionData() LinkTokenCreateInstitutionData {
	if o == nil || o.InstitutionData == nil {
		var ret LinkTokenCreateInstitutionData
		return ret
	}
	return *o.InstitutionData
}

// GetInstitutionDataOk returns a tuple with the InstitutionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenGetMetadataResponse) GetInstitutionDataOk() (*LinkTokenCreateInstitutionData, bool) {
	if o == nil || o.InstitutionData == nil {
		return nil, false
	}
	return o.InstitutionData, true
}

// HasInstitutionData returns a boolean if a field has been set.
func (o *LinkTokenGetMetadataResponse) HasInstitutionData() bool {
	if o != nil && o.InstitutionData != nil {
		return true
	}

	return false
}

// SetInstitutionData gets a reference to the given LinkTokenCreateInstitutionData and assigns it to the InstitutionData field.
func (o *LinkTokenGetMetadataResponse) SetInstitutionData(v LinkTokenCreateInstitutionData) {
	o.InstitutionData = &v
}

// GetAccountFilters returns the AccountFilters field value if set, zero value otherwise.
func (o *LinkTokenGetMetadataResponse) GetAccountFilters() AccountFiltersResponse {
	if o == nil || o.AccountFilters == nil {
		var ret AccountFiltersResponse
		return ret
	}
	return *o.AccountFilters
}

// GetAccountFiltersOk returns a tuple with the AccountFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenGetMetadataResponse) GetAccountFiltersOk() (*AccountFiltersResponse, bool) {
	if o == nil || o.AccountFilters == nil {
		return nil, false
	}
	return o.AccountFilters, true
}

// HasAccountFilters returns a boolean if a field has been set.
func (o *LinkTokenGetMetadataResponse) HasAccountFilters() bool {
	if o != nil && o.AccountFilters != nil {
		return true
	}

	return false
}

// SetAccountFilters gets a reference to the given AccountFiltersResponse and assigns it to the AccountFilters field.
func (o *LinkTokenGetMetadataResponse) SetAccountFilters(v AccountFiltersResponse) {
	o.AccountFilters = &v
}

// GetRedirectUri returns the RedirectUri field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LinkTokenGetMetadataResponse) GetRedirectUri() string {
	if o == nil || o.RedirectUri.Get() == nil {
		var ret string
		return ret
	}

	return *o.RedirectUri.Get()
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkTokenGetMetadataResponse) GetRedirectUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RedirectUri.Get(), o.RedirectUri.IsSet()
}

// SetRedirectUri sets field value
func (o *LinkTokenGetMetadataResponse) SetRedirectUri(v string) {
	o.RedirectUri.Set(&v)
}

// GetClientName returns the ClientName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LinkTokenGetMetadataResponse) GetClientName() string {
	if o == nil || o.ClientName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientName.Get()
}

// GetClientNameOk returns a tuple with the ClientName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkTokenGetMetadataResponse) GetClientNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientName.Get(), o.ClientName.IsSet()
}

// SetClientName sets field value
func (o *LinkTokenGetMetadataResponse) SetClientName(v string) {
	o.ClientName.Set(&v)
}

func (o LinkTokenGetMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["initial_products"] = o.InitialProducts
	}
	if true {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	if true {
		toSerialize["country_codes"] = o.CountryCodes
	}
	if true {
		toSerialize["language"] = o.Language.Get()
	}
	if o.InstitutionData != nil {
		toSerialize["institution_data"] = o.InstitutionData
	}
	if o.AccountFilters != nil {
		toSerialize["account_filters"] = o.AccountFilters
	}
	if true {
		toSerialize["redirect_uri"] = o.RedirectUri.Get()
	}
	if true {
		toSerialize["client_name"] = o.ClientName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTokenGetMetadataResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTokenGetMetadataResponse := _LinkTokenGetMetadataResponse{}

	if err = json.Unmarshal(bytes, &varLinkTokenGetMetadataResponse); err == nil {
		*o = LinkTokenGetMetadataResponse(varLinkTokenGetMetadataResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "initial_products")
		delete(additionalProperties, "webhook")
		delete(additionalProperties, "country_codes")
		delete(additionalProperties, "language")
		delete(additionalProperties, "institution_data")
		delete(additionalProperties, "account_filters")
		delete(additionalProperties, "redirect_uri")
		delete(additionalProperties, "client_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTokenGetMetadataResponse struct {
	value *LinkTokenGetMetadataResponse
	isSet bool
}

func (v NullableLinkTokenGetMetadataResponse) Get() *LinkTokenGetMetadataResponse {
	return v.value
}

func (v *NullableLinkTokenGetMetadataResponse) Set(val *LinkTokenGetMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenGetMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenGetMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenGetMetadataResponse(val *LinkTokenGetMetadataResponse) *NullableLinkTokenGetMetadataResponse {
	return &NullableLinkTokenGetMetadataResponse{value: val, isSet: true}
}

func (v NullableLinkTokenGetMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenGetMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


