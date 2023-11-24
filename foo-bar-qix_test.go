package foo_bar_qix

import (
	"fmt"
	"testing"
)

var fbq = FooBarQix{}

func TestFooBarQix(t *testing.T) {
	parameters := []struct{ input, expected string }{
		{"1", "1"},
		{"3", "FooFoo"}, {"6", "Foo"}, {"9", "Foo"},
		{"5", "BarBar"}, {"10", "Bar*"}, {"20", "Bar*"},
		{"7", "QixQix"}, {"14", "Qix"}, {"28", "Qix"},
		{"15", "FooBarBar"}, {"30", "FooBarFoo*"},
		{"21", "FooQix"}, {"42", "FooQix"},
		{"35", "BarQixFooBar"}, {"70", "BarQixQix*"},
		{"105", "FooBarQix*Bar"},
		{"101", "1*1"}, {"10101", "FooQix**"}, {"303", "FooFoo*Foo"},
	}

	for _, p := range parameters {
		t.Run(fmt.Sprintf("%s => %s", p.input, p.expected), func(t *testing.T) {
			actual, err := fbq.compute(p.input)
			if err != nil {
				t.Fatal(err)
			}

			if actual != p.expected {
				t.Logf("given: %s, expected: %s, got: %s", p.input, p.expected, actual)
				t.Fail()
			}
		})
	}
}
