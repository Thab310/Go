package main

func Slice(numbers []int) int {
	sum := 0

	for _, k := range numbers {
		sum += k
	}

	return sum
}
