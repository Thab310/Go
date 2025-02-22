package slice

import (
	"fmt"
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
