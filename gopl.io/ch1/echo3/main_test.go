package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// Pass a slice of strings and an interface to handle the dependency of command
	// line arguments and printing
	args := []string{"filename", "Hello", "World"}

	buffer := &bytes.Buffer{}

	Echo(buffer, args)

	got := buffer.String()
	want := "Hello World\n"

	if got != want {
		t.Errorf("got '%#v' want '%#v'", got, want)
	}
}
