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

// MFA Specifies the multi-factor authentication settings to use with this test account
type MFA struct {
	// Possible values are `device`, `selections`, or `questions`.  If value is `device`, the MFA answer is `1234`.  If value is `selections`, the MFA answer is always the first option.  If value is `questions`, the MFA answer is  `answer_<i>_<j>` for the j-th question in the i-th round, starting from 0. For example, the answer to the first question in the second round is `answer_1_0`.
	Type string `json:"type"`
	// Number of rounds of questions. Required if value of `type` is `questions`. 
	QuestionRounds float32 `json:"question_rounds"`
	// Number of questions per round. Required if value of `type` is `questions`. If value of type is `selections`, default value is 2.
	QuestionsPerRound float32 `json:"questions_per_round"`
	// Number of rounds of selections, used if `type` is `selections`. Defaults to 1.
	SelectionRounds float32 `json:"selection_rounds"`
	// Number of available answers per question, used if `type` is `selection`. Defaults to 2. 
	SelectionsPerQuestion float32 `json:"selections_per_question"`
	AdditionalProperties map[string]interface{}
}

type _MFA MFA

// NewMFA instantiates a new MFA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFA(type_ string, questionRounds float32, questionsPerRound float32, selectionRounds float32, selectionsPerQuestion float32) *MFA {
	this := MFA{}
	this.Type = type_
	this.QuestionRounds = questionRounds
	this.QuestionsPerRound = questionsPerRound
	this.SelectionRounds = selectionRounds
	this.SelectionsPerQuestion = selectionsPerQuestion
	return &this
}

// NewMFAWithDefaults instantiates a new MFA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAWithDefaults() *MFA {
	this := MFA{}
	return &this
}

// GetType returns the Type field value
func (o *MFA) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MFA) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MFA) SetType(v string) {
	o.Type = v
}

// GetQuestionRounds returns the QuestionRounds field value
func (o *MFA) GetQuestionRounds() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.QuestionRounds
}

// GetQuestionRoundsOk returns a tuple with the QuestionRounds field value
// and a boolean to check if the value has been set.
func (o *MFA) GetQuestionRoundsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QuestionRounds, true
}

// SetQuestionRounds sets field value
func (o *MFA) SetQuestionRounds(v float32) {
	o.QuestionRounds = v
}

// GetQuestionsPerRound returns the QuestionsPerRound field value
func (o *MFA) GetQuestionsPerRound() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.QuestionsPerRound
}

// GetQuestionsPerRoundOk returns a tuple with the QuestionsPerRound field value
// and a boolean to check if the value has been set.
func (o *MFA) GetQuestionsPerRoundOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QuestionsPerRound, true
}

// SetQuestionsPerRound sets field value
func (o *MFA) SetQuestionsPerRound(v float32) {
	o.QuestionsPerRound = v
}

// GetSelectionRounds returns the SelectionRounds field value
func (o *MFA) GetSelectionRounds() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SelectionRounds
}

// GetSelectionRoundsOk returns a tuple with the SelectionRounds field value
// and a boolean to check if the value has been set.
func (o *MFA) GetSelectionRoundsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SelectionRounds, true
}

// SetSelectionRounds sets field value
func (o *MFA) SetSelectionRounds(v float32) {
	o.SelectionRounds = v
}

// GetSelectionsPerQuestion returns the SelectionsPerQuestion field value
func (o *MFA) GetSelectionsPerQuestion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SelectionsPerQuestion
}

// GetSelectionsPerQuestionOk returns a tuple with the SelectionsPerQuestion field value
// and a boolean to check if the value has been set.
func (o *MFA) GetSelectionsPerQuestionOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SelectionsPerQuestion, true
}

// SetSelectionsPerQuestion sets field value
func (o *MFA) SetSelectionsPerQuestion(v float32) {
	o.SelectionsPerQuestion = v
}

func (o MFA) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["question_rounds"] = o.QuestionRounds
	}
	if true {
		toSerialize["questions_per_round"] = o.QuestionsPerRound
	}
	if true {
		toSerialize["selection_rounds"] = o.SelectionRounds
	}
	if true {
		toSerialize["selections_per_question"] = o.SelectionsPerQuestion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MFA) UnmarshalJSON(bytes []byte) (err error) {
	varMFA := _MFA{}

	if err = json.Unmarshal(bytes, &varMFA); err == nil {
		*o = MFA(varMFA)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "question_rounds")
		delete(additionalProperties, "questions_per_round")
		delete(additionalProperties, "selection_rounds")
		delete(additionalProperties, "selections_per_question")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMFA struct {
	value *MFA
	isSet bool
}

func (v NullableMFA) Get() *MFA {
	return v.value
}

func (v *NullableMFA) Set(val *MFA) {
	v.value = val
	v.isSet = true
}

func (v NullableMFA) IsSet() bool {
	return v.isSet
}

func (v *NullableMFA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFA(val *MFA) *NullableMFA {
	return &NullableMFA{value: val, isSet: true}
}

func (v NullableMFA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


