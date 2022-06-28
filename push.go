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
			Remarks:    String(remarks),
		},
		CallStates: &CallSystemCodesHolder{
			ShortName: String("Update"),
		},
	}
	return c.PushUpdate(ctx, cd)
}
