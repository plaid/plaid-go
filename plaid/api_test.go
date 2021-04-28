package plaid

import (
	"net/http"
	"os"
)

// Credentials and parameters to be used for testing.
var (
	testClientID = os.Getenv("PLAID_CLIENT_ID")
	testSecret   = os.Getenv("PLAID_SECRET")
	testEnv      = Sandbox

	sandboxInstitution                               = "ins_109508"
	sandboxInstitutionQuery                          = "Platypus"
	paymentInitiationMetadataSandboxInstitution      = "ins_117650"
	paymentInitiationMetadataSandboxInstitutionQuery = "Royal Bank of Plaid"
	testProducts                                     = []string{"auth", "identity", "transactions"}
)

var testOptions = ClientOptions{
	testClientID,
	testSecret,
	testEnv,
	&http.Client{},
}
var testClient, _ = NewClient(testOptions)
