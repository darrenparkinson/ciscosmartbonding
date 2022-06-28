package ciscosmartbonding

import (
	"encoding/json"
	"testing"
	"time"
)

type TestCiscoDateTime struct {
	Time CiscoDateTime `json:"time"`
}

func TestCiscoTime(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  time.Time
	}{
		{"cisco time", `{"time": "2019-05-02T18:20:43.000+0000"}`, time.Date(2019, time.May, 2, 18, 20, 43, 0, time.FixedZone("", 0))},
		{"RFC3339Nano time", `{"time": "2016-08-08T21:35:14.000052975+02:00"}`, time.Date(2016, time.August, 8, 21, 35, 14, 52975, time.FixedZone("", 7200))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got TestCiscoDateTime
			err := json.Unmarshal([]byte(tt.input), &got)
			if err != nil {
				t.Fatal(err)
			}
			if !got.Time.Time.Equal(tt.want) {
				t.Fatalf("returned %+v, wanted %+v", got, tt.want)

			}
		})
	}
}

func TestCiscoTimeNull(t *testing.T) {
	var got CiscoDateTime
	err := json.Unmarshal([]byte("null"), &got)
	if err != nil {
		t.Fatalf("returned %+v, wanted %+v", err, nil)
	}
}

func TestCiscoTimeError(t *testing.T) {
	var got CiscoDateTime
	err := json.Unmarshal([]byte(`"2022-01-01"`), &got)
	if err == nil {
		t.Fatalf("returned %+v, wanted error", nil)
	}
}
