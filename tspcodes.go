package ciscosmartbonding

import (
	"context"
)

// TspCode defines model for TspCode.
type TspCode struct {
	ChangeFlag             *string `json:"changeFlag,omitempty"`
	EditTimeUtc            *string `json:"editTimeUtc,omitempty"`
	Id                     *int64  `json:"id,omitempty"`
	ProblemCodeDescription *string `json:"problemCodeDescription,omitempty"`
	ProblemCodeName        *string `json:"problemCodeName,omitempty"`
	SubTechId              *int32  `json:"subTechId,omitempty"`
	SubTechName            *string `json:"subTechName,omitempty"`
	TechId                 *int32  `json:"techId,omitempty"`
	TechName               *string `json:"techName,omitempty"`
}

// Link defines the model for pagination links
// type Link struct {
// 	Rel    string `json:"rel,omitempty"`
// 	Method string `json:"method,omitempty"`
// 	HRef   string `json:"href,omitempty"`
// }

// TspCodeList defines model for TspCodeList.
type TspCodeList struct {
	TspCodes []TspCode `json:"tspCodes,omitempty"`
	// Links    []Link    `json:"_links,omitempty"`
}

// ListParams defines parameters for List.
// type ListParams struct {
// 	// Only retrieve records with ids greater than since_id
// 	SinceId *int64 `url:"since_id,omitempty" json:"since_id,omitempty"`

// 	// Only retrieve records with ids up to max_id
// 	MaxId *int64 `url:"max_id,omitempty" json:"max_id,omitempty"`

// 	// Limit number of returned records. Only nonnegative values allowed. Maximum 10000
// 	Limit *int64 `url:"limit,omitempty" json:"limit,omitempty"`
// }

// String function to enable nice output during testing
// func (l *ListParams) String() string {
// 	return fmt.Sprintf("SinceId:%d MaxId:%d Limit:%d", l.GetSinceId(), l.GetMaxId(), l.GetLimit())
// }

// GetAllTSPCodes is a convenience function to retrieve all TSP Codes
func (c *Client) GetAllTSPCodes(ctx context.Context) ([]TspCode, error) {
	slug := "/tsp/api/v1/xylem/tspcodes"
	var tcl TspCodeList
	err := c.makeListRequest(ctx, slug, &tcl)
	return tcl.TspCodes, err
}
