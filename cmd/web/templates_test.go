package main

import (
	"testing"
	"time"

	"snippetbox.owezzy.tech/internal/assert"
)

func TestHumanDate(t *testing.T) {

	tm := time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC)
	hd := humanDate(tm)

	// Check that the output from the humanDate function is in the format we
	// expect. If it isn't what we expect, use the t.Errorf() function to
	// indicate that the test has failed and log the expected and actual
	// values.

	if hd != "17 Mar 2022 at 10:15" {
		t.Errorf("got %q; want %q", hd, "17 Mar 2022 at 10:15")
	}

	// Create a slice of anonymous structs containing the test case name,
	// input to our humanDate() function (the tm field), and expected output
	// (the want field).
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2022 at 10:15",
		}, {
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2022 at 09:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)

		})
	}
}
