package main

import "testing"

// func TestSlice(t *testing.T) {
// 	numbersA := []int{2, 4, 6, 8}
// 	numbersB := []int{2,}

// 	got := Slice(numbersA)
// 	want := 30

// 	if got != want {
// 		t.Errorf("got %d want %d ", got, want)
// 	}
// }
func assertCorrectMessage(t testing.TB, got, want int, input []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d input %v", got, want, input)
	}
}

func TestSlice(t *testing.T) {
	t.Run("Testing 1st group of slices", func(t *testing.T) {
		input := []int{10, 5, 5}
		got := Slice(input)
		want := 20
		assertCorrectMessage(t, got, want, input)
	})

	t.Run("Testing second group of slices", func(t *testing.T) {
		input := []int{2, 5, 5, 5, 6}
		got := Slice(input)
		want := 23
		assertCorrectMessage(t, got, want, input)
	})
}
