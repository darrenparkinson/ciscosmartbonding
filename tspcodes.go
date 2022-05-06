package ciscosmartbonding

import (
	"context"
	"fmt"
)

// TspCode defines model for TspCode.
type TspCode struct {
	ChangeFlag             *string `json:"changeFlag,omitempty"`
	EditTimeUtc            *string `json:"editTimeUtc,omitempty"`
	Id                     *int32  `json:"id,omitempty"`
	ProblemCodeDescription *string `json:"problemCodeDescription,omitempty"`
	ProblemCodeName        *string `json:"problemCodeName,omitempty"`
	SubTechId              *int32  `json:"subTechId,omitempty"`
	SubTechName            *string `json:"subTechName,omitempty"`
	TechId                 *int32  `json:"techId,omitempty"`
	TechName               *string `json:"techName,omitempty"`
}

// Link defines the model for pagination links
type Link struct {
	Rel    string `json:"rel,omitempty"`
	Method string `json:"method,omitempty"`
	HRef   string `json:"href,omitempty"`
}

// TspCodeList defines model for TspCodeList.
type TspCodeList struct {
	TspCodes []TspCode `json:"tspCodes,omitempty"`
	Links    []Link    `json:"_links,omitempty"`
}

// ListParams defines parameters for List.
type ListParams struct {
	// Only retrieve records with ids greater than since_id
	SinceId *int64 `url:"since_id,omitempty" json:"since_id,omitempty"`

	// Only retrieve records with ids up to max_id
	MaxId *int64 `url:"max_id,omitempty" json:"max_id,omitempty"`

	// Limit number of returned records. Only nonnegative values allowed. Maximum 10000
	Limit *int64 `url:"limit,omitempty" json:"limit,omitempty"`
}

// String function to enable nice output during testing
func (l *ListParams) String() string {
	return fmt.Sprintf("SinceId:%d MaxId:%d Limit:%d", l.GetSinceId(), l.GetMaxId(), l.GetLimit())
}

// GetAllTSPCodes is a convenience function to retrieve all TSP Codes
func (c *Client) GetAllTSPCodes(ctx context.Context) ([]TspCode, error) {
	slug := "/rest/v1/tsp-codes/"
	codes := []TspCode{}
	lp := &ListParams{
		Limit: Int64(500),
	}
	for {
		var tcl TspCodeList
		err := c.makeListRequest(ctx, slug, &tcl, lp)
		if err != nil {
			return nil, err
		}
		codes = append(codes, tcl.TspCodes...)
		link, err := getNextLink(tcl.Links)
		if err != nil {
			return nil, err
		}
		if link == nil {
			break
		}
		lp = link
	}
	return codes, nil
}

// GetTSPCodes retrieves a list of TSP Codes with pagination options
func (c *Client) GetTSPCodes(ctx context.Context, listParams ...*ListParams) (*TspCodeList, error) {
	slug := "/rest/v1/tsp-codes/"
	var tcl TspCodeList
	var lp *ListParams
	if len(listParams) > 0 {
		lp = listParams[0]
	}
	err := c.makeListRequest(ctx, slug, &tcl, lp)
	return &tcl, err
}
