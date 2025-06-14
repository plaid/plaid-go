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

// CashflowReportMonthlySummary Monthly summary statistics derived from transaction-level data.
type CashflowReportMonthlySummary struct {
	// The start date of the period covered in this monthly summary.  This date will be the first day of the month, unless the month being covered is a partial month because it is the first month included in the summary and the date range being requested does not begin with the first day of the month.  The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	StartDate string `json:"start_date"`
	// The end date of the period included in this monthly summary.  This date will be the last day of the month, unless the month being covered is a partial month because it is the last month included in the summary and the date range being requested does not end with the last day of the month.  The date will be returned in an ISO 8601 format (YYYY-MM-DD).
	EndDate string `json:"end_date"`
	StartingBalance NullableCashflowReportMonthlySummaryStartingBalance `json:"starting_balance"`
	EndingBalance NullableCashflowReportMonthlySummaryEndingBalance `json:"ending_balance"`
	AverageDailyEndingBalance CashflowReportMonthlySummaryAverageDailyEndingBalance `json:"average_daily_ending_balance"`
	AverageDailyInflowAmount CashflowReportMonthlySummaryAverageDailyInflowAmount `json:"average_daily_inflow_amount"`
	AverageDailyOutflowAmount CashflowReportMonthlySummaryAverageDailyOutflowAmount `json:"average_daily_outflow_amount"`
	AverageDailyNetCashflowAmount CashflowReportMonthlySummaryAverageDailyNetCashflowAmount `json:"average_daily_net_cashflow_amount"`
	// The average count of the number of daily inflow transactions. Rounded to 2 decimal places.
	AverageDailyInflowTransactionCount float32 `json:"average_daily_inflow_transaction_count"`
	// The average count of the number of daily outflow transactions. Rounded to 2 decimal places.
	AverageDailyOutflowTransactionCount float32 `json:"average_daily_outflow_transaction_count"`
	TotalRevenue CashflowReportMonthlySummaryTotalRevenue `json:"total_revenue"`
	TotalLoanPayment CashflowReportMonthlySummaryTotalLoanPayment `json:"total_loan_payment"`
	TotalVariableExpense CashflowReportMonthlySummaryTotalVariableExpense `json:"total_variable_expense"`
	TotalPayroll CashflowReportMonthlySummaryTotalPayroll `json:"total_payroll"`
	// The total number of all NSF transactions during this month.
	NsfTransactionCount int32 `json:"nsf_transaction_count"`
	// The total number of all overdraft transactions during this month.
	OverdraftTransactionCount int32 `json:"overdraft_transaction_count"`
	// The number of days with a negative daily average ending balance. The daily average is calculated across all valid accounts. Values will be in the range [0, 31].
	NegativeEndingBalanceDayCount int32 `json:"negative_ending_balance_day_count"`
	AdditionalProperties map[string]interface{}
}

type _CashflowReportMonthlySummary CashflowReportMonthlySummary

// NewCashflowReportMonthlySummary instantiates a new CashflowReportMonthlySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashflowReportMonthlySummary(startDate string, endDate string, startingBalance NullableCashflowReportMonthlySummaryStartingBalance, endingBalance NullableCashflowReportMonthlySummaryEndingBalance, averageDailyEndingBalance CashflowReportMonthlySummaryAverageDailyEndingBalance, averageDailyInflowAmount CashflowReportMonthlySummaryAverageDailyInflowAmount, averageDailyOutflowAmount CashflowReportMonthlySummaryAverageDailyOutflowAmount, averageDailyNetCashflowAmount CashflowReportMonthlySummaryAverageDailyNetCashflowAmount, averageDailyInflowTransactionCount float32, averageDailyOutflowTransactionCount float32, totalRevenue CashflowReportMonthlySummaryTotalRevenue, totalLoanPayment CashflowReportMonthlySummaryTotalLoanPayment, totalVariableExpense CashflowReportMonthlySummaryTotalVariableExpense, totalPayroll CashflowReportMonthlySummaryTotalPayroll, nsfTransactionCount int32, overdraftTransactionCount int32, negativeEndingBalanceDayCount int32) *CashflowReportMonthlySummary {
	this := CashflowReportMonthlySummary{}
	this.StartDate = startDate
	this.EndDate = endDate
	this.StartingBalance = startingBalance
	this.EndingBalance = endingBalance
	this.AverageDailyEndingBalance = averageDailyEndingBalance
	this.AverageDailyInflowAmount = averageDailyInflowAmount
	this.AverageDailyOutflowAmount = averageDailyOutflowAmount
	this.AverageDailyNetCashflowAmount = averageDailyNetCashflowAmount
	this.AverageDailyInflowTransactionCount = averageDailyInflowTransactionCount
	this.AverageDailyOutflowTransactionCount = averageDailyOutflowTransactionCount
	this.TotalRevenue = totalRevenue
	this.TotalLoanPayment = totalLoanPayment
	this.TotalVariableExpense = totalVariableExpense
	this.TotalPayroll = totalPayroll
	this.NsfTransactionCount = nsfTransactionCount
	this.OverdraftTransactionCount = overdraftTransactionCount
	this.NegativeEndingBalanceDayCount = negativeEndingBalanceDayCount
	return &this
}

