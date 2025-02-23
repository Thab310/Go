package slice

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum of 4 numbers", func(t *testing.T) {
		numbers := []int{2, 2, 2, 2}
		got := Sum(numbers)
		want := 8
		assertCorrectSum(t, got, want, numbers)

	})
}

func assertCorrectSum(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v given %v", got, want, numbers)
	}
}

func ExampleSum() {
	numbers := []int{2, 2, 2, 2, 2}
	fmt.Println(Sum(numbers))
	//output: 10
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{2, 8}, []int{2, 2, 6})
	want := []int{10, 10}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {
	assertCorrectSumAll := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("make sums of some slices", func(t *testing.T) {
		got := SumAllTail([]int{2, 2, 4}, []int{4, 2, 2, 6})
		want := []int{6, 10}

		assertCorrectSumAll(t, got, want)
	})

	t.Run("safely sum of empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{2, 2, 6})
		want := []int{}

		assertCorrectSumAll(t, got, want)

	})

}
