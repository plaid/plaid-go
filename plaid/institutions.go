package plaid

import (
	"encoding/json"
	"errors"
	"time"
)

type Institution struct {
	Credentials    []Credential `json:"credentials"`
	HasMFA         bool         `json:"has_mfa"`
	ID             string       `json:"institution_id"`
	MFA            []string     `json:"mfa"`
	Name           string       `json:"name"`
	Products       []string     `json:"products"`
	CountryCodes   []string     `json:"country_codes"`
	OAuth          bool         `json:"oauth"`
	RoutingNumbers []string     `json:"routing_numbers"`

	// Included when `options.include_status` is true.
	InstitutionStatus *InstitutionStatus `json:"status,omitempty"`

	// Included when `options.include_optional_metadata` is true.
	PrimaryColor string `json:"primary_color,omitempty"`
	// Included when `options.include_optional_metadata` is true.
	URL string `json:"url,omitempty"`
	// Included when `options.include_optional_metadata` is true.
	Logo string `json:"logo,omitempty"`
}

type InstitutionStatus struct {
	ItemLogins          ItemLogins              `json:"item_logins"`
	TransactionsUpdates InstitutionStatusFields `json:"transactions_updates"`
	Auth                InstitutionStatusFields `json:"auth"`
	Balance             InstitutionStatusFields `json:"balance"`
	Identity            InstitutionStatusFields `json:"identity"`
}

type ItemLogins struct {
	Status           string                     `json:"status"`
	LastStatusChange time.Time                  `json:"last_status_change"`
	Breakdown        InstitutionStatusBreakdown `json:"breakdown"`
}

type InstitutionStatusFields struct {
	Status           string                     `json:"status"`
	LastStatusChange time.Time                  `json:"last_status_change"`
	Breakdown        InstitutionStatusBreakdown `json:"breakdown"`
}

type InstitutionStatusBreakdown struct {
	Success          float64 `json:"success"`
	ErrorPlaid       float64 `json:"error_plaid"`
	ErrorInstitution float64 `json:"error_institution"`
	RefreshInterval  string  `json:"refresh_interval"` // only applicable to TransactionsUpdates status
}

type Credential struct {
	Label string `json:"label"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

type getInstitutionsRequest struct {
	ClientID     string                 `json:"client_id"`
	Secret       string                 `json:"secret"`
	Count        int                    `json:"count"`
	Offset       int                    `json:"offset"`
	CountryCodes []string               `json:"country_codes"`
	Options      GetInstitutionsOptions `json:"options,omitempty"`
}

type GetInstitutionsOptions struct {
	Products                []string `json:"products"`
	IncludeOptionalMetadata bool     `json:"include_optional_metadata"`
	OAuth                   *bool    `json:"oauth"`
	RoutingNumbers          []string `json:"routing_numbers"`
}

type GetInstitutionsResponse struct {
	APIResponse
	Institutions []Institution `json:"institutions"`
	Total        int           `json:"total"`
}

type getInstitutionByIDRequest struct {
	ID           string                    `json:"institution_id"`
	CountryCodes []string                  `json:"country_codes"`
	ClientID     string                    `json:"client_id"`
	Secret       string                    `json:"secret"`
	Options      GetInstitutionByIDOptions `json:"options,omitempty"`
}

type GetInstitutionByIDOptions struct {
	IncludeOptionalMetadata bool `json:"include_optional_metadata"`
	IncludeStatus           bool `json:"include_status"`
}

type GetInstitutionByIDResponse struct {
	APIResponse
	Institution Institution `json:"institution"`
}

type searchInstitutionsRequest struct {
	Query        string                    `json:"query"`
	CountryCodes []string                  `json:"country_codes"`
	Products     []string                  `json:"products"`
	ClientID     string                    `json:"client_id"`
	Secret       string                    `json:"secret"`
	Options      SearchInstitutionsOptions `json:"options,omitempty"`
}

type SearchInstitutionsOptions struct {
	IncludeOptionalMetadata bool                   `json:"include_optional_metadata"`
	AccountFilter           map[string]interface{} `json:"account_filter"`
	OAuth                   *bool                  `json:"oauth"`
}

type SearchInstitutionsResponse struct {
	APIResponse
	Institutions []Institution `json:"institutions"`
}

// GetInstitutionByID returns information for a single institution given an ID.
// See https://plaid.com/docs/api/institutions/#institutionsget_by_id.
func (c *Client) GetInstitutionByID(
	id string,
	countryCodes []string,
) (resp GetInstitutionByIDResponse, err error) {
	return c.GetInstitutionByIDWithOptions(id, countryCodes, GetInstitutionByIDOptions{})
}

// GetInstitutionByIDWithOptions returns information for a single institution given an ID.
// See https://plaid.com/docs/api/institutions/#institutionsget_by_id.
func (c *Client) GetInstitutionByIDWithOptions(
	id string,
	countryCodes []string,
	options GetInstitutionByIDOptions,
) (resp GetInstitutionByIDResponse, err error) {
	if id == "" {
		return resp, errors.New("/institutions/get_by_id - institution id must be specified")
	}

	jsonBody, err := json.Marshal(getInstitutionByIDRequest{
		ID:           id,
		CountryCodes: countryCodes,
		ClientID:     c.clientID,
		Secret:       c.secret,
		Options:      options,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/institutions/get_by_id", jsonBody, &resp)
	return resp, err
}

// GetInstitutions returns information for all institutions supported by Plaid.
// See https://plaid.com/docs/api/institutions/#institutionsget.
func (c *Client) GetInstitutions(count, offset int, countryCodes []string) (resp GetInstitutionsResponse, err error) {
	return c.GetInstitutionsWithOptions(count, offset, countryCodes, GetInstitutionsOptions{})
}

// GetInstitutionsWithOptions returns information for all institutions supported by Plaid.
// See https://plaid.com/docs/api/institutions/#institutionsget.
func (c *Client) GetInstitutionsWithOptions(
	count int,
	offset int,
	countryCodes []string,
	options GetInstitutionsOptions,
) (resp GetInstitutionsResponse, err error) {
	if count == 0 {
		count = 50
	}

	jsonBody, err := json.Marshal(getInstitutionsRequest{
		ClientID:     c.clientID,
		Secret:       c.secret,
		Count:        count,
		Offset:       offset,
		CountryCodes: countryCodes,
		Options:      options,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/institutions/get", jsonBody, &resp)
	return resp, err
}

// SearchInstitutions returns institutions corresponding to a query string and
// supported products.
// See https://plaid.com/docs/api/institutions/#institutionssearch.
func (c *Client) SearchInstitutions(
	query string,
	products []string,
	countryCodes []string,
) (resp SearchInstitutionsResponse, err error) {
	return c.SearchInstitutionsWithOptions(query, products, countryCodes, SearchInstitutionsOptions{})
}

// SearchInstitutionsWithOptions returns institutions corresponding to a query string and
// supported products.
// See https://plaid.com/docs/api/institutions/#institutionssearch.
func (c *Client) SearchInstitutionsWithOptions(
	query string,
	products []string,
	countryCodes []string,
	options SearchInstitutionsOptions,
) (resp SearchInstitutionsResponse, err error) {
	if query == "" {
		return resp, errors.New("/institutions/search - query must be specified")
	}

	jsonBody, err := json.Marshal(searchInstitutionsRequest{
		Query:        query,
		Products:     products,
		CountryCodes: countryCodes,
		ClientID:     c.clientID,
		Secret:       c.secret,
		Options:      options,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/institutions/search", jsonBody, &resp)
	return resp, err
}
