package main

import "testing"

var testData = []struct {

	in, in2 []int
	out, out2 int
}{
	{[]int{1,2,3},[]int{2,3,4},0,3},
	{[]int{0,1,5},[]int{1,0,5},1,1},
	{[]int{4,4,4},[]int{1,1,1},3,0},
}

func TestCompareTheTriplets(t *testing.T) {
	for _, pair := range testData {
		var expected, expected2 = pair.out, pair.out2
		var actual, actual2 = compareTheTriplets(pair.in, pair.in2)
		if actual != expected || actual2 != expected2 {
			t.Error("Test Failed with arg:", pair.in, pair.in2, ": Expected: ", expected, expected2,
			": Got: ", actual, actual2)
		}
	}
}