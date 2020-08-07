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
		GetInstitutionsOptions{
			CountryCodes: []string{"GB"},
			OAuth:        &oauthTrue,
		},
		GetInstitutionsOptions{
			RoutingNumbers: []string{"011000138"}, // ins_1
		},
	} {
		t.Run(fmt.Sprintf("%#v", options), func(t *testing.T) {
			instsResp, err := testClient.GetInstitutionsWithOptions(2, 1, options)
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

			if options.RoutingNumbers != nil {
				for _, inst := range instsResp.Institutions {
					assert.Equal(t, inst.ID, "ins_1")
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
			CountryCodes: []string{"GB"},
			OAuth:        &oauthTrue,
		},
	} {
		t.Run(fmt.Sprintf("%#v", options), func(t *testing.T) {
			p := []string{"transactions"}
			instsResp, err := testClient.SearchInstitutionsWithOptions(sandboxInstitutionQuery, p, options)
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

			instResp, err := testClient.GetInstitutionByIDWithOptions(institutionID, options)
			assert.Nil(t, err)
			assert.True(t, len(instResp.Institution.Products) > 0)

			if options.IncludeOptionalMetadata {
				assert.NotEmpty(t, instResp.Institution.URL)
			}
		})
	}
}
