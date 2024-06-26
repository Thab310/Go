package main

import (
	"reflect"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want int, input []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d input %v", got, want, input)
	}
}

func TestSlice(t *testing.T) {
	t.Run("Testing slices", func(t *testing.T) {
		input := []int{10, 5, 5}
		got := Sum(input)
		want := 20
		assertCorrectMessage(t, got, want, input)
	})

	t.Run("Test sum of all slices", func(t *testing.T) {
		inputA := []int{2, 4, 6}
		inputB := []int{5, 5}
		inputC := []int{2, 1, 5, 2}

		got := SumAll(inputA, inputB, inputC)
		want := []int{12, 10, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test sum of all tails in collection", func(t *testing.T) {
		inputA := []int{2, 3, 4, 5}
		inputB := []int{3, 5}
		inputC := []int{5, 5, 5}

		got := SumAllTails(inputA, inputB, inputC)
		want := []int{12, 5, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
