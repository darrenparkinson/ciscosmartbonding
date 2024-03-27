// Copyright 2022 Darren Parkinson. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:generate go run gen-accessors.go
package ciscosmartbonding

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/time/rate"
)

// const apiURL = "https://cisco-test.solvedirect.com/ws/"
const apiURL = "https://stage.sbnprd.xylem.cisco.com/sb-partner-oauth-proxy-api/"

// const tokenURL = "https://cloudsso.cisco.com/as/token.oauth2"

type token struct {
	AccessToken string    `json:"access_token"`
	Scope       string    `json:"scope"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresAt   time.Time // used to keep track of when the token expires, based on ExpiresIn at the time of creation
}

// ciscoautherror represents an error we might get from the auth process
type ciscoautherror struct {
	Error       string `json:"error"`
	Description string `json:"error_description"`
}

// CiscoError represents the cisco error we might get back from Cisco for requests
type CiscoError struct {
	Message string `json:"message"`
	//TODO: They have on occasion sent an int instead of string, so we might need to cater for that
	// Status  int    `json:"status"`
	Status string `json:"status"`
}

// Client manages communication with the Cisco Smart Bonding API.
type Client struct {

	//RestyClient provides access to the resty client for using extra features
	RestyClient *resty.Client

	//TokenURL allows you to set your own token URL after initialising the client in case Cisco changes it
	TokenURL string

	//AuthType allows you to set the auth type after initialising the client in case Cisco changes it
	//It should be one of "query" or "form" for sending the data as query parameters or x-www-form-urlencoded data
	//respectively. Default is query for the old authentication.
	AuthType string

	clientID string
	secret   string

	token *token
	mu    sync.Mutex

	lim *rate.Limiter
}

// NewClient creates a new API client.  You can pass in your own resty client, or use nil for a default one.
// You can change the default test base URL after initialisation using client.RestyClient.SetBaseURL.
func NewClient(clientID, secret string, r *resty.Client) *Client {
	if r == nil {
		r = resty.New()
	}
	r.SetBaseURL(apiURL)
	r.SetTimeout(30 * time.Second)

	rl := rate.NewLimiter(150, 1) // this is not documented, so we'll limit to 150/s

	return &Client{
		RestyClient: r,
		TokenURL:    "https://cloudsso.cisco.com/as/token.oauth2",
		AuthType:    "query",
		lim:         rl,
		clientID:    clientID,
		secret:      secret,
	}

}

func (c *Client) makeListRequest(ctx context.Context, slug string, v interface{}) error {
	err := c.checkLimitAndGetToken(ctx)
	if err != nil {
		return err
	}
	resp, err := c.RestyClient.R().
		SetContext(ctx).
		SetResult(&v).
		SetError(&CiscoError{}).
		Get(slug)
	if err != nil {
		return err
	}
	if resp.IsSuccess() {
		return nil
	}
	return parseErrorResponse(resp)
}

func parseErrorResponse(resp *resty.Response) error {
	var smartbondingerr error
	switch resp.StatusCode() {
	case 400:
		smartbondingerr = ErrBadRequest
	case 401:
		smartbondingerr = ErrUnauthorized
	case 403:
		smartbondingerr = ErrForbidden
	case 404:
		smartbondingerr = ErrNotFound
	case 500:
		smartbondingerr = ErrInternalError
	default:
		smartbondingerr = ErrUnknown
	}
	if ce := resp.Error().(*CiscoError); ce != nil && ce.Message != "" {
		return fmt.Errorf("%w: %s", smartbondingerr, ce.Message)
	}
	return smartbondingerr
}

func (c *Client) checkLimitAndGetToken(ctx context.Context) error {
	if !c.lim.Allow() {
		c.lim.Wait(ctx)
	}
	err := c.getToken()
	if err != nil {
		return err
	}
	return nil
}

// getToken is a helper function to reuse an existing or retrieve a new token
func (c *Client) getToken() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.token != nil && c.token.ExpiresAt.After(time.Now().Add(time.Duration(time.Minute*5))) {
		return nil
	}
	t, err := c.generateAuthToken()
	if err != nil {
		log.Printf("error retrieving token: %s", err)
		return err
	}
	c.token = t
	c.RestyClient.SetAuthToken(t.AccessToken)
	return nil
}

func (c *Client) generateAuthToken() (*token, error) {
	var t token
	var e ciscoautherror

	if c.AuthType == "form" {
		_, err := c.RestyClient.R().
			SetFormData(map[string]string{"grant_type": "client_credentials", "client_id": c.clientID, "client_secret": c.secret}).
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetResult(&t).
			SetError(&e).
			Post(c.TokenURL)

		if err != nil {
			return &t, err
		}
	} else {
		_, err := c.RestyClient.R().
			SetQueryParams(map[string]string{"grant_type": "client_credentials", "client_id": c.clientID, "client_secret": c.secret}).
			SetHeader("Content-Type", "application/json").
			SetResult(&t).
			SetError(&e).
			Post(c.TokenURL)

		if err != nil {
			return &t, err
		}
	}
	// so a 401 isn't actually an error of course
	// we will still need to check the response
	if e.Error != "" {
		return nil, errors.New(e.Description)
	}
	t.ExpiresAt = time.Now().Add(time.Duration(t.ExpiresIn) * time.Second)
	return &t, nil
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
// func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
// func Int(v int) *int { return &v }

// Int32 is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it.
func Int32(v int32) *int32 { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Float32 is a helper routine that allocates a new Float32 value
// to store v and returns a pointer to it.
func Float32(v float32) *float32 { return &v }

// Float64 is a helper routine that allocates a new Float64 value
// to store v and returns a pointer to it.
// func Float64(v float64) *float64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
