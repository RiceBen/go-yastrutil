package numerical

import (
	"strings"
)

var (
	lessThan20 = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens       = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
)

// Convert numerical to alphabet
func ItoAlpha(numerical int) string {
	if numerical == 0 {
		return "zero"
	}

	if numerical < 0 {
		return "minus " + ItoAlpha(-numerical)
	}

	var parts []string

	if numerical >= 1000000000 {
		billion := numerical / 1000000000
		parts = append(parts, ItoAlpha(billion)+" billion")
		numerical %= 1000000000
	}

	if numerical >= 1000000 {
		million := numerical / 1000000
		parts = append(parts, ItoAlpha(million)+" million")
		numerical %= 1000000
	}

	if numerical >= 1000 {
		thousand := numerical / 1000
		parts = append(parts, numberToWords(thousand)+" thousand")
		numerical %= 1000
	}

	if numerical > 0 {
		hundred := numerical / 100
		remainder := numerical % 100
		if hundred > 0 {
			parts = append(parts, lessThan20[hundred]+" hundred")
		}
		if remainder > 0 {
			parts = append(parts, numberToWords(remainder))
		}
	}

	return strings.Join(parts, " and ")
}

func numberToWords(numerical int) string {
	if numerical < 20 {
		return lessThan20[numerical]
	}
	if numerical < 100 {
		ten := numerical / 10
		unit := numerical % 10
		if unit == 0 {
			return tens[ten]
		}
		return tens[ten] + "-" + lessThan20[unit]
	}
	if numerical < 1000 {
		hundred := numerical / 100
		remainder := numerical % 100
		if remainder == 0 {
			return lessThan20[hundred] + " hundred"
		}
		return lessThan20[hundred] + " hundred and " + numberToWords(remainder)
	}

	return ""
}
