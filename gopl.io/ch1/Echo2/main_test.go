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

	// slice of string that works as command line arguments
	args := []string{"Filename", "Hello", "World"}

	// bytes buffer to store the output
	buffer := bytes.Buffer{}

	// call our function
	Echo(&buffer, args)

	// what we get
	got := buffer.String()

	// what we want
	want := args[1] + " " + args[2] + "\n"

	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}