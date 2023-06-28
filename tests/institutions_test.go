package tests

import (
	"context"
	"testing"

	"github.com/plaid/plaid-go/v14/plaid"
	"github.com/stretchr/testify/assert"
)

func TestInstitutionsGet(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	testCases := []struct {
		desc           string
		countryCodes   []plaid.CountryCode
		count          int32
		offset         int32
		options        plaid.InstitutionsGetRequestOptions
		expectedLength int // expected length of results
	}{
		{
			desc:           "succeeds without options",
			countryCodes:   []plaid.CountryCode{plaid.COUNTRYCODE_US},
			count:          3,
			offset:         4,
			options:        plaid.InstitutionsGetRequestOptions{},
			expectedLength: 3,
		},
		{
			desc:           "succeeds with optional metadata",
			countryCodes:   []plaid.CountryCode{plaid.COUNTRYCODE_US},
			count:          3,
			offset:         4,
			options:        plaid.InstitutionsGetRequestOptions{IncludeOptionalMetadata: plaid.PtrBool(true)},
			expectedLength: 3,
		},
		{
			desc:           "succeeds for oauth institutions",
			countryCodes:   []plaid.CountryCode{plaid.COUNTRYCODE_GB},
			count:          3,
			offset:         1,
			options:        plaid.InstitutionsGetRequestOptions{Oauth: *plaid.NewNullableBool(plaid.PtrBool(true))},
			expectedLength: 3,
		},
		{
			desc:           "errors with invalid offset",
			count:          2,
			offset:         -1,
			options:        plaid.InstitutionsGetRequestOptions{},
			expectedLength: 0,
		},
		{
			desc:         "succeeds with routing numbers",
			countryCodes: []plaid.CountryCode{plaid.COUNTRYCODE_US},
			count:        1,
			offset:       0,
			options: plaid.InstitutionsGetRequestOptions{
				RoutingNumbers: []string{"021200339", "052001633"},
			},
			expectedLength: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			request := plaid.NewInstitutionsGetRequest(tc.count, tc.offset, tc.countryCodes)
			request.SetOptions(tc.options)
			instsResp, _, err := testClient.PlaidApi.InstitutionsGet(ctx).InstitutionsGetRequest(*request).Execute()
			if len(tc.countryCodes) == 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

				assert.Len(t, instsResp.Institutions, tc.expectedLength)
				for _, inst := range instsResp.Institutions {
					assert.NotEmpty(t, inst.Name)
				}

				if tc.options.IncludeOptionalMetadata != nil && *tc.options.IncludeOptionalMetadata {
					for _, inst := range instsResp.Institutions {
						assert.NotEmpty(t, inst.Url)
					}
				}
				if tc.options.Oauth.IsSet() {
					for _, inst := range instsResp.Institutions {
						assert.Equal(t, inst.Oauth, *tc.options.Oauth.Get())
					}
				}
			}
		})
	}
}

