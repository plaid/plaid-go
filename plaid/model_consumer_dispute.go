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

// ConsumerDispute The information about a previously submitted valid dispute statement by the consumer
type ConsumerDispute struct {
	// (Deprecated) A unique identifier (UUID) of the consumer dispute that can be used for troubleshooting
	ConsumerDisputeId string `json:"consumer_dispute_id"`
	// Date of the disputed field (e.g. transaction date), in an ISO 8601 format (YYYY-MM-DD)
	DisputeFieldCreateDate string `json:"dispute_field_create_date"`
	Category ConsumerDisputeCategory `json:"category"`
	// Text content of dispute
	Statement string `json:"statement"`
	AdditionalProperties map[string]interface{}
}

type _ConsumerDispute ConsumerDispute

// NewConsumerDispute instantiates a new ConsumerDispute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerDispute(consumerDisputeId string, disputeFieldCreateDate string, category ConsumerDisputeCategory, statement string) *ConsumerDispute {
	this := ConsumerDispute{}
	this.ConsumerDisputeId = consumerDisputeId
	this.DisputeFieldCreateDate = disputeFieldCreateDate
	this.Category = category
	this.Statement = statement
	return &this
}

// NewConsumerDisputeWithDefaults instantiates a new ConsumerDispute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerDisputeWithDefaults() *ConsumerDispute {
	this := ConsumerDispute{}
	return &this
}

// GetConsumerDisputeId returns the ConsumerDisputeId field value
func (o *ConsumerDispute) GetConsumerDisputeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerDisputeId
}

// GetConsumerDisputeIdOk returns a tuple with the ConsumerDisputeId field value
// and a boolean to check if the value has been set.
func (o *ConsumerDispute) GetConsumerDisputeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConsumerDisputeId, true
}

// SetConsumerDisputeId sets field value
func (o *ConsumerDispute) SetConsumerDisputeId(v string) {
	o.ConsumerDisputeId = v
}

// GetDisputeFieldCreateDate returns the DisputeFieldCreateDate field value
func (o *ConsumerDispute) GetDisputeFieldCreateDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisputeFieldCreateDate
}

// GetDisputeFieldCreateDateOk returns a tuple with the DisputeFieldCreateDate field value
// and a boolean to check if the value has been set.
func (o *ConsumerDispute) GetDisputeFieldCreateDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisputeFieldCreateDate, true
}

// SetDisputeFieldCreateDate sets field value
func (o *ConsumerDispute) SetDisputeFieldCreateDate(v string) {
	o.DisputeFieldCreateDate = v
}

// GetCategory returns the Category field value
func (o *ConsumerDispute) GetCategory() ConsumerDisputeCategory {
	if o == nil {
		var ret ConsumerDisputeCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ConsumerDispute) GetCategoryOk() (*ConsumerDisputeCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *ConsumerDispute) SetCategory(v ConsumerDisputeCategory) {
	o.Category = v
}

// GetStatement returns the Statement field value
func (o *ConsumerDispute) GetStatement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Statement
}

// GetStatementOk returns a tuple with the Statement field value
// and a boolean to check if the value has been set.
func (o *ConsumerDispute) GetStatementOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Statement, true
}

// SetStatement sets field value
func (o *ConsumerDispute) SetStatement(v string) {
	o.Statement = v
}

func (o ConsumerDispute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["consumer_dispute_id"] = o.ConsumerDisputeId
	}
	if true {
		toSerialize["dispute_field_create_date"] = o.DisputeFieldCreateDate
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["statement"] = o.Statement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConsumerDispute) UnmarshalJSON(bytes []byte) (err error) {
	varConsumerDispute := _ConsumerDispute{}

	if err = json.Unmarshal(bytes, &varConsumerDispute); err == nil {
		*o = ConsumerDispute(varConsumerDispute)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "consumer_dispute_id")
		delete(additionalProperties, "dispute_field_create_date")
		delete(additionalProperties, "category")
		delete(additionalProperties, "statement")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConsumerDispute struct {
	value *ConsumerDispute
	isSet bool
}

func (v NullableConsumerDispute) Get() *ConsumerDispute {
	return v.value
}

func (v *NullableConsumerDispute) Set(val *ConsumerDispute) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerDispute) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerDispute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerDispute(val *ConsumerDispute) *NullableConsumerDispute {
	return &NullableConsumerDispute{value: val, isSet: true}
}

func (v NullableConsumerDispute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerDispute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