// NewCashflowReportMonthlySummaryWithDefaults instantiates a new CashflowReportMonthlySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashflowReportMonthlySummaryWithDefaults() *CashflowReportMonthlySummary {
	this := CashflowReportMonthlySummary{}
	return &this
}

// GetStartDate returns the StartDate field value
func (o *CashflowReportMonthlySummary) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *CashflowReportMonthlySummary) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *CashflowReportMonthlySummary) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *CashflowReportMonthlySummary) SetEndDate(v string) {
	o.EndDate = v
}

// GetStartingBalance returns the StartingBalance field value
// If the value is explicit nil, the zero value for CashflowReportMonthlySummaryStartingBalance will be returned
func (o *CashflowReportMonthlySummary) GetStartingBalance() CashflowReportMonthlySummaryStartingBalance {
	if o == nil || o.StartingBalance.Get() == nil {
		var ret CashflowReportMonthlySummaryStartingBalance
		return ret
	}

	return *o.StartingBalance.Get()
}

// GetStartingBalanceOk returns a tuple with the StartingBalance field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashflowReportMonthlySummary) GetStartingBalanceOk() (*CashflowReportMonthlySummaryStartingBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartingBalance.Get(), o.StartingBalance.IsSet()
}

// SetStartingBalance sets field value
func (o *CashflowReportMonthlySummary) SetStartingBalance(v CashflowReportMonthlySummaryStartingBalance) {
	o.StartingBalance.Set(&v)
}

// GetEndingBalance returns the EndingBalance field value
// If the value is explicit nil, the zero value for CashflowReportMonthlySummaryEndingBalance will be returned
func (o *CashflowReportMonthlySummary) GetEndingBalance() CashflowReportMonthlySummaryEndingBalance {
	if o == nil || o.EndingBalance.Get() == nil {
		var ret CashflowReportMonthlySummaryEndingBalance
		return ret
	}

	return *o.EndingBalance.Get()
}

// GetEndingBalanceOk returns a tuple with the EndingBalance field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CashflowReportMonthlySummary) GetEndingBalanceOk() (*CashflowReportMonthlySummaryEndingBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndingBalance.Get(), o.EndingBalance.IsSet()
}

// SetEndingBalance sets field value
func (o *CashflowReportMonthlySummary) SetEndingBalance(v CashflowReportMonthlySummaryEndingBalance) {
	o.EndingBalance.Set(&v)
}

// GetAverageDailyEndingBalance returns the AverageDailyEndingBalance field value
func (o *CashflowReportMonthlySummary) GetAverageDailyEndingBalance() CashflowReportMonthlySummaryAverageDailyEndingBalance {
	if o == nil {
		var ret CashflowReportMonthlySummaryAverageDailyEndingBalance
		return ret
	}

	return o.AverageDailyEndingBalance
}

// GetAverageDailyEndingBalanceOk returns a tuple with the AverageDailyEndingBalance field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetAverageDailyEndingBalanceOk() (*CashflowReportMonthlySummaryAverageDailyEndingBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AverageDailyEndingBalance, true
}

