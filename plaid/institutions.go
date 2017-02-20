package plaid

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetInstitution returns information for a single institution given an ID.
// See: https://plaid.com/docs/api/#institutions-by-id
func GetInstitution(environment environmentURL, id string) (inst institution, err error) {
	if id == "" {
		return inst, errors.New("/institutions/all/:id - institution id must be specified")
	}

	err = getAndUnmarshal(environment, "/institutions/all/"+id, &inst)
	return inst, err
}

// GetInstitutionsSearch returns all institutions that match the query parameters.
// If product parameter is included, results are filtered by product.
// If institution id option is specified, query and product parameters are ignored.
// See: https://plaid.com/docs/api/#institution-search
func GetInstitutionsSearch(environment environmentURL, query, product, id string) (institutions []institutionExtended, err error) {
	if query == "" && id == "" {
		return nil, errors.New("/institutions/all/ - query or institution id must be specified")
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

	err = getAndUnmarshal(environment, "/institutions/all/search?"+v.Encode(), &institutions)
	return
}

// Returns all financial institutions currently supported by Plaid.
// If not specified, count defaults to 50.
// See: https://plaid.com/docs/api/#all-institutions
func (c *Client) GetInstitutions(environment environmentURL, products []string, count int, offset int) (institutions []institution, err error) {
	// Default to count=50.
	if count == 0 {
		count = 50
	}

	jsonText, err := json.Marshal(institutionsJson{
		ClientID: c.clientID,
		Secret:   c.secret,
		Products: products,
		Count:    count,
		Offset:   offset,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", string(c.environment)+"/institutions/all", bytes.NewReader(jsonText))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "plaid-go")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch {
	case res.StatusCode == 200:
		// Successful response.
		var institutionsJsonResp allInstitutionsJson
		if err = json.Unmarshal(raw, &institutionsJsonResp); err != nil {
			return nil, err
		}
		return institutionsJsonResp.Results, nil
	case res.StatusCode >= 400:
		// Attempt to unmarshal into Plaid error format
		var plaidErr plaidError
		if err = json.Unmarshal(raw, &plaidErr); err != nil {
			return nil, err
		}
		plaidErr.StatusCode = res.StatusCode
		return nil, plaidErr
	}

	return nil, errors.New("Unknown Plaid Error - Status:" + string(res.StatusCode))
}

type allInstitutionsJson struct {
	TotalCount int           `json:"total_count"`
	Results    []institution `json:"results"`
}

type institutionsJson struct {
	ClientID string   `json:"client_id"`
	Secret   string   `json:"secret"`
	Products []string `json:"products"` // e.g. ["connect", "auth", "info", "income", "risk"]
	Count    int      `json:"count"`
	Offset   int      `json:"offset"`
}

type institutionExtended struct {
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
		Light    string   `json:"light"`
		Primary  string   `json:"primary"`
		Dark     string   `json:"dark"`
		Darker   string   `json:"darker"`
		Gradient []string `json:"gradient"`
	} `json:"colors"`
	Logo      string `json:"logo"`
	Namebreak int    `json:"nameBreak"`
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
