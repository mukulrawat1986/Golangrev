package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// we will mock command line arguments by using a slice of strings.
	// The output will be stored in a bytes buffer

	assertEquality := func(t *testing.T, buffer *bytes.Buffer, want string) {
		t.Helper()
		got := buffer.String()

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	t.Run("when command line arguments are given", func(t *testing.T) {
		args := []string{"filename", "hello", "world"}

		// create a bytes buffer to store output
		buffer := bytes.Buffer{}

		// pass it to the function
		Echo(&buffer, args)

		want := args[1] + " " + args[2] + "\n"

		assertEquality(t, &buffer, want)
	})

	t.Run("when command line arguments are not given", func(t *testing.T) {
		args := []string{"filename"}

		// create a bytes buffer to store output
		buffer := bytes.Buffer{}

		// pass it to the function
		Echo(&buffer, args)

		want := "\n"

		assertEquality(t, &buffer, want)
	})
}
