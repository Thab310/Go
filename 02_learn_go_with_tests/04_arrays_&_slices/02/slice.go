package main

func Sum(numbers []int) int {
	sum := 0

	for _, k := range numbers {
		sum += k
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	var sums []int
	for _, k := range slices {
		sums = append(sums, Sum(k))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int
	for _, k := range slices {
		if len(k) == 0 {
			sums = append(sums, 0)
		} else {
			tail := k[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
