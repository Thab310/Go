package main

import "testing"

func TestSum(t *testing.T) {

	numbers := [5]int{5, 6, 5, 2, 2}
	got := Sum(numbers)
	want := 20

	if want != got {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}
