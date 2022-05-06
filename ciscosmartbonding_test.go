package ciscosmartbonding

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	baseURLPath := "/ws"
	apiHandler := http.NewServeMux()
	server := httptest.NewServer(apiHandler)
	client = NewClient("dummy", "dummy", nil)
	burl, _ := url.Parse(server.URL + baseURLPath)
	client.RestyClient.SetBaseURL(burl.String())
	// client.RestyClient.SetDebug(true)
	client.token = &token{AccessToken: "dummy", Scope: "data", TokenType: "Bearer", ExpiresIn: 3600, ExpiresAt: time.Now().Add(time.Duration(3600) * time.Second)}
	return client, apiHandler, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestErrorResponses(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	ctx := context.Background()

	var tests = []struct {
		name string
		code int
		want error
	}{
		{"badrequest", http.StatusBadRequest, ErrBadRequest},
		{"unauthorized", http.StatusUnauthorized, ErrUnauthorized},
		{"forbidden", http.StatusForbidden, ErrForbidden},
		{"notfound", http.StatusNotFound, ErrNotFound},
		{"internalerror", http.StatusInternalServerError, ErrInternalError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			turl := fmt.Sprintf("/ws/testing/%s", tt.name)
			mux.HandleFunc(turl, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				errordata := `{"message":"something bad happened","status":400}`
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.code)
				fmt.Fprint(w, errordata)
			})
			var e CiscoError
			err := client.makeListRequest(ctx, fmt.Sprintf("/testing/%s", tt.name), &e)
			if err == nil {
				t.Fatal("Expected error")
			}
			if !errors.Is(err, tt.want) {
				t.Fatalf("Returned %+v, wanted %+v", err, tt.want)
			}
		})
	}
}

func TestGetNextLink(t *testing.T) {
	links := []Link{{
		Rel:    "next",
		Method: "GET",
		HRef:   "https://cisco-test.solvedirect.com/ws/rest/v1/tsp-codes?max_id=494480506&limit=5",
	}}
	want := &ListParams{
		SinceId: nil,
		MaxId:   Int64(494480506),
		Limit:   Int64(5),
	}
	got, err := getNextLink(links)
	if err != nil {
		t.Fatalf("Returned %+v, wanted %+v", err, want)
	}
	if !cmp.Equal(got, want) {
		t.Errorf("getNextLink returned %+v, want %+v", got, want)
	}
}

func TestGetToken(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	mux.HandleFunc("/ws/rest/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		data := `{"access_token":"dummy","scope":"data","token_type":"Bearer","expires_in":3600}`
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, data)
	})
	client.token = nil

	want := token{
		AccessToken: "dummy",
		Scope:       "data",
		TokenType:   "Bearer",
		ExpiresIn:   3600,
		ExpiresAt:   time.Now().Add(time.Duration(3600) * time.Second),
	}

	client.getToken()
	if !cmp.Equal(client.token.AccessToken, want.AccessToken) {
		t.Errorf("GetToken returned %+v, want %+v", client.token.AccessToken, want)
	}
	// check we have at least an hour on the token
	if time.Until(client.token.ExpiresAt) < time.Duration(50*time.Minute) {
		t.Errorf("GetToken returned %+v, want %+v", client.token.AccessToken, want)
	}
}
