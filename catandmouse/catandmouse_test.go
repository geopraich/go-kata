package main

import "testing"

var testData = []struct {
	in []int
	out string
}{
	{[]int{1,2,3},"Cat B"},
	{[]int{1,3,2},"Mouse C"},
	{[]int{2,1,3},"Cat A"},
	{[]int{5,1,2},"Cat B"},
	{[]int{1,6,3},"Cat A"},
	{[]int{4,5,1},"Cat A"},
	{[]int{5,4,1},"Cat B"},
	{[]int{2,3,4},"Cat B"},
	{[]int{3,1,2},"Mouse C"},
	{[]int{5,1,3},"Mouse C"},
	{[]int{1,11,6},"Mouse C"},
	{[]int{0,4,2},"Mouse C"},
	{[]int{12,0,6},"Mouse C"},
	{[]int{1,2,0},"Cat A"},
	{[]int{2,5,3},"Cat A"},
	{[]int{1,4,3},"Cat B"},
	{[]int{22, 75, 70},"Cat B"},
	{[]int{33, 86, 59},"Cat A"},
	{[]int{47, 29, 89}, "Cat A"},
	{[]int{18, 19, 82},"Cat B"},
	{[]int{84, 17, 18},"Cat B"},
	{[]int{100, 11, 55},"Cat B"},
	{[]int{4, 44, 44}, "Cat B"},
	{[]int{16,16,16},"Mouse C"},
	{[]int{16,40,16},"Cat A"},
	{[]int{4,44,45},"Cat B"},
	{[]int{10,10,52},"Mouse C"},
}

func TestCatAndMouse(t *testing.T) {
	for _, pair := range testData {
		expected := pair.out
		actual := mouseAndCat(pair.in)
		if actual != expected {
			t.Error("Test Failed with arg:", pair.in, ": Expected: ", expected,
				": Got: ", actual)
		}
	}
}