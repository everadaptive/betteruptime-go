package api

import (
	"context"
	"fmt"
	"net/url"
)

type MonitorGroup struct {
	Paused    bool   `json:"paused,omitempty"`
	Name      string `json:"name,omitempty"`
	SortIndex int    `json:"sort_index,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type MonitorGroupHTTPResponse struct {
	Data struct {
		ID         string       `json:"id"`
		Attributes MonitorGroup `json:"attributes"`
	} `json:"data"`
}

type MonitorGroupListHTTPResponse struct {
	Data []struct {
		ID         string       `json:"id"`
		Attributes MonitorGroup `json:"attributes"`
	} `json:"data"`
}

func MonitorGroupCreate(ctx context.Context, meta interface{}, in MonitorGroup) (MonitorGroupHTTPResponse, error) {
	var out MonitorGroupHTTPResponse
	if err := resourceCreate(ctx, meta, "/api/v2/monitor-groups", &in, &out); err != nil {
		return out, err
	}

	return out, nil
}

func MonitorGroupGet(ctx context.Context, meta interface{}, id string) (MonitorGroupListHTTPResponse, error) {
	var out MonitorGroupListHTTPResponse
	if err, ok := resourceRead(ctx, meta, fmt.Sprintf("/api/v2/monitor-groups/%s", url.PathEscape(id)), &out); err != nil {
		return out, err
	} else if !ok {
		return MonitorGroupListHTTPResponse{}, nil
	}

	return out, nil
}

func MonitorGroupUpdate(ctx context.Context, meta interface{}, id string, req MonitorGroup) (MonitorGroupHTTPResponse, error) {
	var out MonitorGroupHTTPResponse
	if err := resourceUpdate(ctx, meta, fmt.Sprintf("/api/v2/monitor-groups/%s", url.PathEscape(id)), &req, &out); err != nil {
		return out, err
	}

	return out, nil
}

func MonitorGroupDelete(ctx context.Context, meta interface{}, id string) error {
	return resourceDelete(ctx, meta, fmt.Sprintf("/api/v2/monitor-groups/%s", url.PathEscape(id)))
}
