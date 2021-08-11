package hackerrank

import "testing"

func TestBirthdayBlowCandles(t *testing.T) {
	var testdata = []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3, 3, 4}, 1},
		{[]int{2, 2, 2, 2, 2}, 5},
		{[]int{1, 2, 3, 3}, 2},
		{[]int{0, 34, 100, 100}, 2},
		{[]int{0, 0}, 2},
	}
	for _, pair := range testdata {
		expected := pair.out
		actual := birthdayBlowCandles(pair.in)
		if actual != expected {
			t.Error("Test Failed with arg:", pair.in, ": Expected: ", expected,
				": Got: ", actual)
		}
	}
}
