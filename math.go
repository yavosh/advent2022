package advent2022

func SumInts(ints []int) int {
	acc := 0

	for _, v := range ints {
		acc += v
	}

	return acc
}
