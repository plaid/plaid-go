package plaid

import (
	"encoding/json"
	"errors"
)

// StudentLoanLiability contains student loan liability data.
type StudentLoanLiability struct {
	AccountID                  string                     `json:"account_id"`
	AccountNumber              string                     `json:"account_number"`
	DisbursementDates          []string                   `json:"disbursement_dates"`
	ExpectedPayoffDate         string                     `json:"expected_payoff_date"`
	Guarantor                  string                     `json:"guarantor"`
	InterestRatePercentage     float64                    `json:"interest_rate_percentage"`
	IsOverdue                  bool                       `json:"is_overdue"`
	LastPaymentAmount          float64                    `json:"last_payment_amount"`
	LastPaymentDate            string                     `json:"last_payment_date"`
	LastStatementBalance       float64                    `json:"last_statement_balance"`
	LastStatementIssueDate     string                     `json:"last_statement_issue_date"`
	LoanName                   string                     `json:"loan_name"`
	LoanStatus                 StudentLoanStatus          `json:"loan_status"`
	MinimumPaymentAmount       float64                    `json:"minimum_payment_amount"`
	NextPaymentDueDate         string                     `json:"next_payment_due_date"`
	OriginationDate            string                     `json:"origination_date"`
	OriginationPrincipalAmount float64                    `json:"origination_principal_amount"`
	OutstandingInterestAmount  float64                    `json:"outstanding_interest_amount"`
	PaymentReferenceNumber     string                     `json:"payment_reference_number"`
	PSLFStatus                 PSLFStatus                 `json:"pslf_status"`
	RepaymentPlan              StudentLoanRepaymentPlan   `json:"repayment_plan"`
	SequenceNumber             string                     `json:"sequence_number"`
	ServicerAddress            StudentLoanServicerAddress `json:"servicer_address"`
	YTDInterestPaid            float64                    `json:"ytd_interest_paid"`
	YTDPrincipalPaid           float64                    `json:"ytd_principal_paid"`
}

// PSLFStatus contains information about the student's eligibility in the
// Public Service Loan Forgiveness program.
type PSLFStatus struct {
	EstimatedEligibilityDate string `json:"estimated_eligibility_date"`
	PaymentsMade             int64  `json:"payments_made"`
	PaymentsRemaining        int64  `json:"payments_remaining"`
}

// StudentLoanServicerAddress is the address of the servicer.
type StudentLoanServicerAddress struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
	Region     string `json:"region"`
	Street     string `json:"street"`
}

// StudentLoanStatus contains details about the status of the student loan.
type StudentLoanStatus struct {
	Type    string `json:"type"`
	EndDate string `json:"end_date"`
}

// StudentLoanRepaymentPlan contains details about the repayment plan of the
// loan.
type StudentLoanRepaymentPlan struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type getLiabilitiesRequestOptions struct {
	AccountIDs []string `json:"account_ids,omitempty"`
}

type getLiabilitiesRequest struct {
	ClientID    string                       `json:"client_id"`
	Secret      string                       `json:"secret"`
	AccessToken string                       `json:"access_token"`
	Options     getLiabilitiesRequestOptions `json:"options,omitempty"`
}

// GetLiabilitiesResponse is the response from /liabilities/get.
type GetLiabilitiesResponse struct {
	APIResponse
	Accounts    []Account `json:"accounts"`
	Item        Item      `json:"item"`
	Liabilities struct {
		Student []StudentLoanLiability `json:"student"`
	} `json:"liabilities"`
}

// GetLiabilitiesOptions contains options for /liabilities/get.
type GetLiabilitiesOptions struct {
	// AccountIDs is used to filter accounts included in the response. A nil or
	// zero-length slice does not result in any filter being applied.
	AccountIDs []string
}

// GetLiabilitiesWithOptions retrieves liability data. See
// https://plaid.com/docs/api/#liabilities.
func (c *Client) GetLiabilitiesWithOptions(
	accessToken string,
	options GetLiabilitiesOptions,
) (resp GetLiabilitiesResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/liabilities/get - access token must be specified")
	}

	req := getLiabilitiesRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
		Options:     getLiabilitiesRequestOptions{},
	}
	if len(options.AccountIDs) > 0 {
		req.Options.AccountIDs = options.AccountIDs
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	err = c.Call("/liabilities/get", jsonBody, &resp)
	return resp, err
}

// GetLiabilities retrieves liability data. See
// https://plaid.com/docs/api/#liabilities.
func (c *Client) GetLiabilities(accessToken string) (resp GetLiabilitiesResponse, err error) {
	return c.GetLiabilitiesWithOptions(accessToken, GetLiabilitiesOptions{
		AccountIDs: []string{},
	})
}
