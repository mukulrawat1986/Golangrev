package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {

	// We will pass an interface to the Echo function along with a slice of strings to read
	// from which will be the arguments provided to the program
	args := []string{"filename", "Hello", "World"}

	buffer := &bytes.Buffer{}

	Echo(buffer, args)

	got := buffer.String()
	want := "Hello World\n"

	if got != want {
		t.Errorf("got '%#v', want '%#v'", got, want)
	}
}