// SetAverageDailyEndingBalance sets field value
func (o *CashflowReportMonthlySummary) SetAverageDailyEndingBalance(v CashflowReportMonthlySummaryAverageDailyEndingBalance) {
	o.AverageDailyEndingBalance = v
}

// GetAverageDailyInflowAmount returns the AverageDailyInflowAmount field value
func (o *CashflowReportMonthlySummary) GetAverageDailyInflowAmount() CashflowReportMonthlySummaryAverageDailyInflowAmount {
	if o == nil {
		var ret CashflowReportMonthlySummaryAverageDailyInflowAmount
		return ret
	}

	return o.AverageDailyInflowAmount
}

// GetAverageDailyInflowAmountOk returns a tuple with the AverageDailyInflowAmount field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetAverageDailyInflowAmountOk() (*CashflowReportMonthlySummaryAverageDailyInflowAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AverageDailyInflowAmount, true
}

// SetAverageDailyInflowAmount sets field value
func (o *CashflowReportMonthlySummary) SetAverageDailyInflowAmount(v CashflowReportMonthlySummaryAverageDailyInflowAmount) {
	o.AverageDailyInflowAmount = v
}

// GetAverageDailyOutflowAmount returns the AverageDailyOutflowAmount field value
func (o *CashflowReportMonthlySummary) GetAverageDailyOutflowAmount() CashflowReportMonthlySummaryAverageDailyOutflowAmount {
	if o == nil {
		var ret CashflowReportMonthlySummaryAverageDailyOutflowAmount
		return ret
	}

	return o.AverageDailyOutflowAmount
}

// GetAverageDailyOutflowAmountOk returns a tuple with the AverageDailyOutflowAmount field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetAverageDailyOutflowAmountOk() (*CashflowReportMonthlySummaryAverageDailyOutflowAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AverageDailyOutflowAmount, true
}

// SetAverageDailyOutflowAmount sets field value
func (o *CashflowReportMonthlySummary) SetAverageDailyOutflowAmount(v CashflowReportMonthlySummaryAverageDailyOutflowAmount) {
	o.AverageDailyOutflowAmount = v
}

// GetAverageDailyNetCashflowAmount returns the AverageDailyNetCashflowAmount field value
func (o *CashflowReportMonthlySummary) GetAverageDailyNetCashflowAmount() CashflowReportMonthlySummaryAverageDailyNetCashflowAmount {
	if o == nil {
		var ret CashflowReportMonthlySummaryAverageDailyNetCashflowAmount
		return ret
	}

	return o.AverageDailyNetCashflowAmount
}

// GetAverageDailyNetCashflowAmountOk returns a tuple with the AverageDailyNetCashflowAmount field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetAverageDailyNetCashflowAmountOk() (*CashflowReportMonthlySummaryAverageDailyNetCashflowAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AverageDailyNetCashflowAmount, true
}

// SetAverageDailyNetCashflowAmount sets field value
func (o *CashflowReportMonthlySummary) SetAverageDailyNetCashflowAmount(v CashflowReportMonthlySummaryAverageDailyNetCashflowAmount) {
	o.AverageDailyNetCashflowAmount = v
}

// GetAverageDailyInflowTransactionCount returns the AverageDailyInflowTransactionCount field value
func (o *CashflowReportMonthlySummary) GetAverageDailyInflowTransactionCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AverageDailyInflowTransactionCount
}

// GetAverageDailyInflowTransactionCountOk returns a tuple with the AverageDailyInflowTransactionCount field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetAverageDailyInflowTransactionCountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AverageDailyInflowTransactionCount, true
}

// SetAverageDailyInflowTransactionCount sets field value
func (o *CashflowReportMonthlySummary) SetAverageDailyInflowTransactionCount(v float32) {
	o.AverageDailyInflowTransactionCount = v
}

// GetAverageDailyOutflowTransactionCount returns the AverageDailyOutflowTransactionCount field value
func (o *CashflowReportMonthlySummary) GetAverageDailyOutflowTransactionCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AverageDailyOutflowTransactionCount
}

// GetAverageDailyOutflowTransactionCountOk returns a tuple with the AverageDailyOutflowTransactionCount field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetAverageDailyOutflowTransactionCountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AverageDailyOutflowTransactionCount, true
}

