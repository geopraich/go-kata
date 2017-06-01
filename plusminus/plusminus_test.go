package main

import "testing"

var testData = []struct {
	in int
	in2 []float64
	out, out2, out3 float64
}{
	{6,[]float64{1,0,-3,-4,5,4},3.0/6.0,2.0/6.0,1.0/6.0},
}

func TestPlusMinusGivenTestData(t *testing.T) {
	for _, pair := range testData {
		expected, expected2, expected3 := pair.out, pair.out2, pair.out3
		actual, actual2, actual3 := fraction(pair.in, pair.in2)
		if actual != expected || actual2 != expected2 || actual3 != expected3 {
			t.Error("Test Failed with arg:", pair.in, pair.in2, ": Expected: ", expected, expected2,
				expected3, ": Got: ", actual, actual2, actual3)
		}
	}
}