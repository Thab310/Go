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

////////////////////////////////////////////////////////
/*
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Thabelo")
	want := "Hello Thabelo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
*/
////////////////////////////////////////////////////////
/*
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello people", func(t *testing.T) {
		got := Hello("Thabelo")
		want := "Hello Thabelo"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Hello world when string is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
*/
////////////////////////////////////////////////////////

package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Hello people", func(t *testing.T) {
		got := Hello("Thabelo", "Venda")
		want := "Ndaa Thabelo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello world if string is empty", func(t *testing.T) {
		got := Hello("", "Venda")
		want := "Ndaa world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
