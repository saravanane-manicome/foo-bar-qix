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
	expected := "1"

	result, err := fbq.compute("1")
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestBar(t *testing.T) {
	expected := "Bar"

	result, err := fbq.compute("10")
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestNotBar(t *testing.T) {
	expected := "1"

	result, err := fbq.compute("1")
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}