func TestInstitutionsSearch(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	sandboxInstitutionQuery := "Platypus"
	paymentInitiationMetadataSandboxInstitutionQuery := "Royal Bank of Plaid"
	products := []plaid.Products{plaid.PRODUCTS_TRANSACTIONS}
	testCases := []struct {
		desc         string
		query        string
		countryCodes []plaid.CountryCode
		options      plaid.InstitutionsSearchRequestOptions
	}{
		{
			desc:         "succeeds without options",
			query:        sandboxInstitutionQuery,
			countryCodes: []plaid.CountryCode{plaid.COUNTRYCODE_US},
			options:      plaid.InstitutionsSearchRequestOptions{},
		},
		{
			desc:         "succeeds with optional metadata",
			query:        sandboxInstitutionQuery,
			countryCodes: []plaid.CountryCode{plaid.COUNTRYCODE_US},
			options:      plaid.InstitutionsSearchRequestOptions{IncludeOptionalMetadata: plaid.PtrBool(true)},
		},
		{
			desc:         "succeeds with payment initiation metadata",
			query:        paymentInitiationMetadataSandboxInstitutionQuery,
			countryCodes: []plaid.CountryCode{plaid.COUNTRYCODE_GB},
			options:      plaid.InstitutionsSearchRequestOptions{IncludePaymentInitiationMetadata: *plaid.NewNullableBool(plaid.PtrBool(true))},
		},
		{
			desc:         "succeeds for oauth institutions",
			query:        sandboxInstitutionQuery,
			countryCodes: []plaid.CountryCode{plaid.COUNTRYCODE_GB},
			options:      plaid.InstitutionsSearchRequestOptions{Oauth: *plaid.NewNullableBool(plaid.PtrBool(true))},
		},
		{
			desc:    "errors without empty query",
			query:   "",
			options: plaid.InstitutionsSearchRequestOptions{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			request := plaid.NewInstitutionsSearchRequest(tc.query, products, tc.countryCodes)
			request.SetOptions(tc.options)
			instsResp, _, err := testClient.PlaidApi.InstitutionsSearch(ctx).InstitutionsSearchRequest(*request).Execute()
			if len(tc.countryCodes) == 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, instsResp.Institutions)

				if tc.options.IncludeOptionalMetadata != nil {
					for _, inst := range instsResp.Institutions {
						assert.NotEmpty(t, inst.Url)
					}
				}
				if tc.options.IncludePaymentInitiationMetadata.IsSet() {
					for _, inst := range instsResp.Institutions {
						assert.NotEmpty(t, inst.PaymentInitiationMetadata)
					}
				}
			}
		})
	}
}

func TestInstitutionsGetById(t *testing.T) {
	testClient := NewTestClient()
	ctx := context.Background()

	paymentInitiationMetadataSandboxInstitution := "ins_117650"
	testCases := []struct {
		desc          string
		institutionID string
		countryCodes  []plaid.CountryCode
		options       plaid.InstitutionsGetByIdRequestOptions
	}{
		{
			desc: "succeeds without options",
			// can't use the normal sandbox institution because it only returns the ItemLogins status.
			institutionID: "ins_12",
			countryCodes:  []plaid.CountryCode{plaid.COUNTRYCODE_US},
			options:       plaid.InstitutionsGetByIdRequestOptions{},
		},
		{
			desc:          "succeeds with optional metadata",
			institutionID: "ins_12",
			countryCodes:  []plaid.CountryCode{plaid.COUNTRYCODE_US},
			options:       plaid.InstitutionsGetByIdRequestOptions{IncludeOptionalMetadata: plaid.PtrBool(true)},
		},
		{
			desc:          "succeeds with payment initiation metadata",
			institutionID: paymentInitiationMetadataSandboxInstitution,
			countryCodes:  []plaid.CountryCode{plaid.COUNTRYCODE_GB},
			options:       plaid.InstitutionsGetByIdRequestOptions{IncludePaymentInitiationMetadata: plaid.PtrBool(true)},
		},
		{
			desc:          "errors with invalid ins ID",
			institutionID: "ins_-1",
			options:       plaid.InstitutionsGetByIdRequestOptions{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			request := plaid.NewInstitutionsGetByIdRequest(tc.institutionID, tc.countryCodes)
			request.SetOptions(tc.options)
			instResp, _, err := testClient.PlaidApi.InstitutionsGetById(ctx).InstitutionsGetByIdRequest(*request).Execute()
			if len(tc.countryCodes) == 0 {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, instResp.Institution.Products)

				if tc.options.IncludeOptionalMetadata != nil {
					assert.NotEmpty(t, instResp.Institution.Url)
				}

				if tc.options.IncludePaymentInitiationMetadata != nil {
					assert.NotEmpty(t, instResp.Institution.PaymentInitiationMetadata)
				}
			}
		})
	}
}
