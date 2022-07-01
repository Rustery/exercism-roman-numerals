package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3000 {
		return "", errors.New("wrong input")
	}
	var result string
	d := digits(input)
	f := [4][3]string{{"I", "V", "X"}, {"X", "L", "C"}, {"C", "D", "M"}, {"M", " ", " "}}
	for i, num := range d {
		s := f[len(d)-i-1]
		switch {
		case num <= 3:
			result += strings.Repeat(s[0], num)
		case num == 9:
			result += s[0] + s[2]
		default:
			if 5-num > 0 {
				result += strings.Repeat(s[0], 5-num)
			}
			result += s[1]
			if num-5 > 0 {
				result += strings.Repeat(s[0], num-5)
			}
		}
	}
	return result, nil
}

func digits(i int) (result []int) {
	for i > 0 {
		result = append([]int{i % 10}, result...)
		i /= 10
	}
	return result
}
