package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

var oauthTrue = true

func TestGetInstitutions(t *testing.T) {
	testCases := []struct {
		desc         string
		countryCodes []string
		count        int
		offset       int
		options      GetInstitutionsOptions
		wantLength   int // expected length of results
	}{
		{
			desc:         "succeeds without options",
			countryCodes: []string{"US"},
			count:        2,
			offset:       1,
			options:      GetInstitutionsOptions{},
			wantLength:   2,
		},
		{
			desc:         "succeeds with optional metadata",
			countryCodes: []string{"US"},
			count:        2,
			offset:       1,
			options:      GetInstitutionsOptions{IncludeOptionalMetadata: true},
			wantLength:   2,
		},
		{
			desc:         "succeeds for oauth institutions",
			countryCodes: []string{"GB"},
			count:        2,
			offset:       1,
			options:      GetInstitutionsOptions{OAuth: &oauthTrue},
			wantLength:   2,
		},
		{
			desc:         "errors without country codes",
			countryCodes: []string{},
			count:        2,
			offset:       1,
			options:      GetInstitutionsOptions{},
			wantLength:   0,
		},
		{
			desc:         "succeeds with routing numbers",
			countryCodes: []string{"US"},
			count:        1,
			offset:       0,
			options: GetInstitutionsOptions{
				RoutingNumbers: []string{"021200339", "052001633"},
			},
			wantLength: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			instsResp, err := testClient.GetInstitutionsWithOptions(tc.count, tc.offset, tc.countryCodes, tc.options)
			if len(tc.countryCodes) == 0 {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)

				assert.Len(t, instsResp.Institutions, tc.wantLength)
				for _, inst := range instsResp.Institutions {
					assert.NotEmpty(t, inst.Name)
				}

				if tc.options.IncludeOptionalMetadata {
					for _, inst := range instsResp.Institutions {
						assert.NotEmpty(t, inst.URL)
					}
				}
				if tc.options.OAuth != nil {
					for _, inst := range instsResp.Institutions {
						assert.Equal(t, inst.OAuth, *tc.options.OAuth)
					}
				}
			}
		})
	}
}

func TestSearchInstitutions(t *testing.T) {
	testCases := []struct {
		desc         string
		countryCodes []string
		options      SearchInstitutionsOptions
	}{
		{
			desc:         "succeeds without options",
			countryCodes: []string{"US"},
			options:      SearchInstitutionsOptions{},
		},
		{
			desc:         "succeeds with optional metadata",
			countryCodes: []string{"US"},
			options:      SearchInstitutionsOptions{IncludeOptionalMetadata: true},
		},
		{
			desc:         "succeeds for oauth institutions",
			countryCodes: []string{"GB"},
			options:      SearchInstitutionsOptions{OAuth: &oauthTrue},
		},
		{
			desc:         "errors without country codes",
			countryCodes: []string{},
			options:      SearchInstitutionsOptions{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			p := []string{"transactions"}
			instsResp, err := testClient.SearchInstitutionsWithOptions(sandboxInstitutionQuery, p, tc.countryCodes, tc.options)
			if len(tc.countryCodes) == 0 {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.True(t, len(instsResp.Institutions) > 0)

				if tc.options.IncludeOptionalMetadata {
					for _, inst := range instsResp.Institutions {
						assert.NotEmpty(t, inst.URL)
					}
				}
				if tc.options.OAuth != nil {
					for _, inst := range instsResp.Institutions {
						assert.Equal(t, inst.OAuth, *tc.options.OAuth)
					}
				}
			}
		})
	}
}

func TestGetInstitutionsByID(t *testing.T) {
	testCases := []struct {
		desc         string
		countryCodes []string
		options      GetInstitutionByIDOptions
	}{
		{
			desc:         "succeeds without options",
			countryCodes: []string{"US"},
			options:      GetInstitutionByIDOptions{},
		},
		{
			desc:         "succeeds with optional metadata",
			countryCodes: []string{"US"},
			options:      GetInstitutionByIDOptions{IncludeOptionalMetadata: true},
		},
		{
			desc:         "errors without country codes",
			countryCodes: []string{},
			options:      GetInstitutionByIDOptions{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// can't use the normal sandbox institution because it only returns the ItemLogins status.
			institutionID := "ins_12"
			instResp, err := testClient.GetInstitutionByIDWithOptions(institutionID, tc.countryCodes, tc.options)
			if len(tc.countryCodes) == 0 {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.True(t, len(instResp.Institution.Products) > 0)

				if tc.options.IncludeOptionalMetadata {
					assert.NotEmpty(t, instResp.Institution.URL)
				}
			}
		})
	}
}
