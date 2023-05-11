package ciscosmartbonding

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// PushUpdate will POST a CallData object to Cisco Smart Bonding.
// It returns the resty response and any errors, but no data is returned.
// You must use PullUpdate to retrieve status updates for any requests.
func (c *Client) PushUpdate(ctx context.Context, callData *CallData) (*resty.Response, error) {
	slug := "/rest/v1/push/call"
	c.checkLimitAndGetToken(ctx)
	resp, err := c.RestyClient.R().
		SetContext(ctx).
		SetBody(callData).
		SetError(&CiscoError{}).
		Post(slug)
	if err != nil {
		return resp, err
	}
	if resp.IsSuccess() {
		return resp, nil
	}
	return resp, parseErrorResponse(resp)
}

// UpdateTicketWithWorkNotes is a helper function to send worknotes to an existing ticket.
func (c *Client) UpdateTicketWithWorkNotes(ctx context.Context, ticketID, remarks string) (*resty.Response, error) {
	cd := &CallData{
		Calls: &InboundCallsHolder{
			CustCallID: String(ticketID),
			// Remarks:    String(remarks),
			Remarks: &StringOrSliceOfErrors{RemarkString: remarks},
		},
		CallStates: &CallSystemCodesHolder{
			ShortName: String("Update"),
		},
	}
	return c.PushUpdate(ctx, cd)
}

// UpdateTicketWithWorkNotes is a helper function to send worknotes to an existing ticket along with an attachment.
// The attachment must be a <10MB base64 encoded string
func (c *Client) UpdateTicketWithWorkNotesAndAttachment(ctx context.Context, ticketID, remarks, filename, content string) (*resty.Response, error) {
	cd := &CallData{
		Calls: &InboundCallsHolder{
			CustCallID: String(ticketID),
			// Remarks:    String(remarks),
			Remarks: &StringOrSliceOfErrors{RemarkString: remarks},
		},
		CallStates: &CallSystemCodesHolder{
			ShortName: String("Update"),
		},
		Attachments: &[]AttachmentsHolder{{
			DataBase64: String(content),
			FileName:   String(filename),
			NR:         Float32(1),
		}},
	}
	return c.PushUpdate(ctx, cd)

}

func (c *Client) ResolveTicketWithWorkNotes(ctx context.Context, ticketID, remarks string) (*resty.Response, error) {
	cd := &CallData{
		Calls: &InboundCallsHolder{
			CustCallID: String(ticketID),
			// Remarks:    String(remarks),
			Remarks: &StringOrSliceOfErrors{RemarkString: remarks},
		},
		CallStates: &CallSystemCodesHolder{
			ShortName: String("Resolved"),
		},
	}
	return c.PushUpdate(ctx, cd)
}

// ClostTicketRequest is a helper type for the CloseTicket helper function to avoid
// having to remember the crazy field names.
type CloseTicketRequest struct {
	ProblemDescription string                     // String(32000) - Free Text - Description
	CustomerSymptom    string                     // String(32000) - Free Text - Diagnosis?
	HardwareProductID  string                     // String(50) - Free Text - Hardware Product ID, e.g. "495266". (If unknown, default to 1866083) - PIN
	SoftwareProductID  string                     // String(50) - Free Text - Software Product ID, e.g. "888410" (If unknown, default to 3074572) - Department
	ResolutionSummary  string                     // String(50) - Free Text - Resolution Summary, e.g. "faulty power socket changed" - LocationStreet
	Complexity         CloseTicketComplexity      // String(50) - From List - Complexity, e.g. "1 Level -Basic" - LocationCity
	ResolutionCode     CloseTicketResolutionCode  // String(50) - From List - Resolution Code, e.g. "Customer Education" - Fax
	UnderlyingCause    CloseTicketUnderlyingCause // String(50) - From List - Underlying Cause, e.g. "Licensing" - EMail
}

// CloseTicket is a helper function so you don't have to remember the crazy field names for the values
func (c *Client) CloseTicket(ctx context.Context, ticketID string, ctr CloseTicketRequest) (*resty.Response, error) {
	cd := &CallData{
		Calls: &InboundCallsHolder{
			CustCallID:  String(ticketID),
			Description: String(ctr.ProblemDescription),
			// Solution: "",
			// Diagnosis: "", // Free Text String 32000
			CCP: &PersonsHolder{
				PIN:            String(ctr.HardwareProductID),        // String(50) - Free Text - Hardware Product ID, e.g. "495266". (If unknown, default to 1866083)
				Department:     String(ctr.SoftwareProductID),        // String(50) - Free Text - Software Product ID, e.g. "888410" (If unknown, default to 3074572)
				LocationStreet: String(ctr.ResolutionSummary),        // String(50) - Free Text - Resolution Summary, e.g. "faulty power socket changed"
				LocationCity:   String(ctr.Complexity.String()),      // String(50) - From List - Complexity, e.g. "1 Level -Basic"
				Fax:            String(ctr.ResolutionCode.String()),  // String(50) - From List - Resolution Code, e.g. "Customer Education"
				EMail:          String(ctr.UnderlyingCause.String()), // String(50) - From List - Underlying Cause, e.g. "Licensing"
				// Room:           String(), // String(50) - Free Text - Subscription ID, e.g. "S123456789"
			},
		},
		CallStates: &CallSystemCodesHolder{
			ShortName: String("Closed"),
		},
	}
	return c.PushUpdate(ctx, cd)
}
