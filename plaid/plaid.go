package plaid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewClient(clientID, secret string, environment environmentURL) client {
	return client{clientID, secret, environment}
}

type client struct {
	clientID    string
	secret      string
	environment environmentURL
}

type environmentURL string

var Tartan environmentURL = "https://tartan.plaid.com"
var Production environmentURL = "https://api.plaid.com"

func getAndUnmarshal(environment environmentURL, endpoint string, structure interface{}) error {
	res, err := http.Get(string(environment) + endpoint)
	if err != nil {
		return err
	}
	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	return json.Unmarshal(raw, structure)
}
