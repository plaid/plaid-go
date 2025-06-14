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

// UserAccountRevokedWebhook The `USER_ACCOUNT_REVOKED` webhook is fired when an end user has revoked access to their account on the Data Provider's portal. This webhook is currently sent only for Chase and PNC Items, but may be sent in the future for other financial institutions that allow account-level permissions revocation through their portals. Upon receiving this webhook, it is recommended to delete any Plaid-derived data you have stored that is associated with the revoked account.  If you are using Auth and receive this webhook, this webhook indicates that the TAN associated with the revoked account is no longer valid and cannot be used to create new transfers. You should not create new ACH transfers for the account that was revoked until access has been re-granted.  You can request the user to re-grant access to their account by sending them through [update mode](https://plaid.com/docs/link/update-mode). Alternatively, they may re-grant access directly through the Data Provider's portal.  After the user has re-granted access, Auth customers should call the auth endpoint again to obtain the new TAN.
type UserAccountRevokedWebhook struct {
	// `ITEM`
	WebhookType string `json:"webhook_type"`
	// `USER_ACCOUNT_REVOKED`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	// The external account ID of the affected account
	AccountId string `json:"account_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _UserAccountRevokedWebhook UserAccountRevokedWebhook

// NewUserAccountRevokedWebhook instantiates a new UserAccountRevokedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccountRevokedWebhook(webhookType string, webhookCode string, itemId string, accountId string, environment WebhookEnvironmentValues) *UserAccountRevokedWebhook {
	this := UserAccountRevokedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.AccountId = accountId
	this.Environment = environment
	return &this
}

// NewUserAccountRevokedWebhookWithDefaults instantiates a new UserAccountRevokedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccountRevokedWebhookWithDefaults() *UserAccountRevokedWebhook {
	this := UserAccountRevokedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *UserAccountRevokedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *UserAccountRevokedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *UserAccountRevokedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *UserAccountRevokedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *UserAccountRevokedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *UserAccountRevokedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *UserAccountRevokedWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *UserAccountRevokedWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *UserAccountRevokedWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetAccountId returns the AccountId field value
func (o *UserAccountRevokedWebhook) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *UserAccountRevokedWebhook) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *UserAccountRevokedWebhook) SetAccountId(v string) {
	o.AccountId = v
}

// GetEnvironment returns the Environment field value
func (o *UserAccountRevokedWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *UserAccountRevokedWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *UserAccountRevokedWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o UserAccountRevokedWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserAccountRevokedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varUserAccountRevokedWebhook := _UserAccountRevokedWebhook{}

	if err = json.Unmarshal(bytes, &varUserAccountRevokedWebhook); err == nil {
		*o = UserAccountRevokedWebhook(varUserAccountRevokedWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserAccountRevokedWebhook struct {
	value *UserAccountRevokedWebhook
	isSet bool
}

func (v NullableUserAccountRevokedWebhook) Get() *UserAccountRevokedWebhook {
	return v.value
}

func (v *NullableUserAccountRevokedWebhook) Set(val *UserAccountRevokedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccountRevokedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccountRevokedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccountRevokedWebhook(val *UserAccountRevokedWebhook) *NullableUserAccountRevokedWebhook {
	return &NullableUserAccountRevokedWebhook{value: val, isSet: true}
}

func (v NullableUserAccountRevokedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccountRevokedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


