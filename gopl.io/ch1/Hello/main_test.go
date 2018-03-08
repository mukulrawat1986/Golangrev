// Our test program
package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	// we need to inject  the dependency of printing.
	// We will pass a pointer to a buffer of bytes to our function
	// and the output will be stored there which we can compare.

	t.Run("when we pass an argument to our function", func(t *testing.T) {
		// Initialize an empty struct of bytes.Buffer
		buffer := bytes.Buffer{}

		// pass the buffer to our function
		Hello(&buffer, "Hello World")

		// what was stored in the buffer
		got := buffer.String()

		// what we want
		want := "Hello World\n"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("when we don't pass an argument to our function", func(t *testing.T) {
		// Initiailze an empty struct of bytes.Buffer
		buffer := bytes.Buffer{}

		// pass the buffer to our function
		Hello(&buffer, "")

		// what was stored in the buffer
		got := buffer.String()

		// what we want
		want := "\n"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
