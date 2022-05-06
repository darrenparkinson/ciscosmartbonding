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

	mux.HandleFunc("/ws/rest/v1/tsp-codes/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		data := `{"_links": [],"tspCodes": [{
				"changeFlag": "Unchanged",
		    "editTimeUtc": "2020-04-01 00:00:06",
		    "id": 123456789,
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
	got, err := client.GetTSPCodes(ctx, &ListParams{Limit: Int64(1)})
	if err != nil {
		t.Fatalf("GetTSPCodes returned error: %v", err)
	}
	want := &TspCodeList{
		TspCodes: []TspCode{{
			ChangeFlag:             String("Unchanged"),
			EditTimeUtc:            String("2020-04-01 00:00:06"),
			Id:                     Int32(123456789),
			ProblemCodeDescription: String("Product Feature/Function Question"),
			ProblemCodeName:        String("PRODUCT_QUESTION"),
			SubTechId:              Int32(1234),
			SubTechName:            String("Sub Tech Name"),
			TechId:                 Int32(12),
			TechName:               String("Tech Name")}},
		Links: []Link{},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("GetTSPCodes returned %+v, want %+v", got, want)
	}
}
func TestTSPCodes_GetAll(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	//TODO: Test that it makes multiple requests

	mux.HandleFunc("/ws/rest/v1/tsp-codes/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		data1 := `{"_links": [],"tspCodes": [{
				"changeFlag": "Unchanged",
		    "editTimeUtc": "2020-04-01 00:00:06",
		    "id": 123456789,
		    "problemCodeDescription": "Product Feature/Function Question",
		    "problemCodeName": "PRODUCT_QUESTION",
		    "subTechId": 1234,
		    "subTechName": "Sub Tech Name",
		    "techId": 12,
		    "techName": "Tech Name"
			}]}`
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, data1)
	})

	ctx := context.Background()
	got, err := client.GetAllTSPCodes(ctx)
	if err != nil {
		t.Fatalf("GetTSPCodes returned error: %v", err)
	}
	want := []TspCode{{
		ChangeFlag:             String("Unchanged"),
		EditTimeUtc:            String("2020-04-01 00:00:06"),
		Id:                     Int32(123456789),
		ProblemCodeDescription: String("Product Feature/Function Question"),
		ProblemCodeName:        String("PRODUCT_QUESTION"),
		SubTechId:              Int32(1234),
		SubTechName:            String("Sub Tech Name"),
		TechId:                 Int32(12),
		TechName:               String("Tech Name")},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("GetAllTSPCodes returned %+v, want %+v", got, want)
	}
}

func TestTSPCodes_Get_Error(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()
	mux.HandleFunc("/ws/rest/v1/tsp-codes/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		errordata := `{"message":"something bad happened","status":400}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, errordata)
	})
	ctx := context.Background()
	got, err := client.GetTSPCodes(ctx, &ListParams{Limit: Int64(1)})
	if err == nil {
		t.Fatalf("Returned %+v, wanted error", got)
	}
	ce := CiscoError{
		Message: "something bad happened",
		Status:  400,
	}
	want := fmt.Errorf("%w: %s", ErrBadRequest, ce.Message)
	if !errors.Is(err, ErrBadRequest) {
		t.Fatalf("Returned %+v, wanted %+v", err, want)
	}
	if !strings.Contains(err.Error(), ce.Message) {
		t.Fatalf("Returned %+v, wanted %+v", err, want)
	}

}
