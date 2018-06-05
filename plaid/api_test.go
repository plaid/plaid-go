package plaid

import (
	"net/http"
	"os"
)

// Credentials and parameters to be used for testing.
var (
	testClientID  = os.Getenv("PLAID_CLIENT_ID")
	testSecret    = os.Getenv("PLAID_SECRET")
	testPublicKey = os.Getenv("PLAID_PUBLIC_KEY")
	testEnv       = Sandbox

	sandboxInstitution     = "ins_109508"
	sandboxInstitutionName = "First Platypus Bank"
	testProducts           = []string{"auth", "identity", "income", "transactions"}
)

var testOptions = ClientOptions{
	testClientID,
	testSecret,
	testPublicKey,
	testEnv,
	&http.Client{},
}
var testClient, _ = NewClient(testOptions)
