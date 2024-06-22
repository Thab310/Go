package main

func Sum(numbers [5]int) int {
	sum := 0

	for _, k := range numbers {
		sum += k
	}

	return sum
}
