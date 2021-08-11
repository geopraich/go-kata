package hackerrank

import (
	"fmt"
	"strconv"
)

func staircase(steps int) {
	char := ""
	for c := steps; c > 0; c-- {
		char += "#"
		fmt.Printf("%"+strconv.Itoa(steps)+"s\n", char)
	}
}

func staircaseTwo(steps int) {

	var char string
	var space string
	lst := make([]string, steps-1)
	lst2 := make([]string, steps)
	lst3 := make([]string, steps)

	for c := 0; c < steps-1; c++ {
		space += " "
		lst[c] = space
	}
	for c := 0; c < steps; c++ {
		char += "#"
		lst2[c] = char
	}
	for c := 0; c < steps-1; c++ {
		lst3[c] = lst[len(lst)-1-c] + lst2[c]
	}

	lst3[len(lst3)-1] = lst2[len(lst2)-1]

	for _, seq := range lst3 {
		fmt.Println(seq)
	}

}
