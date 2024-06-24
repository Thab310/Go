package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want int, input []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d input %v", got, want, input)
	}
}

func TestSlice(t *testing.T) {
	t.Run("Testing slices", func(t *testing.T) {
		input := []int{10, 5, 5}
		got := Slice(input)
		want := 20
		assertCorrectMessage(t, got, want, input)
	})
}
