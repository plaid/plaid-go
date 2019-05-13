package plaid

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetInstitutions(t *testing.T) {
	instsResp, err := testClient.GetInstitutions(2, 1)
	assert.Nil(t, err)
	assert.Equal(t, len(instsResp.Institutions), 2)
	assert.Equal(t, instsResp.Institutions[0].Name, "American Express")
}

func TestSearchInstitutions(t *testing.T) {
	p := []string{"transactions"}
	instsResp, err := testClient.SearchInstitutions(sandboxInstitutionName, p)
	assert.Nil(t, err)
	assert.True(t, len(instsResp.Institutions) > 0)
}

func TestGetInstitutionByID(t *testing.T) {
	instResp, err := testClient.GetInstitutionByID(sandboxInstitution)
	assert.Nil(t, err)
	assert.True(t, len(instResp.Institution.Products) > 0)
}

func TestGetInstitutionByIDWithOptions(t *testing.T) {
	opts := GetInstitutionByIDOptions{
		IncludeOptionalMetadata: true,
		IncludeStatus:           true,
	}
	instResp, err := testClient.GetInstitutionByIDWithOptions(sandboxInstitution, opts)
	assert.Nil(t, err)
	assert.True(t, len(instResp.Institution.Logo) > 0)
}
