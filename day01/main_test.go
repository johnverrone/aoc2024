package main

import "testing"

func TestBar(t *testing.T) {
	want := "Bar"
	if got := Bar(); got != want {
		t.Errorf("Bar() = %q, want %q", got, want)
	}
}
