package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// create a buffer of bytes.Buffer which we will inject
	// as the dependency for printing.
	buffer := bytes.Buffer{}

	Greet(&buffer, "Hello World")

	got := buffer.String()
	want := "Hello World\n"
	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}
