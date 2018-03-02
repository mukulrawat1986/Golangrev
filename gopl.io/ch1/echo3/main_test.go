package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// we will mock the dependency of command line arguments using a slice of strings.
	// We will use bytes buffer to mock out printing.

	t.Run("when command line arguments are given", func(t *testing.T) {
		args := []string{"filename", "hello", "world"}

		// create a buffer of bytes
		buffer := bytes.Buffer{}

		// pass it to the function
		Echo(&buffer, args)

		got := buffer.String()
		want := args[0] + " " + args[1] + "\n"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
