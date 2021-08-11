package hackerrank

func applesAndOranges(s, t, Tree int, fruit []int) int {
	fruitTotal := 0
	for _, fruitFall := range fruit {
		if (Tree+fruitFall) >= s && (Tree+fruitFall) <= t {
			fruitTotal++
		}
	}
	return fruitTotal
}
