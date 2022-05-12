package ciscosmartbonding

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPullUpdate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	// load in the test data
	filename := filepath.Join("testdata", "pull-message-example.json")
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal("error reading test data file", filename, err)
	}
	updateResponse := &CallData{}
	err = json.NewDecoder(bytes.NewReader(data)).Decode(&updateResponse)
	if err != nil {
		t.Fatalf("error decoding test data file %s: %v", filename, err)
	}
	// set up our handler
	mux.HandleFunc("/ws/rest/v1/pull/call", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(data))
	})
	// make the request
	ctx := context.Background()
	got, _, err := client.PullUpdate(ctx)
	if err != nil {
		t.Fatalf("PullUpdate returned error: %v", err)
	}
	want := updateResponse
	if !cmp.Equal(*got, *want) {
		t.Errorf("PullUpdate returned\n%+v\n, want\n%+v", got, want)
	}
}
