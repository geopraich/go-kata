package main

import "fmt"

func main() {
	// obtaining the inputs from HackerRank.com
	// two lines of input with three integers on each
	alice := make([]int, 3)
	bob := make([]int, 3)

	for c := 0; c < 3; c++{
		fmt.Scan(&alice[c])
	}
	for c := 0; c < 3; c++ {
		fmt.Scan(&bob[c])
	}

	// runs solution function compareTheTriplets
	aScore, bScore := compareTheTriplets(alice, bob)

	// print to stdout the scores for alice bob separated by space
	fmt.Printf("%v %v", aScore, bScore)

}

func compareTheTriplets(player1, player2 []int) (int, int) {
	var aScore, bScore int
	for count := 0; count < len(player1); count++ {
		if player1[count] > player2[count] {
			aScore++
		} else if player1[count] < player2[count] {
			bScore++
		}
	}
	return aScore, bScore
}
