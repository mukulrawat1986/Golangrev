package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// we will mock the dependency of command line arguments using a slice of strings.
	// We will use bytes buffer to mock out printing.

	assertEquality := func(t *testing.T, buffer *bytes.Buffer, want string) {
		t.Helper()
		got := buffer.String()

		if got != want {
			t.Errorf("got '%#v', want '%#v'", got, want)
		}
	}

	t.Run("when command line arguments are given", func(t *testing.T) {
		args := []string{"filename", "hello", "world"}

		// create a buffer of bytes
		buffer := bytes.Buffer{}

		// pass it to the function
		Echo(&buffer, args)

		want := args[1] + " " + args[2] + "\n"

		assertEquality(t, &buffer, want)
	})

	t.Run("when no command line arguments are given", func(t *testing.T) {
		args := []string{"filename"}

		// create a buffer of bytes
		buffer := bytes.Buffer{}

		// pass it to the function
		Echo(&buffer, args)

		want := "\n"

		assertEquality(t, &buffer, want)
	})
}
