package ciscosmartbonding

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTSPCodes_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/tsp/api/v1/xylem/tspcodes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		data := `{"tspCodes": [{
				"changeFlag": "Unchanged",
		    "editTimeUtc": "2020-04-01 00:00:06",
		    "id": 10000000002,
		    "problemCodeDescription": "Product Feature/Function Question",
		    "problemCodeName": "PRODUCT_QUESTION",
		    "subTechId": 1234,
		    "subTechName": "Sub Tech Name",
		    "techId": 12,
		    "techName": "Tech Name"
			}]}`
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, data)
	})

	ctx := context.Background()
	got, err := client.GetAllTSPCodes(ctx)
	if err != nil {
		t.Fatalf("GetTSPCodes returned error: %v", err)
	}
	want := []TspCode{{
		ChangeFlag:             String("Unchanged"),
		EditTimeUtc:            String("2020-04-01 00:00:06"),
		Id:                     Int64(10000000002),
		ProblemCodeDescription: String("Product Feature/Function Question"),
		ProblemCodeName:        String("PRODUCT_QUESTION"),
		SubTechId:              Int32(1234),
		SubTechName:            String("Sub Tech Name"),
		TechId:                 Int32(12),
		TechName:               String("Tech Name")}}
	if !cmp.Equal(got, want) {
		t.Errorf("GetTSPCodes returned %+v, want %+v", got, want)
	}
}

func TestTSPCodes_Get_Error(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	mux.HandleFunc("/tsp/api/v1/xylem/tspcodes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		errordata := `{"message":"something bad happened","status":"400"}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errordata)
	})
	ctx := context.Background()
	got, err := client.GetAllTSPCodes(ctx)
	if err == nil {
		t.Fatalf("Returned %+v, wanted error", got)
	}
	ce := CiscoError{
		Message: "something bad happened",
		Status:  "400",
	}
	want := fmt.Errorf("%w: %s", ErrBadRequest, ce.Message)
	if !errors.Is(err, ErrBadRequest) {
		t.Fatalf("Returned %+v, wanted %+v", err, want)
	}
	if !strings.Contains(err.Error(), ce.Message) {
		t.Fatalf("Returned %+v, wanted %+v", err, want)
	}

}
