package hackerrank

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConvertor(s string) string {

	lst := strings.Split(s, ":")
	b := lst[2][2:4]
	lst[2] = lst[2][:2]

	if lst[0] == "12" && b == "AM" {
		lst[0] = "00"
	} else if lst[0] != "12" && b == "PM" {
		x, _ := strconv.Atoi(lst[0])
		lst[0] = strconv.Itoa(x + 12)
	}

	return fmt.Sprintf("%v:%v:%v", lst[0], lst[1], lst[2])

}
