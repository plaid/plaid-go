package plaid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/plaid/plaid-go/internal"
)

// APIVersion holds the latest version of the Plaid API
const APIVersion = "2020-09-14"

// Client holds information required to interact with the Plaid API.
// Note: Client is only exported for method documentation purposes.
// Instances should only be created through the 'NewClient' function.
//
// See https://github.com/golang/go/issues/7823.
type Client struct {
	clientID    string
	secret      string
	environment Environment
	httpClient  *http.Client
}

type ClientOptions struct {
	ClientID    string
	Secret      string
	Environment Environment
	HTTPClient  *http.Client
}

// NewClient instantiates a Client associated with a client id, secret and environment.
func NewClient(options ClientOptions) (client *Client, err error) {
	if !options.Environment.Valid() {
		fmt.Printf("WARNING: Invalid environment specified: %s, please use one of: plaid.Sandbox, plaid.Development or plaid.Production\n", string(options.Environment))
	}

	if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{}
	}

	return &Client{
		clientID:    options.ClientID,
		secret:      options.Secret,
		environment: options.Environment,
		httpClient:  options.HTTPClient,
	}, nil
}

func (c *Client) Call(endpoint string, body []byte, v interface{}) error {
	req, err := c.newRequest(endpoint, bytes.NewReader(body), v)
	if err != nil {
		return err
	}

	return c.do(req, v)
}

// newRequest is used by Call to generate a http.Request with appropriate headers.
func (c *Client) newRequest(endpoint string, body io.Reader, v interface{}) (*http.Request, error) {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	req, err := http.NewRequest("POST", string(c.environment)+endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Plaid Go v"+internal.Version)

	// Add header for Plaid API version
	req.Header.Add("Plaid-Version", APIVersion)

	return req, nil
}

// do is used by Call to execute an http.Request and parse its response .
// Also handles parsing of the plaid error format.
func (c *Client) do(req *http.Request, v interface{}) error {
	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	// Successful response
	if res.StatusCode == 200 {
		return json.NewDecoder(res.Body).Decode(v)
	}
	// Attempt to unmarshal into Plaid error format
	var plaidErr Error
	if err = json.NewDecoder(res.Body).Decode(&plaidErr); err != nil {
		return err
	}
	plaidErr.StatusCode = res.StatusCode
	return plaidErr
}
