package main

import (
	"sort"
	"fmt"
)

func main() {
	var input int
	fmt.Scan(&input)
	input2 := make([]int, input)
	for c := 0; c < input; c++ {
		fmt.Scan(&input2[c])
	}

	fmt.Println(birthdayBlowCandles(input2))
}

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
