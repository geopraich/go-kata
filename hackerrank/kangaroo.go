package hackerrank

// brute force
func kangarooJumpTogether(jumps []int) string {
	var lst = make([]int, 10000)
	var lst2 = make([]int, 10000)
	pos2 := jumps[3]
	c2 := 0
	c3 := jumps[2]
	for c := jumps[0]; c2 < len(lst); c += jumps[1] {
		lst[c2] = c
		c2++
	}
	for c := 0; c < len(lst); c++ {
		lst2[c] = c3
		if lst[c] == lst2[c] {
			return "YES"
		}
		c3 += pos2
	}
	return "NO"
}
