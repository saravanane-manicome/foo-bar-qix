package foo_bar_qix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type FooBarQix struct{}

var (
	charMapping = map[int32]string{
		'3': "Foo",
		'5': "Bar",
		'7': "Qix",
	}
)

func (fbq FooBarQix) compute(input string) (string, error) {
	value, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return "", errors.New(fmt.Sprintf("invalid input: %s, expected integer", input))
	}

	getTrailingWithZeroAs := zeroReplacerForTrailingOf(input)

	output := ""
	if value%3 == 0 {
		output = output + "Foo"
	}

	if value%5 == 0 {
		output = output + "Bar"
	}

	if value%7 == 0 {
		output = output + "Qix"
	}

	if output == "" {
		return strings.Replace(input, "0", "*", -1) + getTrailingWithZeroAs(""), nil
	}

	return output + getTrailingWithZeroAs("*"), err
}

func zeroReplacerForTrailingOf(input string) func(string) string {
	return func(zeroReplacement string) string {
		trailing := ""
		for _, c := range input {
			if c == '0' {
				trailing += zeroReplacement
			} else {
				trailing += charMapping[c]
			}
		}
		return trailing
	}
}
