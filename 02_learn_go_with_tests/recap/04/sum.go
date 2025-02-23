package slice

import "slices"

func Sum(n []int) int {
	sum := 0

	for _, number := range n {
		sum += number
	}
	return sum
}

func SumAll(s ...[]int) []int {

	var sums []int

	for _, slice := range s {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTail(s ...[]int) []int {
	var sums []int

	for _, slice := range s {
		if slices.Equal(slice, nil) {
			return []int{}
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
