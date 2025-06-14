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

// TransactionsGetResponse TransactionsGetResponse defines the response schema for `/transactions/get`
type TransactionsGetResponse struct {
	// An array containing the `accounts` associated with the Item for which transactions are being returned. Each transaction can be mapped to its corresponding account via the `account_id` field.
	Accounts []AccountBase `json:"accounts"`
	// An array containing transactions from the account. Transactions are returned in reverse chronological order, with the most recent at the beginning of the array. The maximum number of transactions returned is determined by the `count` parameter.
	Transactions []Transaction `json:"transactions"`
	// The total number of transactions available within the date range specified. If `total_transactions` is larger than the size of the `transactions` array, more transactions are available and can be fetched via manipulating the `offset` parameter.
	TotalTransactions int32 `json:"total_transactions"`
	Item Item `json:"item"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransactionsGetResponse TransactionsGetResponse

// NewTransactionsGetResponse instantiates a new TransactionsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsGetResponse(accounts []AccountBase, transactions []Transaction, totalTransactions int32, item Item, requestId string) *TransactionsGetResponse {
	this := TransactionsGetResponse{}
	this.Accounts = accounts
	this.Transactions = transactions
	this.TotalTransactions = totalTransactions
	this.Item = item
	this.RequestId = requestId
	return &this
}

// NewTransactionsGetResponseWithDefaults instantiates a new TransactionsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsGetResponseWithDefaults() *TransactionsGetResponse {
	this := TransactionsGetResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *TransactionsGetResponse) GetAccounts() []AccountBase {
	if o == nil {
		var ret []AccountBase
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *TransactionsGetResponse) GetAccountsOk() (*[]AccountBase, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *TransactionsGetResponse) SetAccounts(v []AccountBase) {
	o.Accounts = v
}

// GetTransactions returns the Transactions field value
func (o *TransactionsGetResponse) GetTransactions() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *TransactionsGetResponse) GetTransactionsOk() (*[]Transaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *TransactionsGetResponse) SetTransactions(v []Transaction) {
	o.Transactions = v
}

// GetTotalTransactions returns the TotalTransactions field value
func (o *TransactionsGetResponse) GetTotalTransactions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTransactions
}

// GetTotalTransactionsOk returns a tuple with the TotalTransactions field value
// and a boolean to check if the value has been set.
func (o *TransactionsGetResponse) GetTotalTransactionsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalTransactions, true
}

// SetTotalTransactions sets field value
func (o *TransactionsGetResponse) SetTotalTransactions(v int32) {
	o.TotalTransactions = v
}

// GetItem returns the Item field value
func (o *TransactionsGetResponse) GetItem() Item {
	if o == nil {
		var ret Item
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *TransactionsGetResponse) GetItemOk() (*Item, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *TransactionsGetResponse) SetItem(v Item) {
	o.Item = v
}

// GetRequestId returns the RequestId field value
func (o *TransactionsGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransactionsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransactionsGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransactionsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	if true {
		toSerialize["total_transactions"] = o.TotalTransactions
	}
	if true {
		toSerialize["item"] = o.Item
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionsGetResponse := _TransactionsGetResponse{}

	if err = json.Unmarshal(bytes, &varTransactionsGetResponse); err == nil {
		*o = TransactionsGetResponse(varTransactionsGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "transactions")
		delete(additionalProperties, "total_transactions")
		delete(additionalProperties, "item")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionsGetResponse struct {
	value *TransactionsGetResponse
	isSet bool
}

func (v NullableTransactionsGetResponse) Get() *TransactionsGetResponse {
	return v.value
}

func (v *NullableTransactionsGetResponse) Set(val *TransactionsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsGetResponse(val *TransactionsGetResponse) *NullableTransactionsGetResponse {
	return &NullableTransactionsGetResponse{value: val, isSet: true}
}

func (v NullableTransactionsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


