package main

func Sum(s []int) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}
