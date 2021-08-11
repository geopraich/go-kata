package hackerrank

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := make([]float64, a)
	for c := 0; c < a; c++ {
		fmt.Scan(&b[c])
	}

	pF, nF, zF := fraction(a, b)
	fmt.Printf("%v\n%v\n%v", pF, nF, zF)
}

func fraction(denom int, s []float64) (float64, float64, float64) {

	var pT, nT, zT float64
	d2 := float64(denom)

	for _, num := range s {
		if num < 0 {
			nT++
		}
		if num > 0 {
			pT++
		}
		if num == 0 {
			zT++
		}
	}

	return pT / d2, nT / d2, zT / d2
}
