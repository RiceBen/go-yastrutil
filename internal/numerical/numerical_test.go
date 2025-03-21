package numerical_test

import (
	"testing"

	numericals "github.com/RiceBen/go-yastrutil/internal/numerical"
)

func Test_Zero(t *testing.T) {
	target := 0
	expect := "zero"

	actual := numericals.ItoAlpha(target)

	if actual != expect {
		t.Errorf("Expect %s got %s", expect, actual)
	}
}

func Test_Datasets(t *testing.T) {
	dataSets := map[string]struct {
		input  int
		output string
	}{
		"1": {
			input:  1,
			output: "one",
		},
		"99": {
			input:  99,
			output: "ninety-nine",
		},
		"888": {
			input:  888,
			output: "eight hundred and eighty-eight",
		},
		"7777": {
			input:  7777,
			output: "seven thousand and seven hundred and seventy-seven",
		},
		"66666": {
			input:  66666,
			output: "sixty-six thousand and six hundred and sixty-six",
		},
		"555555": {
			input:  555555,
			output: "five hundred and fifty-five thousand and five hundred and fifty-five",
		},
		"4444444": {
			input:  4444444,
			output: "four million and four hundred and forty-four thousand and four hundred and forty-four",
		},
		"33333333": {
			input:  33333333,
			output: "thirty-three million and three hundred and thirty-three thousand and three hundred and thirty-three",
		},
	}

	for name, test := range dataSets {
		t.Run(name, func(t *testing.T) {
			actual := numericals.ItoAlpha(test.input)

			if actual != test.output {
				t.Errorf("Expect %s got %s", test.output, actual)
			}
		})
	}
}