// SetAverageDailyOutflowTransactionCount sets field value
func (o *CashflowReportMonthlySummary) SetAverageDailyOutflowTransactionCount(v float32) {
	o.AverageDailyOutflowTransactionCount = v
}

// GetTotalRevenue returns the TotalRevenue field value
func (o *CashflowReportMonthlySummary) GetTotalRevenue() CashflowReportMonthlySummaryTotalRevenue {
	if o == nil {
		var ret CashflowReportMonthlySummaryTotalRevenue
		return ret
	}

	return o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetTotalRevenueOk() (*CashflowReportMonthlySummaryTotalRevenue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalRevenue, true
}

// SetTotalRevenue sets field value
func (o *CashflowReportMonthlySummary) SetTotalRevenue(v CashflowReportMonthlySummaryTotalRevenue) {
	o.TotalRevenue = v
}

// GetTotalLoanPayment returns the TotalLoanPayment field value
func (o *CashflowReportMonthlySummary) GetTotalLoanPayment() CashflowReportMonthlySummaryTotalLoanPayment {
	if o == nil {
		var ret CashflowReportMonthlySummaryTotalLoanPayment
		return ret
	}

	return o.TotalLoanPayment
}

// GetTotalLoanPaymentOk returns a tuple with the TotalLoanPayment field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetTotalLoanPaymentOk() (*CashflowReportMonthlySummaryTotalLoanPayment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalLoanPayment, true
}

// SetTotalLoanPayment sets field value
func (o *CashflowReportMonthlySummary) SetTotalLoanPayment(v CashflowReportMonthlySummaryTotalLoanPayment) {
	o.TotalLoanPayment = v
}

// GetTotalVariableExpense returns the TotalVariableExpense field value
func (o *CashflowReportMonthlySummary) GetTotalVariableExpense() CashflowReportMonthlySummaryTotalVariableExpense {
	if o == nil {
		var ret CashflowReportMonthlySummaryTotalVariableExpense
		return ret
	}

	return o.TotalVariableExpense
}

// GetTotalVariableExpenseOk returns a tuple with the TotalVariableExpense field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetTotalVariableExpenseOk() (*CashflowReportMonthlySummaryTotalVariableExpense, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalVariableExpense, true
}

// SetTotalVariableExpense sets field value
func (o *CashflowReportMonthlySummary) SetTotalVariableExpense(v CashflowReportMonthlySummaryTotalVariableExpense) {
	o.TotalVariableExpense = v
}

// GetTotalPayroll returns the TotalPayroll field value
func (o *CashflowReportMonthlySummary) GetTotalPayroll() CashflowReportMonthlySummaryTotalPayroll {
	if o == nil {
		var ret CashflowReportMonthlySummaryTotalPayroll
		return ret
	}

	return o.TotalPayroll
}

// GetTotalPayrollOk returns a tuple with the TotalPayroll field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetTotalPayrollOk() (*CashflowReportMonthlySummaryTotalPayroll, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalPayroll, true
}

// SetTotalPayroll sets field value
func (o *CashflowReportMonthlySummary) SetTotalPayroll(v CashflowReportMonthlySummaryTotalPayroll) {
	o.TotalPayroll = v
}

// GetNsfTransactionCount returns the NsfTransactionCount field value
func (o *CashflowReportMonthlySummary) GetNsfTransactionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NsfTransactionCount
}

// GetNsfTransactionCountOk returns a tuple with the NsfTransactionCount field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetNsfTransactionCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NsfTransactionCount, true
}

// SetNsfTransactionCount sets field value
func (o *CashflowReportMonthlySummary) SetNsfTransactionCount(v int32) {
	o.NsfTransactionCount = v
}

// GetOverdraftTransactionCount returns the OverdraftTransactionCount field value
func (o *CashflowReportMonthlySummary) GetOverdraftTransactionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OverdraftTransactionCount
}

// GetOverdraftTransactionCountOk returns a tuple with the OverdraftTransactionCount field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetOverdraftTransactionCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OverdraftTransactionCount, true
}

// SetOverdraftTransactionCount sets field value
func (o *CashflowReportMonthlySummary) SetOverdraftTransactionCount(v int32) {
	o.OverdraftTransactionCount = v
}

