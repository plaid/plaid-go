package plaid

import (
	"encoding/json"
	"errors"
	"io"
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
var TartanSandbox environmentURL = "https://tartan.plaid.com"
var Production environmentURL = "https://api.plaid.com"

type account struct {
	ID      string `json:"_id"`
	ItemID  string `json:"_item"`
	UserID  string `json:"_user"`
	Balance struct {
		Available float64 `json:"available"`
		Current   float64 `json:"current"`
	}
	Meta struct {
		Number string `json:"number"`
		Name   string `json:"name"`
	}
	Numbers struct {
		Account     string `json:"account"`
		Routing     string `json:"routing"`
		WireRouting string `json:"wireRouting"`
	}
	Type            string `json:"type"`
	InstitutionType string `json:"institution_type"`
}

type mfaIntermediate struct {
	AccessToken string      `json:"access_token"`
	MFA         interface{} `json:"mfa"`
	Type        string      `json:"type"`
}
type mfaDevice struct {
	Message string
}
type mfaList struct {
	Mask string
	Type string
}
type mfaQuestion struct {
	Question string
}
type mfaSelection struct {
	Answers  []string
	Question string
}

// 'mfa' contains the union of all possible mfa types
// Users should switch on the 'Type' field
type mfaResponse struct {
	AccessToken string
	Type        string

	Device     mfaDevice
	List       []mfaList
	Questions  []mfaQuestion
	Selections []mfaSelection
}

type postResponse struct {
	// Normal response fields
	AccessToken string    `json:"access_token"`
	Accounts    []account `json:"accounts"`
	MFA         string    `json:"mfa"`
}

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

	// Successful response
	if res.StatusCode == 200 {
		if err = json.Unmarshal(raw, structure); err != nil {
			return err
		}
		return nil
	}
	// Attempt to unmarshal into Plaid error format
	var plaidErr plaidError
	if err = json.Unmarshal(raw, &plaidErr); err != nil {
		return err
	}
	plaidErr.StatusCode = res.StatusCode
	return plaidErr
}

func postAndUnmarshal(environment environmentURL, endpoint string,
	body io.Reader) (*postResponse, *mfaResponse, error) {
	// Read response body
	res, err := http.Post(string(environment)+endpoint, "application/json", body)
	if err != nil {
		return nil, nil, err
	}
	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	res.Body.Close()

	// Different marshaling cases
	var mfaInter mfaIntermediate
	var postRes postResponse
	switch {
	// Successful response
	case res.StatusCode == 200:
		if err = json.Unmarshal(raw, &postRes); err != nil {
			return nil, nil, err
		}
		return &postRes, nil, nil

	// MFA case
	case res.StatusCode == 201:
		if err = json.Unmarshal(raw, &mfaInter); err != nil {
			return nil, nil, err
		}
		mfaRes := mfaResponse{Type: mfaInter.Type, AccessToken: mfaInter.AccessToken}
		switch mfaInter.Type {
		case "device":
			temp, ok := mfaInter.MFA.(interface{})
			if !ok {
				return nil, nil, errors.New("Could not decode device mfa")
			}
			deviceStruct, ok := temp.(map[string]interface{})
			if !ok {
				return nil, nil, errors.New("Could not decode device mfa")
			}
			deviceText, ok := deviceStruct["message"].(string)
			if !ok {
				return nil, nil, errors.New("Could not decode device mfa")
			}
			mfaRes.Device.Message = deviceText

		case "list":
			temp, ok := mfaInter.MFA.([]interface{})
			if !ok {
				return nil, nil, errors.New("Could not decode list mfa")
			}
			for _, v := range temp {
				listArray, ok := v.(map[string]interface{})
				if !ok {
					return nil, nil, errors.New("Could not decode list mfa")
				}
				maskText, ok := listArray["mask"].(string)
				if !ok {
					return nil, nil, errors.New("Could not decode list mfa")
				}
				typeText, ok := listArray["type"].(string)
				if !ok {
					return nil, nil, errors.New("Could not decode list mfa")
				}
				mfaRes.List = append(mfaRes.List, mfaList{maskText, typeText})
			}

		case "questions":
			questions, ok := mfaInter.MFA.([]interface{})
			if !ok {
				return nil, nil, errors.New("Could not decode questions mfa")
			}
			for _, v := range questions {
				q, ok := v.(map[string]interface{})
				if !ok {
					return nil, nil, errors.New("Could not decode questions mfa")
				}
				questionText, ok := q["question"].(string)
				if !ok {
					return nil, nil, errors.New("Could not decode questions mfa question")
				}
				mfaRes.Questions = append(mfaRes.Questions, mfaQuestion{questionText})
			}

		case "selections":
			selections, ok := mfaInter.MFA.([]interface{})
			if !ok {
				return nil, nil, errors.New("Could not decode selections mfa")
			}
			for _, v := range selections {
				s, ok := v.(map[string]interface{})
				if !ok {
					return nil, nil, errors.New("Could not decode selections mfa")
				}
				tempAnswers, ok := s["answers"].([]interface{})
				if !ok {
					return nil, nil, errors.New("Could not decode selections answers")
				}
				answers := make([]string, len(tempAnswers))
				for i, a := range tempAnswers {
					answers[i], ok = a.(string)
				}
				if !ok {
					return nil, nil, errors.New("Could not decode selections answers")
				}
				question, ok := s["question"].(string)
				if !ok {
					return nil, nil, errors.New("Could not decode selections questions")
				}
				mfaRes.Selections = append(mfaRes.Selections, mfaSelection{answers, question})
			}
		}
		return nil, &mfaRes, nil

	// Error case, attempt to unmarshal into Plaid error format
	case res.StatusCode >= 400:
		var plaidErr plaidError
		if err = json.Unmarshal(raw, &plaidErr); err != nil {
			return nil, nil, err
		}
		plaidErr.StatusCode = res.StatusCode
		return nil, nil, plaidErr
	}
	return nil, nil, errors.New("Unknown Plaid Error - Status:" + string(res.StatusCode))
}
