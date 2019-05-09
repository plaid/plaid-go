package plaid

import (
	"sort"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetInstitutions(t *testing.T) {
	instsResp, err := testClient.GetInstitutions(2, 1)
	assert.Nil(t, err)

	expectedNames := []string{
		"Amegy Bank of Texas",
		"American Express",
	}
	outputNames := []string{}
	for _, inst := range instsResp.Institutions {
		outputNames = append(outputNames, inst.Name)
	}
	sort.Slice(expectedNames, func(i, j int) bool { return expectedNames[i] < expectedNames[j] })
	sort.Slice(outputNames, func(i, j int) bool { return outputNames[i] < outputNames[j] })
	assert.Equal(t, expectedNames, outputNames)
}

func TestSearchInstitutions(t *testing.T) {
	p := []string{"transactions"}
	instsResp, err := testClient.SearchInstitutions(sandboxInstitutionName, p)
	assert.Nil(t, err)
	assert.True(t, len(instsResp.Institutions) > 0)
}

func TestGetInstitutionsByID(t *testing.T) {
	instResp, err := testClient.GetInstitutionByID(sandboxInstitution)
	assert.Nil(t, err)
	assert.True(t, len(instResp.Institution.Products) > 0)
}
