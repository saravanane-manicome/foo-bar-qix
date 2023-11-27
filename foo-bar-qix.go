package foo_bar_qix

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type FooBarQix struct{}

var (
	labelByDivider = map[int64]string{
		3: "Foo",
		5: "Bar",
		7: "Qix",
	}
)

func (fbq FooBarQix) compute(input string) (string, error) {
	value, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return "", errors.New(fmt.Sprintf("invalid input: %s, expected integer", input))
	}

	trailingWithZeroAs := trailingWithZeroMappingForInput(input)

	output := parseMultiples(value)

	if output == "" {
		return strings.ReplaceAll(input, "0", "*") + trailingWithZeroAs(""), nil
	}

	return output + trailingWithZeroAs("*"), nil
}

func parseMultiples(input int64) string {
	keys := sortedDividers()

	output := ""
	for _, divider := range keys {
		if input%divider == 0 {
			output += labelByDivider[divider]
		}
	}
	return output
}

func sortedDividers() []int64 {
	var keys []int64
	for k := range labelByDivider {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func trailingWithZeroMappingForInput(input string) func(string) string {
	trailing := ""
	for _, c := range input {
		if c == '0' {
			trailing += "0"
		} else {
			charAsInt := int64(c - '0')
			trailing += labelByDivider[charAsInt]
		}
	}

	return func(zeroReplacement string) string {
		return strings.ReplaceAll(trailing, "0", zeroReplacement)
	}
}
