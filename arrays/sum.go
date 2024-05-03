package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func SumAll(args ...[]int) []int {
	var sums []int

	for _, arg := range args {
		sums = append(sums, Sum(arg))
	}

	return sums
}
