package plaid

import (
	"fmt"
	"sort"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetInstitutions(t *testing.T) {
	for _, options := range []GetInstitutionsOptions{
		GetInstitutionsOptions{},
		GetInstitutionsOptions{IncludeOptionalMetadata: true},
	} {
		t.Run(fmt.Sprintf("%#v", options), func(t *testing.T) {
			instsResp, err := testClient.GetInstitutionsWithOptions(2, 1, options)
			assert.Nil(t, err)

			expectedNames := []string{
				"Amegy Bank of Texas - Personal Banking",
				"American Express",
			}
			outputNames := []string{}
			for _, inst := range instsResp.Institutions {
				outputNames = append(outputNames, inst.Name)
			}
			sort.Slice(expectedNames, func(i, j int) bool { return expectedNames[i] < expectedNames[j] })
			sort.Slice(outputNames, func(i, j int) bool { return outputNames[i] < outputNames[j] })
			assert.Equal(t, expectedNames, outputNames)

			if options.IncludeOptionalMetadata {
				for _, inst := range instsResp.Institutions {
					assert.NotEmpty(t, inst.URL)
				}
			}
		})
	}
}

func TestSearchInstitutions(t *testing.T) {
	for _, options := range []SearchInstitutionsOptions{
		SearchInstitutionsOptions{},
		SearchInstitutionsOptions{IncludeOptionalMetadata: true},
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
			instResp, err := testClient.GetInstitutionByIDWithOptions(sandboxInstitution, options)
			assert.Nil(t, err)
			assert.True(t, len(instResp.Institution.Products) > 0)

			if options.IncludeOptionalMetadata {
				assert.NotEmpty(t, instResp.Institution.URL)
			}

			if options.IncludeStatus {
				assert.NotEmpty(t, instResp.Institution.InstitutionStatus)
				assert.True(t, instResp.Institution.InstitutionStatus.ItemLogins.LastStatusChange.Unix() > 0)
			}
		})
	}
}
