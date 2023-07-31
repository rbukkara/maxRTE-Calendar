package weekday

import (
	"testing"
)

type IO struct {
	input      string
	output     string
	errMessage string
}

func TestCreditCardValidity(t *testing.T) {

	testCases := []IO{
		{
			input:      "09-01-2022",
			output:     "Thursday",
			errMessage: "",
		},
		{
			input:      "07-24-2023",
			output:     "Monday",
			errMessage: "",
		},
		{
			input:      "07-34-2023",
			output:     "",
			errMessage: "invalid date",
		},
		{
			input:      "01-01-2022",
			output:     "Saturday",
			errMessage: "",
		},
	}

	for _, testCase := range testCases {

		got, err := WeekDay(testCase.input)
		if err != nil {
			if err.Error() != testCase.errMessage {
				t.Errorf("%v expected and got: %v", testCase.errMessage, err)
			}
		} else {
			if got != testCase.output {
				t.Errorf("%v expected and got: %v", testCase.output, got)
			}
		}
	}

}
