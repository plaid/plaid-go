package plaid_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Keep the mock server's expectations independent of generated client setup:
// tests must fail if regeneration starts inheriting credentials from the shell.
func TestWireMockRequiresDummyCredentials(t *testing.T) {
	data, err := os.ReadFile("wiremock/wiremock-mappings.json")
	require.NoError(t, err)
	var configuration struct {
		Mappings []struct {
			Request struct {
				URLPath string `json:"urlPath"`
				Headers map[string]struct {
					EqualTo string `json:"equalTo"`
				} `json:"headers"`
			} `json:"request"`
		} `json:"mappings"`
	}
	require.NoError(t, json.Unmarshal(data, &configuration))
	require.NotEmpty(t, configuration.Mappings)
	for _, mapping := range configuration.Mappings {
		assert.Equal(t, "wiremock-client-id", mapping.Request.Headers["PLAID-CLIENT-ID"].EqualTo, mapping.Request.URLPath)
		assert.Equal(t, "wiremock-secret", mapping.Request.Headers["PLAID-SECRET"].EqualTo, mapping.Request.URLPath)
	}
}

func TestWireMockOnlyBindsLoopback(t *testing.T) {
	data, err := os.ReadFile("wiremock/docker-compose.test.yml")
	require.NoError(t, err)
	assert.Contains(t, string(data), `"127.0.0.1:0:8080"`)
	assert.NotContains(t, string(data), `"0:8080"`)
	assert.NotContains(t, string(data), "--verbose")
}
