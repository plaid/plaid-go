package plaid

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/plaid/plaid-go/internal"
)

// APIVersion holds the latest version of the Plaid API
const APIVersion = "2019-05-29"

// Client holds information required to interact with the Plaid API.
// Note: Client is only exported for method documentation purposes.
// Instances should only be created through the 'NewClient' function.
//
// See https://github.com/golang/go/issues/7823.
type Client struct {
	clientID    string
	secret      string
	publicKey   string
	environment Environment
	httpClient  *http.Client
}

type ClientOptions struct {
	ClientID    string
	Secret      string
	PublicKey   string
	Environment Environment
	HTTPClient  *http.Client
}

// NewClient instantiates a Client associated with a client id, secret and environment.
func NewClient(options ClientOptions) (client *Client, err error) {
	if !options.Environment.Valid() {
		return nil, errors.New("Invalid environment specified: " + string(options.Environment))
	}

	if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{}
	}

	return &Client{
		clientID:    options.ClientID,
		secret:      options.Secret,
		publicKey:   options.PublicKey,
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

	req, err := http.NewRequest("POST", "https://"+string(c.environment)+endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Plaid Go v"+internal.Version)

	// Add header for Plaid API version
	req.Header.Add("Plaid-Version", APIVersion)

	return req, nil
}

// do is used by Call to execute an http.Request and parse its response.
// Also handles parsing of the plaid error format.
func (c *Client) do(req *http.Request, v interface{}) error {
	res, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Successful response
	if res.StatusCode == 200 {
		return json.Unmarshal(resBody, v)
	}

	// Attempt to unmarshal into Plaid error format
	var plaidErr Error
	if err = json.Unmarshal(resBody, &plaidErr); err != nil {
		return err
	}
	plaidErr.StatusCode = res.StatusCode
	return plaidErr
}
