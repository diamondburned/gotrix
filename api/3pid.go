package api

import (
	"fmt"

	"github.com/diamondburned/gotrix/api/httputil"
	"github.com/diamondburned/gotrix/matrix"
)

// ThirdpartyIdentifier is an instance of thirdparty identifier returned by the
// homeserver.
type ThirdpartyIdentifier struct {
	Medium      string           `json:"medium"`
	Address     string           `json:"address"`
	ValidatedAt matrix.Timestamp `json:"validated_at"`
	AddedAt     matrix.Timestamp `json:"added_at"`
}

// ThreePID returns all thirdparty identifiers associated with
// the current token.
func (c *Client) ThreePID() ([]ThirdpartyIdentifier, error) {
	resp := []ThirdpartyIdentifier{}
	err := c.Request(
		"GET", c.Endpoints.Account3PID(), &resp,
		httputil.WithToken(),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting 3PID: %w", err)
	}
	return resp, nil
}

// ThreePIDAdd adds the third party identifier associated with
// the client secret and session ID to the current token.
func (c *Client) ThreePIDAdd(clientSecret string, sessionID string) (*UserInteractiveAuthAPI, error) {
	var req struct {
		Auth         interface{} `json:"auth,omitempty"`
		ClientSecret string      `json:"client_secret"`
		SessionID    string      `json:"sid"`
	}
	req.ClientSecret = clientSecret
	req.SessionID = sessionID

	uiaa := &UserInteractiveAuthAPI{}
	uiaa.Request = func(auth, to interface{}) error {
		req.Auth = auth
		return c.Request(
			"POST", c.Endpoints.Account3PIDAdd(), to,
			httputil.WithToken(), httputil.WithJSONBody(req),
		)
	}
	uiaa.RequestThreePID = func(authType string, auth, to interface{}) error {
		return c.Request(
			"POST", c.Endpoints.Account3PIDRequestToken(authType), to,
			httputil.WithJSONBody(auth),
		)
	}

	err := uiaa.Auth(nil)

	return uiaa, err
}

// ThreePIDBindArg represents all required arguments of (*Client).ThreePIDBind.
type ThreePIDBindArg struct {
	ClientSecret  string `json:"client_secret"`
	IDServer      string `json:"id_server"`
	IDAccessToken string `json:"id_access_token"`
	SessionID     string `json:"sid"`
}

// ThreePIDBind binds a third party identifier connected to an identity server
// to the current token.
func (c *Client) ThreePIDBind(req ThreePIDBindArg) error {
	err := c.Request(
		"POST", c.Endpoints.Account3PIDBind(), nil,
		httputil.WithToken(), httputil.WithJSONBody(req),
	)
	if err != nil {
		return fmt.Errorf("error binding 3PID: %w", err)
	}
	return nil
}

// ThreePIDDeleteArg represents all possible arguments of (*Client).ThreePIDDelete.
type ThreePIDDeleteArg struct {
	IDServer string `json:"id_server,omitempty"`
	Medium   string `json:"medium"`
	Address  string `json:"address"`
}

// ThreePIDDelete deletes a third party identifier from the current token.
func (c *Client) ThreePIDDelete(req ThreePIDDeleteArg) (matrix.IDServerUnbindResult, error) {
	var resp struct {
		IDServerUnbindResult matrix.IDServerUnbindResult `json:"id_server_unbind_result"`
	}

	err := c.Request(
		"POST", "_matrix/client/v3/account/3pid/delete", &resp,
		httputil.WithToken(), httputil.WithJSONBody(req),
	)
	if err != nil {
		return "", fmt.Errorf("error deleting 3PID: %w", err)
	}

	return resp.IDServerUnbindResult, nil
}

// ThreePIDUnbindArg represents all possible arguments of (*Client).ThreePIDUnbind.
type ThreePIDUnbindArg struct {
	IDServer string `json:"id_server,omitempty"`
	Medium   string `json:"medium"`
	Address  string `json:"address"`
}

// ThreePIDUnbind unbinds a third party identifier from the current token.
func (c *Client) ThreePIDUnbind(req ThreePIDUnbindArg) (matrix.IDServerUnbindResult, error) {
	var resp struct {
		IDServerUnbindResult matrix.IDServerUnbindResult `json:"id_server_unbind_result"`
	}

	err := c.Request(
		"POST", "_matrix/client/v3/account/3pid/unbind", &resp,
		httputil.WithToken(), httputil.WithJSONBody(req),
	)
	if err != nil {
		return "", fmt.Errorf("error unbinding 3PID: %w", err)
	}

	return resp.IDServerUnbindResult, nil
}
