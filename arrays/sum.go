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

func SumAllTails(args ...[]int) []int {
	var sums []int

	for _, arg := range args {
		if len(arg) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := arg[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
