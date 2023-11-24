package foo_bar_qix

import (
	"errors"
	"fmt"
	"strconv"
)

type FooBarQix struct{}

func (fbq FooBarQix) compute(input string) (string, error) {
	value, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return "", errors.New(fmt.Sprintf("invalid input: %s, expected integer", input))
	}

	multipleOf3 := value%3 == 0
	multipleOf5 := value%5 == 0

	if !multipleOf3 && !multipleOf5 {
		return input, nil
	}

	output := ""
	if multipleOf3 {
		output = output + "Foo"
	}

	if multipleOf5 {
		output = output + "Bar"
	}

	return output, err
}
