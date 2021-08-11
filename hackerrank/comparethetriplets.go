package hackerrank

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
