package main

func Sum(nums []int) int {
	sum := 0
	for _, item := range nums {
		sum += item
	}
	return sum
}

func SumAll(sliceofarrays ...[]int) []int {

	finalSlice := make([]int, len(sliceofarrays))
	for i, item := range sliceofarrays {
		finalSlice[i] = Sum(item)
	}
	return finalSlice
}
