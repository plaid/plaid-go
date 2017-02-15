package plaid

import (
	"errors"
	"net/url"
)

// GetInstitution returns information for a single institution given an ID.
func GetInstitution(environment environmentURL, id string) (inst institution, err error) {
	err = getAndUnmarshal(environment, "/institutions/all/"+id, &inst)
	return
}

// GetInstitutionsSearch returns all institutions that match the query parameters.
// If product parameter is included, results are filtered by product.
// If institution id option is specified, query and product parameters are ignored.
func GetInstitutionsSearch(environment environmentURL, query, product, id string) (institutions []institutionJson, err error) {
	if query == "" && id == "" {
		return nil, errors.New("Query or institution id must be specified")
	}

	v := url.Values{}

	if query != "" {
		v.Add("q", query)
	}

	if product != "" {
		v.Add("p", product)
	}

	if id != "" {
		v.Add("id", id)
	}

	payload := v.Encode()

	err = getAndUnmarshal(environment, "/institutions/all/search?"+payload, &institutions)
	return
}

type institutionJson struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Products struct {
		Auth    bool `json:"auth"`
		Balance bool `json:"balance"`
		Connect bool `json:"connect"`
		Info    bool `json:"info"`
	} `json:"products"`
	ForgottenPassword string `json:"forgottenPassword"`
	AccountLocked     string `json:"accountLocked"`
	AccountSetup      string `json:"accountSetup"`
	Video             string `json:"video"`
	Fields            []struct {
		Name  string `json:"name"`
		Label string `json:"label"`
		Type  string `json:"type"`
	} `json:"fields"`
	Colors struct {
		Light   string `json:"light"`
		Primary string `json:"primary"`
		Dark    string `json:"dark"`
		Darker  string `json:"darker"`
	} `json:"colors"`
	Logo      string `json:"logo"`
	Namebreak string `json:"nameBreak"`
	Type      string `json:"type"`
}

type institution struct {
	Credentials struct {
		Password string `json:"password"` // e.g.: "Password"
		PIN      string `json:"pin"`      // e.g.: "PIN"
		Username string `json:"username"` // e.g.: "Online ID"
	}
	Name     string   `json:"name"`     // e.g.: "Bank of America"
	HasMFA   bool     `json:"has_mfa"`  // e.g.: true
	ID       string   `json:"id"`       // e.g.: "5301a93ac140de84910000e0"
	MFA      []string `json:"mfa"`      // e.g.: ["code", "list", "questions"]
	Products []string `json:"products"` // e.g.: ["connect", "auth", "balance"]
	Type     string   `json:"type"`     // e.g.: "bofa"
}
