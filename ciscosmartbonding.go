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
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"golang.org/x/time/rate"
)

const apiURL = "https://cisco-test.solvedirect.com/ws/"

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
	Status  int    `json:"status"`
}

// Client manages communication with the Cisco Smart Bonding API.
type Client struct {
	// BaseURL for the API.  Set to https://cisco-test.solvedirect.com/ws/ using `ciscosmartbonding.NewClient()`, or set directly.
	// BaseURL string

	//RestyClient provides access to the resty client for using extra features
	RestyClient *resty.Client

	username string
	password string

	token *token
	mu    sync.Mutex

	lim *rate.Limiter
}

// NewClient creates a new API client.  You can pass in your own resty client, or use nil for a default one.
func NewClient(username, password string, r *resty.Client) *Client {
	if r == nil {
		r = resty.New()
	}
	r.SetBaseURL(apiURL)
	r.SetTimeout(30 * time.Second)

	rl := rate.NewLimiter(150, 1) // this is not documented, so we'll limit to 150/s

	// token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	// log.Println(token)
	return &Client{
		// BaseURL:     apiURL,
		RestyClient: r,
		lim:         rl,
		username:    username,
		password:    password,
		// basicAuthToken: token,
	}

}

func getNextLink(links []Link) (*ListParams, error) {
	for _, l := range links {
		if l.Rel == "next" {
			u, err := url.Parse(l.HRef)
			if err != nil {
				return nil, err
			}
			qp := u.Query()
			lp := &ListParams{}
			if v, ok := qp["max_id"]; ok {
				id, err := strconv.ParseInt(v[0], 10, 64)
				if err != nil {
					return nil, fmt.Errorf("error parsing max_id: %s", err)
				}
				lp.MaxId = Int64(id)
			}
			if v, ok := qp["limit"]; ok {
				lim, err := strconv.ParseInt(v[0], 10, 64)
				if err != nil {
					return nil, fmt.Errorf("error parsing limit: %s", err)
				}
				lp.Limit = Int64(lim)
			}
			if lp.GetMaxId() == 0 || lp.GetLimit() == 0 {
				return nil, errors.New("missing max id or limit in next link")
			}
			return lp, nil
		}
	}
	return nil, nil
}

func (c *Client) makeListRequest(ctx context.Context, slug string, v interface{}, listParams ...*ListParams) error {
	if !c.lim.Allow() {
		c.lim.Wait(ctx)
	}
	err := c.getToken()
	if err != nil {
		return err
	}
	listParamsString := url.Values{}
	if len(listParams) > 0 {
		listParamsString, err = query.Values(listParams[0])
	}
	if err != nil {
		return fmt.Errorf("%w: %s", ErrProcessingListParams, err)
	}
	resp, err := c.RestyClient.R().
		SetContext(ctx).
		SetQueryString(listParamsString.Encode()).
		SetResult(&v).
		SetError(&CiscoError{}).
		Get(slug)
	if err != nil {
		return err
	}
	if resp.IsSuccess() {
		return nil
	}
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

// getToken is a helper function to reuse an existing or retrieve a new token
func (c *Client) getToken() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.token != nil && c.token.ExpiresAt.After(time.Now().Add(time.Duration(time.Minute*5))) {
		return nil
	}
	t, err := c.generateAuthToken()
	if err != nil {
		log.Println("error retrieving token")
		return err
	}
	c.token = t
	c.RestyClient.SetAuthToken(t.AccessToken)
	return nil
}

func (c *Client) generateAuthToken() (*token, error) {
	var t token
	var e ciscoautherror
	_, err := c.RestyClient.R().
		SetQueryParams(map[string]string{"grant_type": "client_credentials"}).
		SetBasicAuth(c.username, c.password).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(&t).
		SetError(&e).
		Post("/rest/oauth/token")

	if err != nil {
		return &t, err
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

// Float64 is a helper routine that allocates a new Float64 value
// to store v and returns a pointer to it.
// func Float64(v float64) *float64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
