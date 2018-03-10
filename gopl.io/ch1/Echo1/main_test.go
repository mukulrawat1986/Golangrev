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

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}

	t.Run("when command line arguments are given", func(t *testing.T) {
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

		assertCorrectMessage(t, got, want)
	})

	t.Run("when command line arguments are not given", func(t *testing.T) {
		// create a slice of strings
		args := []string{"Filename"}

		// create a buffer of bytes
		buffer := bytes.Buffer{}

		// pass it to our function
		Echo(&buffer, args)

		// what we got
		got := buffer.String()

		// what we want
		want := "\n"

		assertCorrectMessage(t, got, want)
	})
}
