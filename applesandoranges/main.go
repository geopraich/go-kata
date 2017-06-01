package main

import "fmt"

func main() {
	var s, t, aTree, oTree, input, input2 int
	fmt.Scanln(&s, &t)
	fmt.Scanln(&aTree, &oTree)
	fmt.Scanln(&input, &input2)
	var apples = make([]int, input)
	var oranges = make([]int, input2)
	for c := 0; c < input; c++ {
		fmt.Scan(&apples[c])
	}
	for c := 0; c < input2; c++ {
		fmt.Scan(&oranges[c])
	}
	fmt.Println(applesAndOranges(s,t,aTree,apples))
	fmt.Println(applesAndOranges(s,t,oTree,oranges))
}

func applesAndOranges(s,t,Tree int, fruit []int) int {
	fruitTotal := 0
	for _, fruitFall := range fruit {
			if (Tree+fruitFall) >= s && (Tree+fruitFall) <= t {
				fruitTotal++
			}
		}
	return fruitTotal
}