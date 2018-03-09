package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// The dependency in this case is command line arguments and printing
	// to stdout.
	// So we will inject the dependency by using a slice of strings and
	// bytes.Buffer{} for printing

	// create a slice of strings
	args := []string{"Filename", "Hello", "World"}

	// create a buffer of bytes
	buffer := bytes.Buffer{}

	// pass it to our function
	Echo(&buffer, args)

	// what we got
	got := buffer.String()

	// what we want
	want := args[1] + " " + args[2] + "\n"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
