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
		want := args[1] + " " + args[2] + "\n"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("when no command line arguments are given", func(t *testing.T) {
		args := []string{"filename"}

		// create a buffer of bytes
		buffer := bytes.Buffer{}

		// pass it to the function
		Echo(&buffer, args)

		got := buffer.String()
		want := "\n"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
