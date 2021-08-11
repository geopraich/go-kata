package hackerrank

import "testing"

func TestTimeConvertorGiven12HourTimeInHHMMSSXXFormat(t *testing.T) {
	var testData = []struct {
		in  string
		out string
	}{
		{"12:32:44AM", "00:32:44"},
		{"12:05:33PM", "12:05:33"},
		{"11:59:59PM", "23:59:59"},
		{"11:59:59AM", "11:59:59"},
		{"05:00:00PM", "17:00:00"},
		{"05:00:00AM", "05:00:00"},
	}
	for _, pair := range testData {
		expected := pair.out
		actual := timeConvertor(pair.in)
		if actual != expected {
			t.Error("Test Failed with arg:", pair.in, ": Expected: ", expected,
				": Got: ", actual)
		}
	}
}
