package ciscosmartbonding

import "time"

// CiscoTime represents the date time with offset values that Cisco returns in their data.
type CiscoDateTime struct {
	time.Time
}

// DateTimeWithOffsetNoColonFormat represents the datetime field provided in the Cisco results.
const DateTimeWithOffsetNoColonFormat = "2006-01-02T15:04:05.999999999Z0700"

func (d *CiscoDateTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	s := string(b)
	t, err := time.Parse(time.RFC3339Nano, s[1:len(s)-1])
	if err != nil {
		t, err = time.Parse(DateTimeWithOffsetNoColonFormat, s[1:len(s)-1])
		if err != nil {
			return err
		}
	}
	d.Time = t
	return nil
}
