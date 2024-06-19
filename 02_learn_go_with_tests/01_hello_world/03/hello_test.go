package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Thabelo", "")
		want := "Hello, Thabelo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, gents' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, gents"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Venda", func(t *testing.T) {
		got := Hello("Thabelo", "TshiVenda")
		want := "Hendaa, Thabelo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Tsonga", func(t *testing.T) {
		got := Hello("Thabelo", "Xitsonga")
		want := "Avuxeni, Thabelo"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
