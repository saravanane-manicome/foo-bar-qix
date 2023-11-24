package foo_bar_qix

import (
	"errors"
	"fmt"
	"strconv"
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

	trailing := getTrailing(input)

	multipleOf3 := value%3 == 0
	multipleOf5 := value%5 == 0
	multipleOf7 := value%7 == 0

	if !multipleOf3 && !multipleOf5 && !multipleOf7 {
		return input + trailing, nil
	}

	output := ""
	if multipleOf3 {
		output = output + "Foo"
	}

	if multipleOf5 {
		output = output + "Bar"
	}

	if multipleOf7 {
		output = output + "Qix"
	}

	return output + trailing, err
}

func getTrailing(input string) string {
	trailing := ""
	for _, c := range input {
		trailing += charMapping[c]
	}
	return trailing
}
