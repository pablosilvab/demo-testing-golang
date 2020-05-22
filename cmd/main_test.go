package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Pablo")
	want := "Hello, Pablo"

	if got != want {
		t.Errorf("Got: %q, Expected: %q", got, want)
	}
}
