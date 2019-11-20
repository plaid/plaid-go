package plaid

import (
	"encoding/json"
	"errors"
)

type Income struct {
	IncomeStreams                       []IncomeStream `json:"income_streams"`
	LastYearIncome                      int            `json:"last_year_income"`
	LastYearIncomeBeforeTax             int            `json:"last_year_income_before_tax"`
	ProjectedYearlyIncome               int            `json:"projected_yearly_income"`
	ProjectedYearlyIncomeBeforeTax      int            `json:"projected_yearly_income_before_tax"`
	MaxNumberOfOverlappingIncomeStreams int            `json:"max_number_of_overlapping_income_streams"`
	NumberOfIncomeStreams               int            `json:"number_of_income_streams"`
}

type IncomeStream struct {
	Confidence    float64 `json:"confidence"`
	Days          int     `json:"days"`
	MonthlyIncome int     `json:"monthly_income"`
	Name          string  `json:"name"`
}

type getIncomeRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

type GetIncomeResponse struct {
	APIResponse
	Income Income `json:"income"`
}

// GetIncome retrieves information pertaining to an Item's income.
// See https://plaid.com/docs/api/#income.
func (c *Client) GetIncome(accessToken string) (resp GetIncomeResponse, err error) {
	if accessToken == "" {
		return resp, errors.New("/income/get - access token must be specified")
	}

	jsonBody, err := json.Marshal(getIncomeRequest{
		ClientID:    c.clientID,
		Secret:      c.secret,
		AccessToken: accessToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/income/get", jsonBody, &resp)
	return resp, err
}
