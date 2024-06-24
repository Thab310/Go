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

func main() {
	a := []int{2, 3, 4}
	b := []int{5, 5}
	c := []int{1, 1, 1, 1}

	SumAll(a)
	SumAll(a, b)
	SumAll(a, b, c, c)
}
