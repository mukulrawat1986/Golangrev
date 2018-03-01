package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// we will mock our command line arguments by using a slice of strings
	// In our mock the fist string is filename, not to be printed.

	// Helper function to assertEquality
	assertEquality := func(t *testing.T, buffer *bytes.Buffer, want string) {
		t.Helper()
		got := buffer.String()

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("test with some arguments", func(t *testing.T) {
		args := []string{"filename", "hello", "world"}

		// buffer to hold our output
		buffer := bytes.Buffer{}

		// call the function
		Echo(&buffer, args)

		want := args[1] + " " + args[2] + "\n"

		assertEquality(t, &buffer, want)
	})

	t.Run("test with no arguments", func(t *testing.T) {
		args := []string{"filename"}

		// buffer to hold our output
		buffer := bytes.Buffer{}

		// call the function
		Echo(&buffer, args)

		want := "\n"

		assertEquality(t, &buffer, want)
	})
}
