package ciscosmartbonding

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPushUpdate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	createRequest := &CallData{}
	// set up the handler
	// we're only testing that the function makes a POST request to Cisco.  Since they don't
	// actually send us any response, there isn't much for us to test here.
	mux.HandleFunc("/rest/v1/push/call", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
	})
	ctx := context.Background()
	got, err := client.PushUpdate(ctx, createRequest)
	if err != nil {
		t.Fatalf("PushUpdate returned error: %v", err)
	}
	want := 200
	if got.StatusCode() != 200 {
		t.Errorf("PushUpdate returned\n%+v\n, want\n%+v", got.StatusCode(), want)
	}
}

func TestUpdateTicketWithWorkNotes(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	inputTicket := "PartnerTicket001"
	inputRemark := "Test Remark"

	// set up the handler
	mux.HandleFunc("/rest/v1/push/call", func(w http.ResponseWriter, r *http.Request) {
		// decode the incoming body
		var got CallData
		json.NewDecoder(r.Body).Decode(&got)

		testMethod(t, r, "POST")
		want := CallData{
			Calls: &InboundCallsHolder{
				CustCallID: String(inputTicket),
				// Remarks:    String(inputRemark),
				Remarks: &StringOrSliceOfErrors{RemarkString: inputRemark},
			},
			CallStates: &CallSystemCodesHolder{
				ShortName: String("Update"),
			},
		}
		// check we got the expected body
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("Request body mismatch (-want +got):\n%s", diff)
		}
	})
	ctx := context.Background()
	got, err := client.UpdateTicketWithWorkNotes(ctx, inputTicket, inputRemark)
	if err != nil {
		t.Fatalf("UpdateTicketWithWorkNotes returned error: %v", err)
	}
	want := 200
	if got.StatusCode() != 200 {
		t.Errorf("UpdateTicketWithWorkNotes returned\n%+v\n, want\n%+v", got.StatusCode(), want)
	}
}
