package hackerrank

import "testing"

func TestGradeStudentsRoundingUpToNearestMultipleOfFiveWhenAboveThirtyEight(t *testing.T) {
	var testData = []struct {
		in  []float64
		out []float64
	}{
		{[]float64{8, 28, 38, 45, 47, 48, 99, 88, 52, 58, 104}, []float64{8, 28, 40, 45, 47, 50, 100, 90, 52, 60, 104}},
	}
	for i, pair := range testData {
		expected := pair.out[i]
		actual := gradeStudentsWithoutModulus(pair.in)
		if actual[i] != expected {
			t.Error("Test Failed with arg:", pair.in, ": Expected: ", expected,
				": Got: ", actual)
		}
	}
}

func TestGradeStudentsWithModulus(t *testing.T) {
	var testData2 = []struct {
		in  []int
		out []int
	}{
		{[]int{28, 38, 40, 55, 58, 53, 52, 99, 100, 104}, []int{28, 40, 55, 60, 55, 52, 100, 100, 104}},
	}
	for i, pair := range testData2 {
		expected := pair.out[i]
		actual := gradeStudentsWithModulus(pair.in)
		if actual[i] != expected {
			t.Error("Test Failed with arg:", pair.in, ": Expected: ", expected,
				": Got: ", actual)
		}
	}
}
