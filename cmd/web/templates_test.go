package main

import (
	"testing"
	"time"
)

func TestHumanData(t *testing.T) {
	// Initialize a new time.Time object and pass it to the humanDate func.
	tm := time.Date(2021, 12, 16, 10, 0, 0, 0, time.UTC)
	hd := humanDate(tm)
	if hd != "16 Dec 2021 at 10:00" {
		t.Errorf("want %q; got %q", "16 Dec 2021 at 10:00", hd)
	}
}
