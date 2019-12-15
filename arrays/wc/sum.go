package main

func Sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(slices ...[]int) []int {
	slicesLength := len(slices)
	result := make([]int, slicesLength)
	for i := 0; i < slicesLength; i++ {
		result[i] = Sum(slices[i])
	}
	return result
}

func SumAllTails(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		if len(slice) > 0 {
			tail := slice[1:]
			result = append(result, Sum(tail))
		} else {
			result = append(result, 0)
		}
	}
	return result
}
