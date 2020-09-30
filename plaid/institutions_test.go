package plaid

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
)

var oauthTrue = true

func TestGetInstitutions(t *testing.T) {
	for _, options := range []GetInstitutionsOptions{
		GetInstitutionsOptions{},
		GetInstitutionsOptions{IncludeOptionalMetadata: true},
		GetInstitutionsOptions{OAuth: &oauthTrue},
	} {
		t.Run(fmt.Sprintf("%#v", options), func(t *testing.T) {
			countryCodes := []string{"US"}
			if options.OAuth == true {
				countryCodes = []string["GB"]
			}
			instsResp, err := testClient.GetInstitutionsWithOptions(2, 1, countryCodes, options)
			assert.Nil(t, err)

			assert.Len(t, instsResp.Institutions, 2)
			for _, inst := range instsResp.Institutions {
				assert.NotEmpty(t, inst.Name)
			}

			if options.IncludeOptionalMetadata {
				for _, inst := range instsResp.Institutions {
					assert.NotEmpty(t, inst.URL)
				}
			}
			if options.OAuth != nil {
				for _, inst := range instsResp.Institutions {
					assert.Equal(t, inst.OAuth, *options.OAuth)
				}
			}
		})
	}
}

func TestSearchInstitutions(t *testing.T) {
	for _, options := range []SearchInstitutionsOptions{
		SearchInstitutionsOptions{},
		SearchInstitutionsOptions{IncludeOptionalMetadata: true},
		SearchInstitutionsOptions{
			OAuth:        &oauthTrue,
		},
	} {
		t.Run(fmt.Sprintf("%#v", options), func(t *testing.T) {
			p := []string{"transactions"}
			countryCodes := []string{"US"}
			if options.OAuth == true {
				countryCodes = []string["GB"]
			}
			instsResp, err := testClient.SearchInstitutionsWithOptions(sandboxInstitutionQuery, p, countryCodes, options)
			assert.Nil(t, err)
			assert.True(t, len(instsResp.Institutions) > 0)

			if options.IncludeOptionalMetadata {
				for _, inst := range instsResp.Institutions {
					assert.NotEmpty(t, inst.URL)
				}
			}
			if options.OAuth != nil {
				for _, inst := range instsResp.Institutions {
					assert.Equal(t, inst.OAuth, *options.OAuth)
				}
			}
		})
	}
}

func TestGetInstitutionsByID(t *testing.T) {
	for _, options := range []GetInstitutionByIDOptions{
		GetInstitutionByIDOptions{},
		GetInstitutionByIDOptions{IncludeOptionalMetadata: true},
	} {
		t.Run(fmt.Sprintf("%#v", options), func(t *testing.T) {
			// can't use the normal sandbox institution because it only returns the ItemLogins status.
			institutionID := "ins_12"
			countryCodes := []string{"US"}
			instResp, err := testClient.GetInstitutionByIDWithOptions(institutionID, countryCodes, options)
			assert.Nil(t, err)
			assert.True(t, len(instResp.Institution.Products) > 0)

			if options.IncludeOptionalMetadata {
				assert.NotEmpty(t, instResp.Institution.URL)
			}
		})
	}
}
