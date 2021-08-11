package hackerrank

import (
	"sort"
)

func birthdayBlowCandles(lst []int) int {
	sort.Ints(lst)
	max := lst[len(lst)-1]
	candles := 1

	for _, number := range lst[:len(lst)-1] {
		if number == max {
			candles += 1
		}
	}
	return candles
}
