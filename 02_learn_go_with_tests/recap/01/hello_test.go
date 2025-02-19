/*
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Here we go again :("

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
*/

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mike")
	want := "Here we go again :( Thabelo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
