package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Thabelo")
	want := "Hello, Thabelo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
