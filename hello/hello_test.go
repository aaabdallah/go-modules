package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."

	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, but want %q", got, want)
	}
}
