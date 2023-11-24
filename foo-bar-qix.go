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

	if value%3 == 0 {
		return "Foo", nil
	}

	return input, err
}
