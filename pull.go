package ciscosmartbonding

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// PullUpdate makes a request for any ticket updates. It returns a CallData object, and also the
// resty response so that you can check for a 204 No Content in the resp.StatusCode() since that
// is the only way you can tell to stop pulling for data.
func (c *Client) PullUpdate(ctx context.Context) (*CallData, *resty.Response, error) {
	slug := "/rest/v1/pull/call"
	update := CallData{}
	c.checkLimitAndGetToken(ctx)
	resp, err := c.RestyClient.R().
		SetContext(ctx).
		SetResult(&update).
		SetError(&CiscoError{}).
		Post(slug)
	if err != nil {
		return nil, resp, err
	}
	if resp.IsSuccess() {
		return &update, resp, nil
	}
	return nil, resp, parseErrorResponse(resp)
}
