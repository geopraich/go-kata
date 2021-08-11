package hackerrank

import (
	"sort"
)

func maxMinSum(lst []int64) (int64, int64) {
	var total int64
	for _, number := range lst {
		total += number
	}

	var min int64 = total - lst[0]
	var max int64 = total - lst[0]

	for count := 1; count < len(lst); count++ {
		temp := total - lst[count]
		if temp > max {
			max = temp
		}
		if temp < min {
			min = temp
		}
	}
	return min, max
}

func miniMaxSum(lst []int) (int64, int64) {
	var min int64
	var max int64

	sort.Ints(lst)

	for _, number := range lst[1:] {
		max += int64(number)
	}
	for _, number := range lst[0 : len(lst)-1] {
		min += int64(number)
	}

	return min, max
}

func miniMaxSum2(lst []int) (int64, int64) {
	var max int64
	var min int64

	sort.Ints(lst)

	for _, number := range lst {
		max += int64(number)
	}

	min = max - int64(lst[len(lst)-1])
	max = max - int64(lst[0])

	return min, max
}
