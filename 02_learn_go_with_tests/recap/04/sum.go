package slice

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