// GetNegativeEndingBalanceDayCount returns the NegativeEndingBalanceDayCount field value
func (o *CashflowReportMonthlySummary) GetNegativeEndingBalanceDayCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NegativeEndingBalanceDayCount
}

// GetNegativeEndingBalanceDayCountOk returns a tuple with the NegativeEndingBalanceDayCount field value
// and a boolean to check if the value has been set.
func (o *CashflowReportMonthlySummary) GetNegativeEndingBalanceDayCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NegativeEndingBalanceDayCount, true
}

// SetNegativeEndingBalanceDayCount sets field value
func (o *CashflowReportMonthlySummary) SetNegativeEndingBalanceDayCount(v int32) {
	o.NegativeEndingBalanceDayCount = v
}

func (o CashflowReportMonthlySummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if true {
		toSerialize["end_date"] = o.EndDate
	}
	if true {
		toSerialize["starting_balance"] = o.StartingBalance.Get()
	}
	if true {
		toSerialize["ending_balance"] = o.EndingBalance.Get()
	}
	if true {
		toSerialize["average_daily_ending_balance"] = o.AverageDailyEndingBalance
	}
	if true {
		toSerialize["average_daily_inflow_amount"] = o.AverageDailyInflowAmount
	}
	if true {
		toSerialize["average_daily_outflow_amount"] = o.AverageDailyOutflowAmount
	}
	if true {
		toSerialize["average_daily_net_cashflow_amount"] = o.AverageDailyNetCashflowAmount
	}
	if true {
		toSerialize["average_daily_inflow_transaction_count"] = o.AverageDailyInflowTransactionCount
	}
	if true {
		toSerialize["average_daily_outflow_transaction_count"] = o.AverageDailyOutflowTransactionCount
	}
	if true {
		toSerialize["total_revenue"] = o.TotalRevenue
	}
	if true {
		toSerialize["total_loan_payment"] = o.TotalLoanPayment
	}
	if true {
		toSerialize["total_variable_expense"] = o.TotalVariableExpense
	}
	if true {
		toSerialize["total_payroll"] = o.TotalPayroll
	}
	if true {
		toSerialize["nsf_transaction_count"] = o.NsfTransactionCount
	}
	if true {
		toSerialize["overdraft_transaction_count"] = o.OverdraftTransactionCount
	}
	if true {
		toSerialize["negative_ending_balance_day_count"] = o.NegativeEndingBalanceDayCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CashflowReportMonthlySummary) UnmarshalJSON(bytes []byte) (err error) {
	varCashflowReportMonthlySummary := _CashflowReportMonthlySummary{}

	if err = json.Unmarshal(bytes, &varCashflowReportMonthlySummary); err == nil {
		*o = CashflowReportMonthlySummary(varCashflowReportMonthlySummary)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "starting_balance")
		delete(additionalProperties, "ending_balance")
		delete(additionalProperties, "average_daily_ending_balance")
		delete(additionalProperties, "average_daily_inflow_amount")
		delete(additionalProperties, "average_daily_outflow_amount")
		delete(additionalProperties, "average_daily_net_cashflow_amount")
		delete(additionalProperties, "average_daily_inflow_transaction_count")
		delete(additionalProperties, "average_daily_outflow_transaction_count")
		delete(additionalProperties, "total_revenue")
		delete(additionalProperties, "total_loan_payment")
		delete(additionalProperties, "total_variable_expense")
		delete(additionalProperties, "total_payroll")
		delete(additionalProperties, "nsf_transaction_count")
		delete(additionalProperties, "overdraft_transaction_count")
		delete(additionalProperties, "negative_ending_balance_day_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCashflowReportMonthlySummary struct {
	value *CashflowReportMonthlySummary
	isSet bool
}

func (v NullableCashflowReportMonthlySummary) Get() *CashflowReportMonthlySummary {
	return v.value
}

func (v *NullableCashflowReportMonthlySummary) Set(val *CashflowReportMonthlySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCashflowReportMonthlySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCashflowReportMonthlySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashflowReportMonthlySummary(val *CashflowReportMonthlySummary) *NullableCashflowReportMonthlySummary {
	return &NullableCashflowReportMonthlySummary{value: val, isSet: true}
}

func (v NullableCashflowReportMonthlySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashflowReportMonthlySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


