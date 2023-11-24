package foo_bar_qix

import "testing"

var fbq = FooBarQix{}

func TestFoo(t *testing.T) {
	expected := "Foo"

	result, err := fbq.compute("6")
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestNotFoo(t *testing.T) {
	expected := "5"

	result, err := fbq.compute("5")
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}
