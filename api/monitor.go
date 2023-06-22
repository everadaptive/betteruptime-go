package api

import (
	"context"
	"fmt"
	"net/url"
)

type Monitor struct {
	SSLExpiration       int                  `json:"ssl_expiration,omitempty"`
	DomainExpiration    int                  `json:"domain_expiration,omitempty"`
	PolicyID            string               `json:"policy_id,omitempty"`
	URL                 string               `json:"url,omitempty"`
	MonitorType         string               `json:"monitor_type,omitempty"`
	RequiredKeyword     string               `json:"required_keyword,omitempty"`
	ExpectedStatusCodes *[]int               `json:"expected_status_codes,omitempty"`
	Call                bool                 `json:"call,omitempty"`
	SMS                 bool                 `json:"sms,omitempty"`
	Email               bool                 `json:"email,omitempty"`
	Push                bool                 `json:"push,omitempty"`
	TeamWait            int                  `json:"team_wait,omitempty"`
	Paused              bool                 `json:"paused,omitempty"`
	PausedAt            string               `json:"paused_at,omitempty"`
	FollowRedirects     bool                 `json:"follow_redirects,omitempty"`
	Port                string               `json:"port,omitempty"`
	Regions             *[]string            `json:"regions,omitempty"`
	MonitorGroupID      int                  `json:"monitor_group_id,omitempty"`
	PronounceableName   string               `json:"pronounceable_name,omitempty"`
	RecoveryPeriod      int                  `json:"recovery_period,omitempty"`
	VerifySSL           bool                 `json:"verify_ssl,omitempty"`
	CheckFrequency      int                  `json:"check_frequency,omitempty"`
	ConfirmationPeriod  int                  `json:"confirmation_period,omitempty"`
	HTTPMethod          string               `json:"http_method,omitempty"`
	RequestTimeout      int                  `json:"request_timeout,omitempty"`
	RequestBody         string               `json:"request_body,omitempty"`
	RequestHeaders      *[]map[string]string `json:"request_headers,omitempty"`
	AuthUsername        string               `json:"auth_username,omitempty"`
	AuthPassword        string               `json:"auth_password,omitempty"`
	MaintenanceFrom     string               `json:"maintenance_from,omitempty"`
	MaintenanceTo       string               `json:"maintenance_to,omitempty"`
	MaintenanceTimezone string               `json:"maintenance_timezone,omitempty"`
	RememberCookies     bool                 `json:"remember_cookies,omitempty"`
	LastCheckedAt       string               `json:"last_checked_at,omitempty"`
	Status              string               `json:"status,omitempty"`
	CreatedAt           string               `json:"created_at,omitempty"`
	UpdatedAt           string               `json:"updated_at,omitempty"`
}

type MonitorHTTPResponse struct {
	Data struct {
		ID         string  `json:"id"`
		Attributes Monitor `json:"attributes"`
	} `json:"data"`
}

type MonitorListHTTPResponse struct {
	Data []struct {
		ID         string  `json:"id"`
		Attributes Monitor `json:"attributes"`
	} `json:"data"`
}

func MonitorCreate(ctx context.Context, meta interface{}, in Monitor) (MonitorHTTPResponse, error) {
	var out MonitorHTTPResponse
	if err := resourceCreate(ctx, meta, "/api/v2/monitors", &in, &out); err != nil {
		return out, err
	}

	return out, nil
}

func MonitorGetByID(ctx context.Context, meta interface{}, id string) (MonitorHTTPResponse, error) {
	var out MonitorHTTPResponse
	if err, ok := resourceRead(ctx, meta, fmt.Sprintf("/api/v2/monitors/%s", url.PathEscape(id)), &out); err != nil {
		return out, err
	} else if !ok {
		return MonitorHTTPResponse{}, nil
	}

	return out, nil
}

func MonitorGet(ctx context.Context, meta interface{}, query string) (MonitorListHTTPResponse, error) {
	var out MonitorListHTTPResponse
	if err, ok := resourceRead(ctx, meta, fmt.Sprintf("/api/v2/monitors?%s", query), &out); err != nil {
		return out, err
	} else if !ok {
		return MonitorListHTTPResponse{}, nil
	}

	return out, nil
}

func MonitorUpdate(ctx context.Context, meta interface{}, id string, req Monitor) (MonitorHTTPResponse, error) {
	var out MonitorHTTPResponse
	if err := resourceUpdate(ctx, meta, fmt.Sprintf("/api/v2/monitors/%s", url.PathEscape(id)), &req, &out); err != nil {
		return out, err
	}

	return out, nil
}

func MonitorDelete(ctx context.Context, meta interface{}, id string) error {
	return resourceDelete(ctx, meta, fmt.Sprintf("/api/v2/monitors/%s", url.PathEscape(id)))
}
