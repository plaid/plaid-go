package plaid

import (
	"encoding/json"
	"errors"
)

type Institution struct {
	Credentials []Credential `json:"credentials"`
	HasMFA      bool         `json:"has_mfa"`
	ID          string       `json:"institution_id"`
	MFA         []string     `json:"mfa"`
	Name        string       `json:"name"`
	Products    []string     `json:"products"`

	// included when options.include_status is true
	InstitutionStatus InstitutionStatus `json:"status"`

	// included when options.include_optional_metadata is true
	PrimaryColor string `json:"primary_color"`
	URL          string `json:"url"`
	Logo         string `json:"logo"`
}

type InstitutionStatus struct {
	ItemLogins ItemLogins `json:"item_logins"`
}

type ItemLogins struct {
	Status           string                     `json:"status"`
	LastStatusChange string                     `json:"last_status_change"`
	Breakdown        InstitutionStatusBreakdown `json:"breakdown"`
}

type InstitutionStatusBreakdown struct {
	Success          float64 `json:"success"`
	ErrorPlaid       float64 `json:"error_plaid"`
	ErrorInstitution float64 `json:"error_institution"`
}

type Credential struct {
	Label string `json:"label"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

type getInstitutionsRequest struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	Count    int    `json:"count"`
	Offset   int    `json:"offset"`
}

type GetInstitutionsResponse struct {
	APIResponse
	Institutions []Institution `json:"institutions"`
	Total        int           `json:"total"`
}

type getInstitutionByIDRequest struct {
	ID        string                           `json:"institution_id"`
	PublicKey string                           `json:"public_key"`
	Options   getInstitutionByIDRequestOptions `json:"options,omitempty"`
}

type getInstitutionByIDRequestOptions struct {
	IncludeOptionalMetadata bool `json:"include_optional_metadata"`
	IncludeStatus           bool `json:"include_status"`
}

type GetInstitutionByIDResponse struct {
	APIResponse
	Institution Institution `json:"institution"`
}

type GetInstitutionByIDOptions struct {
	IncludeOptionalMetadata bool `json:"include_optional_metadata"`
	IncludeStatus           bool `json:"include_status"`
}

type searchInstitutionsRequest struct {
	Query     string   `json:"query"`
	Products  []string `json:"products"`
	PublicKey string   `json:"public_key"`
}

type SearchInstitutionsResponse struct {
	APIResponse
	Institutions []Institution `json:"institutions"`
}

// GetInstitutionByID returns information for a single institution given an ID.
// See https://plaid.com/docs/api/#institutions-by-id.
func (c *Client) GetInstitutionByID(id string) (resp GetInstitutionByIDResponse, err error) {
	if id == "" {
		return resp, errors.New("/institutions/get_by_id - institution id must be specified")
	}

	jsonBody, err := json.Marshal(getInstitutionByIDRequest{
		ID:        id,
		PublicKey: c.publicKey,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/institutions/get_by_id", jsonBody, &resp)
	return resp, err
}

// GetInstitutionByIDWithOptions returns information for a single institution
// given an ID.
// See https://plaid.com/docs/api/#institutions-by-id.
func (c *Client) GetInstitutionByIDWithOptions(id string, options GetInstitutionByIDOptions) (resp GetInstitutionByIDResponse, err error) {
	if id == "" {
		return resp, errors.New("/institutions/get_by_id - institution id must be specified")
	}

	jsonBody, err := json.Marshal(getInstitutionByIDRequest{
		ID:        id,
		PublicKey: c.publicKey,
		Options: getInstitutionByIDRequestOptions{
			IncludeOptionalMetadata: options.IncludeOptionalMetadata,
			IncludeStatus:           options.IncludeStatus,
		},
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/institutions/get_by_id", jsonBody, &resp)
	return resp, err
}

// GetInstitutions returns information for all institutions supported by Plaid.
// See https://plaid.com/docs/api/#all-institutions.
func (c *Client) GetInstitutions(count, offset int) (resp GetInstitutionsResponse, err error) {
	if count == 0 {
		count = 50
	}

	jsonBody, err := json.Marshal(getInstitutionsRequest{
		ClientID: c.clientID,
		Secret:   c.secret,
		Count:    count,
		Offset:   offset,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/institutions/get", jsonBody, &resp)
	return resp, err
}

// SearchInstitutions returns institutions corresponding to a query string and
// supported products.
// See https://plaid.com/docs/api/#institution-search.
func (c *Client) SearchInstitutions(query string, products []string) (resp SearchInstitutionsResponse, err error) {
	if query == "" {
		return resp, errors.New("/institutions/search - query must be specified")
	}

	jsonBody, err := json.Marshal(searchInstitutionsRequest{
		Query:     query,
		Products:  products,
		PublicKey: c.publicKey,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/institutions/search", jsonBody, &resp)
	return resp, err
}
