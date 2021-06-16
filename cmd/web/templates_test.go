package main

import (
	"testing"
	"time"
)

// func TestHumanData(t *testing.T) {
// 	// Initialize a new time.Time object and pass it to the humanDate func.
// 	tm := time.Date(2021, 12, 16, 10, 0, 0, 0, time.UTC)
// 	hd := humanDate(tm)
// 	if hd != "16 Dec 2021 at 10:00" {
// 		t.Errorf("want %q; got %q", "16 Dec 2021 at 10:00", hd)
// 	}
// }

func TestHumanData(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name.
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC),
			want: "17 Dec 2020 at 10:00",
		},
		{
			name: "Empy",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CST",
			tm:   time.Date(2020, 12, 17, 10, 0, 0, 0, time.FixedZone("CST", 8*60*60)),
			want: "17 Dec 2020 at 02:00",
		},
	}

	// Loop over the test cases.
	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case.
		// The first parameter is the name of the test (which is used to
		// identify the sub-test in any log output) and
		// the second parameter is anonymous function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			if hd != tt.want {
				t.Errorf("want %q; got %q", tt.want, hd)
			}
		})
	}
}
