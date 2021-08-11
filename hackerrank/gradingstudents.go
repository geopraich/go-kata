package hackerrank

func gradeStudentsWithoutModulus(lst []float64) []float64 {
	var tensNum float64
	var unitNum float64
	for index, grade := range lst {
		if grade >= 38 && grade < 100 {
			for _, unit := range []float64{1, 2, 3, 4, 5, 6, 7, 8, 9} {
				if (grade/10)-unit > 0 && (grade/10)-unit < 1 {
					tensNum = unit * 10
					unitNum = grade - (unit * 10)
				} else if (grade/10)-unit == 0 {
					tensNum = unit * 10
					unitNum = 0
				}
			}

			if unitNum > 2 && unitNum < 5 {
				lst[index] = tensNum + 5
			} else if unitNum > 7 {
				lst[index] = tensNum + 10
			}
		}
	}
	return lst
}

func gradeStudentsWithModulus(lst []int) []int {
	var modResult int
	for index, grade := range lst {
		modResult = grade % 5
		if grade >= 38 && grade < 100 {
			if modResult > 2 {
				lst[index] = (grade + 5) - modResult
			}
		}
	}
	return lst
}
