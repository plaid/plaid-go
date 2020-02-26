package plaid

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetInstitutions(t *testing.T) {
	for _, options := range []GetInstitutionsOptions{
		GetInstitutionsOptions{},
		GetInstitutionsOptions{IncludeOptionalMetadata: true},
		GetInstitutionsOptions{
			CountryCodes: []string{"GB"},
			OAuth:        true,
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

			if options.OAuth {
				for _, inst := range instsResp.Institutions {
					assert.True(t, inst.OAuth)
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
			OAuth:        true,
		},
	} {
		t.Run(fmt.Sprintf("%#v", options), func(t *testing.T) {
			p := []string{"transactions"}
			instsResp, err := testClient.SearchInstitutionsWithOptions(sandboxInstitutionName, p, options)
			assert.Nil(t, err)
			assert.True(t, len(instsResp.Institutions) > 0)

			if options.IncludeOptionalMetadata {
				for _, inst := range instsResp.Institutions {
					assert.NotEmpty(t, inst.URL)
				}
			}

			if options.OAuth {
				for _, inst := range instsResp.Institutions {
					assert.True(t, inst.OAuth)
				}
			}
		})
	}
}

func TestGetInstitutionsByID(t *testing.T) {
	for _, options := range []GetInstitutionByIDOptions{
		GetInstitutionByIDOptions{},
		GetInstitutionByIDOptions{IncludeOptionalMetadata: true},
		GetInstitutionByIDOptions{IncludeOptionalMetadata: true, IncludeStatus: true},
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

			if options.IncludeStatus {
				assert.NotEmpty(t, instResp.Institution.InstitutionStatus)
				assert.True(t, instResp.Institution.InstitutionStatus.ItemLogins.LastStatusChange.Unix() > 0)
				assert.True(t, instResp.Institution.InstitutionStatus.TransactionsUpdates.LastStatusChange.Unix() > 0)
				assert.True(t, instResp.Institution.InstitutionStatus.Auth.LastStatusChange.Unix() > 0)
				assert.True(t, instResp.Institution.InstitutionStatus.Balance.LastStatusChange.Unix() > 0)
				assert.True(t, instResp.Institution.InstitutionStatus.Identity.LastStatusChange.Unix() > 0)
			}
		})
	}
}
