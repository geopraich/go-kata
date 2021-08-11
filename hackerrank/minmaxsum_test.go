package hackerrank

import "testing"

func TestMinMaxSum(t *testing.T) {
	var testData = []struct {
		in        []int64
		out, out2 int64
	}{
		{[]int64{1, 2, 3, 4, 5}, 10, 14},
		{[]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 36, 44},
	}

	for _, pair := range testData {
		expected, expected2 := pair.out, pair.out2
		actual, actual2 := maxMinSum(pair.in)
		if actual != expected || actual2 != expected2 {
			t.Error("Test Failed with arg:", pair.in, ": Expected: ", expected, expected2,
				": Got: ", actual, actual2)
		}
	}
}

func TestMiniMaxSum(t *testing.T) {
	var testData2 = []struct {
		in        []int
		out, out2 int64
	}{
		{[]int{1, 2, 3, 4, 5}, 10, 14},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 36, 44},
	}

	for _, pair := range testData2 {
		expected, expected2 := pair.out, pair.out2
		actual, actual2 := miniMaxSum(pair.in)
		if actual != expected || actual2 != expected2 {
			t.Error("Test Failed with arg:", pair.in, ": Expected: ", expected, expected2,
				": Got: ", actual, actual2)
		}
	}
}

func TestMiniMaxSum2(t *testing.T) {
	var testData2 = []struct {
		in        []int
		out, out2 int64
	}{
		{[]int{1, 2, 3, 4, 5}, 10, 14},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 36, 44},
	}

	for _, pair := range testData2 {
		expected, expected2 := pair.out, pair.out2
		actual, actual2 := miniMaxSum2(pair.in)
		if actual != expected || actual2 != expected2 {
			t.Error("Test Failed with arg:", pair.in, ": Expected: ", expected, expected2,
				": Got: ", actual, actual2)
		}
	}
}
