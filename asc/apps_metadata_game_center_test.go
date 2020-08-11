package asc

import (
	"context"
	"testing"
)

func TestListGameCenterEnabledVersionsForApp(t *testing.T) {
	testEndpointWithResponse(t, "{}", &GameCenterEnabledVersionsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListGameCenterEnabledVersionsForApp(ctx, "10", &ListGameCenterEnabledVersionsForAppQuery{})
	})
}

func TestListCompatibleVersionsForGameCenterEnabledVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &GameCenterEnabledVersionsResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListCompatibleVersionsForGameCenterEnabledVersion(ctx, "10", &ListCompatibleVersionsForGameCenterEnabledVersionQuery{})
	})
}

func TestListCompatibleVersionIDsForGameCenterEnabledVersion(t *testing.T) {
	testEndpointWithResponse(t, "{}", &GameCenterEnabledVersionCompatibleVersionsLinkagesResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.Apps.ListCompatibleVersionIDsForGameCenterEnabledVersion(ctx, "10", &ListCompatibleVersionIDsForGameCenterEnabledVersionQuery{})
	})
}

func TestCreateCompatibleVersionsForGameCenterEnabledVersion(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.CreateCompatibleVersionsForGameCenterEnabledVersion(ctx, "10", GameCenterEnabledVersionCompatibleVersionsLinkagesRequest{})
	})
}

func TestCreateCompatibleVersionsForGameCenterEnabledVersionNoRequest(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.CreateCompatibleVersionsForGameCenterEnabledVersion(ctx, "10", nil)
	})
}

func TestUpdateCompatibleVersionsForGameCenterEnabledVersion(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.UpdateCompatibleVersionsForGameCenterEnabledVersion(ctx, "10", GameCenterEnabledVersionCompatibleVersionsLinkagesRequest{})
	})
}

func TestUpdateCompatibleVersionsForGameCenterEnabledVersionNoRequest(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.UpdateCompatibleVersionsForGameCenterEnabledVersion(ctx, "10", nil)
	})
}

func TestRemoveCompatibleVersionsForGameCenterEnabledVersion(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.RemoveCompatibleVersionsForGameCenterEnabledVersion(ctx, "10", GameCenterEnabledVersionCompatibleVersionsLinkagesRequest{})
	})
}

func TestRemoveCompatibleVersionsForGameCenterEnabledVersionNoRequest(t *testing.T) {
	testEndpointWithNoContent(t, func(ctx context.Context, client *Client) (*Response, error) {
		return client.Apps.RemoveCompatibleVersionsForGameCenterEnabledVersion(ctx, "10", nil)
	})
}