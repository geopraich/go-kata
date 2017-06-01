package main

import (
	"testing"
)

var testData = []struct{
	in []int
	out string
}{
	{[]int{0,3,4,2}, "YES"},
	{[]int{0,2,5,3},"NO"},
}

func TestKangarooJumpTogether(t *testing.T) {
	for _, pair := range testData {
		expected := pair.out
		actual := kangarooJumpTogether(pair.in)
		if actual != expected {
			t.Error("Failed with args: ", pair.in, " expected: ", pair.out,
			"Got: ", actual)
		}
	}
}
