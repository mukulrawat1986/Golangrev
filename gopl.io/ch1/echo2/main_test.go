package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// we will mock command line arguments by using a slice of strings.
	// The output will be stored in a bytes buffer

	t.Run("when command line arguments are given", func(t *testing.T) {
		args := []string{"filename", "hello", "world"}

		// create a bytes buffer to store output
		buffer := bytes.Buffer{}

		// pass it to the function
		Echo(&buffer, args)

		got := buffer.String()
		want := args[1] + " " + args[2] + "\n"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("when command line arguments are not given", func(t *testing.T) {
		args := []string{"filename"}

		// create a bytes buffer to store output
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
