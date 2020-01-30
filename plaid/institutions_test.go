package plaid

import (
	"context"
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetInstitutions(t *testing.T) {
	for _, options := range []GetInstitutionsOptions{
		GetInstitutionsOptions{},
		GetInstitutionsOptions{IncludeOptionalMetadata: true},
	} {
		t.Run(fmt.Sprintf("%#v", options), func(t *testing.T) {
			instsResp, err := testClient.GetInstitutionsWithOptions(context.Background(), 2, 1, options)
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
			instsResp, err := testClient.SearchInstitutionsWithOptions(context.Background(), sandboxInstitutionName, p, options)
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
	} {
		t.Run(fmt.Sprintf("%#v", options), func(t *testing.T) {
			instResp, err := testClient.GetInstitutionByIDWithOptions(context.Background(), sandboxInstitution, options)
			assert.Nil(t, err)
			assert.True(t, len(instResp.Institution.Products) > 0)

			if options.IncludeOptionalMetadata {
				assert.NotEmpty(t, instResp.Institution.URL)
			}
		})
	}
}
