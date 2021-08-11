package hackerrank

import (
	"fmt"
	"math"
)

func catAndMouse(lst []int) string {
	// mouse between cats
	if lst[2] >= lst[0] && lst[2] <= lst[1] {
		if lst[2]-lst[0] == lst[1]-lst[2] {
			return fmt.Sprint("Mouse C")
		} else {
			catA := lst[2] - lst[0]
			catB := lst[1] - lst[2]
			if catA < catB {
				return fmt.Sprint("Cat A")
			} else {
				return fmt.Sprint("Cat B")
			}
		}
	} else if lst[2] <= lst[0] && lst[2] >= lst[1] {
		if lst[2]-lst[1] == lst[0]-lst[2] {
			return fmt.Sprint("Mouse C")
		} else {
			catB := lst[2] - lst[1]
			catA := lst[0] - lst[2]
			if catB < catA {
				return fmt.Sprint("Cat B")
			} else {
				return fmt.Sprint("Cat A")
			}
		}
	}
	// mouse further/closer away/to (from) cats
	if lst[2] > lst[0] && lst[2] > lst[1] {
		if lst[0] > lst[1] {
			return fmt.Sprint("Cat A")
		} else if lst[1] > lst[0] {
			return fmt.Sprint("Cat B")
		} else {
			return fmt.Sprint("Mouse C")
		}
	}
	if lst[2] < lst[0] && lst[2] < lst[1] {
		if lst[0] < lst[1] {
			return fmt.Sprint("Cat A")
		} else if lst[1] < lst[0] {
			return fmt.Sprint("Cat B")
		}
	}
	return fmt.Sprint("Mouse C")
}

func mouseAndCat(lst []int) string {
	catA := math.Abs(float64(lst[2]) - float64(lst[0]))
	catB := math.Abs(float64(lst[2]) - float64(lst[1]))
	if catA == catB {
		return fmt.Sprint("Mouse C")
	}
	if catA < catB {
		return fmt.Sprint("Cat A")
	}
	return fmt.Sprint("Cat B")
}
