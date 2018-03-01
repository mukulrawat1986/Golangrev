package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// we will mock our command line arguments by using a slice of strings
	// In our mock the fist string is filename, not to be printed.
	args := []string{"filename", "hello", "world"}

	// buffer to hold our output
	buffer := bytes.Buffer{}

	// call the function
	Echo(&buffer, args)

	got := buffer.String()
	want := args[1] + " " + args[2] + "\n"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
