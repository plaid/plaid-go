package plaid

import (
	"encoding/json"
	"errors"
)

type AssetReport struct {
	AssetReportID  string            `json:"asset_report_id"`
	ClientReportID string            `json:"client_report_id"`
	DateGenerated  string            `json:"date_generated"`
	DaysRequested  int               `json:"days_requested"`
	Items          []AssetReportItem `json:"items"`
	User           AssetReportUser   `json:"user"`
}

type AssetReportItem struct {
	Accounts        []Account `json:"accounts"`
	DateLastUpdated string    `json:"date_last_updated"`
	InstitutionID   string    `json:"institution_id"`
	InstitutionName string    `json:"institution_name"`
	ItemID          string    `json:"item_id"`
}

type AssetReportUser struct {
	ClientID    string `json:"client_user_id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	PhoneNumber string `json:"phone_number"`
	SSN         string `json:"ssn"`
}

type getAssetReportRequest struct {
	ClientID         string `json:"client_id"`
	Secret           string `json:"secret"`
	AssetReportToken string `json:"asset_report_token"`
}

type GetAssetReportResponse struct {
	APIResponse
	Report   AssetReport `json:"report"`
	Warnings []string    `json:"warnings"`
}

type removeAssetReportRequest struct {
	ClientID         string `json:"client_id"`
	Secret           string `json:"secret"`
	AssetReportToken string `json:"asset_report_token"`
}

type RemoveAssetReportResponse struct {
	APIResponse
	Removed bool `json:"removed"`
}

type createAuditCopyRequest struct {
	ClientID         string `json:"client_id"`
	Secret           string `json:"secret"`
	AssetReportToken string `json:"asset_report_token"`
	AuditorID        string `json:"auditor_id"`
}

type CreateAuditCopyTokenResponse struct {
	APIResponse
	AuditCopyToken string `json:"audit_copy_token"`
}

func (c *Client) GetAssetReport(assetReportToken string) (resp GetAssetReportResponse, err error) {
	if assetReportToken == "" {
		return resp, errors.New("/asset_report/get - asset report token must be specified")
	}

	jsonBody, err := json.Marshal(getAssetReportRequest{
		ClientID:         c.clientID,
		Secret:           c.secret,
		AssetReportToken: assetReportToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/asset_report/get", jsonBody, &resp)
	return resp, err
}

func (c *Client) CreateAuditCopy(assetReportToken, auditorID string) (resp CreateAuditCopyTokenResponse, err error) {
	if assetReportToken == "" || auditorID == "" {
		return resp, errors.New("/asset_report/audit_copy/create - asset report token and auditor id must be specified")
	}

	jsonBody, err := json.Marshal(createAuditCopyRequest{
		ClientID:         c.clientID,
		Secret:           c.secret,
		AssetReportToken: assetReportToken,
		AuditorID:        auditorID,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/asset_report/audit_copy/create", jsonBody, &resp)
	return resp, err
}

func (c *Client) RemoveAssetReport(assetReportToken string) (resp RemoveAssetReportResponse, err error) {
	if assetReportToken == "" {
		return resp, errors.New("/asset_report/remove - asset report token must be specified")
	}

	jsonBody, err := json.Marshal(removeAssetReportRequest{
		ClientID:         c.clientID,
		Secret:           c.secret,
		AssetReportToken: assetReportToken,
	})

	if err != nil {
		return resp, err
	}

	err = c.Call("/asset_report/remove", jsonBody, &resp)
	return resp, err
}
