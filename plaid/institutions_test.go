package plaid

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
)

var oauthTrue = true

func TestGetInstitutions(t *testing.T) {
	var tests = []struct {
		options    GetInstitutionsOptions
		count      int
		offset     int
		wantLength int // expected length of results
	}{
		{options: GetInstitutionsOptions{}, count: 2, offset: 1, wantLength: 2},
		{options: GetInstitutionsOptions{IncludeOptionalMetadata: true}, count: 2, offset: 1, wantLength: 2},
		{options: GetInstitutionsOptions{
			CountryCodes: []string{"GB"},
			OAuth:        &oauthTrue,
		}, count: 2, offset: 1, wantLength: 2},
		{options: GetInstitutionsOptions{
			RoutingNumbers: []string{"021200339", "052001633"},
		}, count: 1, offset: 0, wantLength: 1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#v", tt.options), func(t *testing.T) {
			instsResp, err := testClient.GetInstitutionsWithOptions(tt.count, tt.offset, tt.options)
			assert.Nil(t, err)

			assert.Len(t, instsResp.Institutions, tt.wantLength)
			for _, inst := range instsResp.Institutions {
				assert.NotEmpty(t, inst.Name)
			}

			if tt.options.IncludeOptionalMetadata {
				for _, inst := range instsResp.Institutions {
					assert.NotEmpty(t, inst.URL)
				}
			}
			if tt.options.OAuth != nil {
				for _, inst := range instsResp.Institutions {
					assert.Equal(t, inst.OAuth, *tt.options.OAuth)
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
